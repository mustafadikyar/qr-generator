package generator

import (
    "fmt"
    "image/png"
    "os"
    "path/filepath"
    
    "qr-generator-golang/internal/config"
    "qr-generator-golang/internal/domain"
    "github.com/boombuler/barcode"
    "github.com/boombuler/barcode/qr"
    "github.com/google/uuid"
    "go.uber.org/zap"
)

type QRGenerator struct {
    config *config.Config
    logger domain.Logger
}

func NewQRGenerator(config *config.Config, logger domain.Logger) domain.QRGenerator {
    return &QRGenerator{
        config: config,
        logger: logger,
    }
}

func (g *QRGenerator) Generate() error {
    g.logger.Info("starting QR code generation", 
        zap.String("content", g.config.Content),
        zap.Int("width", g.config.Width))

    if err := g.config.Validate(); err != nil {
        return fmt.Errorf("validation error: %w", err)
    }

    qrCode, err := g.createQRCode()
    if err != nil {
        return err
    }

    return g.saveToFile(qrCode)
}

func (g *QRGenerator) createQRCode() (barcode.Barcode, error) {
    qrCode, err := qr.Encode(g.config.Content, g.config.GetErrorLevel(), qr.Auto)
    if err != nil {
        return nil, fmt.Errorf("QR encoding error: %w", err)
    }

    qrCode, err = barcode.Scale(qrCode, g.config.Width, g.config.Height)
    if err != nil {
        return nil, fmt.Errorf("QR scaling error: %w", err)
    }

    return qrCode, nil
}

func (g *QRGenerator) saveToFile(qrCode barcode.Barcode) error {
    if err := os.MkdirAll(g.config.OutputDir, 0755); err != nil {
        return fmt.Errorf("directory creation error: %w", err)
    }

    filename := fmt.Sprintf("%s.png", uuid.New().String())
    outputPath := filepath.Join(g.config.OutputDir, filename)

    file, err := os.Create(outputPath)
    if err != nil {
        return fmt.Errorf("file creation error: %w", err)
    }
    defer file.Close()

    if err := png.Encode(file, qrCode); err != nil {
        return fmt.Errorf("PNG encoding error: %w", err)
    }

    g.logger.Info("QR code generated successfully", zap.String("path", outputPath))
    return nil
}