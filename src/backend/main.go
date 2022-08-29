package backend

import (
	"database/sql"
	"net/http"
	"zdamtosam/src/backend/db"
)

type Handler struct {
	db *sql.DB
}

func NewHandler() *Handler {
	return &Handler{db.NewDatabaseClient()}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	rows, err := h.db.Query("SELECT * FROM users")
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var data1 []byte
		var data2 []byte
		var data3 []byte
		var data4 []byte
		var data5 []byte
		rows.Scan(&data1, &data2, &data3, &data4, &data5)
		w.Write(data1)
		w.Write(data2)
		w.Write(data3)
		w.Write(data4)
		w.Write(data5)
	}

}
