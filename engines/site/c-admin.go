package site

import (
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

func (p *Engine) postAdminSiteInfo(c *gin.Context) (interface{}, error) {
	lang := c.MustGet(web.LOCALE).(string)
	var fm fmSiteInfo
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	for k, v := range map[string]string{
		"title":       fm.Title,
		"subTitle":    fm.SubTitle,
		"keywords":    fm.Keywords,
		"description": fm.Description,
		"copyright":   fm.Copyright,
	} {
		if err := p.I18n.Set(lang, "site."+k, v); err != nil {
			return nil, err
		}
	}
	return gin.H{}, nil
}

type fmSiteAuthor struct {
	Name  string `form:"name"`
	Email string `form:"email"`
}

func (p *Engine) postAdminSiteAuthor(c *gin.Context) (interface{}, error) {
	lang := c.MustGet(web.LOCALE).(string)
	var fm fmSiteAuthor
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	for k, v := range map[string]string{
		"name":  fm.Name,
		"email": fm.Email,
	} {
		if err := p.I18n.Set(lang, "site.author."+k, v); err != nil {
			return nil, err
		}
	}
	return gin.H{}, nil
}

type fmSiteSeo struct {
	GoogleVerifyCode string `form:"googleVerifyCode"`
	BaiduVerifyCode  string `form:"baiduVerifyCode"`
}

func (p *Engine) getAdminSiteSeo(c *gin.Context) (interface{}, error) {
	var gc string
	var bc string
	p.Settings.Get("site.google.verify.code", &gc)
	p.Settings.Get("site.baidu.verify.code", &bc)
	return gin.H{
		"googleVerifyCode": gc,
		"baiduVerifyCode":  bc,
	}, nil
}

func (p *Engine) postAdminSiteSeo(c *gin.Context) (interface{}, error) {

	var fm fmSiteSeo
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	for k, v := range map[string]string{
		"google.verify.code": fm.GoogleVerifyCode,
		"baidu.verify.code":  fm.BaiduVerifyCode,
	} {
		if err := p.Settings.Set("site."+k, v, true); err != nil {
			return nil, err
		}
	}

	return gin.H{}, nil
}

type fmSiteSMTP struct {
	Host                 string `form:"host"`
	Port                 int    `form:"port"`
	Ssl                  bool   `form:"ssl"`
	Username             string `form:"username"`
	Password             string `form:"password" binding:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=Password"`
}

func (p *Engine) getAdminSiteSMTP(c *gin.Context) (interface{}, error) {

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

	return smtp, nil

}
func (p *Engine) postAdminSiteSMTP(c *gin.Context) (interface{}, error) {

	var fm fmSiteSMTP
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	val := map[string]interface{}{
		"host":     fm.Host,
		"port":     fm.Port,
		"username": fm.Username,
		"password": fm.Password,
		"ssl":      fm.Ssl,
	}
	if err := p.Settings.Set("site.smtp", val, true); err != nil {
		return nil, err
	}

	return gin.H{}, nil
}
