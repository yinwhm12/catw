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
	expireToken := time.Now().Add(time.Hour * 24).Unix()

	key := []byte(token_key)

	claims := MyCustomClaims{
		email,
		jwt.StandardClaims{
			ExpiresAt:expireToken,
			Issuer:"yin",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenStr, err = token.SignedString(key)
	return
}

func CheckToken(tokenStr string) (flag int, err error) {
	token, err := jwt.ParseWithClaims(tokenStr,&MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
			fmt.Println("11111")
			return nil ,fmt.Errorf("Unexpected signing method %v",token.Header["alg"])
		}
		fmt.Println("2222")
		return []byte(token_key), nil
	})
	token.Claims.Valid()

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid{
		fmt.Println("%v %v %v",claims.Email,claims.StandardClaims.ExpiresAt)
		return OK,err
	}else{

		fmt.Println("email:",claims.Email)
			if ve, ok := err.(*jwt.ValidationError); ok{
				if ve.Errors&jwt.ValidationErrorMalformed != 0{
					fmt.Println("4444")
					return Fail, err
				} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0{
					fmt.Println("5555")
					return TimeOver, err
				}else {
					fmt.Println("6666")
					return Fail, err
				}
		}
	}
	return

	//if token.Valid{
	//	fmt.Println("3333")
	//	return OK, err
	//}else if ve, ok := err.(*jwt.ValidationError); ok {
	//	if ve.Errors&jwt.ValidationErrorMalformed != 0{
	//		fmt.Println("4444")
	//		return Fail, err
	//	}else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0{
	//		fmt.Println("5555")
	//		return TimeOver, err
	//	}else {
	//		fmt.Println("6666")
	//		return Fail, err
	//	}
	//}else {
	//	fmt.Println("7777")
	//	return Fail,err
	//}

	//if _, ok := token.Claims.(*MyCustomClaims); ok && token.Valid{
	//	//return claims,ok
	//	return
	//}else {
	//	return
	//}
}

//func Login(email, pwd string)(u models.User,err error)  {
//
//}