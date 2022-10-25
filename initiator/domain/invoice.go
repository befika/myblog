package domain

import (
	InvoiceHandler "blog/internal/adapter/http/rest/server/invoice"
	"blog/internal/adapter/storage/invoice"
	"blog/internal/adapter/storage/subscription"
	utils "blog/internal/constant/model/init"
	routing "blog/internal/glue"
	InvoiceUsecase "blog/internal/module/invoice"

	"github.com/gin-gonic/gin"
	"gopkg.in/robfig/cron.v2"
)

func InvoiceInit(utils utils.Utils, router *gin.RouterGroup) {
	subPercistance := subscription.BlogSubInit(utils.Conn)
	invoicePercistance := invoice.InvoiceInit(utils.Conn)
	cron := cron.New()
	invoiceUsecase := InvoiceUsecase.InvoiceInit(cron, invoicePercistance, subPercistance, utils)
	invoiceUsecase.SendInvoice()
	invoiceHandler := InvoiceHandler.NewInvoiceHandler(invoiceUsecase)
	routing.InvoiceRoutes(router, invoiceHandler)

}
