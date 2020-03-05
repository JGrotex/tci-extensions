package getemails

import (
	"io/ioutil"
	"log"
	"net/mail"
	"strconv"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	imap "github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	//Input declaration
	server := context.GetInput("server").(string)
	port := context.GetInput("port").(int)
	username := context.GetInput("username").(string)
	password := context.GetInput("password").(string)
	mailbox := context.GetInput("mailbox").(string)

	flag := ""

	log.Println("Connecting to server ... " + mailbox)

	// Connect to server
	c, err := client.DialTLS(server+":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")

	// Don't forget to logout
	defer c.Logout()

	// Login
	if err := c.Login(username, password); err != nil {
		log.Fatal(err)
	}
	log.Println("Logged in")

	// List mailboxes
	mailboxes := make(chan *imap.MailboxInfo, 10)
	idone := make(chan error, 1)
	go func() {
		idone <- c.List("", "*", mailboxes)
	}()

	log.Println("Mailboxes:")
	for m := range mailboxes {
		log.Println("* " + m.Name)
	}

	if err := <-idone; err != nil {
		log.Fatal(err)
	}

	// Select INBOX
	mbox, err := c.Select(mailbox, false)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Flags for "+mailbox+":", mbox.Flags)

	// Get the last message
	if mbox.Messages == 0 {
		flag = "none"
	} else {
		seqset := new(imap.SeqSet)
		seqset.AddRange(mbox.Messages, mbox.Messages)

		// Get the whole message body
		section := &imap.BodySectionName{}
		items := []imap.FetchItem{section.FetchItem()}

		messages := make(chan *imap.Message, 1)
		idone = make(chan error, 1)
		go func() {
			idone <- c.Fetch(seqset, items, messages)
		}()

		log.Println("Last message:")
		msg := <-messages
		r := msg.GetBody(section)
		if r == nil {
			context.SetOutput("flag", "error")
			log.Fatal("Server didn't returned message body")
		}

		if err := <-idone; err != nil {
			log.Fatal(err)
		}

		m, err := mail.ReadMessage(r)
		if err != nil {
			log.Fatal(err)
		}

		header := m.Header
		log.Println("Date:", header.Get("Date"))
		log.Println("From:", header.Get("From"))
		log.Println("To:", header.Get("To"))
		log.Println("Subject:", header.Get("Subject"))

		body, err := ioutil.ReadAll(m.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Body:", string(body))
		log.Println("Email read Done!")
		flag = "new"

		// First mark the message as deleted
		item := imap.FormatFlagsOp(imap.AddFlags, true)
		flags := []interface{}{imap.DeletedFlag}
		if err := c.Store(seqset, item, flags, nil); err != nil {
			log.Fatal(err)
		}

		// Then delete it
		if err := c.Expunge(nil); err != nil {
			log.Fatal(err)
		}
		log.Println("Last message has been deleted")

		context.SetOutput("date", header.Get("Date"))
		context.SetOutput("from", header.Get("From"))
		context.SetOutput("to", header.Get("To"))
		context.SetOutput("subject", header.Get("Subject"))
		context.SetOutput("body", string(body))
	}
	context.SetOutput("flag", flag)
	return true, nil
}
