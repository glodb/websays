package models

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (cat Category) GetID() int {
	return cat.ID
}
