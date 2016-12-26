package manager

import (
	"gameproject/global/config"
	"gameproject/global/plat"
)

type PlatSet struct {
	setID       int
	defaultPlat int
	authorMap   map[string]plat.IClass
	callbackMap map[string]plat.IClass
}

func (this *PlatSet) InitCfg(platMgr *PlatMgr, tmpCfgMap map[int]config.PlatConfig, cfg *config.PlatSet) {
	this.authorMap = make(map[string]plat.IClass)
	this.callbackMap = make(map[string]plat.IClass)
	this.setID = cfg.SetID
	this.defaultPlat = cfg.DefaultPlat
	for _, v := range cfg.PlatIDs {
		author, callbackUrl, class := platMgr.ClonePlatByID(tmpCfgMap, v)
		if class != nil {
			this.authorMap[author] = class
			this.callbackMap[callbackUrl] = class
		}
	}
	if this.defaultPlat > 0 {
		author, callbackUrl, class := platMgr.ClonePlatByID(tmpCfgMap, this.defaultPlat)
		if class != nil {
			this.authorMap[author] = class
			this.callbackMap[callbackUrl] = class
		}
	}
}
