package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"dadandev.com/test/domain"
)

type userService struct {
	db *sql.DB
}

func NewUser(db *sql.DB) domain.UserService {
	return userService{
		db: db,
	}
}

// Login implements domain.UserService.
func (u userService) Login(w http.ResponseWriter, r *http.Request, user domain.User) error {
	_, err := json.Marshal(user)
	if err != nil {
		return err
	}
	rows, _ := u.db.Query("SELECT * FROM users")
	type Data struct {
		Username  string
		FirstName string
	}
	for rows.Next() {
		var Data Data
		err := rows.Scan(&Data.Username, &Data.FirstName)
		if err != nil {
			panic("Closed")
		}
		fmt.Println(Data.FirstName)
	}
	return err
}

// Register implements domain.UserService.
func (u userService) Register() {
	panic("unimplemented")
}
