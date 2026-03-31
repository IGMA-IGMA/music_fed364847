package main

<<<<<<< Updated upstream
func main() {
	gensintjson(10000)
=======
import (
	"fmt"
)

func init() {
	initLog()
	initDB()
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