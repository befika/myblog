package login_test

import (
	"bytes"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Login Functionality", func() {
	method := "POST"
	endpoint := "localhost:9090/v1/auth/login"

	Describe("Login request", func() {
		Context("with empty data", func() {
			// user := model.User{Email: "", Password: ""}
			It("should return error", func() {
				req, _ := http.NewRequest(method, endpoint, bytes.NewBuffer([]byte(ReqJson)))

				c, _ := gin.CreateTestContext(Resp)
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

				AuthHandler.Login(c)
				// Resp
				var err error
				if Resp.Code >= 400 {
					err = errors.New("empty data")
				}
				Expect(err).NotTo(BeNil())
				Expect(Resp).To(BeNil())
			})
		})
	})
})
