package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestGeneratePasswordHash(t *testing.T) {
    // Define a test password
    testPassword := "mytestpassword"

    // Call the function to generate a password hash
    hash, err := GeneratePasswordHash(testPassword)

    // Assert that the error is nil, indicating a successful hash generation
    assert.NoError(t, err)

    // Assert that the generated hash is not empty
    assert.NotEmpty(t, hash)

    // TO DO: add more assertions as needed, such as checking the hash format or length
}

func TestVerifyPassword(t *testing.T) {
    // Define a test password and its corresponding hash
    testPassword := "mytestpassword"
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(testPassword), bcrypt.DefaultCost)
    if err != nil {
        t.Fatal("Failed to generate test password hash")
    }

    // Call the VerifyPassword function with the test password and its hash
    isMatch := VerifyPassword(testPassword, string(hashedPassword))

    // Assert that the password matches the hash
    assert.True(t, isMatch)

    // Test with a wrong password
    isMatch = VerifyPassword("wrongpassword", string(hashedPassword))
    
    // Assert that the wrong password doesn't match the hash
    assert.False(t, isMatch)
}