package base

import (
	"errors"
	"fmt"
	"html/template"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/beego/i18n"
)

const (
	// LOCALE locale key
	LOCALE = "locale"

	localeDataKey = "l"
)

// Controller base
type Controller struct {
	beego.Controller
}

// Abort abort
func (p *Controller) Abort(s int) {
	p.Controller.Abort(strconv.Itoa(s))
}

// ParseForm parse form
func (p *Controller) ParseForm(fm interface{}) (ok bool, fsh *beego.FlashData) {
	fsh = beego.NewFlash()
	err := p.Controller.ParseForm(fm)
	if err != nil {
		fsh.Error(err.Error())
		return
	}
	valid := validation.Validation{}
	ok, err = valid.Valid(fm)
	if err != nil {
		fsh.Error(err.Error())
		return
	}
	if !ok {
		var msg []string
		for _, err := range valid.Errors {
			msg = append(msg, fmt.Sprintf("%s: %s", err.Key, err.Message))
		}
		fsh.Error(strings.Join(msg, "<br/>"))
		return
	}
	ok = true
	return
}

// E i18n error
func (p *Controller) E(format string, args ...interface{}) error {
	return errors.New(p.T(format, args...))
}

// T t
func (p *Controller) T(code string, args ...interface{}) string {
	beego.Debug("T", code)
	beego.Debug("DATA", p.Data)
	return Tr(p.Data[localeDataKey].(string), code, args...)
}

// Prepare prepare
func (p *Controller) Prepare() {
	p.Layout = "application.html"
	p.setLang()
	p.setXSRF()
	p.setEngines()
}

func (p *Controller) setEngines() {
	p.Data["engines"] = beego.AppConfig.Strings("engines")
}

// HTML render html
func (p *Controller) HTML(title, tpl string) {
	beego.ReadFromRequest(&p.Controller)
	p.TplName = tpl
	p.Data["title"] = title
}

func (p *Controller) setXSRF() {
	p.Data["xsrf"] = template.HTML(p.XSRFFormHTML())
	p.Data["xsrf_token"] = p.XSRFToken()
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
