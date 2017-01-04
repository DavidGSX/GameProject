package manager

import (
	"gameproject/global/config"
	"log"
	"net"
	"strconv"
	"sync"
)

var authorMgr *AuthorMgr
var authorMgrLock sync.RWMutex

type AuthorMgr struct {
	zoneId2Links map[uint32]*Server //ZoneId到连接信息的映射
}

func GetAuthorMgr() *AuthorMgr {
	authorMgrLock.Lock()
	defer authorMgrLock.Unlock()

	if authorMgr == nil {
		authorMgr = new(AuthorMgr)
		authorMgr.zoneId2Links = make(map[uint32]*Server)
	}
	return authorMgr
}

func (this *AuthorMgr) AddLink(zoneId uint32, l *Server) {
	authorMgrLock.Lock()
	defer authorMgrLock.Unlock()

	v, ok := this.zoneId2Links[zoneId]
	if ok {
		v.Close()
		delete(this.zoneId2Links, zoneId)
	}

	this.zoneId2Links[zoneId] = l
}

func (this *AuthorMgr) DelLink(zoneId uint32) {
	authorMgrLock.Lock()
	defer authorMgrLock.Unlock()

	delete(this.zoneId2Links, zoneId)
}

func (this *AuthorMgr) GetLink(zoneId uint32) *Server {
	authorMgrLock.RLock()
	defer authorMgrLock.RUnlock()

	v, ok := this.zoneId2Links[zoneId]
	if ok {
		return v
	} else {
		return nil
	}
}

func InitAuthor(cfg *config.GlobalConfig) {
	ip := cfg.HttpConfig.AuthorizeIP
	port := cfg.HttpConfig.AuthorizePort
	l, err := net.Listen("tcp", ip+":"+strconv.Itoa(port))
	if err != nil {
		log.Fatal("Author Listen Error:", err)
	}
	log.Println("Author Listen ", ip, port)
	defer func() {
		if err := recover(); err != nil {
			log.Println("Author Error -> ", err)
		}
	}()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Panic("Author Accept Error:", err)
		}
		go NewLink(conn).Process()
	}
}
