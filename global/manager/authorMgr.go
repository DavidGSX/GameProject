package manager

import (
	"gameproject/global/config"
	"log"
	"net"
	"strconv"
)

func InitAuthor(cfg *config.GlobalConfig) {
	ip := cfg.HttpConfig.AuthorizeIP
	port := cfg.HttpConfig.AuthorizePort
	l, err := net.Listen("tcp", ip+":"+strconv.Itoa(port))
	if err != nil {
		log.Fatal("Author Listen Error:", err)
	}
	log.Println("Author Listen ", ip, port)
	defer func() {
		if err := recover(); err != nil {
			log.Println("Author Error -> ", err)
		}
	}()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Panic("Author Accept Error:", err)
		}
		go handleAuthor(conn)
	}
}

func handleAuthor(conn net.Conn) {
	defer func() {
		log.Println("handleAuthor disconnected", conn.RemoteAddr().String())
		conn.Close()
		if err := recover(); err != nil {
			log.Println("handleAuthor -> ", err)
		}
	}()
	log.Println("handleAuthor connected", conn.RemoteAddr().String())

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
