package ftpsendfile

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strconv"
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
	// server=<your FTP Server>
	// username=<your username>
	// password=<your secret>
	// pathsrc=<dest folder>
	// filename=<filename>
	// pathdest=<target folder>

	props, err := ReadPropertiesFile("c:\\GODev\\FTPApp.properties")
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

	//setup attrs
	server := gprops["server"]
	port := 21
	url := server + ":" + strconv.Itoa(port)

	tc.SetInput("server", server)
	tc.SetInput("port", port)
	tc.SetInput("username", gprops["username"])
	tc.SetInput("password", gprops["password"])
	tc.SetInput("pathsrc", gprops["pathsrc"])
	tc.SetInput("filesrc", gprops["filename"])
	tc.SetInput("pathdest", gprops["pathdest"])
	tc.SetInput("filedest", gprops["filename"])

	act.Eval(tc)

	//check result attr
	result := tc.GetOutput("output")
	assert.Equal(t, result, "Successfully connected to "+url)
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
