package service

import (
	"bwastartup-crowdfunding/model"
	"errors"
	"github.com/golang-jwt/jwt"
)

var (
	JWT_SIGNATURE_KEY  = []byte("4rya Yun4nT4 H4Nd50M3")
	APPLICATION_NAME   = "Crowdfunding"
	JWT_SIGNING_METHOD = jwt.SigningMethodHS256
)

type AuthServiceImpl struct {
}

func NewAuthServiceImpl() *AuthServiceImpl {
	return &AuthServiceImpl{}
}

func (a *AuthServiceImpl) GenerateToken(id uint32) (string, error) {
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

func (a *AuthServiceImpl) ValidateToken(encodedToken string) (*jwt.Token, error) {
	parseToken, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("signing method invalid")
		} else if method != JWT_SIGNING_METHOD {
			return nil, errors.New("signing method invalid")
		}
		return JWT_SIGNATURE_KEY, nil
	})

	if err != nil {
		return nil, err
	}

	return parseToken, nil
}

