package main

func CreateTable() string {
	return `CREATE TABLE User (
		id SERIAL PRIMARY KEY
		username VARCHAR(50) NOT NULL 
		email text NOT NULL UNIQUE
		pwd text NOT NULL
		create_at DEFAULT NOW()
	)`
}
