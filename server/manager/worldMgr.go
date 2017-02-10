package manager

import (
	"gameproject/common"
	"gameproject/server/config"
	"gameproject/server/world/swmsg"
	"gameproject/world/msgProto"
	"log"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
)

var worldConn *WorldConn
var worldConnLock sync.RWMutex

type WorldConn struct {
	cfg     *config.ServerConfig
	conn    net.Conn
	recvBuf []byte
	sendBuf []byte
}

func GetWorldConn() *WorldConn {
	worldConnLock.Lock()
	defer worldConnLock.Unlock()

	return worldConn
}

func ReConnectWorld() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("ReConnectWorld Error:", err)
		}
	}()
	<-time.After(1e10) // 10秒钟后重连World
	if worldConn != nil && worldConn.cfg != nil {
		WorldMgrInit(worldConn.cfg)
	} else {
		log.Println("ReConnect World Failed, cfg is nil")
	}
}

func WorldMgrInit(cfg *config.ServerConfig) {
	worldConnLock.Lock()
	defer worldConnLock.Unlock()

	if worldConn == nil {
		worldConn = new(WorldConn)
	}
	worldConn.cfg = cfg

	ip := cfg.WorldConfig.WorldIP
	port := cfg.WorldConfig.WorldPort
	conn, err := net.Dial("tcp", ip+":"+strconv.Itoa(int(port)))
	if err != nil {
		log.Println("Connect World Error:", err)
		go ReConnectWorld()
		return
	}
	log.Println("Connect World Success ", ip, port)

	worldConn.conn = conn
	worldConn.recvBuf = make([]byte, 0)
	worldConn.sendBuf = make([]byte, 0)
	worldConn.OnRegisterWorld(cfg)
	go worldConn.Process()
}

func (this *WorldConn) OnRegisterWorld(cfg *config.ServerConfig) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("OnRegisterWorld Error:", err)
		}
	}()

	send := &msgProto.S2WServerStart{}
	send.ZoneId = cfg.GetZoneId()
	data, err := proto.Marshal(send)
	if err != nil {
		log.Println("Marshal GS2WServerStart error:", err)
		return
	}

	oct := &common.Octets{}
	oct.MarshalUint32(uint32(len(data)))
	oct.MarshalUint32(101)
	oct.MarshalBytesOnly(data)
	this.sendBuf = append(this.sendBuf, oct.GetBuf()...)
}

func (this *WorldConn) Process() {
	defer func() {
		if this.conn != nil {
			this.conn.Close()
		}
		go ReConnectWorld()
		if err := recover(); err != nil {
			log.Println("World Exception -> ", err)
		}
	}()
	for {
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

func (this *WorldConn) OnReceive() {
	reader := make([]byte, 16384)
	n, err := this.conn.Read(reader)
	if err != nil {
		// 超时不处理，继续执行
		if nerr, ok := err.(*net.OpError); !ok || !nerr.Timeout() {
			log.Panic("Read Error:", err, " ", this.conn.RemoteAddr().String())
		}
	}
	if n > 0 {
		this.recvBuf = append(this.recvBuf, reader[:n]...)
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

		msg := swmsg.GetMsg(int(msgType))
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

	// 缓冲区大于10M，记录日志
	if len(this.recvBuf) > 1e7 {
		log.Println("Receive Buffer too Big, size:", len(this.recvBuf))
	}
}

func (this *WorldConn) OnRegisted(b []byte) {
	res := &msgProto.W2SServerStartRes{}
	err := proto.Unmarshal(b, res)
	if err != nil {
		log.Println("unmarshal W2GSServerStartRes error:", err)
		return
	}
}

func (this *WorldConn) OnSend() {
	worldConnLock.Lock()
	defer worldConnLock.Unlock()

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

func (this *WorldConn) Send(x []byte) {
	worldConnLock.Lock()
	defer worldConnLock.Unlock()

	this.sendBuf = append(this.sendBuf, x...)
}

func (this *WorldConn) SendByRoleIds(roleIds []uint64, b []byte) {
	GetLinkMgr().SendByRoleIds(roleIds, b)
}

func (this *WorldConn) SetUserId(u string) {
	// just use to adjust interface ISend
}

func (this *WorldConn) GetUserId() string {
	// just use to adjust interface ISend
	return ""
}

func (this *WorldConn) SetRoleId(r uint64) {
	// just use to adjust interface ISend
}

func (this *WorldConn) GetRoleId() uint64 {
	// just use to adjust interface ISend
	return 0
}

func (this *WorldConn) SetZoneId(r uint32) {
	// just use to adjust interface ISend
}

func (this *WorldConn) GetZoneId() uint32 {
	// just use to adjust interface ISend
	return 0
}

func (this *WorldConn) SendByZoneIds(zoneIds []uint32, b []byte) {
	// just use to adjust interface ISend
}
