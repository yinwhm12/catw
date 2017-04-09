package controllers

import (
	"yinwhm.com/yin/catw/models"
	"yinwhm.com/yin/catw/models/bean"
)

type LevelTypeController struct {
	BaseController
}

func (c *LevelTypeController)URLMapping()  {

}

// @Get ...
// @Title 类型
// @Description 等级 如大学 小学
// @Success 200
// @Failure 403
// @router / [get]
func (c *LevelTypeController)Get()  {

	levelType, err := models.GetAllLevelTypeInfo()
	if err != nil{
		c.RespJSON(bean.CODE_Forbidden,err.Error())
		return
	}
	c.RespJSONData(levelType)
}