package main

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	//
	_ "github.com/lib/pq"
)

// PostgresDBConfig ...
type PostgresDBConfig struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
}

// PostgresDB ...
type PostgresDB struct {
	config PostgresDBConfig
}

func (pdb *PostgresDB) getConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", pdb.config.host, pdb.config.port, pdb.config.user, pdb.config.password, pdb.config.dbname)
}

// AddShortenedURL ...
func (pdb *PostgresDB) AddShortenedURL(su ShortenedURL) error {
	db, err := sql.Open("postgres", pdb.getConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare(`INSERT INTO shortened_urls (short_code, long_url) VALUES($1, $2)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(su.ShortCode, su.LongURL)
	if err != nil {
		return err
	}
	if res == nil {
		return errors.New("Error adding new URL to db")
	}

	return nil
}

// GetShortenedURL ...
func (pdb *PostgresDB) GetShortenedURL(shortCode string) (*ShortenedURL, error) {
	db, err := sql.Open("postgres", pdb.getConnectionString())
	if err != nil {
		return nil, err
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

// LogAccess will record access to a given shortCode to the database
func (pdb *PostgresDB) LogAccess(shortCode string) error {
	db, err := sql.Open("postgres", pdb.getConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare(`
	INSERT INTO access_logs (shortened_url_id, access_date_time)
	SELECT id, NOW()
	FROM shortened_urls
	WHERE short_code = $1;
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(shortCode)
	if err != nil {
		return err
	}
	if res == nil {
		return errors.New("Error logging access")
	}

	return nil
}

// GetAccessLogs will return access logs for the given shortCode in the given time frame
func (pdb *PostgresDB) GetAccessLogs(shortCode string, from time.Time, to time.Time) ([]AccessLog, error) {
	db, err := sql.Open("postgres", pdb.getConnectionString())
	if err != nil {
		return nil, err
	}
	defer db.Close()

	fromFormatted := from.Format("2006-01-02 15:04:05")
	toFormatted := to.Format("2006-01-02 15:04:05")

	rows, err := db.Query(`
	SELECT al.access_date_time
	FROM access_logs al
	INNER JOIN shortened_urls sl on al.shortened_url_id = sl.id
	WHERE sl.short_code = $1
	AND access_date_time BETWEEN $2 AND $3
	ORDER BY al.access_date_time DESC;`, shortCode, fromFormatted, toFormatted)

	if err != nil {
		return nil, err
	}

	var accessLog []AccessLog

	for rows.Next() {
		var dateTime time.Time
		err = rows.Scan(&dateTime)
		if err != nil {
			return nil, err
		}
		accessLog = append(accessLog, AccessLog{DateTime: dateTime})
	}

	return accessLog, nil
}
