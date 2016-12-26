package plat

import (
	"gameproject/global/config"
	"log"
)

/*
	Author    string // 平台类型，填账号的后缀
	CallbackUrl string // 平台充值回调的标记
	LoginUrl    string // 登录认证用的Url
	AppID       string // 游戏在渠道平台注册后的ID
	AppKey      string // 游戏在渠道平台注册后的key
	AppSecret   string // 游戏在渠道平台注册后的密钥（有些平台不需要）
	ConnTimeOut int    // 连接超时时间，单位秒，默认为5秒
	ReadTimeOut int    // 读取超时时间，单位秒，默认为3秒
*/
type OneSdk struct {
	Plat
}

func (this *OneSdk) Init(cfg *config.PlatConfig) {
	this.Author = cfg.Author
	this.CallbackUrl = cfg.CallbackUrl
	this.LoginUrl = cfg.LoginUrl
	this.AppID = cfg.AppID
	this.AppKey = cfg.AppKey
	this.AppSecret = cfg.AppSecret
	this.ConnTimeOut = cfg.ConnTimeOut
	this.ReadTimeOut = cfg.ReadTimeOut

	if this.Author == "" {
		log.Panic("OneSdk Author not set")
	}
	if this.CallbackUrl == "" {
		log.Panic("OneSdk CallbackUrl not set")
	}
	if this.LoginUrl == "" {
		log.Panic("OneSdk LoginUrl not set")
	}
	if this.AppID == "" {
		log.Panic("OneSdk AppID not set")
	}
	if this.AppKey == "" {
		log.Panic("OneSdk AppKey not set")
	}
	if this.AppSecret == "" {
		log.Panic("OneSdk AppSecret not set")
	}
	if this.ConnTimeOut == 0 {
		this.ConnTimeOut = 5
	}
	if this.ReadTimeOut == 0 {
		this.ReadTimeOut = 3
	}
}

func (this *OneSdk) Clone() IClass {
	return new(OneSdk)
}

func (this *OneSdk) Authorize(uid, token string) {

}

func (this *OneSdk) Callback(msg string) {

}
