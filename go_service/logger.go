package main

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	logger *zap.Logger
}

func NewLogger(filename string) (*zap.Logger, error) {
	file, err := createFile(LogDirPath, filename)
	if err != nil {
		return nil, fmt.Errorf("failed to create log file: %w", err)
	}

	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.LineEnding = ",\n"
	fileEncoder := zapcore.NewJSONEncoder(config.EncoderConfig)
	writer := zapcore.AddSync(file)
	defaultLogLevel := zapcore.DebugLevel

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),

	)

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	return logger, nil
}
