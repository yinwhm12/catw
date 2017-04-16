package filters

import (
	"github.com/astaxie/beego/context"
	"yinwhm.com/yin/catw/client"
	"github.com/kataras/go-errors"
	"yinwhm.com/yin/catw/models"
)

var BeforeWrite = func(ctx *context.Context) {

	token := ctx.Request.Header.Get("Authorization")
	if token ==""{
		client.AllowCrows(ctx,errors.New("Miss Token,No Login!"))
		return
	}else{
		flag, _ := client.CheckToken(token)
		if flag == client.OK{
			uid, err := models.GetUserIdByToken(token)
			if err != nil{
				client.AllowCrows(ctx,err)
				return
			}
			ctx.Input.SetData("uid",uid)
		}else if flag == client.TimeOver{
			client.AllowCrows(ctx,errors.New("No Long to See!"))
			return

		}else if flag == client.Fail{
			client.AllowCrows(ctx,errors.New("Token Wrong!"))
			//ctx.Redirect(401,"/")
			return
		}else {
			client.AllowCrows(ctx,errors.New("Token Wrong!"))
			return
		}
	}
}
