package connection

import (
	"database/sql"
	"uas_neoj/helper"

	_ "github.com/lib/pq"
)

func GetConnection() (*sql.DB, error) {

	dataSource := helper.LoadEnv("DATABASE_URL")

	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		return nil, err
	}

	return db, nil
}
