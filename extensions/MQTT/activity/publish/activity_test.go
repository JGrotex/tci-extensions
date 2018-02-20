package publish

import (
	"fmt"
	"io/ioutil"
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
	//setup attrs

	fmt.Println("Publishing a TCI test message to a topic")

	tc.SetInput("broker", "tcp://<your MQTT host>:1883")
	tc.SetInput("id", "tci_tester")
	tc.SetInput("user", "")
	tc.SetInput("password", "")
	tc.SetInput("topic", "<your topic>/temp")
	tc.SetInput("qos", 0)
	tc.SetInput("message", "some message ... ")

	act.Eval(tc)

	//check result attr
	result := tc.GetOutput("send")
	assert.Equal(t, true, result)

	t.Log(result)

}
