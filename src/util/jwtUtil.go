package util

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const signature = "KumaJie"
const expiration = time.Hour * 24 * 7

type UsrClaims struct {
	UserId int64
	jwt.RegisteredClaims
}

func GenerateToken(userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &UsrClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration)),
		},
	})
	tokenString, err := token.SignedString([]byte(signature))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) (*UsrClaims, error) {
	parseedToken, err := jwt.ParseWithClaims(tokenString, &UsrClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signature), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := parseedToken.Claims.(*UsrClaims)
	if !ok || !parseedToken.Valid {
		return nil, jwt.ErrSignatureInvalid
	}
	return claims, nil
}
