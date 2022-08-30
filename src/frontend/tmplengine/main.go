package tmplengine

import (
	"html/template"
	"net/http"
)

func Render(w http.ResponseWriter, data interface{}, templates ...string) {
	tmpl := template.Must(template.New("").Funcs(template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
		"mod": func(a, b int) int {
			return a % b
		},
	}).ParseFiles(append(templates, "templates/base.html")...))
	tmpl.ExecuteTemplate(w, "base", data)
}
