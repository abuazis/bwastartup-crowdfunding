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
