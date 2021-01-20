package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
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
	http.HandleFunc("/file", fileHandler)
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

// fileHandler downloads a file on the client machine
func fileHandler(w http.ResponseWriter, r *http.Request) {

	imageFile, _ := os.Open("./Sample.png")

	// its for the file to download automatically as "Sample.png"
	w.Header().Set("Content-Disposition", "attachment; filename=Sample.png")

	w.Header().Set("Content-Type", "image/png")

	defer imageFile.Close()

	// copying the imageFile to response writer
	io.Copy(w, imageFile)
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
