package auth

import "github.com/gorilla/sessions"

type Session struct {
	Store *sessions.CookieStore `inject:""`
}
