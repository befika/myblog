package helper

import (
	"strconv"

	utils "blog/internal/constant/model/init"

	"github.com/gin-gonic/gin"
)

func GeneratePaginationRequest(context *gin.Context) *utils.PageParam {
	// default limit, page & sort parameter
	limit := 2
	page := 1
	sort := "created_at desc"
	var filter []string
	query := context.Request.URL.Query()

	for key, value := range query {
		queryValue := value[len(value)-1]

		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = queryValue
			break
		case "filter":
			filter = append(filter, queryValue)
			break
		}

	}

	return &utils.PageParam{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}
}
