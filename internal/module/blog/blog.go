package blog

import (
	Percistance "blog/internal/adapter/storage/blog"
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

type BlogUsecase interface {
	CreateBlog(c context.Context, blog model.Blog) (*model.Blog, error)
	Blog(c context.Context, blog model.Blog) (*model.Blog, error)
	Blogs(c context.Context, pgnFlt *rest.FilterParams) ([]model.Blog, error)
}

type service struct {
	contextTimeout time.Duration
	enforcer       *casbin.Enforcer
	validator      *validator.Validate
	trans          ut.Translator
	blogPercistace Percistance.Blog
}

func BlogInit(blogPercistace Percistance.Blog, utils utils.Utils) BlogUsecase {
	return &service{
		contextTimeout: utils.Timeout,
		enforcer:       utils.Enforcer,
		validator:      utils.GoValidator,
		trans:          utils.Translator,
		blogPercistace: blogPercistace,
	}
}

func (b service) CreateBlog(c context.Context, blog model.Blog) (*model.Blog, error) {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()
	return b.blogPercistace.CreateBlog(ctx, &blog)
}
func (b service) Blogs(c context.Context, pgnFlt *rest.FilterParams) ([]model.Blog, error) {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()
	return b.blogPercistace.Blogs(ctx, pgnFlt)
}
func (b service) Blog(c context.Context, blog model.Blog) (*model.Blog, error) {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()
	if uuid.Equal(blog.ID, uuid.Nil) {
		return nil, errors.New("empty blog id provided")
	}
	return b.blogPercistace.Blog(ctx, &blog)
}
