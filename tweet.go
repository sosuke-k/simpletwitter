package simpletwitter

import (
	"errors"
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

// Tweet is tweet object
type Tweet struct {
	ItemID     string
	ScreenName string
	Name       string
	Time       string
	Text       string
	Before     []Tweet
	After      []Tweet
}

// NewTweet returns tweet object by url
func NewTweet(url string) (tweet Tweet, err error) {

	var (
		doc *goquery.Document
	)

	doc, err = goquery.NewDocument(url)
	if err != nil {
		err = &Error{
			Op:     "Request",
			GetURL: url,
			Err:    err,
		}
		return
	}

	if url != doc.Url.String() {
		err = &Error{
			Op:            "Redirected",
			GetURL:        url,
			RedirectedURL: doc.Url.String(),
		}
		return
	}

	tweet, err = parse(doc.Find(".permalink-tweet-container .tweet"))
	if err != nil {
		if notExists := checkNotExisting(doc); notExists {
			err = errors.New("this page not exists")
			err = &Error{Op: "Parse", GetURL: url, RedirectedURL: doc.Url.String(), Err: err}
		} else {
			err = &Error{Op: "Parse", GetURL: url, RedirectedURL: doc.Url.String(), Err: err}
		}
		return
	}

	return
}

func parse(s *goquery.Selection) (tweet Tweet, err error) {
	success := false
	attrs := []string{
		"data-item-id",
		"data-screen-name",
		"data-name",
	}
	data := map[string]string{}

	for _, attr := range attrs {
		var value string
		if value, success = s.Attr(attr); !success {
			err = fmt.Errorf("not having %s attribute", attr)
			return
		}
		data[attr] = value
	}

	tweet = Tweet{}
	tweet.ItemID = data["data-item-id"]
	tweet.ScreenName = data["data-screen-name"]
	tweet.Name = data["data-name"]

	// if could get the above attribues, allow the following values to be blank.
	tweet.Time, _ = s.Find("._timestamp").Attr("data-time")
	tweet.Text = s.Find(".tweet-text").Text()
	return
}

func checkNotExisting(doc *goquery.Document) (notExisting bool) {
	if doc.Find(".body-content h1").Text() == "Sorry, that page doesn’t exist!" {
		notExisting = true
	}
	return
}
