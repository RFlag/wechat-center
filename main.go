package main

import (
	"ftgo"

	"wechat-center/router"
)

func main() {
	ftgo.Run(":80", router.Router)
}
