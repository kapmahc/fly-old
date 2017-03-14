package auth

import (
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Mount web mount-points
func (p *Engine) Mount(rt *gin.Engine) {
	rt.GET("/users", p.indexUsers)

	rt.GET("/attachment/*name", p.getAttachment)
	rt.GET("/attachments", p.Jwt.MustSignInMiddleware, p.indexAttachments)
	rt.POST("/attachments", p.Jwt.MustSignInMiddleware, p.createAttachment)
	rt.GET("/attachments/:id", p.showAttachment)
	rt.POST("/attachments/:id", p.Jwt.MustSignInMiddleware, p.canEditAttachment, p.updateAttachment)
	rt.DELETE("/attachments/:id", p.Jwt.MustSignInMiddleware, p.canEditAttachment, p.destroyAttachment)

	ung := rt.Group("/users")
	ung.POST("/sign-in", p.postUsersSignIn)
	ung.POST("/sign-up", p.postUsersSignUp)
	ung.GET("/confirm/:token", web.Redirect(p.getUsersConfirm))
	ung.POST("/confirm", p.postUsersConfirm)
	ung.GET("/unlock/:token", web.Redirect(p.getUsersUnlock))
	ung.POST("/unlock", p.postUsersUnlock)
	ung.POST("/forgot-password", p.postUsersForgotPassword)
	ung.POST("/reset-password/:token", p.postUsersResetPassword)

	umg := rt.Group("/users", p.Jwt.MustSignInMiddleware)
	umg.POST("/info", p.postUsersInfo)
	umg.GET("/info", p.getUsersInfo)
	umg.POST("/change-password", p.postUsersChangePassword)
	umg.GET("/logs", p.getUsersLogs)
	umg.DELETE("/sign-out", p.deleteUsersSignOut)

}
