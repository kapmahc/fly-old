package vpn

import (
	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/web"
	"github.com/spf13/viper"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) getReadme(c *gin.Context) (interface{}, error) {
	data["title"] = p.I18n.T(lang, "ops.vpn.readme.title")
	tpl := "ops-vpn-readme"
	data["user"] = c.MustGet(auth.CurrentUser)
	data["name"] = viper.Get("server.name")
	data["home"] = web.Home()
	data["port"] = 1194
	data["network"] = "10.18.0"

	token, err := p.generateToken(10)
	if err != nil {
		return nil, err
	}
	data["token"] = string(token)
	return tpl, nil
}
