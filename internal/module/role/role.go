package role

import (
	Percistance "blog/internal/adapter/storage/role"
	model "blog/internal/constant/model/db"
	utils "blog/internal/constant/model/init"
	"blog/internal/constant/model/rest"
	"context"
	"errors"
	"time"

	ut "github.com/go-playground/universal-translator"
	uuid "github.com/satori/go.uuid"

	"github.com/casbin/casbin/v2"
	"github.com/go-playground/validator/v10"
)

type RoleUsecase interface {
	AddRole(c context.Context, r model.Role) (*model.Role, error)
	Roles(c context.Context, pgnFlt *rest.FilterParams) ([]model.Role, error)
	Role(c context.Context, r model.Role) (*model.Role, error)
	UpdateRole(c context.Context, r model.Role) (*model.Role, error)
	DeleteRole(c context.Context, r model.Role) error
	AssignRole(c context.Context, role model.Role, user model.User) (*model.UserRoles, error)
	GetUserRoles(c context.Context, userroles model.UserRoles) ([]model.UserRoles, error)
}

type service struct {
	contextTimeout  time.Duration
	enforcer        *casbin.Enforcer
	validator       *validator.Validate
	trans           ut.Translator
	rolePercistance Percistance.Role
}

func InitializeRole(rolePercistance Percistance.Role, utils utils.Utils) RoleUsecase {
	return &service{
		contextTimeout:  utils.Timeout,
		enforcer:        utils.Enforcer,
		validator:       utils.GoValidator,
		trans:           utils.Translator,
		rolePercistance: rolePercistance,
	}
}

func (s service) AddRole(c context.Context, r model.Role) (*model.Role, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.rolePercistance.AddRole(ctx, &r)
}
func (s service) Roles(c context.Context, pgnFlt *rest.FilterParams) ([]model.Role, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.rolePercistance.Roles(ctx, pgnFlt)
}

func (s service) Role(c context.Context, r model.Role) (*model.Role, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	if uuid.Equal(r.ID, uuid.Nil) {
		return nil, errors.New("empty role id provided")
	}
	return s.rolePercistance.Role(ctx, &r)
}
func (s service) UpdateRole(c context.Context, r model.Role) (*model.Role, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.rolePercistance.UpdateRole(ctx, &r)
}

func (s service) DeleteRole(c context.Context, r model.Role) error {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.rolePercistance.DeleteRole(ctx, &r)
}
func (s service) GetUserRoles(c context.Context, userroles model.UserRoles) ([]model.UserRoles, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	if uuid.Equal(userroles.UserID, uuid.Nil) {
		return nil, errors.New("empty userId provided")
	}
	return s.rolePercistance.GetUserRoles(ctx, userroles)
}
func (s service) AssignRole(c context.Context, role model.Role, user model.User) (*model.UserRoles, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	if uuid.Equal(role.ID, uuid.Nil) {
		return nil, errors.New("empty roleId provided")
	} else if uuid.Equal(user.ID, uuid.Nil) {
		return nil, errors.New("empty userId provided")
	}
	return s.rolePercistance.AssignRole(ctx, role, user)
}
