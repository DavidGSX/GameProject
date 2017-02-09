package manager

import (
	"gameproject/server/config"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"
)

var jmxWG *sync.WaitGroup

func JMXInit(cfg *config.ServerConfig, wg *sync.WaitGroup) {
	ip := cfg.JMXConfig.JMXIP
	port := cfg.JMXConfig.JMXPort
	l, err := net.Listen("tcp", ip+":"+strconv.Itoa(int(port)))
	if err != nil {
		log.Fatal("JMX Listen Error:", err)
	}
	log.Println("JMX Listen ", ip, port)
	jmxWG = wg
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
		go handleJMX(conn)
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

	for {
		reader := make([]byte, 1024)
		n, err := conn.Read(reader)
		if err != nil {
			log.Panic("Read Error:", err)
		}

		if n > 0 {
			cmd := string(reader[:n])
			cmd = strings.Trim(cmd, "\r\n")
			log.Println(n, reader[:n], cmd)
			switch strings.ToLower(cmd) {
			case "stop":
				jmxWG.Done()
			case "reload":
				if config.ReloadConfig() == true {
					log.Println(cmd + " success!")
					conn.Write([]byte(cmd + " success!\n"))
				} else {
					log.Println(cmd + " failed!")
					conn.Write([]byte(cmd + " failed!\n"))
				}
			case "quit":
				log.Panic("JMX Quit")
			default:
				conn.Write([]byte("Unknow Command " + cmd + "\n"))
			}
		}
	}
}
