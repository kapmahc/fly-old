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
		// ---------
		d["l"] = l
		d["languages"] = viper.GetStringSlice("languages")
		// ---------
		var ds []*web.Dropdown
		web.Walk(func(en web.Engine) error {
			ds = append(ds, en.Dashboard(c))
			return nil
		})
		d["dashboard"] = ds
		// ---------
		d[csrf.TemplateTag] = csrf.TemplateField(c.Request)
		token := csrf.Token(c.Request)
		d["csrf"] = token
		c.Writer.Header().Set("X-CSRF-Token", token)
		// ---------
		if user, ok := c.Get(CurrentUser); ok {
			d[CurrentUser] = user
		}
		if admin, ok := c.Get(IsAdmin); ok {
			d[IsAdmin] = admin
		}
		// ---------
		if v != "" {
			c.HTML(http.StatusOK, v, d)
		}
	}
}
