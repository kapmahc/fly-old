package vpn

import (
	"net/http"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

//Jwt jwt helper
type Jwt struct {
	Key    []byte               `inject:"jwt.key"`
	Method crypto.SigningMethod `inject:"jwt.method"`
	I18n   *web.I18n            `inject:""`
}

// Middleware middleware
func (p *Jwt) Middleware(c *gin.Context) {
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
