package auth

import "github.com/gorilla/mux"

// Mount web mount points
func (p *Engine) Mount(rt *mux.Router) {
	ug := rt.PathPrefix("/users").Subrouter()
	ug.HandleFunc("/sign-in", p.getSignIn).Name("users.sign-in")
}
