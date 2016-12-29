package main

import (
	"gameproject/common"
	"gameproject/server/protocol"
	"log"
	"net"
	"time"

	"github.com/golang/protobuf/proto"
)

func main() {
	goRobot("star")
}

func goRobot(userid string) {

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

	for c := 1; c <= 100; c++ {
		go AddMoney(int64(c), 1)
	}

	<-time.After(3e10) // 30秒后退出
	log.Println("all finish")
}

func AddMoney(roleId int64, num int32) {
	conn, err := net.Dial("tcp", "127.0.0.1:29000")
	if err != nil {
		log.Panic("connect server error!")
	}

	for i := 1; i <= 10000; i++ {
		sendInfo := &protocol.CAddMoney{}
		sendInfo.RoleId = roleId
		sendInfo.Num = num
		data, err := proto.Marshal(sendInfo)
		if err != nil {
			log.Panic("marshal data error:", err)
		}
		oct := &common.Octets{}
		oct.CompactUint32(uint32(len(data)))
		oct.MarshalBytes(data)
		conn.Write(oct.GetBuf())
	}

	<-time.After(1e10) // 10秒后退出
	log.Println("roldId:", roleId, " finish")
}
