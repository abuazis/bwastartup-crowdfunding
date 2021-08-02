package service

import (
	"fmt"
	"testing"
)

var authService AuthService

func TestAuthServiceImpl_GenerateToken(t *testing.T) {
	generateToken, err := authService.GenerateToken(1)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(generateToken)
}

func TestAuthServiceImpl_ValidateToken(t *testing.T) {
	encodedToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJDcm93ZGZ1bmRpbmciLCJ1c2VyX2lkIjoxfQ.V_Ysz6FVjT8mbwsD8ZNY4cx2rThTxLc62EZUGFG3tNQ"
	token, err := authService.ValidateToken(encodedToken)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(token)
}
func TestAuthServiceImpl_ValidateTokenFail(t *testing.T) {
	encodedToken := "WRONG TOKEN"
	_, err := authService.ValidateToken(encodedToken)
	if err == nil {
		t.Error()
	}
	fmt.Println(err.Error())
}

func TestAuthServiceImpl_ValidateTokenFail2(t *testing.T) {
	encodedToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVC9.eyJpc3MiOiJDcm93ZGZ1bmRpbmciLCJ1c2VyX2lkIjoxfQ.V_Ysz6FVjT8mbwsD8ZNY4cx2rThTxLc62EZUGFG3tNQ"
	_, err := authService.ValidateToken(encodedToken)
	if err == nil {
		t.Fatal()
	}
	fmt.Println(err.Error())
}
