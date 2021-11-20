package main

import (
	"fmt"
	"time"

	"github.com/DelusionalOptimist/pastebean/create/router"
	"github.com/DelusionalOptimist/pastebean/pkg/common/models"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	time.Sleep(time.Second * 30)
	db, err := initDBConnection()
	if err != nil {
		panic(err)
	}
	router := router.NewRouter(db)

	router.Run()
}

func initDBConnection() (*gorm.DB, error) {
	viper.SetDefault("DB_DRIVER", "postgres")
	viper.SetDefault("DB_SOURCE", "postgres://postgres:password@postgres:5432/pastebeandb?sslmode=disable")

	viper.AutomaticEnv()

	connStr := fmt.Sprintf("%s", viper.GetString("DB_SOURCE"))

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
  if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Paste{})

	return db, nil
}
