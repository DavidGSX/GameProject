package manager

import (
	"gameproject/global/config"
	"gameproject/global/plat"
	"log"
	"net/url"
	"strings"
	"sync"
)

var platMgr *PlatMgr
var platMgrLock sync.RWMutex

type PlatMgr struct {
	classMap    map[string]plat.IClass
	configMap   map[int]config.PlatConfig
	authorMap   map[string]plat.IClass
	callbackMap map[string]plat.IClass
	platSetMap  map[uint32]config.PlatSet
}

func GetPlatMgr() *PlatMgr {
	platMgrLock.Lock()
	defer platMgrLock.Unlock()
	if platMgr == nil {
		platMgr = new(PlatMgr)
		platMgr.classMap = make(map[string]plat.IClass)
		platMgr.configMap = make(map[int]config.PlatConfig)
		platMgr.authorMap = make(map[string]plat.IClass)
		platMgr.callbackMap = make(map[string]plat.IClass)
		platMgr.platSetMap = make(map[uint32]config.PlatSet)
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

	platMgrLock.Lock()
	defer func() {
		platMgrLock.Unlock()
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

	tmpAuthorMap := make(map[string]plat.IClass)
	tmpCallbackMap := make(map[string]plat.IClass)
	for _, v := range cfg.PlatConfigs {
		author := v.Author
		_, ok := tmpAuthorMap[author]
		if ok {
			log.Panic("PlatMgr.LoadCfg plat author duplicate, author:", author)
		}

		callback := v.CallbackUrl
		_, ok = tmpCallbackMap[callback]
		if ok && callback != "" {
			log.Panic("PlatMgr.LoadCfg plat callback duplicate, callback:", callback)
		}

		c, ok := this.classMap[v.ClassName]
		if ok == false {
			log.Panic("PlatMgr.LoadCfg class not exist:", v.ClassName)
		}

		class := c.Clone()
		class.Init(&v)
		tmpAuthorMap[author] = class
		if callback != "" {
			tmpCallbackMap[callback] = class
		}
	}

	this.LoadPlatSet(cfg)

	this.configMap = tmpCfgMap
	this.authorMap = tmpAuthorMap
	this.callbackMap = tmpCallbackMap
	return nil
}

// 加载所有的PlatSetInfo到Map中，并初始化处理的class，用SetID作为唯一标记
func (this *PlatMgr) LoadPlatSet(cfg *config.GlobalConfig) {
	if cfg == nil {
		log.Panic("PlatMgr.LoadPlatSet cfg is nil")
	}

	tmpPlatSet := make(map[uint32]config.PlatSet)
	for _, v := range cfg.PlatSetInfo {
		id := v.SetID
		_, ok := tmpPlatSet[uint32(id)]
		if ok {
			log.Panic("PlatMgr.LoadPlatSet set id duplicate, id:", id)
		}
		tmpPlatSet[uint32(id)] = v
		v.Show("")
	}

	this.platSetMap = tmpPlatSet
}

func (this *PlatMgr) ProcessCallback(path string, param url.Values) string {
	platMgrLock.RLock()
	defer platMgrLock.RUnlock()

	log.Println("path:", path, " param:", param)
	callback := strings.TrimPrefix(path, "/")
	class, ok := this.callbackMap[callback]
	if ok == false {
		log.Println("PlatMgr.ProcessCallback plat not exist, callback:", callback)
		return string("Invalid URI.Path " + path)
	}

	return class.Callback(param)
}

func (this *PlatMgr) ProcessAuthor(plats uint32, userId, token string) {
	platSet, ok := this.platSetMap[plats]
	if ok == false {
		log.Println("ProcessAuthor wrong plats:", plats)
		return
	}

	strs := strings.Split(userId, "$")
	if len(strs) != 2 {
		log.Println("ProcessAuthor wrong userId:", userId)
		return
	}

	class, ok := this.authorMap[strs[1]]
	if ok == false {
		log.Println("ProcessAuthor wrong user plat:", strs[1])
		return
	}

	// todo
	_ = platSet
	_ = class
	log.Println("to do")
}
