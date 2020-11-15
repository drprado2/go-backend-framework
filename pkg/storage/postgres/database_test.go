package postgres

import (
	"github.com/drprado2/go-backend-framework/pkg/storage"
	"github.com/drprado2/go-backend-framework/pkg/tests/testutilities"
	"testing"
)

var connectionString string
var connectionWithoutDb string
var databaseName string
var transaction storage.TransactionInterface
var connection storage.ConnectionInterface
var database storage.DatabaseInterface

func setup(t *testing.T){
	dbName, conn, connWithoutDb, err := testutilities.CreateRandomDB()
	if err != nil {
		t.Error("Error in setup", err)
	}
	connectionString = conn
	connectionWithoutDb = connWithoutDb
	databaseName = dbName
}

func teardown(t *testing.T){
	factory := NewDatabaseFactory(connectionWithoutDb)
	db, err := factory.GetDB()
	defer db.Close()
	if err != nil {
		t.Error("Error on terardown", err)
	}
	_, err = db.Exec(`drop database "` + databaseName + `" WITH (FORCE);`)
	if err != nil {
		t.Error("Error on terardown", err)
	}
}

func TestDatabaseWorks(t *testing.T){
	setup(t)
	defer teardown(t)

	factory := NewDatabaseFactory(connectionString)
	fullDb, err := factory.GetDB()
	defer fullDb.Close()
	if err != nil {
		t.Error(err)
	}

	database = fullDb

	if _, err := database.Exec("create table testTable (id int primary key, name varchar)"); err != nil {
		t.Error("Error creating table", err)
	}
	var id int
	err = database.QueryRow("insert into testTable values (1, 'adriano') RETURNING id").Scan(&id)
	if id != 1 || err != nil {
		t.Error("Error inserting on table", id, err)
	}
}


