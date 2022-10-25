package auth

import (
	model "blog/internal/constant/model/db"
	"blog/internal/constant/model/rest"
	"blog/internal/module/auth"
	"blog/internal/module/role"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gravitational/trace"
)

type AuthHandler interface {
	Login(c *gin.Context)
}
type authHandler struct {
	authUsecase auth.LoginUseCase
	roleUsecase role.RoleUsecase
}

func NewAuthHandler(authUsecase auth.LoginUseCase, roleUsecase role.RoleUsecase) AuthHandler {
	return authHandler{
		authUsecase: authUsecase,
		roleUsecase: roleUsecase,
	}
}

func (a authHandler) Login(c *gin.Context) {
	ctx := c.Request.Context()
	authUser := model.User{}
	err := c.Bind(&authUser)
	if err != nil {
		rest.ErrorResponseJson(c, err.Error(), trace.ErrorToCode(err))
		return
	}
	res, err := a.authUsecase.Login(ctx, authUser.Email, authUser.Password)
	if err != nil {
		rest.ErrorResponseJson(c, err.Error(), trace.ErrorToCode(err))
		return
	}
	rest.SuccessResponseJson(c, nil, res, http.StatusOK)
}
