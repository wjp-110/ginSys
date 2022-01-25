package common

import (
	"ginSys/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("das_sdm_ffd")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

//生成token
func ReleaseToken(user model.User) (string, error) {
	expriationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expriationTime.Unix(),
			IssuedAt: time.Now().Unix(),
			Issuer: "wjp",
			Subject: "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString,nil
}
//解析token
func ParseToken(tokenString string) (*jwt.Token, *Claims,error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString,claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}