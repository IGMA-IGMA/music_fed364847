package main

import (
	"fmt"
	"os"
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
	_, err := os.Create("dataperson.json")
	defer file.close()
	if err != nil{
		return err
	}
	
}
