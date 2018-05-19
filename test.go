package main

import (
	"ftgo"

	"wechat-center/router"
)

func main() {
	ftgo.Run(":8080", router.Router)
}
