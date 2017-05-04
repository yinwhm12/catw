package controllers

import (
	"yinwhm.com/yin/catw/models"
	"yinwhm.com/yin/catw/models/bean"
	"yinwhm.com/yin/catw/client"
	"encoding/json"
	"sort"
	"yinwhm.com/yin/catw/utils"
	"sync"
	"yinwhm.com/yin/catw/tool"
)

type MessageController struct {
	BaseController
}

func (c *MessageController)URLMapping()  {
	
}

// @Description 一级回复 仅向楼主通知
// @router /toOwner [post]
func (c *MessageController)PostOne()  {
	var messageJSON client.MessageJSON
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &messageJSON); err !=nil{
		c.RespJSON(bean.CODE_Params_Err, err.Error())
		return
	}else{
		//接收者信息 楼主
		to_user, err := models.GetUserById(messageJSON.ToUserID)
		if err != nil{
			c.RespJSON(bean.CODE_Forbidden, err.Error())
			return
		}
		//发送者信息
		from_user, err := models.GetUserById(c.Uid())
		if err != nil{
			c.RespJSON(bean.CODE_Forbidden, err.Error())
			return
		}
		article, err := models.GetArticleById(messageJSON.ArticleId)
		if err != nil{
			c.RespJSON(bean.CODE_Forbidden, err.Error())
			return
		}
		message := models.Message{
			ToUser:to_user,
			FromUser:from_user,
			Article:article,
		}
		//保存这条信息
		err = models.AddOneMessage(&message)
		if err != nil{
			c.RespJSON(bean.CODE_Forbidden, err.Error())
			return
		}

	}

}

// @Description 二级回复 进行所有参与二级的都通知
// @router /toMany [post]
func (c *MessageController)PostMany()  {
	var messageJSON client.MessageJSON
	if err := json.Unmarshal(c.Ctx.Input.RequestBody,&messageJSON); err != nil{
		c.RespJSON(bean.CODE_Params_Err, err.Error())
		return
	}else {
		article, err := models.GetArticleById(messageJSON.ArticleId)
		if err != nil{
			c.RespJSON(bean.CODE_Forbidden, err.Error())
			return
		}
		from_user, err := models.GetUserById(c.Uid())//顶者
		if err != nil{
			c.RespJSON(bean.CODE_Forbidden, err.Error())
			return
		}

		userIds, err := models.GetAllUserIdsByRespondOne(messageJSON.RespondOneId)
		if err != nil{
			c.RespJSON(bean.CODE_Forbidden, err.Error())
			return
		}
		sort.Ints(userIds)
		links := utils.Duplicate(userIds)//过滤重复的userid


		message := models.Message{
			Article:article,
			FromUser:from_user,
		}
		var wg sync.WaitGroup


		//查找用户信息并进行插入数据
		for i:=0; i<len(links); i++{
			wg.Add(1)
			var id int
			switch v := links[i].(type) {
				case int:
					id = v
				default:
					continue
			}
			go tool.DealAddManyMessage(id,message,&wg)
			//user, err := models.GetUserById(id)
			//if err != nil{//其中 获取某个用户的信息 出问题时 直接跳出
			//	c.RespJSON(bean.CODE_Forbidden,err.Error())
			//	return
			//}

			//err = models.AddOneMessage(&message)
			//if err != nil{
			//	c.RespJSON(bean.CODE_Forbidden, err.Error())
			//	return
			//}

		}
		wg.Wait()


	}
}

// @Description 分页界面 同一入口 参入page=0未读 page=1已读 page=3全部
// @router / [get]
func (c *MessageController)GetAll()  {
	page, err := c.GetInt("page")
	if err != nil{
		c.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	limit, err := c.GetInt("limit")
	if err != nil{
		c.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	offset, err := c.GetInt("offset")
	if err != nil{
		c.RespJSON(bean.CODE_Forbidden,err.Error())
		return
	}
	user, err := models.GetUserById(c.Uid())
	if err != nil{
		c.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	var total int64
	var messages []*models.Message
	if page == 0{//未读页面
		messages, total, err = models.GetMessageUnreadPageByToUser(user,limit, offset)
	}else if page == 1{//已读页面
		messages, total, err = models.GetMessageReadPageByToUser(user, limit, offset)
	}else if page == 3{//全部页面
		messages, total, err = models.GetAllMessagesByToUser(user, limit, offset)
	}else{//page 参数有误
		c.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	if err != nil{
		c.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}else {
		c.RespJSONDataWithTotal(messages, total)
	}
}

// @Description 获取首页 未读通知数量标记
// @router /getNews [get]
func (c *MessageController)GetNews()  {
	user, err := models.GetUserById(c.Uid())
	if err != nil{
		c.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	total, err := models.GetMessageUnreadCountByToUser(user)
	if err != nil{
		c.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	c.RespJSONData(total)
}
