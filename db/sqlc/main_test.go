package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:7912@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	//config, err := util.LoadConfig("../..")
	//if err != nil {
	//	log.Fatal("cannot load conf", err)
	//}

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connct to db", err)
	}
	testQueries = New(conn)
	os.Exit(m.Run())
}
