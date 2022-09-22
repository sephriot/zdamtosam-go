package frontend

import (
	"net/http"
	"zdamtosam.pl/src/db"
	"zdamtosam.pl/src/frontend/tmplengine"
	"zdamtosam.pl/src/model"
)

func (h *Handler) Profile(w http.ResponseWriter, r *http.Request) {
	data := h.prepareTemplateData(r)
	data["PageTitle"] = "ZdamToSam | Profil użytkownika"
	data["User7dStats"] = db.GetUserStatsForLast7Days(h.Db, data["LoggedUser"].(model.User).Id)

	tmplengine.Render(w, data,
		tmplengine.FS_PATH_PREFIX+"templates/profile.html",
		tmplengine.FS_PATH_PREFIX+"templates/categories.html",
		tmplengine.FS_PATH_PREFIX+"templates/navbar.html")
}
