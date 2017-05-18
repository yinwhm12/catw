package controllers

import (
	"yinwhm.com/yin/catw/models/bean"
	"yinwhm.com/yin/catw/models"
	"yinwhm.com/yin/catw/tool"
	"yinwhm.com/yin/catw/utils"
	"sort"
)

type SearchController struct {
	BaseController
}

func (c *SearchController)URLMapping()  {

}

// @Description 搜索 参数 作者关键字
// @router /searchName [get]
func (c *SearchController)SearchName()  {
	auther := c.GetString("name")
	limit, err := c.GetInt("limit")
	if err != nil{
		c.RespJSON(bean.CODE_Params_Err, err.Error())
		return
	}
	offset, err := c.GetInt("offset")
	if err != nil{
		c.RespJSON(bean.CODE_Params_Err, err.Error())
		return
	}
	articles, total, err := models.SearchAuthers(auther, limit,offset)
	if err != nil{
		c.RespJSON(bean.CODE_Params_Err, err.Error())
		return
	}
	length := len(articles)
	//fmt.Println("----total----",length)
	if length <= 0{
		total = 0
	}else {
		Dlinks := make([]int,length)
		vlinks := make([]int, length)
		for _,s := range articles{
			Dlinks = append(Dlinks,s.User.Id)
			vlinks = append(vlinks,s.ValueArticle.ValueArticleId)//价值
		}
		sort.Ints(Dlinks)
		uLinks := utils.Duplicate(Dlinks)
		//获取作者
		userMap, err := models.GetUsersByIds(uLinks)
		if err != nil{
			c.RespJSON(bean.CODE_Params_Err, err.Error())
			return
		}
		//获取价值
		valueMap, err := models.GetAllValueByIds(vlinks)
		if err != nil{
			c.RespJSON(bean.CODE_Params_Err, err.Error())
			return
		}

		for i, s := range articles{
			//用户 赋值
			if u, ok  := userMap[s.User.Id]; ok{
				articles[i].User = &u
			}
			//价值 赋值
			if v, ok := valueMap[s.ValueArticle.ValueArticleId]; ok{
				articles[i].ValueArticle = &v
			}
		}
	}
	c.RespJSONDataWithTotal(articles, total)
}

// @Description 按标题搜索 不分大小写
// @router /searchTitle [get]
func (c *SearchController)SearchTitle()  {
	title := c.GetString("title")
	limit, err := c.GetInt("limit")
	if err != nil{
		c.RespJSON(bean.CODE_Params_Err, err.Error())
		return
	}
	offset, err := c.GetInt("offset")
	if err != nil{
		c.RespJSON(bean.CODE_Params_Err, err.Error())
		return
	}
	articles, total, err := models.SearchTitle(title,limit,offset)
	if err != nil{
		c.RespJSON(bean.CODE_Params_Err, err.Error())
		return
	}
	length := len(articles)
	if length <=0{
		total = 0

	}else {
		Dlinks := make([]int,len(articles))
		vlinks := make([]int, len(articles))
		for _,s := range articles{
			Dlinks = append(Dlinks,s.User.Id)
			vlinks = append(vlinks,s.ValueArticle.ValueArticleId)//价值
		}
		sort.Ints(Dlinks)
		uLinks := utils.Duplicate(Dlinks)
		//获取作者
		userMap, err := models.GetUsersByIds(uLinks)
		if err != nil{
			c.RespJSON(bean.CODE_Params_Err, err.Error())
			return
		}
		//获取价值
		valueMap, err := models.GetAllValueByIds(vlinks)
		if err != nil{
			c.RespJSON(bean.CODE_Params_Err, err.Error())
			return
		}

		for i, s := range articles{
			//用户 赋值
			if u, ok  := userMap[s.User.Id]; ok{
				articles[i].User = &u
			}
			//价值 赋值
			if v, ok := valueMap[s.ValueArticle.ValueArticleId]; ok{
				articles[i].ValueArticle = &v
			}
		}

	}
	c.RespJSONDataWithTotal(articles, total)
}

// @Description 内容搜索
// @router /searchContent [get]
func (c *SearchController)SearchContent()  {
	keyString := c.GetString("content")

	limit, err := c.GetInt("limit")
	if err != nil{
		c.RespJSON(bean.CODE_Params_Err, err.Error())
		return
	}
	offset, err := c.GetInt("offset")
	if err != nil{
		c.RespJSON(bean.CODE_Params_Err, err.Error())
		return
	}

	//获取全部的内容
	articles, err := models.SearchAllContent()
	if err != nil{
		c.RespJSON(bean.CODE_Params_Err, err.Error())
		return
	}

	//获取 符合的ids
	ids := tool.Search(articles, keyString)
	length := len(ids)
	if length <= 0{
		c.RespJSON(bean.CODE_Params_Err,"没有响应的数据")
		return
	}
	links := utils.Duplicate(ids)

	resultArticles, err := models.SearchAllContentPageByIds(links, limit, offset)
	if err != nil{
		c.RespJSON(bean.CODE_Params_Err, err.Error())
		return
	}

	resultLength := len(resultArticles)
	if length <= 0{
		length = 0
	}else {
		Dlinks := make([]int,resultLength)
		vlinks := make([]int, resultLength)
		for _,s := range articles{
			Dlinks = append(Dlinks,s.User.Id)
			vlinks = append(vlinks,s.ValueArticle.ValueArticleId)//价值
		}
		sort.Ints(Dlinks)
		uLinks := utils.Duplicate(Dlinks)
		//获取作者
		userMap, err := models.GetUsersByIds(uLinks)
		if err != nil{
			c.RespJSON(bean.CODE_Params_Err, err.Error())
			return
		}
		//获取价值
		valueMap, err := models.GetAllValueByIds(vlinks)
		if err != nil{
			c.RespJSON(bean.CODE_Params_Err, err.Error())
			return
		}

		for i, s := range articles{
			//用户 赋值
			if u, ok  := userMap[s.User.Id]; ok{
				articles[i].User = &u
			}
			//价值 赋值
			if v, ok := valueMap[s.ValueArticle.ValueArticleId]; ok{
				articles[i].ValueArticle = &v
			}
		}
	}
	c.RespJSONDataWithTotal(resultArticles, int64(length))
	return


}