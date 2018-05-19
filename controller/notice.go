package controller

import (
	"wechat-center/model/notice"

	"github.com/gin-gonic/gin"
)

func Notice(c *gin.Context) {
	var Param struct {
		ToUserName   string `json:"toUserName"`
		FromUserName string `json:"fromUserName"`
	}
	err := c.Bind(&Param)
	if err != nil {
		return
	}
	c.Data(200, "application/xml; charset=utf-8", []byte(notice.GetReplyMessage(Param.ToUserName, Param.FromUserName, "")))

}
