package approver

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
	ivLoc      = "location"
	ivUser     = "user"
	ivPass     = "pass"
	ivclientid = "clientid"
	ivCaseType = "caseType"
	ovCases    = "cases"
)

var activityLog = logger.GetLogger("liveapps-activity-approver")
var baseurlStr = "liveapps.cloud.tibco.com"

type approverActivity struct {
	metadata *activity.Metadata
}

//NewActivity Flogo Activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &approverActivity{metadata: metadata}
}

func (a *approverActivity) Metadata() *activity.Metadata {
	return a.metadata
}

//Eval Flogo Activity
func (a *approverActivity) Eval(context activity.Context) (done bool, err error) {
	activityLog.Info("Executing LiveApps activity")
	//Read Inputs
	if context.GetInput(ivLoc) == nil {
		// return error to the engine
		return false, activity.NewError("LiveApps Location string is not configured", "LiveApps-4001", nil)
	}
	Location := context.GetInput(ivLoc).(string)

	if context.GetInput(ivUser) == nil {
		// return error to the engine
		return false, activity.NewError("LiveApps Username string is not configured", "LiveApps-4002", nil)
	}
	User := context.GetInput(ivUser).(string)

	if context.GetInput(ivPass) == nil {
		// return error to the engine
		return false, activity.NewError("LiveApps Password string is not configured", "LiveApps-4003", nil)
	}
	Pass := context.GetInput(ivPass).(string)

	if context.GetInput(ivclientid) == nil {
		// return error to the engine
		return false, activity.NewError("TIBCO Cloud ClientID string is not configured", "LiveApps-4004", nil)
	}
	ClientID := context.GetInput(ivclientid).(string)

	if context.GetInput(ivCaseType) == nil {
		// return error to the engine
		return false, activity.NewError("LiveApps Application Name string is not configured", "LiveApps-4005", nil)
	}
	//CaseType := context.GetInput(ivCaseType).(string)

	// 1.) *** get SSO TIBCO Cloud AccessToken, only for Login V1 and V2
	/*	accessToken, err := getAccessToken(User, Pass)
		if err != nil {
			panic("Error while getting TC Access_Token")
		}
		activityLog.Info("*** AccessToken: " + accessToken)
	*/
	// 2.) *** login to LiveApps Location/Region
	sessionCookie, err := sessionLogin(Location, User, Pass, ClientID)
	if err != nil {
		panic("Error while getting Session Cookie")
	}
	activityLog.Info("*** Session Cookie: " + sessionCookie)

	// 3.) *** get Organisation Claims
	org, err := getOrgClaims(sessionCookie)
	if err != nil {
		panic("Error while getting Org. Claims")
	}
	activityLog.Info("*** Session Org. Claims: " + org)

	// 4.) *** get Applications / CaseTypes
	//TODO

	// 5.) *** get Case Instances
	//TODO

	// 6.) *** map Result Case List
	//TODO
	context.SetOutput(ovCases, "")

	return true, nil
}

// *** FUNCTION ... get SSO TIBCO Cloud AccessToken
func getAccessToken(user string, pass string) (accessToken string, error error) {

	// execute validation - Start
	urlStr := "https://sso-ext.tibco.com/as/token.oauth2?grant_type=password"

	// Build out the data for our message
	v := url.Values{}
	v.Set("username", user)
	v.Set("password", pass)
	v.Set("client_id", "ropc_ipass")

	rb := *strings.NewReader(v.Encode())

	// Create client
	client := &http.Client{}

	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	activityLog.Info("*** getting Access Token")

	// Make request
	resp, _ := client.Do(req)

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(bodyBytes, &data)
		if err == nil {
			//activityLog.Info("D: " + data["access_token"].(string))
			//fmt.Print("D: " + data["access_token"].(string))

			return data["access_token"].(string), nil
		}
	} else {
		activityLog.Error("LiveApps Status: " + resp.Status)
	}
	return "", nil
}

// *** FUNCTION ... login to LiveApps Location/Region
func sessionLogin(location string, username string, password string, clientid string) (sessionCookie string, error error) {

	// execute validation - Start

	// change base URL to loacation
	if location == "eu" {
		baseurlStr = "https://eu." + baseurlStr
	} else if location == "au" {
		baseurlStr = "https://au." + baseurlStr
	} else {
		baseurlStr = "https://" + baseurlStr
	}

	//add operation
	urlStr := baseurlStr + "/idm/v3/login-oauth"

	// Build out the data for our message
	v := url.Values{}
	v.Set("TenantId", "bpm")
	v.Set("ClientID", clientid)
	v.Set("Email", username)
	v.Set("Password", password)

	rb := *strings.NewReader(v.Encode())

	// Create client
	client := &http.Client{}

	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")

	activityLog.Info("*** getting Session Cookie")

	// Make request
	resp, _ := client.Do(req)

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {

		cookies := resp.Header["Set-Cookie"]

		tsc := strings.Split(cookies[1], ";")[0]
		domain := strings.Split(cookies[2], ";")[0]

		return tsc + ";" + domain, nil

	} else {
		activityLog.Error("LiveApps Status: " + resp.Status)
	}
	return "", nil
}

type Claims struct {
	Email     string `json:"email"`
	Sandboxes []struct {
		Id   string `json:"id"`
		Type string `json:"type"`
	} `json:"sandboxes"`
}

// *** FUNCTION ... get Organisation Claims
func getOrgClaims(sessionCookie string) (org string, error error) {

	// execute validation - Start

	//add operation
	urlStr := baseurlStr + "/organisation/claims"

	// Create client
	client := &http.Client{}

	req, _ := http.NewRequest("GET", urlStr, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Cookie", sessionCookie)

	activityLog.Info("*** getting Org. Claims")

	// Make request
	resp, _ := client.Do(req)

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {

		var data Claims

		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(bodyBytes, &data)
		if err == nil {

			//ToDO: get Sandbox ID of Type Production Sandbox

			sandboxID := data.Sandboxes[0].Id
			activityLog.Info("*** Sandbox ID Result: " + sandboxID)

			return sandboxID, nil
		}
	} else {
		activityLog.Error("LiveApps Status: " + resp.Status)
	}
	return "", nil
}

// *** FUNCTION ... get Applications / CaseTypes
//not yet fully implemented
//TODO

// *** FUNCTION ... get Case Instances
//not yet fully implemented
//TODO
