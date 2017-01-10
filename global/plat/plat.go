package plat

import (
	"gameproject/global/config"
	"net/url"
)

type Plat struct {
	Author      string // 平台类型，填账号的后缀
	CallbackUrl string // 平台充值回调的标记
	LoginUrl    string // 登录认证用的Url
	AppID       string // 游戏在渠道平台注册后的ID
	AppKey      string // 游戏在渠道平台注册后的key
	AppSecret   string // 游戏在渠道平台注册后的密钥（有些平台不需要）
	ConnTimeOut uint32 // 连接超时时间，单位秒，默认为5秒
	ReadTimeOut uint32 // 读取超时时间，单位秒，默认为3秒
}

type IClass interface {
	Init(*config.PlatConfig)
	Clone() IClass
	Authorize(string, string) string
	Callback(url.Values) string
}

type IPlatMgr interface {
	AddClass(string, IClass)
}

// 所有已实现的渠道，都需要在此处注册
func InitClass(platMgr IPlatMgr) {
	platMgr.AddClass("AllApp", new(AllApp))
	platMgr.AddClass("OneSdk", new(OneSdk))
	platMgr.AddClass("Laohu", new(Laohu))
}
