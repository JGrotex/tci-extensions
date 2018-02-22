# TCI MQTT Extension
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

first Version with allows to publish a message to a MQTT Gateway from TIBCO Cloud Integration Web Integrator.

Similar Implemenation to the [Flogo MQTT](https://github.com/jvanderl/flogo-components/tree/master/activity/mqtt) one from Jan van der Lugt.<br>
This Version is just prepackaged for TIBCO Cloud Integration Web Integrator, and adds a Icons, Category, and some Input Parameter under Configuration Tab.

![MQTT Publish image](../../screenshots/MQTT-pub.png?raw=true "TCI WI MQTT Publish Screenshot") 

## Activities
available Activities so far
### MQTT Publish
Sample publish to a MQTT topic,

Input
- host          string (host,port)
- id            string (client ID)
- username      string
- password      string
- qos           number 0,1,2
- topic         string
- message       string

<i>Hint:</i> none yet.

Output
- send               bool   `json:"send"`
  
Sample Input Data

- host `tcp://<your MQTT host>:1883`
- topic `sometopic/xyz`

Sample Output Data

``json:
{"send":true}
``

<hr>
<sub><b>Note:</b> more TCI Extensions can be found here: https://tibcosoftware.github.io/tci-awesome/ </sub>
