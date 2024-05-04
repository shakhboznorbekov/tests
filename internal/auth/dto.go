package auth

import "github.com/golang-jwt/jwt"

type SignIn struct {
	FirstName string `json:"first_name"`
	Password  string `json:"password"`
}

type Claims struct {
	FirstName string `json:"first_name"`
	jwt.StandardClaims
}

type GenerateToken struct {
	FirstName string
	Username  string
}

type TokenData struct {
	FirstName string
	UserId    string
}
type AuthResponse struct {
	Token string `json:"token"`
}
