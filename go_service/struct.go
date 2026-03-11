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
	path_env = "../config/.env"
	path_db_config = "../config/db_config.yaml"
	path_fakesintjson = "../data/dataperson.json"
	path_data_dir = "../data"
	path_log_dir = "../logs"
)