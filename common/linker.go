package common

import (
	"log"
	"net"
)

type Linker struct {
	conn    net.Conn
	sid     uint32
	recvbuf []byte
	sendbuf []byte
}

func NewLinker(c net.Conn) *Linker {
	l := new(Linker)
	l.conn = c
	l.recvbuf = make([]byte, 0)
	l.sendbuf = make([]byte, 0)
	return l
}

func (l *Linker) SetSid(s uint32) {
	l.sid = s
}

func (l *Linker) GetSid() uint32 {
	return l.sid
}

func (l *Linker) Receive(x []byte) {
	l.recvbuf = append(l.recvbuf, x...)

	for {
		if len(l.recvbuf) == 0 {
			break
		}

		oct := NewOctets(l.recvbuf)
		t := oct.UncompactUint32()
		size := oct.UncompactUint32()
		if size > uint32(oct.Remain()) {
			break
		}

		mark := oct.Pos() + int(size)
		l.recvbuf = l.recvbuf[mark:]

		sid := oct.UnmarshalUint32()
		l.SetSid(sid & 0x7fffffff)

		GetProtoMgr().Process(int(t), oct, l)

		if oct.Pos() != mark {
			log.Panic("ProtoMgr.Process type", t, "mark", mark, "error pos", oct.Pos())
		}
	}
}

func (l *Linker) Send(x []byte) {
	l.sendbuf = append(l.sendbuf, x...)
	n, err := l.conn.Write(l.sendbuf)
	if err != nil {
		log.Panic("Write Error:", err)
	}
	log.Println(n, l.sendbuf[:n])
	l.sendbuf = l.sendbuf[n:]
}
