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
	CreateUser(ctx context.Context, user *UserJS) error
	ReadUserByEmail(ctx context.Context, email string) (*UserJS, error)
	ReadUserById(ctx context.Context, id int32) (*UserJS, error)
	DeleteUser(ctx context.Context, user *UserJS) error
	UpdateUser(ctx context.Context, user *UserJS) error
	isUser(ctx context.Context, user *UserJS) bool

	AddLike(ctx context.Context, id_user, id_artist int32) error
	RemoveLike(ctx context.Context, id_user, id_artist int32) error
	Close()
}

func NewConnect(cfg *DBConfig) (*PostgresDB, error) {
	connStr := fmt.Sprintf(DBConnStringFormat, cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	loggerDB.Info("Configuration loaded from environment")

	db, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		loggerDB.Fatal("Connection pool not established", zap.Error(err))
		return nil, err
	}

	if err = db.Ping(context.Background()); err != nil {
		loggerDB.Fatal("Database connection failed", zap.Error(err))
		db.Close()
		return nil, err
	}

	if _, err = db.Exec(context.Background(), QueryDropTable()); err != nil {
		loggerDB.Error("Failed to drop table", zap.Error(err))
	} else {
		loggerDB.Info("Table dropped successfully")
	}

	if _, err = db.Exec(context.Background(), QueryCreateTable()); err != nil {
		loggerDB.Error("Failed to create table", zap.Error(err))
		return nil, fmt.Errorf("failed to create table: %w", err)
	}
	loggerDB.Info("User storage created")

	return &PostgresDB{pool: db}, nil
}

func (db *PostgresDB) Close() {
	loggerDB.Info("Closing database connection")
	db.pool.Close()
	loggerDB.Info("Database connection closed")
}

func (db *PostgresDB) CreateUser(ctx context.Context, user *UserJS) error {
	var err error
	user.Pwd, err = HashPassword(user.Pwd)
	if err != nil {
		loggerDB.Error("Failed to hash password",
			zap.String("username", user.Username),
			zap.String("email", user.Email),
			zap.Error(err))
		return fmt.Errorf("ошибка хеширования пароля: %w", err)
	}

	err = db.pool.QueryRow(ctx, QueryCreateUser(),
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
		loggerDB.Error("Failed to create user",
			zap.String("username", user.Username),
			zap.String("email", user.Email),
			zap.Error(err))
		return fmt.Errorf("ошибка создания пользователя: %w", err)
	}

	loggerDB.Info("User created successfully",
		zap.Int32("id", user.ID),
		zap.String("username", user.Username),
		zap.String("email", user.Email),
	)

	return nil
}

func (db *PostgresDB) ReadUserByEmail(ctx context.Context, email string) (*UserJS, error) {
	var user UserJS
	loggerDB.Debug("Reading user by email", zap.String("email", email))

	err := db.pool.QueryRow(ctx, QueryInfoUserByEmail(), email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Pwd,
		&user.CreatedAt,
	)
	if err != nil {
		loggerDB.Error("Failed to read user by email",
			zap.String("email", email),
			zap.Error(err))
		return nil, err
	}

	loggerDB.Debug("User found by email",
		zap.Int32("id", user.ID),
		zap.String("email", email))
	return &user, nil
}

func (db *PostgresDB) ReadUserById(ctx context.Context, id int32) (*UserJS, error) {
	var user UserJS
	loggerDB.Debug("Reading user by ID", zap.Int32("id", id))

	err := db.pool.QueryRow(ctx, QueryInfoUserById(), id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Pwd,
		&user.CreatedAt,
	)
	if err != nil {
		loggerDB.Error("Failed to read user by ID",
			zap.Int32("id", id),
			zap.Error(err))
		return nil, err
	}

	loggerDB.Debug("User found by ID",
		zap.Int32("id", user.ID),
		zap.String("username", user.Username))
	return &user, nil
}

func (db *PostgresDB) DeleteUser(ctx context.Context, user *UserJS) error {
	loggerDB.Info("Deleting user",
		zap.Int32("id", user.ID),
		zap.String("username", user.Username))

	_, err := db.pool.Exec(ctx, QueryDeleteUser(), user.ID)
	if err != nil {
		loggerDB.Error("Failed to delete user",
			zap.Int32("id", user.ID),
			zap.String("username", user.Username),
			zap.Error(err))
		return fmt.Errorf("ошибка удаления пользователя: %w", err)
	}

	loggerDB.Info("User deleted successfully",
		zap.Int32("id", user.ID),
		zap.String("username", user.Username))
	return nil
}

func (db *PostgresDB) isUser(ctx context.Context, user *UserJS) bool {
	loggerDB.Info("Checking if user exists",
		zap.Int32("id", user.ID),
		zap.String("username", user.Username),
		zap.String("email", user.Email))

	var exists bool
	err := db.pool.QueryRow(ctx, QueryIsUser(), user.ID, user.Username, user.Email).Scan(&exists)
	if err != nil {
		loggerDB.Error("Failed to check user existence",
			zap.Int32("id", user.ID),
			zap.String("username", user.Username),
			zap.String("email", user.Email),
			zap.Error(err))
		return false
	}

	if exists {
		loggerDB.Debug("User exists in database",
			zap.Int32("id", user.ID),
			zap.String("username", user.Username),
			zap.String("email", user.Email))
	} else {
		loggerDB.Debug("User does not exist in database",
			zap.Int32("id", user.ID),
			zap.String("username", user.Username),
			zap.String("email", user.Email))
	}

	return exists
}

func (db *PostgresDB) AddLike(ctx context.Context, id_user, id_artist int32) error {
	loggerDB.Debug("Adding like",
		zap.Int32("user_id", id_user),
		zap.Int32("artist_id", id_artist))

	_, err := db.pool.Exec(ctx, QueryAddLike(), id_user, id_artist)
	if err != nil {
		loggerDB.Error("Failed to add like",
			zap.Int32("user_id", id_user),
			zap.Int32("artist_id", id_artist),
			zap.Error(err))
		return fmt.Errorf("ошибка добавления лайка: %w", err)
	}

	loggerDB.Info("Like added successfully",
		zap.Int32("user_id", id_user),
		zap.Int32("artist_id", id_artist))
	return nil
}

func (db *PostgresDB) RemoveLike(ctx context.Context, id_user, id_artist int32) error {
	loggerDB.Debug("Removing like",
		zap.Int32("user_id", id_user),
		zap.Int32("artist_id", id_artist))

	_, err := db.pool.Exec(ctx, QueryRemoveLike(), id_user, id_artist)
	if err != nil {
		loggerDB.Error("Failed to remove like",
			zap.Int32("user_id", id_user),
			zap.Int32("artist_id", id_artist),
			zap.Error(err))
		return fmt.Errorf("ошибка удаления лайка: %w", err)
	}

	loggerDB.Info("Like removed successfully",
		zap.Int32("user_id", id_user),
		zap.Int32("artist_id", id_artist))
	return nil
}
