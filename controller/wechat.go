package controller

import (
	"crypto/sha1"
	"fmt"
	"ftgo/ftapi"
	"io"
	"log"
	"sort"
	"strings"

	"wechat-center/conf"
	"wechat-center/model/notice"

	"github.com/gin-gonic/gin"
)

var callbackUrl = map[string][]string{
	"无参数关注": []string{
		"/noParam/subscribe",
	},
	"参数关注": []string{
		"/param/subscribe",
	},
}

func CheckServer(c *gin.Context) {
	wechat := c.Param("wechat")
	token := conf.PublicNum[wechat].Token
	var Param struct {
		Signature string `form:"signature"`
		Timestamp string `form:"timestamp"`
		Nonce     string `form:"nonce"`
		Echostr   string `form:"echostr"`
	}

	err := c.BindQuery(&Param)
	if err != nil {
		log.Println("ParamError")
		return
	}
	log.Println(Param.Echostr)
	sl := []string{token, Param.Timestamp, Param.Nonce}
	sort.Strings(sl)
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))
	signatureGen := fmt.Sprintf("%x", s.Sum(nil))
	if signatureGen != Param.Signature {
		return
	}
	c.String(200, Param.Echostr)
}

func Wechat(c *gin.Context) {
	var Param struct {
		ToUserName   string `xml:"ToUserName"`   //开发者微信号
		FromUserName string `xml:"FromUserName"` //发送方帐号(一个OpenID)
		CreateTime   int64  `xml:"CreateTime"`   //消息创建时间(整型)
		MsgType      string `xml:"MsgType"`      //消息类型,event
		Event        string `xml:"Event"`        //事件类型,subscribe(订阅).unsubiscribe(取消订阅)
		EventKey     string `xml:"EventKey"`     //事件KEY值，qrscene_为前缀，后面为二维码的参数值
		Ticket       string `xml:"Ticket"`       //二维码的ticket，可用来换取二维码图片
	}

	err := c.Bind(&Param)
	if err != nil {
		log.Println("ParamError")
		return
	}
	if Param.MsgType != "event" { //事件推送
		return
	}
	if Param.Event == "subscribe" { //关注或扫码关注
		var replyMessage string
		if Param.EventKey == "" { //关注
			//model层函数调用
			replyMessage = notice.GetReplyMessage(Param.ToUserName, Param.FromUserName, "")
			log.Println("@@@@", replyMessage)
			if replyMessage == "" {
				// 第三方插件
				for _, url := range callbackUrl["无参数关注"] {
					// TODO: 回调 url

					var result struct {
						ftapi.ResultList
						Data struct {
							Reply string `json:"reply"`
						} `json:"data"`
					}
					err = ftapi.Post(url, Param, &result)
					if err != nil {
						return
					}
					c.Data(200, "application/xml; charset=utf-8", []byte(result.Data.Reply))
					return
				}
			}
			c.Data(200, "application/xml; charset=utf-8", []byte(replyMessage))

		} else {
			//model层函数调用
			scene := strings.Split(Param.EventKey, "qrscene_")

			inviteOpenId := conf.QrCodeParam[scene[1]].OpenId
			replyMessage = notice.GetReplyMessage(Param.ToUserName, Param.FromUserName, inviteOpenId)
			if replyMessage == "" {
				for _, url := range callbackUrl["参数关注"] {
					// TODO: 回调 url
					var result struct {
						ftapi.ResultList
						Data struct {
							Reply string `json:"reply"`
						} `json:"data"`
					}
					err = ftapi.Post(url, Param, &result)
					if err != nil {
						return
					}

					c.Data(200, "application/xml; charset=utf-8", []byte(result.Data.Reply))
					return
				}
			}
			c.Data(200, "application/xml; charset=utf-8", []byte(replyMessage))

		}
	}
}
