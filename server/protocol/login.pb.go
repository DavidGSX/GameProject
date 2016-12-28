// Code generated by protoc-gen-go.
// source: login.proto
// DO NOT EDIT!

/*
Package protocol is a generated protocol buffer package.

It is generated from these files:
	login.proto

It has these top-level messages:
	CUserLogin
	SUserLogin
	CRoleList
	SRoleList
*/
package protocol

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

type CUserLogin_PlatformType int32

const (
	CUserLogin_UNKNOW  CUserLogin_PlatformType = 0
	CUserLogin_IOS     CUserLogin_PlatformType = 1
	CUserLogin_ANDROID CUserLogin_PlatformType = 2
	CUserLogin_YYB     CUserLogin_PlatformType = 3
	CUserLogin_YHLM    CUserLogin_PlatformType = 4
)

var CUserLogin_PlatformType_name = map[int32]string{
	0: "UNKNOW",
	1: "IOS",
	2: "ANDROID",
	3: "YYB",
	4: "YHLM",
}
var CUserLogin_PlatformType_value = map[string]int32{
	"UNKNOW":  0,
	"IOS":     1,
	"ANDROID": 2,
	"YYB":     3,
	"YHLM":    4,
}

func (x CUserLogin_PlatformType) String() string {
	return proto.EnumName(CUserLogin_PlatformType_name, int32(x))
}
func (CUserLogin_PlatformType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type SUserLogin_LoginResType int32

const (
	SUserLogin_UNKNOW_ERR  SUserLogin_LoginResType = 0
	SUserLogin_SUCCESS     SUserLogin_LoginResType = 1
	SUserLogin_PASSWD_ERR  SUserLogin_LoginResType = 2
	SUserLogin_TIME_EXCEED SUserLogin_LoginResType = 3
)

var SUserLogin_LoginResType_name = map[int32]string{
	0: "UNKNOW_ERR",
	1: "SUCCESS",
	2: "PASSWD_ERR",
	3: "TIME_EXCEED",
}
var SUserLogin_LoginResType_value = map[string]int32{
	"UNKNOW_ERR":  0,
	"SUCCESS":     1,
	"PASSWD_ERR":  2,
	"TIME_EXCEED": 3,
}

func (x SUserLogin_LoginResType) String() string {
	return proto.EnumName(SUserLogin_LoginResType_name, int32(x))
}
func (SUserLogin_LoginResType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type CUserLogin struct {
	UserId   string                  `protobuf:"bytes,1,opt,name=userId" json:"userId,omitempty"`
	Token    string                  `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
	Zoneid   int32                   `protobuf:"varint,3,opt,name=zoneid" json:"zoneid,omitempty"`
	Platform CUserLogin_PlatformType `protobuf:"varint,4,opt,name=platform,enum=protocol.CUserLogin_PlatformType" json:"platform,omitempty"`
}

func (m *CUserLogin) Reset()                    { *m = CUserLogin{} }
func (m *CUserLogin) String() string            { return proto.CompactTextString(m) }
func (*CUserLogin) ProtoMessage()               {}
func (*CUserLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CUserLogin) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *CUserLogin) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CUserLogin) GetZoneid() int32 {
	if m != nil {
		return m.Zoneid
	}
	return 0
}

func (m *CUserLogin) GetPlatform() CUserLogin_PlatformType {
	if m != nil {
		return m.Platform
	}
	return CUserLogin_UNKNOW
}

type SUserLogin struct {
	LoginRes SUserLogin_LoginResType `protobuf:"varint,1,opt,name=loginRes,enum=protocol.SUserLogin_LoginResType" json:"loginRes,omitempty"`
}

func (m *SUserLogin) Reset()                    { *m = SUserLogin{} }
func (m *SUserLogin) String() string            { return proto.CompactTextString(m) }
func (*SUserLogin) ProtoMessage()               {}
func (*SUserLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SUserLogin) GetLoginRes() SUserLogin_LoginResType {
	if m != nil {
		return m.LoginRes
	}
	return SUserLogin_UNKNOW_ERR
}

type CRoleList struct {
	SelectRoleId int32 `protobuf:"varint,1,opt,name=selectRoleId" json:"selectRoleId,omitempty"`
}

func (m *CRoleList) Reset()                    { *m = CRoleList{} }
func (m *CRoleList) String() string            { return proto.CompactTextString(m) }
func (*CRoleList) ProtoMessage()               {}
func (*CRoleList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CRoleList) GetSelectRoleId() int32 {
	if m != nil {
		return m.SelectRoleId
	}
	return 0
}

type SRoleList struct {
	Roles           []*SRoleList_RoleInfo `protobuf:"bytes,1,rep,name=roles" json:"roles,omitempty"`
	PreLoginRoleId  int64                 `protobuf:"varint,2,opt,name=preLoginRoleId" json:"preLoginRoleId,omitempty"`
	PreRoleInBattle bool                  `protobuf:"varint,3,opt,name=preRoleInBattle" json:"preRoleInBattle,omitempty"`
}

func (m *SRoleList) Reset()                    { *m = SRoleList{} }
func (m *SRoleList) String() string            { return proto.CompactTextString(m) }
func (*SRoleList) ProtoMessage()               {}
func (*SRoleList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SRoleList) GetRoles() []*SRoleList_RoleInfo {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *SRoleList) GetPreLoginRoleId() int64 {
	if m != nil {
		return m.PreLoginRoleId
	}
	return 0
}

func (m *SRoleList) GetPreRoleInBattle() bool {
	if m != nil {
		return m.PreRoleInBattle
	}
	return false
}

type SRoleList_RoleInfo struct {
	RoleId      int64           `protobuf:"varint,1,opt,name=roleId" json:"roleId,omitempty"`
	RoleName    string          `protobuf:"bytes,2,opt,name=roleName" json:"roleName,omitempty"`
	Level       int32           `protobuf:"varint,3,opt,name=level" json:"level,omitempty"`
	School      int32           `protobuf:"varint,4,opt,name=school" json:"school,omitempty"`
	Components  map[int32]int32 `protobuf:"bytes,5,rep,name=components" json:"components,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	ShowFashion bool            `protobuf:"varint,6,opt,name=showFashion" json:"showFashion,omitempty"`
}

func (m *SRoleList_RoleInfo) Reset()                    { *m = SRoleList_RoleInfo{} }
func (m *SRoleList_RoleInfo) String() string            { return proto.CompactTextString(m) }
func (*SRoleList_RoleInfo) ProtoMessage()               {}
func (*SRoleList_RoleInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

func (m *SRoleList_RoleInfo) GetRoleId() int64 {
	if m != nil {
		return m.RoleId
	}
	return 0
}

func (m *SRoleList_RoleInfo) GetRoleName() string {
	if m != nil {
		return m.RoleName
	}
	return ""
}

func (m *SRoleList_RoleInfo) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *SRoleList_RoleInfo) GetSchool() int32 {
	if m != nil {
		return m.School
	}
	return 0
}

func (m *SRoleList_RoleInfo) GetComponents() map[int32]int32 {
	if m != nil {
		return m.Components
	}
	return nil
}

func (m *SRoleList_RoleInfo) GetShowFashion() bool {
	if m != nil {
		return m.ShowFashion
	}
	return false
}

func init() {
	proto.RegisterType((*CUserLogin)(nil), "protocol.CUserLogin")
	proto.RegisterType((*SUserLogin)(nil), "protocol.SUserLogin")
	proto.RegisterType((*CRoleList)(nil), "protocol.CRoleList")
	proto.RegisterType((*SRoleList)(nil), "protocol.SRoleList")
	proto.RegisterType((*SRoleList_RoleInfo)(nil), "protocol.SRoleList.RoleInfo")
	proto.RegisterEnum("protocol.CUserLogin_PlatformType", CUserLogin_PlatformType_name, CUserLogin_PlatformType_value)
	proto.RegisterEnum("protocol.SUserLogin_LoginResType", SUserLogin_LoginResType_name, SUserLogin_LoginResType_value)
}

func init() { proto.RegisterFile("login.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 504 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xad, 0xed, 0x3a, 0x75, 0xc6, 0x51, 0x62, 0xad, 0x10, 0xb2, 0x22, 0x0e, 0xc1, 0x07, 0xe4,
	0x03, 0x32, 0x52, 0xb8, 0x20, 0xa4, 0x1c, 0x5a, 0xc7, 0x88, 0x08, 0x37, 0xa9, 0xd6, 0x8d, 0x4a,
	0x4e, 0x95, 0x49, 0xb7, 0x24, 0xea, 0xc6, 0x6b, 0xd9, 0xdb, 0xa2, 0xf0, 0x4f, 0x10, 0x3f, 0x82,
	0x7f, 0xc3, 0xef, 0x41, 0xfb, 0x91, 0xc4, 0xcd, 0x81, 0x93, 0xfd, 0xde, 0xbc, 0xd9, 0x79, 0x6f,
	0x77, 0xc0, 0xa5, 0xec, 0xfb, 0xba, 0x88, 0xca, 0x8a, 0x71, 0x86, 0x1c, 0xf9, 0x59, 0x32, 0x1a,
	0xfc, 0x35, 0x00, 0xe2, 0x79, 0x4d, 0xaa, 0x54, 0x94, 0xd1, 0x4b, 0x68, 0x3d, 0xd6, 0xa4, 0x9a,
	0xdc, 0xf9, 0xc6, 0xc0, 0x08, 0xdb, 0x58, 0x23, 0xf4, 0x02, 0x6c, 0xce, 0x1e, 0x48, 0xe1, 0x9b,
	0x92, 0x56, 0x40, 0xa8, 0x7f, 0xb2, 0x82, 0xac, 0xef, 0x7c, 0x6b, 0x60, 0x84, 0x36, 0xd6, 0x08,
	0x8d, 0xc0, 0x29, 0x69, 0xce, 0xef, 0x59, 0xb5, 0xf1, 0x4f, 0x07, 0x46, 0xd8, 0x1d, 0xbe, 0x8e,
	0x76, 0x13, 0xa3, 0xc3, 0xb4, 0xe8, 0x4a, 0x8b, 0xae, 0xb7, 0x25, 0xc1, 0xfb, 0x96, 0x20, 0x86,
	0x4e, 0xb3, 0x82, 0x00, 0x5a, 0xf3, 0xe9, 0x97, 0xe9, 0xec, 0xc6, 0x3b, 0x41, 0x67, 0x60, 0x4d,
	0x66, 0x99, 0x67, 0x20, 0x17, 0xce, 0xce, 0xa7, 0x63, 0x3c, 0x9b, 0x8c, 0x3d, 0x53, 0xb0, 0x8b,
	0xc5, 0x85, 0x67, 0x21, 0x07, 0x4e, 0x17, 0x9f, 0xd3, 0x4b, 0xef, 0x34, 0xf8, 0x65, 0x00, 0x64,
	0x87, 0x60, 0x23, 0x70, 0xe4, 0x05, 0x60, 0x52, 0xcb, 0x68, 0xcf, 0x2c, 0x1d, 0x74, 0x51, 0xaa,
	0x45, 0xca, 0xd2, 0xae, 0x25, 0x48, 0xa1, 0xd3, 0xac, 0xa0, 0x2e, 0x80, 0xb2, 0x74, 0x9b, 0x60,
	0xec, 0x9d, 0x08, 0x37, 0xd9, 0x3c, 0x8e, 0x93, 0x4c, 0x58, 0xeb, 0x02, 0x5c, 0x9d, 0x67, 0xd9,
	0xcd, 0x58, 0x16, 0x4d, 0xd4, 0x03, 0xf7, 0x7a, 0x72, 0x99, 0xdc, 0x26, 0x5f, 0xe3, 0x24, 0x19,
	0x7b, 0x56, 0xf0, 0x0e, 0xda, 0x31, 0x66, 0x94, 0xa4, 0xeb, 0x9a, 0xa3, 0x00, 0x3a, 0x35, 0xa1,
	0x64, 0xc9, 0x05, 0xa3, 0x2f, 0xde, 0xc6, 0xcf, 0xb8, 0xe0, 0x8f, 0x05, 0xed, 0x6c, 0xdf, 0x31,
	0x04, 0xbb, 0x62, 0x54, 0x06, 0xb1, 0x42, 0x77, 0xf8, 0xaa, 0x11, 0x64, 0xa7, 0x89, 0x64, 0x63,
	0x71, 0xcf, 0xb0, 0x92, 0xa2, 0x37, 0xd0, 0x2d, 0x2b, 0xa2, 0x32, 0xa8, 0x39, 0xe2, 0x25, 0x2d,
	0x7c, 0xc4, 0xa2, 0x10, 0x7a, 0x65, 0x45, 0x54, 0xf7, 0x45, 0xce, 0x39, 0x25, 0xf2, 0x6d, 0x1d,
	0x7c, 0x4c, 0xf7, 0x7f, 0x9b, 0xe0, 0xec, 0xa6, 0x88, 0x4d, 0xa8, 0x0e, 0xf6, 0x2d, 0xac, 0x11,
	0xea, 0x83, 0x23, 0xfe, 0xa6, 0xf9, 0x86, 0xe8, 0xd5, 0xd9, 0x63, 0xb1, 0x53, 0x94, 0x3c, 0x11,
	0xaa, 0x97, 0x47, 0x01, 0x71, 0x52, 0xbd, 0x5c, 0x31, 0x46, 0xe5, 0xe6, 0xd8, 0x58, 0x23, 0x94,
	0x02, 0x2c, 0xd9, 0xa6, 0x64, 0x05, 0x29, 0x78, 0xed, 0xdb, 0x32, 0xf9, 0xdb, 0xff, 0x25, 0x8f,
	0xe2, 0xbd, 0x3c, 0x29, 0x78, 0xb5, 0xc5, 0x8d, 0x7e, 0x34, 0x00, 0xb7, 0x5e, 0xb1, 0x1f, 0x9f,
	0xf2, 0x7a, 0xb5, 0x66, 0x85, 0xdf, 0x92, 0x11, 0x9b, 0x54, 0x7f, 0x04, 0xbd, 0xa3, 0x03, 0x90,
	0x07, 0xd6, 0x03, 0xd9, 0xea, 0x07, 0x12, 0xbf, 0x22, 0xc2, 0x53, 0x4e, 0x1f, 0x55, 0x36, 0x1b,
	0x2b, 0xf0, 0xd1, 0xfc, 0x60, 0x7c, 0x6b, 0x49, 0x67, 0xef, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff,
	0xfa, 0xf2, 0xef, 0x7b, 0x77, 0x03, 0x00, 0x00,
}
