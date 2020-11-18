package testutilities

import (
	"database/sql"
	"fmt"
	"github.com/drprado2/go-backend-framework/pkg/configs"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

const (
	emptyString = ""
)

func CreateRandomDB() (*sql.DB, *sql.DB, string, error) {
	dbName, err := uuid.NewUUID()
	if err != nil {
		return nil, nil, emptyString, err
	}

	config, err := configs.GetConfig()
	connStringWithoutDB := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable",
		config.DatabaseHost, config.DatabasePort, config.DatabaseUser, config.DatabasePassword)

	connectionWithoutDB, err := sql.Open("postgres", connStringWithoutDB)
	if err != nil {
		return nil, nil, emptyString, err
	}

	sqlCreateDB := `CREATE DATABASE "` + dbName.String() + `";`
	if _, err := connectionWithoutDB.Exec(sqlCreateDB); err != nil {
		return nil, nil, emptyString, err
	}

	connStringWithDB := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DatabaseHost, config.DatabasePort, config.DatabaseUser, config.DatabasePassword, dbName.String())

	connectionWithDB, err := sql.Open("postgres", connStringWithDB)
	if err != nil {
		return nil, nil, emptyString, err
	}

	return connectionWithoutDB, connectionWithDB, dbName.String(), nil
}
