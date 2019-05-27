package drone

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	s "strings"
	"time"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"

	"gobot.io/x/gobot/platforms/parrot/bebop/client"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

var activityLog = logger.GetLogger("activity-drone")

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

func ffmpeg() (stdin io.WriteCloser, stderr io.ReadCloser, err error) {
	ffmpeg := exec.Command("ffmpeg", "-i", "pipe:0", "http://localhost:8090/bebop.ffm")

	stderr, err = ffmpeg.StderrPipe()

	if err != nil {
		return
	}

	stdin, err = ffmpeg.StdinPipe()

	if err != nil {
		return
	}

	if err = ffmpeg.Start(); err != nil {
		return
	}

	go func() {
		for {
			buf, err := ioutil.ReadAll(stderr)
			if err != nil {
				fmt.Println(err)
			}
			if len(buf) > 0 {
				fmt.Println(string(buf))
			}
		}
	}()

	return stdin, stderr, nil
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {
	activityLog.Info("Executing Drone activity")

	username := s.TrimSpace(context.GetInput("username").(string))

	dfunction := context.GetInput("function").(string)
	activityLog.Info("Drone Function: " + dfunction)

	if len(username) == 0 {

		context.SetOutput("statusCode", 190)
		context.SetOutput("message", "username field is blank")

	} else {
		var msg string
		var code int

		switch dfunction {
		case "justFlight":
			{
				// Implemenation following this Sample
				// https://github.com/hybridgroup/gobot/blob/master/platforms/parrot/bebop/client/examples/video.go

				tempfolder := s.TrimSpace(context.GetInput("tempfolder").(string))
				flighttime := s.TrimSpace(context.GetInput("flighttime").(string))

				if len(username) == 0 {

					code = 100
					msg = "username cannot be blank"

				} else {

					var err = os.Remove(tempfolder + "img-" + username + ".jpg")
					if err != nil {
						fmt.Println(err)
					}

					time.Sleep(1 * time.Second)

					bebop := client.New()
					if err := bebop.Connect(); err != nil {
						fmt.Println("Connect Err: ", err)
					}

					bebop.HullProtection(false)
					bebop.Outdoor(false)

					fmt.Println("takeoff")
					if err := bebop.TakeOff(); err != nil {
						fmt.Println(err)
						fmt.Println("fail")
					}

					secs, _ := time.ParseDuration(flighttime + "s")
					time.Sleep(secs)

					if err := bebop.Clockwise(10); err != nil {
						fmt.Println(err)
						fmt.Println("fail")
					}

					time.Sleep(secs)

					if err := bebop.CounterClockwise(20); err != nil {
						fmt.Println(err)
						fmt.Println("fail")
					}

					time.Sleep(secs)

					fmt.Println("land")
					if err := bebop.Land(); err != nil {
						fmt.Println(err)
					}

					code = 200
					msg = ""

				}
			}
		case "Flight":
			{
				// Implemenation following this Sample
				// https://github.com/hybridgroup/gobot/blob/master/platforms/parrot/bebop/client/examples/video.go

				tempfolder := s.TrimSpace(context.GetInput("tempfolder").(string))
				flighttime := s.TrimSpace(context.GetInput("flighttime").(string))

				if len(username) == 0 {

					code = 100
					msg = "username cannot be blank"

				} else {

					var err = os.Remove(tempfolder + "img-" + username + ".jpg")
					if err != nil {
						fmt.Println(err)
					}

					time.Sleep(1 * time.Second)

					bebop := client.New()
					if err := bebop.Connect(); err != nil {
						fmt.Println("Connect Err: ", err)
					}
					if err := bebop.VideoEnable(true); err != nil {
						fmt.Println("Video Err: ", err)
					}
					if err := bebop.VideoStreamMode(0); err != nil {
						fmt.Println("StreamMode Err: ", err)
					}

					bebop.HullProtection(false)
					bebop.Outdoor(false)

					fmt.Println("takeoff")
					if err := bebop.TakeOff(); err != nil {
						fmt.Println(err)
						fmt.Println("fail")
					}

					secs, _ := time.ParseDuration(flighttime + "s")
					time.Sleep(secs)

					cmd := exec.Command("ffmpeg", "-protocol_whitelist", "file,rtp,udp", "-i", tempfolder+"drone.sdp", "-r", "30", tempfolder+"img-"+username+".jpg")
					cmd.Run()
					cmd = nil

					fmt.Println("land")
					if err := bebop.Land(); err != nil {
						fmt.Println(err)
					}

					time.Sleep(2 * time.Second)

					/*		if err := bebop.VideoEnable(false); err != nil {
								fmt.Println(err)
							}

							if err := bebop.Close(); err != nil {
								fmt.Println(err)
							}
					*/

					code = 200
					msg = ""

				}
			}
		case "AdvFlight":
			{
				// Implemenation following this Sample
				// https://github.com/hybridgroup/gobot/blob/master/platforms/parrot/bebop/client/examples/video.go

				fmt.Println("AdvFlight")

				tempfolder := s.TrimSpace(context.GetInput("tempfolder").(string))
				flighttime := s.TrimSpace(context.GetInput("flighttime").(string))

				if len(username) == 0 {

					code = 100
					msg = "username cannot be blank"

				} else {

					var err = os.Remove(tempfolder + "img-" + username + ".jpg")
					if err != nil {
						fmt.Println(err)
					}

					time.Sleep(1 * time.Second)

					bebop := client.New()
					if err := bebop.Connect(); err != nil {
						fmt.Println("Connect Err: ", err)
					}
					if err := bebop.VideoEnable(true); err != nil {
						fmt.Println("Video Err: ", err)
					}
					if err := bebop.VideoStreamMode(0); err != nil {
						fmt.Println("StreamMode Err: ", err)
					}

					bebop.HullProtection(false)
					bebop.Outdoor(false)

					fmt.Println("takeoff")
					if err := bebop.TakeOff(); err != nil {
						fmt.Println(err)
						fmt.Println("fail")
					}

					secs, _ := time.ParseDuration(flighttime + "s")
					time.Sleep(secs)

					fmt.Println("Clockwise")
					if err := bebop.Clockwise(10); err != nil {
						fmt.Println(err)
						fmt.Println("fail")
					}

					time.Sleep(secs)

					fmt.Println("CounterClockwise")
					if err := bebop.CounterClockwise(20); err != nil {
						fmt.Println(err)
						fmt.Println("fail")
					}

					time.Sleep(secs)

					cmd := exec.Command("ffmpeg", "-protocol_whitelist", "file,rtp,udp", "-i", tempfolder+"drone.sdp", "-r", "30", tempfolder+"img-"+username+".jpg")
					cmd.Run()
					cmd = nil

					fmt.Println("land")
					if err := bebop.Land(); err != nil {
						fmt.Println(err)
					}

					time.Sleep(2 * time.Second)

					/*		if err := bebop.VideoEnable(false); err != nil {
								fmt.Println(err)
							}

							if err := bebop.Close(); err != nil {
								fmt.Println(err)
							}
					*/

					code = 200
					msg = ""

				}
			}
		case "Picture":
			{
				tempfolder := s.TrimSpace(context.GetInput("tempfolder").(string))
				if len(username) == 0 {

					code = 101
					msg = "username cannot be blank"

				} else {

					var err = os.Remove(tempfolder + "img-" + username + ".jpg")
					if err != nil {
						//fmt.Println(err)
					}

					time.Sleep(1 * time.Second)

					bebop := client.New()
					if err := bebop.Connect(); err != nil {
						fmt.Println("Connect Err: ", err)
					}
					if err := bebop.VideoEnable(true); err != nil {
						fmt.Println("Video Err: ", err)
					}
					if err := bebop.VideoStreamMode(0); err != nil {
						fmt.Println("StreamMode Err: ", err)
					}

					/* take Picture */
					cmd := exec.Command("ffmpeg", "-protocol_whitelist", "file,rtp,udp", "-i", tempfolder+"drone.sdp", "-r", "30", tempfolder+"img-"+username+".jpg")
					cmd.Run()
					cmd = nil

					time.Sleep(2 * time.Second)

					if err := bebop.VideoEnable(false); err != nil {
						fmt.Println(err)
					}

					/*	if err := bebop.Close(); err != nil {
							fmt.Println(err)
						}
					*/
					code = 200
					msg = ""
				}
			}

		case "PictureNow":
			{
				tempfolder := s.TrimSpace(context.GetInput("tempfolder").(string))
				if len(username) == 0 {

					code = 101
					msg = "username cannot be blank"

				} else {

					var err = os.Remove(tempfolder + "img-" + username + ".jpg")
					if err != nil {
						//fmt.Println(err)
					}

					/* take Picture */
					cmd := exec.Command("ffmpeg", "-protocol_whitelist", "file,rtp,udp", "-i", tempfolder+"drone.sdp", "-r", "30", tempfolder+"img-"+username+".jpg")
					cmd.Run()
					cmd = nil

					code = 200
					msg = ""
				}
			}

		default:
			{
				code = 150
				msg = "Function field cannot be blank"
			}
		}

		activityLog.Info("Drone Activity executed with Status Code: " + "200")

		context.SetOutput("statusCode", code)
		context.SetOutput("message", msg)
	}

	return true, err
}
