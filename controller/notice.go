package controller

import(
	"github.com/gin-gonic/gin"
	"log"

	"wechat/model"
)

func CheckServer(c *gin.Context){
	var param struct{
		Signature string
		Timestamp int64
		Nonce string
		Echostr string
	}
	        // signatureGen := model.CheckServer(param.Timestamp, param.Nonce)
	        //
	        // signatureIn := strings.Join(r.Form["signature"], "")
	        // if signatureGen != signatureIn {
	        //         return false
	        // }
	        // echostr := strings.Join(r.Form["echostr"], "")
	        // fmt.Fprintf(w, echostr)
	       c.String(200, param.Echostr)
}


func Notice(c *gin.Context) {
	var ParamQRCode struct {
		ToUserName   string `xml:"ToUserName"`   //发送方帐号(一个OpenID)
		FromUserName string `xml:"FromUserName"` //开发者微信号
		CreateTime   int64 `xml:"CreateTime"`   //消息创建时间(整型)
		MsgType      string `xml:"MsgType"`      //消息类型,event
		Event        string `xml:"Event"`        //事件类型,subscribe(订阅).unsubiscribe(取消订阅)
		EventKey     string `xml:"EventKey"`     //事件KEY值，qrscene_为前缀，后面为二维码的参数值
		Ticket       string `xml:"Ticket"`       //二维码的ticket，可用来换取二维码图片
	}

    err:=c.Bind(&ParamQRCode)
	if err != nil {
		log.Println("ParamError")
		return
	}
   log.Println("###ToUserName :",ParamQRCode.ToUserName)

   if ParamQRCode.MsgType == "event" { //事件推送
	   if ParamQRCode.Event == "subscrie" { //关注或扫码关注
		   if ParamQRCode.Ticket == "" { //关注
             replyMessage:= model.Notice(ParamQRCode.ToUserName,ParamQRCode.FromUserName,"nihaoa")
			log.Println("!!!!replyMessage:",replyMessage)
			 c.XML(200, replyMessage)
		   } else { //扫码关注
			   return
		   }
	   } else if ParamQRCode.Event == "LOCATION" { //上报地理位置
		  return
	   } else if ParamQRCode.Event == "CLICK" { //自定义菜单事件
          return
	   }
   }
   return
}
