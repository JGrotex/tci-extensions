# TCI Twitter Extension
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

This activity allows you to:
- Tweet, Retweet a particular Tweet,
- Follow/Unfollow a User on Twitter,
- Block/Unblock a User on Twitter, 
- DirectMessage a user on Twitter,

## Parameter Details
| Setting           | Required  | Description                         |
|:------------------|:----------|:------------------------------------|
| consumerKey       | true      | consumerKey of your Twitter account |         
| consumerSecret    | true      | consumerSecret of your Twitter account |
| accessToken       | true      | accessToken of your Twitter account |
| accessTokenSecret | true      | accessTokenSecret of your Twitter account |
| Function          | true      | Actions you want to perform, the possible values are "Tweet", "TweetMedia","ReTweet","Block","Unblock","Follow","Unfollow","DirectMessage" |
| User              | false     | Use this field to provide tweetId for ReTweet, User for Block, UnBlock, Follow, UnFollow, DirectMessage |
| Text              | false     | Use this field to provide an text for you Tweet or Direct Message |
| MediaURL          | false     | Use this field to provide an external Image Media URL |

Note: 
You can generate your Twitter Keys and Tokens from here, 
- direct Link: https://apps.twitter.com/app/new 
- Create your Twitter Application, and
- click the Create my Access Token button

## Third-party libraries used
package anaconda - "github.com/ChimeraCoder/anaconda" :: Thanks to Anaconda it is simple accessing version 1.1 of the Twitter API.

## Thanks
This Implementation is inspired and based on a TIBCO PSG Hackathon Source Implementation, thanks to FLOGO AllStars team.