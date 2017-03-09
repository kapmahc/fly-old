package web

import (
	"net/http"

	gin "gopkg.in/gin-gonic/gin.v1"
)

// TEXT text render
func TEXT(c *gin.Context, s string, e error) {
	if e == nil {
		c.String(http.StatusOK, s)
	} else {
		c.String(http.StatusInternalServerError, e.Error())
	}
}

// JSON json render
func JSON(c *gin.Context, o interface{}, e error) {
	if e == nil {
		if o == nil {
			o = gin.H{}
		}
		c.JSON(http.StatusOK, o)
	} else {
		c.String(http.StatusInternalServerError, e.Error())
	}
}
