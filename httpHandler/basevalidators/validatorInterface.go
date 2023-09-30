package basevalidators

// ValidatorInterface is an interface that defines the contract for data validation in the application.
type ValidatorInterface interface {
	// Validate performs validation on the provided data for a specific API or operation identified by 'apiName'.
	// It returns an error if the data is invalid according to the validation rules.
	//
	// Parameters:
	//   - apiName: A string representing the name or identifier of the API or operation being validated.
	//   - data: An interface{} representing the data to be validated.
	//
	// Returns:
	//   - error: An error indicating validation failure if the data is invalid; otherwise, it returns nil.
	//
	// Example Usage:
	//   err := MyValidator.Validate("CreateArticle", atricleData)
	Validate(apiName string, data interface{}) error
}
