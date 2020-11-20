package postgres

import (
	"context"
	"fmt"
	"github.com/drprado2/go-backend-framework/pkg/configs"
	"github.com/drprado2/go-backend-framework/pkg/storage"
	"github.com/drprado2/go-backend-framework/pkg/tests/testutilities"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"testing"
)


type databaseFixture struct {
	databaseName string
	transaction storage.TransactionInterface
	connection storage.ConnectionInterface
	fullDB storage.FullDatabaseInterface
	DB storage.DatabaseInterface
	connectionWithoutDB storage.FullDatabaseInterface
}

func (fixture *databaseFixture) setup(t *testing.T) {
	connWithoutDB, connWithDB, dbName, err := testutilities.CreateRandomDB()
	if err != nil {
		t.Fatal("Error in setup", err)
	}

	fixture.connectionWithoutDB = &Database{
		DB: *connWithoutDB,
	}
	fixture.fullDB = &Database{
		DB: *connWithDB,
	}
	fixture.DB = connWithDB
	fixture.databaseName = dbName
}

func (fixture *databaseFixture) teardown(t *testing.T) {
	defer fixture.fullDB.Close()
	defer fixture.connectionWithoutDB.Close()
	_, err := fixture.connectionWithoutDB.Exec(`drop database "` + fixture.databaseName + `" WITH (FORCE);`)
	if err != nil {
		t.Error("Error on terardown", err)
	}
}

func TestDatabaseExecution(t *testing.T) {
	fixture := databaseFixture{}
	fixture.setup(t)
	defer fixture.teardown(t)

	if _, err := fixture.DB.Exec(`create table testTable (id int primary key, name varchar)`); err != nil {
		t.Fatal("Error creating table", err)
	}
	var id int
	if err := fixture.DB.QueryRow(`insert into testTable values (1, 'adriano') RETURNING id`).Scan(&id); id != 1 || err != nil {
		t.Fatal("Error inserting on table", id, err)
	}
}

func TestTransactionRollback(t *testing.T) {
	fixture := databaseFixture{}
	fixture.setup(t)
	defer fixture.teardown(t)

	tableName, _ := uuid.NewUUID()
	if row := fixture.DB.QueryRow(fmt.Sprintf(`create table "%s" (id int primary key, name varchar)`, tableName.String())); row.Err() != nil {
		t.Fatal("Error creating table", row.Err())
	}

	tx, err := fixture.fullDB.Begin()
	fixture.transaction = tx
	if err != nil {
		t.Fatalf("Fail creating transaction\n%s", err)
	}
	var id int
	err = fixture.transaction.QueryRow(fmt.Sprintf(`insert into "%s" values (1, 'adriano') RETURNING id`, tableName.String())).Scan(&id)
	if id != 1 || err != nil {
		t.Fatal("Error inserting on table", id, err)
	}
	countQuery := fmt.Sprintf(`select count(*) from "%s"`, tableName.String())

	var countInTx int
	fixture.transaction.QueryRow(countQuery).Scan(&countInTx)
	if countInTx != 1 {
		t.Fatalf("Rows count in transaction should be 1 got %v", countInTx)
	}

	err = fixture.transaction.Rollback()

	if err != nil {
		t.Fatalf("Fail rollback\nError: %s", err)
	}

	var countAfterRollback int
	fixture.DB.QueryRow(countQuery).Scan(&countAfterRollback)
	if countAfterRollback != 0 {
		t.Fatalf("Rows count in transaction after rollback should be 0 got %v", countAfterRollback)
	}
}

func TestTransactionCommit(t *testing.T) {
	fixture := databaseFixture{}
	fixture.setup(t)
	defer fixture.teardown(t)

	tableName, _ := uuid.NewUUID()
	if row := fixture.DB.QueryRow(fmt.Sprintf(`create table "%s" (id int primary key, name varchar)`, tableName.String())); row.Err() != nil {
		t.Fatal("Error creating table", row.Err())
	}

	tx, err := fixture.fullDB.Begin()
	fixture.transaction = tx
	if err != nil {
		t.Fatalf("Fail creating transaction\n%s", err)
	}
	var id int
	err = fixture.transaction.QueryRow(fmt.Sprintf(`insert into "%s" values (1, 'adriano') RETURNING id`, tableName.String())).Scan(&id)
	if id != 1 || err != nil {
		t.Fatal("Error inserting on table", id, err)
	}
	countQuery := fmt.Sprintf(`select count(*) from "%s"`, tableName.String())

	var countInTx int
	fixture.transaction.QueryRow(countQuery).Scan(&countInTx)
	if countInTx != 1 {
		t.Fatalf("Rows count in transaction should be 1 got %v", countInTx)
	}

	err = fixture.transaction.Commit()

	if err != nil {
		t.Fatalf("Fail rollback\nError: %s", err)
	}

	var countAfterRollback int
	fixture.DB.QueryRow(countQuery).Scan(&countAfterRollback)
	if countAfterRollback != 1 {
		t.Fatalf("Rows count in transaction after commit should be 1 got %v", countAfterRollback)
	}
}

