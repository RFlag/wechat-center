package main

import (
	"ftgo"

	"wechat/router"
)

func main() {

	ftgo.Run(":80", router.Router)
}
