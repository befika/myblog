package subscription

import (
	model "blog/internal/constant/model/db"
	"blog/internal/constant/model/rest"
	"blog/internal/module/subscription"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gravitational/trace"
	uuid "github.com/satori/go.uuid"
)

type BlogSubHandler interface {
	AddBlogSubscription(c *gin.Context)
	BlogSubscriptions(c *gin.Context)
}

type blogsubscription struct {
	BlogSubUsecase subscription.BlogSubUsecase
}

func NewBlogSubHandler(BlogSubUsecase subscription.BlogSubUsecase) BlogSubHandler {
	return &blogsubscription{
		BlogSubUsecase: BlogSubUsecase,
	}
}

func (b blogsubscription) AddBlogSubscription(c *gin.Context) {
	ctx := c.Request.Context()
	uid := uuid.FromStringOrNil(c.Param("user-id"))
	bid := uuid.FromStringOrNil(c.Param("blog-id"))

	bs := model.BlogSubscription{UserID: uid, BlogID: bid}
	err := c.Bind(&bs)
	if err != nil {
		rest.ErrorResponseJson(c, err.Error(), trace.ErrorToCode(err))
		return
	}
	detail, err := b.BlogSubUsecase.AddBlogSubscription(ctx, bs)
	if err != nil {
		rest.ErrorResponseJson(c, err.Error(), trace.ErrorToCode(err))
		return
	}
	rest.SuccessResponseJson(c, nil, detail, http.StatusOK)
}

func (b blogsubscription) BlogSubscriptions(c *gin.Context) {
	detail, err := b.BlogSubUsecase.BlogSubscriptions()
	if err != nil {
		rest.ErrorResponseJson(c, err.Error(), trace.ErrorToCode(err))
		return
	}
	rest.SuccessResponseJson(c, nil, detail, http.StatusOK)
}

// func (ss ssubscription) Subscription(c *gin.Context) {
// 	ctx := c.Request.Context()
// 	id := uuid.FromStringOrNil(c.Param("id"))
// 	sub := model.SystemSubscription{ID: id}
// 	detail, err := ss.SSUsecase.Subscription(ctx, sub)
// 	if err != nil {
// 		rest.ErrorResponseJson(c, err.Error(), trace.ErrorToCode(err))
// 		return
// 	}
// 	rest.SuccessResponseJson(c, nil, detail, http.StatusOK)
// }
