package main

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

var pdb PostgresDB
var runTests = true

// Postgres must be running at localhost:5432 to run these tests
// If it isn't, return error and set runTests to false to avoid having to reevaluate
func setupPostgresTests() error {
	if !runTests {
		return errors.New("Skipping tests, Postgres should be running for these tests")
	}

	config := PostgresDBConfig{host: "localhost", port: 5432, user: "postgres", password: "docker", dbname: "urlshortener"}
	pdb = PostgresDB{config: config}

	db, err := sql.Open("postgres", pdb.getConnectionString())
	if err != nil {
		runTests = false
		return err
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		runTests = false
		return err
	}

	return nil
}

func TestGetAccessLogsAllTime(t *testing.T) {
	// Arrange
	err := setupPostgresTests()
	if err != nil {
		t.Skip()
	}

	currentTime, err := time.Parse("2006-01-02 15:04:05", "2019-09-08 17:00:00")
	if err != nil {
		panic(err)
	}

	var timeFrom time.Time
	timeTo := currentTime

	// Act
	logs, err := pdb.GetAccessLogs("dfgSa", timeFrom, timeTo)

	// Assert
	if err != nil {
		t.Errorf("%v", err)
	}

	if len(logs) != 5 {
		t.Errorf("Expected 5 - got %v", len(logs))
	}
}

func TestGetAccessLogsPastWeek(t *testing.T) {
	// Arrange
	err := setupPostgresTests()
	if err != nil {
		t.Skip()
	}

	currentTime, err := time.Parse("2006-01-02 15:04:05", "2019-09-08 17:00:00")
	if err != nil {
		panic(err)
	}

	duration, _ := time.ParseDuration("-168h") // 7 days
	timeFrom := currentTime.Add(duration)
	timeTo := currentTime

	// Act
	logs, err := pdb.GetAccessLogs("dfgSa", timeFrom, timeTo)

	// Assert
	if err != nil {
		t.Errorf("%v", err)
	}

	if len(logs) != 3 {
		t.Errorf("Expected 3 - got %v", len(logs))
	}
}

func TestGetAccessLogsTwentyFourHours(t *testing.T) {
	// Arrange
	err := setupPostgresTests()
	if err != nil {
		t.Skip()
	}

	currentTime, err := time.Parse("2006-01-02 15:04:05", "2019-09-08 17:00:00")
	if err != nil {
		panic(err)
	}

	duration, _ := time.ParseDuration("-24h")
	timeFrom := currentTime.Add(duration)
	timeTo := currentTime

	// Act
	logs, err := pdb.GetAccessLogs("dfgSa", timeFrom, timeTo)

	// Assert

	if len(logs) != 2 {
		t.Errorf("Expected 2 - got %v", len(logs))
	}
}

func TestLogAccess(t *testing.T) {
	// Arrange
	err := setupPostgresTests()
	if err != nil {
		t.Skip()
	}
	initialCount := getAccessLogCount(3)

	// Act
	err = pdb.LogAccess("EIFdf")

	// Assert
	if err != nil {
		t.Errorf("%v", err)
	}

	finalCount := getAccessLogCount(3)

	if finalCount != initialCount+1 {
		t.Errorf("Mismatch in expected and actual log count - expected %v, got %v", initialCount+1, finalCount)
	}
}

func getAccessLogCount(id int) int {
	db, err := sql.Open("postgres", pdb.getConnectionString())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var count int
	row := db.QueryRow("SELECT COUNT(*) FROM access_logs WHERE shortened_url_id = $1", id)
	err = row.Scan(&count)
	if err != nil {
		panic(err)
	}

	return count
}
