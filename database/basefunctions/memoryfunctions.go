package basefunctions

import (
	"errors"
	"strconv"
	"sync"
	"websays/database/basemodels"
	"websays/database/basetypes"
)

var mapInitiater sync.Once

type MemoryFunctions struct {
	lock sync.Mutex
	data map[string]interface{}
	id   int
}

func (u *MemoryFunctions) GetFunctions() BaseFucntionsInterface {
	mapInitiater.Do(func() {
		u.data = map[string]interface{}{}
	})
	return u
}

func (u *MemoryFunctions) EnsureIndex(dbName basetypes.DBName, collectionName basetypes.CollectionName, data interface{}) error {
	return nil
}

func (u *MemoryFunctions) GetNextID() int {
	u.id++
	return u.id
}

func (u *MemoryFunctions) Add(dbName basetypes.DBName, collectionName basetypes.CollectionName, data interface{}) error {
	idData := data.(basemodels.BaseModels)
	u.lock.Lock()
	defer u.lock.Unlock()
	if _, ok := u.data[strconv.FormatInt(int64(idData.GetID()), 10)+"_"+string(collectionName)]; ok {
		return errors.New("Id already exists")
	}
	u.data[strconv.FormatInt(int64(idData.GetID()), 10)+"_"+string(collectionName)] = idData
	return nil
}
func (u *MemoryFunctions) FindOne(dbName basetypes.DBName, collectionName basetypes.CollectionName, condition interface{}) (interface{}, error) {
	idData := condition.(basemodels.BaseModels)
	u.lock.Lock()
	defer u.lock.Unlock()
	if data, ok := u.data[strconv.FormatInt(int64(idData.GetID()), 10)+"_"+string(collectionName)]; ok {
		return data, nil
	} else {
		return nil, errors.New("Not found")
	}
}
func (u *MemoryFunctions) UpdateOne(dbName basetypes.DBName, collectionName basetypes.CollectionName, query interface{}, data interface{}, upsert bool) error {
	u.lock.Lock()
	defer u.lock.Unlock()

	idData := data.(basemodels.BaseModels)
	if _, ok := u.data[strconv.FormatInt(int64(idData.GetID()), 10)+"_"+string(collectionName)]; ok {
		u.data[strconv.FormatInt(int64(idData.GetID()), 10)+"_"+string(collectionName)] = data
	} else {
		return errors.New("data not found")
	}
	return nil
}
func (u *MemoryFunctions) DeleteOne(dbName basetypes.DBName, collectionName basetypes.CollectionName, data interface{}) error {
	u.lock.Lock()
	defer u.lock.Unlock()

	idData := data.(basemodels.BaseModels)
	if _, ok := u.data[strconv.FormatInt(int64(idData.GetID()), 10)+"_"+string(collectionName)]; ok {
		delete(u.data, strconv.FormatInt(int64(idData.GetID()), 10)+"_"+string(collectionName))
	} else {
		return errors.New("data not found")
	}

	return nil
}
