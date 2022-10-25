package invoice

import (
	"blog/internal/constant/model/rest"
	"blog/internal/module/invoice"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gravitational/trace"
)

type InvoiceHandler interface {
	Invoices(c *gin.Context)
}
type invoiceHandler struct {
	invoiceUsecase invoice.Usecase
}

func NewInvoiceHandler(invoiceUsecase invoice.Usecase) InvoiceHandler {
	return &invoiceHandler{
		invoiceUsecase: invoiceUsecase,
	}
}

func (i invoiceHandler) Invoices(c *gin.Context) {
	ctx := c.Request.Context()
	pgnFlt, err := rest.ParsePgn(c)
	if err != nil {
		rest.ErrorResponseJson(c, err, trace.ErrorToCode(err))
		return
	}
	detail, err := i.invoiceUsecase.Invoices(ctx, pgnFlt)
	if err != nil {
		rest.ErrorResponseJson(c, err, trace.ErrorToCode(err))
		return
	}
	rest.SuccessResponseJson(c, pgnFlt, detail, http.StatusOK)
}
