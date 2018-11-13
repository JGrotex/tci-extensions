/*
TCI Flogo twitter package is using github.com/ChimeraCoder/anaconda
*/
package tweet

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/ChimeraCoder/anaconda"
)

// function to Tweet with string messages
func PostTweet(consumerKey, consumerSecret, accessToken, accessTokenSecret, tweet string) (statusCode int, message string) {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
	twt, err := api.PostTweet(tweet, nil)
	if err != nil {
		abc := err.(*anaconda.ApiError)
		return abc.StatusCode, abc.Error()
	}
	return 200, twt.IdStr
}

// function to Tweet with string messages
func PostTweetMedia(consumerKey, consumerSecret, accessToken, accessTokenSecret, tweet string, mediaURL string) (statusCode int, message string) {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)

	//read Image from URL
	res, err := http.Get(mediaURL)
	if err != nil {
		return 133, err.Error()
	}
	defer res.Body.Close()
	mediadata, _ := ioutil.ReadAll(res.Body)

	//read Image from local Filesystem
	// mediadata, err := ioutil.ReadFile(mediaURL)
	// if err != nil {
	// 	return 133, err.Error()
	// }

	//newTwitterApi
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	//upload Media
	mediaResponse, err := api.UploadMedia(base64.StdEncoding.EncodeToString(mediadata))
	if err != nil {
		fmt.Println(err)
	}

	v := url.Values{}
	v.Set("media_ids", strconv.FormatInt(mediaResponse.MediaID, 10))

	// Tweet
	twt, err := api.PostTweet(tweet, v)
	if err != nil {
		abc := err.(*anaconda.ApiError)
		return abc.StatusCode, abc.Error()
	}
	return 200, twt.IdStr
}

// function to retweet a particular tweet
func ReTweet(consumerKey, consumerSecret, accessToken, accessTokenSecret string, tweetId int64) (statusCode int, message string) {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
	twt, err := api.Retweet(tweetId, false)
	if err != nil {
		abc := err.(*anaconda.ApiError)
		return abc.StatusCode, abc.Error()
	}
	return 200, twt.IdStr
}

// function to DirectMessage an user
func DirectMessage(consumerKey, consumerSecret, accessToken, accessTokenSecret, directmsg, user string) (statusCode int, message string) {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
	twt, err := api.PostDMToScreenName(directmsg, user)
	if err != nil {
		abc := err.(*anaconda.ApiError)
		return abc.StatusCode, abc.Error()
	}
	return 200, twt.IdStr
}

// function to Block a user on Twitter
func BlockUser(consumerKey, consumerSecret, accessToken, accessTokenSecret, twitterHandle string) (statusCode int, message string) {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
	resp, err := api.BlockUser(twitterHandle, nil)
	if err != nil {
		abc := err.(*anaconda.ApiError)
		return abc.StatusCode, abc.Error()
	}
	return 200, resp.IdStr
}

// function to UnBlock a user on Twitter.
func UnBlockUser(consumerKey, consumerSecret, accessToken, accessTokenSecret, twitterHandle string) (statusCode int, message string) {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
	resp, err := api.UnblockUser(twitterHandle, nil)
	if err != nil {
		abc := err.(*anaconda.ApiError)
		return abc.StatusCode, abc.Error()
	}
	return 200, resp.IdStr
}

// function to Follow a user on Twitter
func Follow(consumerKey, consumerSecret, accessToken, accessTokenSecret, twitterHandle string) (statusCode int, message string) {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
	resp, err := api.FollowUser(twitterHandle)
	if err != nil {
		abc := err.(*anaconda.ApiError)
		return abc.StatusCode, abc.Error()
	}
	return 200, resp.IdStr
}

// function to UnFollow a user on Twitter
func UnFollow(consumerKey, consumerSecret, accessToken, accessTokenSecret, twitterHandle string) (statusCode int, message string) {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
	resp, err := api.UnfollowUser(twitterHandle)
	if err != nil {
		abc := err.(*anaconda.ApiError)
		return abc.StatusCode, abc.Error()
	}
	return 200, resp.IdStr
}
