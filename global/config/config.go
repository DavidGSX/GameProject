package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"
)

type HttpConfig struct {
	JMXIP         string // JMX监听的IP
	JMXPort       int    // JMX监听的端口
	CallbackIP    string // 充值回调监听的IP
	CallbackPort  int    // 充值回调监听的端口
	AuthorizeIP   string // 与Server通信的IP
	AuthorizePort int    // 与Server通信的端口
}

type PlatConfig struct {
	PlatID      int    // 平台的唯一ID
	ClassName   string // 处理此平台的类名称，需要与程序实现保持一致
	Author      string // 平台类型，填账号的后缀
	CallbackUrl string // 平台充值回调的标记
	LoginUrl    string // 登录认证用的Url
	AppID       string // 游戏在渠道平台注册后的ID
	AppKey      string // 游戏在渠道平台注册后的key
	AppSecret   string // 游戏在渠道平台注册后的密钥（有些平台不需要）
	ConnTimeOut int    // 连接超时时间，单位秒，默认为5秒
	ReadTimeOut int    // 读取超时时间，单位秒，默认为3秒
}

type PlatSet struct {
	SetID       int   // 平台集合的名字，Server连接Global时会携带此信息 （一般为IOS，Android，硬核联盟，应用宝）
	PlatIDs     []int // 集合具体能处理的渠道
	DefaultPlat int   // 在集合中没有找到的渠道类型，用此渠道来处理
}

type GlobalConfig struct {
	HttpConfig  HttpConfig   // 网络通信的配置
	PlatConfigs []PlatConfig // 所有渠道的配置
	PlatSetInfo []PlatSet    // 所有集合的配置
}

var config *GlobalConfig
var l sync.Mutex

func LoadConfig(filename string) error {
	defer func() {
		if err := recover(); err != nil {
			log.Println("LoadConfig >>>>>>", err)
		}
	}()
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Panic("Read Global Config Error:", err, " filename:", filename)
		return err
	}
	cfg := new(GlobalConfig)
	err = json.Unmarshal(content, cfg)
	if err != nil {
		log.Panic("Unmarshal Global Config Error:", err)
		return err
	}
	config = cfg
	return nil
}

func GetConfig() *GlobalConfig {
	l.Lock()
	defer l.Unlock()
	if config == nil {
		LoadConfig("./config/global.conf")
	}
	return config
}

func ReloadConfig() error {
	l.Lock()
	defer l.Unlock()
	return LoadConfig("./config/global.conf")
}

func (h *HttpConfig) Show() {
	if h == nil {
		return
	}
	log.Println("JMXIP:", h.JMXIP)
	log.Println("JMXPort:", h.JMXPort)
	log.Println("CallbackIP:", h.CallbackIP)
	log.Println("CallbackPort:", h.CallbackPort)
	log.Println("AuthorizeIP:", h.AuthorizeIP)
	log.Println("AuthorizePort:", h.AuthorizePort)
	log.Println()
}

func (p *PlatConfig) Show(prefix string) {
	if p == nil {
		return
	}
	log.Println(prefix, "PlatID:", p.PlatID)
	log.Println(prefix, "ClassName:", p.ClassName)
	log.Println(prefix, "Author:", p.Author)
	log.Println(prefix, "CallbackUrl:", p.CallbackUrl)
	log.Println(prefix, "LoginUrl:", p.LoginUrl)
	log.Println(prefix, "AppID:", p.AppID)
	log.Println(prefix, "AppKey:", p.AppKey)
	log.Println(prefix, "AppSecret:", p.AppSecret)
	log.Println(prefix, "ConnTimeOut:", p.ConnTimeOut)
	log.Println(prefix, "ReadTimeOut:", p.ReadTimeOut)
	log.Println(prefix)
}

func (p *PlatSet) Show(prefix string) {
	if p == nil {
		return
	}
	log.Println(prefix, "SetID:", p.SetID)
	for _, v := range p.PlatIDs {
		log.Println(prefix, "	PlatIDs:", v)
	}
	log.Println(prefix, "DefaultPlat:", p.DefaultPlat)
	log.Println(prefix)
}

func (g *GlobalConfig) Show() {
	if g == nil {
		return
	}
	g.HttpConfig.Show()
	for _, v := range g.PlatConfigs {
		v.Show("	")
	}
	for _, v := range g.PlatSetInfo {
		v.Show("	")
	}
}
