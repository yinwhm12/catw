package controllers

import (
	"yinwhm.com/yin/catw/client"
	"encoding/json"
	"yinwhm.com/yin/catw/models/bean"
	"yinwhm.com/yin/catw/models"
	"strconv"
	"sort"
	"yinwhm.com/yin/catw/utils"
	"yinwhm.com/yin/catw/tool"
)

type RespondOneController struct {
	BaseController
}

func (c *RespondOneController)URLMapping()  {

}

// Post...
// @Description 添加一级评论
// @router / [post]
func (c *RespondOneController)Post()  {
	var respondJSON client.RespondJSON
	if err := json.Unmarshal(c.Ctx.Input.RequestBody,&respondJSON); err != nil{
		c.RespJSON(bean.CODE_Forbidden,err.Error())
		return
	}else{
		//获取评论者信息
		user, err := models.GetUserById(c.Uid())
		if err != nil{
			c.RespJSON(bean.CODE_Forbidden, err.Error())
			return
		}
		//获得文章信息
		article, err := models.GetArticleById(respondJSON.ArticleId)
		if err != nil{
			c.RespJSON(bean.CODE_Forbidden, err.Error())
			return
		}
		respondOne := models.RespondOne{
			TextContent:respondJSON.TextContent,
			User:user,
			Article:article,

		}
		err = models.AddRespondOne(&respondOne)
		if err != nil{
			c.RespJSON(bean.CODE_Forbidden, err.Error())
			return
		}
		//文章评论数据增加
		err = models.AddOneByCommentById(article.ValueArticle.ValueArticleId)
		if err != nil{
			c.RespJSON(bean.CODE_Forbidden, err.Error())
			return
		}
		c.RespJSON(bean.CODE_Success, "评论成功!")

	}

}

// GetAll ...
// @Description 获取谋篇文章具体的所有的一级评论  参数为文章id
// @router /getAll/:id [get]
func (c *RespondOneController)GetAllResponds()  {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	rOnes, err := models.GetAllRespondOneByArticleId(id)
	if err != nil{
		c.RespJSON(bean.CODE_Params_Err, err.Error())
		return
	}
	length := len(rOnes)
	if length ==0{
		c.RespJSONData("no_respond")
		return
	}

	//建立管道  获取一级评论对应下的各个二级评论
	flag := make(chan int,1)
	//getTwos := make(chan []*models.RespondTwo,length)
	//获取二级评论 启动线程
	//GetRespondTwos(rOnes,getTwos)
	go tool.GetRespondTwos(&rOnes,flag)

	//获取评论者信息
	Dlinks := make([]int,length)
	for _, s := range rOnes{
		//fmt.Println("=========",s.User.Id)
		Dlinks = append(Dlinks,s.User.Id)
		//查询数据库 获取二级  进行线程查询
		//go   models.GetAllRespondTwoByROne(&s)

	}
	//过滤 重复作者 使作者数组减轻
	sort.Ints(Dlinks)
	links := utils.Duplicate(Dlinks)
	//获取用户信息
	//fmt.Println("--------",links)
	userMap, err := models.GetUsersByIds(links)
	if err != nil{
		c.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	//进行相应的 user赋值
	for i, s := range rOnes{
		if u, ok := userMap[s.User.Id]; ok{
			//fmt.Println("usermap------",u)
			rOnes[i].User =  &u
		}
	}
	//管道 取值赋给 respondOne
	//for i,s := range getTwos{
	//	rOnes[s[i].RespondOne.RespondOneId].RespondTwos = s
	//
	//}
	getFlag := <- flag //获取管道值 是否 已添加完成
	if getFlag == length{

		c.RespJSONData(rOnes)
		//for _, s := range rOnes{
		//	fmt.Println("-------data----",s.RespondTwos)
		//	if len(s.RespondTwos) != 0{
		//		fmt.Println("------len-----",len(s.RespondTwos))
		//		for _, r:=range s.RespondTwos{
		//			fmt.Println("--ssss----",r.User)
		//			fmt.Println("--ssss----",r)
		//		}
		//	}
		//}
	}else {
		c.RespJSON(bean.CODE_Forbidden,"数据不完整!")
	}
	
}

// @router /:id [get]
func (c *RespondOneController)Get()  {
	
}

