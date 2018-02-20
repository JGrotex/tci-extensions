# TCI Scripting Extension
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Early draft Version of a Scripting activity.

First Version will allow to execute JavaScripts using the OTTO Engine. 
Next Version will allow to use GOlang as well - this GOlang exists already as fragment, will complete it later on.

Considering to handover multiple Values as Object as well. Need to do some testing around this.

## Activities
available Activities so far
### JavaScript Execution 
Sample of dynamic Javascript exection, the script is loaded on the fly from any URL.
So it can be changed at any time without any redeployment, etc.

Input
- Input                 string (...later Object)
- ScriptURL             string

Output
- Output                string (...later Object)

### JavaScript Hints 
here a first sample Scripts to how it works ...

Sample Input Data
- Input "some data"
- ScriptURL "http://www.godev.de/logicscripts/dynconcatsample.js"

Sample Output Data
- Output = "Input: some data"

```js 
var feedback = "Input: " + input;
console.log("JS VM - the value " + feedback);
feedback;
```

Sample Input Data
- Input "green"
- ScriptURL "http://www.godev.de/logicscripts/dynifsample.js"

Sample Output Data
- Output = "25"

```js 
var feedback;

if (input=="green"){
	console.log("JS VM: green path selected");
	feedback = "25";
} else {
	console.log("JS VM: other path selected");
	feedback = "75";
};

feedback;
```

<hr>
<sub><b>Note:</b> more TCI Extensions can be found here: https://tibcosoftware.github.io/tci-awesome/ </sub>
