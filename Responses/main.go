package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func main() {
	http.HandleFunc("/", stringHandler)
	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/template", templateHandler)
	fmt.Println("listening on port 8081")
	http.ListenAndServe(":8081", nil)

}

//stringHandler returns http respone in string format.
func stringHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Buddy!")

}

//jsonHandler returns http respone in JSON format.
func jsonHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var u User

	// Unmarshalling the data into User struct
	json.NewDecoder(r.Body).Decode(&u)

	// do some manipulations on User struct

	// sending the final User struct as a response
	json.NewEncoder(w).Encode(u)

}

//templateHandler renders a template and returns as http response.
func templateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Fprintf(w, "Unable to load template")
	}

	user := User{Id: 1,
		Name:  "YourName",
		Email: "YourEmail",
		Phone: "YourNumber"}
	t.Execute(w, user)
}
