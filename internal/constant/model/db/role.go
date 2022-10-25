package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Role struct {
	ID   uuid.UUID `json:"id" gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	Name string    `json:"name" form:"name" validate:"required"`

	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type UserRoles struct {
	UserID uuid.UUID `json:"user_id" form:"user_id" gorm:"primarykey;type:uuid;"`
	RoleID uuid.UUID `json:"role_id" form:"role_id" gorm:"primarykey;type:uuid;"`

	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
