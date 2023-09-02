package config

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"

	"github.com/joho/godotenv"
	logs "github.com/sirupsen/logrus"
)

type Config struct {
	BaseUrl    string
	ApiPrefix  string
	ApiVersion string
	Port       string
}

const (
	projectDirName = "react-ts-go"
)

func LoadConfig() (*Config, error) {
	re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		if err != nil {
			logs.WithFields(logs.Fields{
				"cause": err,
				"cwd":   cwd,
			}).Fatalf("Problem loading .env file %v", err)
		}
	}
	fmt.Printf("Loaded .enf Success %v/.env\n", rootPath)

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
