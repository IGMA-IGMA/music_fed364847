package main

import "go.uber.org/zap"

var (
	loggerDB         *Logger
	loggerGRPC       *Logger
	loggerFileManage *Logger
	db               *PostgresDB
)

func initLog() {
	loggerFileManage = NewLogger("file_manage.json")
	loggerDB = NewLogger("db_info.json")
	loggerGRPC = NewLogger("grpc_server.json")
}

func initDB() {
	cfg, err := ParsConfig()
	if err != nil {
		loggerDB.Fatal("Failed to parse config", zap.Error(err))
	}

	db, err = NewConnect(cfg)
	if err != nil {
		loggerDB.Fatal("Failed to connect to database", zap.Error(err))
	}

	loggerDB.Info("Database connected successfully")
}