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

type Storage interface {
	CreateUser(ctx context.Context, user *UserJS)
	DeleateUser(ctx context.Context, user *UserJS)
	UpdateUser(ctx context.Context, user *UserJS)
	InfoUser(ctx context.Context, user *UserJS)
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
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := pgxpool.New(context.Background(), connStr)
	CheckError(err)
	if err = db.Ping(context.Background()); err != nil {
		db.Close()
		return nil, err
	}
	db.Exec(context.Background(), CreateTable())
	fmt.Println(1)
	return db, nil
}

func (db *PostgresDB) Close() {
	db.pool.Close()
}
