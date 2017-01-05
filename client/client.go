package main

import (
	"gameproject/common"
	"gameproject/server/protocol"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/golang/protobuf/proto"
)

func main() {
	//AddMoney(1, 2)
	for c := 100; c <= 200; c++ {
		go UserLogin("star"+strconv.Itoa(c)+"$apps", strconv.Itoa(c))
	}

	<-time.After(3e11) // 30秒后退出
	log.Println("all finish")
}

func UserLogin(userId, token string) {
	conn, err := net.Dial("tcp", "127.0.0.1:29000")
	if err != nil {
		log.Panic("connect server error!")
	}

	for i := 1; i <= 10000; i++ {
		login := &protocol.CUserLogin{}
		login.UserId = userId
		login.Token = token
		login.ZoneId = 1001
		login.Platform = protocol.CUserLogin_IOS
		data, err := proto.Marshal(login)
		if err != nil {
			log.Panic("marshal CUserLogin error:", err)
		}
		oct := &common.Octets{}
		oct.MarshalUint32(uint32(len(data)))
		oct.MarshalUint32(1001)
		oct.MarshalBytesOnly(data)
		conn.Write(oct.GetBuf())

		readbuf := make([]byte, 1024)
		n, err := conn.Read(readbuf)
		if err != nil {
			log.Panic("read error:", err)
		}

		oct = common.NewOctets(readbuf[:n])
		size := oct.UnmarshalUint32()
		msgType := oct.UnmarshalUint32()
		buf := oct.UnmarshalBytesOnly(int(size))
		loginRes := &protocol.SUserLogin{}
		err = proto.Unmarshal(buf, loginRes)
		if err != nil {
			log.Panic("unmarshal login result error:", err)
		}
		log.Println("UserId:", userId, " Token:", token, " Result:", loginRes.GetLoginRes(), " Size:", size, " MsgType", msgType)
	}

	<-time.After(1e10) // 10秒后退出
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
		oct.MarshalUint32(uint32(len(data)))
		oct.MarshalUint32(1005)
		oct.MarshalBytesOnly(data)
		conn.Write(oct.GetBuf())
	}

	<-time.After(1e10) // 10秒后退出
	log.Println("roldId:", roleId, " finish")
}
