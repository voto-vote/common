package db

import (
	"database/sql"
	"os"

	"github.com/doug-martin/goqu/v9"
)

type PostgresConnector struct {
	DB      *sql.DB
	Dialect goqu.DialectWrapper
}

func GetConnection() (PostgresConnector, error) {
	dbAccess, err := sql.Open("postgres", os.Getenv("POSTGRES_CREDENTIALS"))
	if err != nil {
		return PostgresConnector{DB: dbAccess}, err
	}
	p := PostgresConnector{DB: dbAccess, Dialect: goqu.Dialect("postgres")}
	return p, nil
}
