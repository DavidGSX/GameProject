package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"
)

type BaseConfig struct {
	ZoneId   uint32 // 服务器的唯一ID
	Platform uint32 // 连接Global的平台类型
}

type DBConfig struct {
	DBIP             string // 数据库的IP
	DBPort           uint32 // 数据库的端口
	MinPoolSize      uint32 // 连接池最少的连接数
	MaxPoolSize      uint32 // 连接池最大的连接数
	AcquireIncrement uint32 // 连接池每次增加的数量
}

type JMXConfig struct {
	JMXIP   string // JMX监听的IP
	JMXPort uint32 // JMX监听的端口
}

type LinkConfig struct {
	LinkIP   string // 监听玩家连接的IP
	LinkPort uint32 // 监听玩家连接的端口
}

type GlobalConfig struct {
	GlobalIP   string // 连接Global的IP
	GlobalPort uint32 // 连接Global的端口
}

type RPCConfig struct {
	RPCIP   string // 连接RPC的IP
	RPCPort uint32 // 连接RPC的端口
}

type ServerConfig struct {
	BaseConfig   BaseConfig   // 基础信息配置
	DBConfig     DBConfig     // 数据库配置
	JMXConfig    JMXConfig    // JMX的配置
	LinkConfig   LinkConfig   // Link的配置
	GlobalConfig GlobalConfig // Global的配置
	RPCConfig    RPCConfig    // RPC的配置
}

var config *ServerConfig
var configLock sync.Mutex

func LoadConfig(filename string) bool {
	defer func() {
		if err := recover(); err != nil {
			log.Println("LoadConfig >>>>>>", err)
		}
	}()
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Panic("Read Server Config Error:", err, " filename:", filename)
	}
	cfg := new(ServerConfig)
	err = json.Unmarshal(content, cfg)
	if err != nil {
		log.Panic("Unmarshal Config Error:", err)
	}
	config = cfg
	return true
}

func GetConfig() *ServerConfig {
	configLock.Lock()
	defer configLock.Unlock()
	if config == nil {
		LoadConfig("./config/server.conf")
	}
	return config
}

func ReloadConfig() bool {
	configLock.Lock()
	defer configLock.Unlock()
	return LoadConfig("./config/server.conf")
}

func (this *ServerConfig) GetZoneId() uint32 {
	return this.BaseConfig.ZoneId
}

func (this *ServerConfig) GetPlatform() uint32 {
	return this.BaseConfig.Platform
}

func (this *BaseConfig) Show() {
	if this == nil {
		return
	}
	log.Println("ZoneId:", this.ZoneId)
	log.Println("Platform:", this.Platform)
	log.Println()
}

func (this *DBConfig) Show() {
	if this == nil {
		return
	}
	log.Println("DBIP:", this.DBIP)
	log.Println("DBPort:", this.DBPort)
	log.Println("MinPoolSize:", this.MinPoolSize)
	log.Println("MaxPoolSize:", this.MaxPoolSize)
	log.Println("AcquireIncrement:", this.AcquireIncrement)
	log.Println()
}

func (this *JMXConfig) Show() {
	if this == nil {
		return
	}
	log.Println("JMXIP:", this.JMXIP)
	log.Println("JMXPort:", this.JMXPort)
	log.Println()
}

func (this *LinkConfig) Show() {
	if this == nil {
		return
	}
	log.Println("LinkIP:", this.LinkIP)
	log.Println("LinkPort:", this.LinkPort)
	log.Println()
}

func (this *GlobalConfig) Show() {
	if this == nil {
		return
	}
	log.Println("GlobalIP:", this.GlobalIP)
	log.Println("GlobalPort:", this.GlobalPort)
	log.Println()
}

func (this *RPCConfig) Show() {
	if this == nil {
		return
	}
	log.Println("RPCIP:", this.RPCIP)
	log.Println("RPCPort:", this.RPCPort)
	log.Println()
}

func (this *ServerConfig) Show() {
	if this == nil {
		return
	}
	this.BaseConfig.Show()
	this.DBConfig.Show()
	this.JMXConfig.Show()
	this.LinkConfig.Show()
	this.GlobalConfig.Show()
	this.RPCConfig.Show()
}
