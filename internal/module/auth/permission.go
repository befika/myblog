package auth

import (
	model "blog/internal/constant/model/db"
	utils "blog/internal/constant/model/init"
	"blog/internal/constant/model/permission"

	ut "github.com/go-playground/universal-translator"

	"fmt"
	"time"

	"github.com/casbin/casbin/v2"
	"github.com/go-playground/validator/v10"
	"github.com/gofrs/uuid"
)

type PermissionUseCase interface {
	GetAllPermissions() ([]model.RolePermission, error)
	AddPermissionToRole(permissions model.RolePermission) error
	MigratePermissionsToCasbin() error
	GetRolePermissions(roleID uuid.UUID) ([]model.RolePermission, error)
	DeleteRolePermission(perm model.RolePermission) error
}

type permissionPercistace struct {
	contextTimeout time.Duration
	enforcer       *casbin.Enforcer
	validator      *validator.Validate
	translator     ut.Translator
}

func PermissionInit(utils utils.Utils) PermissionUseCase {
	return &permissionPercistace{
		contextTimeout: utils.Timeout,
		enforcer:       utils.Enforcer,
		validator:      utils.GoValidator,
		translator:     utils.Translator,
	}
}

func (p permissionPercistace) GetAllPermissions() ([]model.RolePermission, error) {
	prmList := permission.PermissionObjects
	prms := []model.RolePermission{}
	for prm, _ := range prmList {

		prms = append(prms, model.RolePermission{
			Name: prm,
		})
	}
	return prms, nil
}

func (p permissionPercistace) AddPermissionToRole(permissions model.RolePermission) error {

	prm := model.RolePermission{
		Role:   permissions.Role,
		Object: permission.PermissionObjects[permissions.Name],
		Action: permission.PermissionActions[permissions.Name],
	}
	fmt.Println(prm)

	_, err := p.enforcer.AddPolicy(prm.Role, prm.Object, prm.Action)

	if err != nil {
		return err
	}

	return nil
}

func (p permissionPercistace) MigratePermissionsToCasbin() error {

	for prm, _ := range permission.PermissionObjects {

		p.enforcer.AddPolicy("3e74763d-abf3-46d3-806b-d9dfcc4fd36f", permission.PermissionObjects[prm], permission.PermissionActions[prm])
	}

	return nil
}
func (p permissionPercistace) DeleteRolePermission(perm model.RolePermission) error {
	_, err := p.enforcer.DeletePermissionForUser(perm.Role, perm.Object, perm.Action)
	if err != nil {
		return err
	}
	return nil
}
func (p permissionPercistace) GetRolePermissions(roleID uuid.UUID) ([]model.RolePermission, error) {
	perm := []model.RolePermission{}
	permissions := p.enforcer.GetPermissionsForUser(roleID.String())
	for _, prm := range permissions {
		permission := model.RolePermission{}
		permission.Name = prm[1]
		permission.Action = prm[2]
		permission.Object = prm[2]
		perm = append(perm, permission)
	}

	return perm, nil
}
