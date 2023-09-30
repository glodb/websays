package baseinterfaces

import (
	"websays/database/basefunctions"
	"websays/database/basetypes"
	"websays/httpHandler/basevalidators"
)

// Controller is the primary interface that defines the contract for all controllers within the application.
// It combines multiple responsibilities and functionalities related to database interactions, validation, and API routing.
type Controller interface {
	// BaseFucntionsInterface provides common database operations such as querying, updating, and deleting records.
	basefunctions.BaseFucntionsInterface

	// BaseControllerFactory allows controllers to access and interact with other controllers when necessary.
	BaseControllerFactory

	// ValidatorInterface is implemented for input validation, ensuring data integrity and consistency.
	basevalidators.ValidatorInterface

	// SetBaseFunctions sets the base database functions that are used by this controller, enabling it to interact with the database.
	SetBaseFunctions(basefunctions.BaseFucntionsInterface)

	// GetCollectionName returns the name of the database collection associated with this controller,
	// facilitating the identification of the data store for this controller.
	GetCollectionName() basetypes.CollectionName

	// DoIndexing is responsible for handling indexing operations specific to the controller's data,
	// optimizing data retrieval and query performance.
	// It may create or update indexes in the underlying database.
	DoIndexing() error

	// RegisterApis is used to define and register API endpoints associated with this controller.
	// These endpoints are responsible for exposing controller-specific functionalities to external clients.
	RegisterApis()

	// GetDBName returns the name of the database where the controller's data is stored.
	// It allows the controller to identify the appropriate database for its operations.
	GetDBName() basetypes.DBName
}
