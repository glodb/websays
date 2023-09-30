package validators

import (
	"errors"
	"websays/app/models"
)

// ProductValidator is a validator specific to product-related APIs.
// It implements the Validator interface and is responsible for validating
// incoming data for different product-related API endpoints.
type ProductValidator struct {
}

// Validate performs data validation for product-related API endpoints.
// It takes the name of the API being invoked (apiName) and the data to be validated.
// Depending on the API name, specific validation rules are applied to ensure the
// data meets the required criteria.
//
// Parameters:
//   - apiName: The name of the API being invoked, which determines the validation rules to apply.
//   - data:    The data to be validated, expected to be of type models.Product.
//
// Returns:
//   - error:   An error is returned if validation fails, indicating the specific validation issue.
func (pro *ProductValidator) Validate(apiName string, data interface{}) error {
	proData := data.(models.Product)

	// Apply validation rules based on the API name
	switch apiName {
	case "/api/addProduct":
		// Validate for adding a product
		if proData.Name == "" {
			return errors.New("Product Name can't be empty")
		}
	case "/api/updateProduct":
		// Validate for updating a product
		if proData.ID <= 0 {
			return errors.New("ID is not proper")
		}
		if proData.Name == "" {
			return errors.New("Product Name can't be empty")
		}
	}

	// If no validation issues are found, return nil indicating successful validation
	return nil
}
