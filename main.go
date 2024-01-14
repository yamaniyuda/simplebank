package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/yamaniyuda/simplebank/api"
	db "github.com/yamaniyuda/simplebank/db/sqlc"
	"github.com/yamaniyuda/simplebank/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(&store)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
