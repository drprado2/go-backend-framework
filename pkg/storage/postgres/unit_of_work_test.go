package postgres

import (
	"fmt"
	"github.com/drprado2/go-backend-framework/pkg/storage"
	"github.com/drprado2/go-backend-framework/pkg/tests/testutilities"
	"github.com/google/uuid"
	"testing"
)

type unitOfWorkFixture struct {
	databaseName string
	unitOfWork storage.UnitOfWorkInterface
	fullDB storage.FullDatabaseInterface
	connectionWithoutDB storage.FullDatabaseInterface
}

func (fixture *unitOfWorkFixture) setup(t *testing.T) {
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
	fixture.databaseName = dbName
	fixture.unitOfWork = NewUnitOfWork(fixture.fullDB)
}

func (fixture *unitOfWorkFixture) teardown(t *testing.T) {
	defer fixture.fullDB.Close()
	defer fixture.connectionWithoutDB.Close()
	_, err := fixture.connectionWithoutDB.Exec(`drop database "` + fixture.databaseName + `" WITH (FORCE);`)
	if err != nil {
		t.Error("Error on terardown", err)
	}
}

func TestQueriesWithoutTransaction(t *testing.T){
	fixture := unitOfWorkFixture{}
	fixture.setup(t)
	defer fixture.teardown(t)

	db := fixture.unitOfWork.GetDatabase()
	tableName, _ := uuid.NewUUID()
	if row := db.QueryRow(fmt.Sprintf(`create table "%s" (id int primary key, name varchar)`, tableName.String())); row.Err() != nil {
		t.Fatal("Error creating table", row.Err())
	}

	var id int
	err := db.QueryRow(fmt.Sprintf(`insert into "%s" values (1, 'adriano') RETURNING id`, tableName.String())).Scan(&id)
	if id != 1 || err != nil {
		t.Fatal("Error inserting on table", id, err)
	}
	countQuery := fmt.Sprintf(`select count(*) from "%s"`, tableName.String())

	var count int
	db.QueryRow(countQuery).Scan(&count)
	if count != 1 {
		t.Fatalf("Rows count without transaction should be 1 got %v", count)
	}
}

func TestQueriesWithRollbackTransaction(t *testing.T){
	fixture := unitOfWorkFixture{}
	fixture.setup(t)
	defer fixture.teardown(t)

	fixture.unitOfWork.BeginTran()
	db := fixture.unitOfWork.GetDatabase()
	tableName, _ := uuid.NewUUID()
	if row := db.QueryRow(fmt.Sprintf(`create table "%s" (id int primary key, name varchar)`, tableName.String())); row.Err() != nil {
		t.Fatal("Error creating table", row.Err())
	}

	var id int
	err := db.QueryRow(fmt.Sprintf(`insert into "%s" values (1, 'adriano') RETURNING id`, tableName.String())).Scan(&id)
	if id != 1 || err != nil {
		t.Fatal("Error inserting on table", id, err)
	}
	countQuery := fmt.Sprintf(`select count(*) from "%s"`, tableName.String())

	var countInTx int
	db.QueryRow(countQuery).Scan(&countInTx)
	if countInTx != 1 {
		t.Fatalf("Rows count without transaction should be 1 got %v", countInTx)
	}

	if err := fixture.unitOfWork.Rollback(); err != nil {
		t.Fatalf("Error on rollback\nError: %s", err)
	}

	var countAfterRollback int
	db.QueryRow(countQuery).Scan(&countAfterRollback)
	if countAfterRollback != 0 {
		t.Fatalf("Rows count without transaction should be 0 got %v", countAfterRollback)
	}
}

func TestQueriesWithCommitTransaction(t *testing.T){
	fixture := unitOfWorkFixture{}
	fixture.setup(t)
	defer fixture.teardown(t)

	fixture.unitOfWork.BeginTran()
	db := fixture.unitOfWork.GetDatabase()
	tableName, _ := uuid.NewUUID()
	if row := db.QueryRow(fmt.Sprintf(`create table "%s" (id int primary key, name varchar)`, tableName.String())); row.Err() != nil {
		t.Fatal("Error creating table", row.Err())
	}

	var id int
	err := db.QueryRow(fmt.Sprintf(`insert into "%s" values (1, 'adriano') RETURNING id`, tableName.String())).Scan(&id)
	if id != 1 || err != nil {
		t.Fatal("Error inserting on table", id, err)
	}
	countQuery := fmt.Sprintf(`select count(*) from "%s"`, tableName.String())

	var countInTx int
	db.QueryRow(countQuery).Scan(&countInTx)
	if countInTx != 1 {
		t.Fatalf("Rows count without transaction should be 1 got %v", countInTx)
	}

	if err := fixture.unitOfWork.Commit(); err != nil {
		t.Fatalf("Error on rollback\nError: %s", err)
	}
	db = fixture.unitOfWork.GetDatabase()

	var countAfterRollback int
	db.QueryRow(countQuery).Scan(&countAfterRollback)
	if countAfterRollback != 1 {
		t.Fatalf("Rows count without transaction should be 1 got %v", countAfterRollback)
	}
}