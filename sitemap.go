package fly

//https://www.sitemaps.org/protocol.html

import (
	"encoding/xml"
	"io"
	"time"
)

const (
	// ALWAYS always
	ALWAYS = "always"
	// HOURLY hourly
	HOURLY = "hourly"
	// DAILY daily
	DAILY = "daily"
	// WEEKLY weekly
	WEEKLY = "weekly"
	// MONTHLY monthly
	MONTHLY = "monthly"
	// YEARLY yearly
	YEARLY = "yearly"
	// NEVER never
	NEVER = "never"
)

// URL sitemap.xml url
type URL struct {
	Loc        string    `xml:"loc"`
	LastMod    time.Time `xml:"lastmod"`
	ChangeFreq string    `xml:"changefreq"`
	Priority   float32   `xml:"priority"`
}

// URLSet sitemap.xml urlset
type URLSet struct {
	XMLName xml.Name `xml:"http://www.sitemaps.org/schemas/sitemap/0.9 urlset"`
	URL     []*URL   `xml:"url"`
}

// Sitemap sitemap.xml
func Sitemap(w io.Writer, u *URLSet) error {
	enc := xml.NewEncoder(w)
	return enc.Encode(u)
}
