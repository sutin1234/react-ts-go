package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BaseUrl    string
	ApiPrefix  string
	ApiVersion string
	Port       string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env.local file")
	}

	cfg := &Config{
		BaseUrl:    os.Getenv("BASE_URL"),
		ApiPrefix:  os.Getenv("API_FREFIX"),
		ApiVersion: os.Getenv("API_VERSION"),
		Port:       os.Getenv("PORT"),
	}
	config, _ := json.Marshal(cfg)
	fmt.Printf("Loaded Config Success %v\n", string(config))
	return cfg, nil
}
