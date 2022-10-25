package blog

import (
	"blog/internal/adapter/http/rest/server/image"
	dbpgn "blog/internal/constant/db"
	model "blog/internal/constant/model/db"
	"blog/internal/constant/model/rest"
	"context"

	"gorm.io/gorm"
)

type Blog interface {
	CreateBlog(ctx context.Context, blog *model.Blog) (*model.Blog, error)
	Blogs(ctx context.Context, pgnFlt *rest.FilterParams) ([]model.Blog, error)
	Blog(ctx context.Context, blog *model.Blog) (*model.Blog, error)
}

type blogPercistance struct {
	conn  *gorm.DB
	Store image.Storage
}

func BlogInit(conn *gorm.DB) Blog {
	return &blogPercistance{
		conn: conn,
	}
}

func (b blogPercistance) CreateBlog(ctx context.Context, blog *model.Blog) (*model.Blog, error) {
	conn := b.conn.WithContext(ctx)
	err := conn.Model(&model.Blog{}).Create(&blog).Error
	if err != nil {
		return nil, err
	}
	return blog, nil
}
func (b blogPercistance) Blogs(ctx context.Context, pgnFlt *rest.FilterParams) ([]model.Blog, error) {
	conn := b.conn.WithContext(ctx)
	blogs := []model.Blog{}
	err := conn.Model(&model.Blog{}).Preload("Pictures").Scopes(dbpgn.Filter(pgnFlt)).Find(&blogs).Error
	if err != nil {
		return nil, err
	}
	if err := conn.Model(&model.Blog{}).Offset(-1).Limit(-1).Count(&pgnFlt.Total).Error; err != nil {
		return nil, err
	}
	return blogs, nil
}
func (b blogPercistance) Blog(ctx context.Context, blog *model.Blog) (*model.Blog, error) {
	conn := b.conn.WithContext(ctx)
	err := conn.Model(&model.Blog{}).Where("id=?", blog.ID).Preload("Pictures").Find(&blog).Error
	if err != nil {
		return nil, err
	}
	return blog, nil
}
