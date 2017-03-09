package auth

import gin "gopkg.in/gin-gonic/gin.v1"

// Mount web mount-points
func (p *Engine) Mount(rt *gin.Engine) {
	ung := rt.Group("/users")
	ung.POST("/sign-in", p.postUsersSignIn)
	ung.POST("/sign-up", p.postUsersSignUp)
	ung.GET("/confirm/:token", p.getUsersConfirm)
	ung.POST("/confirm", p.postUsersConfirm)
	ung.GET("/unlock/:token", p.getUsersUnlock)
	ung.POST("/unlock", p.postUsersUnlock)
	ung.POST("/forgot-password", p.postUsersForgotPassword)
	ung.POST("/reset-password/:token", p.postUsersResetPassword)

	ung.GET("/", p.indexUsers)

	umg := rt.Group("/users", p.Jwt.MustSignInMiddleware)
	umg.POST("/info", p.postUsersInfo)
	umg.POST("/change-password", p.postUsersChangePassword)
	umg.GET("/logs", p.getUsersLogs)
	umg.DELETE("/sign-out", p.deleteUsersSignOut)

}
