package main

import (
	"database/sql"
	"errors"
	"fmt"

	//
	_ "github.com/lib/pq"
)

const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "docker"
	dbname   = "urlshortener"
)

// PostgresDB ...
type PostgresDB struct {
}

// Add ...
func (pdb PostgresDB) Add(su ShortenedURL) error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, err := db.Prepare(`INSERT INTO shortened_urls (short_code, long_url) VALUES($1, $2)`)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(su.ShortCode, su.LongURL)
	if err != nil {
		panic(err)
	}
	if res == nil {
		return errors.New("Error adding new URL to db")
	}

	return nil
}

// Get ...
func (pdb PostgresDB) Get(shortCode string) (*ShortenedURL, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	shortenedURL := ShortenedURL{}

	row := db.QueryRow("SELECT short_code, long_url FROM shortened_urls WHERE short_code = $1;", shortCode)
	err = row.Scan(&shortenedURL.ShortCode, &shortenedURL.LongURL)
	switch err {
	case sql.ErrNoRows:
		return nil, errors.New("Invalid short code supplied")
	case nil:
		return &shortenedURL, nil
	default:
		panic(err)
	}
}
