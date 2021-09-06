package main

import (
	"html/template"
	"log"
	"net/http"
)

type Institute struct {
	InstituteName      string    `json:"instituteName"`
	RepresentativeName string    `json:"representativeName"`
	LearnerList        []Learner `json:"learnerList"`
	InstituteUrl       string    `json:"instituteUrl"`
}

type Learner struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Standard  string `json:"standard"`
	Division  string `json:"division"`
	LoginName string `json:"loginName"`
}

func templateHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./SendEmailToInstitute.html")

	data := Institute{
		InstituteName:      "ABC",
		RepresentativeName: "XYZ",
		InstituteUrl:       "https://example.com",
		LearnerList: []Learner{
			{FirstName: "Learner", LastName: "One", Standard: "10", Division: "A", LoginName: "learnero"},
			{FirstName: "Learner", LastName: "Two", Standard: "9", Division: "A", LoginName: "learnert"},
			{FirstName: "Learner", LastName: "Three", Standard: "8", Division: "A", LoginName: "learnert1"},
		},
	}

	tmpl.Execute(w, data)

}
func main() {
	fs := http.FileServer(http.Dir("images"))

	http.Handle("/images/", http.StripPrefix("/images/", fs))

	http.HandleFunc("/", templateHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
