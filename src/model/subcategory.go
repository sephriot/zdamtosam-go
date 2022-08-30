package model

type Subcategory struct {
	Id         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	CategoryId int    `json:"categoryId,omitempty"`
}
