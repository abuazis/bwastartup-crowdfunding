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
