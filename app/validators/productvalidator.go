package validators

import (
	"errors"
	"websays/app/models"
)

type ProductValidator struct {
}

func (pro *ProductValidator) Validate(apiName string, data interface{}) error {
	proData := data.(models.Product)
	switch apiName {
	case "/api/addProduct":
		if proData.Name == "" {
			return errors.New("Product Name can't be empty")
		}
	}
	return nil
}
