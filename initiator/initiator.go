package initiator

import (
	"blog/initiator/helper"
	"blog/internal/constant"
	"log"

	"blog/initiator/domain"

	"github.com/gin-gonic/gin"
)

const (
	authModel = "/config/rbac_model.conf"
)

func Initialize() {
	DATABASE_URL, err := constant.DbConnectionString()
	if err != nil {
		log.Fatal("database connection failed!")
	}
	common, er := helper.GetConn(DATABASE_URL, authModel)
	if err != nil {
		log.Fatal(er)
	}
	router := gin.Default()
	// router.Use(corsMW())

	v1 := router.Group("/v1")

	// initialize domains
	domain.AuthInit(common, v1)
	domain.RoleInit(common, v1)
	domain.UserInit(common, v1)
	domain.BlogInit(common, v1)
	domain.SubscriptionInit(common, v1)
	domain.InvoiceInit(common, v1)
	router.Run(":9090")

}
