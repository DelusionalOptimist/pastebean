package models

import (
	"os"

	"github.com/spf13/viper"
)

type PostgresConfig struct {
	Host string `mapstructure:"DB_HOST"`
	Port int `mapstructure:"DB_PORT"`
	User string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	DBName string `mapstructure:"DB_NAME"`
}

func LoadDBConfig(configName string) (config *PostgresConfig, err error) {
	currDir, err := os.Getwd()
	if err != nil {
		return
	}
	viper.AddConfigPath(currDir)
	viper.SetConfigName(configName)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	return
}
