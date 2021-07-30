package repository

import (
	"bwastartup-crowdfunding/entity"
	"context"
)

type UserRepository interface {
	Save(ctx context.Context, user entity.User) (entity.User, error)
}
