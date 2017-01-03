package manager

import (
	"gameproject/server/config"
	"log"
	"net"
	"strconv"
	"sync"
)

var globalConn *GlobalConn
var globalConnLock sync.RWMutex

type GlobalConn struct {
	conn    net.Conn
	recvBuf []byte
	sendBuf []byte
}

func GlobalMgrInit(cfg *config.ServerConfig) {
	ip := cfg.GlobalConfig.GlobalIP
	port := cfg.GlobalConfig.GlobalPort
	conn, err := net.Dial("tcp", ip+":"+strconv.Itoa(port))
	if err != nil {
		log.Fatal("Connect Global Error:", err)
	}
	log.Println("Connect Global Success ", ip, port)

	globalConnLock.Lock()
	defer func() {
		globalConnLock.Unlock()
		if err := recover(); err != nil {
			log.Println("Global Init Exception -> ", err)
		}
	}()

	if globalConn == nil {
		globalConn = new(GlobalConn)
	} else if globalConn.conn != nil {
		globalConn.conn.Close()
	}

	globalConn.conn = conn
	globalConn.recvBuf = make([]byte, 0)
	globalConn.sendBuf = make([]byte, 0)

	go GlobalReadProcess()
	go GlobalWriteProcess()
}

func GlobalReadProcess() {

}

func GlobalWriteProcess() {

}
