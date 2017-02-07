// Code generated by protoc-gen-go.
// source: world.proto
// DO NOT EDIT!

/*
Package msgProto is a generated protocol buffer package.

It is generated from these files:
	world.proto

It has these top-level messages:
	GS2WServerStart
	W2GSServerStartRes
	GS2WSendInfo
	W2GSSendInfo
	W2GSDispatch
	W2GSBroadcast
*/
package msgProto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GS2WServerStart struct {
	ZoneId uint32 `protobuf:"varint,1,opt,name=zoneId" json:"zoneId,omitempty"`
}

func (m *GS2WServerStart) Reset()                    { *m = GS2WServerStart{} }
func (m *GS2WServerStart) String() string            { return proto.CompactTextString(m) }
func (*GS2WServerStart) ProtoMessage()               {}
func (*GS2WServerStart) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GS2WServerStart) GetZoneId() uint32 {
	if m != nil {
		return m.ZoneId
	}
	return 0
}

type W2GSServerStartRes struct {
}

func (m *W2GSServerStartRes) Reset()                    { *m = W2GSServerStartRes{} }
func (m *W2GSServerStartRes) String() string            { return proto.CompactTextString(m) }
func (*W2GSServerStartRes) ProtoMessage()               {}
func (*W2GSServerStartRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type GS2WSendInfo struct {
	ZoneId uint32 `protobuf:"varint,1,opt,name=zoneId" json:"zoneId,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=userId" json:"userId,omitempty"`
	Type   uint32 `protobuf:"varint,3,opt,name=type" json:"type,omitempty"`
	Info   string `protobuf:"bytes,4,opt,name=info" json:"info,omitempty"`
}

func (m *GS2WSendInfo) Reset()                    { *m = GS2WSendInfo{} }
func (m *GS2WSendInfo) String() string            { return proto.CompactTextString(m) }
func (*GS2WSendInfo) ProtoMessage()               {}
func (*GS2WSendInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GS2WSendInfo) GetZoneId() uint32 {
	if m != nil {
		return m.ZoneId
	}
	return 0
}

func (m *GS2WSendInfo) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *GS2WSendInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *GS2WSendInfo) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

type W2GSSendInfo struct {
	ZoneId uint32 `protobuf:"varint,1,opt,name=zoneId" json:"zoneId,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=userId" json:"userId,omitempty"`
	Type   uint32 `protobuf:"varint,3,opt,name=type" json:"type,omitempty"`
	Info   string `protobuf:"bytes,4,opt,name=info" json:"info,omitempty"`
}

func (m *W2GSSendInfo) Reset()                    { *m = W2GSSendInfo{} }
func (m *W2GSSendInfo) String() string            { return proto.CompactTextString(m) }
func (*W2GSSendInfo) ProtoMessage()               {}
func (*W2GSSendInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *W2GSSendInfo) GetZoneId() uint32 {
	if m != nil {
		return m.ZoneId
	}
	return 0
}

func (m *W2GSSendInfo) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *W2GSSendInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *W2GSSendInfo) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

type W2GSDispatch struct {
	FromZoneId uint32 `protobuf:"varint,1,opt,name=fromZoneId" json:"fromZoneId,omitempty"`
	ToZoneId   uint32 `protobuf:"varint,2,opt,name=toZoneId" json:"toZoneId,omitempty"`
	Type       uint32 `protobuf:"varint,3,opt,name=type" json:"type,omitempty"`
	Info       string `protobuf:"bytes,4,opt,name=info" json:"info,omitempty"`
}

func (m *W2GSDispatch) Reset()                    { *m = W2GSDispatch{} }
func (m *W2GSDispatch) String() string            { return proto.CompactTextString(m) }
func (*W2GSDispatch) ProtoMessage()               {}
func (*W2GSDispatch) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *W2GSDispatch) GetFromZoneId() uint32 {
	if m != nil {
		return m.FromZoneId
	}
	return 0
}

func (m *W2GSDispatch) GetToZoneId() uint32 {
	if m != nil {
		return m.ToZoneId
	}
	return 0
}

func (m *W2GSDispatch) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *W2GSDispatch) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

