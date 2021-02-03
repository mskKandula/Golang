package middleware

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mskKandula/model"
)

func Auth(creds model.Validation) (string, time.Time, error) {

	var err error

	//Creating Access Token
	os.Setenv("jwtKey", "7yt65U745TR57lo9h%$fre#$TR43EW") //this should be in an env file

	atClaims := jwt.MapClaims{}

	atClaims["authorized"] = true

	atClaims["email"] = creds.Email

	atClaims["password"] = creds.Password

	expirationTime := time.Now().Add(time.Minute * 5)

	atClaims["expireAt"] = expirationTime

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte(os.Getenv("jwtKey")))

	if err != nil {
		return "", expirationTime, err
	}

	return token, expirationTime, nil

}
