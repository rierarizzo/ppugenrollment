package security

import "github.com/golang-jwt/jwt/v5"

type UserClaims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}
