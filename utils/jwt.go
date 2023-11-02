package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWTToken(userID int) (string, error) {
    // Create a new token with a signing method
    token := jwt.New(jwt.SigningMethodHS256)

    // Set the claims (payload) for the token
    claims := token.Claims.(jwt.MapClaims)
    claims["userID"] = userID
    claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

    // Sign and get the complete encoded token as a string
    tokenString, err := token.SignedString([]byte("your-secret-key"))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}