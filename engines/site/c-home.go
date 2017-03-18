package site

import gin "gopkg.in/gin-gonic/gin.v1"

func (p *Engine) getDashboard(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "header.dashboard")
	return "site-dashboard", nil
}
