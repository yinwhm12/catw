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


	User *User `json:"user,omitempty" orm:"rel(fk)"`
	EndType *EndType `json:"end_type,omitempty" orm:"rel(fk)"`
	ValueArticle *ValueArticle `json:"value_article,omitempty" orm:"reverse(one)"`
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

func GetArticlesByUser(id int)(articles *[]Article,err error)  {
	if _,err =orm.NewOrm().QueryTable("article").
		Filter("User",id).RelatedSel().All(&articles); err != nil{
		return nil,err
	}
	return articles,nil
}

func GetArticlesByEndType(id int)(articles *[]Article,err error)  {
	if _,err = orm.NewOrm().QueryTable("article").
		Filter("EndType",id).RelatedSel().All(&articles); err != nil{
		return nil,err
	}
	return articles, nil
}

//通过rootId 类型获得
func GetAticlesByRoot1Id(root1_id int)(articles []Article, err error)  {
	o := orm.NewOrm()
	_, err = o.Raw("SELECT * FROM acticle a INER JOIN end_type e on " +
		"a.end_type_id = e.end_type_id WHERE e.root1_type_id = (" +
		"SELECT root_1_type_id From root_1_type Where root_1_type_id = ?",root1_id).QueryRows(&articles)
	if err != nil{
		return nil, err
	}
	return articles, nil
}

//通过root2Id类型获得
func GetAticlesByRoot2Id(root2_id int)(articles []Article,err error)  {
	o := orm.NewOrm()
	_, err = o.Raw("SELECT * FROM acticle a INER JOIN end_type e on " +
		"a.end_type_id = e.end_type_id WHERE e.root2_type_id = (" +
		"SELECT root_2_type_id From root_2_type Where root_2_type_id = ?",root2_id).QueryRows(&articles)
	if err != nil{
		return
	}
	return articles, nil
}

//通过 LevelTypeId
func GetAticlesByLevelTypeId(id int)(articles []Article,err error)  {
	o := orm.NewOrm()
	_, err = o.Raw("SELECT * FROM acticle a INER JOIN end_type e on " +
		"a.end_type_id = e.end_type_id WHERE e.level_type_id = (" +
		"SELECT level_type_id From level_type Where level_type_id = ?",id).QueryRows(&articles)
	if err != nil{
		return
	}
	return articles, nil
}

//主页 获得的课间操类型数据 根据id 获得相应类型最新的 文章 或者 帖子 或者 课间操
func GetPalyThemeIndex(id int) (articles []Article, err error) {
	o := orm.NewOrm()
	if _, err = o.Raw("SELECT title,created_time FROM article a INNER JOIN" +
		" end_type e ON a.end_type_id = e.end_type_id INNER JOIN root_1_type r " +
		"ON e.root1_type_id = r.root_1_type_id WHERE r.root_1_type_id = ? " +
		"ORDER BY a.created_time DESC LIMIT 0,9",id).QueryRows(&articles); err != nil{
		return
	}
	return articles, nil
}

