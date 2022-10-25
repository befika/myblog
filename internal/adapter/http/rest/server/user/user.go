package user

import (
	model "blog/internal/constant/model/db"
	"blog/internal/constant/model/rest"
	module "blog/internal/module/user"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gravitational/trace"
	uuid "github.com/satori/go.uuid"
)

type UserHandler interface {
	AddUser(c *gin.Context)
	Users(c *gin.Context)
	User(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type userHandler struct {
	userUsecase module.UserUsecase
}

func NewUserHandler(userUsecase module.UserUsecase) UserHandler {
	return &userHandler{
		userUsecase: userUsecase,
	}
}
func (u userHandler) AddUser(c *gin.Context) {
	ctx := c.Request.Context()
	user := model.User{}
	err := c.Bind(&user)
	if err != nil {
		rest.ErrorResponseJson(c, err, trace.ErrorToCode(err))
	}
	detail, err := u.userUsecase.AddUser(ctx, user)
	if err != nil {
		rest.ErrorResponseJson(c, err, trace.ErrorToCode(err))
	}
	rest.SuccessResponseJson(c, nil, detail, http.StatusOK)

}

func (u userHandler) Users(c *gin.Context) {
	ctx := c.Request.Context()
	pgnFlt, er := rest.ParsePgn(c)
	if er != nil {
		rest.ErrorResponseJson(c, er, trace.ErrorToCode(er))
		return
	}
	detail, err := u.userUsecase.Users(ctx, pgnFlt)
	if err != nil {
		log.Println(trace.ErrorToCode(er))
		rest.ErrorResponseJson(c, er, trace.ErrorToCode(er))
		return
	}
	rest.SuccessResponseJson(c, pgnFlt, detail, http.StatusOK)

}

func (u userHandler) User(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	user := model.User{ID: uuid.FromStringOrNil(id)}
	detail, err := u.userUsecase.User(ctx, user)
	if err != nil {
		rest.ErrorResponseJson(c, err.Error(), trace.ErrorToCode(err))
		return
	}
	rest.SuccessResponseJson(c, nil, detail, http.StatusOK)

}

func (u userHandler) UpdateUser(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	user := model.User{ID: uuid.FromStringOrNil(id)}
	err := c.Bind(&user)
	if err != nil {
		rest.ErrorResponseJson(c, err, trace.ErrorToCode(err))
	}
	detail, err := u.userUsecase.UpdateUser(ctx, user)
	if err != nil {
		rest.ErrorResponseJson(c, err.Error(), trace.ErrorToCode(err))
		return
	}
	rest.SuccessResponseJson(c, nil, detail, http.StatusOK)

}

func (u userHandler) DeleteUser(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	user := model.User{ID: uuid.FromStringOrNil(id)}
	err := u.userUsecase.DeleteUser(ctx, user)
	if err != nil {
		rest.ErrorResponseJson(c, err, trace.ErrorToCode(err))
	}
	rest.SuccessResponseJson(c, nil, "user has been deleted", http.StatusOK)

}
