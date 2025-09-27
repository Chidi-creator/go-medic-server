package services

import (
	"fmt"
	"github/Chidi-creator/go-medic-server/config"
	"github/Chidi-creator/go-medic-server/internal/models"

	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWT_SECRET = []byte(config.AppConfig.JWT_SECRET)

// expiry time for initial access and
var (
	AccessTokenExpiry  = time.Hour
	RefreshTokenExpiry = 7 * 24 * time.Hour
)

// jwt claims
type Claims struct {
	UserID string `json:"userid"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(user *models.User, tokenType string) (string, error) {
	var expirationTime time.Duration
	var subject string

	switch tokenType {
	case "access":
		expirationTime = AccessTokenExpiry
		subject = "access"
	case "refresh":
		expirationTime = RefreshTokenExpiry
		subject = "refesh"
	default:
		return "", fmt.Errorf("invalid Token Type")

	}

	claims := &Claims{
		UserID: user.ID.Hex(),
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationTime)),
			Issuer:    "medic-server",
			Subject:   subject,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString(JWT_SECRET)

	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func ValidateToken(tokenstr string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenstr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return JWT_SECRET, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error with parsing tokens: %w", err)
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
