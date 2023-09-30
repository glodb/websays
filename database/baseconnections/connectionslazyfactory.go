package baseconnections

import (
	"sync"
	"websays/database/basetypes"
)

// dbConnections is a struct representing the database connections manager.
type dbConnections struct {
	dbconnections map[basetypes.DbType]*ConnectionInterface
}

var instance *dbConnections
var once sync.Once

// GetInstance returns a single instance of the dbConnections manager.
// It follows the lazy initialization approach, creating database connections only when needed.
func GetInstance() *dbConnections {
	once.Do(func() {
		instance = &dbConnections{}
		instance.dbconnections = make(map[basetypes.DbType]*ConnectionInterface)
	})
	return instance
}

// GetConnection retrieves or creates a connection based on the specified database type.
// It returns the corresponding connection interface.
func (u *dbConnections) GetConnection(dbType basetypes.DbType) ConnectionInterface {
	if connection, ok := u.dbconnections[dbType]; ok {
		return *connection
	}
	switch dbType {
	case basetypes.MYSQL:
		{
			connection := MysqlConnection{}
			mysqlconnector, err := connection.CreateConnection()
			if err != nil {
				return nil
			}
			u.dbconnections[dbType] = &mysqlconnector
			return *u.dbconnections[dbType]
		}
	case basetypes.FILE:
		{ // Allowing file connections
			connection := FileConnection{}
			fileconnector, err := connection.CreateConnection()
			if err != nil {
				return nil
			}
			u.dbconnections[dbType] = &fileconnector
			return *u.dbconnections[dbType]
		}
	case basetypes.MEMORY:
		{ // Allowing memory connection
			connection := MemoryConnection{}
			memoryconnector, err := connection.CreateConnection()
			if err != nil {
				return nil
			}
			u.dbconnections[dbType] = &memoryconnector
			return *u.dbconnections[dbType]
		}
	}
	return nil
}
