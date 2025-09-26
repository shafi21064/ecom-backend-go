package db

import (
	"fmt"

	"e-com/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetDBConnectionString(cnf *config.DBConfig) string {
	conString := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s", cnf.DbUser, cnf.DbPassword, cnf.DbHost, cnf.DbPort, cnf.DbName)
	return conString
}

func NewDBConnection(cnf *config.DBConfig) (*sqlx.DB, error) {

	dbSourse := GetDBConnectionString(cnf)
	dbCon, err := sqlx.Connect("postgres", dbSourse)

	if err != nil {
		println(err)
		return nil, err
	}
	return dbCon, nil
}
