package utils

import (
	"log"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Get the current directory of the running test or app
	envPath, err := filepath.Abs(".env")
	if err != nil {
		log.Fatal("Error getting absolute path to .env file")
	}

	// Load the .env file when not using Docker
	err = godotenv.Load(envPath)
	if err != nil {
		log.Println("Could not load .env from current directory, attempting parent directory...")

		// Try to load from the parent directory (for test environments)
		parentEnvPath, err := filepath.Abs("../../.env")
		if err != nil {
			log.Fatal("Error getting absolute path to .env file in parent directory")
		}

		err = godotenv.Load(parentEnvPath)
		if err != nil {
			log.Fatal("Error loading .env file from parent directory")
		}
	}
}
