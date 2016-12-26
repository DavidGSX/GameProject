package plat

import (
	"gameproject/global/config"
)

type Plat struct {
	Author      string // 平台类型，填账号的后缀
	CallbackUrl string // 平台充值回调的标记
	LoginUrl    string // 登录认证用的Url
	AppID       string // 游戏在渠道平台注册后的ID
	AppKey      string // 游戏在渠道平台注册后的key
	AppSecret   string // 游戏在渠道平台注册后的密钥（有些平台不需要）
	ConnTimeOut int    // 连接超时时间，单位秒，默认为5秒
	ReadTimeOut int    // 读取超时时间，单位秒，默认为3秒
}

type IClass interface {
	Init(*config.PlatConfig)
	Clone() IClass
	Authorize(string, string)
	Callback(string)
}
