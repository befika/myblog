package user

import (
	dbpgn "blog/internal/constant/db"
	model "blog/internal/constant/model/db"
	rest "blog/internal/constant/model/rest"
	"context"

	"gorm.io/gorm"
)

type User interface {
	AddUser(ctx context.Context, prm *model.User) (*model.User, error)
	Users(ctx context.Context, pgnFlt *rest.FilterParams) ([]model.User, error)
	User(ctx context.Context, prm *model.User) (*model.User, error)
	UpdateUser(ctx context.Context, prm *model.User) (*model.User, error)
	DeleteUser(ctx context.Context, prm *model.User) error
}

type userPercistance struct {
	conn *gorm.DB
}

func UserInit(conn *gorm.DB) User {
	return &userPercistance{
		conn: conn,
	}
}

func (u userPercistance) AddUser(ctx context.Context, prm *model.User) (*model.User, error) {
	conn := u.conn.WithContext(ctx)
	err := conn.Model(&model.User{}).Create(&prm).Error
	if err != nil {
		return nil, err
	}
	return prm, nil
}
func (u userPercistance) Users(ctx context.Context, pgnFlt *rest.FilterParams) ([]model.User, error) {
	conn := u.conn.WithContext(ctx)
	users := []model.User{}
	err := conn.Model(&model.User{}).Preload("Roles").Scopes(dbpgn.Filter(pgnFlt)).Find(&users).Error
	if err != nil {
		return nil, err
	}
	if err := conn.Model(&model.User{}).Offset(-1).Limit(-1).Count(&pgnFlt.Total).Error; err != nil {
		return nil, err
	}
	return users, nil
}
func (u userPercistance) User(ctx context.Context, prm *model.User) (*model.User, error) {
	conn := u.conn.WithContext(ctx)
	err := conn.Model(&model.User{}).Where("id=?", prm.ID).Or("email=?", prm.Email).Preload("Roles").First(&prm).Error
	if err != nil {
		return nil, err
	}
	return prm, nil
}

func (u userPercistance) UpdateUser(ctx context.Context, prm *model.User) (*model.User, error) {
	conn := u.conn.WithContext(ctx)
	err := conn.Model(&model.User{}).Where("id=?", prm.ID).Updates(&prm).Error
	if err != nil {
		return nil, err
	}
	return prm, nil
}
func (u userPercistance) DeleteUser(ctx context.Context, prm *model.User) error {
	conn := u.conn.WithContext(ctx)
	err := conn.Model(&model.User{}).Where("id=?", prm.ID).Delete(&prm).Error
	if err != nil {
		return err
	}
	return nil
}
