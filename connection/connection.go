package connection

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

// postgres://otquydnchftins:7b0065dd4e6f984c477f8578741e5769a79f652c5116956d4da6d4b859c126c5@ec2-44-205-41-76.compute-1.amazonaws.com:5432/dpk9ai6qssq9o

func LoadEnv(key string) string {
	return os.Getenv(key)
}

func GetConnection() (*sql.DB, error) {

	dataSource := LoadEnv("DATABASE_URL")

	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		return nil, err
	}

	return db, nil
}
