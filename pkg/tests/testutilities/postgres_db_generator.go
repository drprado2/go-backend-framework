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

func CreateRandomDB() (string, string, string, error) {
	dbName, err := uuid.NewUUID()
	if err != nil {
		return emptyString, emptyString, emptyString, err
	}

	config, err := configs.GetConfig()
	if err != nil {
		return emptyString, emptyString, emptyString, err
	}

	connInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s sslmode=disable",
		config.DatabaseHost, config.DatabasePort, config.DatabaseUser, config.DatabasePassword)
	conn, err := sql.Open("postgres", connInfo)
	defer conn.Close()
	if err != nil {
		return emptyString, emptyString, emptyString, err
	}
	err = conn.Ping()
	if err != nil {
		return emptyString, emptyString, emptyString, err
	}
	sql := `CREATE DATABASE "` + dbName.String() + `";`
	if _, err := conn.Exec(sql); err != nil {
		return emptyString, emptyString, emptyString, err
	}
	connectionWithDb := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.DatabaseHost, config.DatabasePort, config.DatabaseUser, config.DatabasePassword, dbName.String())

	return dbName.String(), connectionWithDb, connInfo, nil
}
