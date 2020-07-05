package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

// AppConfig ...
type AppConfig struct {
	Port                  string `mapstructure:"PORT" validate:"required"`
	DatabaseDriver        string `mapstructure:"DATABASE_DRIVER" validate:"required"`
	DatabaseConnectionURL string `mapstructure:"DATABASE_CONNECTION_URL" validate:"required"`
}

var validate *validator.Validate

// SetAppConfig ...
func SetAppConfig() (*AppConfig, error) {
	appConfig := AppConfig{}
	if err := viper.Unmarshal(&appConfig); err != nil {
		return nil, err
	}

	validate = validator.New()
	if err := validate.Struct(appConfig); err != nil {
		return nil, err
	}

	return &appConfig, nil
}
