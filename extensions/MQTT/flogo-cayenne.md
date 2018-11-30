# How to connect Flogo with Cayenne
regardless if you want to connect to Cayenne Dashboard (https://www.mydevices.com) from TIBCO Cloud Integration Flogo, or from Flogo (https://www.flogo.io) natively on Device. You can just use the MQTT Activity stored here, or this [Flogo MQTT Activity](https://github.com/jvanderl/flogo-components/tree/master/activity/mqtt) from Jan van der Lugt.<br>

Just configure the graphical Flogo Flow Activity to contain your new Device Cayenne Account Details, and you configure your realtime Web Dashboard and/or Mobile App Dashboard. 

## TIBCO Cloud Integration Flogo Sample
Screenshot of TIBCO Cloud Integration Flogo and Cayenne

![TCI Flogo and Cayenne](../../screenshots/TCI-Flogo-Cayenne.png?raw=true "TIBCO Cloud Integration Flogo and Cayenne") 

## TIBCO Flogo.io Sample
Screenshot of TIBCO Flogo.io Web UI and Cayenne

![Flogo and Cayenne](../../screenshots/Flogo-Cayenne.png?raw=true "TIBCO Flogo.io and Cayenne") 

## create new Flogo Device in Cayenne
to add a new Flogo Device or Flogo Cloud Service to Cayenne just follow these simple steps

1. select "Add new..."
2. select "Bring Your Own Thing"

![new Flogo Device](../../screenshots/step1-cayenne-bring-own.png?raw=true "add Flogo Device or Service to Cayenne") 

3. copy and paste your Details to the Flogo MQTT Activity
4. test or send something, as Cayenne is already waiting for you

![waiting for Flogo Device](../../screenshots/step2-cayenne-connect-details.png?raw=true "Cayenne is already waiting for Flogo") 

## Testing
Sample use a HTTP REST Message as Service put with flexible channel selection, with this new channels can be created on the fly.

![Flogo and Cayenne](../../screenshots/TCI-Flogo-Cayenne-testing.png?raw=true "TIBCO Flogo.io and Cayenne") 

``json:
{
    "channel":"1",
    "message":"temp,c=22.5"
}
``

<hr>
<sub><b>Note:</b> more TCI Extensions can be found here: https://tibcosoftware.github.io/tci-awesome/ </sub>
