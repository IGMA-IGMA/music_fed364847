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
    CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);

    CREATE TABLE IF NOT EXISTS users_like (
        id SERIAL PRIMARY KEY,
        id_user INTEGER REFERENCES users(id) ON DELETE CASCADE,
        id_artist INTEGER,
        create_at TIMESTAMP DEFAULT NOW()
    );

    CREATE INDEX IF NOT EXISTS idx_users_like_id_user ON users_like(id_user);
    CREATE INDEX IF NOT EXISTS idx_users_like_id_artist ON users_like(id_artist);
    `
}

func QueryCreateUser() string {
	return `INSERT INTO users(
	username, email, pwd, create_at)
	VALUES
	($1, $2, $3, NOW())
	RETURNING id, username, email, pwd, create_at
	`
}

func QueryDeleteUser() string {
	return `DELETE FROM users WHERE username = $1`
}

func QueryInfoUserByEmail() string {
	return `SELECT * FROM users WHERE email=$1`
}

func QueryInfoUserByName() string {
	return `SELECT * FROM users WHERE username=$1`
}

func QueryInfoUserById() string {
	return `SELECT * FROM users WHERE id=$1`
}

func QueryDropTable() string {
	return `DROP TABLE IF EXISTS users`
}

func QueryIsUser() string{
    return "SELECT EXISTS (SELECT 1 FROM users WHERE id=$1 AND username=$2 AND email=$3)"
}

func QueryAddLike() string {
	return `INSERT INTO users_like(id_user, id_artist, create_at) VALUES ($1, $2, NOW())`
}

func QueryRemoveLike() string {
	return `DELETE FROM users_like WHERE id_user=$1 AND id_artist=$2`
}

func QueryAllLikeUser() string{
    return `SELECT id_artist FROM users_like WHERE id_user=$1`
}