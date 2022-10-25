package subscription

import (
	model "blog/internal/constant/model/db"
	"context"

	"gorm.io/gorm"
)

type BlogSubscription interface {
	AddBlogSubscription(ctx context.Context, prm *model.BlogSubscription) (*model.BlogSubscription, error)
	BlogSubscriptions() ([]model.BlogSubscription, error)
	// Subscription(ctx context.Context, sub *model.SystemSubscription) (*model.SystemSubscription, error)
}
type blogSubPercistance struct {
	conn *gorm.DB
}

func BlogSubInit(conn *gorm.DB) BlogSubscription {
	return &blogSubPercistance{
		conn: conn,
	}
}
func (s blogSubPercistance) AddBlogSubscription(ctx context.Context, prm *model.BlogSubscription) (*model.BlogSubscription, error) {
	conn := s.conn.WithContext(ctx)
	err := conn.Model(&model.BlogSubscription{}).Create(&prm).Error
	if err != nil {
		return nil, err
	}
	return prm, nil
}

func (s blogSubPercistance) BlogSubscriptions() ([]model.BlogSubscription, error) {
	// conn := s.conn.WithContext(ctx)
	sub := []model.BlogSubscription{}
	err := s.conn.Model(&model.BlogSubscription{}).Find(&sub).Error
	if err != nil {
		return nil, err
	}
	return sub, nil
}

// func (s ssPercistance) Subscription(ctx context.Context, sub *model.SystemSubscription) (*model.SystemSubscription, error) {
// 	conn := s.conn.WithContext(ctx)
// 	if uuid.Equal(sub.ID, uuid.Nil) {
// 		return nil, errors.New("empty subscription id provided")
// 	}
// 	err := conn.Model(&model.SystemSubscription{}).Where("id=?", sub.ID).First(&sub).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return sub, nil
// }
