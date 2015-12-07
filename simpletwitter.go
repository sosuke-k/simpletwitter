/*
Package simpletwitter uses goquery(https://github.com/PuerkitoBio/goquery) to fetch tweets.

Example:

  tweet err := simpletwitter.NewTweet(url)
  if err != nil {
    if err.(*simpletwitter.Error).Op = "Redirected" {
      ...
    }
  }
  fmt.Println(tweet.ScreenName + " : " + tweet.Text)

*/
package simpletwitter
