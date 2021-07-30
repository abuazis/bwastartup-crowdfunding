package service

import (
	"bwastartup-crowdfunding/entity"
	"bwastartup-crowdfunding/model"
	"bwastartup-crowdfunding/repository"
	"context"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	repository repository.UserRepository
}

func NewUserServiceImpl(repository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{repository: repository}
}

func (u UserServiceImpl) Register(ctx context.Context, request model.RegisterRequest) (entity.User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	if err != nil {
		return entity.User{}, err
	}
	user := entity.User{
		Name:         request.Name,
		Occupation:   request.Occupation,
		Email:        request.Email,
		PasswordHash: string(passwordHash),
	}

	save, err := u.repository.Save(ctx, user)
	if err != nil {
		return entity.User{}, err
	}

	return save, nil
}
