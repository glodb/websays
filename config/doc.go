// Package config reads configuration values from a file and provides them through a singleton Config instance.
//
// This package is responsible for reading and parsing configuration values from a file
// and making them accessible throughout the application via a singleton Config instance. It ensures that configuration
// data is loaded only once and can be easily accessed by different parts of the application.
//
// Usage:
// To access configuration values, you can use the GetInstance() function to retrieve the Config instance and then
// call methods on it to retrieve specific configuration values.
//
// Example:
//   // Get the Config instance
//   cfg := GetInstance()
//
//   // Access a configuration value
//   databaseName := cfg.Database.DBName
//
package config
