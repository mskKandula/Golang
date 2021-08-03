package main

import (
	"fmt"

	"github.com/mskKandula/protobuf/personDetails"
)

func main() {
	person := getDetails()

	fmt.Println("The Details of a person are : ", person)

}

func getDetails() *personDetails.Person {
	person := &personDetails.Person{
		Age:          28,
		Name:         "ABC",
		MobileNumber: []string{"1234567890", "9876543210"},
		Email:        "example@gmail.com",
	}

	fmt.Println("The mobile numbers of a person are: ", person.GetMobileNumber())

	fmt.Println("The age of a person is: ", person.GetAge())

	return person
}
