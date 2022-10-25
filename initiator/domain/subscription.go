package domain

import (
	SubHandler "blog/internal/adapter/http/rest/server/subscription"
	"blog/internal/adapter/storage/subscription"
	utils "blog/internal/constant/model/init"
	routing "blog/internal/glue"
	SubUsecase "blog/internal/module/subscription"

	"github.com/gin-gonic/gin"
)

func SubscriptionInit(utils utils.Utils, router *gin.RouterGroup) {
	SSubPercistance := subscription.SSInit(utils.Conn)
	SSubUsecase := SubUsecase.InitializeSubscription(SSubPercistance, utils)
	SSubHandler := SubHandler.NewSSHandler(SSubUsecase)

	BlogSubPercistance := subscription.BlogSubInit(utils.Conn)
	BlogSubUsecase := SubUsecase.InitializeBlogSubscription(BlogSubPercistance, utils)
	blogsubHandler := SubHandler.NewBlogSubHandler(BlogSubUsecase)

	routing.SubRoutes(router, SSubHandler)
	routing.BlogSubRoutes(router, blogsubHandler)
}
