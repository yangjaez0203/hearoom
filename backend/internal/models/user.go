package models

import "github.com/golang-jwt/jwt/v5"

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Anonymous bool   `json:"anonymous"`
}

type UserClaims struct {
	Username  string `json:"username"`
	Anonymous bool   `json:"anonymous"`
	jwt.RegisteredClaims
}
