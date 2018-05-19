package notice

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"wechat-center/entity"
	"wechat-center/model"
)

func Notice(toUserName, fromUserName, content string) *entity.ReplyMessage {
	replyMessage := new(entity.ReplyMessage)
	replyMessage.ToUserName = fromUserName
	replyMessage.FromUserName = toUserName
	replyMessage.CreateTime = time.Now().Unix()
	replyMessage.MsgType = "text"
	replyMessage.Content = content
	return replyMessage
}
func GetReplyMessage(toUserName, fromUserName, inviteOpenId string) string {
	accessToken, err := model.GetAccessToken(toUserName) //获取token
	if err != nil {
		return ""
	}
	userInfo, err := model.GetUserInfo(accessToken, fromUserName) //获取用户信息
	if err != nil {
		return ""
	}
	wechatInfo, err := model.GetWechatInfo(accessToken) //获取微信信息
	if err != nil {
		return ""
	}
	var content string
	if inviteOpenId != "" {
		inviteUserInfo, err := model.GetUserInfo(accessToken, fromUserName) //获取用户信息
		if err != nil {
			return ""
		}
		SendInviteMessage(accessToken, inviteOpenId, userInfo.NickName)
		content = model.Message(userInfo.NickName, inviteUserInfo.NickName, wechatInfo.Total)
	} else {
		content = model.Message(userInfo.NickName, "", wechatInfo.Total)
	}
	replyMessage := Notice(toUserName, fromUserName, content)

	reply := "<xml>" +
		"<ToUserName>" + replyMessage.ToUserName + "</ToUserName>" +
		"<FromUserName>" + replyMessage.FromUserName + "</FromUserName>" +
		"<CreateTime>" + strconv.FormatInt(replyMessage.CreateTime, 10) + "</CreateTime>" +
		"<MsgType>" + replyMessage.MsgType + "</MsgType>" +
		"<Content>" + replyMessage.Content + "</Content>" +
		"</xml>"

	return reply
}
func SendInviteMessage(accessToken, inviteOpenId, nickName string) error {
	timeNow := time.Now().Format("2006-01-02 15:04:05")

	param := " {" +
		"\"touser\":\"" + inviteOpenId + "\"," +
		"\"template_id\":\"x_NyJP1erJc1Gm3mRdliAnvjyWO3SKLP1NvH9L3UYbY\"," +
		"\"data\":{" +
		"\"first\": {" +
		"\"value\":\"恭喜你邀请成功！\\n\"," +
		" \"color\":\"#173177\"" +
		"}," +
		" \"keyword1\":{" +
		"\"value\":\"" + nickName + "\\n\"," +
		" \"color\":\"#173177\"" +
		" }," +
		"\"keyword2\": {" +
		" \"value\":\"" + timeNow + "\\n\"," +
		"   \"color\":\"#173177\"" +
		"}" +
		"}" +
		"}"
	var jsonStr = []byte(param)
	body := bytes.NewBuffer([]byte(jsonStr))
	res, err := http.Post("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token="+accessToken, "application/json;charset=utf-8", body)
	if err != nil {
		log.Fatal(err)
		return err
	}
	_, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
