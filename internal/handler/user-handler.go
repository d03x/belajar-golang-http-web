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
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		bytedata, _ := json.Marshal(user)

		w.Write(bytedata)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}
