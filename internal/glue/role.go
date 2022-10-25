package routing

import (
	"blog/internal/adapter/http/rest/middleware"
	handler "blog/internal/adapter/http/rest/server/role"
	"blog/internal/constant/model/permission"

	"github.com/gin-gonic/gin"
)

func RoleRoutes(grp *gin.RouterGroup, authMiddleware middleware.AuthMiddleware, roleHandler handler.RoleHandler) {
	grp.POST("/roles/register", authMiddleware.Authorizer(permission.CreateRole), roleHandler.AddRole)
	grp.GET("/roles", authMiddleware.Authorizer(permission.GetRoles), roleHandler.Roles)
	grp.GET("/roles/:id", authMiddleware.Authorizer(permission.GetRole), roleHandler.Role)
	grp.PUT("/roles/:id", authMiddleware.Authorizer(permission.UpdateRole), roleHandler.UpdateRole)
	grp.DELETE("/roles/:id", authMiddleware.Authorizer(permission.DeleteRole), roleHandler.DeleteRole)
	grp.POST("/roles/:user-id/assignrole/:role-id", roleHandler.AssignRole)
	grp.GET("/roles/userrole/:user-id", roleHandler.GetUserRoles)
}
