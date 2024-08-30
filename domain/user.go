package domain

import (
	"net/http"
)

type UserService interface {
	Login(w http.ResponseWriter, r *http.Request, user User) error
	Register()
}
type User struct {
	Email    string
	Password string
}
