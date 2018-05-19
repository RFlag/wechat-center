package controller

import (
	"wechat-center/conf"

	"github.com/gin-gonic/gin"
)

func Scene(c *gin.Context) {
	var Param struct {
		SceneId string `json:"scene_id"`
	}
	err := c.Bind(&Param)
	if err != nil {
		return
	}
	openId := conf.QrCodeParam[Param.SceneId]
	c.JSON(200, openId)

}
