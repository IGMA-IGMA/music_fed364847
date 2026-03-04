package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type DBConfig struct {
	DBHost     string `yaml:"host"`
	DBPort     string `yaml:"port"` 
	DBName     string `yaml:"name"`
	DBUser     string `yaml:"user"`
	DBPassword string `yaml:"password"`
}

func NewConfig() (*DBConfig, error) {

	if err := godotenv.Load(path_env); err != nil {
		log.Println("Note: .env file not found, using system env")
	}

	data, err := os.ReadFile(path_db_config)
	if err != nil {
		return nil, err
	}

	replacedData := os.ExpandEnv(string(data))

	config := &DBConfig{}
	err = yaml.Unmarshal([]byte(replacedData), config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
