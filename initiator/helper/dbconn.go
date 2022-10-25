package helper

import (
	model "blog/internal/constant/model/db"
	utils "blog/internal/constant/model/init"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConn(dbUrl string, authModel string) (utils.Utils, error) {

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{
		SkipDefaultTransaction: true, //30% performance increases
	})
	if err != nil {
		log.Printf("Error when Opening database connection: %v", err)
		os.Exit(1)
	}
	db.AutoMigrate(
		&model.Role{},
		&model.User{},
		&model.UserRoles{},
		&model.Blog{},
		&model.BlogImages{},
		&model.SystemSubscription{},
		&model.BlogSubscription{},
		&model.Invoice{},
	)
	trans, validate, err := GetValidation()
	if err != nil {
		log.Fatal("*errors.ErrorModel ", err)
	}
	duration, _ := strconv.Atoi(os.Getenv("timeout"))
	timeoutContext := time.Duration(duration) * time.Second
	enforcer := NewEnforcer(db, authModel)
	return utils.Utils{
		Conn:        db,
		Timeout:     timeoutContext,
		Translator:  trans,
		GoValidator: validate,
		Enforcer:    enforcer,
	}, nil
}

// NewEnforcer creates an enforcer via file or DB.
func NewEnforcer(conn *gorm.DB, authModel string) *casbin.Enforcer {
	adapter, err := gormadapter.NewAdapterByDBWithCustomTable(conn, &model.CasbinRule{})
	if err != nil {
		log.Fatal("*errors.ErrorModel ", err)
	}

	enforcer, err := casbin.NewEnforcer(authModel, adapter)
	if err != nil {
		log.Fatal("*errors.ErrorModel ", err)
	}
	enforcer.EnableAutoSave(true)
	enforcer.LoadPolicy()
	return enforcer
}
