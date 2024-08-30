package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"dadandev.com/test/domain"
)

type userHandler struct {
	UserService domain.UserService
}

func UserHandler(srv *http.ServeMux, userService domain.UserService) {
	userHandler := userHandler{
		UserService: userService,
	}
	srv.HandleFunc("/login", userHandler.Login)

}
func (uh *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = uh.UserService.Login(w, r, user)
	if err != nil {
		log.Println(err)
		http.Error(w, "ADA KESALAHAN", http.StatusBadRequest)
	}

}
