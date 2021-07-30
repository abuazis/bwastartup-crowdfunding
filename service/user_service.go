package service

import (
	"bwastartup-crowdfunding/entity"
	"bwastartup-crowdfunding/model"
	"context"
)

type UserService interface {
	Register(ctx context.Context, request model.RegisterRequest) (entity.User, error)
}
