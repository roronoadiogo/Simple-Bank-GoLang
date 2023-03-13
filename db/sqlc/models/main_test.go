package models

import (
	"database/sql"
	"log"
	"testing"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/bank_account?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {

		log.Fatal("Cannot connect to db:", err)
	}

	testQueries = New(conn)

}
