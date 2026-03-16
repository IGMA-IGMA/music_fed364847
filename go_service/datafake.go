package main

import (
	"encoding/json"

	"github.com/go-faker/faker/v4"
)

func gensintjson(n int) error {
	file, err := createFile(DataDirPath, FakeDataJSONPath)
	if err != nil {
		return err
	}

	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")

	users := make([]*UserJS, 0, n)
	for i := 1; i <= n; i++ {
		userJSON := UserJS{ID: i, Username: faker.Username(), Email: faker.Email(), Pwd: faker.Password()}
		users = append(users, &userJSON)
	}

	encoder.Encode(users)

	return nil
}
