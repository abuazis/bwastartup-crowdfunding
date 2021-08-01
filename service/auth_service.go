package service

type AuthService interface {
	GenerateToken(id uint32) (string, error)
}
