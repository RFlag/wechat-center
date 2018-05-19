package controller

import (
	"encoding/json"
	"ftgo"
	"io/ioutil"
	"log"
	"net/http"
	"wechat-center/entity"

	"github.com/gin-gonic/gin"
)

func GetWechatInfo(c *gin.Context) {
	var Param struct {
		AccessToken string `json:"access_token"`
	}
	err := c.Bind(&Param)
	if err != nil {
		return
	}
	wechatInfo := new(entity.WechatInfo)
	req, err := http.NewRequest("GET", "https://api.weixin.qq.com/cgi-bin/user/get?access_token="+Param.AccessToken, nil)
	if err != nil {
		log.Println(err)
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	byteWechatList, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(byteWechatList, &wechatInfo)
	if err != nil {
		return
	}
	c.JSON(200, ftgo.ResultData(wechatInfo))
}
func GetUserInfo(c *gin.Context) {
	var Param struct {
		AccessToken string `json:"access_token"`
		OpenId      string `json:"openId"`
	}
	userInfo := new(entity.UserInfo)
	req, err := http.NewRequest("GET", "https://api.weixin.qq.com/cgi-bin/user/info?access_token="+Param.AccessToken+"&openid="+Param.OpenId, nil)
	if err != nil {
		log.Println(err)
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	byteUserInfo, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(byteUserInfo, &userInfo)
	if err != nil {
		return
	}
	c.JSON(200, ftgo.ResultData(userInfo))
}
