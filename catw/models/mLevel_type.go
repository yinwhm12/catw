package models

import "github.com/astaxie/beego/orm"

type LevelType struct {
	LevelTypeId int `json:"level_type_id,omitempty" orm:"pk;column(level_type_id);auto"`
	LevelTypeName string `json:"level_type_name,omitempty" orm:"column(level_type_name);size(11)" json:"-"`

	EndType []*EndType `orm:"reverse(many)"`
}

func (level *LevelType)TableName() string {
	return "level_type"
}

func init()  {
	orm.RegisterModel(new(LevelType))
}

func GetLevelTypeInfoById(id int)(l *LevelType, err error)  {
	o := orm.NewOrm()
	l = &LevelType{LevelTypeId: id}
	if err = o.Read(l); err == nil{
		return l, nil
	}
	return nil, err
}

func GetAllLevelTypeInfo()(l []LevelType,err error)  {
	o := orm.NewOrm()
	if _, err = o.Raw("SELECT * FROM level_type").
		QueryRows(&l);err == nil{
		return l, err
	}
	return nil, err
}

func GetLevelTypeInfoByName(name string)(l *LevelType,err error)  {
	o := orm.NewOrm()
	l = &LevelType{LevelTypeName:name}
	if err = o.Read(l); err == nil{
		return l, nil
	}
	return nil,err

}