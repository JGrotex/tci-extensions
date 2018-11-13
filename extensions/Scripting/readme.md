# TCI Scripting Extension
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Early draft Version of a Scripting activity.

First Version will allow to execute JavaScripts using the OTTO Engine. 
thanks to ... https://github.com/robertkrimen/otto

New Version allow to use GOlang as well - but is not working with TCI, yet. I see no Issues to use it with Flogo!
thanks to ... https://github.com/novalagung/golpal

You can find a experimental Version for one line GO Eval, too. But this really just a working draft.
thanks to ... https://github.com/xtaci/goeval 

Update: In the meantime 'Vijay Nalawade' created as well Javascript Flogo Activity working with TCI Flogo as well.
This one stores the Script into a Designtime Datafield, choise what you like. ...
https://github.com/vijaynalawade/flogo/tree/master/activity/js

## Activities
available Activities so far
### JavaScript Execution 
Sample of dynamic Javascript exection, the script is loaded on the fly from any URL.
So it can be changed at any time without any redeployment, etc.

![Exec Javascript image](../../screenshots/Scripting-JS.png?raw=true "TCI WI execute Javascript Screenshot")

Multiple Values can be handover vis JSON Object as well.

Input
- Input                 string 
- ScriptURL             string

Output
- Output                string 

### GOlang Execution 
Sample of dynamic GOlang exection, the script is loaded on the fly from any URL.
So it can be changed at any time without any redeployment, etc.

<b>Warning:</b> not working in TCI, yet! This is because Golpal needs a temp. Folder, and TCI does not allow that.
Similar Implemenation shoud work in Flogo without any Issues.

Input
- Input                 string (replaces {input} inside the script)
- ScriptURL             string

Output
- Output                string 

### GOlang Eval 
Just for tiny calculations, etc., the one line string is loaded on the fly from any URL.
So it can be changed at any time without any redeployment, etc.

<b>Warning:</b> just for GOlang calculations, yet!

Input
- Input                 string (input is set as Value to the eval)
- ScriptURL             string

Output
- Output                int 

#### JavaScript Execution Hints 
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

Sample Input Data with complex JSON
- Input "{\"data\":\"green\"}"
- ScriptURL "http://www.godev.de/logicscripts/dynifsamplecomplex.js"

Sample Output Data
- Output = "25"

```js 
var feedback;

console.log("JS VM IN: " + input);
var obj = JSON.parse(input);

if (obj.data=="green"){
	console.log("JS VM: green path selected");
	feedback = "25";
} else {
	console.log("JS VM: other path selected");
	feedback = "75";
};

feedback;
```

#### GOlang Execution Hints

Sample Input Data
- Input "green"
- ScriptURL "http://www.godev.de/logicscripts/dynifsample.go"

Sample Output Data
- Output = "25"

```go
input := "{input}"
feedback := ""

if (input=="green"){
	feedback = "25"
	
} else {
	feedback = "75"
}

return feedback
```
<i>^ as 'golpal' not support Input like 'otto' {input} is replaced with the Input String.</i>

#### GOlang Eval Hints

Sample Input Data
- Input 50
- ScriptURL "http://www.godev.de/logicscripts/simplecalc.go"

Sample Output Data
- Output = "25"

```go
feedback := input / 2
```
<i>^ eval support only simple calculations, and loops, but no 'if', so I recomment to stick with just calculations.</i>

<hr>
<sub><b>Note:</b> more TCI Extensions can be found here: https://tibcosoftware.github.io/tci-awesome/ </sub>
