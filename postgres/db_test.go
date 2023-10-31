package utils

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestInitDB(t *testing.T) {
    // Create a mock database driver
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("Failed to create mock: %v", err)
    }
    defer db.Close()

    // Mock expectations
    mock.ExpectPing()

    // Call InitDB with the mock database driver
    Connection = db
    InitDB()

    // Verify that Connection was set
    if Connection == nil {
        t.Error("Connection is nil")
    }

    // Ensure that expectations were met
    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("Expectations were not met: %v", err)
    }
}