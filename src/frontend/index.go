package frontend

import (
	"net/http"
	"zdamtosam/src/frontend/tmplengine"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmplengine.Render(w, nil, "templates/index.html", "templates/navbar.html")
}
