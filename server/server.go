package main

import (
	"gameproject/common"
	"gameproject/server/protocol"
	"log"
	"net"
	"strconv"

	"github.com/golang/protobuf/proto"
)

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:29000")
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
			log.Println(total)

			go modifydb(1001, int(addmoney.GetNum()))

			buffer = buffer[oct.Pos():]
		}
	}
}

func modifydb(roleid, num int) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("defer->>>>>>", err)
		}
	}()

	v := common.GetKV(strconv.Itoa(roleid))
	if v == "" {
		v = "0"
	}
	t, err := strconv.Atoi(v)
	if err != nil {
		log.Panic("strconv.Atoi error:", err)
	}
	t += num
	common.SetKV(strconv.Itoa(roleid), strconv.Itoa(t))
}
