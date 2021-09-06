package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/mskKandula/controller"
)

var (
	err error
)

func init() {
	controller.Db, err = sql.Open("mysql", "userName:password@tcp(address:port)/WebApp")

	if err != nil {
		log.Fatal("Connection Failed to Open")
	}
}

func main() {
	http.HandleFunc("/signup", controller.Signup)
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/logout", controller.Logout)
	fmt.Println("server is running on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
