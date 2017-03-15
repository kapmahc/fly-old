package auth

import (
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Mount web mount-points
func (p *Engine) Mount(rt *gin.Engine) {
	rt.GET("/users", web.JSON(p.indexUsers))

	rt.GET("/attachments", p.Jwt.MustSignInMiddleware, web.JSON(p.indexAttachments))
	rt.POST("/attachments", p.Jwt.MustSignInMiddleware, web.JSON(p.createAttachment))
	rt.GET("/attachments/:id", web.JSON(p.showAttachment))
	rt.POST("/attachments/:id", p.Jwt.MustSignInMiddleware, p.canEditAttachment, web.JSON(p.updateAttachment))
	rt.DELETE("/attachments/:id", p.Jwt.MustSignInMiddleware, p.canEditAttachment, web.JSON(p.destroyAttachment))

	ung := rt.Group("/users")
	ung.POST("/sign-in", web.JSON(p.postUsersSignIn))
	ung.POST("/sign-up", web.JSON(p.postUsersSignUp))
	ung.GET("/confirm/:token", web.Redirect(p.getUsersConfirm))
	ung.POST("/confirm", web.JSON(p.postUsersConfirm))
	ung.GET("/unlock/:token", web.Redirect(p.getUsersUnlock))
	ung.POST("/unlock", web.JSON(p.postUsersUnlock))
	ung.POST("/forgot-password", web.JSON(p.postUsersForgotPassword))
	ung.POST("/reset-password/:token", web.JSON(p.postUsersResetPassword))

	umg := rt.Group("/users", p.Jwt.MustSignInMiddleware)
	umg.POST("/info", web.JSON(p.postUsersInfo))
	umg.GET("/info", web.JSON(p.getUsersInfo))
	umg.POST("/change-password", web.JSON(p.postUsersChangePassword))
	umg.GET("/logs", web.JSON(p.getUsersLogs))
	umg.DELETE("/sign-out", web.JSON(p.deleteUsersSignOut))

}
