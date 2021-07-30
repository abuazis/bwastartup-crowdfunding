package repository

import (
	"bwastartup-crowdfunding/database"
	"bwastartup-crowdfunding/entity"
	"context"
	"fmt"
	"testing"
)

var repository UserRepository

func TestMain(m *testing.M) {
	db := database.GetConnection()
	repository = NewUserRepositoryImpl(db)

	m.Run()
}

func TestUserRepositoryImpl_Save(t *testing.T) {
	user := entity.User{
		Name:           "Test Repository",
		Occupation:     "Student",
		Email:          "test_repository@test.com",
		PasswordHash:   "password",
		AvatarFileName: "myfile.jpg",
		Role:           "user",
	}
	ctx := context.Background()
	saveUser, err := repository.Save(ctx, user)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(saveUser)
}
