package auth

import (
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/kapmahc/fly/web"
	"github.com/spf13/viper"

	gin "gopkg.in/gin-gonic/gin.v1"
)

// HTML html render
func HTML(f func(*gin.Context, string, gin.H) (string, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		l := c.MustGet(web.LOCALE).(string)
		d := gin.H{}
		v, e := f(c, l, d)
		if e != nil {
			d[web.ERROR] = e.Error()
		}
		d["l"] = l
		d["languages"] = viper.GetStringSlice("languages")
		d[csrf.TemplateTag] = csrf.TemplateField(c.Request)
		c.Writer.Header().Set("X-CSRF-Token", csrf.Token(c.Request))

		if v != "" {
			c.HTML(http.StatusOK, v, d)
		}
	}
}
