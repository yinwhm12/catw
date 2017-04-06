package client

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"fmt"
	"yinwhm.com/yin/catw/models"
)


type MyCustomClaims struct {
	Email string `json:"email"`
	Pwd string `json:"pwd"`
	jwt.StandardClaims
}



func SetToken(email ,pwd string) (tokenStr string, err error) {
	expireToken := time.Now().Add(time.Hour * 24).Unix()

	key := []byte(token_key)

	claims := MyCustomClaims{
		email,
		pwd,
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

func Login(email, pwd string)(u models.User,err error)  {

}