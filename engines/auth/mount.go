package auth

import "github.com/kapmahc/sky"

// Mount web mount points
func (p *Engine) Mount(rt *sky.Router) {
	rt.Group("/users", func(r *sky.Router) {
		r.Get("auth.users.sign-in", "/sign-in", p.Layout.Application, p.getUsersSignIn)
		r.Post("auth.users.sign-in", "/sign-in", p.postUsersSignIn)
		r.Get("auth.users.sign-up", "/sign-up", p.Layout.Application, p.getUsersSignUp)
		r.Post("auth.users.sign-up", "/sign-up", p.postUsersSignUp)
		r.Get("auth.users.confirm-token", "/confirm/{token}", p.getUsersConfirm)
		r.Get("auth.users.confirm", "/confirm", p.Layout.Application, p.getUsersEmailForm("confirm"))
		r.Post("auth.users.confirm", "/confirm", p.postUsersConfirm)
		r.Get("auth.users.unlock-token", "/unlock/{token}", p.getUsersUnlock)
		r.Get("auth.users.unlock", "/unlock", p.Layout.Application, p.getUsersEmailForm("unlock"))
		r.Post("auth.users.unlock", "/unlock", p.postUsersUnlock)
		r.Get("auth.users.forgot-password", "/forgot-password", p.Layout.Application, p.getUsersEmailForm("forgot-password"))
		r.Post("auth.users.forgot-password", "/forgot-password", p.postUsersForgotPassword)
		r.Get("auth.users.reset-password", "/reset-password/{token}", p.Layout.Application, p.getUsersResetPassword)
		r.Post("auth.users.reset-password", "/reset-password", p.postUsersResetPassword)
	})

}
