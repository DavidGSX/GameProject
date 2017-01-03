package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"
)

type JMXConfig struct {
	JMXIP   string // JMX监听的IP
	JMXPort int    // JMX监听的端口
}

type LinkConfig struct {
	LinkIP   string // 监听玩家连接的IP
	LinkPort int    // 监听玩家连接的端口
}

type GlobalConfig struct {
	GlobalIP   string // 连接Global的IP
	GlobalPort int    // 连接Global的端口
}

type ServerConfig struct {
	JMXConfig    JMXConfig    // JMX的配置
	LinkConfig   LinkConfig   // Link的配置
	GlobalConfig GlobalConfig // Global的配置
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

func (this *ServerConfig) Show() {
	if this == nil {
		return
	}
	this.JMXConfig.Show()
	this.LinkConfig.Show()
	this.GlobalConfig.Show()
}
