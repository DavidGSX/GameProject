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
	CCreateRole
	SCreateRole
	CAddMoney
	SMoneyInfo
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

type SCreateRole_ResultType int32

const (
	SCreateRole_UNKNOW_ERR      SCreateRole_ResultType = 0
	SCreateRole_SUCCESS         SCreateRole_ResultType = 1
	SCreateRole_NAME_INVALID    SCreateRole_ResultType = 2
	SCreateRole_NAME_DUPLICATED SCreateRole_ResultType = 3
	SCreateRole_NAME_OVERLEN    SCreateRole_ResultType = 4
	SCreateRole_NAME_SHORTLEN   SCreateRole_ResultType = 5
	SCreateRole_OVER_COUNT      SCreateRole_ResultType = 6
)

var SCreateRole_ResultType_name = map[int32]string{
	0: "UNKNOW_ERR",
	1: "SUCCESS",
	2: "NAME_INVALID",
	3: "NAME_DUPLICATED",
	4: "NAME_OVERLEN",
	5: "NAME_SHORTLEN",
	6: "OVER_COUNT",
}
var SCreateRole_ResultType_value = map[string]int32{
	"UNKNOW_ERR":      0,
	"SUCCESS":         1,
	"NAME_INVALID":    2,
	"NAME_DUPLICATED": 3,
	"NAME_OVERLEN":    4,
	"NAME_SHORTLEN":   5,
	"OVER_COUNT":      6,
}

func (x SCreateRole_ResultType) String() string {
	return proto.EnumName(SCreateRole_ResultType_name, int32(x))
}
func (SCreateRole_ResultType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

type CUserLogin struct {
	UserId   string                  `protobuf:"bytes,1,opt,name=userId" json:"userId,omitempty"`
	Token    string                  `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
	ZoneId   uint32                  `protobuf:"varint,3,opt,name=zoneId" json:"zoneId,omitempty"`
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

func (m *CUserLogin) GetZoneId() uint32 {
	if m != nil {
		return m.ZoneId
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
	SelectRoleId uint64 `protobuf:"varint,1,opt,name=selectRoleId" json:"selectRoleId,omitempty"`
}

func (m *CRoleList) Reset()                    { *m = CRoleList{} }
func (m *CRoleList) String() string            { return proto.CompactTextString(m) }
func (*CRoleList) ProtoMessage()               {}
func (*CRoleList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CRoleList) GetSelectRoleId() uint64 {
	if m != nil {
		return m.SelectRoleId
	}
	return 0
}

type SRoleList struct {
	Roles          []*SRoleList_RoleInfo `protobuf:"bytes,1,rep,name=roles" json:"roles,omitempty"`
	PreLoginRoleId uint64                `protobuf:"varint,2,opt,name=preLoginRoleId" json:"preLoginRoleId,omitempty"`
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

func (m *SRoleList) GetPreLoginRoleId() uint64 {
	if m != nil {
		return m.PreLoginRoleId
	}
	return 0
}

type SRoleList_RoleInfo struct {
	RoleId      uint64          `protobuf:"varint,1,opt,name=roleId" json:"roleId,omitempty"`
	RoleName    string          `protobuf:"bytes,2,opt,name=roleName" json:"roleName,omitempty"`
	Level       uint32          `protobuf:"varint,3,opt,name=level" json:"level,omitempty"`
	School      uint32          `protobuf:"varint,4,opt,name=school" json:"school,omitempty"`
	Components  map[int32]int32 `protobuf:"bytes,5,rep,name=components" json:"components,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	ShowFashion bool            `protobuf:"varint,6,opt,name=showFashion" json:"showFashion,omitempty"`
}

func (m *SRoleList_RoleInfo) Reset()                    { *m = SRoleList_RoleInfo{} }
func (m *SRoleList_RoleInfo) String() string            { return proto.CompactTextString(m) }
func (*SRoleList_RoleInfo) ProtoMessage()               {}
func (*SRoleList_RoleInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

func (m *SRoleList_RoleInfo) GetRoleId() uint64 {
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

func (m *SRoleList_RoleInfo) GetLevel() uint32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *SRoleList_RoleInfo) GetSchool() uint32 {
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

type CCreateRole struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	School uint32 `protobuf:"varint,2,opt,name=school" json:"school,omitempty"`
	Sex    uint32 `protobuf:"varint,4,opt,name=sex" json:"sex,omitempty"`
}

func (m *CCreateRole) Reset()                    { *m = CCreateRole{} }
func (m *CCreateRole) String() string            { return proto.CompactTextString(m) }
func (*CCreateRole) ProtoMessage()               {}
func (*CCreateRole) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CCreateRole) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CCreateRole) GetSchool() uint32 {
	if m != nil {
		return m.School
	}
	return 0
}

