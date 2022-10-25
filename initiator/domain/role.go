package domain

import (
	roleHandler "blog/internal/adapter/http/rest/server/role"
	"blog/internal/adapter/storage/role"
	utils "blog/internal/constant/model/init"
	routing "blog/internal/glue"
	roleUsecase "blog/internal/module/role"

	"github.com/gin-gonic/gin"
)

func RoleInit(utils utils.Utils, router *gin.RouterGroup) {
	rolePercistance := role.RoleInit(utils.Conn)
	roleUsecase := roleUsecase.InitializeRole(rolePercistance, utils)
	roleHandler := roleHandler.NewRoleHandler(roleUsecase)
	routing.RoleRoutes(router, AuthMiddleware(utils), roleHandler)
}
