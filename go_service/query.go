package main

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
}

func QueryCreateUser() string {
	return `INSERT INTO users(
	username, 
	email,
	pwd,
	create_at
	)
	VALUES
	($1, $2, $3, NOW())
	RETURNING id, username, email, pwd, create_at
	`
}

func QueryDeleteUser() string {
	return `DELETE FROM users WHERE username = $1`
}

func QueryInfoUser() string {
	return `SELECT * FROM users WHERE email=$1`
}

func QueryDropTable() string{
	return `DROP TABLE if EXISTS users`
}