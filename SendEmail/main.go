package main

import (
	"bytes"
	"html/template"
	"gopkg.in/gomail.v2"
	"fmt"
	"github.com/mskKandula/Variables"
)

type basicDetails struct{
	Name string
	Email string
}

var user = basicDetails{
	Name:"UserName",
	Email:"UserName@gmail.com",
}


func main(){

body,err := makeTemplate(user,"./registrationMailTemplate.html")

if err!= nil{
		panic(err)
	}

m := gomail.NewMessage() 
m.SetHeader("From", "sender@gmail.com")
m.SetHeader("To", "reciever@gmail.com")
m.SetHeader("Subject", "Registration Successfull!")
m.SetBody("text/html", body)

d := gomail.NewPlainDialer("smtp.gmail.com",587,"sender@gmail.com",Variables.Password)

if err := d.DialAndSend(m); err != nil {
	panic(err)
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