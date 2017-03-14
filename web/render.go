package web

import (
	"net/http"

	gin "gopkg.in/gin-gonic/gin.v1"
)

// // TEXT text render
// func TEXT(c *gin.Context, s string, err error) {
// 	if err == nil {
// 		c.String(http.StatusOK, s)
// 	} else {
// 		log.Error(err)
// 		c.String(http.StatusInternalServerError, err.Error())
// 	}
// }

// Redirect redirect
func Redirect(f func(*gin.Context) (u string, e error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		u, e := f(c)
		if e != nil {
			c.String(http.StatusInternalServerError, e.Error())
			return
		}
		c.Redirect(http.StatusFound, u)
	}
}

// JSON json render
func JSON(f func(*gin.Context) (interface{}, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		v, e := f(c)
		if c != nil {
			c.String(http.StatusInternalServerError, e.Error())
			return
		}
		c.JSON(http.StatusOK, v)
	}
}
