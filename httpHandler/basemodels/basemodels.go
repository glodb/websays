package basemodels

// BaseModels is an interface that defines a common method for obtaining the unique identifier (ID) of a model.
// It serves as a contract for any data structure or object that represents a model in the application.
// The GetID method should return an integer representing the unique identifier of the model.
// This interface enables consistent access to IDs across different model types and is often used in database operations
// and interactions with models.
type BaseModels interface {
	// GetID returns the unique identifier (ID) of the model.
	GetID() int
}
