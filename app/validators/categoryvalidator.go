package validators

import (
	"errors"
	"websays/app/models"
)

type CategoryValidator struct {
}

func (cat *CategoryValidator) Validate(apiName string, data interface{}) error {
	categoryData := data.(models.Category)
	switch apiName {
	case "/api/createCategory":
		if categoryData.Name == "" {
			return errors.New("Category Name can't be empty")
		}
	case "/api/updateCategory":
		if categoryData.ID <= 0 {
			return errors.New("Category Id is not correct")
		}
		if categoryData.Name == "" {
			return errors.New("Category Name can't be empty")
		}
	}
	return nil
}
