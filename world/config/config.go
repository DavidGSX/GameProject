package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"
)

type WorldConfig struct {
	ServerIP     string // 监听Server的IP
	ServerPort   uint32 // 监听Server的端口
	JMXIP        string // JMX监听的IP
	JMXPort      uint32 // JMX监听的端口
	DBIP         string // 数据库的IP
	DBPort       uint32 // 数据库的端口
	MinPoolSize  uint32 // 数据库连接池最少的连接数
	MaxPoolSize  uint32 // 数据库连接池最大的连接数
	AcqIncrement uint32 // 数据库连接池每次增加的数量
}

var config *WorldConfig
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
	cfg := new(WorldConfig)
	err = json.Unmarshal(content, cfg)
	if err != nil {
		log.Panic("Unmarshal Config Error:", err)
	}
	config = cfg
	return true
}

func GetConfig() *WorldConfig {
	configLock.Lock()
	defer configLock.Unlock()
	if config == nil {
		LoadConfig("./config/world.conf")
	}
	return config
}

func ReloadConfig() bool {
	configLock.Lock()
	defer configLock.Unlock()
	return LoadConfig("./config/world.conf")
}

func (this *WorldConfig) Show() {
	if this == nil {
		return
	}
	log.Println("ServerIP:", this.ServerIP)
	log.Println("ServerPort:", this.ServerPort)
	log.Println("JMXIP:", this.JMXIP)
	log.Println("JMXPort:", this.JMXPort)
	log.Println("DBIP:", this.DBIP)
	log.Println("DBPort:", this.DBPort)
	log.Println("MinPoolSize:", this.MinPoolSize)
	log.Println("MaxPoolSize:", this.MaxPoolSize)
	log.Println("AcqIncrement:", this.AcqIncrement)
	log.Println()
}
