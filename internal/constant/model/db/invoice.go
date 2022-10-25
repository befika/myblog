package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Invoice struct {
	ID uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`

	Month       time.Time `json:"due_date"`
	BillTo      uuid.UUID `json:"bill_to"`
	PaidTo      uuid.UUID `json:"paid_to"`
	TotalBlogs  int       `json:"total_blogs"`
	Status      string    `json:"status"`
	Total       float64   `json:"total"`
	Description string    `json:"description"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
