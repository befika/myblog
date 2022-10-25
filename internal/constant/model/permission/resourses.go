package permission

import "fmt"

var (
	User         = "user"
	Auth         = "auth"
	Role         = "role"
	Permission   = "permission"
	Blog         = "blog"
	Subscription = "subscription"
)

func Action(resourse string, operation string) string {
	return fmt.Sprintf("%s_%s", resourse, operation)
}

func Object(prmObj string) string {
	return prmObj
}
