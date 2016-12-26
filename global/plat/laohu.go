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
type Laohu struct {
	Plat
}

func (this *Laohu) Init(cfg *config.PlatConfig) {
	this.Author = cfg.Author
	this.CallbackUrl = cfg.CallbackUrl
	this.LoginUrl = cfg.LoginUrl
	this.AppID = cfg.AppID
	this.AppKey = cfg.AppKey
	this.AppSecret = cfg.AppSecret
	this.ConnTimeOut = cfg.ConnTimeOut
	this.ReadTimeOut = cfg.ReadTimeOut

	if this.Author == "" {
		log.Panic("Laohu Author not set")
	}
	if this.CallbackUrl == "" {
		log.Panic("Laohu CallbackUrl not set")
	}
	if this.LoginUrl == "" {
		log.Panic("Laohu LoginUrl not set")
	}
	if this.AppID == "" {
		log.Panic("Laohu AppID not set")
	}
	if this.AppKey == "" {
		log.Panic("Laohu AppKey not set")
	}
	if this.AppSecret == "" {
		log.Panic("Laohu AppSecret not set")
	}
	if this.ConnTimeOut == 0 {
		this.ConnTimeOut = 5
	}
	if this.ReadTimeOut == 0 {
		this.ReadTimeOut = 3
	}
}

func (this *Laohu) Clone() IClass {
	return new(Laohu)
}

func (this *Laohu) Authorize(uid, token string) {

}

func (this *Laohu) Callback(msg string) {

}
