package service

import (
	"bwastartup-crowdfunding/database"
	"bwastartup-crowdfunding/model"
	"bwastartup-crowdfunding/repository"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var userRepository repository.UserRepository
var userService UserService

// Main test for service package
func TestMain(m *testing.M) {
	db := database.GetConnection()
	userRepository = repository.NewUserRepositoryImpl(db)
	userService = NewUserServiceImpl(userRepository)

	authService = NewAuthServiceImpl()

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

func TestUserServiceImpl_GenerateAvatarName(t *testing.T) {
	avatar := userService.GenerateAvatarName(1, "test service image", ".jpg")
	assert.Equal(t, "1-test-service-image-avatar.jpg", avatar)
	fmt.Println(avatar)
}

func TestUserServiceImpl_SaveAvatar(t *testing.T) {
	ctx := context.Background()
	avatar := userService.GenerateAvatarName(1, "test service image", ".txt")
	_, err := userService.SaveAvatar(ctx, 1, avatar)
	if err == nil {
		t.Fail()
	}
	assert.Equal(t, "upload: invalid file extension", err.Error())
	fmt.Println(err.Error())
}
