package repository

import (
	"bwastartup-crowdfunding/entity"
	"context"
)

type UserRepository interface {
	Save(ctx context.Context, user entity.User) (entity.User, error)
	FindByEmail(ctx context.Context, email string) (entity.User,error)
}
