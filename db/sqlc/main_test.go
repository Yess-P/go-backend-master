package db

import (
	"backend-master/db/util"
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
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load conf", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	//var err error

	//testDB, err = sql.Open(dbDriver, dbSource)
	//if err != nil {
	//	log.Fatal("cannot connct to db", err)
	//}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
