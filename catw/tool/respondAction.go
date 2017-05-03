package tool

import (
	"sync"
	"sort"
	"yinwhm.com/yin/catw/models"
	"yinwhm.com/yin/catw/utils"
)

//var wg sync.WaitGroup
//线程 从数据库获取数据
func GetRespondTwos(rOne *[]models.RespondOne,flag chan int)(err error) {
	//var i *int
	var wg sync.WaitGroup
	for i, s := range  *rOne{
		wg.Add(1)
		//go func() {
		//	fmt.Println("range------rone",s,"------i--",i)
		//	fmt.Println("----------ssi---",s.RespondOneId)
		//	rTwo, err :=models.GetAllRespondTwoByROne(s)
		//	if err != nil{
		//		return
		//	}
		//	fmt.Println("------erji---",rTwo)
		//	fmt.Println("------erjiddd---",len(rTwo))
		//	flag:=GetRespondTwoUser(rTwo)//二级评论者信息
		//	if  flag == "some"{
		//		rOne[i].RespondTwos = rTwo
		//	}
		//	wg.Done()
		//}()
		go GetRespondOneWg(&wg,s,&(*rOne)[i])
	}
	wg.Wait()
	flag <- len(*rOne)
	return
}

//获取二级评论信息
func GetRespondOneWg(wg *sync.WaitGroup,one models.RespondOne,rOne *models.RespondOne)  {
	rTwo, err :=models.GetAllRespondTwoByROne(one)
	if err != nil{
		return
	}
	//fmt.Println("------erji---",rTwo)
	//fmt.Println("------erjiddd---",len(rTwo))
	flag:=GetRespondTwoUser(rTwo)//二级评论者信息
	//fmt.Println("-----------flag-------",flag)
	if  flag == "some"{
		rOne.RespondTwos = rTwo
		//fmt.Println("rone____retwo----",rOne.RespondTwos)
	}
	wg.Done()
}

//获取二级评论者的信息
func GetRespondTwoUser(rTwo []*models.RespondTwo)(flag string)  {
	if len(rTwo) ==0{
		return "none"
	}
	Dlinks := make([]int,len(rTwo))
	for _, s := range rTwo{
		Dlinks = append(Dlinks,s.User.Id)
		//fmt.Println("==========id-------",s.User.Id)
		//fmt.Println("==========idssss-------",s)
	}
	sort.Ints(Dlinks)
	links := utils.Duplicate(Dlinks)
	//fmt.Println("=========links-----",len(links))
	userMap, err := models.GetUsersByIds(links)
	//fmt.Println("-----go -----userMap----",userMap)
	if err != nil{
		//c.RespJSON(bean.CODE_Forbidden, err.Error())
		return "none"
	}
	//进行相应的 user赋值
	for i, s := range rTwo{
		if u, ok := userMap[s.User.Id]; ok{
			rTwo[i].User =  &u
		}
	}
	return "some"
}