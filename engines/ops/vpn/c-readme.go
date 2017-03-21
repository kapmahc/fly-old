package vpn

import gin "gopkg.in/gin-gonic/gin.v1"

func (p *Engine) getReadme(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "vpn.logs.index.title")
	tpl := "vpn-readme"

	return tpl, nil
}
