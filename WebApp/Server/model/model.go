package model

//user registration

type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age,omitempty"`
	Email    string `json:"email"`
	PhoneNo  string `json:"phoneNo,omitempty"`
	Password string `json:"password"`
}
