package logging

import (
    "qr-generator-golang/internal/domain"
    "go.uber.org/zap"
)

type Logger struct {
    *zap.Logger
}

func NewLogger() domain.Logger {
    logger, _ := zap.NewProduction()
    return &Logger{Logger: logger}
}

func (l *Logger) Error(msg string, err error, fields ...zap.Field) {
    fields = append(fields, zap.Error(err))
    l.Logger.Error(msg, fields...)
}

func (l *Logger) Fatal(msg string, err error, fields ...zap.Field) {
    fields = append(fields, zap.Error(err))
    l.Logger.Fatal(msg, fields...)
}