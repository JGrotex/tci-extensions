
package sms

import (
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/stretchr/testify/assert"
)

var activityMetadata *activity.Metadata
var connectionData = ``

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

func TestActivityRegistration(t *testing.T) {
	act := NewActivity(getActivityMetadata())
	if act == nil {
		t.Error("Activity Not Registered")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {
	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(act.Metadata())

	// *** for testing, replace all in <> with your Account Details!

	//setup attrs
	tc.SetInput("apiKey", "<your key>")
	tc.SetInput("apiSecret", "<your Screcet>")
	tc.SetInput("FromNumber", "NEXMO")
	tc.SetInput("ToNumber", "+49171....<your number>")
	tc.SetInput("SMStext", "hi from GODev")

	_, err := act.Eval(tc)
	assert.Nil(t, err)

	result := tc.GetOutput("send")
	assert.Equal(t, result, true)

	t.Log(result)
}
