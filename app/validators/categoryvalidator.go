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
	case "/api/addCategory":
		if categoryData.Name == "" {
			return errors.New("Category Name can't be empty")
		}
	}
	return nil
}
