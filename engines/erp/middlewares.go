package erp

import (
	"net/http"

	"github.com/kapmahc/fly/engines/auth"
	"github.com/spf13/viper"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) isSeller(c *gin.Context) bool {
	user, ok := c.Get(auth.CurrentUser)
	if !ok {
		return false
	}
	if viper.GetString("erp.mode") == "public" {
		return true
	}
	if admin, ok := c.Get(auth.IsAdmin); ok && admin.(bool) {
		return true
	}
	if p.Dao.Is(user.(*auth.User).ID, RoleSeller) {
		return true
	}
	return false
}
func (p *Engine) mustSellerMiddleware(c *gin.Context) {
	if p.isSeller(c) {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusForbidden)
	}
}

func init() {
	viper.SetDefault("erp.mode", "private")
}
