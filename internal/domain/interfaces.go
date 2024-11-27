package domain

import "go.uber.org/zap"

type QRGenerator interface {
    Generate() error
}

type ConfigLoader interface {
    Load() error
}

type Logger interface {
    Info(msg string, fields ...zap.Field)
    Error(msg string, err error, fields ...zap.Field)
    Fatal(msg string, err error, fields ...zap.Field)
    Sync() error
}
