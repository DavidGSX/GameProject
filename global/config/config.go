package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"
)

type BaseConfig struct {
	JMXIP            string // JMX监听的IP
	JMXPort          uint32 // JMX监听的端口
	CallbackIP       string // 充值回调监听的IP
	CallbackPort     uint32 // 充值回调监听的端口
	AuthorizeIP      string // 与Server通信的IP
	AuthorizePort    uint32 // 与Server通信的端口
	RPCIP            string // RPC监听的IP
	RPCPort          uint32 // RPC监听的端口
	DBIP             string // 数据库的IP
	DBPort           uint32 // 数据库的端口
	MinPoolSize      uint32 // 连接池最少的连接数
	MaxPoolSize      uint32 // 连接池最大的连接数
	AcquireIncrement uint32 // 连接池每次增加的数量
}

type PlatConfig struct {
	PlatID      uint32 // 平台的唯一ID
	ClassName   string // 处理此平台的类名称，需要与程序实现保持一致
	Author      string // 平台类型，填账号的后缀
	CallbackUrl string // 平台充值回调的标记
	LoginUrl    string // 登录认证用的Url
	AppID       string // 游戏在渠道平台注册后的ID
	AppKey      string // 游戏在渠道平台注册后的key
	AppSecret   string // 游戏在渠道平台注册后的密钥（有些平台不需要）
	ConnTimeOut uint32 // 连接超时时间，单位秒，默认为5秒
	ReadTimeOut uint32 // 读取超时时间，单位秒，默认为3秒
}

type PlatSet struct {
	SetID   uint32   // 平台集合的名字，Server连接Global时会携带此信息 （一般为IOS，Android，硬核联盟，应用宝）
	PlatIDs []uint32 // 集合具体能处理的渠道
}

type GlobalConfig struct {
	BaseConfig  BaseConfig   // 基础配置
	PlatConfigs []PlatConfig // 所有渠道的配置
	PlatSetInfo []PlatSet    // 所有集合的配置
}

var config *GlobalConfig
var configLock sync.Mutex

func LoadConfig(filename string) bool {
	defer func() {
		if err := recover(); err != nil {
			log.Println("LoadConfig >>>>>>", err)
		}
	}()
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Panic("Read Global Config Error:", err, " filename:", filename)
	}
	cfg := new(GlobalConfig)
	err = json.Unmarshal(content, cfg)
	if err != nil {
		log.Panic("Unmarshal Config Error:", err)
	}
	config = cfg
	return true
}

func GetConfig() *GlobalConfig {
	configLock.Lock()
	defer configLock.Unlock()
	if config == nil {
		LoadConfig("./config/global.conf")
	}
	return config
}

func ReloadConfig() bool {
	configLock.Lock()
	defer configLock.Unlock()
	return LoadConfig("./config/global.conf")
}

func (this *BaseConfig) Show() {
	if this == nil {
		return
	}
	log.Println("JMXIP:", this.JMXIP)
	log.Println("JMXPort:", this.JMXPort)
	log.Println("CallbackIP:", this.CallbackIP)
	log.Println("CallbackPort:", this.CallbackPort)
	log.Println("AuthorizeIP:", this.AuthorizeIP)
	log.Println("AuthorizePort:", this.AuthorizePort)
	log.Println("RPCIP:", this.RPCIP)
	log.Println("RPCPort:", this.RPCPort)
	log.Println("DBIP:", this.DBIP)
	log.Println("DBPort:", this.DBPort)
	log.Println("MinPoolSize:", this.MinPoolSize)
	log.Println("MaxPoolSize:", this.MaxPoolSize)
	log.Println("AcquireIncrement:", this.AcquireIncrement)
	log.Println()
}

func (this *PlatConfig) Show(prefix string) {
	if this == nil {
		return
	}
	log.Println(prefix, "PlatID:", this.PlatID)
	log.Println(prefix, "ClassName:", this.ClassName)
	log.Println(prefix, "Author:", this.Author)
	log.Println(prefix, "CallbackUrl:", this.CallbackUrl)
	log.Println(prefix, "LoginUrl:", this.LoginUrl)
	log.Println(prefix, "AppID:", this.AppID)
	log.Println(prefix, "AppKey:", this.AppKey)
	log.Println(prefix, "AppSecret:", this.AppSecret)
	log.Println(prefix, "ConnTimeOut:", this.ConnTimeOut)
	log.Println(prefix, "ReadTimeOut:", this.ReadTimeOut)
	log.Println(prefix)
}

func (this *PlatSet) Show(prefix string) {
	if this == nil {
		return
	}
	log.Println(prefix, "SetID:", this.SetID)
	for _, v := range this.PlatIDs {
		log.Println(prefix, "	PlatIDs:", v)
	}
	log.Println(prefix)
}

func (this *GlobalConfig) Show() {
	if this == nil {
		return
	}
	this.BaseConfig.Show()
	for _, v := range this.PlatConfigs {
		v.Show("	")
	}
	for _, v := range this.PlatSetInfo {
		v.Show("	")
	}
}
