package main

import (
	"github.com/DelusionalOptimist/pastebean/delete/router"
	"github.com/DelusionalOptimist/pastebean/pkg/common/utils"
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
