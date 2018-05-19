package controller

import (
	"ftgo/ftsync"
	"log"
	"wechat-center/model"

	"github.com/gin-gonic/gin"
)

func AccessToken(c *gin.Context) {
	var Param struct {
		ToUserName string `binding:"required" json:"toUserName" from:"toUserName"` //开发者微信号
	}
	err := c.Bind(&Param)
	if err != nil {
		return
	}
	m, err := ftsync.Lock("wechat-center.AccessToken")
	if err != nil {
		return
	}
	defer m.Unlock()

	//微信号获取AppId AppSecret
	accessToken, err := model.GetAccessToken(Param.ToUserName)
	if err != nil {
		log.Println(err)
		return
	}
	c.String(200, accessToken)
}
