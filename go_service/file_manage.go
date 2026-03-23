package main

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

func createFile(dir string, filename string) (*os.File, error) {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return nil, err
	}

	pathFile := FileJoin(dir, filename)
	_, err = os.Stat(pathFile)
	if err == nil {
		err := os.Remove(pathFile)
		if err != nil {
			return nil, err
		}

	}
	file, err := os.Create(pathFile)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func FileJoin(dir string, filename string) string {
	return filepath.Join(dir, filename)
}

func ParsConfig() (*DBConfig, error) {
	if err := godotenv.Load(EnvFilePath); err != nil {
		loggerFileManage.Error("Note: .env file not found, using system env")
	}

	data, err := os.ReadFile(DBConfigPath)
	if err != nil {
		return nil, err
	}

	replacedData := os.ExpandEnv(string(data))

	config := &DBConfig{}
	err = yaml.Unmarshal([]byte(replacedData), config)
	if err != nil {
		loggerFileManage.Error("ENV: The configuration was not received successfully", zap.String("error", err.Error()), zap.String("EnvFile", DBConfigPath))
		return nil, err
	}

	return config, nil
}
