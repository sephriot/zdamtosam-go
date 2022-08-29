package tmplengine

import (
	"html/template"
	"net/http"
)

func Render(w http.ResponseWriter, data interface{}, templates ...string) {
	tmpl := template.Must(template.ParseFiles(append(templates, "templates/base.html")...))
	tmpl.ExecuteTemplate(w, "base", data)
}
