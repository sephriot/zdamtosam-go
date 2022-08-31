package model

type Exercise struct {
	Id            int      `json:"id,omitempty"`
	Task          string   `json:"task,omitempty"`
	Answer        string   `json:"answer,omitempty"`
	Options       []string `json:"options,omitempty"`
	Hint          string   `json:"hint,omitempty"`
	StepByStep    string   `json:"stepByStep,omitempty"`
	Image         string   `json:"image,omitempty"`
	LevelId       int      `json:"levelId,omitempty"`
	SubcategoryId int      `json:"subcategoryId,omitempty"`
	UserId        string   `json:"userId,omitempty"`
	Visible       int      `json:"visible,omitempty"`
	Date          string   `json:"date,omitempty"`
	NextId        int      `json:"nextId,omitempty"`
	PreviousId    int      `json:"previousId,omitempty"`
}
