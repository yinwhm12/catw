package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"github.com/kataras/go-errors"
	"yinwhm.com/yin/catw/utils"
	"strings"
)

type User struct {
	Id int `json:"id,omitempty" orm:"pk;column(id);auto"`
	Name string `json:"name,omitempty" orm:"column(name);null"`
	Pwd string `json:"pwd,omitempty" orm:"column(pwd);null"`
	CreatedTime int `json:"created_time,omitempty" orm:"column(created_time);null"`
	UpdatedTime int `json:"updated_time,omitempty" orm:"column(updated_time);null"`
	Email string `json:"email,omitempty" orm:"column(email);null"`
	AccessToken string `json:"access_token,omitempty" orm:"column(access_token);size(255);null" `
	RefreshToken string `json:"refresh_token,omitempty" orm:"column(refresh_token);size(255);null"`
	
	//文章收藏 存入文章id 格式如: ,1,2,3,4
	CollectArticles string `json:"collect_articles,omitempty" orm:"column(collect_articles);size(255);null"`
	//文章点赞 存入文章id 格式如: ,1,2,3,4
	UpArticles string `json:"up_articles,omitempty" orm:"column(up_articles);size(255);null"`

	Article []*Article `orm:"reverse(many)" json:"article,omitempty"`
}

func (u *User)TableName() string  {
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
//
func GetUserById(id int)(u *User,err error)  {
	o := orm.NewOrm()
	u = &User{Id: id}
	if err = o.Read(u); err == nil{
		return u, nil
	}
	return nil, err
}
//通过对象获得 user全部信息
func GetUserByUser(user *User)(err error)  {
	o := orm.NewOrm()
	err = o.Read(user)
	return
}
// 获取用户name email
func GetUserInfoById(id int)(u User, err error)  {
	o := orm.NewOrm()
	err = o.QueryTable(new(User)).Filter("Id",id).One(&u,"Id","Name","Email")
	return

}

func GetUserEmailNameById(id int)(u *User,err error)  {
	o := orm.NewOrm()
	u = &User{Id:id}
	err = o.Read(u)
	return

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

func GetUserInfoByEmail(email string) (u *User, err error){
	o := orm.NewOrm()
	u = &User{Email:email}
	if err = o.Read(u, "Email"); err == nil{
		return u, nil
	}
	return nil, err
}

func UpdateUserToken(u *User)(err error)  {
	o := orm.NewOrm()
	f := utils.GetNotEmptyFields(u, "AccessToken")
	num, err := o.Update(u,f...)
	if err != nil{
		return
	}
	if num == 0{
		err = errors.New("not found")
	}
	return
}

func GetUserInfoByToken(token string)(u *User, err error)  {
	o := orm.NewOrm()
	u = &User{AccessToken:token}
	if err = o.Read(u, "AccessToken"); err == nil{
		return  u, err
	}
	return nil, err

}

//检验该邮箱是否已注册过
func CheckEmailForRegister(email string)(noRow bool)  {
	o := orm.NewOrm()
	u := &User{Email:email}
	err := o.Read(u, "Email")
	if err == orm.ErrNoRows{//查询不到
		noRow = true
		return
	}else if err == orm.ErrMissPK{//找不到主键
		noRow = true
	}else {
	}
	noRow = false
	return
}

//验证密码是否有效
func Login(email, pwd string)(u *User, err error)  {
	o := orm.NewOrm()
	u = &User{Email:email,Pwd:pwd}
	if err := o.Read(u,"Email","Pwd"); err == nil{
		return u, nil
	}
	return  nil, err

}

//通过token 获得用户的id
func GetUserIdByToken(token string)(uid int, err error)  {
	o := orm.NewOrm()
	var u User
	u = User{AccessToken:token}
	if err = o.Read(&u,"AccessToken"); err == nil{
		return u.Id,nil
	}
	return 0,err
}

//通过ids 获得一批 用户的信息 仅有Id email name
func GetUsersByIds(ids []interface{})(userMap map[int]User, err error)  {
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))
	var users []User
	if _,err = qs.Filter("Id__in",ids).All(&users,"Id","Name","Email"); err != nil{
		return
	}

	userMap = map[int]User{}
	for _, u := range users {
		userMap[u.Id] = u
	}
	return userMap,nil
}

//获取用户点赞的文章upArticles 以及收藏的文章
func GetArticleStateById(id int)(upArticleStr,collectArticleStr string,err error)  {
	o := orm.NewOrm()
	var u User
	u = User{Id:id}
	if err = o.Read(&u); err == nil{
		upArticleStr = strings.TrimSpace(u.UpArticles)
		collectArticleStr = strings.TrimSpace(u.CollectArticles)
		return upArticleStr,collectArticleStr, err
	}
	return
}

func GetUpArticlesById(id int)(upArticleStr string,err error)  {
	o := orm.NewOrm()
	var u User
	u = User{Id:id}
	if err = o.Read(&u); err == nil{
		upArticleStr = strings.TrimSpace(u.UpArticles)
		return upArticleStr, err
	}
	return
}

//获取用户收藏的文章信息 collectArticles
//func GetCollectArticlesById(id int)(collectArticleStr string, err error)  {
//	o := orm.NewOrm()
//
//
//}

// 点赞 更新
func UpdateUpArticles(uid int, upArticleStr string)(err error)  {
	o := orm.NewOrm()
	_, err = o.QueryTable(new(User)).Filter("Id",uid).
		Update(orm.Params{"UpArticles":upArticleStr})
	return

}

// 收藏 更新
func UpdateCollectArticles(uid int, collectArticleStr string)(err error)  {
	o := orm.NewOrm()
	_, err = o.QueryTable(new(User)).Filter("Id",uid).
		Update(orm.Params{"CollectArticles":collectArticleStr})
	return
}

// 获取收藏 总数 string
func GetCollectArticles(id int)(collectArticles string,err error)  {
	o := orm.NewOrm()
	var u User
	u = User{Id:id}
	if err = o.Read(&u); err == nil{
		collectArticles = strings.TrimSpace(u.CollectArticles)
		return collectArticles, err
	}
	return
}