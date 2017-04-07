package auth

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/kapmahc/sky"
	"github.com/spf13/viper"
)

const (
	navLinks = "links"
)

// Layout layout
type Layout struct {
	Router *mux.Router `inject:""`
}

// URLFor builds a url for the route.
func (p *Layout) URLFor(name string, pairs ...interface{}) string {

	rt := p.Router.Get(name)
	if rt == nil {
		return name
	}
	var params []string
	for _, v := range pairs {
		switch t := v.(type) {
		case string:
			params = append(params, v.(string))
		default:
			log.Warn("unknown type", t)
		}
	}
	url, err := rt.URL(params...)
	if err != nil {
		log.Error(err)
		return name
	}
	return url.String()
}

// Application application-layout
func (p *Layout) Application(c *sky.Context) error {
	var items []*sky.Dropdown
	sky.Walk(func(en sky.Engine) error {
		tmp := en.Application(c)
		items = append(items, tmp...)
		return nil
	})
	return p.payload(c, items)
}

// Dashboard dashboard-layout
func (p *Layout) Dashboard(c *sky.Context) error {
	var items []*sky.Dropdown
	sky.Walk(func(en sky.Engine) error {
		tmp := en.Dashboard(c)
		items = append(items, tmp...)
		return nil
	})
	return p.payload(c, items)
}

func (p *Layout) payload(c *sky.Context, links []*sky.Dropdown) error {
	c.Set(sky.DATA, sky.H{
		"navbar":    links,
		CurrentUser: c.Get(CurrentUser),
		"l":         c.Get(sky.LOCALE),
		"csrf":      csrf.Token(c.Request),
		"languages": viper.GetStringSlice("languages"),
	})
	return c.Next()

}
