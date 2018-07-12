package models

// AuthRequest represents login credentials
type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
