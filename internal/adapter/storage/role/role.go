package role

import (
	dbpgn "blog/internal/constant/db"
	model "blog/internal/constant/model/db"
	"blog/internal/constant/model/rest"
	"context"

	"gorm.io/gorm"
)

type Role interface {
	AddRole(ctx context.Context, prm *model.Role) (*model.Role, error)
	Roles(ctx context.Context, pgnFlt *rest.FilterParams) ([]model.Role, error)
	Role(ctx context.Context, prm *model.Role) (*model.Role, error)
	UpdateRole(ctx context.Context, prm *model.Role) (*model.Role, error)
	DeleteRole(ctx context.Context, prm *model.Role) error
	AssignRole(ctx context.Context, role model.Role, user model.User) (*model.UserRoles, error)
	GetUserRoles(ctx context.Context, userroles model.UserRoles) ([]model.UserRoles, error)
}

type rolePercistace struct {
	conn *gorm.DB
}

func RoleInit(conn *gorm.DB) Role {
	return &rolePercistace{
		conn: conn,
	}
}

func (r rolePercistace) AddRole(ctx context.Context, prm *model.Role) (*model.Role, error) {
	conn := r.conn.WithContext(ctx)
	err := conn.Model(&model.Role{}).Create(&prm).Error
	if err != nil {
		return nil, err
	}
	return prm, nil
}
func (r rolePercistace) Roles(ctx context.Context, pgnFlt *rest.FilterParams) ([]model.Role, error) {
	conn := r.conn.WithContext(ctx)
	roles := []model.Role{}
	err := conn.Model(&model.Role{}).Scopes(dbpgn.Filter(pgnFlt)).Find(&roles).Error
	if err != nil {
		return nil, err
	}
	if err := conn.Model(&model.Role{}).Offset(-1).Limit(-1).Count(&pgnFlt.Total).Error; err != nil {
		return nil, err
	}
	return roles, nil

}
func (r rolePercistace) Role(ctx context.Context, prm *model.Role) (*model.Role, error) {
	conn := r.conn.WithContext(ctx)
	err := conn.Model(&model.Role{}).Where("id=?", prm.ID).Find(&prm).Error
	if err != nil {
		return nil, err
	}
	return prm, nil

}
func (r rolePercistace) UpdateRole(ctx context.Context, prm *model.Role) (*model.Role, error) {
	conn := r.conn.WithContext(ctx)
	err := conn.Model(&model.Role{}).Where("id=?", prm.ID).Updates(&prm).Error
	if err != nil {
		return nil, err
	}
	return prm, nil

}
func (r rolePercistace) DeleteRole(ctx context.Context, prm *model.Role) error {
	conn := r.conn.WithContext(ctx)
	err := conn.Model(&model.Role{}).Where("id=?", prm.ID).Delete(&prm).Error
	if err != nil {
		return err
	}
	return nil
}
func (r rolePercistace) GetUserRoles(ctx context.Context, userroles model.UserRoles) ([]model.UserRoles, error) {
	conn := r.conn.WithContext(ctx)
	userRoles := []model.UserRoles{}
	err := conn.Where("user_id=?", userroles.UserID).Find(&userRoles).Error
	if err != nil {
		return nil, err
	}
	return userRoles, nil
}

func (r rolePercistace) AssignRole(ctx context.Context, role model.Role, user model.User) (*model.UserRoles, error) {
	conn := r.conn.WithContext(ctx)
	userRoles := &model.UserRoles{UserID: user.ID, RoleID: role.ID}
	err := conn.Where(&model.UserRoles{}).Create(userRoles).Error
	if err != nil {
		return nil, err
	}
	return userRoles, nil
}
