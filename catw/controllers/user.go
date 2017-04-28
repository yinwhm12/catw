package controllers

import (
	"yinwhm.com/yin/catw/models"

	"yinwhm.com/yin/catw/models/bean"
	"strings"
	"strconv"
)

// Operations about Users
type UserController struct {
	BaseController
}


// @Description 点赞  参数flag = (yes---增加   no---删除)
// @router /sayUp/:id [put]
func (u *UserController)SayUpArticles()  {
	flagStr := u.Ctx.Input.Param("flag")
	idStr := u.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	user, err := models.GetUserById(u.Uid())
	if err != nil{
		u.RespJSON(bean.CODE_Forbidden,err.Error())
		return
	}
	//获取点赞string
	upArticleStr, err := models.GetUpArticlesById(user.Id)
	if flagStr == "yes"{//增加

		upArticleStr +=","+idStr
		err = models.UpdateUpArticles(user.Id,upArticleStr)
		if err != nil{
			u.RespJSON(bean.CODE_Params_Err,err.Error())
			return
		}
		err = models.AddOneByUpById(id)//文章点赞次数减一
		if err != nil{
			u.RespJSON(bean.CODE_Params_Err,err.Error())
			return
		}
		u.RespJSONData("点赞成功!")

	}else if flagStr == "no"{
		if strings.Contains(upArticleStr,","+idStr){
			ss := strings.Replace(upArticleStr,","+idStr,"",1)
			err = models.UpdateUpArticles(user.Id,ss)
			if err != nil{
				u.RespJSON(bean.CODE_Params_Err,err.Error())
				return
			}
			err = models.DeletOneByUpId(id)//点赞次数增加
			if err != nil{
				u.RespJSON(bean.CODE_Params_Err,err.Error())
				return
			}

		}else {
			u.RespJSON(bean.CODE_Forbidden,"错误请求!")
			return
		}
		u.RespJSONData("取消点赞成功!")
	}else {
		u.RespJSON(bean.CODE_Params_Err, "接口有误!")
		return
	}
}

// @Description 点击收藏 参数flag  = (yes---增加   no---删除)
// @router /sayCollect/:id [put]
func (u *UserController)SayCollect()  {
	flagStr := u.Ctx.Input.Param("flag")
	idStr := u.Ctx.Input.Param(":id")
	id,_ := strconv.Atoi(idStr)
	user, err := models.GetUserById(u.Uid())
	if err != nil{
		u.RespJSON(bean.CODE_Forbidden,err.Error())
		return
	}
	//获取收藏string
	collectArticles, err := models.GetCollectArticles(user.Id)
	if flagStr == "yes"{//增加
		collectArticles +=","+idStr
		err = models.UpdateCollectArticles(user.Id,collectArticles)
		if err != nil{
			u.RespJSON(bean.CODE_Params_Err,err.Error())
			return
		}
		err = models.AddOneByCollectId(id)//收藏次数增加  文章
		if err != nil{
			u.RespJSON(bean.CODE_Params_Err,err.Error())
			return
		}
		u.RespJSONData("收藏成功!")

	}else if flagStr == "no"{
		if strings.Contains(collectArticles,","+idStr){
			ss := strings.Replace(collectArticles,","+idStr,"",1)
			err = models.UpdateCollectArticles(user.Id,ss)
			if err != nil{
				u.RespJSON(bean.CODE_Params_Err,err.Error())
				return
			}
			err = models.DeletOneByCollectId(id) //文章收藏次数 减少
			if err != nil{
				u.RespJSON(bean.CODE_Params_Err, err.Error())
				return
			}

		}else {
			u.RespJSON(bean.CODE_Forbidden,"错误请求!")
			return
		}
		u.RespJSONData("已移除收藏!")
	}else {
		u.RespJSON(bean.CODE_Params_Err, "接口有误!")
		return
	}
}

// @Description 判断是否已经点过赞 已点过赞的话返回yes 否则no
// @router /getUpState/:id [get]
func (u *UserController)GetUpState()  {
	idStr := u.Ctx.Input.Param(":id")
	user, err := models.GetUserById(u.Uid())
	if err != nil{
		u.RespJSON(bean.CODE_Forbidden,err.Error())
		return
	}
	upArticles, err := models.GetUpArticlesById(user.Id)
	if err != nil{
		u.RespJSON(bean.CODE_Params_Err, err.Error())
		return
	}
	if strings.Contains(upArticles,idStr+","){
		u.RespJSONData("yes")
		return
	}else {
		u.RespJSONData("no")
		return
	}
}

// @Description 获取用户点赞过的文章
// @router /getUpArticles [get]
func (u *UserController)GetUpArticles()  {

}

// @Description 获取用户收藏的文章
// @router /getCollectArticles [get]
func (u *UserController)GetCollectArticles()  {

}