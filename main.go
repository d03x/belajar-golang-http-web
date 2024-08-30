package main

import (
	"log"
	"net/http"

	"dadandev.com/test/internal/handler"
	"dadandev.com/test/internal/service"
)

func main() {
	userService := service.NewUser()

	server := http.NewServeMux()

	handler.UserHandler(server, userService)

	er := http.ListenAndServe(":8000", server)
	if er != nil {
		log.Println(er.Error())
	}
}
