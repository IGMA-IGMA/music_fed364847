package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-faker/faker/v4"
)

func gensintjson(n int) error {

	_, err := os.Stat("dataperson.json")
	if !os.IsExist(err) {
		os.Remove("dataperson.json")
	}

	file, err := os.Create("dataperson.json")
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
