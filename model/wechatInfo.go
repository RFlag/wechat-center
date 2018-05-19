package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"wechat-center/entity"
)

func GetWechatInfo(accessToken string) (*entity.WechatInfo, error) {
	wechatInfo := new(entity.WechatInfo)
	req, err := http.NewRequest("GET", "https://api.weixin.qq.com/cgi-bin/user/get?access_token="+accessToken, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	byteWechatList, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = json.Unmarshal(byteWechatList, &wechatInfo)
	if err != nil {
		return nil, err
	}
	return wechatInfo, nil
}
func GetUserInfo(accessToken, openId string) (*entity.UserInfo, error) {
	userInfo := new(entity.UserInfo)
	req, err := http.NewRequest("GET", "https://api.weixin.qq.com/cgi-bin/user/info?access_token="+accessToken+"&openid="+openId, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	byteUserInfo, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = json.Unmarshal(byteUserInfo, &userInfo)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}
func Message(nickName, inviteNickName string, total int) string {
	if inviteNickName == "" {
		return "很高兴遇见你(" + nickName + ")!\n现已成为第" + strconv.Itoa(total) + "位粉丝"
	}
	return "很高兴遇见你(" + nickName + ")!\n恭喜您通过(" + inviteNickName + ")的分享,\n现已成为第" + strconv.Itoa(total) + "位粉丝"
}
