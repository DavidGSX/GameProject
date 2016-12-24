package common

import (
	"bytes"
	"math"
	"strings"
	"testing"
)

func TestNewOctets(t *testing.T) {
	p1 := NewOctets([]byte("just a test"))
	p2 := NewOctets([]byte("just a test"))
	p3 := NewOctets([]byte("just a test!"))
	if p1.Equals(p2) == false {
		t.Errorf("p1 != p2")
	}
	if p1.Equals(p3) {
		t.Errorf("p1 == p3")
	}
	if p2.Equals(p3) {
		t.Errorf("p2 == p3")
	}
}

func TestEquals(t *testing.T) {
	p1 := NewOctets([]byte("just a test"))
	p2 := NewOctets([]byte("just a test"))
	p3 := NewOctets([]byte("just a test!"))
	if p1.Equals(p2) == false {
		t.Errorf("p1 != p2")
	}
	if p1.Equals(p3) {
		t.Errorf("p1 == p3")
	}
	if p2.Equals(p3) {
		t.Errorf("p2 == p3")
	}
}

func TestGetBuf(t *testing.T) {
	p := NewOctets([]byte("just a test"))
	b := p.GetBuf()
	if bytes.Equal(b, []byte("just a test")) == false {
		t.Errorf("failed!")
	}
	if bytes.Equal(b, []byte("just a test!")) {
		t.Errorf("failed!")
	}
}

func TestMarshalByte(t *testing.T) {
	p := NewOctets([]byte(""))
	p.MarshalByte(byte('a'))
	p.MarshalByte(byte('b'))
	p.MarshalByte(byte('c'))
	if bytes.Equal(p.GetBuf(), []byte("abc")) == false {
		t.Errorf("failed!")
	}
	if bytes.Equal(p.GetBuf(), []byte("aaa")) {
		t.Errorf("failed!")
	}
}

func TestUnmarshalByte(t *testing.T) {
	p := NewOctets([]byte("abc"))
	if p.UnmarshalByte() != byte('a') {
		t.Errorf("failed!")
	}
	if p.Pos() != 1 {
		t.Errorf("failed!")
	}
	if p.UnmarshalByte() != byte('b') {
		t.Errorf("failed!")
	}
	if p.Pos() != 2 {
		t.Errorf("failed!")
	}
	if p.UnmarshalByte() != byte('c') {
		t.Errorf("failed!")
	}
	if p.Pos() != 3 {
		t.Errorf("failed!")
	}
}

func TestMarshalBool(t *testing.T) {
	p := NewOctets([]byte(""))
	p.MarshalBool(true)
	p.MarshalBool(true)
	p.MarshalBool(false)
	if bytes.Equal(p.GetBuf(), []byte{1, 1, 0}) == false {
		t.Errorf("failed!")
	}
	if bytes.Equal(p.GetBuf(), []byte("110")) {
		t.Errorf("failed!")
	}
}

func TestUnmarshalBool(t *testing.T) {
	p := NewOctets([]byte{1, 1, 0})
	if p.UnmarshalBool() != true {
		t.Errorf("failed!")
	}
	if p.Pos() != 1 {
		t.Errorf("failed!")
	}
	if p.UnmarshalBool() != true {
		t.Errorf("failed!")
	}
	if p.Pos() != 2 {
		t.Errorf("failed!")
	}
	if p.UnmarshalBool() {
		t.Errorf("failed!")
	}
	if p.Pos() != 3 {
		t.Errorf("failed!")
	}
}

func TestMarshalUint16(t *testing.T) {
	p := NewOctets([]byte(""))
	p.MarshalUint16(0x1234)
	p.MarshalUint16(0x5678)
	p.MarshalUint16(0x9abc)
	p.MarshalUint16(0xdef0)
	if bytes.Equal(p.GetBuf(), []byte{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0}) == false {
		t.Errorf("failed!")
	}
	if bytes.Equal(p.GetBuf(), []byte("123456789abcdef0")) {
		t.Errorf("failed!")
	}
}

func TestUnmarshalUint16(t *testing.T) {
	p := NewOctets([]byte{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0})
	if p.UnmarshalUint16() != 0x1234 {
		t.Errorf("failed!")
	}
	if p.Pos() != 2 {
		t.Errorf("failed!")
	}
	if p.UnmarshalUint16() != 0x5678 {
		t.Errorf("failed!")
	}
	if p.Pos() != 4 {
		t.Errorf("failed!")
	}
	if p.UnmarshalUint16() != 0x9abc {
		t.Errorf("failed!")
	}
	if p.Pos() != 6 {
		t.Errorf("failed!")
	}
	if p.UnmarshalUint16() != 0xdef0 {
		t.Errorf("failed!")
	}
	if p.Pos() != 8 {
		t.Errorf("failed!")
	}
}

