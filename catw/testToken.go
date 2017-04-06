package main
//
//import (
//	"fmt"
//
//	jwt "github.com/dgrijalva/jwt-go"
//	"time"
//)
//
//type MyCustomClaims struct {
//	Email string `json:"email"`
//	jwt.StandardClaims
//}
//
//func main()  {
//	//crutime := time.Now().Unix()
//	//fmt.Println("crutime-->",crutime)
//	//
//	//h := md5.New()
//	//fmt.Println("h-->",h)
//	//
//	//fmt.Println("strconv.FormatInt(crutime,10)-->",strconv.FormatInt(crutime,10))
//	//io.WriteString(h, strconv.FormatInt(crutime,10))
//	//
//	//fmt.Println("h-->",h)
//	//
//	//token := fmt.Sprintf("%x",h.Sum(nil))
//	//fmt.Println("token-->",token)
//	//
//	//fmt.Println(len("8e1a188743c6077110da3c9778183031"))
//
//
//
//
//	//mySiginingKey := []byte("hzwy23")
//	//claims := &jwt.StandardClaims{
//	//	NotBefore: int64(time.Now().Unix() - 1000),
//	//	ExpiresAt: int64(time.Now().Unix() + 1000),
//	//	Issuer: "test1",
//	//}
//	//
//	//token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
//	//ss, err := token.SignedString(mySiginingKey)
//	//fmt.Println("sigin :", ss)
//	//t, err := jwt.Parse(ss, func(*jwt.Token) (interface{}, error) {
//	//	return mySiginingKey, nil
//	//})
//	//
//	//if err != nil{
//	//	fmt.Println("param fail ",err)
//	//	return
//	//}
//	//fmt.Println("origin ",t.Claims)
//	mySignKey := []byte("yin")
//	//mySignKey := "yin"
//	claims := MyCustomClaims{
//		"yin",
//		jwt.StandardClaims{
//			NotBefore: int64(time.Now().Unix() - 1000),
//			ExpiresAt: int64(time.Now().Unix() + 1000),
//			Issuer: "test",
//		},
//	}
//
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
//	ss, err := token.SignedString(mySignKey)
//	fmt.Println("%v",ss)
//
//	tt, err := jwt.ParseWithClaims(ss, &MyCustomClaims{}, func(tt *jwt.Token) (interface{}, error) {
//		return mySignKey, nil
//	})
//	if claims, ok := token.Claims.(*MyCustomClaims); ok && tt.Valid{
//		fmt.Println("-----",claims.Email)
//	}else {
//		fmt.Println("---------",err)
//	}
//
//}
