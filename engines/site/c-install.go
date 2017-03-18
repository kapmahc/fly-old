package site

import (
	"net/http"

	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) getInstall(c *gin.Context) {
	c.HTML(http.StatusOK, "site-install", gin.H{})
}
func (p *Engine) postInstall(c *gin.Context) {

}
