package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	gensintjson(100)
	cfg, err := ParsConfig()
	if err != nil {
		log.Fatalf("Ошибка при парсинге конфига: %v", err)
	}

	db, err := NewConnect(cfg)
	if err != nil {
		log.Fatalf("Ошибка при подключении к БД: %v", err)
	}
	defer db.Close()
	jsonFile, _ := os.Open(FakeDataJSONPath)
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var users []UserJS
	json.Unmarshal(byteValue, &users)
	for _, user := range users {
		db.CreateUser(context.Background(), &user)
	}

}
