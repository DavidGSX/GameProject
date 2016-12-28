package main

import (
	"gameproject/common"
	"gameproject/server/protocol"
	"log"
	"net"

	"github.com/golang/protobuf/proto"
)

func main() {
	goRobot("star")
}

func goRobot(userid string) {
	conn, err := net.Dial("tcp", "127.0.0.1:29000")
	if err != nil {
		log.Panic("connect server error!")
	}

	/*
		login := &protocol.CUserLogin{}
		login.UserId = userid
		login.Token = "123456"
		login.Zoneid = 1001
		login.Platform = protocol.CUserLogin_IOS
		loginData, err := proto.Marshal(login)
		if err != nil {
			log.Panic("marshal login error:", err)
		}
		oct := &common.Octets{}
		oct.CompactUint32(uint32(len(loginData)))
		oct.MarshalBytes(loginData)
		conn.Write(oct.GetBuf())

		readbuf := make([]byte, 1024)
		n, err := conn.Read(readbuf)
		if err != nil {
			log.Panic("read error:", err)
		}

		oct = common.NewOctets(readbuf[:n])
		oct.UncompactUint32()
		buf := oct.UnmarshalBytes()
		loginRes := &protocol.SUserLogin{}
		err = proto.Unmarshal(buf, loginRes)
		if err != nil {
			log.Panic("unmarshal login result error:", err)
		}
		log.Println(loginRes.GetLoginRes())
	*/

	for i := 0; i < 10000; i++ {
		sendInfo := &protocol.CAddMoney{}
		sendInfo.Num = 1
		data, err := proto.Marshal(sendInfo)
		if err != nil {
			log.Panic("marshal data error:", err)
		}
		oct := &common.Octets{}
		oct.CompactUint32(uint32(len(data)))
		oct.MarshalBytes(data)
		conn.Write(oct.GetBuf())
	}

	readbuf := make([]byte, 1024)
	_, err = conn.Read(readbuf)
	if err != nil {
		log.Panic("read error:", err)
	}
}
