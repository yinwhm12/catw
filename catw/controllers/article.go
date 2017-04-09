package controllers

import (
	"yinwhm.com/yin/catw/models"
	"encoding/json"
	"yinwhm.com/yin/catw/models/bean"
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
	var article models.Article
	if err := json.Unmarshal(c.Ctx.Input.RequestBody,&article); err != nil{
		c.RespJSON(bean.CODE_Forbidden,err.Error())
		return
	}else{//参数 有效

		if _,err := models.AddArticle(&article); err == nil{
			c.RespJSONData("OK")
		}else {
			c.RespJSON(bean.CODE_Forbidden,err.Error())
		}
	}
}
