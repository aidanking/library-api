package main

import (
	"github.com/aidanking/library-api/api"
	"github.com/aidanking/library-api/storage"
)

func main() {

	db := storage.Connect()
	defer db.Close()

	server := api.ApiServer{
		ListenAddr: ":8081",
		DB:         db,
	}

	server.Start()
}
