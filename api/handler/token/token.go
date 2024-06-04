package token

import (
	"context"
	"time"

	"github.com/bubblepays/models"

	"github.com/golang-jwt/jwt"

)

// Claims represents the structure of the JWT claims
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

//this should come from environment variable
var jwtKey = []byte("my_secret_key")


// GenerateToken creates a fake token for demonstration purposes
func GenerateToken(ctx *context.Context) (*models.AuthResponse, error) {
    // In a real application, this should generate a secure token
   // token :=  fmt.Sprintf("fake-token-%d", time.Now().Unix())

	// Set token expiration time
	expirationTime := time.Now().Add(5 * time.Minute)

	// Create the JWT claims, which includes the username and expiry time
	username := "test user name" //@TODO this will come http request
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Create the token with the specified claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return nil, err
	}

	//return tokenString, nil

	data := &models.AuthResponse{
		AccessToken: tokenString,
	}

	return data, nil
}