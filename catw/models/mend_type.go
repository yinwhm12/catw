package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

type EndType struct {
	EndTypeId int `json:"end_type_id,omitempty" orm:"pk;column(end_type_id);auto"`
	LevelName string `json:"level_name,omitempty" orm:"column(level_name);size(11)"`
	//Root1TypeId int `json:"root_1_type_id,omitempty" orm:"column(root_1_type_id);"`
	//Root2TypeId int `json:"root_2_type_id,omitempty" orm:"column(root_2_type_id)"`
	//LevelTypeId int `json:"level_type_id,omitempty" orm:"column(level_type_id)"`


	Root1Type *Root1Type `json:"root_1_type,omitempty" orm:"rel(fk)"`
	Root2Type *Root2Type `json:"root_2_type,omitempty" orm:"rel(fk)"`
	LevelType *LevelType `json:"level_type,omitempty" orm:"rel(fk)"`

	//TagsRoot1 []*Tag `orm:"rel(m2m)"`
	Article []*Article `orm:"reverse(many)" json:"article,omitempty"`

}


func (end *EndType)TableName()string  {
	return "end_type"
}

func init()  {
	orm.RegisterModel(new(EndType))

}

func AddEndType(endType *EndType) (err error) {
	//o := orm.NewOrm()
	fmt.Println("-=----n")
	//if _, err = o.InsertMulti(2, endType.Root1Type); err != nil{
	//	return err
	//}
	fmt.Println("--mmmm")
	//if _, err = o.InsertMulti(2, endType.Root2Type); err != nil{
	//	return err
	//}
	//if _, err = o.InsertMulti(2, endType.Root2Type); err != nil{
	//	return err
	//}


	return nil
}



func GetEndTypeInfoById(id int)(endType *EndType,err error)  {
	o := orm.NewOrm()
	endType = &EndType{EndTypeId:id}
	if err = o.Read(endType); err == nil{
		return endType, nil
	}
	return nil,err
}

func GetEndTypeInfoByAllFK(root1,root2,level int)(endType EndType,err error)  {
	o := orm.NewOrm()
	qs := o.QueryTable(new(EndType))
	err = qs.Filter("root1_type_id",root1).Filter("root2_type_id",root2).
		Filter("level_type_id",level).Limit(1).One(&endType)
	return

}




