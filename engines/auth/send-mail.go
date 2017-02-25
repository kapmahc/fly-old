package auth

const (
	actConfirm       = "confirm"
	actUnlock        = "unlock"
	actResetPassword = "reset-password"
)

func (p *Engine) sendEmail(lang string, user *User, act string) {

}

func (p *Engine) parseToken(lang, token, act string) (*User, error) {

}
