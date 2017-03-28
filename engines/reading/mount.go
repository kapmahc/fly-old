package reading

import (
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Mount web mount-points
func (p *Engine) Mount(rt *gin.Engine) {
	rt.GET("/status", p.Jwt.MustAdminMiddleware, web.JSON(p.getStatus))

	rg := rt.Group("/reading")
	rg.GET("/books", web.JSON(p.indexBooks))
	rg.GET("/books/:id", web.JSON(p.showBook))
	rg.GET("/pages/:id/*href", p.showPage)
	rg.DELETE("/books/:id", p.Jwt.MustAdminMiddleware, web.JSON(p.destroyBook))

	rg.POST("/dict", web.JSON(p.postDict))

	rg.GET("/notes", web.JSON(p.indexNotes))
	rg.POST("/notes", p.Jwt.MustSignInMiddleware, web.JSON(p.createNote))
	rg.POST("/notes/:id", p.Jwt.MustSignInMiddleware, p.canEditNote, web.JSON(p.updateNote))
	rg.DELETE("/notes/:id", p.Jwt.MustSignInMiddleware, p.canEditNote, web.JSON(p.destroyNote))

	mg := rg.Group("/my", p.Jwt.MustSignInMiddleware)
	mg.GET("/notes", p.Jwt.MustSignInMiddleware, web.JSON(p.myNotes))

}
