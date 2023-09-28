package models

type Article struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (art Article) GetID() int {
	return art.ID
}
