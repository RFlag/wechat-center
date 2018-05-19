package entity

//扫描带参数二维码事件
type jsonParamQRCode struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Event        string
	EventKey     string
	Ticket       string
}

type ParamQRCode struct {
	ToUserName   string `xml:"ToUserName"`   //发送方帐号(一个OpenID)
	FromUserName string `xml:"FromUserName"` //开发者微信号
	CreateTime   int64  `xml:"CreateTime"`   //消息创建时间(整型)
	MsgType      string `xml:"MsgType"`      //消息类型,event
	Event        string `xml:"Event"`        //事件类型,subscribe(订阅).unsubscribe(取消订阅)
	EventKey     string `xml:"EventKey"`     //事件KEY值，qrscene_为前缀，后面为二维码的参数值
	Ticket       string `xml:"Ticket"`       //二维码的ticket，可用来换取二维码图片
}

type ReplyMessage struct {
	ToUserName   string `xml:"ToUserName"`   //发送方帐号(一个OpenID)
	FromUserName string `xml:"FromUserName"` //开发者微信号
	CreateTime   int64  `xml:"CreateTime"`   //消息创建时间(整型)
	MsgType      string `xml:"MsgType"`      //text 文本
	Content      string `xml:"Content"`      //回复的消息内容（换行：在content中能够换行，微信客户端就支持换行显示）
}
