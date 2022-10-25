package invoice

import (
	"blog/internal/adapter/storage/invoice"
	"blog/internal/adapter/storage/subscription"
	"context"
	"errors"
	"time"

	model "blog/internal/constant/model/db"
	utils "blog/internal/constant/model/init"
	"blog/internal/constant/model/rest"

	"github.com/casbin/casbin/v2"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/robfig/cron.v2"
)

type Usecase interface {
	SendInvoice()
	MonthlyJob()
	Invoices(ctx context.Context, pgnFlt *rest.FilterParams) ([]model.Invoice, error)
	CheckInvoice(ctx context.Context, prm *model.Invoice) (*model.Invoice, error)
}
type service struct {
	contextTimeout     time.Duration
	enforcer           *casbin.Enforcer
	validator          *validator.Validate
	trans              ut.Translator
	cron               *cron.Cron
	invoicePercistance invoice.Invoice
	subPercistance     subscription.BlogSubscription
}

func InvoiceInit(cron *cron.Cron, invoicePercistance invoice.Invoice, subPercistance subscription.BlogSubscription, utils utils.Utils) Usecase {
	return &service{
		contextTimeout:     utils.Timeout,
		enforcer:           utils.Enforcer,
		validator:          utils.GoValidator,
		trans:              utils.Translator,
		cron:               cron,
		invoicePercistance: invoicePercistance,
		subPercistance:     subPercistance,
	}
}

func (s service) SendInvoice() {
	// s.cron.AddFunc("@every 1m", s.MonthlyJob)
	s.cron.AddFunc("0 7 1 * *", s.MonthlyJob)
	s.cron.Start()
}

func (s service) MonthlyJob() {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	Month := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	blogSub, err := s.subPercistance.BlogSubscriptions()
	if err != nil {
		return
	}
	prm := []model.Invoice{}
	count := 1
	for i, b := range blogSub {

		for j := 0; j < len(blogSub); j = j + 1 {
			if b.UserID == blogSub[j].UserID {
				count = count + 1
			}
		}

		if i == 0 {
			// log.Println("subscription at:", i)
			prm = append(prm, model.Invoice{
				Month:       Month,
				BillTo:      b.UserID,
				PaidTo:      b.BlogID,
				TotalBlogs:  count,
				Total:       float64(count) * 20,
				Status:      "Pending",
				Description: "monthly subscription payment for blogs",
			})
		}

		for _, in := range prm {
			for k := 0; k < len(prm); k = k + 1 {
				if in.BillTo != prm[k].BillTo {

					prm = append(prm, model.Invoice{
						Month:       Month,
						BillTo:      b.UserID,
						PaidTo:      b.BlogID,
						TotalBlogs:  count,
						Total:       float64(count) * 20,
						Description: "monthly subscription payment for blogs",
					})
				}
			}
		}
	}

	for _, in := range prm {
		_, err := s.invoicePercistance.SendInvoice(&in)
		if err != nil {
			return
		}

	}
}

func (s service) Invoices(c context.Context, pgnFlt *rest.FilterParams) ([]model.Invoice, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.invoicePercistance.Invoices(ctx, pgnFlt)
}
func (s service) CheckInvoice(c context.Context, prm *model.Invoice) (*model.Invoice, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	if uuid.Equal(prm.ID, uuid.Nil) || prm.Status == "" {
		return nil, errors.New("empty invoice id or status provided")
	}
	return s.invoicePercistance.CheckInvoice(ctx, prm)
}
