package subscription

import (
	Percistance "blog/internal/adapter/storage/subscription"
	model "blog/internal/constant/model/db"
	utils "blog/internal/constant/model/init"
	"context"
	"time"

	"github.com/casbin/casbin/v2"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type BlogSubUsecase interface {
	AddBlogSubscription(ctx context.Context, sub model.BlogSubscription) (*model.BlogSubscription, error)
	BlogSubscriptions() ([]model.BlogSubscription, error)
}
type blogSubservice struct {
	contextTimeout   time.Duration
	enforcer         *casbin.Enforcer
	validator        *validator.Validate
	trans            ut.Translator
	blogSubscription Percistance.BlogSubscription
}

func InitializeBlogSubscription(blogSubscription Percistance.BlogSubscription, utils utils.Utils) BlogSubUsecase {
	return &blogSubservice{
		contextTimeout:   utils.Timeout,
		enforcer:         utils.Enforcer,
		validator:        utils.GoValidator,
		trans:            utils.Translator,
		blogSubscription: blogSubscription,
	}
}

func (s blogSubservice) AddBlogSubscription(c context.Context, sub model.BlogSubscription) (*model.BlogSubscription, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.blogSubscription.AddBlogSubscription(ctx, &sub)
}
func (s blogSubservice) BlogSubscriptions() ([]model.BlogSubscription, error) {
	return s.blogSubscription.BlogSubscriptions()
}
