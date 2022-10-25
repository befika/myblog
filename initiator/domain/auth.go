package domain

import (
	"blog/internal/adapter/http/rest/middleware"
	authHandler "blog/internal/adapter/http/rest/server/auth"
	"blog/internal/adapter/storage/role"
	"blog/internal/adapter/storage/user"
	utils "blog/internal/constant/model/init"
	routing "blog/internal/glue"
	"blog/internal/module/auth"
	roleUsecase "blog/internal/module/role"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(utils utils.Utils) middleware.AuthMiddleware {
	usrPersistence := user.UserInit(utils.Conn)
	jwtManager := auth.NewJWTManager("secret")
	loginUseCase := auth.Initialize(usrPersistence, *jwtManager, utils)

	authMiddleWare := middleware.NewAuthMiddleware(loginUseCase, utils)
	return authMiddleWare
}

func AuthInit(utils utils.Utils, router *gin.RouterGroup) {
	rolePercistance := role.RoleInit(utils.Conn)
	usrPersistence := user.UserInit(utils.Conn)
	jwtManager := auth.NewJWTManager("secret")

	authUsecase := auth.Initialize(usrPersistence, *jwtManager, utils)
	roleUsecase := roleUsecase.InitializeRole(rolePercistance, utils)
	authhandler := authHandler.NewAuthHandler(authUsecase, roleUsecase)

	permissionUsecase := auth.PermissionInit(utils)
	permissionUsecase.MigratePermissionsToCasbin()
	permissionHandler := authHandler.NewPermissionHandler(permissionUsecase, utils)

	routing.AuthRoutes(router, AuthMiddleware(utils), permissionHandler, authhandler)
}
