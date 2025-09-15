package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Model  string `json:"model"`
	ApiURL string `json:"api_url"`
	ApiKey string `json:"api_key"`
}

var configDir string
var configFilename = "ggit.json"

func init() {
	var err error
	configDir, err = os.UserConfigDir()
	if err != nil {
		log.Fatal("failed to get user Confige Directory: %w", err)
	}
}

func GetConfig() (*Config, error) {
	config := Config{
		Model:  "gemini-2.5-flash-preview-05-20",
		ApiURL: "https://generativelanguage.googleapis.com/v1beta/models/",
		ApiKey: "API_KEY",
	}

	if err := os.MkdirAll(configDir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("fauled to create directory: %w", err)
	}

	filePath := filepath.Join(configDir, configFilename)
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		// file does not exist
		data, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			return nil, fmt.Errorf("failed to marashal config: %w", err)
		}
		if err = ioutil.WriteFile(filePath, data, 0644); err != nil {
			return nil, fmt.Errorf("failed to create file: %w", err)
		}

		return &config, nil

	} else if err != nil {
		// other error
		return nil,fmt.Errorf("failed to check file: %w", err)
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil,fmt.Errorf("failed to read file: %w", err)
	}

	var cnfg Config
	err = json.Unmarshal(data, &cnfg)
	if err != nil {
		return nil,fmt.Errorf("failed to unmarash data: %w", err)
	}

	// file exists

	return &cnfg, nil
}
