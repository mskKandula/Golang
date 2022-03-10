package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"

	"github.com/mskKandula/Variables"
	"gopkg.in/gomail.v2"
)

type basicDetails struct {
	Name  string
	Email string
}

var user = basicDetails{
	Name:  "RecieverUserName",
	Email: "RecieverEmailId",
}

func main() {

	body, err := makeTemplate(user, "./registrationMailTemplate.html")

	if err != nil {
		panic(err)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", Variables.SenderEmail)
	m.SetHeader("To", user.Email)
	m.SetHeader("Subject", "Registration Successfull!")
	m.SetBody("text/html", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, Variables.SenderEmail, Variables.SenderEmailPassword)

	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
		return
	}

	fmt.Println("Email Sent")

}

func makeTemplate(profileObj basicDetails, templatePath string) (string, error) {

	parsedTemplate, err := template.ParseFiles(templatePath)

	if err != nil {
		return "", err
	}

	var buff bytes.Buffer

	parseErr := parsedTemplate.Execute(&buff, profileObj)

	if parseErr != nil {
		return "", parseErr
	}

	body := buff.String()

	return body, nil
}
