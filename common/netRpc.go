package common

import (
	"log"
	"net"
)

func NetRpc(ip, port string) {
	ln, err := net.Listen("tcp", ip+":"+port)
	if err != nil {
		log.Fatal("NetRpc Listen Error:", err)
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println("NetRpc defer recover:", err)
		}
	}()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Panic("NetRpc Accept Error:", err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func() {
		log.Println("Client disconnected", conn.RemoteAddr().String())
		conn.Close()
		if err := recover(); err != nil {
			log.Println("handleConnection defer recover:", err)
		}
	}()

	log.Println("Client connected", conn.RemoteAddr().String())

	l := NewLinker(conn)
	for {
		rbuf := make([]byte, 1024)
		n, err := conn.Read(rbuf)
		if err != nil {
			log.Panic("Read Error:", err)
		}

		log.Println(n, rbuf[:n])
		l.Receive(rbuf[:n])
	}
}
