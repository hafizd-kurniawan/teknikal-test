package middleware

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	bearer = "Bearer"
)

var (
	ErrUnAuthorized = errors.New("JWT token tidak tersedia")
)

type Claims struct {
	EmployeeID   int    `json:"employee_id"`
	EmployeeCode string `json:"employee_code"`
	EmployeeName string `json:"employee_name"`
	jwt.RegisteredClaims
}

var jwtSecretKey = ""

func SetJWTSecretKey(key string) {
	jwtSecretKey = key
}
func GenereateJWT2(employeeCode string) (string, error) {
	atSecretKey := "atSecretKeyxyz"

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["employee_code"] = employeeCode
	// claims["role_name"] = roleName
	claims["exp"] = time.Now().Add(time.Hour * 24 * 10).Unix()

	tokenString, err := token.SignedString([]byte(atSecretKey))
	if err != nil {

		return "", err
	}

	return tokenString, nil
}

func GenerateNewJWT(claims *Claims) (signedToken string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func GetJWTClaims(tokenString string) (claims *Claims, err error) {
	splitToken := strings.Split(tokenString, bearer)
	if len(splitToken) != 2 {
		return nil, ErrUnAuthorized
	}

	reqToken := strings.TrimSpace(splitToken[1])
	claims = &Claims{}
	token, err := jwt.ParseWithClaims(reqToken, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})
	if err != nil {
		return nil, ErrUnAuthorized
	}

	if !token.Valid {
		return nil, ErrUnAuthorized
	}

	return claims, nil
}