func TestMarshalUint32(t *testing.T) {
	p := NewOctets([]byte(""))
	p.MarshalUint32(0x12345678)
	p.MarshalUint32(0x9abcdef0)
	if bytes.Equal(p.GetBuf(), []byte{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0}) == false {
		t.Errorf("failed!")
	}
	if bytes.Equal(p.GetBuf(), []byte("123456789abcdef0")) {
		t.Errorf("failed!")
	}
}

func TestUnmarshalUint32(t *testing.T) {
	p := NewOctets([]byte{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0})
	if p.UnmarshalUint32() != 0x12345678 {
		t.Errorf("failed!")
	}
	if p.Pos() != 4 {
		t.Errorf("failed!")
	}
	if p.UnmarshalUint32() != 0x9abcdef0 {
		t.Errorf("failed!")
	}
	if p.Pos() != 8 {
		t.Errorf("failed!")
	}
}

func TestMarshalUint64(t *testing.T) {
	p := NewOctets([]byte(""))
	p.MarshalUint64(0x123456789abcdef0)
	if bytes.Equal(p.GetBuf(), []byte{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0}) == false {
		t.Errorf("failed!")
	}
	if bytes.Equal(p.GetBuf(), []byte("123456789abcdef0")) {
		t.Errorf("failed!")
	}
}

func TestUnmarshalUint64(t *testing.T) {
	p := NewOctets([]byte{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0})
	if p.UnmarshalUint64() != uint64(0x123456789abcdef0) {
		t.Errorf("failed!")
	}
	if p.Pos() != 8 {
		t.Errorf("failed!")
	}
}

func TestMarshalFloat32(t *testing.T) {
	p := NewOctets([]byte(""))
	p.MarshalFloat32(1.23456)
	if bytes.Equal(p.GetBuf(), []byte{0x3f, 0x9e, 0x06, 0x10}) == false {
		t.Errorf("failed!")
	}
	if bytes.Equal(p.GetBuf(), []byte("1.23456")) {
		t.Errorf("failed!")
	}
}

func TestUnmarshalFloat32(t *testing.T) {
	p := NewOctets([]byte{0x3f, 0x9e, 0x06, 0x10})
	if math.Abs(float64(p.UnmarshalFloat32()-1.23456)) > 1e-8 {
		t.Errorf("failed!")
	}
	if p.Pos() != 4 {
		t.Errorf("failed!")
	}
}

func TestMarshalFloat64(t *testing.T) {
	p := NewOctets([]byte(""))
	p.MarshalFloat64(1.234567890123456)
	if bytes.Equal(p.GetBuf(), []byte{0x3f, 0xf3, 0xc0, 0xca, 0x42, 0x8c, 0x59, 0xf8}) == false {
		t.Errorf("failed!")
	}
	if bytes.Equal(p.GetBuf(), []byte("1.234567890123456")) {
		t.Errorf("failed!")
	}
}

func TestUnmarshalFloat64(t *testing.T) {
	p := NewOctets([]byte{0x3f, 0xf3, 0xc0, 0xca, 0x42, 0x8c, 0x59, 0xf8})
	if math.Abs(p.UnmarshalFloat64()-1.234567890123456) > 1e-18 {
		t.Errorf("failed!")
	}
	if p.Pos() != 8 {
		t.Errorf("failed!")
	}
}

func TestMarshalBytes(t *testing.T) {
	p := NewOctets([]byte(""))
	p.MarshalBytes([]byte{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0})
	if bytes.Equal(p.GetBuf(), []byte{0x08, 0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0}) == false {
		t.Errorf("failed!")
	}
	if bytes.Equal(p.GetBuf(), []byte("123456789abcdef0")) {
		t.Errorf("failed!")
	}
}

func TestUnmarshalBytes(t *testing.T) {
	p := NewOctets([]byte{0x08, 0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0})
	b := p.UnmarshalBytes()
	if bytes.Equal(b, []byte{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0}) == false {
		t.Errorf("failed!")
	}
	if p.Pos() != 9 {
		t.Errorf("failed!")
	}
}

