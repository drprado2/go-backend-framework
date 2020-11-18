package postgres

import (
	"context"
	"fmt"
	"github.com/drprado2/go-backend-framework/pkg/storage"
	"github.com/drprado2/go-backend-framework/pkg/tests/testutilities"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"testing"
)

var databaseName string
var transaction storage.TransactionInterface
var connection storage.ConnectionInterface
var fullDB storage.FullDatabaseInterface
var DB storage.DatabaseInterface
var connectionWithoutDB storage.FullDatabaseInterface

func setup(t *testing.T) {
	connWithoutDB, connWithDB, dbName, err := testutilities.CreateRandomDB()

	connectionWithoutDB = &Database{
		DB: *connWithoutDB,
	}
	fullDB = &Database{
		DB: *connWithDB,
	}
	DB = connWithDB

	if err != nil {
		t.Fatal("Error in setup", err)
	}
	databaseName = dbName
}

func teardown(t *testing.T) {
	defer fullDB.Close()
	defer connectionWithoutDB.Close()
	_, err := connectionWithoutDB.Exec(`drop database "` + databaseName + `" WITH (FORCE);`)
	if err != nil {
		t.Error("Error on terardown", err)
	}
}

func TestDatabaseExecution(t *testing.T) {
	setup(t)
	defer teardown(t)

	if _, err := DB.Exec(`create table testTable (id int primary key, name varchar)`); err != nil {
		t.Fatal("Error creating table", err)
	}
	var id int
	if err := DB.QueryRow(`insert into testTable values (1, 'adriano') RETURNING id`).Scan(&id); id != 1 || err != nil {
		t.Fatal("Error inserting on table", id, err)
	}
}

func TestTransactionRollback(t *testing.T) {
	setup(t)
	defer teardown(t)

	tableName, _ := uuid.NewUUID()
	if row := DB.QueryRow(fmt.Sprintf(`create table "%s" (id int primary key, name varchar)`, tableName.String())); row.Err() != nil {
		t.Fatal("Error creating table", row.Err())
	}

	tx, err := fullDB.Begin()
	transaction = tx
	if err != nil {
		t.Fatalf("Fail creating transaction\n%s", err)
	}
	var id int
	err = transaction.QueryRow(fmt.Sprintf(`insert into "%s" values (1, 'adriano') RETURNING id`, tableName.String())).Scan(&id)
	if id != 1 || err != nil {
		t.Fatal("Error inserting on table", id, err)
	}
	countQuery := fmt.Sprintf(`select count(*) from "%s"`, tableName.String())

	var countInTx int
	transaction.QueryRow(countQuery).Scan(&countInTx)
	if countInTx != 1 {
		t.Fatalf("Rows count in transaction should be 1 got %v", countInTx)
	}

	err = transaction.Rollback()

	if err != nil {
		t.Fatalf("Fail rollback\nError: %s", err)
	}

	var countAfterRollback int
	DB.QueryRow(countQuery).Scan(&countAfterRollback)
	if countAfterRollback != 0 {
		t.Fatalf("Rows count in transaction after rollback should be 0 got %v", countAfterRollback)
	}
}

func TestTransactionCommit(t *testing.T) {
	setup(t)
	defer teardown(t)

	tableName, _ := uuid.NewUUID()
	if row := DB.QueryRow(fmt.Sprintf(`create table "%s" (id int primary key, name varchar)`, tableName.String())); row.Err() != nil {
		t.Fatal("Error creating table", row.Err())
	}

	tx, err := fullDB.Begin()
	transaction = tx
	if err != nil {
		t.Fatalf("Fail creating transaction\n%s", err)
	}
	var id int
	err = transaction.QueryRow(fmt.Sprintf(`insert into "%s" values (1, 'adriano') RETURNING id`, tableName.String())).Scan(&id)
	if id != 1 || err != nil {
		t.Fatal("Error inserting on table", id, err)
	}
	countQuery := fmt.Sprintf(`select count(*) from "%s"`, tableName.String())

	var countInTx int
	transaction.QueryRow(countQuery).Scan(&countInTx)
	if countInTx != 1 {
		t.Fatalf("Rows count in transaction should be 1 got %v", countInTx)
	}

	err = transaction.Commit()

	if err != nil {
		t.Fatalf("Fail rollback\nError: %s", err)
	}

	var countAfterRollback int
	DB.QueryRow(countQuery).Scan(&countAfterRollback)
	if countAfterRollback != 1 {
		t.Fatalf("Rows count in transaction after commit should be 1 got %v", countAfterRollback)
	}
}

func TestDatabaseConnectionExecution(t *testing.T) {
	setup(t)
	defer teardown(t)

	ctx := context.Background()
	conn, err := fullDB.Conn(ctx)
	if err != nil {
		t.Fatal("Error creating connection", err)
	}
	if _, err := conn.ExecContext(ctx, `create table testTable (id int primary key, name varchar)`); err != nil {
		t.Fatal("Error creating table", err)
	}
	var id int
	if err := conn.QueryRowContext(ctx, `insert into testTable values (1, 'adriano') RETURNING id`).Scan(&id); id != 1 || err != nil {
		t.Fatal("Error inserting on table", id, err)
	}
}
