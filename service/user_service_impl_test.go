package service

import (
	"bwastartup-crowdfunding/database"
	"bwastartup-crowdfunding/model"
	"bwastartup-crowdfunding/repository"
	"context"
	"fmt"
	"testing"
)

var userRepository repository.UserRepository
var userService UserService

func TestMain(m *testing.M) {
	db := database.GetConnection()
	userRepository = repository.NewUserRepositoryImpl(db)
	userService = NewUserServiceImpl(userRepository)

	m.Run()
}

func TestUserServiceImpl_Register(t *testing.T) {
	ctx := context.Background()
	request := model.RegisterRequest{
		Name:       "Test From User Service",
		Occupation: "Student",
		Email:      "service@test.com",
		Password:   "itsSecureTrustMe12",
	}

	register, err := userService.Register(ctx, request)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(register)
}

func TestUserServiceImpl_Login(t *testing.T) {
	ctx := context.Background()
	request := model.LoginRequest{
		Email:    "service@test.com",
		Password: "itsSecureTrustMe12",
	}
	loginResponse, err := userService.Login(ctx, request)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(loginResponse)
}

func TestUserServiceImpl_LoginEmailFail(t *testing.T) {
	ctx := context.Background()
	request := model.LoginRequest{
		Email:    "notfound@test.com",
		Password: "itsSecureTrustMe12",
	}
	loginResponse, err := userService.Login(ctx, request)
	if err == nil {
		t.Fail()
	}
	fmt.Println(loginResponse)
}

func TestUserServiceImpl_LoginPasswordFailAndSQLInjection(t *testing.T) {
	ctx := context.Background()
	request := model.LoginRequest{
		Email:    "service@test.com",
		Password: "wrongPassword#;",
	}
	loginResponse, err := userService.Login(ctx, request)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(loginResponse)
}

func TestUserServiceImpl_CheckEmail(t *testing.T) {
	ctx := context.Background()
	email := "service@test.com"
	checkEmail, err := userService.CheckEmail(ctx, email)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(checkEmail)
}
