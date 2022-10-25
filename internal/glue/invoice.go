package routing

import (
	"blog/internal/adapter/http/rest/server/invoice"

	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(grp *gin.RouterGroup, invoiceHandler invoice.InvoiceHandler) {
	grp.GET("/invoices", invoiceHandler.Invoices)
}
