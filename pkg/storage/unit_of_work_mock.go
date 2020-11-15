package storage

type UnitOfWorkMock struct {
	BeginTranMock func() error
	RollbackMock func() error
	CommitMock func() error
	GetDatabaseMock func() *DatabaseInterface
}

func (mock *UnitOfWorkMock) BeginTran() error{
	return mock.BeginTranMock()
}

func (mock *UnitOfWorkMock) Rollback() error{
	return mock.RollbackMock()
}

func (mock *UnitOfWorkMock) Commit() error{
	return mock.CommitMock()
}

func (mock *UnitOfWorkMock) GetDatabase() *DatabaseInterface{
	return mock.GetDatabaseMock()
}

