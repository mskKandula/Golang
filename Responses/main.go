package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
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
	http.HandleFunc("/jsonVariant", jsonVariantHandler)
	http.HandleFunc("/template", templateHandler)
	http.HandleFunc("/file", fileHandler)
	fmt.Println("listening on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))

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
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// do some manipulations on User struct

	// sending the final User struct as a response
	json.NewEncoder(w).Encode(u)

}

func jsonVariantHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var name, last string

	response, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonList := gjson.ParseBytes(response)

	var dataToSend []gjson.Result

	for _, val := range jsonList.Array() {
		name = val.Get("name").String()
		last = val.Get("last").String()
		data := prepareResult([]string{"name", "last"}, []interface{}{name, last})
		dataToSend = append(dataToSend, data)
	}
	json.NewEncoder(w).Encode(dataToSend)
}

func prepareResult(keys []string, vals []interface{}) gjson.Result {
	var data string
	for i, k := range keys {
		data, _ = sjson.Set(data, k, vals[i])
	}

	return gjson.Parse(data)
}

// fileHandler downloads a file on the client machine
func fileHandler(w http.ResponseWriter, r *http.Request) {

	imageFile, er := os.Open("./Sample.png")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := User{Id: 1,
		Name:  "YourName",
		Email: "YourEmail",
		Phone: "YourNumber"}
	t.Execute(w, user)
}
