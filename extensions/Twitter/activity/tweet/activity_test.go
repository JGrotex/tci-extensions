package tweet

import (
	"bufio"
	"fmt"
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
	//  consumerKey=<your consumerKey>
	//  consumerSecret=<your consumerSecret>
	//  accessToken=<your accessToken>
	//  accessTokenSecret=<your accessTokenSecret>

	props, err := ReadPropertiesFile("c:\\GODev\\twitterApp.properties")
	gprops = props
	if err != nil {
		panic("Error while reading properties file")
	}

	//if gprops["consumerKey"] == "" || gprops["consumerSecret"] == "" || gprops["accessToken"] == "" || gprops["accessTokenSecret"] == "" {
	//	panic("Error properties not loaded correctly")
	//}
	//fmt.Print("... using ..." + gprops["consumerKey"] + " :: " + gprops["accessToken"])

	//read Flogo Metadata
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

func TestTwitterFunction_Tweet(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(act.Metadata())

	//setup attributes
	tc.SetInput("consumerKey", gprops["consumerKey"])
	tc.SetInput("consumerSecret", gprops["consumerSecret"])
	tc.SetInput("accessToken", gprops["accessToken"])
	tc.SetInput("accessTokenSecret", gprops["accessTokenSecret"])
	tc.SetInput("function", "Tweet")
	//tc.SetInput("user", "<no id needed>")
	tc.SetInput("text", "testing a new TIBCO Cloud Integration Twitter Activity")

	act.Eval(tc)

	//check result attr
	code := tc.GetOutput("statusCode")
	msg := tc.GetOutput("message")
	fmt.Print(msg)
	assert.Equal(t, 200, code)
}

func TestTwitterFunction_TweetMedia(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(act.Metadata())

	//setup attributes
	tc.SetInput("consumerKey", gprops["consumerKey"])
	tc.SetInput("consumerSecret", gprops["consumerSecret"])
	tc.SetInput("accessToken", gprops["accessToken"])
	tc.SetInput("accessTokenSecret", gprops["accessTokenSecret"])
	tc.SetInput("function", "TweetMedia")
	//tc.SetInput("user", "<no id needed>")
	tc.SetInput("text", "testing a new TIBCO Cloud Integration Twitter Activity")
	tc.SetInput("mediaURL", "http://www.godev.de/GOLib/assets/img/NUC%20Skull%20Canyon%20(NUC6i7KYK).jpg")

	act.Eval(tc)

	//check result attr
	code := tc.GetOutput("statusCode")
	msg := tc.GetOutput("message")
	fmt.Print(msg)
	assert.Equal(t, 200, code)
}

func TestTwitterFunction_ReTweet(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(act.Metadata())

	//setup attributes
	tc.SetInput("consumerKey", gprops["consumerKey"])
	tc.SetInput("consumerSecret", gprops["consumerSecret"])
	tc.SetInput("accessToken", gprops["accessToken"])
	tc.SetInput("accessTokenSecret", gprops["accessTokenSecret"])
	tc.SetInput("function", "ReTweet")
	tc.SetInput("user", "<tweet id>")
	tc.SetInput("text", "Retweeting the tweet")

	act.Eval(tc)

	//check result attr
	code := tc.GetOutput("statusCode")
	msg := tc.GetOutput("message")
	fmt.Print(msg)
	assert.Equal(t, 200, code)
}

func TestTwitterFunction_Block(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(act.Metadata())

	//setup attributes
	tc.SetInput("consumerKey", gprops["consumerKey"])
	tc.SetInput("consumerSecret", gprops["consumerSecret"])
	tc.SetInput("accessToken", gprops["accessToken"])
	tc.SetInput("accessTokenSecret", gprops["accessTokenSecret"])
	tc.SetInput("function", "Block")
	tc.SetInput("user", "@IBM")
	//tc.SetInput("text", "")

	act.Eval(tc)

	//check result attr
	code := tc.GetOutput("statusCode")
	msg := tc.GetOutput("message")
	fmt.Print(msg)
	assert.Equal(t, 200, code)
}

func TestTwitterFunction_Unblock(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(act.Metadata())

	//setup attributes
	tc.SetInput("consumerKey", gprops["consumerKey"])
	tc.SetInput("consumerSecret", gprops["consumerSecret"])
	tc.SetInput("accessToken", gprops["accessToken"])
	tc.SetInput("accessTokenSecret", gprops["accessTokenSecret"])
	tc.SetInput("function", "Unblock")
	tc.SetInput("user", "@IBM")
	//tc.SetInput("text", "")

	act.Eval(tc)

	//check result attr
	code := tc.GetOutput("statusCode")
	msg := tc.GetOutput("message")
	fmt.Print(msg)
	assert.Equal(t, 200, code)
}

func TestTwitterFunction_Follow(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(act.Metadata())

	//setup attributes

	tc.SetInput("consumerKey", gprops["consumerKey"])
	tc.SetInput("consumerSecret", gprops["consumerSecret"])
	tc.SetInput("accessToken", gprops["accessToken"])
	tc.SetInput("accessTokenSecret", gprops["accessTokenSecret"])
	tc.SetInput("function", "Follow")
	tc.SetInput("user", "@Google")
	//tc.SetInput("text", "")

	act.Eval(tc)

	//check result attr
	code := tc.GetOutput("statusCode")
	msg := tc.GetOutput("message")
	fmt.Print(msg)
	assert.Equal(t, 200, code)
}

func TestTwitterFunction_Unfollow(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(act.Metadata())

	//setup attributes
	tc.SetInput("consumerKey", gprops["consumerKey"])
	tc.SetInput("consumerSecret", gprops["consumerSecret"])
	tc.SetInput("accessToken", gprops["accessToken"])
	tc.SetInput("accessTokenSecret", gprops["accessTokenSecret"])
	tc.SetInput("function", "Unfollow")
	tc.SetInput("user", "@Google")
	//tc.SetInput("text", "")

	act.Eval(tc)

	//check result attr
	code := tc.GetOutput("statusCode")
	msg := tc.GetOutput("message")
	fmt.Print(msg)
	assert.Equal(t, 200, code)
}

func TestTwitterFunction_directMessage(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(act.Metadata())

	//setup attributes
	tc.SetInput("consumerKey", gprops["consumerKey"])
	tc.SetInput("consumerSecret", gprops["consumerSecret"])
	tc.SetInput("accessToken", gprops["accessToken"])
	tc.SetInput("accessTokenSecret", gprops["accessTokenSecret"])
	tc.SetInput("function", "directMessage")
	tc.SetInput("user", "@someone")
	tc.SetInput("text", "hello on Twitter 2")

	act.Eval(tc)

	//check result attr
	code := tc.GetOutput("statusCode")
	msg := tc.GetOutput("message")
	fmt.Print(msg)
	assert.Equal(t, 200, code)
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
