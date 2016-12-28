package main

import (
	"gameproject/server/protocol"
	"log"
	"net"

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
	for {
		reader := make([]byte, 1024)
		n, err := conn.Read(reader)
		if err != nil {
			log.Panic("Read Error:", err)
		}
		buffer = append(buffer, reader[:n]...)
		log.Println(n, reader)

		login := &protocol.CUserLogin{}
		err = proto.Unmarshal(buffer, login)
		if err != nil {
			log.Panic("unmarshal login error:", err)
		}
		log.Println("UserId:", login.UserId)
		log.Println("Token:", login.Token)
		log.Println("Zoneid:", login.Zoneid)
		log.Println("Platform:", login.Platform)

		loginRes := new(protocol.SUserLogin)
		loginRes.LoginRes = protocol.SUserLogin_SUCCESS

		loginResData, err := proto.Marshal(loginRes)
		if err != nil {
			log.Panic("marshal login result error:", err)
		}

		conn.Write(loginResData)
	}
}
