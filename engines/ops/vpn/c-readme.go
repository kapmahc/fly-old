package vpn

import (
	"net/http"

	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/web"
	"github.com/spf13/viper"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) getReadme(c *gin.Context) {

	data := gin.H{}
	data["user"] = c.MustGet(auth.CurrentUser)
	data["name"] = viper.Get("server.name")
	data["home"] = web.Home()
	data["port"] = 1194
	data["network"] = "10.18.0"

	token, err := p.generateToken(10)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	data["token"] = string(token)

	c.HTML(http.StatusOK, "ops.vpn.readme", data)
}
