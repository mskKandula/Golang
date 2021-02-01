package controller

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mskKandula/model"
)

var (
	Db  *sql.DB
	err error
)

func Signup(w http.ResponseWriter, r *http.Request) {

	var user model.User

	json.NewDecoder(r.Body).Decode(&user)

	query, err := Db.Prepare("INSERT INTO USERS(name, age, email, phoneNo, password) VALUES(?,?,?,?,?)")
	if err != nil {
		log.Panic(err)
	}
	query.Exec(user.Name, user.Age, user.Email, user.PhoneNo, user.Password)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)

}
func Login(w http.ResponseWriter, r *http.Request) {

}
func Logout(w http.ResponseWriter, r *http.Request) {

}
