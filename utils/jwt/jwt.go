package jwt_helper

import (
	"MBFacto/config"
	"errors"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var once sync.Once

var accessSecret []byte
var refreshSecret []byte
var accessExpire time.Duration
var refreshExpire time.Duration

func InitJWT() {
	once.Do(func() {
		accessSecret = []byte(config.Cfg.JWT.AccessKey)
		refreshSecret = []byte(config.Cfg.JWT.RefreshKey)
		accessExpire = parseDuration(config.Cfg.JWT.AccessExpire)
		refreshExpire = parseDuration(config.Cfg.JWT.RefreshExpire)
	})
}

type JWTClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateTokens creates both access and refresh tokens
func GenerateTokens(userID uint) (string, string, error) {
	InitJWT()

	accessToken, err := generateJWT(userID, accessExpire, accessSecret)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := generateJWT(userID, refreshExpire, refreshSecret)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// Generate a JWT token with expiration
func generateJWT(userID uint, expiry time.Duration, secret []byte) (string, error) {
	claims := JWTClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func ValidateToken(tokenString string, isRefresh bool) (*JWTClaims, error) {
	InitJWT()

	secret := accessSecret
	if isRefresh {
		secret = refreshSecret
	}

	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func parseDuration(expiryStr string) time.Duration {
	log.Println(expiryStr)
	expiryInt, err := strconv.Atoi(expiryStr)
	if err != nil {
		log.Fatalf("Invalid value: %v", err)
	}
	return time.Minute * time.Duration(expiryInt)
}
