package main

import (
    "os"
    "qr-generator-golang/internal/config"
    "qr-generator-golang/internal/generator"
    "qr-generator-golang/internal/logging"
)

func main() {
    logger := logging.NewLogger()
    defer logger.Sync()

    configLoader := config.NewLoader(logger)
    if err := configLoader.Load(); err != nil {
        logger.Fatal("failed to load configuration", err)
        os.Exit(1)
    }

    gen := generator.NewQRGenerator(configLoader.GetConfig(), logger)
    if err := gen.Generate(); err != nil {
        logger.Fatal("failed to generate QR code", err)
        os.Exit(1)
    }
}