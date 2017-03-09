package web

import (
	"net/http"

	gin "gopkg.in/gin-gonic/gin.v1"
)

// TEXT text render
func TEXT(c *gin.Context, s string, err error) {
	if err == nil {
		c.String(http.StatusOK, s)
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
}

// JSON json render
func JSON(c *gin.Context, o interface{}, err error) {
	if err == nil {
		if o == nil {
			o = gin.H{}
		}
		c.JSON(http.StatusOK, o)
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
}
