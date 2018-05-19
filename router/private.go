package router

import (
	"wechat-center/controller"

	"github.com/gin-gonic/gin"
)

func private(g *gin.RouterGroup) {
	g.POST("/token", controller.AccessToken)
	g.POST("/notice", controller.Notice)
}
