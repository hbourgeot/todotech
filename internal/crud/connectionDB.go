package crud

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func makeCN() (*sql.DB, error) {
	conn := "host=localhost port=5432 user=postgres password=reyshell dbname=todotech sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
