# GObot Extension
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Background
This Implementation is inspired by a Drone Remote Inspection Use case. Where Flogo is used as a IoT Gateway to connect TIBCO Cloud Messaging with a WiFi Drone to contoll it remotely.

More Details about the full Flow Implemenation later, as time allows.

This activity allows you to:
- start a Parrot BeBop 2 Drone,
- take a picture from Stream,
- land the Drone safely.

## Parameter Details
| Setting           | Required  | Description                              |
|:------------------|:----------|:-----------------------------------------|
| username          | true      | Imagename build with the Name            |
| function          | true      | 'Picture' or 'Flight' with Picture taken |
| tempfolder        | true      | Folder to store Images                   |
| flighttime        | true      | Time to fly in Seconds  2-5sec           |       

## Third-party libraries used
package GoBot - [https://gobot.io](https://gobot.io) :: Thanks to the GoBot Team Members for having this available using tiny GO.
