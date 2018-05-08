package router

import (
	"wechat/controller"

	"github.com/gin-gonic/gin"
)

func Router(g *gin.Engine) {
	// g.ServeHTTP(w http.ResponseWriter, req *http.Request)
	g.GET("/", controller.CheckServer)
	g.POST("/", controller.Notice)
}
