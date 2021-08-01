package service

import (
	"bwastartup-crowdfunding/model"
	"github.com/golang-jwt/jwt"
)

var (
	JWT_SIGNATURE_KEY         = []byte("4rya Yun4nT4 H4Nd50M3")
	APPLICATION_NAME          = "Crowdfunding"
	JWT_SIGNING_METHOD        = jwt.SigningMethodHS256
)

type AuthServiceImpl struct {
}

func NewAuthServiceImpl() *AuthServiceImpl {
	return &AuthServiceImpl{}
}

func (J *AuthServiceImpl) GenerateToken(id uint32) (string, error) {
	claims := model.AuthClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer: APPLICATION_NAME,
		},
		UserId: id,
	}

	token := jwt.NewWithClaims(JWT_SIGNING_METHOD, claims)
	signedString, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		return "", err
	}
	return signedString, nil
}
