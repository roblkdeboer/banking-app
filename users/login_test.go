package users

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/roblkdeboer/banking-app/models"
)

func TestGetUserByEmail(t *testing.T) {
    // Create a mock database driver
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("Failed to create mock: %v", err)
    }
    defer db.Close()

	testUser := &models.User{
        ID:       16,
		Email:"test8@gmail.com",
		Password:"$2a$10$Tyj0S9JTYs3cTK6OikKqeeic2rPrlzkoEZpw7lj5GxPglGjOq5M6W",
    }
    // Set up expectations for the mock
    mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, email, password FROM users WHERE email=$1`)).
        WithArgs(testUser.Email).
        WillReturnRows(sqlmock.NewRows([]string{"id", "email", "password"}).
            AddRow(testUser.ID, testUser.Email, testUser.Password))

    // Call GetUserByEmail with the mock database
    resultUser, err := GetUserByEmail(db, testUser.Email)

    // Verify that the returned user matches the expected user
    if resultUser == nil {
        t.Error("Expected user, got nil")
    } else {
        if resultUser.ID != testUser.ID || resultUser.Email != testUser.Email || resultUser.Password != testUser.Password {
            t.Error("Returned user doesn't match the expected user")
        }
    }

    // Ensure that expectations were met
    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("Expectations were not met: %v", err)
    }
}