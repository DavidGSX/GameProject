package manager

import (
	"gameproject/common"
	"gameproject/server/protocol"
	"log"
	"net"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
)

type Link struct {
	conn     net.Conn
	recvBuf  []byte
	sendBuf  []byte
	sendLock sync.Mutex
	lastTime time.Time
	roleId   uint64
	authored bool
}

func NewLink(c net.Conn) *Link {
	l := new(Link)
	l.conn = c
	l.recvBuf = make([]byte, 0)
	l.sendBuf = make([]byte, 0)
	l.lastTime = time.Now()
	return l
}

func (this *Link) Close() {
	this.conn.Close()
	GetLinkMgr().DelLink(this.roleId)
}

func (this *Link) Process() {
	log.Println("Link connected ", this.conn.RemoteAddr().String())
	defer func() {
		log.Println("Link disconnected", this.roleId, this.conn.RemoteAddr().String())
		this.Close()
		if err := recover(); err != nil {
			log.Println("Link Process Exception ", err)
		}
	}()

	for {
		// 超过5分钟没有通信，断开连接
		if time.Since(this.lastTime) > 5*time.Minute {
			log.Println("Link disconnected keep alive exceed 5 minute, roleID:", this.roleId, this.conn.RemoteAddr().String())
			this.Close()
			return
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
			<-time.After(20*time.Millisecond - time.Since(begin))
		}
	}
}

func (this *Link) OnReceive() {
	reader := make([]byte, 1024)
	n, err := this.conn.Read(reader)
	if err != nil {
		if nerr, ok := err.(*net.OpError); !ok || !nerr.Timeout() {
			log.Panic("Read Error:", err, this.conn.RemoteAddr().String())
		}
	}
	this.recvBuf = append(this.recvBuf, reader[:n]...)

	// 每帧最多处理2条协议
	for i := 0; i < 2; i++ {
		if len(this.recvBuf) < 4 {
			break
		}

		oct := common.NewOctets(this.recvBuf)
		size := oct.UncompactUint32()
		if oct.Remain() < int(size) {
			break
		}

		data := oct.UnmarshalBytes()
		addmoney := &protocol.CAddMoney{}
		err = proto.Unmarshal(data, addmoney)
		if err != nil {
			log.Panic("unmarshal CAddMoney error:", err)
		}

		if addmoney.GetNum()%1000 == 0 {
			log.Println("RoleId:", addmoney.GetRoleId(), " Num:", addmoney.GetNum())
		}

		this.recvBuf = this.recvBuf[oct.Pos():]
	}

	// 缓冲区大于1M，断开客户端连接
	if len(this.recvBuf) > 1e6 {
		log.Panic("Receive Buffer too Big!")
	}
}

func (this *Link) OnSend() {
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
	//log.Println(n, this.sendbuf[:n])
	this.sendBuf = this.sendBuf[n:]
}

func (this *Link) Send(x []byte) {
	this.sendLock.Lock()
	defer this.sendLock.Unlock()

	this.sendBuf = append(this.sendBuf, x...)
}
