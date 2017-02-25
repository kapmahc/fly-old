package auth

const (
	actConfirm       = "confirm"
	actUnlock        = "unlock"
	actResetPassword = "reset-password"
)

func (p *Engine) sendEmail(lang string, user *User, act string) {

}
