package auth

import (
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Mount web mount-points
func (p *Engine) Mount(rt *gin.Engine) {
	rt.GET("/users", web.JSON(p.indexUsers))

	ung := rt.Group("/users")

	ung.POST("/sign-in", web.JSON(p.postUsersSignIn))
	ung.POST("/sign-up", web.JSON(p.postUsersSignUp))
	ung.GET("/confirm/:token", web.JSON(p.getUsersConfirm))
	ung.POST("/confirm", web.JSON(p.postUsersConfirm))
	ung.GET("/unlock/:token", web.JSON(p.getUsersUnlock))
	ung.POST("/unlock", web.JSON(p.postUsersUnlock))
	ung.POST("/forgot-password", web.JSON(p.postUsersForgotPassword))
	ung.POST("/reset-password", web.JSON(p.postUsersResetPassword))

	umg := rt.Group("/users", p.Jwt.MustSignInMiddleware)
	umg.POST("/info", web.JSON(p.postUsersInfo))
	umg.POST("/change-password", web.JSON(p.postUsersChangePassword))
	umg.GET("/logs", web.JSON(p.getUsersLogs))
	umg.DELETE("/sign-out", web.JSON(p.deleteUsersSignOut))

	rt.GET("/attachments", p.Jwt.MustSignInMiddleware, web.JSON(p.indexAttachments))
	rt.POST("/attachments", p.Jwt.MustSignInMiddleware, web.JSON(p.createAttachment))
	rt.GET("/attachments/edit/:id", p.Jwt.MustSignInMiddleware, p.canEditAttachment, web.JSON(p.updateAttachment))
	rt.POST("/attachments/edit/:id", p.Jwt.MustSignInMiddleware, p.canEditAttachment, web.JSON(p.updateAttachment))
	rt.DELETE("/attachments/:id", p.Jwt.MustSignInMiddleware, p.canEditAttachment, web.JSON(p.destroyAttachment))

}