func (m *CCreateRole) GetSex() uint32 {
	if m != nil {
		return m.Sex
	}
	return 0
}

type SCreateRole struct {
	Res  SCreateRole_ResultType `protobuf:"varint,1,opt,name=res,enum=protocol.SCreateRole_ResultType" json:"res,omitempty"`
	Info *SRoleList_RoleInfo    `protobuf:"bytes,2,opt,name=info" json:"info,omitempty"`
}

func (m *SCreateRole) Reset()                    { *m = SCreateRole{} }
func (m *SCreateRole) String() string            { return proto.CompactTextString(m) }
func (*SCreateRole) ProtoMessage()               {}
func (*SCreateRole) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SCreateRole) GetRes() SCreateRole_ResultType {
	if m != nil {
		return m.Res
	}
	return SCreateRole_UNKNOW_ERR
}

func (m *SCreateRole) GetInfo() *SRoleList_RoleInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type CAddMoney struct {
	RoleId uint64 `protobuf:"varint,1,opt,name=roleId" json:"roleId,omitempty"`
	Num    uint32 `protobuf:"varint,2,opt,name=num" json:"num,omitempty"`
}

func (m *CAddMoney) Reset()                    { *m = CAddMoney{} }
func (m *CAddMoney) String() string            { return proto.CompactTextString(m) }
func (*CAddMoney) ProtoMessage()               {}
func (*CAddMoney) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CAddMoney) GetRoleId() uint64 {
	if m != nil {
		return m.RoleId
	}
	return 0
}

func (m *CAddMoney) GetNum() uint32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type SMoneyInfo struct {
	RoleId uint64 `protobuf:"varint,1,opt,name=roleId" json:"roleId,omitempty"`
	Total  uint32 `protobuf:"varint,2,opt,name=total" json:"total,omitempty"`
}

func (m *SMoneyInfo) Reset()                    { *m = SMoneyInfo{} }
func (m *SMoneyInfo) String() string            { return proto.CompactTextString(m) }
func (*SMoneyInfo) ProtoMessage()               {}
func (*SMoneyInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *SMoneyInfo) GetRoleId() uint64 {
	if m != nil {
		return m.RoleId
	}
	return 0
}

func (m *SMoneyInfo) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func init() {
	proto.RegisterType((*CUserLogin)(nil), "protocol.CUserLogin")
	proto.RegisterType((*SUserLogin)(nil), "protocol.SUserLogin")
	proto.RegisterType((*CRoleList)(nil), "protocol.CRoleList")
	proto.RegisterType((*SRoleList)(nil), "protocol.SRoleList")
	proto.RegisterType((*SRoleList_RoleInfo)(nil), "protocol.SRoleList.RoleInfo")
	proto.RegisterType((*CCreateRole)(nil), "protocol.CCreateRole")
	proto.RegisterType((*SCreateRole)(nil), "protocol.SCreateRole")
	proto.RegisterType((*CAddMoney)(nil), "protocol.CAddMoney")
	proto.RegisterType((*SMoneyInfo)(nil), "protocol.SMoneyInfo")
	proto.RegisterEnum("protocol.CUserLogin_PlatformType", CUserLogin_PlatformType_name, CUserLogin_PlatformType_value)
	proto.RegisterEnum("protocol.SUserLogin_LoginResType", SUserLogin_LoginResType_name, SUserLogin_LoginResType_value)
	proto.RegisterEnum("protocol.SCreateRole_ResultType", SCreateRole_ResultType_name, SCreateRole_ResultType_value)
}

