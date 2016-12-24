package manager

import (
	"gameproject/global/config"
	"log"
	"net"
	"strconv"
	"sync"
)

func InitJMX(cfg *config.GlobalConfig, wg sync.WaitGroup) {
	ip := cfg.HttpConfig.CallbackIP
	port := cfg.HttpConfig.CallbackPort
	l, err := net.Listen("tcp", ip+":"+strconv.Itoa(port))
	if err != nil {
		log.Fatal("JMX Listen Error:", err)
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println("JMX Error -> ", err)
		}
	}()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Panic("JMX Accept Error:", err)
		}
		go handleAuthor(conn)
	}
}

func handleJMX(conn net.Conn) {
	defer func() {
		log.Println("handleJMX disconnected", conn.RemoteAddr().String())
		conn.Close()
		if err := recover(); err != nil {
			log.Println("handleJMX -> ", err)
		}
	}()
	log.Println("handleJMX connected", conn.RemoteAddr().String())

	buffer := make([]byte, 0)
	for {
		reader := make([]byte, 1024)
		n, err := conn.Read(reader)
		if err != nil {
			log.Panic("Read Error:", err)
		}
		buffer = append(buffer, reader[:n]...)
		//log.Println(n, reader)

	}
}
