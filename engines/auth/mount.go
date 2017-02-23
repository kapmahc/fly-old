package auth

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Mount web mount points
func (p *Engine) Mount(rt *mux.Router) {
	ug := rt.PathPrefix("/users").Subrouter()
	ug.HandleFunc("/sign-in", p.signIn).Methods(http.MethodGet, http.MethodPost).Name("users.sign-in")
	ug.HandleFunc("/sign-up", p.signUp).Methods(http.MethodGet, http.MethodPost).Name("users.sign-up")
	ug.HandleFunc("/confirm", p.confirm).Methods(http.MethodGet, http.MethodPost).Name("users.confirm")
	ug.HandleFunc("/unlock", p.unlock).Methods(http.MethodGet, http.MethodPost).Name("users.unlock")
	ug.HandleFunc("/reset-password", p.resetPassword).Methods(http.MethodGet, http.MethodPost).Name("users.reset-password")
	ug.HandleFunc("/forgot-password", p.forgotPassword).Methods(http.MethodGet, http.MethodPost).Name("users.forgot-password")
}