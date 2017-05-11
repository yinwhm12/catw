package controllers

import (
	"yinwhm.com/yin/catw/models"

	"yinwhm.com/yin/catw/models/bean"
	"strings"
	"strconv"
	"yinwhm.com/yin/catw/client"
	"encoding/json"
	"sort"
	"yinwhm.com/yin/catw/utils"
)

// Operations about Users
type UserController struct {
	BaseController
}


// @Description 点赞  参数flag = (yes---增加   no---删除)
// @router /sayUp/:id [put]
func (u *UserController)SayUpArticles()  {
	flagStr := u.GetString("flag")
	idStr := u.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	user, err := models.GetUserById(u.Uid())
	if err != nil{
		u.RespJSON(bean.CODE_Forbidden,err.Error())
		return
	}
	//获取点赞string
	upArticleStr, err := models.GetUpArticlesById(user.Id)
	if err != nil{
		u.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	//根据id 判断是否有该文章
	artcle, err := models.GetArticleById(id)
	if err != nil{
		u.RespJSON(bean.CODE_Forbidden,err.Error())
		return
	}
	if flagStr == "yes"{//增加
		if strings.Contains(upArticleStr,","+idStr){//已经点过赞了
			u.RespJSON(bean.CODE_Forbidden,"已经点过赞了!")
			return
		}
		upArticleStr +=","+idStr
		err = models.UpdateUpArticles(user.Id,upArticleStr)
		if err != nil{
			u.RespJSON(bean.CODE_Params_Err,err.Error())
			return
		}
		err = models.AddOneByUpById(artcle.ValueArticle.ValueArticleId)//文章点赞次数减一
		if err != nil{
			u.RespJSON(bean.CODE_Params_Err,err.Error())
			return
		}
		u.RespJSONData("点赞成功!")

	}else if flagStr == "no"{
		if strings.Contains(upArticleStr,","+idStr){
			ss := strings.Replace(upArticleStr,","+idStr,"",1)
			err = models.UpdateUpArticles(user.Id,ss)
			if err != nil{
				u.RespJSON(bean.CODE_Params_Err,err.Error())
				return
			}
			err = models.DeletOneByUpId(artcle.ValueArticle.ValueArticleId)//点赞次数增加
			if err != nil{
				u.RespJSON(bean.CODE_Params_Err,err.Error())
				return
			}

		}else {
			u.RespJSON(bean.CODE_Forbidden,"错误请求!")
			return
		}
		u.RespJSONData("取消点赞成功!")
	}else {
		u.RespJSON(bean.CODE_Params_Err, "接口有误!")
		return
	}
}

// @Description 点击收藏 参数flag  = (yes---增加   no---删除)
// @router /sayCollect/:id [put]
func (u *UserController)SayCollect()  {
	flagStr := u.GetString("flag")
	idStr := u.Ctx.Input.Param(":id")
	id,_ := strconv.Atoi(idStr)
	user, err := models.GetUserById(u.Uid())
	if err != nil{
		u.RespJSON(bean.CODE_Forbidden,err.Error())
		return
	}
	//获取收藏string
	collectArticles, err := models.GetCollectArticles(user.Id)
	if err != nil{
		u.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	//根据id 判断是否有该文章
	artcle, err := models.GetArticleById(id)
	if err != nil{
		u.RespJSON(bean.CODE_Forbidden,err.Error())
		return
	}
	if flagStr == "yes"{//增加
		if strings.Contains(collectArticles,","+idStr){//已经点过赞了
			u.RespJSON(bean.CODE_Forbidden,"已经收藏了!")
			return
		}
		collectArticles +=","+idStr
		err = models.UpdateCollectArticles(user.Id,collectArticles)
		if err != nil{
			u.RespJSON(bean.CODE_Params_Err,err.Error())
			return
		}
		err = models.AddOneByCollectId(artcle.ValueArticle.ValueArticleId)//收藏次数增加  文章
		if err != nil{
			u.RespJSON(bean.CODE_Params_Err,err.Error())
			return
		}
		u.RespJSONData("收藏成功!")

	}else if flagStr == "no"{
		if strings.Contains(collectArticles,","+idStr){
			ss := strings.Replace(collectArticles,","+idStr,"",1)
			err = models.UpdateCollectArticles(user.Id,ss)
			if err != nil{
				u.RespJSON(bean.CODE_Params_Err,err.Error())
				return
			}
			err = models.DeletOneByCollectId(artcle.ValueArticle.ValueArticleId) //文章收藏次数 减少
			if err != nil{
				u.RespJSON(bean.CODE_Params_Err, err.Error())
				return
			}

		}else {
			u.RespJSON(bean.CODE_Forbidden,"错误请求!")
			return
		}
		u.RespJSONData("已移除收藏!")
	}else {
		u.RespJSON(bean.CODE_Params_Err, "接口有误!")
		return
	}
}

// @Description 点击关注 参数flag  = (no---增加   yes---删除)
// @router /sayUser/:id [put]
func (u *UserController)SayUser()  {
	flagStr := u.GetString("flag")
	idStr := u.Ctx.Input.Param(":id")
	id,_ := strconv.Atoi(idStr)
	user, err := models.GetUserById(u.Uid())
	if err != nil{
		u.RespJSON(bean.CODE_Forbidden,err.Error())
		return
	}
	//获取关注string
	collectUsers, err := models.GetCollectUsersById(user.Id)
	if err != nil{
		u.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	//根据id 判断是否有该用户
	usered, err := models.GetUserById(id)
	if err != nil{
		u.RespJSON(bean.CODE_Forbidden,err.Error())
		return
	}
	if flagStr == "no"{//增加
		if strings.Contains(collectUsers,","+idStr){//已经关注过了
			u.RespJSON(bean.CODE_Forbidden,"已经关注过了!")
			return
		}
		collectUsers +=","+idStr
		err = models.UpdateCollectUsers(user.Id,collectUsers)
		if err != nil{
			u.RespJSON(bean.CODE_Params_Err,err.Error())
			return
		}
		//err = models.AddOneByCollectId(usered.ValueArticle.ValueArticleId)//收藏次数增加  文章
		err = models.UpdateUserFan(id,usered.Fan+1)//粉丝数量加一
		if err != nil{
			u.RespJSON(bean.CODE_Params_Err,err.Error())
			return
		}
		u.RespJSONData("关注成功!")

	}else if flagStr == "yes"{
		if strings.Contains(collectUsers,","+idStr){
			ss := strings.Replace(collectUsers,","+idStr,"",1)
			err = models.UpdateCollectUsers(user.Id,ss)
			if err != nil{
				u.RespJSON(bean.CODE_Params_Err,err.Error())
				return
			}
			//err = models.DeletOneByCollectId(artcle.ValueArticle.ValueArticleId) //文章收藏次数 减少
			err = models.UpdateUserFan(id,usered.Fan-1)//粉丝数量减一
			if err != nil{
				u.RespJSON(bean.CODE_Params_Err, err.Error())
				return
			}

		}else {
			u.RespJSON(bean.CODE_Forbidden,"错误请求!")
			return
		}
		u.RespJSONData("已取消关注!")
	}else {
		u.RespJSON(bean.CODE_Params_Err, "接口有误!")
		return
	}
}

type ArticleState struct {
	UpState	int	`json:"up_state"`
	CollectState	int	`json:"collect_state"`
}

// @Description 判断是否已经点过赞或者已经收藏 已点过赞的话返回yes 否则no
// @router /getUpState/:id [get]
func (u *UserController)GetUpState()  {
	idStr := u.Ctx.Input.Param(":id")//文章 id
	user, err := models.GetUserById(u.Uid())
	if err != nil{
		u.RespJSON(bean.CODE_Forbidden,err.Error())
		return
	}
	articleState := ArticleState{
		UpState:0,
		CollectState:0,
	}
	upArticles,collectArticles, err := models.GetArticleStateById(user.Id)
	if err != nil{
		u.RespJSON(bean.CODE_Params_Err, err.Error())
		return
	}
	if strings.Contains(upArticles,","+idStr){
		//u.RespJSONData("yes")
		articleState.UpState = 1
	}
	if strings.Contains(collectArticles,","+idStr){
		articleState.CollectState = 1
	}
	u.RespJSONData(articleState)
}


// @Description 获取用户点赞过的文章
// @router /getUpArticles [get]
func (u *UserController)GetUpArticles()  {
	limit, err := u.GetInt("limit")
	if err != nil{
		u.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	offset, err := u.GetInt("offset")
	if err != nil{
		u.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	upArticleStr, err  := models.GetUpArticlesById(u.Uid())
	if err != nil{
		u.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	upArticles := strings.Split(upArticleStr,",")
	length := len(upArticles)
	if length <=1{
		u.RespJSON(bean.CODE_Forbidden, "暂无点赞过任何文章!")
		return
	}else{
		links := make([]int, length)
		//string z转化为 int
		for i := 1; i<length; i++{//从1开始，0为空格
			s,err  := strconv.Atoi(upArticles[i])
			if err != nil{
				u.RespJSON(bean.CODE_Forbidden,err.Error())
				return
			}
			links = append(links,s)
		}
		articles,total, err := models.GetArticlesPageByIds(links,limit,offset)
		if err != nil{
			u.RespJSON(bean.CODE_Forbidden, err.Error())
			return
		}
		u.RespJSONDataWithTotal(articles, total)
	}
}

// @Description 获取用户收藏的文章
// @router /getCollectArticles [get]
func (u *UserController)GetCollectArticles()  {
	limit, err := u.GetInt("limit")
	if err != nil{
		u.RespJSON(bean.CODE_Forbidden,err.Error())
		return
	}
	offset, err := u.GetInt("offset")
	if err != nil{
		u.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	collectArticleStr, err := models.GetCollectArticles(u.Uid())
	if err != nil{
		u.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	collectArticles := strings.Split(collectArticleStr,",")
	length := len(collectArticles)
	if length <= 1{
		u.RespJSON(bean.CODE_Params_Err, "暂无任何文章的收藏!")
		return
	}else {
		links := make([]int, length)
		for i:=1; i<length; i++{
			s,err := strconv.Atoi(collectArticles[i])
			if err != nil{
				u.RespJSON(bean.CODE_Forbidden, err.Error())
				return
			}
			links = append(links, s)
		}
		articles, total, err := models.GetArticlesPageByIds(links, limit, offset)
		if err != nil{
			u.RespJSON(bean.CODE_Forbidden, err.Error())
			return
		}
		rlength := len(articles)
		if rlength <=0{
			u.RespJSON(bean.CODE_Params_Err, "暂无任何文章的收藏!")
			return
		}
		Dlinks := make([]int, len(articles))
		vlinks := make([]int, len(articles))
		for _, s := range articles{
			Dlinks = append(Dlinks, s.User.Id)
			vlinks = append(vlinks, s.ValueArticle.ValueArticleId)
		}
		sort.Ints(Dlinks)
		uLinks := utils.Duplicate(Dlinks)
		//获取作者信息
		userMap, err := models.GetUsersByIds(uLinks)
		if err != nil{
			u.RespJSON(bean.CODE_Forbidden, err.Error())
			return
		}
		//获取价值
		valueMap, err := models.GetAllValueByIds(vlinks)
		if err != nil{
			u.RespJSON(bean.CODE_Forbidden, err.Error())
			return
		}
		for i, s := range articles{
			//用户 赋值
			if u, ok := userMap[s.User.Id]; ok{
				articles[i].User =&u
			}
			//价值赋值
			if v, ok := valueMap[s.ValueArticle.ValueArticleId]; ok {
				articles[i].ValueArticle = &v
			}
		}
		u.RespJSONDataWithTotal(articles, total)
	}

}

// @Description 获取用户自己的文章
// @router /getMyArticles [get]
func (u *UserController)GetMyArticles()  {
	limit, err := u.GetInt("limit")
	if err != nil{
		u.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	offset, err := u.GetInt("offset")
	if err != nil{
		u.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}

	articles, total, err := models.GetMyArticlesById(u.Uid(),limit, offset)
	if err != nil{
		u.RespJSON(bean.CODE_Forbidden,err.Error())
		return
	}else {
		length := len(articles)
		if length <= 0{
			total = 0
			u.RespJSONDataWithTotal(articles, total)
			return
		}
		vlinks := make([]int, length)
		for _, s := range articles{
			vlinks = append(vlinks,s.ValueArticle.ValueArticleId)
		}
		//获取价值
		valueMap, err := models.GetAllValueByIds(vlinks)
		if err != nil{
			u.RespJSON(bean.CODE_Forbidden, err.Error())
			return
		}
		for i, s := range articles{
			//价值赋值
			if v, ok := valueMap[s.ValueArticle.ValueArticleId]; ok {
				articles[i].ValueArticle = &v
			}
		}
		u.RespJSONDataWithTotal(articles, total)
		return
	}

}

// @Description 修改用户的密码 首先获得旧密码
// @router /pwd [put]
func (u *UserController)PutPWD()  {
	var userJSON client.UserJSON
	if err := json.Unmarshal(u.Ctx.Input.RequestBody, &userJSON); err != nil{
		u.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	if userJSON.OldPWD == ""|| userJSON.NewPWD == "" ||userJSON.NewPWDMore == ""{
		u.RespJSON(bean.CODE_Params_Err,"缺少参数!")
		return
	}
	if userJSON.NewPWDMore != userJSON.NewPWD{
		u.RespJSON(bean.CODE_Params_Err,"两次密码不一致！")
		return
	}
	user, err := models.GetUserById(u.Uid())
	if err != nil{
		u.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	if user.Pwd != userJSON.OldPWD{
		u.RespJSON(bean.CODE_Forbidden, "密码错误!")
		return
	}else{
		user.Pwd = userJSON.NewPWD
		err := models.UpdateUserById(user)
		if err != nil{
			u.RespJSON(bean.CODE_Forbidden, err.Error())
			return
		}
		u.RespJSON(bean.CODE_Success,"修改成功!")
	}

}

// @Description 密码修改 先提交正确的密码才能进行以下步骤
// @router /putOld [put]
func (u *UserController)PutOld()  {
	oldPWD := u.GetString("oldPWD")
	user, err := models.GetUserById(u.Uid())
	if err != nil{
		u.RespJSON(bean.CODE_Forbidden,"登录出错!")
		return
	}
	if user.Pwd != oldPWD{
		u.RespJSON(bean.CODE_Forbidden,"密码错误!")
		return
	}else {
		u.RespJSON(bean.CODE_Success,"OK")
		return
	}

}

// @Description 提交密码 重新设置密码
// @router /newPWD [put]
func (u *UserController)NewPWD()  {
	var pwd client.UserNewPWD
	if err := json.Unmarshal(u.Ctx.Input.RequestBody,&pwd); err != nil{
		u.RespJSON(bean.CODE_Params_Err, err.Error())
		return
	}
	if pwd.InputNewPWD != pwd.InputAgainPWD{
		u.RespJSON(bean.CODE_Params_Err, "两次密码不一致!")
		return
	}
	if len(pwd.InputNewPWD)<6{
		u.RespJSON(bean.CODE_Params_Err, "密码长度小于6位!")
		return
	}
	user, err := models.GetUserById(u.Uid())
	if err != nil{
		u.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	user.Pwd = pwd.InputNewPWD
	err = models.UpdateUserById(user)
	if err != nil{
		u.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	u.RespJSON(bean.CODE_Success, "密码修改成功!")
	return


}


// @Description 获取自己的信息 除密码 收藏 点赞 关注 token
// @router /getSelf [get]
func (u *UserController)GetSelf()  {
	user, err := models.GetUserNotKeyInfoById(u.Uid())
	if err != nil{
		u.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	u.RespJSONData(user)
}

// @Description 修改其他信息除密码以外的
// @router /otherInfo [put]
func (u *UserController)PutInfo()  {
	//var userNotKeyJSON client.UserNotKeyJSON
	//if err := json.Unmarshal(c.Ctx.Input.RequestBody, &userNotKeyJSON); err != nil{
	//	c.RespJSON(bean.CODE_Forbidden, err.Error())
	//	return
	//}
	user, err := models.GetUserById(u.Uid())
	if err != nil{
		u.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	//var userNotKeyJSON client.UserNotKeyJSON
	if err := json.Unmarshal(u.Ctx.Input.RequestBody, &user); err != nil{
		u.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	err = models.UpdateUserInfoById(user)
	if err != nil{
		u.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	u.RespJSON(bean.CODE_Success, "修改成功!")
}

// @Description 查看关注了所有的人的信息
// @router /getAllUser [get]
func (u *UserController)GetAllUser()  {
	limit, err := u.GetInt("limit")
	if err != nil{
		u.RespJSON(bean.CODE_Params_Err, err.Error())
		return
	}
	offset, err := u.GetInt("offset")
	if err != nil{
		u.RespJSON(bean.CODE_Params_Err, err.Error())
		return
	}
	user, err := models.GetUserById(u.Uid())
	if err != nil{
		u.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	collectUserStr, err := models.GetCollectUsersById(user.Id)
	if err != nil{
		u.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	ss := strings.Split(collectUserStr, ",")
	length := len(ss)
	if length <= 1{
		u.RespJSON(bean.CODE_Params_Err,"暂无关注任何人!")
		return
	}else{
		links := make([]int,length)
		for i := 1; i<length; i++{
			sInt,_ := strconv.Atoi(ss[i])
			links = append(links,sInt)
		}
		users, total, err := models.GetAllCollectUsersByIds(links, limit, offset)
		if err != nil{
			u.RespJSON(bean.CODE_Forbidden, err.Error())
			return
		}
		u.RespJSONDataWithTotal(users, total)
	}
}

// @Description 查看是否 已经关注过了 返回状态码
// @router /getCollectUserState/:id [get]
func (u *UserController)GetCollectUserState()  {
	idStr := u.Ctx.Input.Param(":id")
	user, err := models.GetUserById(u.Uid())
	if err != nil{
		u.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	collectUsers, err := models.GetCollectUsersById(user.Id)
	if err != nil{
		u.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	if strings.Contains(collectUsers,","+idStr){
		u.RespJSONData("yes")

	}else {
		u.RespJSONData("no")
	}
}