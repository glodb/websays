package baseconnections

import "websays/database/basetypes"

// The ConnectionInterface is the primary interface for the package, defining essential methods
// for creating connections to underlying storage systems and retrieving the underlying database.
type ConnectionInterface interface {
	// CreateConnection creates a new connection to the underlying storage and returns it.
	// It may return an error if the connection cannot be established.
	CreateConnection() (ConnectionInterface, error)

	// GetDB returns the underlying database instance based on the specified database type.
	// It takes a DbType parameter and returns the corresponding database interface.
	GetDB(dbType basetypes.DbType) interface{}
}
