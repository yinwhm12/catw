package main

//import (
//	"yinwhm.com/yin/catw/client"
//	"fmt"
//	"yinwhm.com/yin/catw/models"
//	"github.com/astaxie/beego/orm"
//	_"github.com/go-sql-driver/mysql"
//	"github.com/astaxie/beego"
//)

//func init()  {
//	link := fmt.Sprintf("%s:%s@(%s:%s)/%s", beego.AppConfig.String("mysqluser"),
//		beego.AppConfig.String("mysqlpass"), beego.AppConfig.String("mysqlurls"),
//		beego.AppConfig.String("mysqlport"), beego.AppConfig.String("mysqldb"))
//	orm.RegisterDataBase("default", "mysql", link)
//
//	orm.Debug = beego.BConfig.RunMode == "dev"
//}
//
//func main()  {
//
//	//var u models.User
//	var email string
//	//var pwd string
//	email = "yinwhm@163.com"
//	//token, err := client.SetToken(email, pwd); if err != nil {
//	//	fmt.Println("can't not set token")
//	//	return
//	//}
//
//
//	 user, err := models.GetUserInfoByEmail(email)
//	if err != nil{
//		fmt.Println("no user")
//		return
//	}
//	token := user.AccessToken
//
//
//
//	fmt.Println("token---",token)
//	//u.Email = email
//	//u.AccessToken = token
//	//id,err := models.AddUser(&u)
//	//if err !=nil{
//	//	fmt.Println("error")
//	//	return
//	//}
//	//fmt.Println(id)
//
//	flag, err := client.CheckToken(token)
//	if err != nil{
//		if client.Fail == flag{
//			fmt.Println("pwd error")
//			return
//		}else if client.TimeOver == flag{
//			fmt.Println("time over")
//			return
//		}
//	}
//	fmt.Println("---ok")
//}