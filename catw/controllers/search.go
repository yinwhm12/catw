package controllers

import (
	"yinwhm.com/yin/catw/models/bean"
	"yinwhm.com/yin/catw/models"
	"yinwhm.com/yin/catw/tool"
	"yinwhm.com/yin/catw/utils"
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
	c.RespJSONDataWithTotal(resultArticles, int64(length))
	return


}