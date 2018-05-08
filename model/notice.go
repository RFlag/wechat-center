package model

import (
	"time"

	"wechat/entity"
)

func Notice(toUserName, fromUserName, content string) *entity.ReplyMessage {
	var replyMessage *entity.ReplyMessage
	replyMessage.ToUserName = toUserName
	replyMessage.FromUserName = fromUserName
	replyMessage.CreateTime = time.Now().Unix()
	replyMessage.MsgType = "text"
	replyMessage.Content = content
	return replyMessage
}
