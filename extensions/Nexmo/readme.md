# TCI Nexmo Extension
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

first draft Version with just a SMS sender using Nexmo API.

Please create your own Access Key on Nexmo (https://www.nexmo.com) to enter into the Activity Details.
A TCI Connector is planed for later. 

## Activities
available Activities so far
### SMS Sender
Sample SMS smartphone Screen,

![Nexmo SMS image](../../screenshots/nexmo-SMS.png?raw=true "TCI WI Nexmo SMS Screenshot")

Input
- apiKey                string
- apiSecret             string
- FromPhonenumber       string
- ToPhonenumber         string
- SMStext               string

<i>Hint:</i> Sender is working as well if you specify only a String as FromPhonenumber e.g. 'TIBCO Casemanagement'.

Output
- send               bool   `json:"send"`
  
Sample Input Data
your Nexmo Account Data for apiKey and apiSecret
+49171.... 
"some text ..."

Sample Output Data

``json:
{"send":true}
``

<hr>
<sub><b>Note:</b> more TCI Extensions can be found here: https://tibcosoftware.github.io/tci-awesome/ </sub>
