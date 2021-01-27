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

func decodeJwt(tokenString string) (interface{}, error) {

	// Initialize a new instance of `Claims`
	claims := jwt.MapClaims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("jwtKey")), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", err
		}
		return "", err
	}

	if !token.Valid {
		return "", err
	}

	// can get the other values as well
	email := claims["email"]

	return email, nil
}

func main() {

	tokenStr, expirationTime, err := generateJwt()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("token :-", tokenStr)

	fmt.Println("The token will expire at : ", expirationTime)

	email, err := decodeJwt(tokenStr)

	if err != nil {
		log.Fatal(err)
	}

	//Perform Authorization or Authentication based on the email & password from decoded token

	fmt.Println("Email is :", email)
}
