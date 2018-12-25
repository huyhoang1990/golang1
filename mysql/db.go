package mysql

import (
	"../conf"
	"database/sql"
)

func ConnectDatabase(config *conf.Config) (*sql.DB, error) {
	return sql.Open("mysql", config.DatabasePath)
}
