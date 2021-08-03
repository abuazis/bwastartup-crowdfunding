package service

import (
	"bwastartup-crowdfunding/entity"
	"bwastartup-crowdfunding/model"
	"context"
)

type UserService interface {
	Register(ctx context.Context, request model.RegisterRequest) (model.RegisterResponse, error)
	Login(ctx context.Context, request model.LoginRequest) (model.LoginResponse, error)
	CheckEmail(ctx context.Context, email string) (bool, error)
	FindById(ctx context.Context, id uint32) (entity.User, error)
	SaveAvatar(ctx context.Context, id uint32, fileName string) (string, error)
	GenerateAvatarName(id uint32, name string, extension string) string
}
