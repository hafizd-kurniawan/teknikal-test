package middleware

import (
	"errors"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

const (
	bearer = "Bearer"
)

var (
	ErrUnAuthorized = errors.New("JWT token tidak tersedia")
)

type Claims struct {
	EmployeeId   int    `json:"id"`
	EmployeeCode string `json:"email"`
	jwt.RegisteredClaims
}

var jwtSecreteKey = ""

func SetJWTSecretKey(key string) {
	jwtSecreteKey = key
}

func GenerateNewJWT(claims *Claims) (signedToken string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString([]byte(jwtSecreteKey))
	if err != nil {
		return
	}
	return
}

func GetJWTClaims(tokenString string) (claims *Claims, err error) {
	splitToken := strings.Split(tokenString, bearer)
	if len(splitToken) != 2 {
		return claims, ErrUnAuthorized
	}

	reqToken := strings.TrimSpace(splitToken[1])
	claims = &Claims{}
	token, err := jwt.ParseWithClaims(reqToken, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSecreteKey), nil
	})
	if err != nil {
		return nil, ErrUnAuthorized
	}

	if !token.Valid {
		return nil, ErrUnAuthorized
	}

	return claims, nil

}
