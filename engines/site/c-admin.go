package site

import (
	"net/http"

	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

type fmSiteInfo struct {
	Title       string `form:"title"`
	SubTitle    string `form:"subTitle"`
	Keywords    string `form:"keywords"`
	Description string `form:"description"`
	Copyright   string `form:"copyright"`
}

func (p *Engine) formAdminSiteInfo(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "site.admin.info.title")
	tpl := "site-admin-info"
	if c.Request.Method == http.MethodPost {
		var fm fmSiteInfo
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}

		for k, v := range map[string]string{
			"title":       fm.Title,
			"subTitle":    fm.SubTitle,
			"keywords":    fm.Keywords,
			"description": fm.Description,
			"copyright":   fm.Copyright,
		} {
			if err := p.I18n.Set(lang, "site."+k, v); err != nil {
				return tpl, err
			}
		}
		data[web.NOTICE] = p.I18n.T(lang, "success")
	}

	return tpl, nil
}

type fmSiteAuthor struct {
	Name  string `form:"name"`
	Email string `form:"email"`
}

func (p *Engine) formAdminSiteAuthor(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "site.admin.author.title")
	tpl := "site-admin-author"
	if c.Request.Method == http.MethodPost {

		var fm fmSiteAuthor
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}

		for k, v := range map[string]string{
			"name":  fm.Name,
			"email": fm.Email,
		} {
			if err := p.I18n.Set(lang, "site.author."+k, v); err != nil {
				return tpl, err
			}
		}
		data[web.NOTICE] = p.I18n.T(lang, "success")
	}
	return tpl, nil
}

type fmSiteSeo struct {
	GoogleVerifyCode string `form:"googleVerifyCode"`
	BaiduVerifyCode  string `form:"baiduVerifyCode"`
}

func (p *Engine) formAdminSiteSeo(c *gin.Context, lang string, data gin.H) (string, error) {

	data["title"] = p.I18n.T(lang, "site.admin.seo.title")
	tpl := "site-admin-seo"
	if c.Request.Method == http.MethodPost {
		var fm fmSiteSeo
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}

		for k, v := range map[string]string{
			"google.verify.code": fm.GoogleVerifyCode,
			"baidu.verify.code":  fm.BaiduVerifyCode,
		} {
			if err := p.Settings.Set("site."+k, v, true); err != nil {
				return tpl, err
			}
		}
		data[web.NOTICE] = p.I18n.T(lang, "success")
	}

	var gc string
	var bc string
	p.Settings.Get("site.google.verify.code", &gc)
	p.Settings.Get("site.baidu.verify.code", &bc)
	data["googleVerifyCode"] = gc
	data["baiduVerifyCode"] = bc
	return tpl, nil
}

type fmSiteSMTP struct {
	Host                 string `form:"host"`
	Port                 int    `form:"port"`
	Ssl                  bool   `form:"ssl"`
	Username             string `form:"username"`
	Password             string `form:"password" binding:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=Password"`
}

func (p *Engine) formAdminSiteSMTP(c *gin.Context, lang string, data gin.H) (string, error) {

	data["title"] = p.I18n.T(lang, "site.admin.smtp.title")
	tpl := "site-admin-smtp"
	if c.Request.Method == http.MethodPost {
		var fm fmSiteSMTP
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}

		val := map[string]interface{}{
			"host":     fm.Host,
			"port":     fm.Port,
			"username": fm.Username,
			"password": fm.Password,
			"ssl":      fm.Ssl,
		}
		if err := p.Settings.Set("site.smtp", val, true); err != nil {
			return tpl, err
		}
		data[web.NOTICE] = p.I18n.T(lang, "success")
	}

	smtp := make(map[string]interface{})
	if err := p.Settings.Get("site.smtp", &smtp); err == nil {
		smtp["password"] = ""
	} else {
		smtp["host"] = "localhost"
		smtp["port"] = 25
		smtp["ssl"] = false
		smtp["username"] = "no-reply@change-me.com"
		smtp["password"] = ""
	}
	data["smtp"] = smtp
	data["ports"] = []int{25, 465, 587, 2525, 2526}
	return tpl, nil
}
