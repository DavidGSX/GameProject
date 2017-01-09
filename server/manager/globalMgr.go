package manager

import (
	"gameproject/common"
	"gameproject/global/protocol"
	"gameproject/server/config"
	"gameproject/server/msgProto"
	"log"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
)

var globalConn *GlobalConn
var globalConnLock sync.RWMutex

type GlobalConn struct {
	cfg     *config.ServerConfig
	conn    net.Conn
	recvBuf []byte
	sendBuf []byte
}

func GetGlobalConn() *GlobalConn {
	globalConnLock.Lock()
	defer globalConnLock.Unlock()

	return globalConn
}

func ReConnectGlobal() {
	<-time.After(1e10) // 10秒钟后重连Global
	if globalConn != nil && globalConn.cfg != nil {
		GlobalMgrInit(globalConn.cfg)
	} else {
		log.Println("ReConnect Global Failed, cfg is nil")
	}
}

func GlobalMgrInit(cfg *config.ServerConfig) {
	ip := cfg.GlobalConfig.GlobalIP
	port := cfg.GlobalConfig.GlobalPort
	conn, err := net.Dial("tcp", ip+":"+strconv.Itoa(int(port)))
	if err != nil {
		log.Fatal("Connect Global Error:", err)
	}
	log.Println("Connect Global Success ", ip, port)

	globalConnLock.Lock()
	defer globalConnLock.Unlock()

	if globalConn == nil {
		globalConn = new(GlobalConn)
	}

	globalConn.cfg = cfg
	globalConn.conn = conn
	globalConn.recvBuf = make([]byte, 0)
	globalConn.sendBuf = make([]byte, 0)
	globalConn.OnRegisterGlobal(cfg)
	go globalConn.Process()
}

func (this *GlobalConn) OnRegisterGlobal(cfg *config.ServerConfig) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("OnRegisterGlobal Error:", err)
		}
	}()

	send := &protocol.SGServerStart{}
	send.ZoneId = cfg.GetZoneId()
	send.Plat = cfg.GetPlatform()
	data, err := proto.Marshal(send)
	if err != nil {
		log.Println("Marshal SGServerStart error:", err)
		return
	}

	oct := &common.Octets{}
	oct.MarshalUint32(uint32(len(data)))
	oct.MarshalUint32(1)
	oct.MarshalBytesOnly(data)
	this.sendBuf = append(this.sendBuf, oct.GetBuf()...)
}

func (this *GlobalConn) Process() {
	defer func() {
		if this.conn != nil {
			this.conn.Close()
		}
		go ReConnectGlobal()
		if err := recover(); err != nil {
			log.Println("Global Exception -> ", err)
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

func (this *GlobalConn) OnReceive() {
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

		switch msgType {
		case 3:
			this.OnAuthResult(data)
		default:
			// 如果协议错误，就清空接受缓冲区
			log.Println("Invalid Msg Type:", msgType)
			this.recvBuf = make([]byte, 0)
			return
		}

		this.recvBuf = this.recvBuf[oct.Pos():]
	}

	// 缓冲区大于10M，记录日志
	if len(this.recvBuf) > 1e7 {
		log.Println("Receive Buffer too Big, size:", len(this.recvBuf))
	}
}

func (this *GlobalConn) OnAuthResult(b []byte) {
	res := &protocol.GSAuthResult{}
	err := proto.Unmarshal(b, res)
	if err != nil {
		log.Println("unmarshal GSAuthResult error:", err)
		return
	}

	userId := res.UserId
	plat := res.Plat
	result := res.Result
	log.Println("OnAuthResult userId:", userId, " plat:", plat, " result:", result)
	link := GetLinkMgr().GetLinkByUserId(userId)

	loginRes := &msgProto.SUserLogin{}
	if result == "ok" {
		loginRes.LoginRes = msgProto.SUserLogin_SUCCESS
		if link != nil {
			link.SetAuthored()
		}
	} else {
		loginRes.LoginRes = msgProto.SUserLogin_PASSWD_ERR
	}
	data, err := proto.Marshal(loginRes)
	if err != nil {
		log.Panic("marshal CUserLogin error:", err)
	}
	oct := &common.Octets{}
	oct.MarshalUint32(uint32(len(data)))
	oct.MarshalUint32(1002)
	oct.MarshalBytesOnly(data)
	if link != nil {
		link.Send(oct.GetBuf())
	} else {
		log.Println("OnAuthResult userId:", userId, " link is nil")
	}
}

func (this *GlobalConn) OnSend() {
	globalConnLock.Lock()
	defer globalConnLock.Unlock()

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

func (this *GlobalConn) Send(x []byte) {
	globalConnLock.Lock()
	defer globalConnLock.Unlock()

	this.sendBuf = append(this.sendBuf, x...)
}

func (this *GlobalConn) SetUserId(u string) {
	// just use to adjust interface ISend
}

func (this *GlobalConn) GetUserId() string {
	// just use to adjust interface ISend
	return ""
}
