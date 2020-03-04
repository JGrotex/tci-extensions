package getemails

import (
	"log"
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
	mbox, err := c.Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Flags for INBOX:", mbox.Flags)

	// Get the last 4 messages
	from := uint32(1)
	to := mbox.Messages
	if mbox.Messages > 3 {
		// We're using unsigned integers here, only substract if the result is > 0
		from = mbox.Messages - 3
	}
	seqset := new(imap.SeqSet)
	seqset.AddRange(from, to)

	messages := make(chan *imap.Message, 10)
	idone = make(chan error, 1)
	go func() {
		idone <- c.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope}, messages)
	}()

	var subject = ""
	log.Println("Last 4 messages:")
	for msg := range messages {
		log.Println("* " + msg.Envelope.Subject)
		subject = msg.Envelope.Subject
	}

	if err := <-idone; err != nil {
		log.Fatal(err)
	}

	log.Println("Done!")

	context.SetOutput("from", "Successfully connected")
	context.SetOutput("subject", subject)
	context.SetOutput("body", "")
	return true, nil
}
