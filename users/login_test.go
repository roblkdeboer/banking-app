package users

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/roblkdeboer/banking-app/models"
)

func TestGetUserByPassword(t *testing.T) {
    // Create a mock database driver
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("Failed to create mock: %v", err)
    }
    defer db.Close()

	testUser := &models.User{
		Email:"test@gmail.com",
		Password:"$2a$10$1bjyURp9wbScoWi4dvZNQeQ0XsvQ/XtUjiOV.vyn.N1wRvA5ZwcH2",
    }
    // Set up expectations for the mock
    mock.ExpectQuery(regexp.QuoteMeta(`SELECT email, password FROM users WHERE email=$1`)).
        WithArgs(testUser.Email).
        WillReturnRows(sqlmock.NewRows([]string{"email", "password"}).
            AddRow(testUser.Email, testUser.Password))

    // Call GetUserByEmail with the mock database
    resultUser, _ := GetUserPassword(db, testUser.Email)

    // Verify that the returned user matches the expected user
    if resultUser == nil {
        t.Error("Expected user, got nil")
    } else {
        if resultUser.Email != testUser.Email || resultUser.Password != testUser.Password {
            t.Error("Returned user doesn't match the expected user")
        }
    }

    // Ensure that expectations were met
    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("Expectations were not met: %v", err)
    }
}