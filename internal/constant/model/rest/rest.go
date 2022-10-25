package rest

import (
	"github.com/gin-gonic/gin"
)

type Links struct {
	Self     string `json:"self"`
	First    string `json:"first"`
	Previous string `json:"previous"`
	Next     string `json:"next"`
	Last     string `json:"last"`
}
type Pagination struct {
	Page       uint   `form:"page" json:"page"`
	Limit      int    `form:"limit" json:"limit"`
	Sort       string `form:"sort" json:"sort"`
	OrderBy    string `form:"order_by" json:"order_by"`
	FilterBy   string `form:"filter_by" json:"filter_by"`
	Filterkey  string `form:"filter_key" json:"filter_key"`
	TotalCount uint   `json:"total_count"`
	Links      Links  `json:"links,omitempty"`
}

func ParsePgn(c *gin.Context) (*FilterParams, error) {
	filterParam := QueryParams{}
	err := c.BindQuery(&filterParam)
	if err != nil {
		return nil, err
	}
	param, err := filterParam.Get()
	if err != nil {
		return nil, err
	}
	return param, nil
}

type Response struct {
	MetaData interface{} `json:"meta_data,omitempty"`
	Data     interface{} `json:"data"`
}

type ErrData struct {
	Error interface{} `json:"error"`
}

//ResponseJson creates new json object
func ErrorResponseJson(c *gin.Context, errData interface{}, statusCode int) {
	c.AbortWithStatusJSON(statusCode, ErrData{errData})
	return
}

//ResponseJson creates new json object
func SuccessResponseJson(c *gin.Context, metaData interface{}, responseData interface{}, statusCode int) {
	c.JSON(statusCode, Response{
		MetaData: metaData,
		Data:     responseData,
	})
	return
}
