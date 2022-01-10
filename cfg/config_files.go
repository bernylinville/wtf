package cfg

import (
	"fmt"
	"os"

	"github.com/olebedev/config"
)

// Initialize takes care of settings up the initial state of WTF configuration
// It ensures necessary directories and files exist
func Initialize(hasCustom bool) {}

// LoadWtfConfigFile loads the specified config file
func LoadWtfConfigFile(filePath string) *config.Config {
	cfg, err := config.ParseYamlFile(filePath)
	if err != nil {
		fmt.Printf("ParseYamlFile")
		fmt.Println(filePath, err)
		os.Exit(1)
	}

	return cfg
}
