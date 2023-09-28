package basefunctions

import (
	"database/sql"
	"log"
	"websays/database/baseconnections"
	"websays/database/basetypes"
)

type MemoryFunctions struct {
}

func (u *MemoryFunctions) GetFunctions() BaseFucntionsInterface {
	return u
}

func (u *MemoryFunctions) EnsureIndex(dbName basetypes.DBName, collectionName basetypes.CollectionName, data interface{}) error {
	return nil
}

func (u *MemoryFunctions) Add(dbName basetypes.DBName, collectionName basetypes.CollectionName, data interface{}) error {
	return nil
}
func (u *MemoryFunctions) FindOne(dbName basetypes.DBName, collectionName basetypes.CollectionName, condition map[string]interface{}, result interface{}) (interface{}, error) {

	return nil, nil
}
func (u *MemoryFunctions) UpdateOne(dbName basetypes.DBName, collectionName basetypes.CollectionName, query string, data []interface{}, upsert bool) error {
	conn := baseconnections.GetInstance().GetConnection(basetypes.MYSQL).GetDB(basetypes.MYSQL).(*sql.DB)
	_, err := conn.Exec(query, data...)
	return err
}
func (u *MemoryFunctions) DeleteOne(dbName basetypes.DBName, collectionName basetypes.CollectionName, query interface{}) error {
	log.Println("Unimplemented DeleteOne MySql")
	return nil
}
