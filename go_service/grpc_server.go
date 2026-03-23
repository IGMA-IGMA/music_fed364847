package main

import (
    "context"
    
    pb "go_service/gen/protos"
    "go.uber.org/zap"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
)

// userServer реализует gRPC интерфейс
type userServer struct {
    pb.UnimplementedApiUsersServer
}

// CreateUser — обработчик создания пользователя
func (s *userServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.StatusCode, error) {
    loggerGRPC.Info("CreateUser called",
        zap.String("username", req.GetUserName()),
        zap.String("email", req.GetUserEmail()),
    )

    // Валидация
    if req.GetUserName() == "" || req.GetUserEmail() == "" {
        loggerGRPC.Error("Validation failed: username/email is empty")
        return &pb.StatusCode{Code: 400}, nil
    }

    // Подготовка данных
    user := &UserJS{
        Username: req.GetUserName(),
        Email:    req.GetUserEmail(),
        Pwd:      req.GetUserPassword(),
    }

    // Сохранение в БД
    err := db.CreateUser(context.Background(), user)
    if err != nil {
        loggerGRPC.Error("Failed to create user in database",
            zap.String("error", err.Error()),
            zap.String("user", req.GetUserName()))
        return &pb.StatusCode{Code: 500}, err  // ← исправил: 500, а не 400
    }

    loggerGRPC.Info("User created successfully",
        zap.Int("id", user.ID),
        zap.String("username", user.Username))

    return &pb.StatusCode{Code: 201}, nil  // ← исправил: 201 Created
}

// NewGRPCServer создает и настраивает gRPC сервер
func NewGRPCServer() *grpc.Server {
    grpcServer := grpc.NewServer()
    
    // Регистрируем наш сервис
    pb.RegisterApiUsersServer(grpcServer, &userServer{})
    
    // Включаем рефлексию для отладки
    reflection.Register(grpcServer)
    
    return grpcServer
}