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
	for c := 100; c < 160; c++ {
		go Robot("star"+strconv.Itoa(c)+"$apps", strconv.Itoa(c))
	}

	<-time.After(3e10) // 30秒后退出
	log.Println("all finish")
}

func Robot(userId, token string) {
	conn, err := net.Dial("tcp", "127.0.0.1:29000")
	if err != nil {
		log.Panic("connect server error!")
	}

	if UserLogin(conn, userId, token) == true {
		roleId := GetRoleId(conn, userId)
		if roleId > 0 {
			AddMoney(conn, roleId, 1)
		} else {
			log.Println("UserId:", userId, " GetRoleId Failed!")
		}
	} else {
		log.Println("UserId:", userId, " UserLogin Failed!")
	}

	<-time.After(1e10) // 10秒后退出
}

func ConnSend(conn net.Conn, msgType uint32, msg proto.Message) {
	data, err := proto.Marshal(msg)
	if err != nil {
		log.Panic("marshal error:", err)
	}
	oct := &common.Octets{}
	oct.MarshalUint32(uint32(len(data)))
	oct.MarshalUint32(msgType)
	oct.MarshalBytesOnly(data)
	conn.Write(oct.GetBuf())
}

func ConnRead(conn net.Conn, msg proto.Message) (msgType uint32) {
	readbuf := make([]byte, 1024)
	n, err := conn.Read(readbuf)
	if err != nil {
		log.Panic("read error:", err)
	}

	oct := common.NewOctets(readbuf[:n])
	size := oct.UnmarshalUint32()
	msgType = oct.UnmarshalUint32()
	data := oct.UnmarshalBytesOnly(size)

	err = proto.Unmarshal(data, msg)
	if err != nil {
		log.Panic("unmarshal result error:", err)
	}
	return msgType
}

func UserLogin(conn net.Conn, userId, token string) bool {
	sendInfo := &protocol.CUserLogin{}
	sendInfo.UserId = userId
	sendInfo.Token = token
	sendInfo.ZoneId = 1001
	sendInfo.Platform = protocol.CUserLogin_IOS
	ConnSend(conn, 1001, sendInfo)

	recvInfo := &protocol.SUserLogin{}
	msgType := ConnRead(conn, recvInfo)
	log.Println("MsgType", msgType, " recvInfo:", recvInfo, " UserId:", userId, " Token:", token)

	if recvInfo.GetLoginRes() == protocol.SUserLogin_SUCCESS {
		return true
	} else {
		return false
	}
}

func GetRoleId(conn net.Conn, userId string) (roleId uint64) {
	sendInfo := &protocol.CRoleList{}
	sendInfo.SelectRoleId = 0
	ConnSend(conn, 1003, sendInfo)

	recvInfo := &protocol.SRoleList{}
	msgType := ConnRead(conn, recvInfo)
	log.Println("MsgType", msgType, " recvInfo:", recvInfo)

	if len(recvInfo.GetRoles()) > 0 {
		return recvInfo.GetRoles()[0].RoleId
	} else {
		//创建角色
		sendInfo := &protocol.CCreateRole{}
		sendInfo.Name = "robot" + userId
		sendInfo.School = 1
		sendInfo.Sex = 2
		ConnSend(conn, 1005, sendInfo)

		recvInfo := &protocol.SCreateRole{}
		msgType := ConnRead(conn, recvInfo)
		log.Println("MsgType", msgType, " recvInfo:", recvInfo)

		if recvInfo.Res == protocol.SCreateRole_SUCCESS {
			return recvInfo.Info.RoleId
		} else {
			return 0
		}
	}
}

func AddMoney(conn net.Conn, roleId uint64, num uint32) {
	for i := 0; i < 1000; i++ {
		sendInfo := &protocol.CAddMoney{}
		sendInfo.RoleId = roleId
		sendInfo.Num = num
		ConnSend(conn, 1007, sendInfo)

		recvInfo := &protocol.SMoneyInfo{}
		msgType := ConnRead(conn, recvInfo)
		log.Println("MsgType", msgType, " recvInfo:", recvInfo)
	}
}
