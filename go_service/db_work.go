package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresDB struct {
	pool *pgxpool.Pool
}

type Storage interface {
	CreateUser(ctx context.Context, user *UserJS) //C
	GetUser(ctx context.Context, user *UserJS) //R
	DeleateUser(ctx context.Context, user *UserJS) //D
	UpdateUser(ctx context.Context, user *UserJS) //U
	
	Close()
}

func NewConnect(cfg *DBConfig) (*PostgresDB, error) {
	connStr := fmt.Sprintf(DBConnStringFormat, cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(context.Background()); err != nil {
		db.Close()
		return nil, err
	}
	db.Exec(context.Background(), CreateTable())
	return &PostgresDB{pool: db}, nil
}

func (db *PostgresDB) Close() {
	db.pool.Close()
}

func (db *PostgresDB) CreateUser(ctx context.Context, user *UserJS) error {

	err := db.pool.QueryRow(ctx, CreateUser(),
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
		return fmt.Errorf("ошибка создания пользователя: %w", err)
	}

	return nil
}
