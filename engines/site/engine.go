package site

import (
	"github.com/facebookgo/inject"
	"github.com/gorilla/mux"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/sky"
	"github.com/kapmahc/sky/cache"
	"github.com/kapmahc/sky/i18n"
	"github.com/kapmahc/sky/job"
	"github.com/kapmahc/sky/settings"
	"golang.org/x/text/language"
	"golang.org/x/tools/blog/atom"
)

// Engine engine
type Engine struct {
	Db         *gorm.DB           `inject:""`
	Settings   *settings.Settings `inject:""`
	CacheStore cache.Store        `inject:""`
	Cache      *cache.Cache       `inject:""`
	I18nStore  i18n.Store         `inject:""`
	I18n       *i18n.I18n         `inject:""`
	Queue      job.Queue          `inject:""`
	Server     *job.Server        `inject:""`
	Jwt        *auth.Jwt          `inject:""`
	Matcher    language.Matcher   `inject:""`
	Router     *mux.Router        `inject:""`
	Layout     *auth.Layout       `inject:""`
	Dao        *auth.Dao          `inject:""`
}

// Map map object
func (p *Engine) Map(*inject.Graph) error {
	return nil
}

// Workers job workers
func (p *Engine) Workers() map[string]job.Handler {
	return map[string]job.Handler{}
}

// Atom rss.atom
func (p *Engine) Atom(lang string) ([]*atom.Entry, error) {
	return []*atom.Entry{}, nil
}

// Sitemap sitemap.xml.gz
func (p *Engine) Sitemap() ([]stm.URL, error) {
	return []stm.URL{}, nil
}

// Application application
func (p *Engine) Application(c *sky.Context) []*sky.Dropdown {
	lang := c.Get(sky.LOCALE).(string)
	return []*sky.Dropdown{
		&sky.Dropdown{Label: p.I18n.T(lang, "site.notices.index.title"), Href: p.Layout.URLFor("site.notices.index")},
	}
}

// Dashboard dashboard
func (p *Engine) Dashboard(c *sky.Context) []*sky.Dropdown {
	lang := c.Get(sky.LOCALE).(string)

	var items []*sky.Dropdown
	if admin, ok := c.Get(auth.IsAdmin).(bool); ok && admin {
		items = append(
			items,
			&sky.Dropdown{
				Label: p.I18n.T(lang, "site.dashboard.title"),
				Links: []*sky.Link{
					&sky.Link{Label: p.I18n.T(lang, "site.admin.status.title"), Href: p.Layout.URLFor("site.admin.status")},
					nil,
					&sky.Link{Label: p.I18n.T(lang, "site.admin.info.title"), Href: p.Layout.URLFor("site.admin.info")},
					&sky.Link{Label: p.I18n.T(lang, "site.admin.author.title"), Href: p.Layout.URLFor("site.admin.author")},
					&sky.Link{Label: p.I18n.T(lang, "site.admin.seo.title"), Href: p.Layout.URLFor("site.admin.seo")},
					&sky.Link{Label: p.I18n.T(lang, "site.admin.smtp.title"), Href: p.Layout.URLFor("site.admin.smtp")},
					nil,
					&sky.Link{Label: p.I18n.T(lang, "site.admin.users.index.title"), Href: p.Layout.URLFor("site.admin.users.index")},
					&sky.Link{Label: p.I18n.T(lang, "site.admin.notices.index.title"), Href: p.Layout.URLFor("site.admin.notices.index")},
					&sky.Link{Label: p.I18n.T(lang, "site.admin.leave-words.index.title"), Href: p.Layout.URLFor("site.admin.leave-words.index")},
					&sky.Link{Label: p.I18n.T(lang, "site.admin.locales.index.title"), Href: p.Layout.URLFor("site.admin.locales.index")},
				},
			},
		)
	}
	return items
}

func init() {
	sky.Register(&Engine{})
}
