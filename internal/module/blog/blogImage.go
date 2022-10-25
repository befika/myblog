package blog

import (
	"blog/internal/adapter/http/rest/server/image"
	Percistance "blog/internal/adapter/storage/blog"
	model "blog/internal/constant/model/db"
	utils "blog/internal/constant/model/init"
	"blog/internal/constant/model/rest"
	"context"
	"errors"
	"log"
	"mime/multipart"
	"path/filepath"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/casbin/casbin/v2"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gravitational/trace"
)

type BlogImg interface {
	AddBlogImg(ctx context.Context, blogImg *model.BlogImages) (*model.BlogImages, error)
	BlogImages(ctx context.Context, pgnFlt *rest.FilterParams) ([]model.BlogImages, error)
	BlogImg(ctx context.Context, blogImg *model.BlogImages) (*model.BlogImages, error)
}

type BlogImgservice struct {
	contextTimeout    time.Duration
	enforcer          *casbin.Enforcer
	validator         *validator.Validate
	trans             ut.Translator
	blogImgPercistace Percistance.BlogImg
	Store             image.Storage
}

func BlogImgInit(blogImgPercistace Percistance.BlogImg, utils utils.Utils) BlogImg {
	return &BlogImgservice{
		contextTimeout:    utils.Timeout,
		enforcer:          utils.Enforcer,
		validator:         utils.GoValidator,
		trans:             utils.Translator,
		blogImgPercistace: blogImgPercistace,
	}
}
func (b BlogImgservice) AddBlogImg(c context.Context, blogImg *model.BlogImages) (*model.BlogImages, error) {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()
	log.Println("Image size", blogImg.File.Size)
	if blogImg.File == nil {
		return nil, errors.New("uneble to upload blog image")
	} else if blogImg.File.Size > 50000 {
		return nil, errors.New("please upload image less than 10mb")
	}
	id := uuid.NewV4().String()
	blogimage := uuid.NewV4().String() + filepath.Ext(blogImg.File.Filename)
	blogImg.Picture = id + "/" + blogimage
	err := b.saveFile(id, blogimage, blogImg.File)
	if err != nil {
		return nil, err
	}
	return b.blogImgPercistace.AddBlogImg(ctx, blogImg)
}
func (b BlogImgservice) BlogImages(c context.Context, pgnFlt *rest.FilterParams) ([]model.BlogImages, error) {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()
	return b.blogImgPercistace.BlogImages(ctx, pgnFlt)
}
func (b BlogImgservice) BlogImg(c context.Context, blogImg *model.BlogImages) (*model.BlogImages, error) {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()
	if uuid.Equal(blogImg.ID, uuid.Nil) {
		return nil, errors.New("empty blog image id provided")
	}
	return b.blogImgPercistace.BlogImg(ctx, blogImg)
}

func (b BlogImgservice) saveFile(id, path string, f *multipart.FileHeader) error {

	fp := filepath.Join("assets", "images", "book_cover", id, path)
	r, err := f.Open()
	if err != nil {
		return trace.Wrap(err)
	}

	err = b.Store.Save(fp, r)
	if err != nil {
		return trace.Wrap(err)
	}

	return nil
}
