package getemails

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/stretchr/testify/assert"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	//read secure Test Properties, to store them outside any GitHub
	//feel free to adjust Path and enter your Keys. Tokens and Secrets
	//File Content Should look like follows:
	// server=<your IMAP Server>
	// username=<your username>
	// password=<your secret>

	props, err := ReadPropertiesFile("c:\\GODev\\ImapApp.properties")
	gprops = props
	if err != nil {
		panic("Error while reading properties file")
	}

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup test attrs
	port := 993
	inbox := "INBOX"

	tc.SetInput("server", gprops["server"])
	tc.SetInput("port", port)
	tc.SetInput("username", gprops["username"])
	tc.SetInput("password", gprops["password"])
	tc.SetInput("mailbox", inbox)

	act.Eval(tc)

	//check result attr
	result := tc.GetOutput("flag")
	assert.Equal(t, result, "none")
}

//Helper Functions
// read Security Settings from external Propery File
//

type ConfigProperties map[string]string

var gprops ConfigProperties

func ReadPropertiesFile(filepath string) (ConfigProperties, error) {
	config := ConfigProperties{}

	if len(filepath) == 0 {
		return config, nil
	}
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line := scan.Text()
		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				config[key] = value
			}
		}
	}

	if err := scan.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return config, nil
}
