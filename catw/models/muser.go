package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"github.com/kataras/go-errors"
	"yinwhm.com/yin/catw/utils"
)

type User struct {
	Id int `json:"id,omitempty" orm:"column(id);atuo"`
	Name string `json:"name,omitempty" orm:"column(name);null"`
	Pwd string `json:"pwd,omitempty" orm:"column(pwd);null"`
	CreatedTime int `json:"created_time,omitempty" orm:"column(created_time);null"`
	UpdatedTime int `json:"updated_time,omitempty" orm:"column(updated_time);null"`
}

func (u *User)TablName() string  {
	return "user"
}

func init()  {
	orm.RegisterModel(new(User))
}

func AddUser(u *User)(id int64,err error)  {
	u.CreatedTime = int(time.Now().Unix())
	o := orm.NewOrm()
	id,err = o.Insert(u)
	return

}

func GetUserInfoByName(name string) (u *User,err error)  {
	o := orm.NewOrm()
	u = &User{Name:name}
	if err = o.Read(u,"Name"); err == nil{
		return  u, nil
	}
	return nil, err

}

func UpdateUserById(u *User) (err error)  {
	o := orm.NewOrm()
	u.UpdatedTime = int(time.Now().Unix())
	fields := utils.GetNotEmptyFields(u,"Pwd")
	num, err := o.Update(u, fields...)
	if err != nil{
		return
	}
	if num == 0{
		err = errors.New("not found")
	}
	return

}