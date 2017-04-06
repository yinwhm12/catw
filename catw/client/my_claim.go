package client

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"fmt"
)


type MyCustomClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}



func SetToken(email string) (tokenStr string, err error) {
	expireToken := time.Now().Add(time.Minute * 5).Unix()

	key := []byte(token_key)

	claims := MyCustomClaims{
		email,
		jwt.StandardClaims{
			ExpiresAt:expireToken,
			Issuer:"yinwhm12@163.com",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenStr, err = token.SignedString(key)
	return
}

func CheckToken(tokenStr string) (flag int, err error) {
	token, err := jwt.ParseWithClaims(tokenStr,&MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil ,fmt.Errorf("Unexpected signing method %v",token.Header["alg"])
		}
		return []byte(token_key), nil
	})

	if token.Valid{
		return OK, err
	}else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0{
			return Fail, err
		}else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0{
			return TimeOver, err
		}else {
			return Fail, err
		}
	}else {
		return Fail,err
	}

	//if _, ok := token.Claims.(*MyCustomClaims); ok && token.Valid{
	//	//return claims,ok
	//	return
	//}else {
	//	return
	//}
}