package main

import (
	"backend-master/api"
	"backend-master/db/util"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/techschool/simplebank/api"
	db "github.com/techschool/simplebank/db/sqlc"
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

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
