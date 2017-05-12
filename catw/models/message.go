package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Message struct {
	MessageId	int	`json:"message_id,omitempty" orm:"pk;column(message_id);auto"`
	State	int	`json:"state,omitempty" orm:"column(state);default(0)"`
	Type	int	`json:"type,omitempty" orm:"column(type);default(0)"`
	CreatedTime	int	`json:"created_time,omitempty" orm:"column(created_time);"`
	Content	string	`json:"content,omitempty" orm:"column(content);size(255);null"`

	FromUser	*User	`json:"from_user,omitempty" orm:"rel(fk)"`
	ToUser	*User	`json:"to_user,omitempty" orm:"rel(fk)"`
	Article	*Article	`json:"article,omitempty" orm:"rel(fk)"`


	//state 默认0 表示未读信息 1表示已读信息
	//from_user 表示发送者	to_user表示接收者
	//type 消息类型 0表示回复文章类型的消息 其他类型 暂无定义


	//一级回复 只有楼主有权收到消息
	//二级回复 有当前一级回复的人以及楼主收到信息
}

func (m *Message)TableName()string  {
	return "message"
}

func init()  {
	orm.RegisterModel(new(Message))

}

//增加一条消息 仅是一级回复时 调用
func AddOneMessage(message *Message)(err error)  {
	message.CreatedTime = int(time.Now().Unix())
	o := orm.NewOrm()
	_, err = o.Insert(message)
	return
}

// 增加多条信息 当回复时二级信息时
func AddMoreMessage(messages []*Message)(err error)  {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Message))
	i, _ := qs.PrepareInsert()
	for _, message := range messages{
		_, err = i.Insert(message)
		if err != nil{
			return err
		}
	}
	i.Close()
	return nil
}

// 获取接受者信息数量
func GetMessageCountByToUser(user *User) (err error) {
	o := orm.NewOrm()
	err = o.QueryTable(new(Message)).Filter("ToUser",user.Id).RelatedSel().One(user)
	return
}

// 获取接受者未读信息 数量 首页
func GetMessageUnreadCountByToUser(user *User)(total int64,err error)  {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Message))
	qs = qs.Filter("ToUser",user.Id).Filter("State",0)
	total, err = qs.Count()
	return

}
// 获取接受者未读信息 分页 按时间先后
func GetMessageUnreadPageByToUser(user *User, limit, offset int)(messages []*Message, total int64, err error)  {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Message))
	qs = qs.Filter("ToUser",user.Id).Filter("State",0).OrderBy("CreatedTime")
	total, err = qs.Count()
	if err != nil{
		return
	}
	_, err = qs.Limit(limit).Offset(offset).All(&messages)
	return
	
}

// 获取已读的信息 数量 以及具体的 已读分页 每页10条
func GetMessageReadPageByToUser(user *User,limit, offset int)(messages []*Message, total int64, err error)  {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Message))
	qs = qs.Filter("ToUser",user.Id).Filter("State",1).OrderBy("-CreatedTime")
	total, err = qs.Count()
	if err != nil{
		return 
	}
	_, err = qs.Limit(limit).Offset(offset).All(&messages)
	return 

}

// 获取全部信息 分页 10条
func GetAllMessagesByToUser(user *User, limit, offset int)(messages []*Message, total int64, err error)  {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Message))
	qs = qs.Filter("ToUser", user.Id).OrderBy("-State","CreatedTime")
	total, err = qs.Count()
	if err != nil{
		return 
	}
	_, err = qs.Limit(limit).Offset(offset).All(&messages)
	return 

}

// 获取具体的某条信息 
func GetOneMessageById(id int)(message Message, err error)  {
	o := orm.NewOrm()
	message = Message{MessageId:id}
	if err = o.Read(message); err == nil{
		return message, nil
	}

	return
}

// 删除某条具体的信息
func DeletOneMessageById(id int)(err error)  {
	o := orm.NewOrm()
	_, err = o.Delete(&Message{MessageId:id})
	return

}

// 改变某条信息的状态
func ChangeMessageStateById(id,state int)(err error)  {
	o := orm.NewOrm()
	_, err = o.QueryTable(new(Message)).Filter("MessageId",id).
		Update(orm.Params{"State":state})
	return
}