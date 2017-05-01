package models

import (
	"github.com/astaxie/beego/orm"
)

type ValueArticle struct {
	ValueArticleId int `json:"value_article_id,omitempty" orm:"pk;column(value_article_id);auto"`
	ReadCount int `json:"read_count,omitempty" orm:"column(read_count)"` //阅读次数
	UpVout int `json:"up_vout,omitempty" orm:"column(up_vout)"` //赞
	CollectedCount int `json:"collected_count,omitempty" orm:"column(collected_count)"` //收藏次数


	//Article *Article `json:"article,omitempty" orm:"reverse(one)"`
	//Article *Article `json:"article,omitempty" orm:"rel(one)"`
}

func (v *ValueArticle)TableName() string  {
	return "value_article"
}

func init()  {
	orm.RegisterModel(new(ValueArticle))
}

//添加一条该文章价值记录
func AddValueArticle()(value_article ValueArticle,err error)  {
	o := orm.NewOrm()
	value_article = ValueArticle{ReadCount:1}
	_,err = o.Insert(&value_article)
	return
}

//删除某条记录
func DeletValueArticleById(valueArticle *ValueArticle)(err error)  {
	o := orm.NewOrm()
	 _, err = o.Delete(&valueArticle)
	return
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
//获取文章 价值信息
func GetOneValueById(v *ValueArticle)(err error)  {
	o := orm.NewOrm()
	//v = &ValueArticle{ValueArticleId:id}
	err = o.Read(v)
	return
}

//获取文章 所有的价值 参数为数组
func GetAllValueByIds(ids []int)(valueMap map[int]ValueArticle,err error)  {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ValueArticle))
	var valuse []ValueArticle
	if _, err = qs.Filter("ValueArticleId__in",ids).All(&valuse); err != nil{
		return
	}

	valueMap = map[int]ValueArticle{}
	for _, v :=range valuse{
		valueMap[v.ValueArticleId] = v
	}
	return valueMap, nil
}

//点赞次数 增加
func AddOneByUpById(id int)(err error)  {
	o := orm.NewOrm()
	//if o.Read(&valueArticle) ==nil{
	//	valueArticle.UpVout++
	//	fmt.Println("-------uuuu",valueArticle.UpVout)
	//	if _, err = o.Update(&valueArticle,"UpVout"); err == nil{
	//		fmt.Println("------fffiiiii")
	//		return nil
	//	}
	//}
	_, err = o.Raw("UPDATE  value_article v set v.up_vout = " +
		"v.up_vout + 1 WHERE v.value_article_id = ?",id).Exec()

	return err
}

//点赞次数 减一
func DeletOneByUpId(id int)(err error)  {
	o := orm.NewOrm()
	valueArticle := ValueArticle{ValueArticleId:id}
	if o.Read(&valueArticle) == nil{
		valueArticle.UpVout--
		if _, err = o.Update(&valueArticle); err != nil{
			return err
		}
	}
	return nil
}

//收藏数 加一
func AddOneByCollectId(id int)(err error)  {
	o := orm.NewOrm()
	valueArticle := ValueArticle{ValueArticleId:id}
	if o.Read(&valueArticle) ==nil{
		valueArticle.CollectedCount++
		if _, err = o.Update(&valueArticle); err != nil{
			return err
		}
	}
	return nil
}

//收藏数 减一
func DeletOneByCollectId(id int)(err error)  {
	o := orm.NewOrm()
	valueArticle := ValueArticle{ValueArticleId:id}
	if o.Read(&valueArticle) == nil{
		valueArticle.CollectedCount--
		if _, err = o.Update(&valueArticle); err != nil{
			return err
		}
	}
	return nil

}

