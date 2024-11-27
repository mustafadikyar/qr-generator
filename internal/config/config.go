package config

import (
    "fmt"
    "github.com/boombuler/barcode/qr"
)

type Config struct {
    Width      int    `mapstructure:"width"`
    Height     int    `mapstructure:"height"`
    Content    string `mapstructure:"content"`
    ErrorLevel string `mapstructure:"errorLevel"`
    OutputDir  string `mapstructure:"outputDir"`
}

func (c *Config) GetErrorLevel() qr.ErrorCorrectionLevel {
    switch c.ErrorLevel {
    case "L":
        return qr.L
    case "M":
        return qr.M
    case "Q":
        return qr.Q
    case "H":
        return qr.H
    default:
        return qr.M
    }
}

func (c *Config) Validate() error {
    if c.Content == "" {
        return fmt.Errorf("content cannot be empty")
    }
    if c.Width <= 0 || c.Height <= 0 {
        return fmt.Errorf("invalid dimensions")
    }
    return nil
}