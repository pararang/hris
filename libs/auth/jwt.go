package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// JWTClaims represents the claims in a JWT
type JWTClaims struct {
	UserID   uuid.UUID `json:"user_id"`
	Username string    `json:"username"`
	Role     string    `json:"role"`
	jwt.RegisteredClaims
}

// JWTService provides JWT token generation and validation
type JWTService struct {
	secretKey     string
	tokenDuration time.Duration
}

// NewJWTService creates a new instance of JWTService
func NewJWTService(secretKey string, tokenDuration time.Duration) *JWTService {
	return &JWTService{
		secretKey:     secretKey,
		tokenDuration: tokenDuration,
	}
}

// GenerateToken generates a new JWT token
func (s *JWTService) GenerateToken(userID uuid.UUID, username, role string) (string, error) {
	claims := &JWTClaims{
		UserID:   userID,
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.tokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(s.secretKey))
}

// ValidateToken validates a JWT token
func (s *JWTService) ValidateToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&JWTClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(s.secretKey), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
