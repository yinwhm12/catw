package controllers

import (
	"yinwhm.com/yin/catw/models"
	"yinwhm.com/yin/catw/models/bean"
	"strconv"
	"yinwhm.com/yin/catw/client"
	"encoding/json"
	"fmt"
	"sort"
	"yinwhm.com/yin/catw/utils"
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

	//c.AllowCross()
	//token := c.Ctx.Input.Header("Authorization")
	//if token == "" {
	//	token = c.Ctx.Input.Query("_token")
	//}
	//fmt.Println("====",token)
	//oo:=c.Ctx.Request.Header.Get("Authorization")
	//fmt.Println("00000000",oo)

	var articleJSON client.ArticleJSON
	if err := json.Unmarshal(c.Ctx.Input.RequestBody,&articleJSON); err != nil{
		c.RespJSON(bean.CODE_Forbidden,err.Error())
		return
	}else {
		user, err := models.GetUserById(c.Uid())
		if err != nil{
			c.RespJSON(bean.CODE_Forbidden,"用户数据有误，请重登录!")
			return
		}
		endType,err :=models.GetEndTypeInfoByAllFK(articleJSON.ArticleRoot1,articleJSON.ArticleRoot2,articleJSON.ArticleLevel)
		if err != nil{
			c.RespJSON(bean.CODE_Forbidden,err.Error())
			return
		}
		var article  models.Article
		article.EndType = &endType
		article.User = user
		article.TextContent = articleJSON.ArticleContent
		article.Title = articleJSON.ArticleTitle
		if _,err = models.AddArticle(&article);err != nil{
			c.RespJSON(bean.CODE_Forbidden,err.Error())
			return
		}
	}
	c.RespJSON(bean.CODE_Success,"OK")

}

// Get ...
// @Title 具体文章 帖子
// @Param	id		path 	string	true
// @Success 200
// @Failure 403
// @router /getOne/:id [get]
func (c *ArticleController)GetOne()  {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	fmt.Println("-----ssss",id)
	article, err := models.GetArticleById(id)
	if err != nil{
		c.RespJSON(bean.CODE_Forbidden,err.Error())
	}else{
		if err = models.GetUserByUser(article.User);err != nil{
			c.RespJSON(bean.CODE_Params_Err,err.Error())
			return
		}
		article.User.Pwd=""
		//fmt.Println("---user",article.User)
		//获取作者信息 评价等信息

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

// Get ...
// @Title 最新主题(课间操)
// @Description 获取最新的 课间操的主题 按时间排序 取最新 9 个
// @Param	id		path 	string	true
// @Param	body		body 		true
// @router /getNine/:id [get]
func (c *ArticleController)GetPlayTheme()  {
	idStr := c.Ctx.Input.Param(":id" )
	id, _ := strconv.Atoi(idStr)
	articles, err := models.GetPalyThemeIndex(id)
	if err != nil{
		c.RespJSON(bean.CODE_Params_Err,err.Error())
		return
	}
	if len(articles) == 0{
		c.RespJSON(bean.CODE_Params_Err,"数据不存在!")
		return
	}
	c.RespJSONData(articles)
}

// @Description 主页面获取 文章 具体的5个
// @router /getThemes/:id [get]
func (c *ArticleController)GetThemesIndex()  {
	idStr := c.Ctx.Input.Param(":id")
	flag := c.GetString("flag")
	id, _ := strconv.Atoi(idStr)
	fmt.Println("----ffff",flag)
	articles, err := models.GetThemesByRoot1Id(flag,id)
	if err != nil{
		c.RespJSON(bean.CODE_Params_Err,"暂无数据!")
		return
	}
	length := len(articles)
	if length == 0{
		c.RespJSON(bean.CODE_Params_Err,"暂无数据!")
		return
	}
	Dlinks := make([]int,length)
	for i, s := range articles{
		Dlinks[i] = s.User.Id
	}
	sort.Ints(Dlinks)
	links := utils.Duplicate(Dlinks)
	//获取作者
	userMap, err := models.GetUsersByIds(links)
	if err != nil{
		c.RespJSON(bean.CODE_Params_Err,err.Error())
		return
	}
	for i, s := range articles{
		if u, ok  := userMap[s.User.Id]; ok{
			articles[i].User = &u
		}
	}

	c.RespJSONData(articles)
}

// @Description 获取全部 参数类型两种类型 其一为root1 其二为root2
// @router /getAll [get]
func (c *ArticleController)GetAll()  {
	var mapIds map[string]int
	flagBool := 0
	mapIds = make(map[string]int)
	root1,err  := c.GetInt("root1")
	if err == nil{
		flagBool++
		mapIds["root1"] = root1
	}
	root2, err := c.GetInt("root2")
	if err == nil{
		flagBool++
		mapIds["root2"] = root2
	}
	if flagBool == 0{
		c.RespJSON(bean.CODE_Params_Err, err.Error())
		return
	}
	limit, err := c.GetInt("limit")
	if err != nil{
		c.RespJSON(bean.CODE_Params_Err, err.Error())
		return
	}
	offset, err := c.GetInt("offset")
	if err != nil {
		c.RespJSON(bean.CODE_Params_Err,err.Error())
		return
	}
	//获取具体的类型 即end_type
	endTypes, err := models.GetEndTypeIds(mapIds)
	if err != nil{
		c.RespJSON(bean.CODE_Params_Err,err.Error())
		return
	}
	links := make([]int, len(endTypes))
	for _, s := range endTypes {
		links = append(links,s.EndTypeId)
	}
	articles, total, err := models.GetIndexAllByPage(links,limit,offset)
	if err != nil{
		c.RespJSON(bean.CODE_Params_Err,err.Error())
		return
	}
	Dlinks := make([]int,len(articles))
	for _,s := range articles{
		Dlinks = append(Dlinks,s.User.Id)
	}
	sort.Ints(Dlinks)
	uLinks := utils.Duplicate(Dlinks)
	//获取作者
	userMap, err := models.GetUsersByIds(uLinks)
	if err != nil{
		c.RespJSON(bean.CODE_Params_Err, err.Error())
		return
	}
	for i, s := range articles{
		if u, ok  := userMap[s.User.Id]; ok{
			articles[i].User = &u
		}
	}
	c.RespJSONDataWithTotal(articles,total)

}


