package main

import (
	"fmt"
	"yinwhm.com/yin/catw/models"
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
)

func init()  {
	link := fmt.Sprintf("%s:%s@(%s:%s)/%s", beego.AppConfig.String("mysqluser"),
		beego.AppConfig.String("mysqlpass"), beego.AppConfig.String("mysqlurls"),
		beego.AppConfig.String("mysqlport"), beego.AppConfig.String("mysqldb"))
	orm.RegisterDataBase("default", "mysql", link)

	orm.Debug = beego.BConfig.RunMode == "dev"
}

func main()  {

	//articles, err := models.GetArticleByUser(3)
	//if err !=nil{
	//	fmt.Println("----err")
	//	return
	//}
	//fmt.Println("len",len(articles))
	//for _,v := range articles{
	//	fmt.Println("user:",v.User,"email",v.User.Email)
	//}

	//ok
	//articles, err := models.GetArticleByEndType(4);
	//if err != nil{
	//	fmt.Println("----err")
	//	return
	//}
	//fmt.Println("len",len(articles))
	//for _,v := range articles{
	//	fmt.Println("type",v.EndType,"name",v.EndType.Root1Type)
	//}

	//ok
	//var article models.Article
	//
	//user,err := models.GetUserById(3)
	//if err != nil{
	//	fmt.Println("err --")
	//	return
	//}
	//
	//endType, err := models.GetEndTypeInfoById(4)
	//if err != nil{
	//	fmt.Println("end---")
	//	return
	//}
	//
	//article.User = user
	//article.EndType = endType
	//article.Title = "test"
	//if _,err = models.AddArticle(&article); err != nil{
	//	fmt.Println("add ")
	//	return
	//}

	root, err := models.GetAllRoot1TypeInfo(); if err != nil{
		fmt.Println("errr---")
		return
	}
	for id,v := range root{
		fmt.Println("id=",id,"name:",v.Root1TypeName)
	}

	fmt.Println("ok")
}

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