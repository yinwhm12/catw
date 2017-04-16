package client

import (
	"github.com/astaxie/beego/context"
	"yinwhm.com/yin/catw/models/bean"
)


func AllowCrows(ctx *context.Context,err error)  {
	if ctx.Request.Method == "OPTIONS" {
		ctx.Input.SetData("uid", 0)
		return
	}
	ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")                               //允许访问源
	ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST,DELETE, GET, PUT, OPTIONS") //允许post访问
	ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")     //header的类型
	ctx.ResponseWriter.Header().Set("Access-Control-Max-Age", "1728000")
	ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
	ctx.ResponseWriter.Header().Set("content-type", "application/json") //返回数据格式是json
	ctx.ResponseWriter.ResponseWriter.WriteHeader(bean.CODE_Unauthorized)
	ctx.WriteString(err.Error())
}
