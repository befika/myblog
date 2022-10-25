package init

import (
	"time"

	"github.com/casbin/casbin/v2"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	"gorm.io/gorm"
)

type Utils struct {
	Conn        *gorm.DB
	GoValidator *validator.Validate
	Translator  ut.Translator
	Timeout     time.Duration
	Enforcer    *casbin.Enforcer
}
type Filter []string
type PageParam struct {
	Limit  int    `json:"limit"`
	Page   int    `json:"page"`
	Sort   string `json:"sort"`
	Filter Filter `json:"filter"`
}
