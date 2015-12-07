# simpletwitter
The parser of tweets by getting HTML

## Install

```
go get github.com/sosuke-k/simpletwitter
```

## Usage

```
...

import "github.com/sosuke-k/simpletwitter"

...

tweet err := simpletwitter.NewTweet(url)
if err != nil {
  if err.(*simpletwitter.Error).Op = "Redirected" {
    ...
  }
}
fmt.Println(tweet.ScreenName + " : " + tweet.Text)
```
