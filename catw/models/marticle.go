package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Article struct {
	Tid int `json:"tid,omitempty" orm:"pk;column(tid);auto"`
	Title string `json:"title,omitempty" orm:"column(title);size(50)"`
	CreatedTime int `json:"created_time,omitempty" orm:"column(created_time);"`
	TextContent string `json:"text_content,omitempty" orm:"column(text_content);null"`
	ImgContent string `json:"img_content,omitempty" orm:"column(img_content)"`


	User *User `json:"user" orm:"rel(fk)"`
	EndType *EndType `json:"end_type" orm:"rel(fk)"`
}

func (a *Article)TableName() string {
	return "article"
}

func init()  {
	orm.RegisterModel(new(Article))
}

func AddArticle(a *Article)(id int64,err error)  {
	a.CreatedTime = int(time.Now().Unix())
	o := orm.NewOrm()
	id, err = o.Insert(a)
	return 
}

func GetArticleById(id int) (a *Article,err error)  {
	o := orm.NewOrm()
	a = &Article{Tid:id}
	if err = o.Read(a); err == nil{
		return a, nil
	}
	return nil, err
}

func GetArticleByUser(id int)(articles []*Article,err error)  {
	if _,err =orm.NewOrm().QueryTable("article").
		Filter("User",id).RelatedSel().All(&articles); err != nil{
		return nil,err
	}
	return articles,nil
}

func GetArticleByEndType(id int)(articles []*Article,err error)  {
	if _,err = orm.NewOrm().QueryTable("article").
		Filter("EndType",id).RelatedSel().All(&articles); err != nil{
		return nil,err
	}
	return articles, nil
}