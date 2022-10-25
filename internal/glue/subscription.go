package routing

import (
	handler "blog/internal/adapter/http/rest/server/subscription"

	"github.com/gin-gonic/gin"
)

func SubRoutes(grp *gin.RouterGroup, ssubHandler handler.SSHandler) {
	grp.POST("/users/:id/subscriptions/:price", ssubHandler.AddSSubscription)
	grp.GET("/users/subscriptions", ssubHandler.SSubscriptions)
	grp.GET("/users/subscriptions/:id", ssubHandler.Subscription)

}
func BlogSubRoutes(grp *gin.RouterGroup, blogSubHandler handler.BlogSubHandler) {
	grp.POST("/blogs/blogsubscription/:blog-id/:user-id", blogSubHandler.AddBlogSubscription)
	grp.GET("/blogs/blogsubscription", blogSubHandler.BlogSubscriptions)

}
