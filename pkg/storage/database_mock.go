package storage

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"time"
)

type DatabaseMock struct {
	ExecMock func(string, ...interface{}) (sql.Result, error)
	PrepareMock func(string) (*sql.Stmt, error)
	QueryMock func(string, ...interface{}) (*sql.Rows, error)
	QueryRowMock func(string, ...interface{}) *sql.Row
	BeginMock func() (*TransactionInterface, error)
	BeginTxMock func(context.Context, *sql.TxOptions) (*TransactionInterface, error)
	CloseMock func() error
	ConnMock func(ctx context.Context) (*ConnectionInterface, error)
	DriverMock func() driver.Driver
	ExecContextMock func(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PingMock func() error
	PingContextMock func(ctx context.Context) error
	PrepareContextMock func(ctx context.Context, query string) (*sql.Stmt, error)
	QueryContextMock func(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContextMock func(ctx context.Context, query string, args ...interface{}) *sql.Row
	SetConnMaxIdleTimeMock func(d time.Duration)
	SetConnMaxLifetimeMock func(d time.Duration)
	SetMaxIdleConnsMock func(n int)
	SetMaxOpenConnsMock func(n int)
	StatsMock func() sql.DBStats
}

func (mock *DatabaseMock) Exec(query string, args ...interface{}) (sql.Result, error){
	return mock.ExecMock(query, args)
}

func (mock *DatabaseMock) Prepare(query string) (*sql.Stmt, error){
	return mock.PrepareMock(query)
}

func (mock *DatabaseMock) Query(query string, args ...interface{}) (*sql.Rows, error){
	return mock.QueryMock(query, args)
}

func (mock *DatabaseMock) QueryRow(query string, args ...interface{}) *sql.Row{
	return mock.QueryRowMock(query, args)
}

func (mock *DatabaseMock) Begin() (*TransactionInterface, error){
	return mock.BeginMock()
}

func (mock *DatabaseMock) BeginTx(ctx context.Context, opts *sql.TxOptions) (*TransactionInterface, error){
	return mock.BeginTxMock(ctx, opts)
}

func (mock *DatabaseMock) Close() error{
	return mock.CloseMock()
}

func (mock *DatabaseMock) Conn(ctx context.Context) (*ConnectionInterface, error){
	return mock.ConnMock(ctx)
}

func (mock *DatabaseMock) Driver() driver.Driver{
	return mock.DriverMock()
}

func (mock *DatabaseMock) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error){
	return mock.ExecContextMock(ctx, query, args)
}

func (mock *DatabaseMock) Ping() error{
	return mock.PingMock()
}

func (mock *DatabaseMock) PingContext(ctx context.Context) error{
	return mock.PingContextMock(ctx)
}

func (mock *DatabaseMock) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error){
	return mock.PrepareContextMock(ctx, query)
}

func (mock *DatabaseMock) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error){
	return mock.QueryContextMock(ctx, query, args)
}

func (mock *DatabaseMock) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row{
	return mock.QueryRowContextMock(ctx, query, args)
}

func (mock *DatabaseMock) SetConnMaxIdleTime(d time.Duration){
	mock.SetConnMaxIdleTimeMock(d)
}

func (mock *DatabaseMock) SetConnMaxLifetime(d time.Duration){
	mock.SetConnMaxLifetimeMock(d)
}

func (mock *DatabaseMock) SetMaxIdleConns(n int){
	mock.SetMaxIdleConnsMock(n)
}

func (mock *DatabaseMock) SetMaxOpenConns(n int){
	mock.SetMaxOpenConnsMock(n)
}

func (mock *DatabaseMock) Stats() sql.DBStats{
	return mock.StatsMock()
}


