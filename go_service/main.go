package main

import (
	"net"

	"go.uber.org/zap"
)

func init() {
	initLog()
	initDB()
}


func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		loggerGRPC.Fatal("Failed to listen", zap.String("error", err.Error()))
	}

	grpcServer := NewGRPCServer()

	loggerGRPC.Info("🚀 gRPC server is running", zap.String("port", ":50051"))

	if err := grpcServer.Serve(listener); err != nil {
		loggerGRPC.Fatal("Failed to serve", zap.String("error", err.Error()))
	}
}
