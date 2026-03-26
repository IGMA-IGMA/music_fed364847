package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type PostgresDB struct {
	pool *pgxpool.Pool
}

<<<<<<< Updated upstream
type Storage interface {
	CreateUser(ctx context.Context, user *UserJS)
	DeleateUser(ctx context.Context, user *UserJS)
	UpdateUser(ctx context.Context, user *UserJS)
	InfoUser(ctx context.Context, user *UserJS)
=======
type StorageDB interface {
	CreateUser(ctx context.Context, user *UserJS) error                 // C
	ReadUserByEmail(ctx context.Context, user *UserJS) (*UserJS, error) // R
	ReadUserById(ctx context.Context, user *UserJS) (*UserJS, error)    // R
	DeleateUser(ctx context.Context, user *UserJS) error                // D
	UpdateUser(ctx context.Context, user *UserJS) error                 // U

>>>>>>> Stashed changes
	Close()
}

func ParsConfig() (*DBConfig, error) {
	if err := godotenv.Load(path_env); err != nil {
		log.Println("Note: .env file not found, using system env")
	}

	data, err := os.ReadFile(path_db_config)
	if err != nil {
		return nil, err
	}

	replacedData := os.ExpandEnv(string(data))

	config := &DBConfig{}
	err = yaml.Unmarshal([]byte(replacedData), config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func NewConnect(cfg *DBConfig) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf(path_conn_db, cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := pgxpool.New(context.Background(), connStr)
	CheckError(err)
	if err = db.Ping(context.Background()); err != nil {
		db.Close()
		return nil, err
	}
	db.Exec(context.Background(), CreateTable())
	return db, nil
}

func (db *PostgresDB) Close() {
	db.pool.Close()
}
<<<<<<< Updated upstream
=======

func (db *PostgresDB) CreateUser(ctx context.Context, user *UserJS) error {
	user.Pwd, _ = HashPassword(user.Pwd)
	err := db.pool.QueryRow(ctx, QueryCreateUser(),
		user.Username,
		user.Email,
		user.Pwd,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Pwd,
		&user.CreatedAt,
	)
	if err != nil {
		loggerDB.Error("Create",
			zap.String("username", user.Username),
			zap.String("email", user.Email),
		)
		return fmt.Errorf("ошибка создания пользователя: %w", err)
	}

	loggerDB.Info("Create",
		zap.String("username", user.Username),
		zap.String("email", user.Email),
	)

	return nil
}

func (db *PostgresDB) InfoUser(ctx context.Context, email string) (*UserJS, error) {
	var user UserJS
	err := db.pool.QueryRow(ctx, QueryInfoUser(), email).Scan(&user.ID,
		&user.Username,
		&user.Email,
		&user.Pwd,
		&user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
>>>>>>> Stashed changes