func TestDatabaseConnectionExecution(t *testing.T) {
	fixture := databaseFixture{}
	fixture.setup(t)
	defer fixture.teardown(t)

	ctx := context.Background()
	conn, err := fixture.fullDB.Conn(ctx)
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

func TestTransactionWithContext(t *testing.T) {
	fixture := databaseFixture{}
	fixture.setup(t)
	defer fixture.teardown(t)

	tableName, _ := uuid.NewUUID()
	if row := fixture.DB.QueryRow(fmt.Sprintf(`create table "%s" (id int primary key, name varchar)`, tableName.String())); row.Err() != nil {
		t.Fatal("Error creating table", row.Err())
	}

	ctx := context.Background()
	tx, err := fixture.fullDB.BeginTx(ctx, nil)
	fixture.transaction = tx
	if err != nil {
		t.Fatalf("Fail creating transaction\n%s", err)
	}
	var id int
	err = fixture.transaction.QueryRowContext(ctx, fmt.Sprintf(`insert into "%s" values (1, 'adriano') RETURNING id`, tableName.String())).Scan(&id)
	if id != 1 || err != nil {
		t.Fatal("Error inserting on table", id, err)
	}
	countQuery := fmt.Sprintf(`select count(*) from "%s"`, tableName.String())

	var countInTx int
	fixture.transaction.QueryRowContext(ctx, countQuery).Scan(&countInTx)
	if countInTx != 1 {
		t.Fatalf("Rows count in transaction should be 1 got %v", countInTx)
	}

	err = fixture.transaction.Commit()

	if err != nil {
		t.Fatalf("Fail rollback\nError: %s", err)
	}

	var countAfterRollback int
	fixture.DB.QueryRowContext(ctx, countQuery).Scan(&countAfterRollback)
	if countAfterRollback != 1 {
		t.Fatalf("Rows count in transaction after commit should be 1 got %v", countAfterRollback)
	}
}

func TestDatabaseFactory(t *testing.T) {
	fixture := databaseFixture{}
	fixture.setup(t)
	defer fixture.teardown(t)

	config, _ := configs.GetConfig()
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DatabaseHost, config.DatabasePort, config.DatabaseUser, config.DatabasePassword, fixture.databaseName)

	factory := NewDatabaseFactory(connString)
	db, err := factory.GetDB()
	if err != nil {
		t.Fatalf("Error creating DB\nError: %s", err)
	}

	if _, err := db.Exec(`create table testTable (id int primary key, name varchar)`); err != nil {
		t.Fatal("Error creating table", err)
	}
	var id int
	if err := db.QueryRow(`insert into testTable values (1, 'adriano') RETURNING id`).Scan(&id); id != 1 || err != nil {
		t.Fatal("Error inserting on table", id, err)
	}
}

func TestTransactionWithConnection(t *testing.T) {
	fixture := databaseFixture{}
	fixture.setup(t)
	defer fixture.teardown(t)

	tableName, _ := uuid.NewUUID()
	if row := fixture.DB.QueryRow(fmt.Sprintf(`create table "%s" (id int primary key, name varchar)`, tableName.String())); row.Err() != nil {
		t.Fatal("Error creating table", row.Err())
	}

	ctx := context.Background()

	conn, err := fixture.fullDB.Conn(ctx)
	if err != nil {
		t.Fatal("Error creating connection", err)
	}

	tx, err := conn.BeginTx(ctx, nil)
	fixture.transaction = tx
	if err != nil {
		t.Fatalf("Fail creating transaction\n%s", err)
	}
	var id int
	err = fixture.transaction.QueryRowContext(ctx, fmt.Sprintf(`insert into "%s" values (1, 'adriano') RETURNING id`, tableName.String())).Scan(&id)
	if id != 1 || err != nil {
		t.Fatal("Error inserting on table", id, err)
	}
	countQuery := fmt.Sprintf(`select count(*) from "%s"`, tableName.String())

	var countInTx int
	fixture.transaction.QueryRowContext(ctx, countQuery).Scan(&countInTx)
	if countInTx != 1 {
		t.Fatalf("Rows count in transaction should be 1 got %v", countInTx)
	}

	err = fixture.transaction.Commit()

	if err != nil {
		t.Fatalf("Fail rollback\nError: %s", err)
	}

	var countAfterRollback int
	fixture.DB.QueryRowContext(ctx, countQuery).Scan(&countAfterRollback)
	if countAfterRollback != 1 {
		t.Fatalf("Rows count in transaction after commit should be 1 got %v", countAfterRollback)
	}
}
