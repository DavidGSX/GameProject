package main

import (
	"gameproject/server/config"
	"gameproject/server/manager"
	"sync"
)

var wg sync.WaitGroup

func main() {
	cfg := config.GetConfig()
	cfg.Show()

	go manager.LinkMgrInit(cfg)
	//go manager.GlobalMgrInit(cfg)
	go manager.JMXInit(cfg, &wg)

	wg.Add(1)
	wg.Wait()
}

/*
import (
	"gameproject/common"
	"gameproject/server/manager"
	"gameproject/server/protocol"
	"log"
	"net"
	"strconv"

	"github.com/golang/protobuf/proto"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:29000")
	if err != nil {
		log.Fatal("Server Listen Error:", err)
	}
	defer func() {
		if err := recover(); err != nil {
			log.Println("Server Error -> ", err)
		}
	}()
	common.GetDBPool()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Panic("Linker Accept Error:", err)
		}
		go handleLinker(conn)
	}
}

func handleLinker(conn net.Conn) {
	defer func() {
		log.Println("handleLinker disconnected", conn.RemoteAddr().String())
		conn.Close()
		if err := recover(); err != nil {
			log.Println("handleLinker -> ", err)
		}
	}()
	log.Println("handleLinker connected", conn.RemoteAddr().String())

	buffer := make([]byte, 0)
	total := int32(0)
	for {
		reader := make([]byte, 1024)
		n, err := conn.Read(reader)
		if err != nil {
			log.Panic("Read Error:", err)
		}
		buffer = append(buffer, reader[:n]...)
		//log.Println(len(buffer), buffer)

		for len(buffer) >= 4 {
			oct := common.NewOctets(buffer)
			size := oct.UncompactUint32()
			if oct.Remain() < int(size) {
				break
			}

			data := oct.UnmarshalBytes()
			addmoney := &protocol.CAddMoney{}
			err = proto.Unmarshal(data, addmoney)
			if err != nil {
				log.Panic("unmarshal login error:", err)
			}

			total += addmoney.GetNum()
			Modifydb(int(addmoney.GetRoleId()), "money", int(addmoney.GetNum()))

			buffer = buffer[oct.Pos():]
		}
	}
}

func Modifydb(roleid int, table string, param int) {
	k := table + strconv.Itoa(roleid)
	manager.GetLockMgr().Lock(k)
	defer func() {
		manager.GetLockMgr().Unlock(k)
		if err := recover(); err != nil {
			log.Println("defer->>>>>>", err)
		}
	}()

	v := common.GetKV(k)
	if v == "" {
		v = "0"
	}
	t, err := strconv.Atoi(v)
	if err != nil {
		log.Panic("strconv.Atoi error:", err)
	}
	t += param
	common.SetKV(k, strconv.Itoa(t))
}

*/
