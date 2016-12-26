package manager

import (
	"gameproject/global/config"
	"gameproject/global/plat"
	"log"
	"sync"
)

var platMgr *PlatMgr
var l sync.Mutex

type PlatMgr struct {
	classMap   map[string]plat.IClass
	configMap  map[int]config.PlatConfig
	platSetMap map[int]*PlatSet
}

func GetPlatMgr() *PlatMgr {
	l.Lock()
	defer l.Unlock()
	if platMgr == nil {
		platMgr = new(PlatMgr)
		platMgr.classMap = make(map[string]plat.IClass)
		platMgr.configMap = make(map[int]config.PlatConfig)
		platMgr.platSetMap = make(map[int]*PlatSet)
	}
	return platMgr
}

func (this *PlatMgr) Init(cfg *config.GlobalConfig) {
	this.InitClass()
	if this.LoadCfg(cfg) != nil {
		log.Panic("PlatMgr.Init Error!")
	}
}

// 所有已实现的渠道，都需要在此处注册
func (this *PlatMgr) InitClass() {
	this.AddClass("AllApp", new(plat.AllApp))
	this.AddClass("OneSdk", new(plat.OneSdk))
	this.AddClass("Laohu", new(plat.Laohu))
}

func (this *PlatMgr) AddClass(name string, class plat.IClass) {
	l.Lock()
	defer l.Unlock()
	if this == nil || name == "" || class == nil {
		log.Panic("PlatMgr.AddClass this or name or class is nil!")
	}

	_, ok := this.classMap[name]
	if ok {
		log.Panic("PlatMgr.AddClass class exist name:", name)
	}

	this.classMap[name] = class
	log.Println("PlatMgr.AddClass class name:", name)
}

// 加载所有的PlatConfigs到Map中，用PlatID作为唯一标记
func (this *PlatMgr) LoadCfg(cfg *config.GlobalConfig) error {
	if cfg == nil {
		log.Panic("PlatMgr.LoadCfg cfg is nil")
	}

	l.Lock()
	defer func() {
		l.Unlock()
		if err := recover(); err != nil {
			log.Println("----------PlatMgr.LoadCfg Error:", err)
		}
	}()

	tmpCfgMap := make(map[int]config.PlatConfig)
	for _, v := range cfg.PlatConfigs {
		id := v.PlatID
		_, ok := tmpCfgMap[id]
		if ok {
			log.Panic("PlatMgr.LoadCfg plat id duplicate, id:", id)
		}
		tmpCfgMap[id] = v
		v.Show("")
	}

	this.LoadPlatSet(tmpCfgMap, cfg)

	this.configMap = tmpCfgMap
	return nil
}

// 加载所有的PlatSetInfo到Map中，并初始化处理的class，用SetID作为唯一标记
func (this *PlatMgr) LoadPlatSet(tmpCfgMap map[int]config.PlatConfig, cfg *config.GlobalConfig) {
	if cfg == nil {
		log.Panic("PlatMgr.LoadPlatSet cfg is nil")
	}

	tmpPlatSet := make(map[int]*PlatSet)
	for _, v := range cfg.PlatSetInfo {
		id := v.SetID
		_, ok := tmpPlatSet[id]
		if ok {
			log.Panic("PlatMgr.LoadPlatSet set id duplicate, id:", id)
		}
		platSet := new(PlatSet)
		platSet.InitCfg(this, tmpCfgMap, &v)
		tmpPlatSet[id] = platSet
		v.Show("")
	}

	this.platSetMap = tmpPlatSet
}

func (this *PlatMgr) ClonePlatByID(tmpCfgMap map[int]config.PlatConfig, id int) (string, string, plat.IClass) {
	v, ok := tmpCfgMap[id]
	if ok == false {
		log.Panic("PlatMgr.ClonePlatByID not exist id:", id)
	}

	c, ok := this.classMap[v.ClassName]
	if ok == false {
		log.Panic("PlatMgr.ClonePlatByID not exist class:", v.ClassName)
	}

	class := c.Clone()
	class.Init(&v)
	return v.Author, v.CallbackUrl, class
}
