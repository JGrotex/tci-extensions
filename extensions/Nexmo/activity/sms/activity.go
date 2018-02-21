package sms

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

const (
	ivAPIKey     = "apiKey"
	ivAPISecret  = "apiSecret"
	ivFromNumber = "FromNumber"
	ivToNumber   = "ToNumber"
	ivSMStext    = "SMStext"
	ovsend       = "send"
)

var activityLog = logger.GetLogger("nexmo-activity-sms")

type smsActivity struct {
	metadata *activity.Metadata
}

//NewActivity TCI Wi Activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &smsActivity{metadata: metadata}
}

func (a *smsActivity) Metadata() *activity.Metadata {
	return a.metadata
}
func (a *smsActivity) Eval(context activity.Context) (done bool, err error) {
	activityLog.Info("Executing Nexmo SMS Sender activity")
	//Read Inputs
	if context.GetInput(ivAPIKey) == nil {
		// APIkey string is not configured
		// return error to the engine
		return false, activity.NewError("Nexmo API Key string is not configured", "Nexmo-SMS-4001", nil)
	}
	APIkey := context.GetInput(ivAPIKey).(string)

	if context.GetInput(ivAPISecret) == nil {
		// APIsecret string is not configured
		// return error to the engine
		return false, activity.NewError("Nexmo API Secret string is not configured", "Nexmo-SMS-4002", nil)
	}
	APIsecret := context.GetInput(ivAPISecret).(string)

	if context.GetInput(ivFromNumber) == nil {
		// FromNumber string is not configured
		// return error to the engine
		return false, activity.NewError("Nexmo FromNumber string is not configured", "Nexmo-SMS-4003", nil)
	}
	FromNumber := context.GetInput(ivFromNumber).(string)

	if context.GetInput(ivToNumber) == nil {
		// ToNumber string is not configured
		// return error to the engine
		return false, activity.NewError("Nexmo ToNumber string is not configured", "Nexmo-SMS-4004", nil)
	}
	ToNumber := context.GetInput(ivToNumber).(string)
	activityLog.Info("ToNumber: " + ToNumber)

	if context.GetInput(ivSMStext) == nil {
		// SMStext string is not configured
		// return error to the engine
		return false, activity.NewError("Nexmo SMStext string is not configured", "Nexmo-SMS-4005", nil)
	}
	SMStext := context.GetInput(ivSMStext).(string)
	activityLog.Info("SMStext: " + SMStext)

	// execute validation - Start
	urlStr := "https://rest.nexmo.com/sms/json"

	// Build out the data for our message
	v := url.Values{}
	v.Set("api_key", APIkey)
	v.Set("api_secret", APIsecret)
	v.Set("to", ToNumber)
	v.Set("from", FromNumber)
	v.Set("text", SMStext)
	rb := *strings.NewReader(v.Encode())

	// Create client
	client := &http.Client{}

	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(bodyBytes, &data)
		if err == nil {
			context.SetOutput(ovsend, true)
			//var sid string
			//sid = data["message-id"].(string)
			//activityLog.Info("Nexmo SMS SID: " + sid)
		}
	} else {
		context.SetOutput(ovsend, false)
		activityLog.Error("Nexmo SMS Status: " + resp.Status)
	}

	return true, nil
}
