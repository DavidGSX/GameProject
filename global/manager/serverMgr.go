package manager

import (
	"gameproject/global/config"
	"log"
	"net"
	"strconv"
	"sync"
)

var serverMgr *ServerMgr
var serverMgrLock sync.RWMutex

type ServerMgr struct {
	zoneId2Server map[uint32]*Server //ZoneId到GS连接信息的映射
}

func GetServerMgr() *ServerMgr {
	serverMgrLock.Lock()
	defer serverMgrLock.Unlock()

	if serverMgr == nil {
		serverMgr = new(ServerMgr)
		serverMgr.zoneId2Server = make(map[uint32]*Server)
	}
	return serverMgr
}

func (this *ServerMgr) AddServer(zoneId uint32, l *Server) {
	serverMgrLock.Lock()
	defer serverMgrLock.Unlock()

	v, ok := this.zoneId2Server[zoneId]
	if ok {
		v.Close()
		delete(this.zoneId2Server, zoneId)
	}

	this.zoneId2Server[zoneId] = l
}

func (this *ServerMgr) DelServer(zoneId uint32) {
	serverMgrLock.Lock()
	defer serverMgrLock.Unlock()

	delete(this.zoneId2Server, zoneId)
}

func (this *ServerMgr) GetServer(zoneId uint32) *Server {
	serverMgrLock.RLock()
	defer serverMgrLock.RUnlock()

	v, ok := this.zoneId2Server[zoneId]
	if ok {
		return v
	} else {
		return nil
	}
}

func ServerMgrInit(cfg *config.GlobalConfig) {
	ip := cfg.BaseConfig.ServerIP
	port := cfg.BaseConfig.ServerPort
	l, err := net.Listen("tcp", ip+":"+strconv.Itoa(int(port)))
	if err != nil {
		log.Fatal("Server Listen Error:", err)
	}
	log.Println("Server Listen ", ip, port)
	defer func() {
		if err := recover(); err != nil {
			log.Println("Server Error -> ", err)
		}
	}()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Panic("Server Accept Error:", err)
		}
		go NewServer(conn).Process()
	}
}
