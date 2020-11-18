package postgres

import (
	"context"
	"database/sql"
	"github.com/drprado2/go-backend-framework/pkg/storage"
	_ "github.com/lib/pq"
)

const (
	postgres = "postgres"
)

type DatabaseFactory struct {
	connectionString string
}

type Connection struct {
	sql.Conn
}

func (conn *Connection) BeginTx(ctx context.Context, opts *sql.TxOptions) (storage.TransactionInterface, error) {
	tx, err := conn.Conn.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}

	var transaction storage.TransactionInterface = tx

	return transaction, nil
}

type Database struct {
	sql.DB
}

func (db *Database) Begin() (storage.TransactionInterface, error) {
	tx, err := db.DB.Begin()
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (db *Database) BeginTx(ctx context.Context, opts *sql.TxOptions) (storage.TransactionInterface, error) {
	tx, err := db.DB.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}

	var transaction storage.TransactionInterface = tx

	return transaction, nil
}

func (db *Database) Conn(ctx context.Context) (storage.ConnectionInterface, error) {
	conn, err := db.DB.Conn(ctx)
	if err != nil {
		return nil, err
	}

	var connection storage.ConnectionInterface = &Connection{Conn: *conn}

	return connection, nil
}

func NewDatabaseFactory(connString string) *DatabaseFactory {
	return &DatabaseFactory{
		connectionString: connString,
	}
}

func (factory *DatabaseFactory) GetDB() (storage.FullDatabaseInterface, error) {
	db, err := sql.Open(postgres, factory.connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	var database storage.FullDatabaseInterface = &Database{
		DB: *db,
	}

	return database, nil
}
