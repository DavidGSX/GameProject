// test.go
package main

import (
	"gameproject/common"
	"log"
	"net"
	"runtime"
	"strconv"
	"sync/atomic"
	"time"
	//"unicode/utf16"
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

	runtime.GOMAXPROCS(1)
	v := string("123456789012345678901234567890123456789012345678901234567890")
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 1e9; j++ {
				save2db(strconv.Itoa(i*10000+j), v)
			}
		}()
	}

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

	buffer := make([]byte, 0)
	for {
		reader := make([]byte, 128)
		n, err := conn.Read(reader)
		if err != nil {
			log.Panic("Read Error:", err)
		}

		//log.Println(n, reader)

		buffer = append(buffer, reader[:n]...)

		for len(buffer) > 4 {
			oct := common.NewOctets(buffer)
			size := oct.UnmarshalUint32()
			//log.Println("size:", size)
			if len(buffer) < int(size+oct.Pos()) {
				break
			}

			key := oct.UnmarshalString()
			//log.Println("key:", string(utf16.Decode(key)))
			//state :=
			oct.UnmarshalUint32()
			//log.Println("state:", state)
			//localid :=
			oct.UnmarshalUint32()
			//log.Println("localid:", localid)
			//ip :=
			oct.UnmarshalUint16s()
			//log.Println("ip:", string(utf16.Decode(ip)))
			//tm :=
			oct.UnmarshalUint64()
			//log.Println("tm:", tm)

			//save2db((key), string(buffer[:oct.Pos()]))
			getfromdb(key)
			//save2db(string(utf16.Decode(key)), string(buffer[:oct.Pos()]))
			//getfromdb(string(utf16.Decode(key)))

			//log.Println("buffer len", len(buffer), "pos", oct.Pos())
			buffer = buffer[oct.Pos():]
		}
		//sendoct := new(common.Octets)
		//sendoct.MarshalUint32(1) //result
		//writer := sendoct.GetBuf()
		//n, err = conn.Write(writer)
		//if err != nil {
		//	log.Panic("Write Error:", err)
		//}
	}
}

var count uint32
var exc10 uint32
var exc100 uint32
var notfind uint32

func save2db(k, v string) {
	c, err := common.GetDBPool().NewClient()
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		c.Close()
		if err := recover(); err != nil {
			log.Println("defer->>>>>>", err)
		}
	}()
	now := time.Now()
	err = c.Set(k, v)
	if err != nil {
		log.Println("------->>>>>>>> set error ", err)
		atomic.AddUint32(&notfind, 1)
	}
	use := time.Since(now)
	if use > 10*time.Millisecond {
		//log.Println("time exceed", use)
		atomic.AddUint32(&exc10, 1)
	}
	if use > 100*time.Millisecond {
		log.Println("------------------------time exceed", use)
		atomic.AddUint32(&exc100, 1)
	}
	x := atomic.AddUint32(&count, 1)
	if x%100000 == 0 {
		log.Println("count", count, "exceed 10ms", exc10, "exceed 100ms", exc100, "not find", notfind)
	}
}

func getfromdb(k string) {
	c, err := common.GetDBPool().NewClient()
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		c.Close()
		if err := recover(); err != nil {
			log.Println("defer->>>>>>", err)
		}
	}()
	now := time.Now()
	v, err := c.Get(k)
	if err != nil {
		log.Println("------------------------get error, k=", k)
		return
	}
	if v.String() == string("") {
		atomic.AddUint32(&notfind, 1)
	}
	use := time.Since(now)
	if use > 10*time.Millisecond {
		log.Println("time exceed", use)
		atomic.AddUint32(&exc10, 1)
	}
	if use > 100*time.Millisecond {
		log.Println("------------------------time exceed", use)
		atomic.AddUint32(&exc100, 1)
	}
	x := atomic.AddUint32(&count, 1)
	if x%10000 == 0 {
		log.Println("count", count, "exceed 10ms", exc10, "exceed 100ms", exc100, "not find", notfind, v)
	}
}

/*
import (
	"gameproject/common"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func main() {

	for {
		v := string("123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890")
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		go save2db(r, v)
		time.Sleep(1 * time.Millisecond)
	}
}

func save2db(r *rand.Rand, v string) {

	c, err := common.GetDBPool().NewClient()
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		defer c.Close()
		if err := recover(); err != nil {
			log.Println("defer->>>>>>", err)
		}
	}()

	for i := 0; i < 10; i++ {
		k := r.Intn(100000000)
		now := time.Now()
		c.Set(strconv.Itoa(k), v)
		use := time.Since(now)
		if use > 10*time.Millisecond {
			log.Println("time exceed", use)
		}
		if use > 100*time.Millisecond {
			log.Println("------------------------time exceed", use)
		}
	}
}
*/
