package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"
)

const DefaultPageSize = 10

type Sort []string

type QueryParams struct {
	Sort    string `json:"sort" form:"sort"`
	Filter  string `json:"filter" form:"filter"`
	Page    string `json:"page" form:"page"`
	PerPage string `json:"per_page" form:"per_page"`
}

type FilterParams struct {
	Sort    Sort     `json:"sort"`
	Filter  []Filter `json:"filter"`
	Page    int64    `json:"page"`
	PerPage int64    `json:"per_page"`

	Total int64 `json:"total"`
}

type filter map[string]json.RawMessage
type Filter struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

func (q QueryParams) Get() (*FilterParams, error) {
	res := &FilterParams{}
	res.Filter = []Filter{}
	if q.Sort != "" {
		decSort, err := url.QueryUnescape(q.Sort)
		if err != nil {
			log.Println("error while unescaping sort query param")
			return nil, err
		}
		q.Sort = decSort

		err = json.Unmarshal([]byte(q.Sort), &res.Sort)
		if err != nil {
			log.Println("error while unmarshalling sort json")
			return nil, err
		}
	}

	if q.Filter != "" {
		decFilter, err := url.QueryUnescape(q.Filter)
		if err != nil {
			log.Println("error while unescaping filter query param")
			return nil, err
		}
		q.Filter = decFilter
		var filter filter

		err = json.Unmarshal([]byte(q.Filter), &filter)
		if err != nil {
			log.Println("error while unmarshalling filter json")
			return nil, err
		}
		keys := getKeys(filter)
		values := map[string][]string{}

		for _, key := range keys {
			var val interface{}
			json.Unmarshal(filter[key], &val)
			switch s := val.(type) {
			case []interface{}:
				for _, v := range s {
					values[key] = append(values[key], fmt.Sprintf("%v", v))
				}
			default:
				values[key] = append(values[key], fmt.Sprintf("%v", s))
			}
		}
		res.Filter = []Filter{}

		if len(keys) > 0 {
			i := 0
			for _, key := range keys {
				f := Filter{Key: key}
				if len(values[key]) > 0 {
					f.Values = values[key]
				}
				res.Filter = append(res.Filter, f)

				i++
			}
		}
	}

	if q.Page == "" {
		res.Page = 1
	} else {

		p, err := toInt64(q.Page)
		if err != nil {
			p = 1
		}

		res.Page = p
	}

	if q.PerPage == "" {
		res.PerPage = DefaultPageSize
	} else {

		pp, err := toInt64(q.PerPage)
		if err != nil {
			pp = DefaultPageSize
		}
		res.PerPage = pp
	}

	return res, nil
}

func getKeys(m filter) []string {
	keys := make([]string, len(m))

	i := 0
	for k := range m {
		keys[i] = k

		i++
	}
	return keys
}

func toInt64(s string) (int64, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func (f *FilterParams) ToQuery() url.Values {
	q := url.Values{}
	if f.Page != 0 {
		q.Add("page", fmt.Sprintf("%d", f.Page))
	}

	if f.PerPage != 0 {
		q.Add("per_page", fmt.Sprintf("%d", f.PerPage))
	}

	if len(f.Sort) > 1 {
		q.Add("sort", fmt.Sprintf(`["%v","%v"]`, f.Sort[0], f.Sort[1]))
	} else if len(f.Sort) == 1 {
		q.Add("sort", fmt.Sprintf(`["%v","ASC"]`, f.Sort[0]))
	}

	if f.Filter[0].Key != "" && len(f.Filter[0].Values) > 0 {
		if len(f.Filter[0].Values) == 1 {
			b, _ := json.Marshal(&f.Filter[0].Values[0])
			q.Add("filter", fmt.Sprintf(`{"%v":%v}`, f.Filter[0].Key, string(b)))
		} else {
			b, _ := json.Marshal(&f.Filter[0].Values)
			q.Add("filter", fmt.Sprintf(`{"%v":%v}`, f.Filter[0].Key, string(b)))
		}
	}
	return q
}
