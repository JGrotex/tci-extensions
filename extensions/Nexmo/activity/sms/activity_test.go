/*
 * Copyright © 2018. TIBCO Software Inc. [JGR]
 * This file is subject to the license terms contained
 * in the license file that is distributed with this file.
 */
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
	tc.SetInput("apiKey", "c541eabe")
	tc.SetInput("apiSecret", "1fa0c30037917d0c")
	tc.SetInput("FromNumber", "NEXMO")
	tc.SetInput("ToNumber", "+491715664015")
	tc.SetInput("SMStext", "hi from GODev")

	_, err := act.Eval(tc)
	assert.Nil(t, err)

	result := tc.GetOutput("send")
	assert.Equal(t, true, result)

	t.Log(result)
}
