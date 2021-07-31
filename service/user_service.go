package service

import (
	"bwastartup-crowdfunding/model"
	"context"
)

type UserService interface {
	Register(ctx context.Context, request model.RegisterRequest) (model.RegisterResponse, error)
	Login(ctx context.Context,request model.LoginRequest) (model.LoginResponse, error)
}
