package client

import (
	"net/url"
	"strings"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"errors"
)

type UcResponse struct {
	StatusCode	int
	Body	[]byte
}

func RequestUcApi(tourl,method string,args map[string]string)(ucres UcResponse,err error)  {
	v:=url.Values{}
	for key,value := range args{
		if key != "atoken" {
			v.Set(key,value)
		}
	}
	senddata := v.Encode()
	body := strings.NewReader(senddata)
	client := &http.Client{}
	req, err := http.NewRequest(method,tourl,body); if err != nil{
		return
	}
	token, ok := args["atoken"]; if ok{
		req.Header.Set("x-token",token)
	}
	req.Header.Set("Content-Type","application/x-www-form-urlencoded;param=value")
	response ,err := client.Do(req); if err != nil {
		return
	}
	data, err := ioutil.ReadAll(response.Body); if err != nil{
		return
	}
	defer response.Body.Close()

	ucres.StatusCode = response.StatusCode
	ucres.Body = data
	return

}

func GetResErrInfo(response UcResponse)(err error)  {
	if response.StatusCode != http.StatusOK{
		var msg string
		err = json.Unmarshal(response.Body, &msg); if err != nil{
			return
		}
		return errors.New(msg)
	}
	return nil
	
}