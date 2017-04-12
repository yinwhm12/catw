package models

import "github.com/astaxie/beego/orm"

type Root2Type struct {
	Root2TypeId int `json:"root_2_type_id,omitempty" orm:"pk;column(root_2_type_id);auto"`
	Root2TypeName string `json:"root_2_type_name,omitempty" orm:"column(root_2_type_name);size(11)" json:"-"`

	EndType []*EndType `orm:"reverse(many)"`
}

func (r *Root2Type)TableName() string  {
	return "root_2_type"
}

func init()  {
	orm.RegisterModel(new(Root2Type))
}

func GetRoot2TypeInfoById(id int)(r *Root2Type, err error)  {
	o := orm.NewOrm()
	r = &Root2Type{Root2TypeId: id}
	if err = o.Read(r); err == nil{
		return r, nil
	}
	return nil, err
}

func GetAllRoot2TypeInfo()(r []Root2Type,err error)  {
	o := orm.NewOrm()
	if _, err = o.Raw("SELECT * FROM root_2_type").
		QueryRows(&r); err == nil {
		return r, nil
	}
	return nil, err
}

func GetRoot2TypeInfoByName(name string)(r *Root2Type,err error)  {
	o := orm.NewOrm()
	r = &Root2Type{Root2TypeName:name}
	if err = o.Read(r); err == nil{
		return r, nil
	}
	return nil, err
}