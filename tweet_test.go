package simpletwitter_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/sosuke-k/simpletwitter"
)

func TestNewTweet(t *testing.T) {
	Convey("Given available first post tweet url\n", t, func() {
		url := "https://twitter.com/olha_drm/status/418033807850496002"
		tweet, err := simpletwitter.NewTweet(url)

		So(err, ShouldEqual, nil)
		So(tweet.Text, ShouldEqual, "よるほ　あけましておめでとうございますほー")
		So(tweet.ScreenName, ShouldEqual, "olha_drm")
		So(tweet.Name, ShouldEqual, "織羽")
		// So(len(tweet.After), ShouldEqual, 1)
		// So(tweet.After[0].Text, ShouldEqual, "@olha_drm あけおめっ！ことよろー！！")
		// So(tweet.After[0].ScreenName, ShouldEqual, "reprohonmono")
	})

	Convey("Given non-available tweet url\n", t, func() {
		url := "https://twitter.com/statuses/418033823511629836"
		_, err := simpletwitter.NewTweet(url)

		So(err, ShouldNotEqual, nil)
		So(err.(*simpletwitter.Error).Op, ShouldEqual, "Redirected")
	})
}
