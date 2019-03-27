package runshellscript

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/stretchr/testify/assert"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

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

	// set up a script to run
	// test01.sh output is the 3 parameters passed in.
	//	deleteFile("test01.sh")
	//	createFile("test01.sh")
	//	writeFile("test01.sh", `printf "test01 output: $1 $2 $3"`)
	//	setWritePermissions("test01.sh")

	//setup attrs
	tc.SetInput("cmd", `{"cmd":"c:/godev/test.bat","params":"aaa bbb ccc"}`)
	//	tc.SetInput("cmd", `{"cmd":"./test01.sh","params":"aaa bbb ccc"}` )
	//tc.SetInput("cmd", `{"cmd":"/Users/iain/setQoS.sh"}` )
	act.Eval(tc)

	//check result attr
	result := tc.GetOutput("result")
	assert.Equal(t, "aaa bbb ccc\r\n", result)
}

func createFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func deleteFile(filename string) error {
	err := os.Remove(filename)
	if err != nil {
		return err
	}
	return nil
}

func writeFile(filename string, testLine string) error {
	// open file using READ & WRITE  permission
	var file, err = os.OpenFile(filename, os.O_RDWR, 0644)
	defer file.Close()

	// write some text line-by-line to file
	_, err = file.WriteString("#!/bin/bash\n")
	_, err = file.WriteString(testLine)

	// save changes
	err = file.Sync()
	if err != nil {
		return err
	}
	return nil
}

func setWritePermissions(filename string) error {
	err := os.Chmod(filename, 0755)
	if err != nil {
		return err
	}
	return nil
}
