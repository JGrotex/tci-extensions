# GODev TCI Extensions
[![Go Report Card](https://goreportcard.com/badge/github.com/JGrotex/tci-extensions)](https://goreportcard.com/report/github.com/JGrotex/tci-extensions) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

This Repro. will get my central area for all future TCI Extensions, will move the others as time allows.

## TCI MQTT Extension
allows to publish MQTT Messages via TCI.
[Details here ...](extensions/MQTT/readme.md)

## TCI Scripting Extension
allows to execute dynamically Javascripts within a TCI Activity, loaded form any URL on the fly.
[Details here ...](extensions/Scripting/readme.md)

## TCI Nexmo Extension
first draft Version with just a SMS sender using Nexmo API.
[Details here ...](extensions/Nexmo/readme.md)

## TCI Twilio Extension
first draft Version with just a SMS sender using Twilio API.
[Details here ...](https://github.com/JGrotex/tci-wi-twilio-extension)

## TCI APILayer dot com Extension
validate Phonenumbers using APILayer dot com API.
[Details here ...](https://github.com/JGrotex/tci-wi-apilayer-extension)

## TCI Banking Extension
validate European International IBAN codes.
[Details here ...](https://github.com/JGrotex/tci-wi-banking-extension)

## TCI Small Tools Extension
a number of tiny tools like Concat, validate Email Adr., create HTML, send HTML Emails, etc.
[Details here ...](https://github.com/JGrotex/tci-wi-smalltools-extension)

## TCI LiveApps Extension
allows to connect to any LiveApps Subscription from TCI Flogo and native Flogo
[Details here ...](extensions/LiveApps/readme.md)

## TCI GObot Extension
allows to control a Parrot BeBop Drone from TCI Flogo and native Flogo
[Details here ...](extensions/GObot/readme.md)

## TCI FTP Extension
allows to files via FTP from TCI Flogo and native Flogo
[Details here ...](extensions/FTP/readme.md)

<hr>

## Others

### other TCI Extensions
more TCI Extensions can be found here in the [official Showcase](https://tibcosoftware.github.io/tci-awesome/)

### helpful source for Flogo Extensions
Links to Flogo Extensions, [here ...](flogokowhow.md)

### How to create a TCI Flogo Extension from a Flogo Extension
most times all the Flogo 'native' Activities you can find in the Web e.g. on GitHub coming with no specific folder structure, Icons, or Activity Category. Here some hints to use them in TCI:
- just Zip the Folder of the Activity following this Folder Structure
  - Folder: your custom Activity Name "anyname"
    - Folder: activity <- defines that this is a Flogo Activity
      - Folder: your Activity Name (use GOlang Package Name)
- upload the ZIP to TCI Extensions
- attached to this Page a little Sample of the Flogo JS Activity
- you find the Flogo Extension in TCI after upload in the "Default" Category.

Remark: Activities using something like the following are not supported on TCI Flogo
- Activities with local File access
- Activities opening own Ports

All necessary additional GOlang Packages automatically discovered by TCI Flogo and automatically downloaded.

<hr>
<sub><b>Note:</b> this List is constandly under my review, try always to keep it actual.</sub>
