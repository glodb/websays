package validators

import (
	"errors"
	"websays/app/models"
)

// CategoryValidator is a validator specific to category-related APIs.
// It implements the Validator interface and is responsible for validating
// incoming data for different category-related API endpoints.
type CategoryValidator struct {
}

// Validate performs data validation for category-related API endpoints.
// It takes the name of the API being invoked (apiName) and the data to be validated.
// Depending on the API name, specific validation rules are applied to ensure the
// data meets the required criteria.
//
// Parameters:
//   - apiName: The name of the API being invoked, which determines the validation rules to apply.
//   - data:    The data to be validated, expected to be of type models.Category.
//
// Returns:
//   - error:   An error is returned if validation fails, indicating the specific validation issue.
func (cat *CategoryValidator) Validate(apiName string, data interface{}) error {
	categoryData := data.(models.Category)

	// Apply validation rules based on the API name
	switch apiName {
	case "/api/createCategory":
		// Validate for creating a category
		if categoryData.Name == "" {
			return errors.New("Category Name can't be empty")
		}
	case "/api/updateCategory":
		// Validate for updating a category
		if categoryData.ID <= 0 {
			return errors.New("Category ID is not correct")
		}
		if categoryData.Name == "" {
			return errors.New("Category Name can't be empty")
		}
	}

	// If no validation issues are found, return nil indicating successful validation
	return nil
}
