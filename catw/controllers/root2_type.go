package controllers

import (
	"yinwhm.com/yin/catw/models"
	"yinwhm.com/yin/catw/models/bean"
)

type Root2TypeController struct {
	BaseController
}

func (c Root2TypeController)URLMapping()  {

}

// @Get ...
// @Title 类型
// @Description 获取 语文、数学、物理等type
// @Success 200
// @Failure 403
// @router / [get]
func (c *Root2TypeController)Get()  {

	type2s, err := models.GetAllRoot2TypeInfo()
	if err != nil{
		c.RespJSON(bean.CODE_Forbidden,err.Error())
		return
	}
	c.RespJSONData(type2s)
}
