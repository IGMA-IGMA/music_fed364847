package main

import (
	"context"
	"fmt"
	"log"
)

func init() {
	initLog()
}

// func main() {
// 	gensintjson(100)
// 	cfg, err := ParsConfig()
// 	if err != nil {
// 		log.Fatalf("Ошибка при парсинге конфига: %v", err)
// 	}

// 	db, err := NewConnect(cfg)
// 	if err != nil {
// 		log.Fatalf("Ошибка при подключении к БД: %v", err)
// 	}
// 	defer db.Close()
// 	jsonFile, _ := os.Open(FakeDataJSONPath)
// 	defer jsonFile.Close()
// 	byteValue, _ := ioutil.ReadAll(jsonFile)
// 	var users []UserJS
// 	json.Unmarshal(byteValue, &users)
// 	for _, user := range users {
// 		db.CreateUser(context.Background(), &user)
// 	}

// }

func main() {
	// gensintjson(100)
	cfg, err := ParsConfig()
	if err != nil {
		log.Fatalf("Ошибка при парсинге конфига: %v", err)
	}

	db, err := NewConnect(cfg)
	if err != nil {
		log.Fatalf("Ошибка при подключении к БД: %v", err)
	}
	defer db.Close()

	// Test data
	testUser := &UserJS{
		Username: "testuser",
		Email:    "test@example.com",
		Pwd:      "hashed_password_here", // In real app, this should be hashed
	}

	// Create user
	err = db.CreateUser(context.Background(), testUser)
	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}
	fmt.Println("User created successfully")

	// Retrieve user by email
	retrievedUser, err := db.InfoUser(context.Background(), testUser.Email)
	if err != nil {
		log.Fatalf("Failed to retrieve user: %v", err)
	}

	fmt.Printf("Retrieved user - ID: %d, Username: %s, Email: %s, Password: %s\n",
		retrievedUser.ID, retrievedUser.Username, retrievedUser.Email, retrievedUser.Pwd)

	// Verify data matches
	if retrievedUser.Username != testUser.Username {
		fmt.Println("Warning: Username mismatch!")
	}
}
