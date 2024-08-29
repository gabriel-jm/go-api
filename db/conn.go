package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	connectionString := "host=pg port=5432 user=postgres" +
		" password=postgres dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	println("DB Connected!")

	return db, nil
}
