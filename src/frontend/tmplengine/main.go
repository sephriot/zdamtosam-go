package tmplengine

import (
	"html/template"
	"math"
	"net/http"
	"strconv"
)

func Render(w http.ResponseWriter, data interface{}, templates ...string) {
	tmpl := template.Must(template.New("").Funcs(template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
		"sub": func(a, b int) int {
			return a - b
		},
		"mod": func(a, b int) int {
			return a % b
		},
		"mul": func(a, b int) int {
			return a * b
		},
		"increment": func(a string) string {
			v, _ := strconv.ParseInt(a, 10, 64)
			return strconv.Itoa(int(v + 1))
		},
		"decrement": func(a string) string {
			v, _ := strconv.ParseInt(a, 10, 64)
			return strconv.Itoa(int(v - 1))
		},
		"inc": func(a int) int {
			return a + 1
		},
		"dec": func(a int) int {
			return a - 1
		},
		"toInt": func(a string) int {
			v, _ := strconv.ParseInt(a, 10, 64)
			return int(v)
		},
		"divCeil": func(a, b int) int {
			return int(math.Ceil(float64(a) / float64(b)))
		},
		"toURL": func(v string) template.URL {
			return template.URL(v)
		},
	}).ParseFiles(append(templates, "templates/base.html")...))
	tmpl.ExecuteTemplate(w, "base", data)
}
