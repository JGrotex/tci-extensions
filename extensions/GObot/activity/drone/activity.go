package drone

import (
	"fmt"
	"io/ioutil"
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
		case "Flight":
			{
				tempfolder := s.TrimSpace(context.GetInput("tempfolder").(string))
				flighttime := s.TrimSpace(context.GetInput("flighttime").(string))
				if len(username) == 0 {

					code = 100
					msg = "username cannot be blank"

				} else {

					cmddel := exec.Command("del", tempfolder+"img-"+username+".jpg")
					cmddel.Run()
					cmddel = nil

					time.Sleep(1 * time.Second)

					bebop := client.New()
					if err := bebop.Connect(); err != nil {
						fmt.Println(err)
					}
					if err := bebop.VideoEnable(true); err != nil {
						fmt.Println(err)
					}
					if err := bebop.VideoStreamMode(0); err != nil {
						fmt.Println(err)
					}
					ffmpeg := exec.Command("ffmpeg", "-i", "pipe:0", "http://localhost:8090/bebop.ffm")
					ffmpegErr, err := ffmpeg.StderrPipe()
					if err != nil {
						fmt.Println(err)
					}

					ffmpegIn, err := ffmpeg.StdinPipe()
					if err != nil {
						fmt.Println(err)
					}
					if err := ffmpeg.Start(); err != nil {
						fmt.Println(err)
					}

					go func() {
						for {
							buf, err := ioutil.ReadAll(ffmpegErr)
							if err != nil {
								fmt.Println(err)
							}
							if len(buf) > 0 {
								fmt.Println(string(buf))
							}
						}
					}()

					go func() {
						for {
							if _, err := ffmpegIn.Write(<-bebop.Video()); err != nil {
								fmt.Println(err)
							}
						}
					}()

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

					code = 200
					msg = ""

					bebop = nil
				}
			}
		case "Picture":
			{
				tempfolder := s.TrimSpace(context.GetInput("tempfolder").(string))
				if len(username) == 0 {

					code = 101
					msg = "username cannot be blank"

				} else {

					cmddel := exec.Command("del", tempfolder+"img-"+username+".jpg")
					cmddel.Run()
					cmddel = nil

					time.Sleep(1 * time.Second)

					bebop := client.New()
					if err := bebop.Connect(); err != nil {
						fmt.Println(err)
					}
					if err := bebop.VideoEnable(true); err != nil {
						fmt.Println(err)
					}
					if err := bebop.VideoStreamMode(0); err != nil {
						fmt.Println(err)
					}
					ffmpeg := exec.Command("ffmpeg", "-i", "pipe:0", "http://localhost:8090/bebop.ffm")
					ffmpegErr, err := ffmpeg.StderrPipe()

					if err != nil {
						fmt.Println(err)
					}

					ffmpegIn, err := ffmpeg.StdinPipe()
					if err != nil {
						fmt.Println(err)
					}
					if err := ffmpeg.Start(); err != nil {
						fmt.Println(err)
					}

					go func() {
						for {
							buf, err := ioutil.ReadAll(ffmpegErr)
							if err != nil {
								fmt.Println(err)
							}
							if len(buf) > 0 {
								fmt.Println(string(buf))
							}
						}
					}()

					go func() {
						for {
							if _, err := ffmpegIn.Write(<-bebop.Video()); err != nil {
								fmt.Println(err)
							}
						}
					}()

					cmd := exec.Command("ffmpeg", "-protocol_whitelist", "file,rtp,udp", "-i", tempfolder+"drone.sdp", "-r", "30", tempfolder+"img-"+username+".jpg")
					cmd.Run()
					cmd = nil

					code = 200
					msg = ""

					bebop = nil
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
