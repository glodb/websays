package basefunctions

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"sync"
	"websays/config"
	"websays/database/basemodels"
	"websays/database/basetypes"
)

type FileFunctions struct {
	runningLock sync.Mutex
	filesLock   sync.Mutex
	id          int
}

func (u *FileFunctions) GetFunctions() BaseFucntionsInterface {
	//TODO: read id from file
	return u
}

func (u *FileFunctions) EnsureIndex(dbName basetypes.DBName, collectionName basetypes.CollectionName, data interface{}) error {
	return nil
}

// Read the running number from a file
func (u *FileFunctions) readRunningNumber(filePath string) (int, error) {
	// Read the content of the file
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return 0, err
	}

	// Convert the content to an integer
	runningNumber, err := strconv.Atoi(string(content))
	if err != nil {
		return 0, err
	}

	return runningNumber, nil
}

// Write the running number to a file
func (u *FileFunctions) writeRunningNumber(filePath string, number int) error {
	// Convert the number to a string
	numberStr := strconv.Itoa(number)

	// Write the string to the file
	err := ioutil.WriteFile(filePath, []byte(numberStr), 0644)
	if err != nil {
		return err
	}

	return nil
}

func (u *FileFunctions) GetNextID() int {
	u.runningLock.Lock()
	defer u.runningLock.Unlock()

	filePath := config.GetInstance().FilePath + "/" + config.GetInstance().RunningFileName
	u.id, _ = u.readRunningNumber(filePath)
	u.id++
	u.writeRunningNumber(filePath, u.id)
	return u.id
}

func (u *FileFunctions) Add(dbName basetypes.DBName, collectionName basetypes.CollectionName, data interface{}) error {
	idData := data.(basemodels.BaseModels)

	filePath := config.GetInstance().FilePath + "/" + strconv.FormatInt(int64(idData.GetID()), 10) + "_" + string(collectionName)
	// Use the os.Stat function to get file information
	_, err := os.Stat(filePath)

	if err == nil {
		return errors.New("Id already exists")
	}

	u.filesLock.Lock()
	defer u.filesLock.Unlock()

	file, err := os.Create(filePath)
	if err != nil {
		return errors.New("Error opening file path")
	}
	defer file.Close()

	// Create a JSON encoder
	encoder := json.NewEncoder(file)

	err = encoder.Encode(data)
	if err != nil {
		return errors.New("Error encoding json")
	}
	return nil
}
func (u *FileFunctions) FindOne(dbName basetypes.DBName, collectionName basetypes.CollectionName, data interface{}) (interface{}, error) {
	idData := data.(basemodels.BaseModels)

	filePath := config.GetInstance().FilePath + "/" + strconv.FormatInt(int64(idData.GetID()), 10) + "_" + string(collectionName)
	// Use the os.Stat function to get file information
	_, err := os.Stat(filePath)

	if err != nil {
		return nil, errors.New("Id not found")
	}

	u.filesLock.Lock()
	defer u.filesLock.Unlock()

	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("Error opening file path")
	}
	defer file.Close()

	// Create a JSON decoder
	decoder := json.NewDecoder(file)

	err = decoder.Decode(&data)
	if err != nil {
		return nil, errors.New("Error decoding json")
	}

	return data, nil
}
func (u *FileFunctions) UpdateOne(dbName basetypes.DBName, collectionName basetypes.CollectionName, query string, data interface{}, upsert bool) error {
	idData := data.(basemodels.BaseModels)

	filePath := config.GetInstance().FilePath + "/" + strconv.FormatInt(int64(idData.GetID()), 10) + "_" + string(collectionName)
	// Use the os.Stat function to get file information
	_, err := os.Stat(filePath)

	if err != nil {
		return errors.New("Id not found")
	}

	u.filesLock.Lock()
	defer u.filesLock.Unlock()

	file, err := os.Create(filePath)
	if err != nil {
		return errors.New("Error opening file path")
	}
	defer file.Close()

	// Create a JSON encoder
	encoder := json.NewEncoder(file)

	err = encoder.Encode(data)
	if err != nil {
		return errors.New("Error encoding json")
	}
	return nil
}
func (u *FileFunctions) DeleteOne(dbName basetypes.DBName, collectionName basetypes.CollectionName, data interface{}) error {
	idData := data.(basemodels.BaseModels)

	filePath := config.GetInstance().FilePath + "/" + strconv.FormatInt(int64(idData.GetID()), 10) + "_" + string(collectionName)
	// Use the os.Stat function to get file information
	_, err := os.Stat(filePath)

	if err != nil {
		return errors.New("Id not found")
	}

	u.filesLock.Lock()
	defer u.filesLock.Unlock()

	err = os.Remove(filePath)

	if err != nil {
		return errors.New("file not found")
	}
	return nil
}
