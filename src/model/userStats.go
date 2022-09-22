package model

type UserStat struct {
	CorrectAnswers int
	Answers        int
	Seconds        int
}

type User7dStat struct {
	Name string
	Data []int
}

func NewUser7dStat(name string) User7dStat {
	return User7dStat{Name: name, Data: []int{0, 0, 0, 0, 0, 0, 0}}
}

type User7dStats struct {
	Stats          []User7dStat
	TotalTimeSpent int
	AvgCorrectness int
}
