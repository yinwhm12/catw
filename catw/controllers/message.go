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
	"strconv"
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
		//to_user, err := models.GetUserById(messageJSON.ToUserID)
		//if err != nil{
		//	c.RespJSON(bean.CODE_Forbidden, err.Error())
		//	return
		//}
		//发送者信息
		from_user, err := models.GetUserById(c.Uid())
		if err != nil{
			c.RespJSON(bean.CODE_Forbidden, err.Error())
			return
		}
		//通过文章id 获取作者id
		article, err := models.GetArticleById(messageJSON.ArticleId)
		if err != nil{
			c.RespJSON(bean.CODE_Forbidden, err.Error())
			return
		}
		to_user, err := models.GetUserById(article.User.Id)
		if err != nil{
			c.RespJSON(bean.CODE_Forbidden, err.Error())
			return
		}
		//消息内容
		content := from_user.Email+ "对你的<<"+article.Title+">>文章进行评论!"
		message := models.Message{
			ToUser:to_user,
			FromUser:from_user,
			Article:article,
			Content:content,
		}
		//保存这条信息
		err = models.AddOneMessage(&message)
		if err != nil{
			c.RespJSON(bean.CODE_Forbidden, err.Error())
			return
		}
		c.RespJSON(bean.CODE_Success,"操作成功!")

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

		//获取作者
		auther, err := models.GetUserById(article.User.Id)
		if err != nil{
			c.RespJSON(bean.CODE_Forbidden,err.Error())
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


		content := from_user.Email+"顶起你的评论,于<<"+article.Title+">>"
		message := models.Message{
			Article:article,
			FromUser:from_user,
			Content:content,
		}

		//向作者发送信息
		message.ToUser = auther
		go models.AddOneMessage(&message)

		var wg sync.WaitGroup


		//查找用户信息并进行插入数据
		for i:=0; i<len(links); i++{
			wg.Add(1)
			var id int
			switch v := links[i].(type) {
				case int:
					id = v
					if id == c.Uid() || id == article.User.Id{//除去自己 除去作者
						wg.Done()
						continue
					}
					break
				default:
					wg.Done()
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

		c.RespJSON(bean.CODE_Success,"操作成功!")


	}
}

// @Description 分页界面 同一入口 参入page=1未读 page=2已读 page=3全部
// @router / [get]
func (c *MessageController)GetAll()  {
	page := c.GetString("page")
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
	if page == "1"{//未读页面
		messages, total, err = models.GetMessageUnreadPageByToUser(user,limit, offset)
	}else if page == "2"{//已读页面
		messages, total, err = models.GetMessageReadPageByToUser(user, limit, offset)
	}else if page == "3"{//全部页面
		messages, total, err = models.GetAllMessagesByToUser(user, limit, offset)
	}else{//page 参数有误
		c.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	if err != nil{
		c.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}else {
		//获取发送者信息
		length := len(messages)
		if length <=0{
			total = 0
		}else{

			Dlinks := make([]int, length)
			for _, s := range messages{
				Dlinks = append(Dlinks,s.FromUser.Id)
			}
			sort.Ints(Dlinks)
			links := utils.Duplicate(Dlinks)
			//获取作者
			userMap, err := models.GetUsersByIds(links)
			if err != nil{
				c.RespJSON(bean.CODE_Params_Err, err.Error())
				return
			}
			for i, s := range messages{
				if u, ok := userMap[s.FromUser.Id]; ok{
					messages[i].FromUser = &u
				}
			}
		}
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

// @Description 删除某条具体的信息
// @router /:id [delete]
func (c *MessageController)Delete()  {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil{
		c.RespJSON(bean.CODE_Bad_Request, err.Error())
		return
	}
	err = models.DeletOneMessageById(id)
	if err != nil{
		c.RespJSON(bean.CODE_Bad_Request, err.Error())
		return
	}else {
		c.RespJSONData("成功删除!")
		return
	}
}

// @Description 标记已读信息
// @router /hadReadMessage/:id [put]
func (c *MessageController)HadReadMessage()  {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil{
		c.RespJSON(bean.CODE_Bad_Request, err.Error())
		return
	}
	err = models.ChangeMessageStateById(id,1)//将状态改为已读
	if err != nil{
		c.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}else {
		c.RespJSON(bean.CODE_Success,"操作成功")
		return
	}
}
