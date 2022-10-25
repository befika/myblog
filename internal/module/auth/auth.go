package auth

import (
	"blog/internal/adapter/storage/user"
	"context"
	"errors"
	"fmt"
	"time"

	model "blog/internal/constant/model/db"
	utils "blog/internal/constant/model/init"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type LoginUseCase interface {
	GetClaims(token string) (*model.UserClaims, error)
	Login(c context.Context, Email, password string) (*model.LoginResponse, error)
}
type authservice struct {
	userPersistence user.User
	jwtManager      JWTManager
	contextTimeout  time.Duration
}

func Initialize(userPersistence user.User, jwtManager JWTManager, utils utils.Utils) LoginUseCase {
	return &authservice{
		userPersistence: userPersistence,
		jwtManager:      jwtManager,
		contextTimeout:  utils.Timeout,
	}
}
func (s authservice) GetClaims(token string) (*model.UserClaims, error) {
	claims, err := s.jwtManager.Verify(token)
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func (s authservice) Login(c context.Context, Email, password string) (*model.LoginResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	if Email == "" || password == "" {
		return nil, errors.New("empty login data")
	}
	u := model.User{Email: Email}
	usr, err := s.userPersistence.User(ctx, &u)
	if err != nil {
		return nil, err
	}
	er := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(password))
	if er != nil {
		return nil, er

	}
	for _, role := range *usr.Roles {
		if len(*usr.Roles) > 1 {
			break
		} else if role.Name == "USER" {
			if usr.Status == "Pending" || usr.Status == "" {
				return nil, errors.New("please pay subscription price")

			} else if usr.Status == "Blocked" {
				return nil, errors.New("your account is deactivated please contact the system admin")

			}

		}
	}
	claims := &model.UserClaims{
		StandardClaims: jwt.StandardClaims{
			Subject: usr.ID.String(),
		},
		Email:  Email,
		Roles:  *usr.Roles,
		UserId: usr.ID,
	}
	token, err := s.jwtManager.Generate(claims)
	if err != nil {
		return nil, err
	}
	return &model.LoginResponse{
		Token:  token,
		Name:   fmt.Sprintf("%v %v %v", usr.FirstName, usr.MiddleName, usr.LastName),
		Email:  usr.Email,
		UserId: usr.ID,
	}, nil
}
