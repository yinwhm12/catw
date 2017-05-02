package tool

import (
	"sync"
	"sort"
	"yinwhm.com/yin/catw/models"
	"yinwhm.com/yin/catw/utils"
	"fmt"
)

//线程 从数据库获取数据
func GetRespondTwos(rOne []models.RespondOne,flag chan int)(err error) {
	//var i *int
	var wg sync.WaitGroup
	for i, s := range  rOne{
		fmt.Println("range------rone",s)
		wg.Add(1)
		go func() {
			rTwo, err :=models.GetAllRespondTwoByROne(s)
			if err != nil{
				return
			}
			fmt.Println("------erji---",rTwo)
			flag:=GetRespondTwoUser(rTwo)//二级评论者信息
			if  flag == "some"{
				rOne[i].RespondTwos = rTwo
			}
			wg.Done()
		}()
	}
	wg.Wait()
	flag <- len(rOne)
	return
}

//获取二级评论者的信息
func GetRespondTwoUser(rTwo []*models.RespondTwo)(flag string)  {
	if len(rTwo) ==0{
		return "none"
	}
	Dlinks := make([]int,len(rTwo))
	for _, s := range rTwo{
		Dlinks = append(Dlinks,s.User.Id)
	}
	sort.Ints(Dlinks)
	links := utils.Duplicate(Dlinks)
	fmt.Println("=========links-----",len(links))
	userMap, err := models.GetUsersByIds(links)
	fmt.Println("-----go -----userMap----",userMap)
	if err != nil{
		//c.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	//进行相应的 user赋值
	for i, s := range rTwo{
		if u, ok := userMap[s.User.Id]; ok{
			rTwo[i].User =  &u
		}
	}
	return "some"
}