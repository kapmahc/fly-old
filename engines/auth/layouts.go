package auth

import (
	"github.com/kapmahc/sky"
	"github.com/kapmahc/sky/i18n"
	"github.com/spf13/viper"
)

const (
	navLinks = "links"
)

// Layout layout
type Layout struct {
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
		"links":     links,
		CurrentUser: c.Get(CurrentUser),
		"lang":      c.Get(i18n.LOCALE),
		"languages": viper.GetStringSlice("languages"),
	})
	return c.Next()

}
