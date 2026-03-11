package main

import (
	"fmt"
	"log"

	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.Logger
}

func NewLogger(filename string) (*Logger, error) {
	file, _ := createFile(path_log_dir, filename)

	config := zap.NewProductionConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONE(config)
	writer := zapcore.AddSunc(file)
	defaultLogLevel := zapcore.DebugLevel
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
	)
	logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	
}
