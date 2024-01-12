package main

import (
	"backend-master/api"
	db "backend-master/db/sqlc"
	"backend-master/db/util"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

//const (
//	dbDriver      = "postgres"
//	dbSource      = "postgresql://root:7912@localhost:5432/simple_bank?sslmode=disable"
//	serverAddress = "0.0.0.0:8080"
//)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
