package main

import (
	"fmt"
	"net/http"

	"github.com/mskKandula/controller"
)

var (
	err error
)

func main() {
	http.HandleFunc("/signup", controller.Signup)
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/logout", controller.Logout)
	fmt.Println("server is running on 8080")
	http.ListenAndServe(":8080", nil)
}
