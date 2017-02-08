package manager

import (
	"gameproject/common"
	"gameproject/server/client/csmsg"
	"gameproject/server/transMgr"
	"log"
	"net"
	"sync"
	"time"
)

type Link struct {
	conn     net.Conn
	recvBuf  []byte
	sendBuf  []byte
	sendLock sync.Mutex
	lastTime time.Time
	userId   string
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
	GetLinkMgr().DelLinkByUserId(this.userId)
	GetLinkMgr().DelLinkByRoleId(this.roleId)
}

func (this *Link) SetUserId(u string) {
	if this.userId != "" && this.userId != u {
		log.Println("SetUserId old:", this.userId, " new:", u)
	}
	this.userId = u
	GetLinkMgr().AddLinkByUserId(u, this)
}

func (this *Link) GetUserId() string {
	return this.userId
}

func (this *Link) SetAuthored() {
	this.authored = true
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
			time.Sleep(20*time.Millisecond - time.Since(begin))
		}
	}
}

func (this *Link) OnReceive() {
	reader := make([]byte, 1024)
	n, err := this.conn.Read(reader)
	if err != nil {
		// 超时不处理，继续执行
		if nerr, ok := err.(*net.OpError); !ok || !nerr.Timeout() {
			log.Panic("Read Error:", err, this.conn.RemoteAddr().String())
		}
	}
	if n > 0 {
		this.recvBuf = append(this.recvBuf, reader[:n]...)
		//log.Println("Link Receive Buffer", this.recvBuf)
	}
	// 每帧最多处理3条协议
	for i := 0; i < 3; i++ {
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

		if msgType != 1001 && (this.authored == false || len(this.userId) == 0) {
			log.Panic("Link is not authored, Message Type:", msgType)
		}
		msg := csmsg.GetMsg(int(msgType))
		if msg == nil {
			log.Panic("Unknow Protocol Type:", msgType)
		}
		msg = msg.Clone()
		err = msg.Unmarshal(data)
		if err != nil {
			log.Panic("Unmarshal Protocol Error:", err, " Type:", msgType)
		}
		msg.Setr(this.roleId)
		msg.Setl(this)
		msg.Setg(GetGlobalConn())
		transMgr.NewTrans().Process(msg)

		this.recvBuf = this.recvBuf[oct.Pos():]
	}

	// 接收缓冲区大于100K，断开客户端连接
	if len(this.recvBuf) > 100*1024 {
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
	this.sendBuf = this.sendBuf[n:]
}

func (this *Link) Send(x []byte) {
	this.sendLock.Lock()
	defer this.sendLock.Unlock()

	this.sendBuf = append(this.sendBuf, x...)
	//log.Println("Link Send Buffer", this.sendBuf)

	// 发送缓冲区大于1M，断开客户端连接
	if len(this.sendBuf) > 1024*1024 {
		log.Panic("Send Buffer too Big!")
	}
}
