package main

<<<<<<< Updated upstream
func CreateTable() string {
	return `CREATE TABLE User (
		id SERIAL PRIMARY KEY
		username VARCHAR(50) NOT NULL 
		email text NOT NULL UNIQUE
		pwd text NOT NULL
		create_at DEFAULT NOW()
	)`
=======
func QueryCreateTable() string {
    return `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        username VARCHAR(50) NOT NULL,
        email TEXT NOT NULL UNIQUE,
        pwd TEXT NOT NULL,
        create_at TIMESTAMP DEFAULT NOW()
    );
    
    CREATE INDEX IF NOT EXISTS idx_users_username ON users(username);
    CREATE INDEX IF NOT EXISTS idx_users_created_at ON users(create_at);
	CREATE INDEX IF NOT EXISTS idx_users_email on users(email);
    `
>>>>>>> Stashed changes
}
