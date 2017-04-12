package controllers

import (
	"yinwhm.com/yin/catw/models"
	"yinwhm.com/yin/catw/models/bean"
)

type Root1TypeController struct {
	BaseController
}

func (c *Root1TypeController)URLMapping()  {

}

// Get ...
// @Title 获取文章、帖子类型
// @Success 200
// @Failure 403
// @router / [get]
func (c *Root1TypeController)Get()  {
	//flag := c.GetString("flag")
	//if flag != "root1"{
	//	c.RespJSON(bean.CODE_Forbidden,"请求错误!")
	//	return
	//}
	type1s, err := models.GetAllRoot1TypeInfo()
	if err != nil{
		c.RespJSON(bean.CODE_Forbidden,err.Error())
		return
	}
	c.RespJSONData(type1s)


}

