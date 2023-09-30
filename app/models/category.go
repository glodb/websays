package models

// Category represents a data model for categorizing items with an ID and a name.
type Category struct {
	ID   int    `json:"id"`   // ID uniquely identifies the category.
	Name string `json:"name"` // Name is the descriptive name of the category.
}

// GetID is a method that implements part of the basemodel interface.
// It returns the unique identifier (ID) of the category.
func (cat Category) GetID() int {
	return cat.ID
}
