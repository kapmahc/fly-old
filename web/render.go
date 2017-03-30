package web

import (
	"net/http"

	gin "gopkg.in/gin-gonic/gin.v1"
)

// JSON json render
func JSON(f func(*gin.Context) (interface{}, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		v, e := f(c)
		if e == nil {
			c.JSON(http.StatusOK, v)
		} else {
			c.String(http.StatusInternalServerError, e.Error())
		}

	}
}
