package common

import (
	"log"
)

type Octets struct {
	buf     []byte
	pos     int
	tranpos int
}

func NewOctets(x []byte) *Octets {
	o := new(Octets)
	o.buf = x
	return o
}

func (o *Octets) GetBuf() []byte {
	return o.buf
}

func (o *Octets) Remain() int {
	return len(o.buf) - o.pos
}

func (o *Octets) Pos() int {
	return o.pos
}

func (o *Octets) MarshalByte(x byte) {
	o.buf = append(o.buf, x)
}

func (o *Octets) UnmarshalByte() byte {
	if o.pos+1 > len(o.buf) {
		log.Panic("UnmarshalByte Error pos:", o.pos, "buf len:", len(o.buf))
	}
	v := o.buf[o.pos]
	o.pos++
	return v
}

func (o *Octets) MarshalBool(x bool) {
	if x {
		o.MarshalByte(byte(1))
	} else {
		o.MarshalByte(byte(0))
	}
}

func (o *Octets) UnmarshalBool() bool {
	if o.pos+1 > len(o.buf) {
		log.Panic("UnmarshalBool Error pos:", o.pos, "buf len:", len(o.buf))
	}
	v := o.buf[o.pos]
	o.pos++
	return v == 1
}

func (o *Octets) MarshalUint16(x uint16) {
	o.MarshalByte(byte(x >> 8))
	o.MarshalByte(byte(x))
}

func (o *Octets) UnmarshalUint16() uint16 {
	if o.pos+2 > len(o.buf) {
		log.Panic("UnmarshalUint16 Error pos:", o.pos, "buf len:", len(o.buf))
	}
	v0 := uint16(o.buf[o.pos])
	v1 := uint16(o.buf[o.pos+1])
	o.pos += 2
	return (v0 << 8) | v1
}

func (o *Octets) MarshalUint32(x uint32) {
	o.MarshalByte(byte(x >> 24))
	o.MarshalByte(byte(x >> 16))
	o.MarshalByte(byte(x >> 8))
	o.MarshalByte(byte(x))
}

func (o *Octets) UnmarshalUint32() uint32 {
	if o.pos+4 > len(o.buf) {
		log.Panic("UnmarshalUint32 Error pos:", o.pos, "buf len:", len(o.buf))
	}
	v0 := uint32(o.buf[o.pos])
	v1 := uint32(o.buf[o.pos+1])
	v2 := uint32(o.buf[o.pos+2])
	v3 := uint32(o.buf[o.pos+3])
	o.pos += 4
	return (v0 << 24) | (v1 << 16) | (v2 << 8) | v3
}

func (o *Octets) MarshalUint64(x uint64) {
	o.MarshalByte(byte(x >> 56))
	o.MarshalByte(byte(x >> 48))
	o.MarshalByte(byte(x >> 40))
	o.MarshalByte(byte(x >> 32))
	o.MarshalByte(byte(x >> 24))
	o.MarshalByte(byte(x >> 16))
	o.MarshalByte(byte(x >> 8))
	o.MarshalByte(byte(x))
}

func (o *Octets) UnmarshalUint64() uint64 {
	if o.pos+8 > len(o.buf) {
		log.Panic("UnmarshalUint64 Error pos:", o.pos, "buf len:", len(o.buf))
	}
	v0 := uint64(o.buf[o.pos])
	v1 := uint64(o.buf[o.pos+1])
	v2 := uint64(o.buf[o.pos+2])
	v3 := uint64(o.buf[o.pos+3])
	v4 := uint64(o.buf[o.pos])
	v5 := uint64(o.buf[o.pos+1])
	v6 := uint64(o.buf[o.pos+2])
	v7 := uint64(o.buf[o.pos+3])
	o.pos += 8
	return (v0 << 56) | (v1 << 48) | (v2 << 40) | (v3 << 32) | (v4 << 24) | (v5 << 16) | (v6 << 8) | v7
}

func (o *Octets) MarshalFloat32(x float32) {
}

func (o *Octets) MarshalFloat64(x float64) {
}

func (o *Octets) MarshalBytes(x []byte) {
	o.CompactUint32(uint32(len(x)))
	o.buf = append(o.buf, x...)
}

func (o *Octets) UnmarshalBytes() []byte {
	size := int(o.UncompactUint32())
	if o.pos+size > len(o.buf) {
		log.Panic("UnmarshalBytes Error pos:", o.pos, "size:", size, "buf len:", len(o.buf))
	}
	v := make([]byte, size)
	copy(v, o.buf[o.pos:])
	o.pos += size
	return v
}

func (o *Octets) MarshalUint16s(x []uint16) {
	o.CompactUint32(uint32(len(x)))
	for i := 0; i < len(x); i++ {
		low := byte(x[i] & 0xff)
		high := byte(x[i] >> 8)
		o.buf = append(o.buf, low, high)
	}
}

func (o *Octets) UnmarshalUint16s() []uint16 {
	size := int(o.UncompactUint32())
	if o.pos+size > len(o.buf) {
		log.Panic("UnmarshalUint16s Error pos:", o.pos, "size:", size, "buf len:", len(o.buf))
	}
	if size%2 != 0 {
		log.Panic("UnmarshalUint16s Error size:", size)
	}
	v := make([]uint16, size/2)
	for i, j := o.pos, 0; j < size/2; i, j = i+2, j+1 {
		low := uint16(o.buf[i])
		high := uint16(o.buf[i+1])
		v[j] = low | (high << 8)
	}
	o.pos += size
	return v
}

func (o *Octets) MarshalString(x string) {
	o.MarshalBytes([]byte(x))
}

func (o *Octets) UnmarshalString() string {
	return string(o.UnmarshalBytes())
}

/*
 *  x < 0x40		y=x				y&0xc0 == 0
 *  x < 0x4000		y=x|0x4000		y&0xc0 == 0x40
 *  x < 0x40000000	y=x|0x80000000	y&0xc0 == 0x80
 *  x >=0x40000000	y=0xc0, x		y&0xc0 == 0xc0
 */
func (o *Octets) CompactUint32(x uint32) {
	if x < 0x40 {
		o.MarshalByte(byte(x))
	} else if x < 0x4000 {
		o.MarshalUint16(uint16(x) | 0x4000)
	} else if x < 0x40000000 {
		o.MarshalUint32(uint32(x) | 0x80000000)
	} else {
		o.MarshalByte(byte(0xc0))
		o.MarshalUint32(x)
	}
}

func (o *Octets) UncompactUint32() uint32 {
	if o.pos == len(o.buf) {
		log.Panic("UncompactUint32 Error pos:", o.pos, "buf len:", len(o.buf))
	}

	v := o.buf[o.pos]
	switch v & 0xc0 {
	case 0xc0:
		o.UnmarshalByte()
		return o.UnmarshalUint32()
	case 0x80:
		return o.UnmarshalUint32() & (^uint32(0x80000000))
	case 0x40:
		return uint32(o.UnmarshalUint16() & (^uint16(0x4000)))
	case 0:
		return uint32(o.UnmarshalByte())
	default:
		log.Panic("UncompactUint32 Error v:", byte(v))
	}
	return 0
}
