package config

import (
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

func getConfigDir(configDir string) string {
	if configDir == "" {
		configDir = os.Getenv("KUBESWITCHER_CONFIG_DIR")
		if configDir == "" {
			log.Fatal("KUBESWITCHER_CONFIG_DIR environment variable is not set")
		}
	}
	return configDir
}

// getAllYamlFiles returns all yaml files in the given directory
func getAllYamlFiles(configDir string) ([]string, error) {
	// Read all files in the directory
	files, err := os.ReadDir(configDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory %s: %v", configDir, err)
	}
	log.Debugf("read %d files from directory %s", len(files), configDir)

	// Filter yaml files
	var yamlFiles []string
	for _, file := range files {
		// Skip directories, only process files
		if file.IsDir() {
			log.Debugf("skipping directory %s", file.Name())
			continue
		}

		// Skip files that are not yaml files
		filePath := filepath.Join(configDir, file.Name())
		if filepath.Ext(filePath) != ".yaml" && filepath.Ext(filePath) != ".yml" {
			log.Debugf("skipping file %s as it is not a yaml file", filePath)
			continue
		}

		yamlFiles = append(yamlFiles, filePath)
	}

	log.Debugf("found %d yaml files in directory %s", len(yamlFiles), configDir)
	return yamlFiles, nil
}