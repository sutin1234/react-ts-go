package config

import (
	"os"
	"regexp"
	"testing"

	"github.com/joho/godotenv"
)

func TestConfig(t *testing.T) {
	t.Run("Test Load config env file", func(t *testing.T) {
		re := regexp.MustCompile(`^(.*` + "go-server" + `)`)
		cwd, _ := os.Getwd()
		rootPath := re.Find([]byte(cwd))
		err := godotenv.Load(string(rootPath) + `/.env`)
		if err != nil {
			t.Errorf("load .env failed: %v", err)
		}
	})
}
