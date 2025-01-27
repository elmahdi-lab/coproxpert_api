package auth

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"ithumans.com/coproxpert/models"
)

var (
	alg                    = jwt.SigningMethodHS512
	secret                 = os.Getenv("JWT_ENCRYPTION_KEY")
	tokenExpirationMinutes int
	refreshTokenHours      int
)

func init() {

	// Parse access token expiration
	var err error
	tokenExpirationMinutes, err = strconv.Atoi(os.Getenv("JWT_ACCESS_TOKEN_EXPIRATION_MINUTES"))
	if err != nil || tokenExpirationMinutes <= 0 {
		tokenExpirationMinutes = 15 // Default to 15 minutes
	}

	// Parse refresh token expiration
	refreshTokenHours, err = strconv.Atoi(os.Getenv("JWT_REFRESH_TOKEN_EXPIRATION_HOURS"))
	if err != nil || refreshTokenHours <= 0 {
		refreshTokenHours = 24 // Default to 24 hours
	}
}

// GenerateJWT generates an access token for a user
func GenerateJWT(userID uuid.UUID, createdAt time.Time) (string, error) {
	issuedAt := time.Now()
	expirationTime := issuedAt.Add(time.Duration(tokenExpirationMinutes) * time.Minute)

	claims := jwt.MapClaims{
		"sub": userID,                             // Subject
		"iss": "coproxpert",                       // Issuer
		"nbf": jwt.NewNumericDate(createdAt),      // Not Before
		"iat": jwt.NewNumericDate(issuedAt),       // Issued At
		"exp": jwt.NewNumericDate(expirationTime), // Expiration Time
	}

	token := jwt.NewWithClaims(alg, claims)

	// Sign the token using the secret key
	return token.SignedString([]byte(secret))
}

// ValidateJWT validates a given token and returns the claims if valid
func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	// Extract claims and validate the token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func GenerateRefreshToken(user *models.User) (string, error) {
	issuedAt := time.Now()
	refreshExpirationTime := issuedAt.Add(time.Duration(refreshTokenHours) * time.Hour)

	claims := jwt.MapClaims{
		"sub": user.ID,                                   // Subject
		"iss": "coproxpert",                              // Issuer
		"nbf": jwt.NewNumericDate(user.CreatedAt),        // Not Before
		"iat": jwt.NewNumericDate(issuedAt),              // Issued At
		"exp": jwt.NewNumericDate(refreshExpirationTime), // Expiration Time
		"typ": "refresh",                                 // Token type
	}

	token := jwt.NewWithClaims(alg, claims)

	// Sign the token using the secret key
	return token.SignedString([]byte(secret))
}
func RefreshTokens(refreshToken string) (string, error) {
	claims, err := ValidateJWT(refreshToken)
	if err != nil {
		return "", err
	}

	// Ensure the token is a refresh token
	typ, ok := claims["typ"].(string)
	if !ok || typ != "refresh" {
		return "", errors.New("invalid refresh token type")
	}

	// Ensure the subject (username) is present
	id, ok := claims["sub"].(uuid.UUID)
	if !ok || id == uuid.Nil {
		return "", errors.New("invalid refresh token: missing subject")
	}

	nbf, ok := claims["nbf"].(time.Time)
	if !ok || id == uuid.Nil {
		return "", errors.New("invalid refresh token: missing subject")
	}

	return GenerateJWT(id, nbf)
}
