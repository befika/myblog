package routing

import (
	"blog/internal/adapter/http/rest/middleware"
	handler "blog/internal/adapter/http/rest/server/user"
	"blog/internal/constant/model/permission"

	"github.com/gin-gonic/gin"
)

func UserRoutes(grp *gin.RouterGroup, authMiddleWare middleware.AuthMiddleware, userHandler handler.UserHandler) {
	grp.POST("/users/register", authMiddleWare.Authorizer(permission.CreateUser), userHandler.AddUser)
	grp.GET("/users", authMiddleWare.Authorizer(permission.GetUsers), userHandler.Users)
	grp.GET("/users/:id", authMiddleWare.Authorizer(permission.GetUser), userHandler.User)
	grp.PUT("/users/:id", authMiddleWare.Authorizer(permission.UpdateUser), userHandler.UpdateUser)
	grp.DELETE("/users/:id", authMiddleWare.Authorizer(permission.DeleteUser), userHandler.DeleteUser)
}
