package postgres

import (
	"github.com/drprado2/go-backend-framework/pkg/storage"
	_ "github.com/lib/pq"
)

type UnitOfWork struct {
	DB storage.FullDatabaseInterface
	tx storage.TransactionInterface
}

func NewUnitOfWork(db storage.FullDatabaseInterface) storage.UnitOfWorkInterface {
	return &UnitOfWork{
		DB: db,
	}
}

func (uow *UnitOfWork) BeginTran() error {
	tx, err := uow.DB.Begin()
	uow.tx = tx
	return err
}

func (uow *UnitOfWork) Rollback() error {
	err := uow.tx.Rollback()
	uow.tx = nil
	return err
}

func (uow *UnitOfWork) Commit() error {
	err := uow.tx.Commit()
	uow.tx = nil
	return err
}

func (uow *UnitOfWork) GetDatabase() storage.DatabaseInterface {
	if uow.tx != nil {
		return uow.tx
	}
	return uow.DB
}
