package role

import (
	model "blog/internal/constant/model/db"
	"blog/internal/constant/model/rest"
	module "blog/internal/module/role"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gravitational/trace"
	uuid "github.com/satori/go.uuid"
)

type RoleHandler interface {
	AddRole(c *gin.Context)
	Roles(c *gin.Context)
	Role(c *gin.Context)
	UpdateRole(c *gin.Context)
	DeleteRole(c *gin.Context)
	AssignRole(c *gin.Context)
	GetUserRoles(c *gin.Context)
}

type roleHandler struct {
	roleUseCase module.RoleUsecase
}

func NewRoleHandler(roleUseCase module.RoleUsecase) RoleHandler {
	return &roleHandler{
		roleUseCase: roleUseCase,
	}
}

func (r roleHandler) AddRole(c *gin.Context) {
	ctx := c.Request.Context()
	role := model.Role{}
	err := c.Bind(&role)
	if err != nil {
		rest.ErrorResponseJson(c, err, trace.ErrorToCode(err))
		return
	}
	detail, err := r.roleUseCase.AddRole(ctx, role)
	if err != nil {
		rest.ErrorResponseJson(c, err, trace.ErrorToCode(err))
		return
	}
	rest.SuccessResponseJson(c, nil, detail, http.StatusOK)

}
func (r roleHandler) Roles(c *gin.Context) {
	ctx := c.Request.Context()
	pgnFlt, err := rest.ParsePgn(c)
	if err != nil {
		rest.ErrorResponseJson(c, err, trace.ErrorToCode(err))
		return
	}
	detail, err := r.roleUseCase.Roles(ctx, pgnFlt)
	if err != nil {
		rest.ErrorResponseJson(c, err, trace.ErrorToCode(err))
		return
	}
	rest.SuccessResponseJson(c, nil, detail, http.StatusOK)

}
func (r roleHandler) Role(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	role := model.Role{ID: uuid.FromStringOrNil(id)}
	detail, err := r.roleUseCase.Role(ctx, role)
	if err != nil {
		rest.ErrorResponseJson(c, err, trace.ErrorToCode(err))
		return
	}
	rest.SuccessResponseJson(c, nil, detail, http.StatusOK)
}

func (r roleHandler) UpdateRole(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	role := model.Role{ID: uuid.FromStringOrNil(id)}
	detail, err := r.roleUseCase.UpdateRole(ctx, role)
	if err != nil {
		rest.ErrorResponseJson(c, err, trace.ErrorToCode(err))
		return
	}
	rest.SuccessResponseJson(c, nil, detail, http.StatusOK)

}

func (r roleHandler) DeleteRole(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	role := model.Role{ID: uuid.FromStringOrNil(id)}
	err := r.roleUseCase.DeleteRole(ctx, role)
	if err != nil {
		rest.ErrorResponseJson(c, err, trace.ErrorToCode(err))
		return
	}
	rest.SuccessResponseJson(c, nil, "role has been deleted", http.StatusOK)

}
func (r roleHandler) AssignRole(c *gin.Context) {
	ctx := c.Request.Context()
	roleid := uuid.FromStringOrNil(c.Param("role-id"))
	userid := uuid.FromStringOrNil(c.Param("user-id"))
	userRole := &model.UserRoles{UserID: userid, RoleID: roleid}
	err := c.Bind(&userRole)
	if err != nil {
		rest.ErrorResponseJson(c, err, trace.ErrorToCode(err))
		return
	}
	Role := &model.Role{ID: roleid}
	User := &model.User{ID: userid}
	_, err = r.roleUseCase.AssignRole(ctx, *Role, *User)
	if err != nil {
		rest.ErrorResponseJson(c, err, trace.ErrorToCode(err))
		return
	}
	rest.SuccessResponseJson(c, nil, userRole, http.StatusOK)
}
func (r roleHandler) GetUserRoles(c *gin.Context) {
	ctx := c.Request.Context()
	id := uuid.FromStringOrNil(c.Param("user-id"))
	userRoleX := &model.UserRoles{UserID: id}
	detail, err := r.roleUseCase.GetUserRoles(ctx, *userRoleX)
	if err != nil {
		rest.ErrorResponseJson(c, err, trace.ErrorToCode(err))
	}
	rest.SuccessResponseJson(c, nil, detail, http.StatusOK)
}
