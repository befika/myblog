package login_test

import (
	"blog/initiator/helper"
	authHandler "blog/internal/adapter/http/rest/server/auth"
	rolepersistence "blog/internal/adapter/storage/role"
	userpersistence "blog/internal/adapter/storage/user"
	"blog/internal/constant"
	authModule "blog/internal/module/auth"
	roleModule "blog/internal/module/role"
	"net/http/httptest"

	"log"
	"os"

	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	authModel = "../../config/rbac_model.conf"
)

var (
	jwtManager  authModule.JWTManager
	AuthModule  authModule.LoginUseCase
	AuthHandler authHandler.AuthHandler
	Resp        *httptest.ResponseRecorder
	ReqJson     string
)

func init() {
	DATABASE_URL, err := constant.DbConnectionString()
	if err != nil {
		log.Fatal("database connection failed!")
	}
	dbConn, er := helper.GetConn(DATABASE_URL, authModel)
	if err != nil {
		log.Fatal(er)
	}
	if err != nil {
		os.Exit(1)
	}
	userPersistence := userpersistence.UserInit(dbConn.Conn)
	rolePersistence := rolepersistence.RoleInit(dbConn.Conn)
	roleUseCase := roleModule.InitializeRole(rolePersistence, dbConn)
	authUseCase := authModule.Initialize(userPersistence, jwtManager, dbConn)
	authHandler := authHandler.NewAuthHandler(authUseCase, roleUseCase)
	AuthHandler = authHandler

}
func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ginkgo Suite")
}
