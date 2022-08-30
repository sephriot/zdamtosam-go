package model

type Exercise struct {
	Id            int    `json:"id,omitempty"`
	Task          string `json:"task,omitempty"`
	Hint          string `json:"hint,omitempty"`
	StepByStep    string `json:"stepByStep,omitempty"`
	Image         string `json:"image,omitempty"`
	LevelId       int    `json:"levelId,omitempty"`
	SubcategoryId int    `json:"subcategoryId,omitempty"`
	UserId        string `json:"userId,omitempty"`
	Visible       int    `json:"visible,omitempty"`
	Date          string `json:"date,omitempty"`
}
