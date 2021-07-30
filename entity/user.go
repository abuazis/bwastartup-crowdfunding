package entity

import "time"

type User struct {
	Id                                                          uint32
	Name, Occupation, Email, PasswordHash, AvatarFileName, Role string
	CreatedAt                                                   time.Time
	UpdatedAt                                                   *time.Time
}
