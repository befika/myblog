package domain

import (
	userHandler "blog/internal/adapter/http/rest/server/user"
	"blog/internal/adapter/storage/user"
	utils "blog/internal/constant/model/init"
	routing "blog/internal/glue"
	userUsecase "blog/internal/module/user"

	"github.com/gin-gonic/gin"
)

func UserInit(utils utils.Utils, router *gin.RouterGroup) {
	userPercistance := user.UserInit(utils.Conn)
	userUsecase := userUsecase.InitializeUser(userPercistance, utils)
	userHandler := userHandler.NewUserHandler(userUsecase)
	routing.UserRoutes(router, AuthMiddleware(utils), userHandler)
}
