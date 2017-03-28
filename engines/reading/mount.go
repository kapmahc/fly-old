package reading

import (
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Mount web mount-points
func (p *Engine) Mount(rt *gin.Engine) {
	rg := rt.Group("/reading")
	rg.GET("/books", web.JSON(p.indexBooks))
	rg.GET("/books/:id", web.JSON(p.showBook))
	rg.GET("/pages/:id/*href", p.showPage)
	rg.GET("/dict", web.JSON(p.formDict))
	rg.POST("/dict", web.JSON(p.formDict))

	rg.GET("/notes", web.JSON(p.indexNotes))
	rg.GET("/notes/new", p.Jwt.MustSignInMiddleware, web.JSON(p.createNote))
	rg.POST("/notes/new", p.Jwt.MustSignInMiddleware, web.JSON(p.createNote))
	rg.GET("/notes/edit/:id", p.Jwt.MustSignInMiddleware, p.canEditNote, web.JSON(p.updateNote))
	rg.POST("/notes/edit/:id", p.Jwt.MustSignInMiddleware, p.canEditNote, web.JSON(p.updateNote))
	rg.DELETE("/notes/:id", p.Jwt.MustSignInMiddleware, p.canEditNote, web.JSON(p.destroyNote))
	rg.GET("/notes/my", p.Jwt.MustSignInMiddleware, web.JSON(p.myNotes))

	ag := rg.Group("/admin", p.Jwt.MustAdminMiddleware)
	ag.GET("/status", web.JSON(p.getAdminStatus))
	ag.GET("/books", web.JSON(p.indexAdminBooks))
	ag.DELETE("/books/:id", web.JSON(p.destroyAdminBook))
}
