package mail

import (
	"net/http"

	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) getReadme(c *gin.Context) {
	c.HTML(http.StatusOK, "ops.mail.readme", gin.H{})
}
