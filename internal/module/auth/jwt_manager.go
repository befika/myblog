package auth

import (
	model "blog/internal/constant/model/db"
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type JWTManager struct {
	secretKey string
}

func NewJWTManager(secretKey string) *JWTManager {
	return &JWTManager{secretKey}
}
func (manager *JWTManager) Generate(userClaims *model.UserClaims) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	tk, err := token.SignedString([]byte(manager.secretKey))
	if err != nil {
		return "", errors.New("tocken generate error")
	}
	return tk, nil
}

func (manager *JWTManager) Verify(accessToken string) (*model.UserClaims, error) {
	if accessToken == "" {
		return nil, errors.New("tocken verify error")
	}
	token, err := jwt.ParseWithClaims(
		accessToken,
		&model.UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("tocken verify error")
			}

			return []byte(manager.secretKey), nil
		},
	)
	if err != nil {
		return nil, errors.New("tocken verify error")
	}
	claims, ok := token.Claims.(*model.UserClaims)
	if !ok {
		return nil, errors.New("tocken verify error")
	}
	return claims, nil
}
