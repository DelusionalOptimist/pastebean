package utils

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDBConnection() (*gorm.DB, error) {
	viper.SetDefault("DB_DRIVER", "postgres")
	viper.SetDefault("DB_SOURCE", "postgres://postgres:password@postgres:5432/pastebeandb?sslmode=disable")

	viper.AutomaticEnv()

	connStr := fmt.Sprintf("%s", viper.GetString("DB_SOURCE"))

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	// try waiting for db to get alive
	var retryCount = 1
	if strings.Contains(err.Error(), "connection refused") && retryCount < 5 {
		log.Printf("Failed connecting to database. Retrying in %d seconds.", 2*retryCount)
		time.Sleep(time.Second * time.Duration(2*retryCount))
		retryCount++
	} else if err != nil {
		return nil, err
	}

	return db, nil
}
