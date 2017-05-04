package tool

import (
	"yinwhm.com/yin/catw/models"
	"sync"
)

func DealAddManyMessage(id int,message models.Message,wg *sync.WaitGroup)  {
	user, err := models.GetUserById(id)
	if err != nil{//其中 获取某个用户的信息 出问题时 忽略
		//c.RespJSON(bean.CODE_Forbidden,err.Error())
		//return
		wg.Done()
		return
	}
	message.ToUser = user
	//插入数据
	err = models.AddOneMessage(&message)
	wg.Done()

}
