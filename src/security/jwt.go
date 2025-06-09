package security

import (
	"MVC_DI/config"
	"MVC_DI/util"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TCustomClaims[T any] struct {
	Claims T
	jwt.RegisteredClaims
}

func GenerateJWT[T any](claims T) (string, error) {
	expiresAt := time.Now().Add(util.GetTime(config.Application.Jwt.Expiration))
	customClaims := TCustomClaims[T]{
		Claims: claims,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	return token.SignedString([]byte(config.Application.Jwt.Secret))
}

func ParseJWT[T any](tokenString string) (T, error) {
	claims := TCustomClaims[T]{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (any, error) {
		return []byte(config.Application.Jwt.Secret), nil
	})
	if err != nil {
		return claims.Claims, err
	}
	if claims, ok := token.Claims.(*TCustomClaims[T]); ok && token.Valid {
		return claims.Claims, nil
	}
	return claims.Claims, err
}

func CheckJWT(tokenString string) bool {
	_, err := ParseJWT[any](tokenString)
	return err == nil
}
