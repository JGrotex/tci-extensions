/*
 * Copyright © 2018. TIBCO Software Inc. [JGR]
 * This file is subject to the license terms contained
 * in the license file that is distributed with this file.
 */
package execjs

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

	//setup attrs
	tc.SetInput("Input", "55656")
	tc.SetInput("ScriptURL", "http://www.godev.de/logicscripts/dynconcatsample.js")

	_, err := act.Eval(tc)
	assert.Nil(t, err)

	result := tc.GetOutput("Output")
	assert.Contains(t, result, "55")
	//assert.Equal(t, result, "done")

	t.Log(result)
}

func TestEval2(t *testing.T) {
	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(act.Metadata())

	//setup attrs
	tc.SetInput("Input", "green")
	tc.SetInput("ScriptURL", "http://www.godev.de/logicscripts/dynifsample.js")

	_, err := act.Eval(tc)
	assert.Nil(t, err)

	result := tc.GetOutput("Output")
	assert.Contains(t, result, "25")
	//assert.Equal(t, result, "done")

	t.Log(result)
}

func TestEval3(t *testing.T) {
	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(act.Metadata())

	//setup attrs
	tc.SetInput("Input", "{\"data\":\"green\"}")
	tc.SetInput("ScriptURL", "http://www.godev.de/logicscripts/dynifsamplecomplex.js")

	_, err := act.Eval(tc)
	assert.Nil(t, err)

	result := tc.GetOutput("Output")
	assert.Contains(t, result, "25")
	//assert.Equal(t, result, "done")

	t.Log(result)
}

func TestEval4(t *testing.T) {
	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(act.Metadata())

	//setup attrs
	tc.SetInput("Input", "{\"data\":\"green\"}")
	tc.SetInput("ScriptURL", "http://www.godev.de/logicscripts/fakername.js")

	_, err := act.Eval(tc)
	assert.Nil(t, err)

	result := tc.GetOutput("Output")
	assert.Contains(t, result, "anything")
	//assert.Equal(t, result, "done")

	t.Log(result)
}
