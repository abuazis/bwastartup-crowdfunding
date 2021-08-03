package repository

import (
	"bwastartup-crowdfunding/entity"
	"context"
)

type UserRepository interface {
	Save(ctx context.Context, user entity.User) (entity.User, error)
	FindByEmail(ctx context.Context, email string) (entity.User, error)
	FindById(ctx context.Context, id uint32) (entity.User, error)
	UpdateAvatar(ctx context.Context, id uint32, avatarFileName string) (bool, error)
}
