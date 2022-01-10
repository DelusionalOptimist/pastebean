package main

import (
	"github.com/DelusionalOptimist/pastebean/pkg/common/utils"
	"github.com/DelusionalOptimist/pastebean/update/router"
	_ "github.com/lib/pq"
)

func main() {
	db, err := utils.InitDBConnection()
	if err != nil {
		panic(err)
	}
	router := router.NewRouter(db)

	router.Run()
}
