package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

type JwtCustomClaims struct {
	UserId      uint
	LoginVerify int
	jwt.StandardClaims
}

// jwt签名
func Signed(jwtMap jwt.MapClaims, key interface{}) (string, error) {
	tk := jwt.NewWithClaims(jwt.SigningMethodRS256, jwtMap)
	return tk.SignedString(key)
}

// jwt解密
func Pares(tokenString string, key interface{}) (interface{}, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})
	//if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, true
	} else {
		logrus.Errorf(err.Error())
		return "", false
	}
}
