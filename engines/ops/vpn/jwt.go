package vpn

import (
	"net/http"
	"time"

	"github.com/SermoDigital/jose/jws"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) generateToken(years int) ([]byte, error) {
	now := time.Now()
	cm := jws.Claims{}
	cm.SetNotBefore(now)
	cm.SetExpiration(now.AddDate(years, 0, 0))
	cm.Set("act", "vpn")

	jt := jws.NewJWT(cm, p.Method)
	return jt.Serialize(p.Key)
}

func (p *Engine) tokenMiddleware(c *gin.Context) {
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
