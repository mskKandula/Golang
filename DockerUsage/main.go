package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Person struct {
	Name string `json:"name"`
}

var (
	db  *sql.DB
	err error
)

func main() {
	db, err = sql.Open("mysql", "root:root@tcp(db:3306)/Users")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()
	http.HandleFunc("/", getPosts)
	fmt.Println("listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

//getPosts returns http respone
func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var persons []Person
	result, err := db.Query("SELECT name from Users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer result.Close()
	for result.Next() {
		var person Person
		err := result.Scan(&person.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		persons = append(persons, person)
	}
	json.NewEncoder(w).Encode(persons)

}
