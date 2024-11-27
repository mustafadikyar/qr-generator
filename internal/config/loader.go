package config

import (
    "github.com/spf13/viper"
    "qr-generator-golang/internal/domain"
)

type Loader struct {
    logger domain.Logger
    config *Config
}

func NewLoader(logger domain.Logger) *Loader {
    return &Loader{
        logger: logger,
        config: &Config{},
    }
}

func (l *Loader) GetConfig() *Config {
    return l.config
}

func (l *Loader) Load() error {
    l.logger.Info("loading configuration")
    
    viper.SetConfigName("config")
    viper.AddConfigPath("config")
    viper.SetConfigType("yml")

    if err := viper.ReadInConfig(); err != nil {
        return err
    }

    if err := viper.Unmarshal(l.config); err != nil {
        return err
    }

    return nil
}