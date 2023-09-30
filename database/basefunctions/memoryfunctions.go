package basefunctions

import (
	"errors"
	"strconv"
	"sync"
	"websays/database/basetypes"
	"websays/httpHandler/basemodels"
)

var mapInitiater sync.Once

// MemoryFunctions is a concrete implementation of the BaseFucntionsInterface for in-memory storage.
type MemoryFunctions struct {
	lock sync.Mutex             // Mutex for locking access to the in-memory data store.
	data map[string]interface{} // The in-memory data store where data is stored.
	id   int                    // ID counter for generating unique IDs.
}

// GetFunctions returns the MemoryFunctions instance as a BaseFucntionsInterface.
func (u *MemoryFunctions) GetFunctions() BaseFucntionsInterface {
	mapInitiater.Do(func() {
		u.data = map[string]interface{}{}
	})
	return u
}

// EnsureIndex ensures an index for the specified database and collection.
// Memory storage does not require index creation, so this method does nothing.
func (u *MemoryFunctions) EnsureIndex(dbName basetypes.DBName, collectionName basetypes.CollectionName, data interface{}) error {
	return nil
}

// GetNextID generates and returns the next available ID for in-memory storage.
func (u *MemoryFunctions) GetNextID() int {
	u.id++
	return u.id
}

// Add adds data to the in-memory data store.
func (u *MemoryFunctions) Add(dbName basetypes.DBName, collectionName basetypes.CollectionName, data interface{}) error {
	idData := data.(basemodels.BaseModels)
	u.lock.Lock()
	defer u.lock.Unlock()
	key := strconv.FormatInt(int64(idData.GetID()), 10) + "_" + string(collectionName)
	if _, ok := u.data[key]; ok {
		return errors.New("ID already exists")
	}
	u.data[key] = idData
	return nil
}

// FindOne retrieves data from the in-memory data store by ID.
func (u *MemoryFunctions) FindOne(dbName basetypes.DBName, collectionName basetypes.CollectionName, condition interface{}) (interface{}, error) {
	idData := condition.(basemodels.BaseModels)
	u.lock.Lock()
	defer u.lock.Unlock()
	key := strconv.FormatInt(int64(idData.GetID()), 10) + "_" + string(collectionName)
	if data, ok := u.data[key]; ok {
		return data, nil
	} else {
		return nil, errors.New("Not found")
	}
}

// UpdateOne updates data in the in-memory data store by ID.
func (u *MemoryFunctions) UpdateOne(dbName basetypes.DBName, collectionName basetypes.CollectionName, query interface{}, data interface{}, upsert bool) error {
	u.lock.Lock()
	defer u.lock.Unlock()

	idData := data.(basemodels.BaseModels)
	key := strconv.FormatInt(int64(idData.GetID()), 10) + "_" + string(collectionName)
	if _, ok := u.data[key]; ok {
		u.data[key] = data
	} else {
		return errors.New("Data not found")
	}
	return nil
}

// DeleteOne deletes data from the in-memory data store by ID.
func (u *MemoryFunctions) DeleteOne(dbName basetypes.DBName, collectionName basetypes.CollectionName, data interface{}) error {
	u.lock.Lock()
	defer u.lock.Unlock()

	idData := data.(basemodels.BaseModels)
	key := strconv.FormatInt(int64(idData.GetID()), 10) + "_" + string(collectionName)
	if _, ok := u.data[key]; ok {
		delete(u.data, key)
	} else {
		return errors.New("Data not found")
	}

	return nil
}
