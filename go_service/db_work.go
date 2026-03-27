package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type PostgresDB struct {
	pool *pgxpool.Pool
}

type StorageDB interface {
	CreateUser(ctx context.Context, user *UserJS) error                 // C
	ReadUserByEmail(ctx context.Context, user *UserJS) (*UserJS, error) // R
	ReadUserById(ctx context.Context, user *UserJS) (*UserJS, error)    // R
	DeleateUser(ctx context.Context, user *UserJS) error                // D
	UpdateUser(ctx context.Context, user *UserJS) error                 // U

	Close()
}

func NewConnect(cfg *DBConfig) (*PostgresDB, error) {
	connStr := fmt.Sprintf(DBConnStringFormat, cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	loggerDB.Info("Configuration loaded from environment")

	db, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		loggerDB.Fatal("Connection pool not established")
		return nil, err
	}

	if err = db.Ping(context.Background()); err != nil {
		loggerDB.Fatal("Database connection failed")
		db.Close()
		return nil, err
	}

	db.Exec(context.Background(), QueryDropTable())
	db.Exec(context.Background(), QueryCreateTable())
	loggerDB.Info("User storage created")

	return &PostgresDB{pool: db}, nil
}

func (db *PostgresDB) Close() {
	db.pool.Close()
}

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

func (db *PostgresDB) ReadUserByEmail(ctx context.Context, email string) (*UserJS, error) {
	var user UserJS
	err := db.pool.QueryRow(ctx, QueryInfoUserByEmail(), email).Scan(&user.ID,
		&user.Username,
		&user.Email,
		&user.Pwd,
		&user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}


