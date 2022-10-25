package subscription

import (
	dbpgn "blog/internal/constant/db"
	model "blog/internal/constant/model/db"
	"blog/internal/constant/model/rest"
	"context"

	"gorm.io/gorm"
)

type SystemSubscription interface {
	AddUserSubscription(ctx context.Context, sub *model.SystemSubscription) (*model.SystemSubscription, error)
	SSubscriptions(ctx context.Context, pgnFlt *rest.FilterParams) ([]model.SystemSubscription, error)
	Subscription(ctx context.Context, sub *model.SystemSubscription) (*model.SystemSubscription, error)
}
type ssPercistance struct {
	conn *gorm.DB
}

func SSInit(conn *gorm.DB) SystemSubscription {
	return &ssPercistance{
		conn: conn,
	}
}
func (s ssPercistance) AddUserSubscription(ctx context.Context, sub *model.SystemSubscription) (*model.SystemSubscription, error) {
	conn := s.conn.WithContext(ctx)
	err := conn.Transaction(func(tx *gorm.DB) error {
		err := conn.Model(&model.SystemSubscription{}).Create(&sub).Error
		if err != nil {
			return err
		}
		err = conn.Model(&model.User{}).Where(&model.User{ID: sub.UserID}).Update("status", "Active").Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return sub, nil
}
func (s ssPercistance) SSubscriptions(ctx context.Context, pgnFlt *rest.FilterParams) ([]model.SystemSubscription, error) {
	conn := s.conn.WithContext(ctx)
	sub := []model.SystemSubscription{}
	err := conn.Model(&model.SystemSubscription{}).Scopes(dbpgn.Filter(pgnFlt)).Find(&sub).Error
	if err != nil {
		return nil, err
	}
	if err := conn.Model(&model.SystemSubscription{}).Offset(-1).Limit(-1).Count(&pgnFlt.Total).Error; err != nil {
		return nil, err
	}
	return sub, nil
}

func (s ssPercistance) Subscription(ctx context.Context, sub *model.SystemSubscription) (*model.SystemSubscription, error) {
	conn := s.conn.WithContext(ctx)
	err := conn.Model(&model.SystemSubscription{}).Where("id=?", sub.ID).First(&sub).Error
	if err != nil {
		return nil, err
	}
	return sub, nil
}
