package repository

import (
	"bwastartup-crowdfunding/database"
	"bwastartup-crowdfunding/entity"
	"context"
	"fmt"
	"testing"
)

var userRepository UserRepository

func TestMain(m *testing.M) {
	db := database.GetConnection()
	userRepository = NewUserRepositoryImpl(db)
	campaignRepository = NewCampaignRepositoryImpl(db)

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
	saveUser, err := userRepository.Save(ctx, user)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(saveUser)
}

func TestUserRepositoryImpl_FindByEmail(t *testing.T) {
	ctx := context.Background()
	user, err := userRepository.FindByEmail(ctx, "service@test.com")
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(user)
}

func TestUserRepositoryImpl_FindById(t *testing.T) {
	ctx := context.Background()
	user, err := userRepository.FindById(ctx, 1)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(user)
}

func TestUserRepositoryImpl_FindByIdFail(t *testing.T) {
	ctx := context.Background()
	_, err := userRepository.FindById(ctx, 999999)
	if err == nil {
		t.Fatal()
	}
	fmt.Println(err.Error())
}

func TestUserRepositoryImpl_UpdateAvatar(t *testing.T) {
	ctx := context.Background()
	isUpdate, err := userRepository.UpdateAvatar(ctx, 1, "test-update.jpg")
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(isUpdate)
}
