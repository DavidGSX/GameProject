package main

import (
	"gameproject/common"
	"gameproject/global"
)

func main() {
	go common.NetRpc("0.0.0.0", "1234")
	common.GetProtoMgr().Add(1, new(global.Exist))
	common.GetProtoMgr().Add(2, new(global.Allocate))
	common.GetProtoMgr().Add(3, new(global.Confirm))
	common.GetProtoMgr().Add(4, new(global.Release))
	select {}
}

/*
import (
	"common"
	"log"
	"net"
	"unicode/utf16"
)

func main() {
	ln, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Listen Error:", err)
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println("main defer->>>>>>", err)
		}
	}()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Panic("Accept Error:", err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func() {
		log.Println("Client disconnected", conn.RemoteAddr().String())
		conn.Close()
		if err := recover(); err != nil {
			log.Println("defer->>>>>>", err)
		}
	}()

	log.Println("Client connected", conn.RemoteAddr().String())

	for {
		reader := make([]byte, 64)
		n, err := conn.Read(reader)
		if err != nil {
			log.Panic("Read Error:", err)
		}

		log.Println(n, reader)

		oct := common.NewOctets(reader)
		t := oct.UncompactUint32()
		log.Println("type:", t)
		log.Println("size:", oct.UncompactUint32())
		sid := oct.UnmarshalUint32()
		sid = sid & 0x7fffffff
		log.Println("sid:", sid)
		log.Println("group:", oct.UnmarshalString())
		name := oct.UnmarshalUint16s()
		log.Println("name:", string(utf16.Decode(name)))
		log.Println("localid:", oct.UnmarshalUint32())
		log.Println("reserve:", oct.UnmarshalBytes())

		sendoct := new(common.Octets)
		sendoct.CompactUint32(t)   //type
		sendoct.CompactUint32(8)   //size
		sendoct.MarshalUint32(sid) //sid
		sendoct.MarshalUint32(2)   //result
		writer := sendoct.GetBuf()
		n, err = conn.Write(writer)
		if err != nil {
			log.Panic("Write Error:", err)
		}
		log.Println(len(writer), writer)
	}
}
*/
