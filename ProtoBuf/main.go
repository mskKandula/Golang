package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/mskKandula/protobuf/personDetails"
	"google.golang.org/protobuf/proto"
)

func main() {
	person := getDetails()

	writeToFile(person, "person.bin")
	readFromFile("person.bin")

}

func readFromFile(fname string) error {
	out, err := ioutil.ReadFile(fname)

	if err != nil {
		log.Fatalln("Unable to read Person file: ", err)
		return err
	}

	person := &personDetails.Person{}

	if err = proto.Unmarshal(out, person); err != nil {
		log.Fatalln("Unable to read Person from a file : ", err)
		return err
	}
	fmt.Println("Successfully read a person file: ", person)
	return nil
}

func writeToFile(person *personDetails.Person, fname string) error {
	out, err := proto.Marshal(person)

	if err != nil {
		log.Fatalln("Unable to Marshal Person: ", err)
		return err
	}

	if err = ioutil.WriteFile(fname, out, 0666); err != nil {
		log.Fatalln("Unable to write Person to a file: ", err)
		return err
	}
	fmt.Println("Successfully Written")
	return nil
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
