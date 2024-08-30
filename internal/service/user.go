package service

import "dadandev.com/test/domain"

type userService struct {
}

func NewUser() domain.UserService {
	return userService{}
}

// Login implements domain.UserService.
func (u userService) Login() {
	panic("unimplemented")
}

// Register implements domain.UserService.
func (u userService) Register() {
	panic("unimplemented")
}
