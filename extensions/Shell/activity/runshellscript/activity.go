package runshellscript

import (
	"encoding/json"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"os"
	"os/exec"
	"strings"
)

// Constants
const (
	command = "cmd"
	params  = "params"
	result  = "result"
)

// log is the default package logger which we'll use to log
var log = logger.GetLogger("activity-setQoS")

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	//ivMsg := `{"cmd":"./test.sh","params":"aaa bbb ccc}`
	// put input varable into a slice (note not order guarenteed)
	ivMsg := context.GetInput(command).(string)
	var ivCmdParams map[string]interface{}
	json.Unmarshal([]byte(ivMsg), &ivCmdParams)

	// get the command to execute including path
	cmd, ok := ivCmdParams[command].(string) // this should be the command or script to execute
	if ok == false {
		// no Command to execute
		log.Infof("No Command to execute - check input syntax: [%s]", err)
		context.SetOutput(result, err.Error())
		return true, err
	}
	// We should have the command to execture - check its there
	_, err = os.Stat(cmd)
	if err != nil {
		// If the file doesn't exist return error
		context.SetOutput("result", err.Error())
		log.Infof("File [%s] does not exist", cmd)
		return true, err
	}

	// Get the Commands Params
	var paramsArray [20]string                    // FIX THIS: make dynamic but ordered
	cmdParams, ok := ivCmdParams[params].(string) // this is a string containg space separated parameters
	if ok == false {
		// no params
		log.Infof("No params provided")
	} else {
		// Put into array for exec.Command to use, space separated
		// put command arguments into an array in the order they are entered.  Order is important.
		i := 0
		for _, field := range split(cmdParams, ' ') {
			paramsArray[i] = field
			i++
		}
	}

	// launch the command
	var cmdOut []byte
	if cmdOut, err = exec.Command(cmd, paramsArray[0:]...).Output(); err != nil {
		log.Infof("Error running Flogo setQoS activity: [%s]", err)
		context.SetOutput(result, err.Error())
		return true, err
	}
	rslt := string(cmdOut)
	// Set the result as part of the context
	context.SetOutput(result, rslt)

	// Signal to the Flogo engine that the activity is completed
	return true, nil
}

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	m := make(map[string]int)
	for _, word := range words {
		m[word] += 1
	}
	return m
}

func split(tosplit string, sep rune) []string {
	var fields []string

	last := 0
	for i, c := range tosplit {
		if c == sep {
			// Found the separator, append a slice
			fields = append(fields, string(tosplit[last:i]))
			last = i + 1
		}
	}

	// Don't forget the last field
	fields = append(fields, string(tosplit[last:]))

	return fields
}
