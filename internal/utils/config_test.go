package utils

import (
	"fmt"
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {

	//check if the file is present and correctly formatted
	t.Run("LoadConfig with .env file", func(t *testing.T) {
		wd, err := os.Getwd()
		fmt.Printf("pwd: %v\n", wd)

		config, err := LoadConfig(".")
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		expectedConfig := Config{
			DBUser:   "someone",
			DBName:   "some-db",
			DBSource: "postgresql://someone:ss@localhost:5432/some-db?sslmode=disable",
		}
		if config != expectedConfig {
			t.Errorf("Got %v, want %v", config, expectedConfig)
		}

	})
	t.Run("LoadConfig with missing .env file", func(t *testing.T) {
		_, err := LoadConfig(".")
		if err != nil {
			t.Errorf("Expected error: %v", err)
			return
		}

	})
}
