package subscription

import (
	model "blog/internal/constant/model/db"
	"blog/internal/constant/model/rest"
	"blog/internal/module/subscription"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gravitational/trace"
	uuid "github.com/satori/go.uuid"
)

type SSHandler interface {
	AddSSubscription(c *gin.Context)
	SSubscriptions(c *gin.Context)
	Subscription(c *gin.Context)
}

type ssubscription struct {
	SSUsecase subscription.SSUsecase
}

func NewSSHandler(SSUsecase subscription.SSUsecase) SSHandler {
	return &ssubscription{
		SSUsecase: SSUsecase,
	}
}

func (ss ssubscription) AddSSubscription(c *gin.Context) {
	ctx := c.Request.Context()
	id := uuid.FromStringOrNil(c.Param("id"))
	price, err := strconv.Atoi(c.Param("price"))
	if err != nil {
		rest.ErrorResponseJson(c, err.Error(), trace.ErrorToCode(err))
		return
	}
	us := model.SystemSubscription{UserID: id, Price: float64(price)}
	err = c.Bind(&us)
	if err != nil {
		rest.ErrorResponseJson(c, err.Error(), trace.ErrorToCode(err))
		return
	}
	detail, err := ss.SSUsecase.AddUserSubscription(ctx, us)
	if err != nil {
		rest.ErrorResponseJson(c, err.Error(), trace.ErrorToCode(err))
		return
	}
	rest.SuccessResponseJson(c, nil, detail, http.StatusOK)
}
func (ss ssubscription) SSubscriptions(c *gin.Context) {
	ctx := c.Request.Context()
	pgnFlt, err := rest.ParsePgn(c)
	if err != nil {
		rest.ErrorResponseJson(c, err.Error(), trace.ErrorToCode(err))
		return
	}
	detail, err := ss.SSUsecase.SSubscriptions(ctx, pgnFlt)
	if err != nil {
		rest.ErrorResponseJson(c, err.Error(), trace.ErrorToCode(err))
		return
	}
	rest.SuccessResponseJson(c, nil, detail, http.StatusOK)
}
func (ss ssubscription) Subscription(c *gin.Context) {
	ctx := c.Request.Context()
	id := uuid.FromStringOrNil(c.Param("id"))
	sub := model.SystemSubscription{ID: id}
	detail, err := ss.SSUsecase.Subscription(ctx, sub)
	if err != nil {
		rest.ErrorResponseJson(c, err.Error(), trace.ErrorToCode(err))
		return
	}
	rest.SuccessResponseJson(c, nil, detail, http.StatusOK)
}
