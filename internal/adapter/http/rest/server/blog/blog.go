package blog

import (
	"blog/internal/adapter/http/rest/middleware"
	model "blog/internal/constant/model/db"
	"blog/internal/constant/model/rest"
	authmodule "blog/internal/module/auth"
	module "blog/internal/module/blog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gravitational/trace"
	uuid "github.com/satori/go.uuid"
)

type BlogHandler interface {
	CreateBlog(c *gin.Context)
	Blogs(c *gin.Context)
	Blog(c *gin.Context)
}
type blogHandler struct {
	blogUsecase   module.BlogUsecase
	authMidleware middleware.AuthMiddleware
	authUseCase   authmodule.LoginUseCase
}

func NewBlogHandler(blogUsecase module.BlogUsecase, authMidleware middleware.AuthMiddleware, authUseCase authmodule.LoginUseCase) BlogHandler {
	return &blogHandler{
		blogUsecase:   blogUsecase,
		authMidleware: authMidleware,
		authUseCase:   authUseCase,
	}
}

func (b blogHandler) CreateBlog(c *gin.Context) {
	ctx := c.Request.Context()

	token := b.authMidleware.ExtractToken(c.Request)
	claims, er := b.authUseCase.GetClaims(token)
	if er != nil {
		rest.ErrorResponseJson(c, er.Error(), trace.ErrorToCode(er))
		return
	}

	blog := model.Blog{UserID: claims.UserId}
	err := c.Bind(&blog)
	if err != nil {
		rest.ErrorResponseJson(c, err.Error(), trace.ErrorToCode(err))
		return
	}

	detail, err := b.blogUsecase.CreateBlog(ctx, blog)
	if err != nil {
		rest.ErrorResponseJson(c, err.Error(), trace.ErrorToCode(err))
		return
	}
	rest.SuccessResponseJson(c, nil, detail, http.StatusOK)
}
func (b blogHandler) Blogs(c *gin.Context) {
	ctx := c.Request.Context()
	pgnFlt, er := rest.ParsePgn(c)
	if er != nil {
		rest.ErrorResponseJson(c, er.Error(), trace.ErrorToCode(er))
		return
	}
	detail, err := b.blogUsecase.Blogs(ctx, pgnFlt)
	if err != nil {
		rest.ErrorResponseJson(c, err.Error(), trace.ErrorToCode(err))
		return
	}
	rest.SuccessResponseJson(c, nil, detail, http.StatusOK)
}
func (b blogHandler) Blog(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	blog := model.Blog{ID: uuid.FromStringOrNil(id)}
	detail, err := b.blogUsecase.Blog(ctx, blog)
	if err != nil {
		rest.ErrorResponseJson(c, err.Error(), trace.ErrorToCode(err))
		return
	}
	rest.SuccessResponseJson(c, nil, detail, http.StatusOK)
}
