package basefunctions

import (
	"log"
	"websays/database/basetypes"
)

type FileFunctions struct {
	id int
}

func (u *FileFunctions) GetFunctions() BaseFucntionsInterface {
	//TODO: read id from file
	return u
}

func (u *FileFunctions) EnsureIndex(dbName basetypes.DBName, collectionName basetypes.CollectionName, data interface{}) error {
	return nil
}

func (u *FileFunctions) GetNextID() int {
	u.id++
	return u.id
}

func (u *FileFunctions) Add(dbName basetypes.DBName, collectionName basetypes.CollectionName, data interface{}) error {
	log.Println("File add")
	return nil
}
func (u *FileFunctions) FindOne(dbName basetypes.DBName, collectionName basetypes.CollectionName, condition interface{}) (interface{}, error) {

	return nil, nil
}
func (u *FileFunctions) UpdateOne(dbName basetypes.DBName, collectionName basetypes.CollectionName, query string, data interface{}, upsert bool) error {
	return nil
}
func (u *FileFunctions) DeleteOne(dbName basetypes.DBName, collectionName basetypes.CollectionName, query interface{}) error {
	log.Println("Unimplemented DeleteOne MySql")
	return nil
}
