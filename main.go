package main

import (
	"fmt"

	"github.com/DelusionalOptimist/pastebean/models"
	"github.com/DelusionalOptimist/pastebean/router"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := initDBConnection()
	if err != nil {
		panic(err)
	}
	router := router.NewRouter(db)

	router.Run()
}

func initDBConnection() (*gorm.DB, error) {
	dbConfig, err := models.LoadDBConfig(".env")
	if err != nil {
		return nil, err
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=%s", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName, "Asia/Kolkata")

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
  if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Paste{})

	return db, nil
}
