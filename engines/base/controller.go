package base

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

const (
	// LOCALE locale key
	LOCALE        = "locale"
	localeDataKey = "l"
)

// Controller base
type Controller struct {
	beego.Controller
}

// T t
func (p *Controller) T(code string, args ...interface{}) string {
	return Tr(p.Data[localeDataKey].(string), code, args...)
}

// Prepare prepare
func (p *Controller) Prepare() {
	p.Layout = "application.html"
	p.setLang()
}
func (p *Controller) setLang() {

	hasCookie := false

	// 1. Check URL arguments.
	lang := p.Input().Get(LOCALE)

	// 2. Get language information from cookies.
	if len(lang) == 0 {
		lang = p.Ctx.GetCookie(LOCALE)
		hasCookie = true
	}
	// 3. Get language information from 'Accept-Language'.
	if len(lang) == 0 {
		al := p.Ctx.Request.Header.Get("Accept-Language")
		if len(al) > 4 {
			al = al[:5] // Only compare first 5 letters.
			if i18n.IsExist(al) {
				lang = al
			}
		}
	}

	// Default language
	if !i18n.IsExist(lang) {
		lang = beego.AppConfig.String("defaultlocale")
		hasCookie = false
	}

	// Save language information in cookies.
	if !hasCookie {
		p.Ctx.SetCookie(LOCALE, lang, 1<<31-1, "/")
	}

	// Set language properties.
	p.Data[localeDataKey] = lang
	p.Data["languages"] = i18n.ListLangs()
}
