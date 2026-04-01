package main

import (
	"net"
	"time"

	"go.uber.org/zap"
)

func init() {
	initLog()
	initDB()
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		loggerGRPC.Fatal("Failed to create listener", zap.Error(err))
	}
	defer listener.Close()

	jwtManager := NewJWTManager("IGMA", 24*time.Hour, "IGMA", "User")
	grpcServer := NewGRPCServer(db, jwtManager)

	loggerGRPC.Info("gRPC server started",
		zap.String("address", "127.0.0.1:50051"))

	if err := grpcServer.Serve(listener); err != nil {
		loggerGRPC.Fatal("Failed to start gRPC server", zap.Error(err))
	}
}
