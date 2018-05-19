package router

import (
	"wechat-center/controller"

	"github.com/gin-gonic/gin"
)

func Router(g *gin.Engine) {
	private(g.Group("/private"))

	g.GET("/wechat/:wechat", controller.CheckServer)
	g.POST("/wechat/:wechat", controller.Wechat)
}
