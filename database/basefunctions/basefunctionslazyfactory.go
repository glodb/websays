package basefunctions

import (
	"errors"
	"sync"
	"websays/database/basetypes"
)

/*
 * baseFunctions is a singleton factory for creating instances of BaseFucntionsInterface.
 * It provides lazy initialization of BaseFucntionsInterface objects for different database types.
 * This is part of the flyweight design pattern.
 */
type baseFunctions struct {
	dbfunctions map[basetypes.DbType]*BaseFucntionsInterface
}

var instance *baseFunctions
var once sync.Once

// GetInstance returns a single object of the baseFunctions factory.
// This is a pure lazy factory, and it does not create the functions class until the dbname is specifically passed.
// This is also part of the flyweight design pattern.
func GetInstance() *baseFunctions {
	once.Do(func() {
		instance = &baseFunctions{}
		instance.dbfunctions = make(map[basetypes.DbType]*BaseFucntionsInterface)
	})
	return instance
}

// GetFunctions returns an instance of BaseFucntionsInterface for the specified database type and name.
// If the instance already exists, it returns the existing instance; otherwise, it creates a new one.
func (u *baseFunctions) GetFunctions(dbType basetypes.DbType, dbName basetypes.DBName) (*BaseFucntionsInterface, error) {
	if connection, ok := u.dbfunctions[dbType]; ok {
		return connection, nil
	}
	switch dbType {
	case basetypes.MYSQL:
		{
			connection := MySqlFunctions{}
			functionsInterface := connection.GetFunctions()

			u.dbfunctions[dbType] = &functionsInterface
			return u.dbfunctions[dbType], nil
		}
	case basetypes.FILE:
		{
			connection := FileFunctions{}
			functionsInterface := connection.GetFunctions()

			u.dbfunctions[dbType] = &functionsInterface
			return u.dbfunctions[dbType], nil
		}
	case basetypes.MEMORY:
		{
			connection := MemoryFunctions{}
			functionsInterface := connection.GetFunctions()

			u.dbfunctions[dbType] = &functionsInterface
			return u.dbfunctions[dbType], nil
		}
	}
	return nil, errors.New("Not configured for this db")
}
