package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"fmt"
)

type RespondTwo struct {
	RespondTwoId	int	`json:"respond_two_id,omitempty" orm:"pk;column(respond_two_id);auto"`
	CreatedTime	int	`json:"created_time,omitempty" orm:"column(created_time)"`
	TextContent	string	`json:"text_content,omitempty" orm:"column(text_content)"`
	ImgContent	string	`json:"img_content,omitempty" orm:"column(img_content)"`

	User	*User	`json:"user,omitempty" orm:"rel(fk)"`
	RespondOne	*RespondOne	`json:"respond_one,omitempty" orm:"rel(fk)"`
	//Article	*Article	`json:"article,omitempty" orm:"rel(fk)"`
}

func (r *RespondTwo)TableName()string  {
	return "respond_two"
}

func init()  {
	orm.RegisterModel(new(RespondTwo))
}

//添加一条评论
func AddRespondTwo(rTwo *RespondTwo)(err error)  {
	rTwo.CreatedTime = (int)(time.Now().Unix())
	o := orm.NewOrm()
	_, err = o.Insert(rTwo)
	return
}

//取二级评论
func GetAllRespondTwoByROne(one RespondOne)(rTwo []*RespondTwo, err error)  {
	o := orm.NewOrm()
	qs := o.QueryTable(new(RespondTwo))
	_, err = qs.Filter("respond_one_id",one.RespondOneId).OrderBy("CreatedTime").All(&rTwo)
	//one.RespondTwos = &rTwo
	return
	
}
func GetAllTestRespondTwoByROne(one *RespondOne)(rTwo []*RespondTwo, err error)  {
	o := orm.NewOrm()
	qs := o.QueryTable(new(RespondTwo))
	fmt.Println("respondId----",one.RespondOneId)
	_, err = qs.Filter("respond_one_id",one.RespondOneId).OrderBy("CreatedTime").All(&rTwo)
	//one.RespondTwos = &rTwo
	return

}

// 传入 管道 获取二级
func GetAllRespondTwoByChan(one RespondOne,c chan []*RespondTwo)( err error)  {
	var rTwo []*RespondTwo
	o := orm.NewOrm()
	qs := o.QueryTable(new(RespondTwo))
	_, err = qs.Filter("RespondOne",one.RespondOneId).OrderBy("CreatedTime").RelatedSel().All(&rTwo)
	if err != nil{
		return
	}
	c <- rTwo //往管道写入值
	//close(c)
	//one.RespondTwos = &rTwo
	return

}

//利用sync 处理二级评论获得
func GetAllRespondTwoBySync()  {

}

//通过一级Id 获取所有参与的评论的用户信息
func GetAllUserByRespondOne(id int)(respondTwos []*RespondTwo, err error)  {
	o := orm.NewOrm()
	qs := o.QueryTable(new(RespondTwo))
	_,err = qs.Filter("RespondOne",id).RelatedSel().All(&respondTwos)
	return
}

//通过一级Id 获取所有参与的评论的用户信息id
func GetAllUserIdsByRespondOne(id int)(ids []int, err error)  {
	var respondTwos []*RespondTwo
	o := orm.NewOrm()
	qs := o.QueryTable(new(RespondTwo))
	_, err = qs.Filter("RespondOne",id).All(&respondTwos)
	ids = make([]int,len(respondTwos))
	if err ==nil{
		for i, s := range respondTwos{
			ids[i] = s.User.Id
		}
		return
	}
	return

}