package data

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func getConnection() (*sql.DB, error) {

	pgConnString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		"db",
		"5432",
		"postgres",
		"postgres",
		"postgres",
	)
	return sql.Open("postgres", pgConnString)
}

func MakeMigration(db *sql.DB) error {
	q := `
    CREATE TABLE IF NOT EXISTS beers (
		id serial NOT NULL,
		created_at timestamp DEFAULT now(),
		updated_at timestamp DEFAULT now(),
		name VARCHAR(150) NOT NULL,
		brewery VARCHAR(150) NOT NULL,
		country VARCHAR(150) NOT NULL,
		price NUMERIC(10,2) NOT NULL,
		currency NUMERIC(10,2) NOT NULL
	);
    `

	rows, err := db.Query(string(q))
	if err != nil {
		return err
	}

	return rows.Close()
}
