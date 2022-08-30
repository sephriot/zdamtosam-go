package frontend

import (
	"database/sql"
	"strings"
	zdb "zdamtosam/src/db"
)

// Remember that path has to have odd number of params, following key/value pattern
func getPathParams(path string) map[string]string {
	s := strings.Split(path, "/")
	params := make(map[string]string)
	if len(s) > 2 {
		i := 1
		for i < len(s) {
			params[s[i]] = s[i+1]
			i += 2
		}
	}
	return params
}

func getBreadcrumbs(db *sql.DB, path string) []Breadcrumb {
	pathParams := getPathParams(path)
	breadcrumbs := make([]Breadcrumb, 0)
	breadcrumbs = append(breadcrumbs, Breadcrumb{
		Name: "ZdamToSam",
		Path: "/",
	})

	if path == "/search" {
		breadcrumbs = append(breadcrumbs, Breadcrumb{
			Name: "Wyszukaj",
			Path: "/search",
		})
	}

	if pathParams["level"] != "" {
		breadcrumbs = append(breadcrumbs, Breadcrumb{
			Name: zdb.GetLevelNameById(db, pathParams["level"]),
			Path: "/level/" + pathParams["level"],
		})
	}
	if pathParams["category"] != "" {
		index := strings.Index(path, "/category/"+pathParams["category"])
		basePath := ""
		if index != -1 {
			basePath = path[0:index]
		}
		breadcrumbs = append(breadcrumbs, Breadcrumb{
			Name: zdb.GetCategoryNameById(db, pathParams["category"]),
			Path: basePath + "/category/" + pathParams["category"],
		})
	}

	return breadcrumbs
}
