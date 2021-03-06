package drone

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
	// username=<your email>
	// tempfolder=<your tempfolder>

	props, err := ReadPropertiesFile("c:\\GODev\\DroneApp.properties")
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

func TestDroneFunction_justFlight(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(act.Metadata())

	//setup attributes
	tc.SetInput("username", gprops["username"])
	tc.SetInput("tempfolder", gprops["tempfolder"])
	tc.SetInput("flighttime", "3")
	tc.SetInput("function", "justFlight")

	act.Eval(tc)

	//check result attr
	code := tc.GetOutput("statusCode")
	msg := tc.GetOutput("message")
	fmt.Print(msg)
	assert.Equal(t, 200, code)
}

func TestDroneFunction_Flight(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(act.Metadata())

	//setup attributes
	tc.SetInput("username", gprops["username"])
	tc.SetInput("tempfolder", gprops["tempfolder"])
	tc.SetInput("flighttime", "2")
	tc.SetInput("function", "Flight")

	act.Eval(tc)

	//check result attr
	code := tc.GetOutput("statusCode")
	msg := tc.GetOutput("message")
	fmt.Print(msg)
	assert.Equal(t, 200, code)
}

func TestDroneFunction_AdvFlight(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(act.Metadata())

	//setup attributes
	tc.SetInput("username", gprops["username"])
	tc.SetInput("tempfolder", gprops["tempfolder"])
	tc.SetInput("flighttime", "2")
	tc.SetInput("function", "AdvFlight")

	act.Eval(tc)

	//check result attr
	code := tc.GetOutput("statusCode")
	msg := tc.GetOutput("message")
	fmt.Print(msg)
	assert.Equal(t, 200, code)
}

func TestDroneFunction_Picture(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(act.Metadata())

	//setup attributes
	tc.SetInput("username", gprops["username"])
	tc.SetInput("tempfolder", gprops["tempfolder"])
	tc.SetInput("flighttime", "2")
	tc.SetInput("function", "Picture")

	act.Eval(tc)

	//check result attr
	code := tc.GetOutput("statusCode")
	msg := tc.GetOutput("message")
	fmt.Print(msg)
	assert.Equal(t, 200, code)
}

func TestDroneFunction_PictureNow(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(act.Metadata())

	//setup attributes
	tc.SetInput("username", gprops["username"]+"now")
	tc.SetInput("tempfolder", gprops["tempfolder"])
	tc.SetInput("flighttime", "2")
	tc.SetInput("function", "PictureNow")

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
