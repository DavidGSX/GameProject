package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"
)

type HttpConfig struct {
	CallbackIP    string // 充值回调监听的IP
	CallbackPort  int    // 充值回调监听的端口
	AuthorizeIP   string // 与Server通信的IP
	AuthorizePort int    // 与Server通信的端口
}

type PlatConfig struct {
	PlatType    string // 平台类型，填账号的后缀
	CallbackUrl string // 平台充值回调的标记
	LoginUrl    string // 登录认证用的Url
	AppID       string // 游戏在渠道平台注册后的ID
	AppKey      string // 游戏在渠道平台注册后的key
	AppSecret   string // 游戏在渠道平台注册后的密钥（有些平台不需要）
}

type PlatSet struct {
	Name        string   // 平台集合的名字，Server连接Global时会携带此信息 （一般为IOS，Android，硬核联盟，应用宝）
	PlatTypes   []string // 集合具体能处理的渠道
	DefaultType string   // 在集合中没有找到的渠道类型，用此渠道来处理
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
	log.Println("CallbackIP:", h.CallbackIP)
	log.Println("CallbackPort:", h.CallbackPort)
	log.Println("AuthorizeIP:", h.AuthorizeIP)
	log.Println("AuthorizePort:", h.AuthorizePort)
}

func (p *PlatConfig) Show(prefix string) {
	if p == nil {
		return
	}
	log.Println(prefix, "PlatType:", p.PlatType)
	log.Println(prefix, "CallbackUrl:", p.CallbackUrl)
	log.Println(prefix, "LoginUrl:", p.LoginUrl)
	log.Println(prefix, "AppID:", p.AppID)
	log.Println(prefix, "AppKey:", p.AppKey)
	log.Println(prefix, "AppSecret:", p.AppSecret)
}

func (p *PlatSet) Show(prefix string) {
	if p == nil {
		return
	}
	log.Println(prefix, "Name:", p.Name)
	for _, v := range p.PlatTypes {
		log.Println(prefix, "	PlatTypes:", v)
	}
	log.Println(prefix, "DefaultType:", p.DefaultType)
}

func (g *GlobalConfig) Show() {
	if g == nil {
		return
	}
	g.HttpConfig.Show()
	log.Println("")
	for _, v := range g.PlatConfigs {
		v.Show("	")
		log.Println("")
	}
	for _, v := range g.PlatSetInfo {
		v.Show("	")
		log.Println("")
	}
}
