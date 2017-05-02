package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type RespondOne struct {
	RespondOneId	int 	`json:"respond_one_id,omitempty" orm:"pk;column(respond_one_id);auto"`
	CreatedTime	int	`json:"created_time,omitempty" orm:"column(created_time)"`
	TextContent	string  `json:"text_content,omitempty" orm:"column(text_content);type(text)"`
	ImgContent	string	`json:"img_content,omitempty" orm:"column(img_content)"`

	User	*User	`json:"user,omitempty" orm:"rel(fk)"`
	Article	*Article	`json:"article,omitempty" orm:"rel(fk)"`

	//RespondTwos *[]RespondTwo	`json:"respond_twos,omitempty" orm:"-"`
	RespondTwos []*RespondTwo	`json:"respond_twos,omitempty" orm:"reverse(many)"`

}

func (r *RespondOne)TableName() string  {
	return "respond_one"
}

func init()  {
	orm.RegisterModel(new(RespondOne))
}


//添加 一条评论
func AddRespondOne(rOne *RespondOne)(err error)  {
	rOne.CreatedTime = int(time.Now().Unix())
	o := orm.NewOrm()
	_, err = o.Insert(rOne)
	return
}

//取评论
func GetAllRespondOneByArticleId(tid int)(rOnes []RespondOne,err error)  {
	o := orm.NewOrm()
	qs := o.QueryTable(new(RespondOne))
	_, err = qs.Filter("Article",tid).OrderBy("CreatedTime").All(&rOnes)
	return
}

//获取 相应的二级评论
func GetOne2ManyRespondTwo()  {
	
}

//根据id 获取具体的某个一级评论
func GetOneRespondById(id int)(rOne *RespondOne, err error)  {
	o := orm.NewOrm()
	rOne = &RespondOne{RespondOneId:id}
	err = o.Read(rOne)
	return

}