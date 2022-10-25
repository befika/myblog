package model

import (
	"mime/multipart"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID `json:"id" gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	FirstName   string    `json:"first_name" form:"first_name" validate:"required,min=3,max=20,alpha"`
	MiddleName  string    `json:"middle_name" form:"middle_name" validate:"required,min=3,max=20,alpha"`
	LastName    string    `json:"last_name" form:"last_name" validate:"required,min=3,max=20,alpha"`
	Photo       string    `validate:"required"`
	Email       string    `json:"email" form:"email" validate:"required,email,unique"`
	PhoneNumber string    `json:"phone_number" form:"contact" validate:"required,gte=10;unique"`
	Username    string    `json:"username" form:"username" validate:"required"`
	Password    string    `json:"password" form:"password" validate:"required"`
	Status      string    `json:"status" form:"status"  validate:"required"`

	PhotoFile          *multipart.FileHeader `gorm:"-" form:"photo"`
	Roles              *[]Role               `gorm:"many2many:user_roles"`
	Blogs              *[]Blog
	SystemSubscription *SystemSubscription
	BlogSubscription   *[]BlogSubscription

	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	DeletedAt *gorm.DeletedAt `json:"-" gorm:"index"`
}
type SystemSubscription struct {
	ID     uuid.UUID `json:"id" gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	UserID uuid.UUID `json:"user_id"`
	Price  float64   `json:"price" validate:"required"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
