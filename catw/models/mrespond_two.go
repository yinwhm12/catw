package models

import "github.com/astaxie/beego/orm"

type RespondTwo struct {
	RespondTwoId	int	`json:"respond_two_id,omitempty" orm:"pk;column(respond_two_id);auto"`
	CreatedTime	int	`json:"created_time,omitempty" orm:"column(created_time)"`
	TextContent	string	`json:"text_content,omitempty" orm:"column(text_content)"`
	ImgContent	string	`json:"img_content,omitempty" orm:"column(img_content)"`

	User	*User	`json:"user,omitempty" orm:"rel(fk)"`
	RespondOne	*RespondOne	`json:"respond_one,omitempty" orm:"rel(fk)"`
	Article	*Article	`json:"article,omitempty" orm:"rel(fk)"`
}

func (r *RespondTwo)TableName()string  {
	return "respond_two"
}

func init()  {
	orm.RegisterModel(new(RespondTwo))
}

//添加一条评论
func AddRespondTwo(rTwo *RespondTwo)(err error)  {
	o := orm.NewOrm()
	_, err = o.Insert(&rTwo)
	return
}

//取二级评论
func GetAllRespondTwoByROne(one RespondOne)(rTwo []RespondTwo, err error)  {
	o := orm.NewOrm()
	qs := o.QueryTable(new(RespondTwo))
	_, err = qs.Filter("RespondOne",one.RespondOneId).All(&rTwo)
	return
	
}
