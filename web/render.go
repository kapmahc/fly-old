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
func Redirect(c *gin.Context, u string, err error) {
	if err == nil {
		c.Redirect(http.StatusFound, u)
	} else {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
}

// JSON json render
func JSON(c *gin.Context, data interface{}, err error) {
	if err == nil {
		if data == nil {
			data = gin.H{}
		}
		c.JSON(http.StatusOK, data)
	} else {
		// c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		c.AbortWithError(http.StatusInternalServerError, err)
	}
}
