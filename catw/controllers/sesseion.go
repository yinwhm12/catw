package controllers

import (
	"yinwhm.com/yin/catw/models/bean"
	"encoding/json"
	"net/http"
	"yinwhm.com/yin/catw/tool"
	"yinwhm.com/yin/catw/models"
	"github.com/kataras/go-errors"
	"yinwhm.com/yin/catw/client"
)

//会话 登录 用户简单信息记录 入口
type SessionController struct {
	BaseController
}

func (c *SessionController)URLMapping()  {
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

	if _, err := models.GetUserInfoByEmail(v.Email); err != nil{
		c.RespJSON(http.StatusBadRequest,"用户不存在")
		return
	}

	user, err := models.Login(v.Email,v.Pwd); if err != nil {
		c.RespJSON(bean.CODE_Params_Err,err.Error())
		return
	}

	//重新生成token
	token, err := client.SetToken(user.Email); if err != nil{
		c.RespJSON(bean.CODE_Bad_Request,err.Error())
		return
	}

	user.AccessToken = token
	//更新数据库的token
	err = models.UpdateUserToken(user); if err != nil{
		c.RespJSON(bean.CODE_Params_Err,err.Error())
		return
	}

	user.Email = v.Email
	user.Pwd = v.Pwd
	user.AccessToken = token

	user.Pwd = ""
	v.Pwd = ""


	//expireCookie := time.Now().Add(time.Minute * 3)
	//cookie := http.Cookie{Name:"Auth",Value:token,Expires:expireCookie,HttpOnly:true,Path:"/"}
	//http.SetCookie(c.Ctx.ResponseWriter,&cookie)
	//c.Ctx.SetCookie("Auth",token)
	//name := "yin"
	//go c.Ctx.SetCookie("Auth",name,expireCookie,"/")

	c.AllowCross()
	//client := &http.Client{}
	//req, err := http.NewRequest("Post","http://127.0.0.1:8088/",nil)
	//req.Header.Add("Auth",token)
	//resp, err := client.Do(req)
	//defer resp.Body.Close()
	//c.Ctx.Output.Header("Auth",token)
	c.Ctx.Output.Header("Auth",token)

	c.RespJSON(http.StatusOK,bean.OutPutSession{Uid:user.Id,Token:user.AccessToken,Email:user.Email,Name:user.Name})

	//c.RespJSONData(user)
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
	 token, err := client.SetToken(u.Email); if err != nil {
		c.RespJSON(bean.CODE_Not_Acceptable,"can't not create token")
		return
	}

	var user models.User
	user.Email = u.Email
	user.Pwd = u.Pwd
	user.AccessToken = token
	//注册用户
	if _,err := models.AddUser(&user); err != nil{
		c.RespJSON(bean.CODE_Params_Err,err.Error())
		return
	}

	//expireCookie := time.Now().Add(time.Minute * 5)
	//
	//cookie := http.Cookie{Name:"Auth",Value:token,Expires:expireCookie,HttpOnly:true}
	//http.SetCookie(c.Ctx.ResponseWriter,&cookie)
	c.RespJSON(http.StatusOK,bean.OutPutSession{Uid:user.Id,Token:user.AccessToken,Email:user.Email})

}
