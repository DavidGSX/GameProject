package main

import (
	"gameproject/server/protocol"
	"log"
	"net"

	"github.com/golang/protobuf/proto"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:29000")
	if err != nil {
		log.Panic("connect server error!")
	}

	login := new(protocol.CUserLogin)
	login.UserId = "star"
	login.Token = "123456"
	login.Zoneid = 1001
	login.Platform = protocol.CUserLogin_IOS

	loginData, err := proto.Marshal(login)
	if err != nil {
		log.Panic("marshal login error:", err)
	}

	conn.Write(loginData)

	for {
		readbuf := make([]byte, 1024)
		n, err := conn.Read(readbuf)
		if err != nil {
			log.Panic("read error:", err)
		}
		log.Println(n, readbuf[:n])

		loginRes := &protocol.SUserLogin{}
		err = proto.Unmarshal(readbuf[:n], loginRes)
		if err != nil {
			log.Panic("unmarshal login result error:", err)
		}
		log.Println(loginRes.GetLoginRes())
	}
}
