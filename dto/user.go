package dto

// LoginRequest represents the request to login
type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse represents the response after successful login
type LoginResponse struct {
	Token string `json:"token"`
}
