package main

import (
	"log"
	"strings"

	"github.com/Matltin/git-cli-ai/utils"
)

func main() {
	config, err := utils.GetConfig()
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}

	if strings.Contains(config.Model, "gemini") {
		
	}

}
