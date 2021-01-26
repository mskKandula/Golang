package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func generateJwt() (string, time.Time, error) {

	//Creating Access Token
	os.Setenv("jwtKey", "7yt65U745TR57lo9h%$fre#$TR43EW") //this should be in an env file

	atClaims := jwt.MapClaims{}

	// Mapping Claims
	atClaims["authorized"] = true

	atClaims["email"] = "user@example.org"

	atClaims["password"] = "123456789"

	expirationTime := time.Now().Add(time.Minute * 5)

	atClaims["expireAt"] = expirationTime

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	tokenString, err := token.SignedString([]byte(os.Getenv("jwtKey")))

	if err != nil {
		return "", expirationTime, err
	}

	return tokenString, expirationTime, nil
}

func main() {

	tokenStr, expirationTime, err := generateJwt()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("token :-", tokenStr)

	fmt.Println("The token will expire at : ", expirationTime)

}
