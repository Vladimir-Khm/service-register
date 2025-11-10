package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	LogLevel            string `mapstructure:"LOG_LEVEL"`
	PayloadSignatureKey string `mapstructure:"PAYLOAD_SIGNATURE_KEY"`

	DbHost               string `mapstructure:"DB_HOST"`
	DbPort               string `mapstructure:"DB_PORT"`
	DbUser               string `mapstructure:"DB_USER"`
	DbPassword           string `mapstructure:"DB_PASSWORD"`
	DbName               string `mapstructure:"DB_NAME"`
	DbMaxOpenConns       int    `mapstructure:"DB_MAX_OPEN_CONNS"`
	DbMaxIdleConns       int    `mapstructure:"DB_MAX_IDLE_CONNS"`
	DbConnMaxLifetimeSec int64  `mapstructure:"DB_CONN_MAX_LIFETIME_SEC"`
	DbConnMaxIdleTimeSec int64  `mapstructure:"DB_CONN_MAX_IDLE_TIME_SEC"`
}

func LoadConfig() (*Config, error) {
	config := &Config{}
	viper.SetConfigFile(".env")

	// ignore error to make .env config not mandatory
	viper.ReadInConfig()

	viper.AutomaticEnv()

	err := viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