type W2GSBroadcast struct {
	ZoneIds []uint32 `protobuf:"varint,1,rep,packed,name=zoneIds" json:"zoneIds,omitempty"`
	Type    uint32   `protobuf:"varint,2,opt,name=type" json:"type,omitempty"`
	Info    string   `protobuf:"bytes,3,opt,name=info" json:"info,omitempty"`
}

func (m *W2GSBroadcast) Reset()                    { *m = W2GSBroadcast{} }
func (m *W2GSBroadcast) String() string            { return proto.CompactTextString(m) }
func (*W2GSBroadcast) ProtoMessage()               {}
func (*W2GSBroadcast) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *W2GSBroadcast) GetZoneIds() []uint32 {
	if m != nil {
		return m.ZoneIds
	}
	return nil
}

func (m *W2GSBroadcast) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *W2GSBroadcast) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func init() {
	proto.RegisterType((*GS2WServerStart)(nil), "msgProto.GS2WServerStart")
	proto.RegisterType((*W2GSServerStartRes)(nil), "msgProto.W2GSServerStartRes")
	proto.RegisterType((*GS2WSendInfo)(nil), "msgProto.GS2WSendInfo")
	proto.RegisterType((*W2GSSendInfo)(nil), "msgProto.W2GSSendInfo")
	proto.RegisterType((*W2GSDispatch)(nil), "msgProto.W2GSDispatch")
	proto.RegisterType((*W2GSBroadcast)(nil), "msgProto.W2GSBroadcast")
}

func init() { proto.RegisterFile("world.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0xcf, 0x2f, 0xca,
	0x49, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc8, 0x2d, 0x4e, 0x0f, 0x00, 0xb1, 0x94,
	0x34, 0xb9, 0xf8, 0xdd, 0x83, 0x8d, 0xc2, 0x83, 0x53, 0x8b, 0xca, 0x52, 0x8b, 0x82, 0x4b, 0x12,
	0x8b, 0x4a, 0x84, 0xc4, 0xb8, 0xd8, 0xaa, 0xf2, 0xf3, 0x52, 0x3d, 0x53, 0x24, 0x18, 0x15, 0x18,
	0x35, 0x78, 0x83, 0xa0, 0x3c, 0x25, 0x11, 0x2e, 0xa1, 0x70, 0x23, 0xf7, 0x60, 0x24, 0xa5, 0x41,
	0xa9, 0xc5, 0x4a, 0x69, 0x5c, 0x3c, 0x10, 0x03, 0xf2, 0x52, 0x3c, 0xf3, 0xd2, 0xf2, 0x71, 0xe9,
	0x06, 0x89, 0x97, 0x16, 0xa7, 0x16, 0x79, 0xa6, 0x48, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0x41,
	0x79, 0x42, 0x42, 0x5c, 0x2c, 0x25, 0x95, 0x05, 0xa9, 0x12, 0xcc, 0x60, 0xd5, 0x60, 0x36, 0x48,
	0x2c, 0x33, 0x2f, 0x2d, 0x5f, 0x82, 0x05, 0xac, 0x12, 0xcc, 0x06, 0xd9, 0x03, 0xb1, 0x9d, 0xc6,
	0xf6, 0x14, 0x41, 0xec, 0x71, 0xc9, 0x2c, 0x2e, 0x48, 0x2c, 0x49, 0xce, 0x10, 0x92, 0xe3, 0xe2,
	0x4a, 0x2b, 0xca, 0xcf, 0x8d, 0x42, 0xb6, 0x0b, 0x49, 0x44, 0x48, 0x8a, 0x8b, 0xa3, 0x24, 0x1f,
	0x2a, 0xcb, 0x04, 0x96, 0x85, 0xf3, 0x89, 0xb6, 0x33, 0x90, 0x8b, 0x17, 0x64, 0xa7, 0x53, 0x51,
	0x7e, 0x62, 0x4a, 0x72, 0x62, 0x71, 0x89, 0x90, 0x04, 0x17, 0x3b, 0xc4, 0x3b, 0xc5, 0x12, 0x8c,
	0x0a, 0xcc, 0x1a, 0xbc, 0x41, 0x30, 0x2e, 0xdc, 0x48, 0x26, 0x2c, 0x46, 0x32, 0x23, 0x8c, 0x4c,
	0x62, 0x03, 0x47, 0xb4, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x73, 0xa7, 0xf0, 0xc4, 0xf7, 0x01,
	0x00, 0x00,
}
