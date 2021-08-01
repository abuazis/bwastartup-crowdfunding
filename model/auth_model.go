package model

import "github.com/golang-jwt/jwt"

type AuthClaims struct {
	jwt.StandardClaims
	UserId uint32 `json:"user_id"`
}
