package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	configModels "websays/config/configModels"
)

// config is a singleton struct that holds configuration values for the application.
type config struct {
	Server          configModels.ServerConfig   `json:"server"`
	Database        configModels.DatabaseConfig `json:"database"`
	FilePath        string                      `json:"filesPath"`
	RunningFileName string                      `json:"runningFileName"`
	Controllers     []string                    `json:"controllers"`
}

var (
	instance *config
	once     sync.Once
)

// GetInstance returns the singleton instance of the config struct.
// If the instance does not exist, it is created once.
func GetInstance() *config {
	once.Do(func() {
		instance = &config{}
	})
	return instance
}

// Setup initializes the configuration by loading values from a JSON file.
// It expects the path to the configuration file as a parameter.
//
// Parameters:
//   - path: The path to the JSON configuration file.
func (c *config) Setup(path string) {
	configFile, err := os.Open(path)
	if err != nil {
		log.Println("Error opening config file:", err)
		return
	}
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&c)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
}
