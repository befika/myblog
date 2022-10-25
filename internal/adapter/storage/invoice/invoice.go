package invoice

import (
	dbpgn "blog/internal/constant/db"
	model "blog/internal/constant/model/db"
	"blog/internal/constant/model/rest"
	"context"

	"gorm.io/gorm"
)

type Invoice interface {
	SendInvoice(prm *model.Invoice) (*model.Invoice, error)
	Invoices(ctx context.Context, pgnFlt *rest.FilterParams) ([]model.Invoice, error)
	CheckInvoice(ctx context.Context, prm *model.Invoice) (*model.Invoice, error)
}
type invoicePercistance struct {
	conn *gorm.DB
}

func InvoiceInit(conn *gorm.DB) Invoice {
	return &invoicePercistance{
		conn: conn,
	}
}

func (i invoicePercistance) SendInvoice(prm *model.Invoice) (*model.Invoice, error) {
	err := i.conn.Model(&model.Invoice{}).Create(&prm).Error
	if err != nil {
		return nil, err
	}
	return prm, nil
}
func (i invoicePercistance) Invoices(ctx context.Context, pgnFlt *rest.FilterParams) ([]model.Invoice, error) {
	conn := i.conn.WithContext(ctx)
	invoices := []model.Invoice{}
	err := conn.Model(&model.Invoice{}).Scopes(dbpgn.Filter(pgnFlt)).Find(&invoices).Error
	if err != nil {
		return nil, err
	}
	if err := conn.Model(&model.Invoice{}).Offset(-1).Limit(-1).Count(&pgnFlt.Total).Error; err != nil {
		return nil, err
	}
	return invoices, nil
}
func (i invoicePercistance) CheckInvoice(ctx context.Context, prm *model.Invoice) (*model.Invoice, error) {
	conn := i.conn.WithContext(ctx)
	err := conn.Model(&model.Invoice{}).Where("id=?", prm.ID).Update("status=?", prm.Status).Error
	if err != nil {
		return nil, err
	}
	return prm, nil
}
