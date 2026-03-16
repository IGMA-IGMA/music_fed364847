package main

type UserJS struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Pwd      string `json:"pwd"`
}

type DBConfig struct {
	DBHost     string `yaml:"host"`
	DBPort     string `yaml:"port"`
	DBName     string `yaml:"name"`
	DBUser     string `yaml:"user"`
	DBPassword string `yaml:"password"`
}

const (
	EnvFilePath        = "../config/.env"
	DBConfigPath       = "../config/db_config.yaml"
	FakeDataJSONPath   = "../data/dataperson.json"
	DataDirPath        = "../data"
	LogDirPath         = "../logs"
	DBConnStringFormat = "postgresql://%s:%s@%s:%s/%s"
)

