package basefunctions

import (
	"websays/database/basetypes"
)

/*
 * BaseFucntionsInterface is an interface that defines the common database operations
 * to be implemented by different types of connections. Flyweight in nature
 */
type BaseFucntionsInterface interface {
	// GetFunctions returns the BaseFucntionsInterface instance.
	GetFunctions() BaseFucntionsInterface

	// EnsureIndex creates an index for a collection in the database.
	// Parameters:
	//   - dbName: The name of the database.
	//   - collectionName: The name of the collection.
	//   - indexData: The index data or structure.
	// Returns an error if the operation fails.
	EnsureIndex(dbName basetypes.DBName, collectionName basetypes.CollectionName, indexData interface{}) error

	// Add inserts a new document into a collection in the database.
	// Parameters:
	//   - dbName: The name of the database.
	//   - collectionName: The name of the collection.
	//   - data: The data to be inserted.
	// Returns an error if the operation fails.
	Add(dbName basetypes.DBName, collectionName basetypes.CollectionName, data interface{}) error

	// FindOne retrieves a single document from a collection in the database based on the provided query.
	// Parameters:
	//   - dbName: The name of the database.
	//   - collectionName: The name of the collection.
	//   - query: The query to filter the document to be retrieved.
	// Returns the retrieved document and an error if the operation fails.
	FindOne(dbName basetypes.DBName, collectionName basetypes.CollectionName, query interface{}) (interface{}, error)

	// UpdateOne updates a document in a collection in the database based on the provided query.
	// Parameters:
	//   - dbName: The name of the database.
	//   - collectionName: The name of the collection.
	//   - query: The query to filter the document to be updated.
	//   - data: The data to update the document with.
	//   - upsert: Whether to perform an upsert (insert if not found) operation.
	// Returns an error if the operation fails.
	UpdateOne(dbName basetypes.DBName, collectionName basetypes.CollectionName, query interface{}, data interface{}, upsert bool) error

	// DeleteOne deletes a document from a collection in the database based on the provided query.
	// Parameters:
	//   - dbName: The name of the database.
	//   - collectionName: The name of the collection.
	//   - query: The query to filter the document to be deleted.
	// Returns an error if the operation fails.
	DeleteOne(dbName basetypes.DBName, collectionName basetypes.CollectionName, query interface{}) error

	// GetNextID returns the next available ID for document insertion.
	GetNextID() int
}
