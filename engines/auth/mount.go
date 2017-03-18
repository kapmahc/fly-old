package auth

import (
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Mount web mount-points
func (p *Engine) Mount(rt *gin.Engine) {
	rt.GET("/users", HTML(p.indexUsers))

	ung := rt.Group("/users")
	ung.GET("/sign-in", HTML(p.formUsersSignIn))
	ung.POST("/sign-in", HTML(p.formUsersSignIn))
	ung.GET("/sign-up", HTML(p.formUsersSignUp))
	ung.POST("/sign-up", HTML(p.formUsersSignUp))
	ung.GET("/confirm/:token", HTML(p.formUsersConfirm))
	ung.GET("/confirm", HTML(p.formUsersConfirm))
	ung.POST("/confirm", HTML(p.formUsersConfirm))
	ung.GET("/unlock/:token", HTML(p.formUsersUnlock))
	ung.GET("/unlock", HTML(p.formUsersUnlock))
	ung.POST("/unlock", HTML(p.formUsersUnlock))
	ung.GET("/forgot-password", HTML(p.formUsersForgotPassword))
	ung.POST("/forgot-password", HTML(p.formUsersForgotPassword))
	ung.GET("/reset-password/:token", HTML(p.formUsersResetPassword))
	ung.POST("/reset-password/:token", HTML(p.formUsersResetPassword))

	umg := rt.Group("/users", p.Jwt.MustSignInMiddleware)
	umg.GET("/info", HTML(p.formUsersInfo))
	umg.POST("/info", HTML(p.formUsersInfo))
	umg.GET("/change-password", HTML(p.formUsersChangePassword))
	umg.POST("/change-password", HTML(p.formUsersChangePassword))
	umg.GET("/logs", HTML(p.getUsersLogs))
	umg.DELETE("/sign-out", web.JSON(p.deleteUsersSignOut))

	rt.GET("/attachments", p.Jwt.MustSignInMiddleware, HTML(p.indexAttachments))
	rt.GET("/attachments/new", p.Jwt.MustSignInMiddleware, HTML(p.newAttachment))
	// rt.POST("/attachments", p.Jwt.MustSignInMiddleware, web.JSON(p.createAttachment))
	// rt.POST("/attachments/:id", p.Jwt.MustSignInMiddleware, p.canEditAttachment, web.JSON(p.updateAttachment))
	rt.DELETE("/attachments/:id", p.Jwt.MustSignInMiddleware, p.canEditAttachment, web.JSON(p.destroyAttachment))

}
