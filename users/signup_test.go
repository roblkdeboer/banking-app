package users

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/roblkdeboer/banking-app/models"
	"github.com/stretchr/testify/assert"
)

// For valid user data
func TestInsertUser(t *testing.T) {
    // Create a mock database driver
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("Failed to create mock: %v", err)
    }
    defer db.Close()
	
    // Create a test user and hashed password
    testUser := models.User{
        FirstName: "John",
        LastName:  "Doe",
        Phone:     "1234567890",
        Email:     "testuser@example.com",
    }
    hashedPassword := "known-hashed-password" // You provide the hashed password

    // Set up expectations for the mock
    mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO users (first_name, last_name, phone, email, password) VALUES ($1, $2, $3, $4, $5`)).
        WithArgs(testUser.FirstName, testUser.LastName, testUser.Phone, testUser.Email, hashedPassword).
        WillReturnResult(sqlmock.NewResult(1, 1))

    // Call the InsertUser function with the mock database
    err = InsertUser(db, testUser, hashedPassword)

    // Assert that the function should not return an error for a successful insert
    assert.NoError(t, err)

    // Ensure that expectations were met
    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("Expectations were not met: %v", err)
    }
}