package manager

import (
	"gameproject/common"
	"gameproject/world/message"
	"log"
	"net"
	"sync"
	"time"
)

type Server struct {
	conn     net.Conn
	recvBuf  []byte
	sendBuf  []byte
	sendLock sync.Mutex
	lastTime time.Time
	zoneId   uint32
	plats    map[string]bool
}

func NewServer(c net.Conn) *Server {
	l := new(Server)
	l.conn = c
	l.recvBuf = make([]byte, 0)
	l.sendBuf = make([]byte, 0)
	l.lastTime = time.Now()
	return l
}

func (this *Server) Close() {
	this.conn.Close()
	GetServerMgr().DelServer(this.zoneId)
}

func (this *Server) SetZoneId(z uint32) {
	if this.zoneId != 0 && this.zoneId != z {
		log.Println("SetZoneId old:", this.zoneId, " new:", z)
	}
	this.zoneId = z
	log.Println("Server SetZoneId ", this.zoneId, this.conn.RemoteAddr().String())
	GetServerMgr().AddServer(z, this)
}

func (this *Server) GetZoneId() uint32 {
	return this.zoneId
}

func (this *Server) Process() {
	log.Println("Server connected ", this.conn.RemoteAddr().String())
	defer func() {
		log.Println("Server disconnected", this.zoneId, this.conn.RemoteAddr().String())
		this.Close()
		if err := recover(); err != nil {
			log.Println("Server Process Exception ", err)
		}
	}()

	for {
		// 超过5分钟没有通信，记录日志
		if time.Since(this.lastTime) > 5*time.Minute {
			log.Println("Server last message exceed 5 minute, zoneId:", this.zoneId, this.conn.RemoteAddr().String())
			this.lastTime = time.Now()
		}
		begin := time.Now()
		// 读的超时时间设置为10ms
		this.conn.SetReadDeadline(begin.Add(10 * time.Millisecond))
		this.OnReceive()
		// 读+写的超时时间设置为20ms
		this.conn.SetReadDeadline(begin.Add(20 * time.Millisecond))
		this.OnSend()
		// 每帧20ms的处理时间
		if time.Since(begin) < 20*time.Millisecond {
			time.Sleep(20*time.Millisecond - time.Since(begin))
		}
	}
}

func (this *Server) OnReceive() {
	reader := make([]byte, 16384)
	n, err := this.conn.Read(reader)
	if err != nil {
		// 超时不处理，继续执行
		if nerr, ok := err.(*net.OpError); !ok || !nerr.Timeout() {
			log.Panic("Read Error:", err, this.conn.RemoteAddr().String())
		}
	}
	if n > 0 {
		this.recvBuf = append(this.recvBuf, reader[:n]...)
		//log.Println("Server Receive Buffer", this.recvBuf)
		this.lastTime = time.Now()
	}
	// 每帧最多处理100条协议
	for i := 0; i < 100; i++ {
		if len(this.recvBuf) < 8 {
			break
		}

		oct := common.NewOctets(this.recvBuf)
		size := oct.UnmarshalUint32()
		msgType := oct.UnmarshalUint32()
		if oct.Remain() < size {
			break
		}
		data := oct.UnmarshalBytesOnly(size)

		if msgType != 101 && this.zoneId == 0 {
			log.Panic("Server is not register, Message Type:", msgType)
		}
		msg := message.GetMsg(int(msgType))
		if msg == nil {
			log.Panic("Unknow Protocol Type:", msgType)
		}
		msg = msg.Clone()
		err = msg.Unmarshal(data)
		if err != nil {
			log.Panic("Unmarshal Protocol Error:", err, " Type:", msgType)
		}
		msg.Sets(this)
		common.NewTrans().Process(msg)

		this.recvBuf = this.recvBuf[oct.Pos():]
	}

	// 缓冲区大于1M，记录日志
	if len(this.recvBuf) > 1e6 {
		log.Println("Receive Buffer too Big, size:", len(this.recvBuf))
	}
}

func (this *Server) OnSend() {
	this.sendLock.Lock()
	defer this.sendLock.Unlock()

	if len(this.sendBuf) == 0 {
		return
	}

	n, err := this.conn.Write(this.sendBuf)
	if err != nil {
		if nerr, ok := err.(*net.OpError); !ok || !nerr.Timeout() {
			log.Panic("Write Error:", err, this.conn.RemoteAddr().String())
		}
	}
	this.sendBuf = this.sendBuf[n:]
}

func (this *Server) Send(b []byte) {
	this.sendLock.Lock()
	defer this.sendLock.Unlock()

	this.sendBuf = append(this.sendBuf, b...)
	//log.Println("Server Send Buffer", this.sendBuf)
}

func (this *Server) SendByZoneIds(zoneIds []uint32, b []byte) {
	GetServerMgr().SendByZoneIds(zoneIds, b)
}
