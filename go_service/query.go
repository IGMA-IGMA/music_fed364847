package main

func CreateTable() string {
	return `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(50) NOT NULL,
		email TEXT NOT NULL UNIQUE,
		pwd TEXT NOT NULL,
		create_at TIMESTAMP DEFAULT NOW()
	);`
}

func CreateUser() string {
	return `INSERT INTO User(
	username, 
	email,
	pwd,
	create_at
	)
	VALUES
	($1, $2, $3, NOW())
	`
}

func DeleteUser() string {
	return `DELETE FROM User WHERE username = $1`
}
