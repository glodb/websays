package validators

import (
	"errors"
	"websays/app/models"
)

// ArticleValidator is a validator specific to article-related APIs.
// It implements the Validator interface and is responsible for validating
// incoming data for different article-related API endpoints.
type ArticleValidator struct {
}

// Validate performs data validation for article-related API endpoints.
// It takes the name of the API being invoked (apiName) and the data to be validated.
// Depending on the API name, specific validation rules are applied to ensure the
// data meets the required criteria.
//
// Parameters:
//   - apiName: The name of the API being invoked, which determines the validation rules to apply.
//   - data:    The data to be validated, expected to be of type models.Article.
//
// Returns:
//   - error:   An error is returned if validation fails, indicating the specific validation issue.
func (art *ArticleValidator) Validate(apiName string, data interface{}) error {
	articleData := data.(models.Article)

	// Apply validation rules based on the API name
	switch apiName {
	case "/api/addArticle":
		// Validate for adding an article
		if articleData.Title == "" {
			return errors.New("Title can't be empty")
		}

		if articleData.Body == "" {
			return errors.New("Body can't be empty")
		}
	case "/api/upateArticle":
		// Validate for updating an article
		if articleData.ID <= 0 {
			return errors.New("ID is not valid")
		}

		if articleData.Title == "" {
			return errors.New("Title can't be empty")
		}

		if articleData.Body == "" {
			return errors.New("Body can't be empty")
		}
	}

	// If no validation issues are found, return nil indicating successful validation
	return nil
}
