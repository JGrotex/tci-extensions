# TCI LiveApps custom Extension
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

first draft Version with just a list Cases Activity

> Implemenation is now using latest V3 Login.

## Working already ...
- get SSO Token for TIBCO Cloud
- do Login to LiveApps
- call Org.Claims
- phrase for Sandbox ID of Type Production

## ToDo
- Implement more Features like e.g:
- set new Case Owner or app Approver
- query and list Cases

## Activities
available Activities so far

### set Approver (not implemented yet)
Sample Implementation in GO, to set a new Case Owner using Case Actions

Input
- location              string  (eu, au, us)
- user                  string
- pass                  string
- clientid              string
- approver              string

### get Cases (not implemented yet)
Sample Implementation in GO, to retrieve a specific Case Type List of Case Instances

![image](../../screenshots/nonyet.png?raw=true "TCI Screenshot")

Input
- location              string  (eu, au, us)
- user                  string
- pass                  string
- clientid              string
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
