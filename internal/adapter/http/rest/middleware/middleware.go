package middleware

import (
	model "blog/internal/constant/model/db"
	utils "blog/internal/constant/model/init"
	"blog/internal/constant/model/permission"
	"blog/internal/constant/model/rest"
	"blog/internal/module/auth"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gravitational/trace"
)

type AuthMiddleware interface {
	Authorizer(perm string) gin.HandlerFunc
	ExtractToken(r *http.Request) string
}

type authMiddleWare struct {
	authUseCase auth.LoginUseCase
	utils       utils.Utils
}

func NewAuthMiddleware(authUseCase auth.LoginUseCase, utils utils.Utils) AuthMiddleware {
	return &authMiddleWare{
		authUseCase: authUseCase,
		utils:       utils,
	}
}

//Authorizer is a middleware for authorization
func (n *authMiddleWare) Authorizer(perm string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := []model.Role{}
		token := n.ExtractToken(c.Request)
		claims, er := n.authUseCase.GetClaims(token)
		if er != nil {
			log.Println(er)
			rest.ErrorResponseJson(c, er.Error(), trace.ErrorToCode(er))
			return
		}
		if claims != nil {
			role = claims.Roles
			c.Set("x-user-id", claims.Subject)
			c.Set("x-user-role", role)
		}
		err := n.utils.Enforcer.LoadPolicy()
		if err != nil {
			log.Println(er)
			rest.ErrorResponseJson(c, er.Error(), trace.ErrorToCode(er))
			return
		}
		var Authorized bool
		for _, roles := range role {
			id := roles.ID.String()
			Authorized, _ = n.utils.Enforcer.Enforce(id, permission.PermissionObjects[perm], permission.PermissionActions[perm])
			if Authorized {
				c.Next()
			}
		}
		if !Authorized {
			log.Println("not Authorized")
			rest.ErrorResponseJson(c, errors.New("you are not authorized"), 400)
			return
		}

	}

}
func (n *authMiddleWare) ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("auth")
	if token != "" {
		return token
	}
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}
