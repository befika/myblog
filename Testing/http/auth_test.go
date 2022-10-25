package http

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"

	"blog/initiator/helper"
	authHandler "blog/internal/adapter/http/rest/server/auth"
	rolepersistence "blog/internal/adapter/storage/role"
	userpersistence "blog/internal/adapter/storage/user"
	"blog/internal/constant"

	"github.com/gin-gonic/gin"

	model "blog/internal/constant/model/db"
	utils "blog/internal/constant/model/init"
	authModule "blog/internal/module/auth"
	roleModule "blog/internal/module/role"

	"github.com/cucumber/godog"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

const (
	authModel = "../../config/rbac_model.conf"
)

type authTest struct {
	authHandler authHandler.AuthHandler
	resp        *httptest.ResponseRecorder
	db          *gorm.DB
	reqJson     string
	utils       utils.Utils
	jwtManager  authModule.JWTManager
}

func (a *authTest) resetResponse(*godog.Scenario) {
	a.resp = httptest.NewRecorder()
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
	roleUseCase := roleModule.InitializeRole(rolePersistence, a.utils)
	authUseCase := authModule.Initialize(userPersistence, a.jwtManager, a.utils)
	authHandler := authHandler.NewAuthHandler(authUseCase, roleUseCase)
	a.authHandler = authHandler
	a.db = dbConn.Conn
}
func (a *authTest) iSendHTTPRequestToAuth(method, endpoint string) error {
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer([]byte(a.reqJson)))
	if err != nil {
		return err
	}
	c, _ := gin.CreateTestContext(a.resp)
	c.Params = []gin.Param{
		{
			Key:   "email",
			Value: "befika77@gmail.com",
		},
		{
			Key:   "password",
			Value: "123456",
		},
	}
	c.Request = req
	a.authHandler.Login(c)

	return nil

}
func (a *authTest) theAuthResponseCodeShouldBe(code int) error {
	if code != a.resp.Code {
		if a.resp.Code >= 400 {
			return fmt.Errorf("expected response code to be: %d, but actual is: %d, response message: %s", code, a.resp.Code, a.resp.Body.String())
		}
		return fmt.Errorf("expected response code to be: %d, but actual is: %d", code, a.resp.Code)
	}
	return nil
}
func (a *authTest) theUserEntersEmailAndPassword(users *godog.Table) error {
	head := users.Rows[0].Cells
	for i := 0; i < len(users.Rows); i++ {
		var vals []interface{}
		for n, cell := range users.Rows[i].Cells {
			switch head[n].Value {
			case "email":
				vals = append(vals, model.User{Email: cell.Value})
			case "password":
				vals = append(vals, model.User{Password: cell.Value})
			case "status":
				vals = append(vals, model.User{Status: cell.Value})
			default:
				return fmt.Errorf("unexpected column name: %s", head[0].Value)
			}
		}
	}

	return nil
}
func InitializeAuthScenario(ctx *godog.ScenarioContext) {
	auth := &authTest{}
	ctx.BeforeScenario(auth.resetResponse)
	ctx.Step(`^I send "(GET|POST|PUT|DELETE)" HTTP request to Auth "([^"]*)"$`, auth.iSendHTTPRequestToAuth)
	ctx.Step(`^the auth response code should be (\d+)$`, auth.theAuthResponseCodeShouldBe)
	ctx.Step(`^the user enters email and password:$`, auth.theUserEntersEmailAndPassword)
	ctx.AfterScenario(func(sc *godog.Scenario, err error) {
		auth.reqJson = ""
	})
}
