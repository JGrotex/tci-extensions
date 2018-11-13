package tweet

import (
	"strconv"
	s "strings"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

var activityLog = logger.GetLogger("activity-twitter")

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
	activityLog.Info("Executing Twitter activity")

	consumerKey := s.TrimSpace(context.GetInput("consumerKey").(string))
	consumerSecret := s.TrimSpace(context.GetInput("consumerSecret").(string))
	accessToken := s.TrimSpace(context.GetInput("accessToken").(string))
	accessTokenSecret := s.TrimSpace(context.GetInput("accessTokenSecret").(string))

	tfunction := context.GetInput("function").(string)
	activityLog.Info("Twitter Function: " + tfunction)

	if len(consumerKey) == 0 {

		context.SetOutput("statusCode", 190)
		context.SetOutput("message", "Consumer Key field is blank")

	} else if len(consumerSecret) == 0 {

		context.SetOutput("statusCode", 191)
		context.SetOutput("message", "Consumer Secret field is blank")

	} else if len(accessToken) == 0 {

		context.SetOutput("statusCode", 192)
		context.SetOutput("message", "Access Token field is blank")

	} else if len(accessTokenSecret) == 0 {

		context.SetOutput("statusCode", 193)
		context.SetOutput("message", "Access Token Secret field is blank")

	} else {
		var msg string
		var code int

		switch tfunction {
		case "Tweet":
			{
				tweet := s.TrimSpace(context.GetInput("text").(string))
				if len(tweet) == 0 {

					code = 100
					msg = "Tweet Text cannot be blank"

				} else {
					code, msg = PostTweet(consumerKey, consumerSecret, accessToken, accessTokenSecret, tweet)
				}
			}
		case "TweetMedia":
			{
				tweet := s.TrimSpace(context.GetInput("text").(string))
				mediaURL := s.TrimSpace(context.GetInput("mediaURL").(string))
				if len(tweet) == 0 || len(mediaURL) == 0 {

					code = 101
					msg = "Tweet Text and/or Media URL cannot be blank"

				} else {
					code, msg = PostTweetMedia(consumerKey, consumerSecret, accessToken, accessTokenSecret, tweet, mediaURL)
				}
			}
		case "ReTweet":
			{
				tweetId := context.GetInput("user")
				tid := s.TrimSpace(tweetId.(string))
				if len(tid) == 0 {

					code = 102
					msg = "TweetId cannot be blank"

				} else {
					value, err := strconv.Atoi(tid)

					if err == nil {
						code, msg = ReTweet(consumerKey, consumerSecret, accessToken, accessTokenSecret, int64(value))
					} else {
						code = 103
						msg = err.Error()
					}
				}
			}
		case "Block":
			{
				user := s.TrimSpace(context.GetInput("user").(string))
				if len(user) == 0 {

					code = 104
					msg = "Block user field cannot be blank"

				} else {
					code, msg = BlockUser(consumerKey, consumerSecret, accessToken, accessTokenSecret, user)
				}
			}
		case "Unblock":
			{
				user := s.TrimSpace(context.GetInput("user").(string))
				if len(user) == 0 {

					code = 105
					msg = "UnBlock user field cannot be blank"

				} else {
					code, msg = UnBlockUser(consumerKey, consumerSecret, accessToken, accessTokenSecret, user)
				}
			}
		case "Follow":
			{
				user := s.TrimSpace(context.GetInput("user").(string))
				if len(user) == 0 {

					code = 106
					msg = "Follow user field cannot be blank"

				} else {
					code, msg = Follow(consumerKey, consumerSecret, accessToken, accessTokenSecret, user)
				}
			}
		case "Unfollow":
			{
				user := s.TrimSpace(context.GetInput("user").(string))
				if len(user) == 0 {

					code = 107
					msg = "UnFollow user field cannot be blank"

				} else {
					code, msg = UnFollow(consumerKey, consumerSecret, accessToken, accessTokenSecret, user)
				}
			}
		case "directMessage":
			{
				user := s.TrimSpace(context.GetInput("user").(string))
				directmsg := s.TrimSpace(context.GetInput("text").(string))
				if len(user) == 0 || len(directmsg) == 0 {

					code = 108
					msg = "User and/or Text field cannot be blank"

				} else {
					code, msg = DirectMessage(consumerKey, consumerSecret, accessToken, accessTokenSecret, directmsg, user)
				}
			}

		default:
			{
				code = 150
				msg = "Function field cannot be blank"
			}
		}

		activityLog.Info("Twitter Activity executed with Status Code: " + strconv.Itoa(code))

		context.SetOutput("statusCode", code)
		context.SetOutput("message", msg)
	}

	return true, err
}
