package goeval

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

const (
	ivInput     = "Input"
	ivScriptURL = "ScriptURL"
	ovOutput    = "Output"
)

var activityLog = logger.GetLogger("Scripting-activity-GO")

type goevalActivity struct {
	metadata *activity.Metadata
}

//NewActivity TCI Wi Activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &goevalActivity{metadata: metadata}
}

func (a *goevalActivity) Metadata() *activity.Metadata {
	return a.metadata
}
func (a *goevalActivity) Eval(context activity.Context) (done bool, err error) {
	activityLog.Info("Eval GO Scripting activity")
	//Read Inputs
	if context.GetInput(ivInput) == nil {
		// ivInput string is not configured
		// return error to the engine
		return false, activity.NewError("Scripting ivInput string is not configured", "Scripting-GO-4001", nil)
	}
	input := context.GetInput(ivInput).(string)

	if context.GetInput(ivScriptURL) == nil {
		// APIsecret string is not configured
		// return error to the engine
		return false, activity.NewError("Scripting URL string is not configured", "Scripting-GO-4002", nil)
	}
	scriptURL := context.GetInput(ivScriptURL).(string)

	// execute validation - Start

	// load Script from Web
	res, err := http.Get(scriptURL)
	if err != nil {
		return false, activity.NewError("Scripting URL not found", "Scripting-GO-4003", nil)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var script = string(body)

	// execute
	output := evalgo(input, script)

	// Output
	context.SetOutput(ovOutput, output)

	return true, nil
}

func evalgo(in string, script string) int {

	activityLog.Info("GO Eval script: ", script)

	s := NewScope()
	s.Set("print", fmt.Println) //not used yet, but with this you are able to replace any function, e.g. to output something.
	s.Eval("input :=" + in)
	output, err := s.Eval(script)
	activityLog.Info("activity output: ", output, err)

	ret := output.(int)
	return ret
}
