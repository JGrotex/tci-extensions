package execgo

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

const (
	ivInput     = "Input"
	ivScriptURL = "ScriptURL"
	ovOutput    = "Output"
)

var activityLog = logger.GetLogger("Scripting-activity-GO")

type execgoActivity struct {
	metadata *activity.Metadata
}

//NewActivity TCI Wi Activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &execgoActivity{metadata: metadata}
}

func (a *execgoActivity) Metadata() *activity.Metadata {
	return a.metadata
}
func (a *execgoActivity) Eval(context activity.Context) (done bool, err error) {
	activityLog.Info("Executing GO Scripting activity")
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

	//
	// Draft .... Activity in GO not implemented yet ... will come as time allows
	//

	// Output
	context.SetOutput(ovOutput, "done: "+input+" "+scriptURL)

	return true, nil
}
