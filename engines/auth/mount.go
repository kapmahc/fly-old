package auth

import (
	"net/http"
)

// Mount web mount points
func (p *Engine) Mount() {
	ug := p.Router.PathPrefix("/users").Subrouter()

	ug.HandleFunc("/sign-in", p.signIn).Methods(http.MethodGet, http.MethodPost).Name("auth.users.sign-in")
	ug.HandleFunc("/sign-up", p.signUp).Methods(http.MethodGet, http.MethodPost).Name("auth.users.sign-up")
	ug.HandleFunc("/confirm", p.confirm).Methods(http.MethodGet, http.MethodPost).Name("auth.users.confirm")
	ug.HandleFunc("/unlock", p.unlock).Methods(http.MethodGet, http.MethodPost).Name("auth.users.unlock")
	ug.HandleFunc("/reset-password", p.resetPassword).Methods(http.MethodGet, http.MethodPost).Name("auth.users.reset-password")
	ug.HandleFunc("/forgot-password", p.forgotPassword).Methods(http.MethodGet, http.MethodPost).Name("auth.users.forgot-password")
}
