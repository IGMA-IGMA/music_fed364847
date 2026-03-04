package main


type UserJS struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Pwd      string `json:"pwd"`
}


const ( 
	path_env = "../config/.env"
	path_db_config = "../config/db_config.yaml"
	path_fakesintjson = "../data/dataperson.json"
)