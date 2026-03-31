package main

<<<<<<< HEAD
import (
	"net"
	"time"

	"go.uber.org/zap"
=======
<<<<<<< Updated upstream
func main() {
	gensintjson(10000)
=======
import (
	"fmt"
>>>>>>> 8db0ba0dbecf88af941a3ee9bff345731e3e4735
)

func init() {
	initLog()
	initDB()
<<<<<<< HEAD
=======
}

func main() {
	// listener, err := net.Listen("tcp", "127.0.0.1:50051")
	// if err != nil {
	// 	loggerGRPC.Fatal("Failed to listen", zap.String("error", err.Error()))
	// }

	// grpcServer := NewGRPCServer()

	// loggerGRPC.Info("🚀 gRPC server is running", zap.String("port", ":50051"))

	// if err := grpcServer.Serve(listener); err != nil {
	// 	loggerGRPC.Fatal("Failed to serve", zap.String("error", err.Error()))
	// }
>>>>>>> Stashed changes
>>>>>>> 8db0ba0dbecf88af941a3ee9bff345731e3e4735
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