package main

var (
	loggerDB   *Logger
	loggerGRPC *Logger
	loggerFileManage *Logger
	db         *PostgresDB
)

func initLog() {
    loggerFileManage = NewLogger("file_manage.json")
    loggerDB = NewLogger("db_info.json")
    loggerGRPC = NewLogger("grpc_server.json")
}
	


func initDB() {
	cfg, _ := ParsConfig()
	db, _ = NewConnect(cfg)
}
