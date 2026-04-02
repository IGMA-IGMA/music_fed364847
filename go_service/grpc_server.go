package main

import (
	"context"

	pb "go_service/gen/serverGRPC"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type userServer struct {
	pb.UnimplementedApiUsersServer
	jwtManager *JWTManager
	db         *PostgresDB
}

func (s *userServer) CreateUser(ctx context.Context, req *pb.UserInput) (*pb.StatusCode, error) {
	loggerGRPC.Info("Processing user registration",
		zap.String("username", req.GetUserName()),
		zap.String("email", req.GetUserEmail()),
	)

	if req.GetUserName() == "" || req.GetUserEmail() == "" {
		loggerGRPC.Error("Registration validation failed: missing required fields",
			zap.String("username", req.GetUserName()),
			zap.String("email", req.GetUserEmail()))
		return &pb.StatusCode{Code: 400}, nil
	}

	user := &UserJS{
		Username: req.GetUserName(),
		Email:    req.GetUserEmail(),
		Pwd:      req.GetUserPassword(),
	}

	err := s.db.CreateUser(context.Background(), user)
	if err != nil {
		loggerGRPC.Error("User registration failed",
			zap.String("username", req.GetUserName()),
			zap.String("email", req.GetUserEmail()),
			zap.Error(err))
		return &pb.StatusCode{Code: 500}, nil
	}

	loggerGRPC.Info("User registered successfully",
		zap.Int32("user_id", user.ID),
		zap.String("username", user.Username),
		zap.String("email", user.Email))

	return &pb.StatusCode{Code: 201}, nil
}

func (s *userServer) Login(ctx context.Context, req *pb.UserInput) (*pb.LoginResponse, error) {
	loggerGRPC.Info("Processing login attempt",
		zap.String("username", req.GetUserName()),
		zap.String("email", req.GetUserEmail()))

	if req.GetUserName() == "" || req.GetUserEmail() == "" {
		loggerGRPC.Error("Login validation failed: missing credentials",
			zap.String("username", req.GetUserName()),
			zap.String("email", req.GetUserEmail()))
		return &pb.LoginResponse{
			Status:  &pb.StatusCode{Code: 400},
			Token:   "",
			Message: "Username and email are required",
		}, nil
	}

	userDB, err := s.db.ReadUserByEmail(context.Background(), req.GetUserEmail())
	if err != nil {
		loggerGRPC.Error("Login failed: user not found",
			zap.String("email", req.GetUserEmail()),
			zap.Error(err))
		return &pb.LoginResponse{
			Status:  &pb.StatusCode{Code: 401},
			Token:   "",
			Message: "Invalid credentials",
		}, nil
	}

	if !CheckPassword(req.GetUserPassword(), userDB.Pwd) {
		loggerGRPC.Error("Login failed: invalid password",
			zap.Int32("user_id", userDB.ID),
			zap.String("username", userDB.Username))
		return &pb.LoginResponse{
			Status:  &pb.StatusCode{Code: 401},
			Token:   "",
			Message: "Invalid credentials",
		}, nil
	}

	token, err := s.jwtManager.GenerateToken(userDB, "user")
	if err != nil {
		loggerGRPC.Error("Login failed: token generation error",
			zap.Int32("user_id", userDB.ID),
			zap.String("username", userDB.Username),
			zap.Error(err))
		return &pb.LoginResponse{
			Status:  &pb.StatusCode{Code: 500},
			Token:   "",
			Message: "Internal server error",
		}, nil
	}

	loggerGRPC.Info("User logged in successfully",
		zap.Int32("user_id", userDB.ID),
		zap.String("username", userDB.Username),
		zap.String("email", userDB.Email))

	return &pb.LoginResponse{
		Status:  &pb.StatusCode{Code: 200},
		Token:   token,
		Message: "Login successful",
	}, nil
}

func (s *userServer) AddLike(ctx context.Context, req *pb.LikeRequest) (*pb.StatusCode, error) {
	loggerGRPC.Info("Processing AddLike request",
		zap.Int32("artist_id", req.GetArtistId()),
	)

	token := req.GetToken()
	if token == nil || token.GetToken() == "" {
		loggerGRPC.Error("AddLike failed: token is missing")
		return &pb.StatusCode{Code: 401}, nil
	}

	accessToken := token.GetToken()

	claims, err := s.jwtManager.ValidateToken(accessToken)
	if err != nil {
		loggerGRPC.Error("AddLike failed: invalid token",
			zap.Error(err))
		return &pb.StatusCode{Code: 401}, nil
	}

	loggerGRPC.Debug("Token validated successfully",
		zap.String("username", claims.Username),
		zap.String("email", claims.Email))

	user := &UserJS{
		Username: claims.Username,
		Email:    claims.Email,
	}

	if !s.db.isUser(ctx, user) {
		loggerGRPC.Error("AddLike failed: user not found in database",
			zap.String("username", claims.Username),
			zap.String("email", claims.Email))
		return &pb.StatusCode{Code: 404}, nil
	}

	err = s.db.AddLike(ctx, user.ID, req.GetArtistId())
	if err != nil {
		loggerGRPC.Error("AddLike failed: database error",
			zap.Int32("user_id", user.ID),
			zap.Int32("artist_id", req.GetArtistId()),
			zap.Error(err))
		return &pb.StatusCode{Code: 500}, nil
	}

	loggerGRPC.Info("Like added successfully",
		zap.Int32("user_id", user.ID),
		zap.String("username", user.Username),
		zap.Int32("artist_id", req.GetArtistId()))

	return &pb.StatusCode{Code: 200}, nil
}

func (s *userServer) RemoveLike(ctx context.Context, req *pb.LikeRequest) (*pb.StatusCode, error) {
	loggerGRPC.Info("Processing RemoveLike request",
		zap.Int32("artist_id", req.GetArtistId()),
	)

	token := req.GetToken()
	if token == nil || token.GetToken() == "" {
		loggerGRPC.Error("RemoveLike failed: token is missing")
		return &pb.StatusCode{Code: 401}, nil
	}

	accessToken := token.GetToken()

	claims, err := s.jwtManager.ValidateToken(accessToken)
	if err != nil {
		loggerGRPC.Error("RemoveLike failed: invalid token",
			zap.Error(err))
		return &pb.StatusCode{Code: 401}, nil
	}

	loggerGRPC.Debug("Token validated successfully",
		zap.String("username", claims.Username),
		zap.String("email", claims.Email))

	user := &UserJS{
		Username: claims.Username,
		Email:    claims.Email,
	}

	if !s.db.isUser(ctx, user) {
		loggerGRPC.Error("RemoveLike failed: user not found in database",
			zap.String("username", claims.Username),
			zap.String("email", claims.Email))
		return &pb.StatusCode{Code: 404}, nil
	}

	err = s.db.RemoveLike(ctx, user.ID, req.GetArtistId())
	if err != nil {
		loggerGRPC.Error("RemoveLike failed: database error",
			zap.Int32("user_id", user.ID),
			zap.Int32("artist_id", req.GetArtistId()),
			zap.Error(err))
		return &pb.StatusCode{Code: 500}, nil
	}

	loggerGRPC.Info("RemoveLike successful",
		zap.Int32("user_id", user.ID),
		zap.String("username", user.Username),
		zap.Int32("artist_id", req.GetArtistId()))

	return &pb.StatusCode{Code: 200}, nil
}

func NewGRPCServer(db *PostgresDB, jwtManager *JWTManager) *grpc.Server {
	grpcServer := grpc.NewServer()

	pb.RegisterApiUsersServer(grpcServer, &userServer{
		db:         db,
		jwtManager: jwtManager,
	})

	reflection.Register(grpcServer)

	return grpcServer
}
