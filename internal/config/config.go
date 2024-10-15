package config

import "github.com/spf13/viper"

type Config struct {
	ServiceName    string `mapstructure:"SERVICE_NAME"`
	ServiceVersion string `mapstructure:"SERVICE_VERSION"`
	ServerPort     int    `mapstructure:"SERVER_PORT"`
	DefaultTimeout int    `mapstructure:"DEFAULT_TIMEOUT"`
	DatabaseURL    string `mapstructure:"DATABASE_URL"`
	DatabaseName   string `mapstructure:"DATABASE_NAME"`
}

const configType = "env"

func New() (*Config, error) {
	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigType(configType)
	viper.SetDefault("SERVICE_NAME", "")
	viper.SetDefault("SERVICE_VERSION", "0.1.0")
	viper.SetDefault("SERVER_PORT", 8080)
	viper.SetDefault("DEFAULT_TIMEOUT", 10)
	viper.SetDefault("DATABASE_URL", "")
	viper.SetDefault("DATABASE_NAME", "")

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
