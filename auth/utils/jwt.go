package utils

import (
	"aoisoft/auth/models"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

type jwtClaims struct {
	jwt.RegisteredClaims
	Id    int64
	Email string
}

func (w *JwtWrapper) GenerateToken(user models.User) (singedToken string, err error) {
	claims := &jwtClaims{
		Id:    user.Id,
		Email: user.Email,
		// 更新成新版本JWT自带载荷jwt.StandardClaims -> RegisteredClaims
		// https://blog.csdn.net/memory_qianxiao/article/details/121055248
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * time.Duration(w.ExpirationHours))),
			Issuer:    w.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println(w.SecretKey)

	signedToken, err := token.SignedString([]byte(w.SecretKey))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return signedToken, nil
}

func (w *JwtWrapper) ValidateToken(signedToken string) (claims *jwtClaims, err error) {
	token, err := jwt.ParseWithClaims(signedToken, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(w.SecretKey), nil
	})
	if err != nil {
		// FIXME:  error strings should not be capitalized
		return nil, errors.New("Could not parse claims")
	}

	if claims, ok := token.Claims.(*jwtClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("JWT is expired")
	}
}
