package manager

import (
	"gameproject/common"
	"gameproject/global/config"
	"gameproject/global/plat"
	"gameproject/global/protocol"
	"log"
	"net/url"
	"strings"
	"sync"

	"github.com/golang/protobuf/proto"
)

var platMgr *PlatMgr
var platMgrLock sync.RWMutex

type PlatMgr struct {
	classMap    map[string]plat.IClass
	authorMap   map[string]plat.IClass
	callbackMap map[string]plat.IClass
	platSetMap  map[uint32][]string
}

func GetPlatMgr() *PlatMgr {
	platMgrLock.Lock()
	defer platMgrLock.Unlock()
	if platMgr == nil {
		platMgr = new(PlatMgr)
		platMgr.classMap = make(map[string]plat.IClass)
		platMgr.authorMap = make(map[string]plat.IClass)
		platMgr.callbackMap = make(map[string]plat.IClass)
		platMgr.platSetMap = make(map[uint32][]string)
	}
	return platMgr
}

func (this *PlatMgr) Init(cfg *config.GlobalConfig) {
	this.InitClass()
	if this.LoadCfg(cfg) != nil {
		log.Panic("PlatMgr.Init Error!")
	}
}

func (this *PlatMgr) InitClass() {
	plat.InitClass(this)
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

	tmpAuthorMap := make(map[string]plat.IClass)
	tmpCallbackMap := make(map[string]plat.IClass)
	for _, v := range cfg.PlatConfigs {
		c, ok := this.classMap[v.ClassName]
		if ok == false {
			log.Panic("PlatMgr.LoadCfg class not exist:", v.ClassName)
		}

		author := v.Author
		_, ok = tmpAuthorMap[author]
		if ok {
			log.Panic("PlatMgr.LoadCfg plat author duplicate, author:", author)
		}

		callback := v.CallbackUrl
		_, ok = tmpCallbackMap[callback]
		if ok && callback != "" {
			log.Panic("PlatMgr.LoadCfg plat callback duplicate, callback:", callback)
		}

		class := c.Clone()
		class.Init(&v)
		tmpAuthorMap[author] = class
		if callback != "" {
			tmpCallbackMap[callback] = class
		}
	}

	this.LoadPlatSet(cfg)

	this.authorMap = tmpAuthorMap
	this.callbackMap = tmpCallbackMap
	return nil
}

// 加载所有的PlatSetInfo到Map中，用SetID作为唯一标记
func (this *PlatMgr) LoadPlatSet(cfg *config.GlobalConfig) {
	if cfg == nil {
		log.Panic("PlatMgr.LoadPlatSet cfg is nil")
	}

	tmpCfgMap := make(map[uint32]config.PlatConfig)
	for _, v := range cfg.PlatConfigs {
		id := v.PlatID
		_, ok := tmpCfgMap[id]
		if ok {
			log.Panic("PlatMgr.LoadCfg plat id duplicate, id:", id)
		}
		tmpCfgMap[id] = v
	}

	tmpPlatSet := make(map[uint32][]string)
	for _, v := range cfg.PlatSetInfo {
		id := v.SetID
		_, ok := tmpPlatSet[uint32(id)]
		if ok {
			log.Panic("PlatMgr.LoadPlatSet set id duplicate, id:", id)
		}

		plats := make([]string, 0)
		for _, p := range v.PlatIDs {
			cfg, ok := tmpCfgMap[p]
			if ok == false {
				log.Panic("PlatMgr.LoadCfg PlatSet plat not in configure, plat:", p)
			}
			plats = append(plats, cfg.Author)
		}
		tmpPlatSet[uint32(id)] = plats
	}

	this.platSetMap = tmpPlatSet
}

func (this *PlatMgr) GetPlatsBySetId(id uint32) (plats map[string]bool) {
	platMgrLock.Lock()
	defer platMgrLock.Unlock()

	plats = make(map[string]bool)
	v, ok := this.platSetMap[id]
	if ok == false {
		return plats
	}

	for _, p := range v {
		plats[p] = true
	}
	return plats
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

func (this *PlatMgr) ProcessAuthor(gs *Server, plat, userId, token string) {
	class, ok := this.authorMap[plat]
	if ok == false {
		log.Println("PlatMgr.ProcessAuthor plat not exist, plat:", plat)
		return
	}

	res := class.Authorize(userId, token)
	send := &protocol.GSAuthResult{}
	send.UserId = userId + "$" + plat
	send.Plat = plat
	send.Result = res
	data, err := proto.Marshal(send)
	if err != nil {
		log.Println("Marshal GSAuthResult error:", err)
		return
	}

	oct := &common.Octets{}
	oct.MarshalUint32(uint32(len(data)))
	oct.MarshalUint32(3)
	oct.MarshalBytesOnly(data)
	gs.Send(oct.GetBuf())
}
