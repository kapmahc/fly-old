package vpn

import (
	"net/http"

	"github.com/SermoDigital/jose/jws"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// TokenMiddleware token-middleware
func (p *Engine) TokenMiddleware(c *gin.Context) {
	tk, err := jws.ParseJWTFromRequest(c.Request)
	if err == nil {
		err = tk.Validate(p.Key, p.Method)
	}
	if err == nil {
		err = tk.Validate(p.Key, p.Method)
	}
	if err == nil {
		if act := tk.Claims().Get("act"); act != nil && act.(string) == "vpn" {
			c.Next()
			return
		}
	}
	c.AbortWithStatus(http.StatusForbidden)
}
