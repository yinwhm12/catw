package controllers

import (
	"yinwhm.com/yin/catw/models/bean"
	"encoding/json"
	"net/http"
	"yinwhm.com/yin/catw/tool"
	"yinwhm.com/yin/catw/models"
	"fmt"
	"github.com/kataras/go-errors"
	"yinwhm.com/yin/catw/client"
)

//会话 登录 用户简单信息记录 入口
type SessionController struct {
	BaseController
}

func (c *SessionController)URLMaping()  {
	c.Mapping("Post",c.Post)
	c.Mapping("Put",c.Put)
	c.Mapping("Delete",c.Delete)
}

// Post ...
// @Title 登录
// @Description 用户登录，创建会话
// @Param	body body
// @Success 200
// @Failure 403
// @router / [post]
func (c *SessionController)Post()  {
	var v bean.CreateSession
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); if err != nil{
		c.RespJSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := models.GetUserInfoByEmail(v.Email); if err != nil{
		c.RespJSON(http.StatusBadRequest,"用户不存在")
		return
	}
	//利用email pwd 生成一个token_login 在检验是否 过时或者 错误(假密码)
	//下面生成 token_login


	//验证token
	flag, err := client.CheckToken(user.AccessToken); if err != nil{
		if flag == client.Fail{//token can not handle
			c.RespJSON(http.StatusForbidden, err.Error())
			return
		}else if flag == client.TimeOver{//token超时 重新设置
			token, err := client.SetToken()
		}
	}

	userAuth, err := tool.CreateSession(v); if err != nil{
		c.RespJSON(http.StatusForbidden, err.Error())
		return
	}
	u, err := models.GetUserInfoByEmail(userAuth.Email); if err != nil{
		c.RespJSON(http.StatusBadRequest,err.Error())
		return
	}
	u.AccessToken = userAuth.AccessToken
	u.RefreshToken = userAuth.RefreshToken
	err = models.UpdateUserToken(u); if err != nil{
		c.RespJSON(http.StatusBadRequest,err.Error())
		return
	}
	c.RespJSON(http.StatusOK,bean.OutPutSession{Uid:u.Id,Token:u.AccessToken})
}

//Delete ...
//@Title 注销
//@Description 注销,删除会话
//@Param	x-token
//@Success 200
//@Failure 403 id is empty
//@router /:id [delete]
func (c *SessionController)Delete()  {
	token := tool.GetRequestToken(c.Ctx)
	u, err := models.GetUserInfoByToken(token); if err != nil{
		c.RespJSON(http.StatusBadRequest,err.Error())
		return
	}
	err = tool.OffLine(u.Name); if err != nil{
		c.RespJSON(http.StatusBadRequest,err.Error())
		return
	}
	c.RespJSON(http.StatusOK, http.StatusText(http.StatusOK))
	return
}

// 注册用户
// @Title 用户注册
// @Description 用户注册以及检验是否注册过
// @Param
// @Success 200 {string} "OK"
// @router /register [post]
func (c *SessionController)Register()  {
	var u bean.CreateSession
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &u); err != nil{
		c.RespJSON(bean.CODE_Params_Err,err.Error())
		return
	}
	if noRow := models.CheckEmailForRegister(u.Email); noRow == false{
		c.RespJSON(bean.CODE_Existed_User_Err,errors.New("had the same email"))
		return
	}

	//注册token
	 token, err := client.SetToken(u.Email,u.Pwd); if err != nil {
		c.RespJSON(bean.CODE_Not_Acceptable,"can't not create token")
		return
	}

	fmt.Println("----len:",len(token))
	var user models.User
	user.Email = u.Email
	user.AccessToken = token
	//注册用户
	if _,err := models.AddUser(&user); err != nil{
		c.RespJSON(bean.CODE_Params_Err,err.Error())
		return
	}

	fmt.Println("id----",user.Id)
	//expireCookie := time.Now().Add(time.Minute * 5)
	//
	//cookie := http.Cookie{Name:"Auth",Value:token,Expires:expireCookie,HttpOnly:true}
	//http.SetCookie(c.Ctx.ResponseWriter,&cookie)
	c.RespJSON(http.StatusOK,bean.OutPutSession{Uid:user.Id,Token:user.AccessToken,Email:user.Email})

}
