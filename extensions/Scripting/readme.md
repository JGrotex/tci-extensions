# TCI Scripting Extension
TIBCO Cloud Integration Flogo Scripting activity.

## Overview
There are 3 Option to dynamically executed Scripts stored on a remote location, all allow Script adjustments on the fly with redeploying the Service to TCI.

1. Allow to execute any JavaScripts using the OTTO Engine. Thanks to ... [https://github.com/robertkrimen/otto](https://github.com/robertkrimen/otto)

2. Allow to execute GOlang as well - but is not working within the Cloud, yet. I see no Issues to use it with 'native' Flogo!
Thanks to ... [https://github.com/novalagung/golpal](https://github.com/novalagung/golpal)

3. You can find a experimental Version for one line GO Eval, too. But this really just a working draft.
Thanks to ... [https://github.com/xtaci/goeval](https://github.com/xtaci/goeval) 

## Activities
The first JavaScript Extension is the most Advanced Option here as it supports longer Scripts.

### JavaScript Execution 
Sample of dynamic Javascript execution, the script is loaded on the fly from any URL.
So it can be changed at any time without any redeployment, etc.

![Exec Javascript image](Scripting-JS.png "TCI WI execute Javascript Screenshot")

Multiple Values can be handover vis JSON Object as well.

Input
- Input                 string 
- ScriptURL             string

Output
- Output                string 

### GOlang Execution 
Sample of dynamic GOlang execution, the script is loaded on the fly from any URL.
So it can be changed at any time without any redeployment, etc.

<b>Warning:</b> not working in TCI, yet! This is because Golpal needs a temp. Folder, and TCI does not allow that.
Similar Implementation should work in Flogo without any Issues.

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
here some first sample Scripts to how it works, more can be found in the Logicscripts GitHub Sample Folder.

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
> <i>^ as 'golpal' not support Input like 'otto' {input} is replaced with the Input String.</i>

#### GOlang Eval Hints

Sample Input Data
- Input 50
- ScriptURL "http://www.godev.de/logicscripts/simplecalc.go"

Sample Output Data
- Output = "25"

```go
feedback := input / 2
```
> <i>^ Eval support only simple calculations, and loops, but no 'if', so I recommend to stick with just calculations.</i>

<hr>
<sub><b>Note:</b> more TCI Extensions can be found here: https://tibcosoftware.github.io/tci-awesome/ </sub>

