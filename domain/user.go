package domain

type UserService interface {
	Login()
	Register()
}
type User struct {
	Email    string
	Password string
}
