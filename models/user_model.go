package models

import "github.com/golang-jwt/jwt"

// User model
type User struct {
	ID       int    `json:"-"`
	Username string `gorm:"column:user_name" json:"user_name"`  
	Password string `json:"password"`
	Name     string `json:"-"`
}
// AuthResponse represents the JSON structure of the authentication response
type AuthResponse struct {
	AccessToken string `json:"access_token"`
}
// Jwt model for username and role
type Jwt struct {
	Username string `json:"user_name"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

// TableName overrides the table name used by User to `profiles`
func (User) TableName() string {
	return "user"
}
