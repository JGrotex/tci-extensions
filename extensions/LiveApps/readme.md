# TCI LiveApps custom Extension
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

first draft Version with just a list Cases Activity

## Working already ...
- get SSO Token for TIBCO Cloud
- do Login to LiveApps
- call Org.Claims

## ToDo
- phrase for Sandbox ID of Type Production
- get Application ID, of selected CaseType
- get Case Instance List

## Activities
available Activities so far

### get Cases
Sample Implementation in GO, to retrieve a specific Case Type List of Case Instances

![image](../../screenshots/nonyet.png?raw=true "TCI Screenshot")

Input
- user                  string
- pass                  string
- caseType              string

<i>Hint:</i> caseType is your Application Name e.g. 'Risk Item Approval'.

Output
- cases               bool   `json:"cases"`
  
Sample Output Data

``json:
{"cases":case[]}
``

<hr>
<sub><b>Note:</b> more TCI Extensions can be found here: https://tibcosoftware.github.io/tci-awesome/ </sub>