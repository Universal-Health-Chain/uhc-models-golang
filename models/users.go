package models

import (
	"github.com/dgrijalva/jwt-go"
)

/*
JWT claims struct
*/
type Token struct {
	UserId   string
	Username string
	jwt.StandardClaims
}

//a struct to rep user user
type User struct {
	Username string   `bson:"username" json:"username,omitempty"`
	Password string   `bson:"password" json:"password,omitempty"`
	Token    string   `bson:"token" json:"token,omitempty"`
	ID       string   `json:"id,omitempty"`
	Email    string   `json:"email"`
	Role     []string `json:"role"`
}
