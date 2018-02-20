package execjs

/*
 * Copyright © 2018. TIBCO Software Inc. [JGR]
 * This file is subject to the license terms contained
 * in the license file that is distributed with this file.
 */

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/robertkrimen/otto"
)

const (
	ivInput     = "Input"
	ivScriptURL = "ScriptURL"
	ovOutput    = "Output"
)

var errhalt = errors.New("Stahp")

var activityLog = logger.GetLogger("Scripting-activity-JS")

type execjsActivity struct {
	metadata *activity.Metadata
}

//NewActivity TCI Wi Activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &execjsActivity{metadata: metadata}
}

func (a *execjsActivity) Metadata() *activity.Metadata {
	return a.metadata
}
func (a *execjsActivity) Eval(context activity.Context) (done bool, err error) {
	activityLog.Info("Executing JS Scripting activity")
	//Read Inputs
	if context.GetInput(ivInput) == nil {
		// ivInput string is not configured
		// return error to the engine
		return false, activity.NewError("Scripting ivInput string is not configured", "Scripting-JS-4001", nil)
	}
	input := context.GetInput(ivInput).(string)

	if context.GetInput(ivScriptURL) == nil {
		// APIsecret string is not configured
		// return error to the engine
		return false, activity.NewError("Scripting URL string is not configured", "Scripting-JS-4002", nil)
	}
	scriptURL := context.GetInput(ivScriptURL).(string)

	// load Script from Web
	res, err := http.Get(scriptURL)
	if err != nil {
		return false, activity.NewError("Scripting URL not found", "Scripting-JS-4003", nil)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var script = string(body)

	// execute validation - Start
	output := ottorun(input, script)

	// Output
	context.SetOutput(ovOutput, output)

	return true, nil
}

func ottorun(in string, script string) string {

	start := time.Now()
	defer func() {
		duration := time.Since(start)
		if caught := recover(); caught != nil {
			if caught == errhalt {
				fmt.Fprintf(os.Stderr, "Some code took to long! Stopping after: %v\n", duration)
				return
			}
			panic(caught) // Something else happened, repanic!
		}
		fmt.Fprintf(os.Stderr, "Ran code successfully: %v\n", duration)
	}()

	//new JS VM
	vm := otto.New()
	//Set Input Vars
	vm.Set("input", in)

	//Timeout
	vm.Interrupt = make(chan func(), 1) // The buffer prevents blocking
	go func() {
		time.Sleep(4 * time.Second) // Stop after some seconds
		vm.Interrupt <- func() {
			panic(errhalt)
		}
	}()

	feedback, err := vm.Run(script)

	if err != nil {
		// if there is an error, then value.IsUndefined() is true
		fmt.Printf("f", " Syntax Error", err)
	}

	ret, _ := feedback.ToString()
	return ret
}
