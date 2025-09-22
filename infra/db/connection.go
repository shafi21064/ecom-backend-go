package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetDBConnectionString() string {
	// user     -> postgres
	// password -> 12345678
	// host     -> localhost
	// port     -> 5432
	// Db Name  -> e-com

	return "user=postgres password=12345678 host=localhost port=5432 dbname=e-com"
}

func NewDBConnection() (*sqlx.DB, error) {

	dbSourse := GetDBConnectionString()
	dbCon, err := sqlx.Connect("postgres", dbSourse)

	if err != nil {
		println(err)
		return nil, err
	}
	return dbCon, nil
}
