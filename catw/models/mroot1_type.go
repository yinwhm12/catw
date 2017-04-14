package models

import (
	"github.com/astaxie/beego/orm"
)

type Root1Type struct {
	Root1TypeId int `json:"root_1_type_id,omitempty" orm:"pk;column(root_1_type_id);auto"`
	Root1TypeName string `json:"root_1_type_name,omitempty" orm:"column(root_1_type_name);size(11)"`

	EndType []*EndType `orm:"reverse(many)" json:"end_type,omitempty"`
}

func (r *Root1Type)TableName()string  {
	return "root_1_type"
}

func init()  {
	orm.RegisterModel(new(Root1Type))
}

func GetRoot1TypeInfoById(id int) (r *Root1Type, err error)  {
	o := orm.NewOrm()
	r = &Root1Type{Root1TypeId: id}
	if err = o.Read(r); err == nil{
		return r,nil
	}
	return nil, err
}

func GetAllRoot1TypeInfo()(r []Root1Type,err error)  {

	o := orm.NewOrm()
	if _,err = o.Raw("SELECT * FROM root_1_type").QueryRows(&r); err == nil{
		return r, nil
	}
	return nil, err
}

func GetRoot1TypeInfoByName(name string)(r *Root1Type,err error)  {
	o := orm.NewOrm()
	r = &Root1Type{Root1TypeName:name}
	if err = o.Read(r); err == nil{
		return r, nil
	}
	return nil,err
}