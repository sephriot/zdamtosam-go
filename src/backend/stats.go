package backend

import (
	"encoding/json"
	"log"
	"net/http"
)

type Stats struct {
	Correct       bool  `json:"correct"`
	LevelId       int64 `json:"levelId"`
	Seconds       int64 `json:"seconds"`
	SubcategoryId int64 `json:"subcategoryId"`
}

func (h *Handler) PostStats(w http.ResponseWriter, r *http.Request) {

	user := h.GetLoggedUser(r)
	if user.Id == "" {
		return
	}

	var body Stats
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println(err)
		return
	}

	query := "SELECT id FROM user_stats WHERE user_id = ? AND subcategory_id = ? AND level_id = ? AND date = CURRENT_DATE;"
	rows, err := h.Db.Query(query, user.Id, body.SubcategoryId, body.LevelId)
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	var statId int
	for rows.Next() {
		_ = rows.Scan(&statId)
	}

	var correctAnswers int
	if body.Correct {
		correctAnswers = 1
	}
	if statId == 0 {
		query = "INSERT INTO user_stats SET user_id = ?, subcategory_id = ?, level_id = ?, date = CURRENT_DATE, correctAnswers = ?, answers = 1, seconds = ?;"
		rows, err = h.Db.Query(query, user.Id, body.SubcategoryId, body.LevelId, correctAnswers, body.Seconds)
		defer rows.Close()
		if err != nil {
			panic(err)
		}
		return
	}

	query = "UPDATE user_stats SET correctAnswers = correctAnswers + ?, answers = answers + 1, seconds = seconds + ? WHERE user_id = ? AND subcategory_id = ? AND level_id = ? AND date = CURRENT_DATE;"
	rows, err = h.Db.Query(query, correctAnswers, body.Seconds, user.Id, body.SubcategoryId, body.LevelId)
	defer rows.Close()
	if err != nil {
		panic(err)
	}
}
