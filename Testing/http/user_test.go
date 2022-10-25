package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"time"

	userHandler "blog/internal/adapter/http/rest/server/user"
	persistence "blog/internal/adapter/storage/user"
	"blog/internal/constant"
	model "blog/internal/constant/model/db"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"

	utils "blog/internal/constant/model/init"
	userModule "blog/internal/module/user"

	"github.com/cucumber/godog"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type userTest struct {
	userHandler userHandler.UserHandler
	resp        *httptest.ResponseRecorder
	token       string
	db          *gorm.DB
	reqJson     string
	utils       utils.Utils
}

func (a *userTest) resetResponse(*godog.Scenario) {
	a.resp = httptest.NewRecorder()
	DATABASE_URL, err := constant.DbConnectionString()
	if err != nil {
		log.Fatal("database connection failed!")
	}
	dbConn, err := gorm.Open(postgres.Open(DATABASE_URL), &gorm.Config{
		SkipDefaultTransaction: true, //30% performance increases
	})
	if err != nil {
		os.Exit(1)
	}
	userPersistence := persistence.UserInit(dbConn)
	userUseCase := userModule.InitializeUser(userPersistence, a.utils)
	userHandler := userHandler.NewUserHandler(userUseCase)
	a.userHandler = userHandler
	a.userHandler = userHandler
	a.db = dbConn
}

func (a *userTest) iSendHTTPRequestTo(method, endpoint string) error {
	var err error

	if err != nil {
		return err
	}
	// handle panic
	defer func() {
		switch t := recover().(type) {
		case string:
			err = fmt.Errorf(t)
		case error:
			err = t
		}
	}()

	switch endpoint {
	case "/v1/users":
		if method == "GET" {
			req, err := http.NewRequest(method, endpoint, nil)

			if err != nil {
				return err
			}
			c, _ := gin.CreateTestContext(a.resp)
			c.Request = req
			a.userHandler.Users(c)

		} else if method == "POST" {
			req, err := http.NewRequest(method, endpoint, bytes.NewBuffer([]byte(a.reqJson)))
			req.Header.Set("Content-Type", "application/json")

			if err != nil {
				return err
			}
			c, _ := gin.CreateTestContext(a.resp)
			c.Request = req
			a.userHandler.AddUser(c)
		}

	case "/v1/users/cbf77990-a47c-4022-85d6-99cc76bd705a":

		if method == "PUT" {

			req, err := http.NewRequest(method, endpoint, bytes.NewBuffer([]byte(a.reqJson)))
			req.Header.Set("Content-Type", "application/json")
			if err != nil {
				return err
			}
			c, _ := gin.CreateTestContext(a.resp)
			c.Params = []gin.Param{
				{
					Key:   "id",
					Value: "cbf77990-a47c-4022-85d6-99cc76bd705a",
				},
			}
			c.Request = req
			a.userHandler.UpdateUser(c)
		} else if method == "DELETE" {

			req, err := http.NewRequest(method, endpoint, bytes.NewBuffer([]byte(a.reqJson)))
			req.Header.Set("Content-Type", "application/json")

			if err != nil {
				return err
			}
			c, _ := gin.CreateTestContext(a.resp)
			c.Params = []gin.Param{
				{
					Key:   "id",
					Value: "cbf77990-a47c-4022-85d6-99cc76bd705a",
				},
			}
			c.Request = req
			a.userHandler.DeleteUser(c)
		}

	default:
		err = fmt.Errorf("unknown endpoint: %s", endpoint)
	}
	return nil

}

func (a *userTest) iHaveATokenValue(token string) error {
	a.token = fmt.Sprintf("Bearer %s", token)
	return nil
}

func (a *userTest) theResponseCodeShouldBe(code int) error {
	if code != a.resp.Code {
		if a.resp.Code >= 400 {
			return fmt.Errorf("expected response code to be: %d, but actual is: %d, response message: %s", code, a.resp.Code, a.resp.Body.String())
		}
		return fmt.Errorf("expected response code to be: %d, but actual is: %d", code, a.resp.Code)
	}
	return nil
}

func (a *userTest) theResponseShouldMatchJson(body *godog.DocString) error {
	var expected, actual interface{}

	// re-encode expected response
	if err := json.Unmarshal([]byte(body.Content), &expected); err != nil {

		return err
	}

	// re-encode actual response too
	if err := json.Unmarshal(a.resp.Body.Bytes(), &actual); err != nil {
		return err
	}

	// the matching may be adapted per different requirements.
	if !reflect.DeepEqual(expected, actual) {
		return fmt.Errorf("expected JSON does not match actual, %v vs. %v", expected, actual)
	}
	return nil

}

func (a *userTest) thereAreUsers(users *godog.Table) error {
	head := users.Rows[0].Cells
	for i := 1; i < len(users.Rows); i++ {
		var vals []interface{}
		for n, cell := range users.Rows[i].Cells {
			switch head[n].Value {
			case "id":
				vals = append(vals, model.User{ID: uuid.FromStringOrNil(cell.Value)})
			case "username":
				vals = append(vals, model.User{Username: cell.Value})
			case "password":
				vals = append(vals, model.User{Password: cell.Value})
			case "phone":
				vals = append(vals, model.User{PhoneNumber: cell.Value})
			case "first_name":
				vals = append(vals, model.User{FirstName: cell.Value})
			case "middle_name":
				vals = append(vals, model.User{MiddleName: cell.Value})
			case "last_name":
				vals = append(vals, model.User{LastName: cell.Value})
			case "email":
				vals = append(vals, model.User{Email: cell.Value})
			case "created_at":
				vals = append(vals, model.User{CreatedAt: time.Time{}})
			case "updated_at":
				vals = append(vals, model.User{UpdatedAt: time.Time{}})
			default:
				return fmt.Errorf("unexpected column name: %s", head[0].Value)
			}
		}
		bytes := []byte(vals[0].(model.User).ID.String())
		a.reqJson = string(bytes)
	}

	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	user := &userTest{}
	ctx.BeforeScenario(user.resetResponse)
	ctx.Step(`^I send "(GET|POST|PUT|DELETE)" HTTP request to "([^"]*)"$`, user.iSendHTTPRequestTo)
	ctx.Step(`^I have a token value "([^"]*)"$`, user.iHaveATokenValue)
	ctx.Step(`^the response code should be (\d+)$`, user.theResponseCodeShouldBe)
	ctx.Step(`^the response should match json:$`, user.theResponseShouldMatchJson)
	ctx.Step(`^there are users:$`, user.thereAreUsers)
	ctx.AfterScenario(func(sc *godog.Scenario, err error) {
		user.reqJson = ""
	})
}