func init() { proto.RegisterFile("login.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 664 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x4d, 0x6f, 0xda, 0x40,
	0x10, 0x8d, 0x3f, 0x20, 0x66, 0x4c, 0xc8, 0x76, 0x5b, 0x55, 0x28, 0xea, 0x81, 0xfa, 0x50, 0x71,
	0xa8, 0xdc, 0x8a, 0xaa, 0x52, 0x15, 0x29, 0x07, 0x6a, 0x5c, 0xc5, 0x8a, 0x31, 0xd1, 0x1a, 0x92,
	0xe6, 0x84, 0x28, 0x6c, 0x1a, 0x14, 0xb3, 0x8b, 0x6c, 0x93, 0x96, 0xde, 0xf3, 0x23, 0xaa, 0xfe,
	0xaf, 0xfe, 0x99, 0x5e, 0xaa, 0x5d, 0x1b, 0xe3, 0x44, 0x2a, 0x3d, 0xb1, 0xf3, 0xe6, 0xcd, 0xce,
	0xbc, 0xf1, 0x5b, 0xc0, 0x8c, 0xf8, 0xd7, 0x39, 0xb3, 0x97, 0x31, 0x4f, 0x39, 0x36, 0xe4, 0xcf,
	0x94, 0x47, 0xd6, 0x6f, 0x05, 0xc0, 0x19, 0x25, 0x34, 0xf6, 0x45, 0x1a, 0x3f, 0x87, 0xea, 0x2a,
	0xa1, 0xb1, 0x37, 0x6b, 0x2a, 0x2d, 0xa5, 0x5d, 0x23, 0x79, 0x84, 0x9f, 0x41, 0x25, 0xe5, 0xb7,
	0x94, 0x35, 0x55, 0x09, 0x67, 0x81, 0x60, 0xff, 0xe0, 0x8c, 0x7a, 0xb3, 0xa6, 0xd6, 0x52, 0xda,
	0x07, 0x24, 0x8f, 0xf0, 0x09, 0x18, 0xcb, 0x68, 0x92, 0x5e, 0xf3, 0x78, 0xd1, 0xd4, 0x5b, 0x4a,
	0xbb, 0xd1, 0x79, 0x69, 0x6f, 0x3a, 0xda, 0xdb, 0x6e, 0xf6, 0x79, 0x4e, 0x1a, 0xae, 0x97, 0x94,
	0x14, 0x25, 0x96, 0x03, 0xf5, 0x72, 0x06, 0x03, 0x54, 0x47, 0xc1, 0x59, 0x30, 0xb8, 0x44, 0x7b,
	0x78, 0x1f, 0x34, 0x6f, 0x10, 0x22, 0x05, 0x9b, 0xb0, 0xdf, 0x0d, 0x7a, 0x64, 0xe0, 0xf5, 0x90,
	0x2a, 0xd0, 0xab, 0xab, 0x8f, 0x48, 0xc3, 0x06, 0xe8, 0x57, 0xa7, 0x7e, 0x1f, 0xe9, 0xd6, 0x4f,
	0x05, 0x20, 0xdc, 0x0a, 0x3b, 0x01, 0x43, 0x2e, 0x80, 0xd0, 0x44, 0x4a, 0x7b, 0x30, 0xd2, 0x96,
	0x67, 0xfb, 0x39, 0x29, 0x1b, 0x69, 0x53, 0x62, 0xf9, 0x50, 0x2f, 0x67, 0x70, 0x03, 0x20, 0x1b,
	0x69, 0xec, 0x12, 0x82, 0xf6, 0xc4, 0x34, 0xe1, 0xc8, 0x71, 0xdc, 0x50, 0x8c, 0xd6, 0x00, 0x38,
	0xef, 0x86, 0xe1, 0x65, 0x4f, 0x26, 0x55, 0x7c, 0x08, 0xe6, 0xd0, 0xeb, 0xbb, 0x63, 0xf7, 0xb3,
	0xe3, 0xba, 0x3d, 0xa4, 0x59, 0x6f, 0xa0, 0xe6, 0x10, 0x1e, 0x51, 0x7f, 0x9e, 0xa4, 0xd8, 0x82,
	0x7a, 0x42, 0x23, 0x3a, 0x4d, 0x05, 0x92, 0x2f, 0x5e, 0x27, 0x0f, 0x30, 0xeb, 0x5e, 0x83, 0x5a,
	0x58, 0x54, 0x74, 0xa0, 0x12, 0xf3, 0x48, 0x0a, 0xd1, 0xda, 0x66, 0xe7, 0x45, 0x49, 0xc8, 0x86,
	0x63, 0xcb, 0x42, 0x76, 0xcd, 0x49, 0x46, 0xc5, 0xaf, 0xa0, 0xb1, 0x8c, 0x69, 0xa6, 0x21, 0xeb,
	0xa3, 0xca, 0x3e, 0x8f, 0xd0, 0xa3, 0x5f, 0x2a, 0x18, 0x9b, 0x5a, 0xf1, 0x7d, 0xe3, 0xf2, 0x50,
	0x79, 0x84, 0x8f, 0xc0, 0x10, 0xa7, 0x60, 0xb2, 0xa0, 0xb9, 0x21, 0x8a, 0x58, 0x38, 0x25, 0xa2,
	0x77, 0x34, 0xca, 0x2d, 0x91, 0x05, 0xe2, 0xa6, 0x64, 0x7a, 0xc3, 0x79, 0x24, 0xfd, 0x70, 0x40,
	0xf2, 0x08, 0xfb, 0x00, 0x53, 0xbe, 0x58, 0x72, 0x46, 0x59, 0x9a, 0x34, 0x2b, 0x52, 0xcf, 0xeb,
	0x5d, 0x7a, 0x6c, 0xa7, 0xa0, 0xbb, 0x2c, 0x8d, 0xd7, 0xa4, 0x54, 0x8f, 0x5b, 0x60, 0x26, 0x37,
	0xfc, 0xdb, 0xa7, 0x49, 0x72, 0x33, 0xe7, 0xac, 0x59, 0x6d, 0x29, 0x6d, 0x83, 0x94, 0xa1, 0xa3,
	0x13, 0x38, 0x7c, 0x74, 0x01, 0x46, 0xa0, 0xdd, 0xd2, 0xb5, 0x54, 0x58, 0x21, 0xe2, 0x28, 0x24,
	0xdc, 0x4d, 0xa2, 0x55, 0xa6, 0xad, 0x42, 0xb2, 0xe0, 0x58, 0xfd, 0xa0, 0x58, 0x67, 0x60, 0x3a,
	0x4e, 0x4c, 0x27, 0x29, 0x15, 0xf3, 0x60, 0x0c, 0x3a, 0x13, 0x3b, 0xc8, 0xde, 0x8a, 0x3c, 0x97,
	0x94, 0xaa, 0x0f, 0x94, 0x22, 0xd0, 0x12, 0xfa, 0x3d, 0x97, 0x2f, 0x8e, 0xd6, 0x1f, 0x05, 0xcc,
	0xb0, 0x74, 0x5b, 0x07, 0xb4, 0xb8, 0x70, 0x67, 0xab, 0xb4, 0x84, 0x2d, 0xc7, 0x26, 0x34, 0x59,
	0x45, 0xa9, 0x34, 0xa7, 0x20, 0xe3, 0xb7, 0xa0, 0xcf, 0xd9, 0x35, 0x97, 0xbd, 0xfe, 0xe7, 0x04,
	0xc9, 0xb4, 0xee, 0x15, 0x80, 0xed, 0x2d, 0xbb, 0x8d, 0x8c, 0xa0, 0x1e, 0x74, 0xfb, 0xee, 0xd8,
	0x0b, 0x2e, 0xba, 0xbe, 0x7c, 0x68, 0x4f, 0xe1, 0x50, 0x22, 0xbd, 0xd1, 0xb9, 0xef, 0x39, 0xdd,
	0xa1, 0xb0, 0x73, 0x41, 0x1b, 0x5c, 0xb8, 0xc4, 0x77, 0x03, 0xa4, 0xe3, 0x27, 0x70, 0x20, 0x91,
	0xf0, 0x74, 0x40, 0x86, 0x02, 0xaa, 0x88, 0x46, 0x22, 0x3f, 0x76, 0x06, 0xa3, 0x60, 0x88, 0xaa,
	0xd6, 0x7b, 0xa8, 0x39, 0xdd, 0xd9, 0xac, 0xcf, 0x19, 0x5d, 0xff, 0xd3, 0x68, 0x08, 0x34, 0xb6,
	0x5a, 0xe4, 0x9b, 0x14, 0x47, 0xeb, 0x18, 0x20, 0x94, 0x35, 0x3b, 0x0d, 0x2a, 0xff, 0xae, 0xd2,
	0xc9, 0xe6, 0x1b, 0x64, 0xc1, 0x97, 0xaa, 0xdc, 0xce, 0xbb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x23, 0xc0, 0xc2, 0x8c, 0x0b, 0x05, 0x00, 0x00,
}
