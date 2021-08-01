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
