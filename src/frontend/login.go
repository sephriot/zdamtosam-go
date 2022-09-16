package frontend

import (
	"net/http"
	"zdamtosam.pl/src/frontend/tmplengine"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	data := h.prepareTemplateData(r)
	data["PageTitle"] = "ZdamToSam | Logowanie"
	tmplengine.Render(w, data, tmplengine.FS_PATH_PREFIX+"templates/login.html",
		tmplengine.FS_PATH_PREFIX+"templates/navbar.html")
}
