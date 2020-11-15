package postgres

import "github.com/drprado2/go-backend-framework/pkg/storage"

type UnitOfWork struct {
	DB storage.FullDatabaseInterface
	tx storage.TransactionInterface
}

func NewUnitOfWork(db storage.FullDatabaseInterface) *UnitOfWork {
	return &UnitOfWork{
		DB: db,
	}
}

func (uow *UnitOfWork) BeginTran() error{
	tx, err := uow.DB.Begin()
	uow.tx = *tx
	return err
}

func (uow *UnitOfWork) Rollback() error{
	 return uow.tx.Rollback()
}

func (uow *UnitOfWork) Commit() error{
	return uow.tx.Commit()
}

func (uow *UnitOfWork) GetDatabase() *storage.DatabaseInterface{
	var db storage.DatabaseInterface = uow.tx
	return &db
}
