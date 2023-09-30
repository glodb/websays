package models

// Article is a simple data model representing an article entity with essential attributes.
type Article struct {
	ID    int    `json:"id"`    // ID uniquely identifies the article.
	Title string `json:"title"` // Title is the title or headline of the article.
	Body  string `json:"body"`  // Body contains the main content of the article.
}

// GetID is a method that implements part of the basemodel interface.
// It returns the unique identifier (ID) of the article.
func (art Article) GetID() int {
	return art.ID
}
