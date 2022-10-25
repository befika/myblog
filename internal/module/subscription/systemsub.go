package subscription

import (
	Percistance "blog/internal/adapter/storage/subscription"
	model "blog/internal/constant/model/db"
	utils "blog/internal/constant/model/init"
	"blog/internal/constant/model/rest"
	"context"
	"errors"
	"time"

	"github.com/casbin/casbin/v2"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
)

type SSUsecase interface {
	AddUserSubscription(ctx context.Context, sub model.SystemSubscription) (*model.SystemSubscription, error)
	Subscription(ctx context.Context, sub model.SystemSubscription) (*model.SystemSubscription, error)
	SSubscriptions(ctx context.Context, pgnFlt *rest.FilterParams) ([]model.SystemSubscription, error)
}
type service struct {
	contextTimeout time.Duration
	enforcer       *casbin.Enforcer
	validator      *validator.Validate
	trans          ut.Translator
	ssPercistance  Percistance.SystemSubscription
}

func InitializeSubscription(ssPercistance Percistance.SystemSubscription, utils utils.Utils) SSUsecase {
	return &service{
		contextTimeout: utils.Timeout,
		enforcer:       utils.Enforcer,
		validator:      utils.GoValidator,
		trans:          utils.Translator,
		ssPercistance:  ssPercistance,
	}
}

func (s service) AddUserSubscription(c context.Context, sub model.SystemSubscription) (*model.SystemSubscription, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.ssPercistance.AddUserSubscription(ctx, &sub)
}
func (s service) SSubscriptions(c context.Context, pgnFlt *rest.FilterParams) ([]model.SystemSubscription, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.ssPercistance.SSubscriptions(ctx, pgnFlt)
}
func (s service) Subscription(c context.Context, sub model.SystemSubscription) (*model.SystemSubscription, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	if uuid.Equal(sub.ID, uuid.Nil) {
		return nil, errors.New("empty subscription id provided")
	}
	return s.ssPercistance.Subscription(ctx, &sub)
}
