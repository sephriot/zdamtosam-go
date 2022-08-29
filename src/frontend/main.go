package frontend

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	StaticResources []string
}

func (*Handler) Handle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, ok := vars["section"]
	if !ok {
		Index(w, r)
		return
	}
}
