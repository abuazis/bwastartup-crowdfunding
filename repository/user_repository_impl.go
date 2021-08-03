package repository

import (
	"bwastartup-crowdfunding/entity"
	"context"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{Db: db}
}

func (u *UserRepositoryImpl) Save(ctx context.Context, user entity.User) (entity.User, error) {
	err := u.Db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (u *UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (entity.User, error) {
	var user entity.User
	err := u.Db.WithContext(ctx).Where("email=?", email).First(&user).Error
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (u *UserRepositoryImpl) FindById(ctx context.Context, id uint32) (entity.User, error) {
	var user entity.User
	err := u.Db.WithContext(ctx).Where("id=?", id).First(&user).Error
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (u *UserRepositoryImpl) UpdateAvatar(ctx context.Context, id uint32, avatarFileName string) (bool, error) {
	err := u.Db.WithContext(ctx).Model(&entity.User{}).Where("id=?", id).Update("avatar_file_name", avatarFileName).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
