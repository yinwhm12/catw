package filters

import (
	"github.com/astaxie/beego/context"
	"fmt"
)

var BeforeWrite = func(ctx *context.Context) {

	cookie := ctx.GetCookie("Auth")
	if len(cookie) == 0{
		fmt.Print("eerrrror--")
	}
	fmt.Println("------",cookie)

	//token := ctx.Input.Header("Authorization")
	//if token == ""{
	//	token = ctx.Input.Query("_token")
	//}
	//flag, _ := client.CheckToken(token)
	//if flag == client.OK{
	//	return
	//}else if flag == client.TimeOver{
	//	ctx.Redirect(bean.CODE_Forbidden,"/v1")
	//}else{
	//
	//}
}
