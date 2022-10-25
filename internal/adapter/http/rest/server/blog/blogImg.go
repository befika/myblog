package blog

import (
	model "blog/internal/constant/model/db"
	"blog/internal/constant/model/rest"
	module "blog/internal/module/blog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gravitational/trace"
	uuid "github.com/satori/go.uuid"
)

type BlogImgHandler interface {
	AddBlogImg(c *gin.Context)
	BlogImages(c *gin.Context)
	BlogImg(c *gin.Context)
}
type blogImgHandler struct {
	blogImgUsecase module.BlogImg
}

func NewBlogImgHandler(blogImgUsecase module.BlogImg) BlogImgHandler {
	return &blogImgHandler{
		blogImgUsecase: blogImgUsecase,
	}
}

func (b blogImgHandler) AddBlogImg(c *gin.Context) {
	ctx := c.Request.Context()
	blogID := c.Param("id")
	blogImg := model.BlogImages{BlogID: uuid.FromStringOrNil(blogID)}
	err := c.Bind(&blogImg)
	if err != nil {
		if err != nil {
			c.JSON(http.StatusBadRequest, trace.Wrap(err))
		}
	}
	detail, err := b.blogImgUsecase.AddBlogImg(ctx, &blogImg)
	if err != nil {
		c.JSON(http.StatusBadRequest, trace.Wrap(err))
	}
	c.JSON(http.StatusOK, detail)
}
func (b blogImgHandler) BlogImages(c *gin.Context) {
	ctx := c.Request.Context()
	pgnFlt, er := rest.ParsePgn(c)
	if er != nil {
		rest.ErrorResponseJson(c, er, trace.ErrorToCode(er))
		return
	}
	detail, err := b.blogImgUsecase.BlogImages(ctx, pgnFlt)
	if err != nil {
		c.JSON(http.StatusBadRequest, trace.Wrap(err))
	}
	c.JSON(http.StatusOK, detail)
}
func (b blogImgHandler) BlogImg(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	blog := model.BlogImages{ID: uuid.FromStringOrNil(id)}
	detail, err := b.blogImgUsecase.BlogImg(ctx, &blog)
	if err != nil {
		c.JSON(http.StatusBadRequest, trace.Wrap(err))
	}
	c.JSON(http.StatusOK, detail)
}
