package controllers

import (
	"yinwhm.com/yin/catw/client"
	"encoding/json"
	"yinwhm.com/yin/catw/models/bean"
	"yinwhm.com/yin/catw/models"
)

type RespondTwoController struct {
	BaseController
}

func (c *RespondTwoController)URLMapping()  {

}

// @Description 添加一条新的二级评论
// @router / [post]
func (c *RespondTwoController)Post()  {
	var respondTwoJSON client.RespondTwoJSON
	if err := json.Unmarshal(c.Ctx.Input.RequestBody,&respondTwoJSON); err != nil{
		c.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}else {
		//检验 数据
		if respondTwoJSON.TextContent == "" || respondTwoJSON.RespondOne == 0{
			c.RespJSON(bean.CODE_Params_Err, "参数有误！")
			return
		}
		user, err := models.GetUserById(c.Uid())
		if err != nil{
			c.RespJSON(bean.CODE_Forbidden, err.Error())
			return
		}
		//获取 一级评论信息
			rOne, err := models.GetOneRespondById(respondTwoJSON.RespondOne)
		if err != nil{
			c.RespJSON(bean.CODE_Forbidden, err.Error())
			return
		}
		//判断 是否文章有效
		artcle, err := models.GetArticleById(rOne.Article.Tid)
		if err != nil{
			c.RespJSON(bean.CODE_Forbidden, err.Error())
			return
		}
		respondTwo := models.RespondTwo{
			TextContent:respondTwoJSON.TextContent,
			User:user,
			RespondOne:rOne,
		}
		err = models.AddRespondTwo(&respondTwo)
		if err != nil{
			c.RespJSON(bean.CODE_Forbidden, err.Error())
			return
		}
		//点评数增加  通过文章article 获取价值表的id 从而使评论数加一
		err = models.AddOneByCommentById(artcle.ValueArticle.ValueArticleId)
		if err != nil{
			c.RespJSON(bean.CODE_Forbidden, err.Error())
			return
		}
		c.RespJSON(bean.CODE_Success,"顶起成功!")
	}
	
}
