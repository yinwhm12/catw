package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"fmt"
)

type Article struct {
	Tid int `json:"tid,omitempty" orm:"pk;column(tid);auto"`
	Title string `json:"title,omitempty" orm:"column(title);size(50)"`
	CreatedTime int `json:"created_time,omitempty" orm:"column(created_time);"`
	TextContent string `json:"text_content,omitempty" orm:"column(text_content);null"`
	ImgContent string `json:"img_content,omitempty" orm:"column(img_content)"`



	UserEmail string  `json:"user_email,omitempty" orm:"-"`

	User *User `json:"user,omitempty" orm:"rel(fk)"`
	EndType *EndType `json:"end_type,omitempty" orm:"rel(fk)"`
	ValueArticle *ValueArticle `json:"value_article,omitempty" orm:"rel(fk)"`
}

func (a *Article)TableName() string {
	return "article"
}

func init()  {
	orm.RegisterModel(new(Article))
}

//func AddArticleByOne2One(a *Article)(err error)  {
//	a.CreatedTime = int(time.Now().Unix())
//	valueArticle := ValueArticle{ReadCount:1}
//
//
//	o := orm.NewOrm()
//	valueArticle.Article = a
//	_, err = o.Insert(a)
//
//	_, err = o.Insert(&valueArticle)
//	return
//}
func AddArticle(a *Article)(err error)  {
	a.CreatedTime = int(time.Now().Unix())
	o := orm.NewOrm()
	_, err = o.Insert(a)
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

func GetArticlesByUser(id int)(articles []Article,err error)  {
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

//通过root1Id 类型获得
func GetAticlesByRoot1Id(root1_id int)(articles []Article, err error)  {
	o := orm.NewOrm()
	_, err = o.Raw("SELECT * FROM acticle a INER JOIN end_type e on " +
		"a.end_type_id = e.end_type_id WHERE e.root1_type_id = (" +
		"SELECT root_1_type_id From root_1_type Where root_1_type_id = ?",root1_id).QueryRows(&articles)
	if err != nil{
		return
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
	_, err = o.Raw("SELECT * FROM acticle a INNER JOIN end_type e on " +
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
	if _, err = o.Raw("SELECT tid,title,created_time FROM article a INNER JOIN" +
		" end_type e ON a.end_type_id = e.end_type_id INNER JOIN root_1_type r " +
		"ON e.root1_type_id = r.root_1_type_id WHERE r.root_1_type_id = ? " +
		"ORDER BY a.created_time DESC LIMIT 0,9",id).QueryRows(&articles); err != nil{
		return
	}
	return articles, nil
}

//通过root1Id 获得 不同的数据量    构造查询
func GetThemesByRoot1Id(flag string, id int)(articles []Article, err error)  {
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").From("article").InnerJoin("end_type").
		On("article.end_type_id = end_type.end_type_id").InnerJoin("root_1_type").
		On("end_type.root1_type_id = root_1_type.root_1_type_id").
		Where("root_1_type.root_1_type_id = ?").OrderBy("article.created_time").Desc()
	if flag == "index"{
		qb.Limit(5).Offset(0)
	}
	sql := qb.String()
	o := orm.NewOrm()
	if _, err = o.Raw(sql,id).QueryRows(&articles); err != nil{
		return
	}
	return articles, nil

}

//获取 value_article   有错
func GetValueArticleByArticle(articles *[]Article)(err error)  {
	if len(*articles) == 0{
		return
	}
	 links :=make([]int,len(*articles))
	for i, s := range *articles{
		links[i] = s.Tid
	}
	o := orm.NewOrm()
	_, err = o.QueryTable(new(ValueArticle)).Filter("Article__Tid",links).All(&articles)
	return
}

//通过endType数组获得全部的数据 获得分页 每页数量为10
func GetIndexAllByPage(ids []int,limit, offset int)(articles []Article,total int64, err error){
	o := orm.NewOrm()
	qs := o.QueryTable(new(Article))
	qs = qs.Filter("EndType__in",ids)
	total, err = qs.Count()
	if err != nil{
		return
	}
	_, err = qs.Limit(limit).Offset(offset).All(&articles)

	return

}

//分页 获取点赞过的文章
func GetArticlesPageByIds(ids []int,limit,offset int)(articles []*Article, total int64, err error)  {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Article))
	qs = qs.Filter("Tid__in",ids)
	total, err = qs.Count()
	if err != nil{
		return
	}
	_, err = qs.Limit(limit).Offset(offset).All(&articles)
	return

}


//test 
func GetPages(root1 int)(articles []Article, err error)  {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Article))
	cout, err := qs.Count()
	fmt.Println("all------",cout)
	if err != nil{
		return
	}

	_, err = qs.Filter("EndType__end_type_id",root1).All(&articles)
	return

}






