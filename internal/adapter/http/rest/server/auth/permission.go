package auth

import (
	"net/http"

	model "blog/internal/constant/model/db"
	utils "blog/internal/constant/model/init"
	"blog/internal/constant/model/rest"
	module "blog/internal/module/auth"

	ut "github.com/go-playground/universal-translator"
	"github.com/gofrs/uuid"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gravitational/trace"
)

type PermissionHandler interface {
	GetAllPermissions(c *gin.Context)
	AddPermissionToRole(c *gin.Context)
	GetRolePermissions(c *gin.Context)
	DeleteRolePermission(c *gin.Context)
}
type permissionHandler struct {
	permissionUsecase module.PermissionUseCase
	validate          *validator.Validate
	trans             ut.Translator
}

func NewPermissionHandler(permissionUsecase module.PermissionUseCase, utils utils.Utils) PermissionHandler {
	return &permissionHandler{
		permissionUsecase: permissionUsecase,
		validate:          utils.GoValidator,
		trans:             utils.Translator,
	}
}

func (p permissionHandler) GetAllPermissions(c *gin.Context) {
	prms, err := p.permissionUsecase.GetAllPermissions()
	if err != nil {
		rest.ErrorResponseJson(c, err.Error(), trace.ErrorToCode(err))
		return
	}
	rest.SuccessResponseJson(c, nil, prms, http.StatusOK)
}
func (p permissionHandler) AddPermissionToRole(c *gin.Context) {
	// roleid := uuid.FromStringOrNil(c.Param("role-id"))
	perm := model.RolePermission{}
	err := c.Bind(&perm)
	if err != nil {
		rest.ErrorResponseJson(c, err.Error(), trace.ErrorToCode(err))
		return
	}
	err = p.permissionUsecase.AddPermissionToRole(perm)
	if err != nil {
		rest.ErrorResponseJson(c, err.Error(), trace.ErrorToCode(err))
		return
	}
	rest.SuccessResponseJson(c, nil, perm, http.StatusOK)
}
func (p permissionHandler) GetRolePermissions(c *gin.Context) {
	roleID := uuid.FromStringOrNil(c.Param("role-id"))
	perm, err := p.permissionUsecase.GetRolePermissions(roleID)
	if err != nil {
		rest.ErrorResponseJson(c, err.Error(), trace.ErrorToCode(err))
		return
	}
	rest.SuccessResponseJson(c, nil, perm, http.StatusOK)
}
func (p permissionHandler) DeleteRolePermission(c *gin.Context) {
	roleID := c.Param("role-id")
	perm := c.Param("permission")
	permission := model.RolePermission{Role: roleID, Object: perm}
	err := p.permissionUsecase.DeleteRolePermission(permission)
	if err != nil {
		rest.ErrorResponseJson(c, err.Error(), trace.ErrorToCode(err))
		return
	}
	rest.SuccessResponseJson(c, nil, "permission deleted from role", http.StatusOK)
}
