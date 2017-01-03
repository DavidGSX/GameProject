package manager

import (
	"gameproject/server/config"
	"log"
	"net"
	"strconv"
	"sync"
)

var linkMgr *LinkMgr
var linkMgrLock sync.RWMutex

type LinkMgr struct {
	roleId2Links map[uint64]*Link //角色id到连接信息的映射
}

func GetLinkMgr() *LinkMgr {
	linkMgrLock.Lock()
	defer linkMgrLock.Unlock()

	if linkMgr == nil {
		linkMgr = new(LinkMgr)
		linkMgr.roleId2Links = make(map[uint64]*Link)
	}
	return linkMgr
}

func (this *LinkMgr) AddLink(roleId uint64, l *Link) {
	linkMgrLock.Lock()
	defer linkMgrLock.Unlock()

	v, ok := this.roleId2Links[roleId]
	if ok {
		v.Close()
		delete(this.roleId2Links, roleId)
	}

	this.roleId2Links[roleId] = l
}

func (this *LinkMgr) DelLink(roleId uint64) {
	linkMgrLock.Lock()
	defer linkMgrLock.Unlock()

	delete(this.roleId2Links, roleId)
}

func (this *LinkMgr) GetLink(roleId uint64) *Link {
	linkMgrLock.RLock()
	defer linkMgrLock.RUnlock()

	v, ok := this.roleId2Links[roleId]
	if ok {
		return v
	} else {
		return nil
	}
}

func LinkMgrInit(cfg *config.ServerConfig) {
	ip := cfg.LinkConfig.LinkIP
	port := cfg.LinkConfig.LinkPort
	l, err := net.Listen("tcp", ip+":"+strconv.Itoa(port))
	if err != nil {
		log.Fatal("Link Listen Error ", err)
	}
	log.Println("Link Listen ", ip, port)

	defer func() {
		if err := recover(); err != nil {
			log.Println("LinkMgr Exception ", err)
		}
	}()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Panic("Link Accept Error ", err)
		}

		go NewLink(conn).Process()
	}
}