func TestMarshalUint16s(t *testing.T) {
	p := NewOctets([]byte(""))
	p.MarshalUint16s([]uint16{0x1234, 0x5678, 0x9abc, 0xdef0})
	if bytes.Equal(p.GetBuf(), []byte{0x08, 0x34, 0x12, 0x78, 0x56, 0xbc, 0x9a, 0xf0, 0xde}) == false {
		t.Errorf("failed!")
	}
	if bytes.Equal(p.GetBuf(), []byte("123456789abcdef0")) {
		t.Errorf("failed!")
	}
}

func TestUnmarshalUint16s(t *testing.T) {
	p := NewOctets([]byte{0x08, 0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0})
	b := p.UnmarshalUint16s()
	if b[0] != 0x3412 {
		t.Errorf("failed!")
	}
	if b[1] != 0x7856 {
		t.Errorf("failed!")
	}
	if b[2] != 0xbc9a {
		t.Errorf("failed!")
	}
	if b[3] != 0xf0de {
		t.Errorf("failed!")
	}
	if p.Pos() != 9 {
		t.Errorf("failed!")
	}
}

func TestMarshalString(t *testing.T) {
	p := NewOctets([]byte(""))
	p.MarshalString("123456789")
	if bytes.Equal(p.GetBuf(), []byte{0x09, '1', '2', '3', '4', '5', '6', '7', '8', '9'}) == false {
		t.Errorf("failed!")
	}
	if bytes.Equal(p.GetBuf(), []byte("123456789")) {
		t.Errorf("failed!")
	}
}

func TestUnmarshalString(t *testing.T) {
	p := NewOctets([]byte{0x09, '1', '2', '3', '4', '5', '6', '7', '8', '9'})
	if strings.Compare(p.UnmarshalString(), "123456789") != 0 {
		t.Errorf("failed!")
	}
	if p.Pos() != 10 {
		t.Errorf("failed!")
	}
}

func TestCompactUint32(t *testing.T) {
	p := NewOctets([]byte(""))
	p.CompactUint32(0x01)
	p.CompactUint32(0x3f)
	p.CompactUint32(0x40)
	p.CompactUint32(0x3fff)
	p.CompactUint32(0x4000)
	p.CompactUint32(0x3fffffff)
	p.CompactUint32(0x40000000)
	p.CompactUint32(0xffffffff)
	b := p.GetBuf()
	if bytes.Equal(b, []byte{0x01, 0x3f, 0x40, 0x40, 0x7f, 0xff, 0x80, 0x00, 0x40, 0x00, 0xbf, 0xff, 0xff, 0xff, 0xc0, 0x40, 0x00, 0x00, 0x00, 0xc0, 0xff, 0xff, 0xff, 0xff}) == false {
		t.Errorf("failed!")
	}
}

func TestUncompactUint32(t *testing.T) {
	p := NewOctets([]byte{0x01, 0x3f})
	if p.UncompactUint32() != 0x01 {
		t.Errorf("failed!")
	}
	if p.Pos() != 1 {
		t.Errorf("failed!")
	}
	if p.UncompactUint32() != 0x3f {
		t.Errorf("failed!")
	}
	if p.Pos() != 2 {
		t.Errorf("failed!")
	}

	p = NewOctets([]byte{0x40, 0x40, 0x7f, 0xff})
	if p.UncompactUint32() != 0x40 {
		t.Errorf("failed!")
	}
	if p.Pos() != 2 {
		t.Errorf("failed!")
	}
	if p.UncompactUint32() != 0x3fff {
		t.Errorf("failed!")
	}
	if p.Pos() != 4 {
		t.Errorf("failed!")
	}

	p = NewOctets([]byte{0x80, 0x00, 0x40, 0x00, 0xbf, 0xff, 0xff, 0xff})
	if p.UncompactUint32() != 0x4000 {
		t.Errorf("failed!")
	}
	if p.Pos() != 4 {
		t.Errorf("failed!")
	}
	if p.UncompactUint32() != 0x3fffffff {
		t.Errorf("failed!")
	}
	if p.Pos() != 8 {
		t.Errorf("failed!")
	}

	p = NewOctets([]byte{0xc0, 0x40, 0x00, 0x00, 0x00, 0xc0, 0xff, 0xff, 0xff, 0xff})
	if p.UncompactUint32() != 0x40000000 {
		t.Errorf("failed!")
	}
	if p.Pos() != 5 {
		t.Errorf("failed!")
	}
	if p.UncompactUint32() != 0xffffffff {
		t.Errorf("failed!")
	}
	if p.Pos() != 10 {
		t.Errorf("failed!")
	}
}
