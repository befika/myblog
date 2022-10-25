package user

import (
	Percistance "blog/internal/adapter/storage/user"
	model "blog/internal/constant/model/db"
	utils "blog/internal/constant/model/init"
	"blog/internal/constant/model/rest"
	"context"
	"errors"
	"time"

	ut "github.com/go-playground/universal-translator"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/casbin/casbin/v2"
	"github.com/go-playground/validator/v10"
)

type UserUsecase interface {
	AddUser(c context.Context, r model.User) (*model.User, error)
	Users(c context.Context, pgnFlt *rest.FilterParams) ([]model.User, error)
	User(c context.Context, r model.User) (*model.User, error)
	UpdateUser(c context.Context, r model.User) (*model.User, error)
	DeleteUser(c context.Context, u model.User) error
}

type service struct {
	contextTimeout  time.Duration
	enforcer        *casbin.Enforcer
	validator       *validator.Validate
	trans           ut.Translator
	userPercistance Percistance.User
}

func InitializeUser(userPercistance Percistance.User, utils utils.Utils) UserUsecase {
	return &service{
		contextTimeout:  utils.Timeout,
		enforcer:        utils.Enforcer,
		validator:       utils.GoValidator,
		trans:           utils.Translator,
		userPercistance: userPercistance,
	}
}

func (s service) AddUser(c context.Context, u model.User) (*model.User, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	hashedByte, err := bcrypt.GenerateFromPassword([]byte(u.Password), 8)
	if err != nil {
		return nil, errors.New("error during password hashing")
	}
	u.Password = string(hashedByte)
	u.Status = "Pending"
	return s.userPercistance.AddUser(ctx, &u)
}
func (s service) Users(c context.Context, pgnFlt *rest.FilterParams) ([]model.User, error) {
	// ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	// defer cancel()
	return s.userPercistance.Users(c, pgnFlt)
}

func (s service) User(c context.Context, u model.User) (*model.User, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	if uuid.Equal(u.ID, uuid.Nil) && u.Email == "" {
		return nil, errors.New("empty user id or email provided")
	}
	return s.userPercistance.User(ctx, &u)
}

func (s service) UpdateUser(c context.Context, u model.User) (*model.User, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	if uuid.Equal(u.ID, uuid.Nil) {
		return nil, errors.New("empty user id provided")
	}
	hashedByte, err := bcrypt.GenerateFromPassword([]byte(u.Password), 8)
	if err != nil {
		return nil, errors.New("error during password hashing")
	}
	u.Password = string(hashedByte)
	return s.userPercistance.UpdateUser(ctx, &u)
}

func (s service) DeleteUser(c context.Context, u model.User) error {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	if uuid.Equal(u.ID, uuid.Nil) {
		return errors.New("empty user id provided")
	}
	return s.userPercistance.DeleteUser(ctx, &u)
}
