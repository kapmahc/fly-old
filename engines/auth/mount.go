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
	// ---------------

	rt.GET("/attachments", p.Jwt.MustSignInMiddleware, web.JSON(p.indexAttachments))
	rt.POST("/attachments", p.Jwt.MustSignInMiddleware, web.JSON(p.createAttachment))
	rt.GET("/attachments/:id", web.JSON(p.showAttachment))
	rt.POST("/attachments/:id", p.Jwt.MustSignInMiddleware, p.canEditAttachment, web.JSON(p.updateAttachment))
	rt.DELETE("/attachments/:id", p.Jwt.MustSignInMiddleware, p.canEditAttachment, web.JSON(p.destroyAttachment))

	umg := rt.Group("/users", p.Jwt.MustSignInMiddleware)
	umg.POST("/info", web.JSON(p.postUsersInfo))
	umg.GET("/info", web.JSON(p.getUsersInfo))
	umg.POST("/change-password", web.JSON(p.postUsersChangePassword))
	umg.GET("/logs", web.JSON(p.getUsersLogs))
	umg.DELETE("/sign-out", web.JSON(p.deleteUsersSignOut))

}
