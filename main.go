package main

import (
	"database/sql"
	"log"
	"net/http"

	"dadandev.com/test/internal/handler"
	"dadandev.com/test/internal/service"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, er := sql.Open("mysql", "dadan:root@/dadan")
	if er != nil {
		panic(er.Error())
	}
	defer db.Close()
	userService := service.NewUser(db)

	server := http.NewServeMux()

	handler.UserHandler(server, userService)
	er = http.ListenAndServe(":8000", server)
	if er != nil {
		log.Println(er.Error())
	}
}
