package blog

import (
	model "blog/internal/constant/model/db"
	"blog/internal/constant/model/rest"
	"context"
	"errors"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type BlogImg interface {
	AddBlogImg(ctx context.Context, blogImg *model.BlogImages) (*model.BlogImages, error)
	BlogImages(ctx context.Context, pgnFlt *rest.FilterParams) ([]model.BlogImages, error)
	BlogImg(ctx context.Context, blogImg *model.BlogImages) (*model.BlogImages, error)
}

type blogimgPercistance struct {
	conn *gorm.DB
}

func BlogImgInit(conn *gorm.DB) BlogImg {
	return &blogimgPercistance{
		conn: conn,
	}
}
func (b blogimgPercistance) AddBlogImg(ctx context.Context, blogImg *model.BlogImages) (*model.BlogImages, error) {
	conn := b.conn.WithContext(ctx)
	if uuid.Equal(blogImg.BlogID, uuid.Nil) {
		return nil, errors.New("empty blog id provided")
	}
	err := conn.Model(&model.BlogImages{}).Create(&blogImg).Error
	if err != nil {
		return nil, err
	}
	return blogImg, nil
}
func (b blogimgPercistance) BlogImages(ctx context.Context, pgnFlt *rest.FilterParams) ([]model.BlogImages, error) {
	conn := b.conn.WithContext(ctx)
	blogImgs := []model.BlogImages{}
	err := conn.Model(&model.BlogImages{}).Find(&blogImgs).Error
	if err != nil {
		return nil, err
	}
	if err := conn.Model(&model.BlogImages{}).Offset(-1).Limit(-1).Count(&pgnFlt.Total).Error; err != nil {
		return nil, err
	}
	return blogImgs, nil
}
func (b blogimgPercistance) BlogImg(ctx context.Context, blogImg *model.BlogImages) (*model.BlogImages, error) {
	conn := b.conn.WithContext(ctx)
	err := conn.Model(&model.BlogImages{}).Where("id=?", blogImg.ID).Preload("Pictures").Find(&blogImg).Error
	if err != nil {
		return nil, err
	}
	return blogImg, nil
}
