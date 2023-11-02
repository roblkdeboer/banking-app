package users

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/roblkdeboer/banking-app/models"
	"github.com/stretchr/testify/assert"
)

func TestGetUserPassword(t *testing.T) {
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
    mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, email, password FROM users WHERE email=$1`)).
        WithArgs(testUser.Email).
        WillReturnRows(sqlmock.NewRows([]string{"id", "email", "password"}).
            AddRow(testUser.ID, testUser.Email, testUser.Password))

    // Call GetUserByEmail with the mock database
    resultUser, _ := GetUserPassword(db, testUser.Email)

	assert.NotNil(t, resultUser, "Expected user, got nil")
    assert.Equal(t, testUser.Email, resultUser.Email, "Emails should match")
    assert.Equal(t, testUser.Password, resultUser.Password, "Passwords should match")

    // Ensure that expectations were met
    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("Expectations were not met: %v", err)
    }
}