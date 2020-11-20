package storage

type UnitOfWorkInterface interface{
	BeginTran() error
	Rollback() error
	Commit() error
	GetDatabase() DatabaseInterface
}