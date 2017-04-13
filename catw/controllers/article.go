package controllers

import (
	"yinwhm.com/yin/catw/models"
	"yinwhm.com/yin/catw/models/bean"
	"strconv"
	"fmt"
	"yinwhm.com/yin/catw/client"
	"encoding/json"
)

type ArticleController struct {
	BaseController
}

func (c *ArticleController)URLMapping()  {

}

// Post ...
// @Title 写文章、帖子
// @Description
// @Param body 		models.User models.EndType models.Article
// @Success 200
// @Failure 403
// @router / [post]
func (c *ArticleController)Post()  {

	fmt.Println("uuuuu",c.Uid())
	var articleJSON client.ArticleJSON
	if err := json.Unmarshal(c.Ctx.Input.RequestBody,&articleJSON); err != nil{
		c.RespJSON(bean.CODE_Forbidden,err.Error())
		return
	}else {
		fmt.Println("articleJSON----",articleJSON)
		user, err := models.GetUserById(c.Uid())
		if err != nil{
			c.RespJSON(bean.CODE_Forbidden,"用户数据有误，请重登录!")
			return
		}
		fmt.Println("user--",user)
		//endType,err :=models.GetEndTypeInfoByAllFK(articleJSON.ArticleRoot1,articleJSON.ArticleRoot2,articleJSON.ArticleLevel)
		//if err != nil{
		//	c.RespJSON(bean.CODE_Forbidden,err.Error())
		//	return
		//}
		//fmt.Println("eeee--",endType)
		//var article  models.Article
		//article.EndType = &endType
		//article.User = user
		//article.TextContent = articleJSON.ArticleContent
		//article.Title = articleJSON.ArticleTitle
		//if _,err = models.AddArticle(&article);err != nil{
		//	c.RespJSON(bean.CODE_Forbidden,err.Error())
		//	return
		//}


	}
	c.RespJSON(bean.CODE_Success,"OK")
	//var article models.Article
	//if err := json.Unmarshal(c.Ctx.Input.RequestBody,&article); err != nil{
	//	c.RespJSON(bean.CODE_Forbidden,err.Error())
	//	return
	//}else{//参数 有效
	//
	//	if _,err := models.AddArticle(&article); err == nil{
	//		c.RespJSONData("OK")
	//	}else {
	//		c.RespJSON(bean.CODE_Forbidden,err.Error())
	//	}
	//}
}

// Get ...
// @Title 具体文章 帖子
// @Success 200
// @Failure 403
// @router /:id [get]
func (c *ArticleController)GetOne()  {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	article, err := models.GetArticleById(id)
	if err != nil{
		c.RespJSON(bean.CODE_Forbidden,err.Error())
	}else{
		c.RespJSONData(article)
	}
}

// Get All
// @Title 取某个类型(root1 root2 levelType)的全部内容
// @Params  type 类型
// @router /type [get]
func (c *ArticleController)GetType()  {
	idStr := c.Ctx.Input.Param(":id")
	type_id, _ := strconv.Atoi(idStr)

	typeStr := c.GetString("type")
	var articles []models.Article
	//articles := []models.Article{}
	var err error
	if typeStr == "root1"{
		articles, err = models.GetAticlesByRoot1Id(type_id)
		if err != nil{
			c.RespJSON(bean.CODE_Forbidden,err.Error())
			return
		}

	}else if typeStr == "root2"{
		articles, err = models.GetAticlesByRoot2Id(type_id)
		if err != nil{
			c.RespJSON(bean.CODE_Forbidden,err.Error())
			return
		}
	}else if typeStr == "level"{
		articles, err = models.GetAticlesByLevelTypeId(type_id)
		if err != nil{
			c.RespJSON(bean.CODE_Forbidden, err.Error())
			return
		}
	}else {
		c.RespJSON(bean.CODE_Forbidden,"bad wrong")
		return
	}
	c.RespJSONData(articles)
}




