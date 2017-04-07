package tool

import (
	"yinwhm.com/yin/catw/models/bean"
	"yinwhm.com/yin/catw/client"
	"github.com/astaxie/beego/context"
	"github.com/kataras/go-errors"
	"yinwhm.com/yin/catw/models"
)

func CreateSession(u bean.CreateSession)(userAuth *client.UserAuth, err error)  {
	userAuth, err = client.Login(u.Email,u.Pwd)
	if err != nil {
		return
	}
	return
}

func GetRequestToken(ctx *context.Context)string  {
	token := ctx.Input.Header("Authorization")
	if token == ""{
		token = ctx.Input.Query("_token")
	}
	return token
}

func OffLine(name string)(err error)  {
	err = client.OffLine(name)
	return
}

func VerifyToken(ctx *context.Context)(uid int, err error)  {
	token := GetRequestToken(ctx)
	if token == "" || len(token) <= 5{
		err = errors.New("forget token?")
		return
	}
	u, err := models.GetUserInfoByToken(token)
	if err != nil{
		err = errors.New("token invalid")
		return
	}
	res, err := client.CheckAccessToken(token,u.Name);
	if err == nil && res{
		uid = u.Id
		return
	}
	return
}

func Register(email, pwd string)(resemail string, err error){
	resemail, err = client.RegisterUser(email, pwd)
	if err != nil{
		return
	}
	return
}