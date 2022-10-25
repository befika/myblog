package domain

import (
	blogHandler "blog/internal/adapter/http/rest/server/blog"
	Percistance "blog/internal/adapter/storage/blog"
	"blog/internal/adapter/storage/user"
	utils "blog/internal/constant/model/init"
	routing "blog/internal/glue"
	"blog/internal/module/auth"
	"blog/internal/module/blog"

	"github.com/gin-gonic/gin"
)

func BlogInit(utils utils.Utils, router *gin.RouterGroup) {
	blogPercistance := Percistance.BlogInit(utils.Conn)
	blogUsecase := blog.BlogInit(blogPercistance, utils)
	usrPersistence := user.UserInit(utils.Conn)
	jwtManager := auth.NewJWTManager("secret")
	loginUseCase := auth.Initialize(usrPersistence, *jwtManager, utils)

	blogImgPercistance := Percistance.BlogImgInit(utils.Conn)
	blogImgUsecase := blog.BlogImgInit(blogImgPercistance, utils)
	blogImgHandler := blogHandler.NewBlogImgHandler(blogImgUsecase)

	blogHandler := blogHandler.NewBlogHandler(blogUsecase, AuthMiddleware(utils), loginUseCase)
	routing.BlogRoutes(router, AuthMiddleware(utils), blogImgHandler, blogHandler)
}
