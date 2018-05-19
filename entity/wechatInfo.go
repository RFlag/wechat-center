package entity

type WechatInfo struct {
	Total int `json:"total"`
	Count int `json:"count"`
	Data  InfoData
}
type InfoData struct {
	OpenId     []string `json:"openid"`
	NextOpenId string   `json:"next_openid"`
}
type AccessToken struct {
	AccessToken string `json:"access_token"`
	// Timestamps  uint64 `json:"timestamps"`
	ExpiresIn int `json:"expires_in"`
}
type UserInfo struct {
	Subscribe      int    `json:"subscribe"`
	OpenId         string `json:"openid"`
	NickName       string `json:"nickname"`
	Sex            int    `json:"sex"`
	City           string `json:"city"`
	Country        string `json:"country"`
	Province       string `josn:"province"`
	Language       string `json:"language"`
	HeadImgUrl     string `json:"headimgurl"`
	SubscribeTime  int64  `json:"subscribe_time"`
	Unionid        string `json:"unionid"`
	Remark         string `json:"remark"`
	GroupId        int    `json:"groupid"`
	TagIdList      []int  `json:"tagid_list"`
	SubscribeScene string `json:"subscribe_scene"`
	QrScene        int    `json:"qr_scene"`
	QrSceneStr     string `json:"qr_scene_str"`
}
type PubilcNum struct {
	AppId       string `json:"app_id"`
	AppSecret   string `json:"app_secret"`
	Token       string `json:"token"` //验证token
	AccessToken string `json:"access_token"`
	NextTime    int64  `json:"next_time"`
}
type ValidTimeToken struct {
	AccessToken string `json:"access_token"`
	NextTime    int64  `json:"next_time"`
}
type QrCodeParam struct {
	OpenId string `json:"open_id"`
}
