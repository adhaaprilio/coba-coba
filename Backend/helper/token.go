package helper

import (
	"backend/config"
	"backend/entity"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// var mySigningKey = []byte(config.GetString("JWT_SECRET"))

type JWTClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

func GenerateToken(user *entity.User) (string, error) {
	mySigningKey := []byte(config.GetString("JWT_SECRET"))
	claims := JWTClaims{
		user.Id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	return ss, err
}

func ValidateToken(tokenString string) (*string, error) {
	secret := []byte(config.GetString("JWT_SECRET"))
	claims := &JWTClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("invalid token")
		}
		return nil, errors.New("ioken expired")
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return &claims.ID, nil
}
