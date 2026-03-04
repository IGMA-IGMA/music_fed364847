package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-faker/faker/v4"
)

func gensintjson(n int) error {

	_, err := os.Stat(path_fakesintjson)
	if !os.IsExist(err) {
		os.Remove(path_fakesintjson)
	}

	file, err := os.Create(path_fakesintjson)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")

	users := make([]*UserJS, 0, n)
	for i := 1; i <= n; i++ {
		user_js := UserJS{ID: i, Username: faker.Username(), Email: faker.Email(), Pwd: faker.Password()}
		users = append(users, &user_js)
	}

	encoder.Encode(users)

	return nil
}
