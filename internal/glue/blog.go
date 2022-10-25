package routing

import (
	"blog/internal/adapter/http/rest/middleware"
	handler "blog/internal/adapter/http/rest/server/blog"
	"blog/internal/constant/model/permission"

	"github.com/gin-gonic/gin"
)

func BlogRoutes(grp *gin.RouterGroup, authMiddleware middleware.AuthMiddleware, blgImgHandler handler.BlogImgHandler, blogHandler handler.BlogHandler) {
	grp.POST("/blogs/register", authMiddleware.Authorizer(permission.CreateBlog), blogHandler.CreateBlog)
	grp.GET("/blogs", authMiddleware.Authorizer(permission.GetBlogs), blogHandler.Blogs)
	grp.GET("/blogs/:id", authMiddleware.Authorizer(permission.GetBlog), blogHandler.Blog)

	grp.POST("/blogs/:id/images/register", blgImgHandler.AddBlogImg)
	grp.GET("/blogs/images", blgImgHandler.BlogImages)
	grp.GET("/blogs/images/:id", blgImgHandler.BlogImg)
}
