package service

import (
	"bwastartup-crowdfunding/entity"
	"bwastartup-crowdfunding/model"
	"bwastartup-crowdfunding/repository"
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"path/filepath"
	"strings"
)

type UserServiceImpl struct {
	repository repository.UserRepository
}

func NewUserServiceImpl(repository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{repository: repository}
}

func (u *UserServiceImpl) Register(ctx context.Context, request model.RegisterRequest) (model.RegisterResponse, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	if err != nil {
		return model.RegisterResponse{}, err
	}
	user := entity.User{
		Name:         request.Name,
		Occupation:   request.Occupation,
		Email:        request.Email,
		PasswordHash: string(passwordHash),
		Role:         "user",
	}

	save, err := u.repository.Save(ctx, user)
	if err != nil {
		return model.RegisterResponse{}, err
	}

	return model.RegisterResponse{
		Id:         save.Id,
		Name:       save.Name,
		Occupation: save.Occupation,
		Email:      save.Email,
	}, nil
}

func (u *UserServiceImpl) Login(ctx context.Context, request model.LoginRequest) (model.LoginResponse, error) {
	user, err := u.repository.FindByEmail(ctx, request.Email)
	if err != nil {
		return model.LoginResponse{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password))
	if err != nil {
		return model.LoginResponse{}, err
	}
	return model.LoginResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Token: "Not implemented",
	}, nil
}

func (u *UserServiceImpl) CheckEmail(ctx context.Context, email string) (bool, error) {
	_, err := u.repository.FindByEmail(ctx, email)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u *UserServiceImpl) FindById(ctx context.Context, id uint32) (entity.User, error) {
	user, err := u.repository.FindById(ctx, id)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

// SaveAvatar return generateFileName
func (u *UserServiceImpl) SaveAvatar(ctx context.Context, id uint32, fileName string) (string, error) {
	user, err := u.FindById(ctx, id)
	if err != nil {
		return "", err
	}

	ext := filepath.Ext(fileName)
	// Validate image extension
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		return "", errors.New("upload: invalid file extension")
	}
	generateName := u.GenerateAvatarName(id, user.Name, ext)

	_, err = u.repository.UpdateAvatar(ctx, id, generateName)
	if err != nil {
		return "", err
	}
	return generateName, nil
}

func (u *UserServiceImpl) GenerateAvatarName(id uint32, userName string, extension string) string {
	name := strings.Join(strings.Split(userName, " "), "-")
	return fmt.Sprintf("%d-%s-avatar%s", id, name, extension)
}
