package main

import (
	"gameproject/common"
	"gameproject/server/client/csproto"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/golang/protobuf/proto"
)

func main() {
	ch := make(chan int, 100)
	for c := 1; c < 1e9; c++ {
		ch <- c
		go Robot(ch, "star"+strconv.Itoa(c)+"$apps", strconv.Itoa(c))
	}
	//<-time.After(3e10) // 30秒后退出
	log.Println("all finish")
}

func Robot(ch chan int, userId, token string) {
	conn, err := net.Dial("tcp", "127.0.0.1:29000")
	if err != nil {
		log.Panic("connect server error!")
	}
	defer conn.Close()

	if UserLogin(conn, userId, token) == true {
		roleId := GetRoleId(conn, userId)
		if roleId > 0 {
			EnterWorld(conn, roleId)
			AddMoney(conn, 1)
			AddLevel(conn, 2)
			ReqServerRole(conn)
		} else {
			log.Println("UserId:", userId, " GetRoleId Failed!")
		}
	} else {
		log.Println("UserId:", userId, " UserLogin Failed!")
	}

	<-time.After(1e6) // 1毫秒后退出
	log.Println("Current Process ------>", <-ch)
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
	sendInfo := &csproto.CUserLogin{}
	sendInfo.UserId = userId
	sendInfo.Token = token
	sendInfo.ZoneId = 1001
	sendInfo.Platform = csproto.CUserLogin_IOS
	ConnSend(conn, 1001, sendInfo)

	recvInfo := &csproto.SUserLogin{}
	ConnRead(conn, recvInfo)
	//log.Println("MsgType", msgType, " recvInfo:", recvInfo, " UserId:", userId, " Token:", token)

	if recvInfo.GetLoginRes() == csproto.SUserLogin_SUCCESS {
		return true
	} else {
		return false
	}
}

func GetRoleId(conn net.Conn, userId string) (roleId uint64) {
	sendInfo := &csproto.CRoleList{}
	sendInfo.SelectRoleId = 0
	ConnSend(conn, 1003, sendInfo)

	recvInfo := &csproto.SRoleList{}
	ConnRead(conn, recvInfo)
	//log.Println("MsgType", msgType, " recvInfo:", recvInfo)

	if len(recvInfo.GetRoles()) > 0 {
		return recvInfo.GetRoles()[0].RoleId
	} else {
		//创建角色
		sendInfo := &csproto.CCreateRole{}
		sendInfo.Name = "robot" + userId
		sendInfo.School = 1
		sendInfo.Sex = 2
		ConnSend(conn, 1005, sendInfo)

		recvInfo := &csproto.SCreateRole{}
		ConnRead(conn, recvInfo)
		//log.Println("MsgType", msgType, " recvInfo:", recvInfo)

		if recvInfo.Res == csproto.SCreateRole_SUCCESS {
			return recvInfo.Info.RoleId
		} else {
			return 0
		}
	}
}

func EnterWorld(conn net.Conn, roleId uint64) {
	sendInfo := &csproto.CEnterWorld{}
	sendInfo.RoleId = roleId
	ConnSend(conn, 1007, sendInfo)

	recvInfo := &csproto.SEnterWorld{}
	ConnRead(conn, recvInfo)
	//log.Println("MsgType", msgType, " recvInfo:", recvInfo)
}

func AddMoney(conn net.Conn, num uint32) {
	for i := 0; i < 10; i++ {
		sendInfo := &csproto.CAddMoney{}
		sendInfo.Num = num
		ConnSend(conn, 1009, sendInfo)

		recvInfo := &csproto.SMoneyInfo{}
		ConnRead(conn, recvInfo)
		//log.Println("MsgType", msgType, " recvInfo:", recvInfo)
	}
}

func AddLevel(conn net.Conn, num uint32) {
	sendInfo := &csproto.CAddLevel{}
	sendInfo.Num = num
	ConnSend(conn, 1011, sendInfo)

	recvInfo := &csproto.SLevelInfo{}
	ConnRead(conn, recvInfo)
	//log.Println("MsgType", msgType, " recvInfo:", recvInfo)
}

func ReqServerRole(conn net.Conn) {
	sendInfo := &csproto.CReqServerRoleInfos{}
	ConnSend(conn, 1013, sendInfo)

	recvInfo := &csproto.SServerRoleInfos{}
	ConnRead(conn, recvInfo)
	//log.Println("MsgType", msgType, " recvInfo:", recvInfo)
}
