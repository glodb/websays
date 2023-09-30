package basefunctions

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"sync"
	"websays/config"
	"websays/database/basetypes"
	"websays/httpHandler/basemodels"
)

// FileFunctions implements the BaseFucntionsInterface for file-based storage.
// It provides methods for ensuring indexes, adding, finding, updating, and deleting data.
type FileFunctions struct {
	runningLock sync.Mutex // Mutex for ensuring thread safety when accessing running number
	filesLock   sync.Mutex // Mutex for ensuring thread safety when accessing files
	id          int        // The running ID
}

// GetFunctions returns the FileFunctions instance as a BaseFucntionsInterface.
func (u *FileFunctions) GetFunctions() BaseFucntionsInterface {
	return u
}

// EnsureIndex ensures an index for the specified database and collection.
// File storage does not require index creation, so this method does nothing.
func (u *FileFunctions) EnsureIndex(dbName basetypes.DBName, collectionName basetypes.CollectionName, data interface{}) error {
	return nil
}

// readRunningNumber reads the running number from a file.
// It takes the filePath as a parameter and returns the running number and any error encountered.
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

// writeRunningNumber writes the running number to a file.
// It takes the filePath and the number to be written as parameters and returns any error encountered.
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

// GetNextID returns the next available ID for file-based storage.
// It reads and increments the running number stored in a file and returns the updated ID.
func (u *FileFunctions) GetNextID() int {
	u.runningLock.Lock()
	defer u.runningLock.Unlock()
	filePath := config.GetInstance().FilePath + "/" + config.GetInstance().RunningFileName
	u.id, _ = u.readRunningNumber(filePath)
	u.id++
	u.writeRunningNumber(filePath, u.id)
	return u.id
}

// Add adds data to the file-based storage.
// It takes the dbName, collectionName, and data to be added as parameters and returns any error encountered.
func (u *FileFunctions) Add(dbName basetypes.DBName, collectionName basetypes.CollectionName, data interface{}) (int, error) {
	idData := data.(basemodels.BaseModels)

	filePath := config.GetInstance().FilePath + "/" + strconv.FormatInt(int64(idData.GetID()), 10) + "_" + string(collectionName)

	// Check if the file with the same ID already exists
	_, err := os.Stat(filePath)

	if err == nil {
		return 0, errors.New("ID already exists")
	}

	u.filesLock.Lock()
	defer u.filesLock.Unlock()

	file, err := os.Create(filePath)
	if err != nil {
		return 0, errors.New("Error opening file path")
	}
	defer file.Close()

	// Create a JSON encoder
	encoder := json.NewEncoder(file)

	err = encoder.Encode(data)
	if err != nil {
		return 0, errors.New("Error encoding JSON")
	}
	return 0, nil
}

// FindOne finds data in the file-based storage by ID.
// It takes the dbName, collectionName, and a data structure to store the result as parameters.
// It returns the found data and any error encountered.
func (u *FileFunctions) FindOne(dbName basetypes.DBName, collectionName basetypes.CollectionName, data interface{}) (interface{}, error) {
	idData := data.(basemodels.BaseModels)

	filePath := config.GetInstance().FilePath + "/" + strconv.FormatInt(int64(idData.GetID()), 10) + "_" + string(collectionName)

	// Check if the file with the specified ID exists
	_, err := os.Stat(filePath)

	if err != nil {
		return nil, errors.New("ID not found")
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
		return nil, errors.New("Error decoding JSON")
	}

	return data, nil
}

// UpdateOne updates data in the file-based storage by ID.
// It takes the dbName, collectionName, query, data, and upsert flag as parameters and returns any error encountered.
func (u *FileFunctions) UpdateOne(dbName basetypes.DBName, collectionName basetypes.CollectionName, query interface{}, data interface{}, upsert bool) error {
	idData := data.(basemodels.BaseModels)

	filePath := config.GetInstance().FilePath + "/" + strconv.FormatInt(int64(idData.GetID()), 10) + "_" + string(collectionName)

	// Check if the file with the specified ID exists
	_, err := os.Stat(filePath)

	if err != nil {
		return errors.New("ID not found")
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
		return errors.New("Error encoding JSON")
	}
	return nil
}

// DeleteOne deletes data from the file-based storage by ID.
// It takes the dbName, collectionName, and data to be deleted as parameters and returns any error encountered.
func (u *FileFunctions) DeleteOne(dbName basetypes.DBName, collectionName basetypes.CollectionName, data interface{}) error {
	idData := data.(basemodels.BaseModels)

	filePath := config.GetInstance().FilePath + "/" + strconv.FormatInt(int64(idData.GetID()), 10) + "_" + string(collectionName)

	// Check if the file with the specified ID exists
	_, err := os.Stat(filePath)

	if err != nil {
		return errors.New("ID not found")
	}

	u.filesLock.Lock()
	defer u.filesLock.Unlock()

	err = os.Remove(filePath)

	if err != nil {
		return errors.New("File not found")
	}
	return nil
}
