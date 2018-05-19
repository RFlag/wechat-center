package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"wechat-center/conf"
	"wechat-center/entity"

	redis "gopkg.in/redis.v4"
)

func init() {
	client = redis.NewClient(&redis.Options{Addr: "192.168.0.81:6379", Password: "123456"})
	pong := client.Ping()
	log.Println(pong)
}

var (
	client *redis.Client
)

func GetAccessToken(toUserName string) (string, error) {
	appId := conf.PublicNum[toUserName].AppId
	appSecret := conf.PublicNum[toUserName].AppSecret
	duration, _ := time.ParseDuration("-100m")
	validTime := time.Now().Add(duration).Unix()
	v, err := client.Get(toUserName).Result()
	if err == redis.Nil {
		accessToken := new(entity.AccessToken)
		req, err := http.NewRequest("GET", "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid="+appId+"&secret="+appSecret, nil)
		if err != nil {
			log.Println(err)
			return "", err
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println(err)
			return "", err
		}
		byteToken, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			return "", err
		}
		err = json.Unmarshal(byteToken, &accessToken)
		if err != nil {
			return "", err
		}

		var validTimeToken = map[string]entity.ValidTimeToken{
			toUserName: {
				AccessToken: accessToken.AccessToken,
				NextTime:    time.Now().Unix(),
			},
		}
		byteValidToken, err := json.Marshal(validTimeToken[toUserName])
		if err != nil {
			return "", err
		}
		validToken := string(byteValidToken)
		err = client.Set(toUserName, validToken, time.Hour*2).Err()
		if err != nil {
			return "", err
		}

		return accessToken.AccessToken, nil
	} else if err != nil {
		return "", err
	}

	var validTimeToken entity.ValidTimeToken
	err = json.Unmarshal([]byte(v), &validTimeToken)
	if err != nil {
		return "", err
	}
	if validTime > validTimeToken.NextTime {
		accessToken := new(entity.AccessToken)
		req, err := http.NewRequest("GET", "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid="+appId+"&secret="+appSecret, nil)
		if err != nil {
			log.Println(err)
			return "", err
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println(err)
			return "", err
		}
		byteToken, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			return "", err
		}
		err = json.Unmarshal(byteToken, &accessToken)
		if err != nil {
			return "", err
		}

		var validTimeToken = map[string]entity.ValidTimeToken{
			toUserName: {
				AccessToken: accessToken.AccessToken,
				NextTime:    time.Now().Unix(),
			},
		}
		byteValidToken, err := json.Marshal(validTimeToken[toUserName])
		if err != nil {
			return "", err
		}
		validToken := string(byteValidToken)
		err = client.Set(toUserName, validToken, time.Hour*2).Err()
		if err != nil {
			return "", err
		}

		return accessToken.AccessToken, nil
	} else {
		return validTimeToken.AccessToken, nil
	}

}
