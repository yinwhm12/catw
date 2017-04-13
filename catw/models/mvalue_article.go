package models

import "github.com/astaxie/beego/orm"

type ValueArticle struct {
	ValueArticleId int `json:"value_article_id,omitempty" orm:"pk;column(value_article_id);auto"`
	ReadCount int `json:"read_count,omitempty" orm:"column(read_count)"`
	UpVout int `json:"up_vout,omitempty" orm:"column(up_vout)"`
	CollectedCount int `json:"collected_count,omitempty" orm:"column(collected_count)"`


	Article *Article `json:"article" orm:"rel(one)"`
}

func (v *ValueArticle)TableName() string  {
	return "value_article"
}

func init()  {
	orm.RegisterModel(new(ValueArticle))
}

//阅读次数增加
func AddOneByReadById(id int)(err error)  {
	o := orm.NewOrm()
	valueArticle := ValueArticle{ValueArticleId:id}
	if o.Read(&valueArticle) == nil{
		valueArticle.ReadCount++
		if _,err = o.Update(&valueArticle); err != nil{
			return  err
		}
	}
	return nil

}

//点赞次数 增加
func AddOneByUpById(id int)(err error)  {
	o := orm.NewOrm()
	valueArticle := ValueArticle{ValueArticleId:id}
	if o.Read(&valueArticle) ==nil{
		valueArticle.UpVout++
		if _, err = o.Update(&valueArticle); err != nil{
			return err
		}
	}
	return nil
}
