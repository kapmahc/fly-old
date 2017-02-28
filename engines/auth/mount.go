package auth

// Mount web mount points
func (p *Engine) Mount() {
	ug := p.Mux.Group("/users")

	ug.Form("auth.users.sign-in", "/sign-in", p.signIn)
	ug.Form("auth.users.sign-up", "/sign-up", p.signUp)
	ug.Form("auth.users.confirm", "/confirm", p.confirm)
	ug.Form("auth.users.unlock", "/unlock", p.unlock)
	ug.Form("auth.users.reset-password", "/reset-password", p.resetPassword)
	ug.Form("auth.users.forgot-password", "/forgot-password", p.forgotPassword)

	ug.Form("auth.users.info", "/info", p.info)
	ug.Form("auth.users.change-password", "/change-password", p.changePassword)
	ug.Get("auth.users.logs", "/logs", p.logs)
	ug.Delete("auth.users.sign-out", "/sign-out", p.signOut)
}
