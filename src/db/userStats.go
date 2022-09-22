package db

import (
	"database/sql"
	"log"
	"sort"
	"time"
	"zdamtosam.pl/src/model"
)

func sliceSum(in []int) int {
	var ret int
	for _, v := range in {
		ret += v
	}
	return ret
}

func GetUserStatsForLast7Days(db *sql.DB, userId string) model.User7dStats {
	var finalStats model.User7dStats
	if userId == "" {
		return finalStats
	}

	rows, err := db.Query("SELECT name, date, correctAnswers, answers, seconds from user_stats JOIN subcategories s on user_stats.subcategory_id = s.id WHERE user_id = ? AND date >= DATE_SUB(NOW(), INTERVAL 7 DAY) ORDER BY answers DESC;", userId)
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	dataMap := make(map[string]map[time.Time]model.UserStat)

	for rows.Next() {
		var name string
		var dateString string
		var userStat model.UserStat
		_ = rows.Scan(&name, &dateString, &userStat.CorrectAnswers, &userStat.Answers, &userStat.Seconds)
		if _, ok := dataMap[name]; !ok {
			dataMap[name] = make(map[time.Time]model.UserStat)
		}
		date, err := time.Parse("2006-01-02", dateString)
		if err != nil {
			log.Println("Failed to parse date", date)
		}
		dataMap[name][date] = userStat
	}

	var correctAnswers int
	var answers int

	for name, timeMap := range dataMap {
		finalStats.Stats = append(finalStats.Stats, model.NewUser7dStat(name))
		for date, stats := range timeMap {
			dataIndex := 6 + date.Day() - time.Now().Day()
			finalStats.Stats[len(finalStats.Stats)-1].Data[dataIndex] = stats.Answers
			answers += stats.Answers
			correctAnswers += stats.CorrectAnswers
			finalStats.TotalTimeSpent += stats.Seconds
		}
	}

	sort.Slice(finalStats.Stats, func(i, j int) bool {
		return sliceSum(finalStats.Stats[i].Data) > sliceSum(finalStats.Stats[j].Data)
	})

	if answers > 0 {
		finalStats.AvgCorrectness = 100 * correctAnswers / answers
	}

	return finalStats
}
