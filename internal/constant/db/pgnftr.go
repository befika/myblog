package dbpgn

import (
	"log"
	"strings"

	rest "blog/internal/constant/model/rest"

	"gorm.io/gorm"
)

func Filter(f *rest.FilterParams) func(db *gorm.DB) *gorm.DB {
	log.Printf("FILTER: %#v", f)
	return func(db *gorm.DB) *gorm.DB {
		var offset int64 = 0
		for _, fil := range f.Filter {
			if fil.Key != "" && fil.Values != nil {
				if fil.Key == "q" {
					db = db.Where(`document @@ to_tsquery(?)`, fil.Values[0]+":*")
				} else {
					db = db.Where(`"`+fil.Key+`" IN ?`, fil.Values)
				}
			}
		}

		if f == nil {
			f = &rest.FilterParams{}
		}
		if f.Sort == nil {
			f.Sort = append(f.Sort, "created_at")
		}
		if f.Page > 0 {
			offset = (f.Page - 1) * f.PerPage
			db.Offset(int(offset))
			if f.PerPage <= 0 {
				f.PerPage = 10
			}
			db = db.Limit(int(f.PerPage))
		}
		if f.Sort != nil {
			if f.Sort[0] != "" {
				sort := "ASC"
				if len(f.Sort) > 1 {
					if strings.EqualFold(f.Sort[1], "ASC") || strings.EqualFold(f.Sort[1], "DSC") {
						sort = f.Sort[1]
					}
				} else {
					if len(f.Sort) < 2 {
						f.Sort = append(f.Sort, sort)
					}
				}

				db.Order(strings.Join(f.Sort, " "))
			}
		}

		return db
	}
}
