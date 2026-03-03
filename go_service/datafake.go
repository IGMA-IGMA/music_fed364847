package main

import (
	"fmt"

	"github.com/go-faker/faker/v4"
)

func usergen() *FakeUser {
	user := FakeUser{}
	err := faker.FakeData(&user)
	if err != nil {
		fmt.Println(err)
	}
	return &user
}

func gensintjson(int n) {
	
}
