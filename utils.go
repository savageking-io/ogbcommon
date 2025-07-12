package shared

import (
	"errors"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// SetLogLevel sets the global logging level
func SetLogLevel(level string) error {
	lvl, err := log.ParseLevel(level)
	if err != nil {
		log.SetLevel(log.InfoLevel)
		return err
	}
	log.SetLevel(lvl)
	return nil
}

// ReadYAMLConfig reads a YAML configuration file and unmarshals it into the provided interface
// config parameter should be a pointer to a struct that matches the YAML structure
func ReadYAMLConfig(filePath string, config interface{}) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return errors.New(fmt.Sprintf("configuration file not found: %s", filePath))
	}

	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		return errors.New(fmt.Sprintf("error reading configuration file: %v", err))
	}

	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		return errors.New(fmt.Sprintf("error parsing YAML configuration: %v", err))
	}

	return nil
}
