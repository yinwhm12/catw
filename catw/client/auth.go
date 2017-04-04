package client

import (
	"net/http"
	"encoding/json"
)

type UserAuth struct {
	Email	string
	UserName	string
	AccessToken	string
	RefreshToken	string
	Password	string
}

//登录
func Login(email,pwd string)(userAuth *UserAuth,err error)  {
	args:=map[string]string{
		"Eamil":email,
		"Password":pwd,
	}
	response,err := RequestUcApi(ucenterUrl + apiVersion + routerCheckLogin, "post", args); if err != nil{
		return
	}
	if response.StatusCode == http.StatusOK {
		userAuth = new(UserAuth)
		err = json.Unmarshal(response.Body,userAuth); if err != nil {
			return
		}
		return
	}else {
		return nil, GetResErrInfo(response)
	}
}

//检查AccessToken
func CheckAccessToken(token, username string) (res bool, err error) {
	args := map[string]string{
		"atoken":token,
	}
	response, err := RequestUcApi(ucenterUrl + apiVersion + routerCheckAccessToken + "/" + username, "GET", args); if err != nil {
		return
	}
	if response.StatusCode == http.StatusOK {
		res = true
		return
	} else {
		return false, GetResErrInfo(response)
	}
}

//通过RefreshToken 重置 AccessToken
func ResetAccessToken(rtoken, username string) (atoken string, err error) {
	args := map[string]string{
		"UserName":username,
		"RefreshToken":rtoken,
	}
	response, err := RequestUcApi(ucenterUrl + apiVersion + routerResetAccessToken, "POST", args); if err != nil {
		return
	}
	if response.StatusCode == http.StatusOK {
		err = json.Unmarshal(response.Body,&atoken); if err != nil {
			return
		}
		return
	}else {
		return "",GetResErrInfo(response)
	}
}


func OffLine(name string) (err error) {
	args := map[string]string{
	}
	response, err := RequestUcApi(ucenterUrl + apiVersion + routerOffLine + "/" + name, "GET", args); if err != nil {
		return
	}
	if response.StatusCode == http.StatusOK {
		return
	}else {
		return GetResErrInfo(response)
	}
}

//注册
func RegisterUser(email, pwd string)(resemail string, err error)  {
	args := map[string]string{
		"Email":email,
		"Password":pwd,
	}
	response, err := RequestUcApi(ucenterUrl + apiVersion + routerUser, "POST", args); if err != nil{
		return
	}
	var u UserAuth
	if response.StatusCode == http.StatusOK{
		err = json.Unmarshal(response.Body, &u); if err != nil{
			return
		}
		resemail = u.Email
		return
	}else {
		return "", GetResErrInfo(response)
	}

}
