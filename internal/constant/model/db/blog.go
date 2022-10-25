package model

import (
	"mime/multipart"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Blog struct {
	ID          uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID      uuid.UUID
	Title       string `json:"title" form:"title" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	Content     string `json:"content" form:"content" validate:"required"`
	Status      string `json:"status" form:"status" validate:"required"`

	Pictures         []BlogImages
	BlogSubscription *[]BlogSubscription

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	DeletedBy uuid.UUID      `json:"-"`
}
type BlogImages struct {
	ID          uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	BlogID      uuid.UUID `json:"blog_id"`
	Picture     string
	File        *multipart.FileHeader `gorm:"-" form:"picture"`
	Description string                `json:"description"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type BlogSubscription struct {
	ID     uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	BlogID uuid.UUID `json:"blog_id"`
	UserID uuid.UUID `json:"user_id"`
	Status string    `json:"status"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
