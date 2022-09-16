package frontend

import (
	"database/sql"
	"strings"
	zdb "zdamtosam.pl/src/db"
)

type Breadcrumb struct {
	Name string
	Path string
}

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

	if path == "/login" {
		breadcrumbs = append(breadcrumbs, Breadcrumb{
			Name: "Logowanie",
			Path: "/login",
		})
	}

	if path == "/terms-of-service" {
		breadcrumbs = append(breadcrumbs, Breadcrumb{
			Name: "Warunki korzystania z serwisu",
			Path: "/terms-of-service",
		})
	}

	if path == "/privacy-policy" {
		breadcrumbs = append(breadcrumbs, Breadcrumb{
			Name: "Polityka prywatności",
			Path: "/privacy-policy",
		})
	}

	if pathParams["level"] != "" {
		breadcrumbs = append(breadcrumbs, Breadcrumb{
			Name: zdb.GetLevelNameById(db, pathParams["level"]),
			Path: "/level/" + pathParams["level"],
		})
	}
	subcategoryPath := ""
	if pathParams["category"] != "" {
		index := strings.Index(path, "/category/"+pathParams["category"])
		basePath := ""
		if index != -1 {
			basePath = path[0:index]
		}
		categoryPath := basePath + "/category/" + pathParams["category"]
		breadcrumbs = append(breadcrumbs, Breadcrumb{
			Name: zdb.GetCategoryNameById(db, pathParams["category"]),
			Path: categoryPath,
		})

		if pathParams["subcategory"] != "" {
			subcategoryPath = categoryPath + "/subcategory/" + pathParams["subcategory"]
			breadcrumbs = append(breadcrumbs, Breadcrumb{
				Name: zdb.GetSubcategoryNameById(db, pathParams["subcategory"]),
				Path: subcategoryPath,
			})
		}
	}
	if pathParams["exercise"] != "" {
		breadcrumbs = append(breadcrumbs, Breadcrumb{
			Name: "Zadanie " + pathParams["exercise"],
			Path: subcategoryPath + "/exercise/" + pathParams["exercise"],
		})
	}

	return breadcrumbs
}
