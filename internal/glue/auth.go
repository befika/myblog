package routing

import (
	"blog/internal/adapter/http/rest/middleware"
	handler "blog/internal/adapter/http/rest/server/auth"
	"blog/internal/constant/model/permission"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(grp *gin.RouterGroup, authMiddleware middleware.AuthMiddleware, permissionHandler handler.PermissionHandler, authhandler handler.AuthHandler) {
	grp.POST("/auth/login", authhandler.Login)
	grp.GET("/auth/permissions", authMiddleware.Authorizer(permission.GetAllPermissions), permissionHandler.GetAllPermissions)
	grp.POST("/auth/permissions", authMiddleware.Authorizer(permission.AddPermissionToRole), permissionHandler.AddPermissionToRole)
	grp.GET("/auth/:role-id/permissions", authMiddleware.Authorizer(permission.GetRolePermissions), permissionHandler.GetRolePermissions)
	grp.DELETE("/auth/:role-id/:permission", authMiddleware.Authorizer(permission.DeleteRolePermission), permissionHandler.DeleteRolePermission)

}
