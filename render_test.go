package fly_test

import (
	"os"
	"testing"
	"time"

	"github.com/kapmahc/fly"
	"golang.org/x/tools/blog/atom"
)

var rdr = fly.Render{Pretty: true}

func TestAtom(t *testing.T) {
	feed := atom.Feed{Title: "title", Updated: atom.Time(time.Now())}
	rdr.XML(os.Stdout, &feed)
}

func TestSitemap(t *testing.T) {
	var us fly.URLSet
	us.URL = append(us.URL,
		&fly.URL{
			Loc:        "http://www.change-me.com",
			LastMod:    time.Now(),
			ChangeFreq: fly.DAILY,
			Priority:   0.9,
		}, &fly.URL{
			Loc:        "http://www.change-me.com/about",
			LastMod:    time.Now(),
			ChangeFreq: fly.WEEKLY,
			Priority:   0.2,
		},
	)
	rdr.XML(os.Stdout, &us)
}
