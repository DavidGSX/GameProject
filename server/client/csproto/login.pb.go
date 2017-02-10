// Code generated by protoc-gen-go.
// source: login.proto
// DO NOT EDIT!

/*
Package csproto is a generated protocol buffer package.

It is generated from these files:
	login.proto

It has these top-level messages:
	CUserLogin
	SUserLogin
	CRoleList
	SRoleList
	CCreateRole
	SCreateRole
	CEnterWorld
	SEnterWorld
	CAddMoney
	SMoneyInfo
	CAddLevel
	SLevelInfo
	CReqServerRoleInfos
	SServerRoleInfos
*/
package csproto

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
	Platform CUserLogin_PlatformType `protobuf:"varint,4,opt,name=platform,enum=csproto.CUserLogin_PlatformType" json:"platform,omitempty"`
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
	LoginRes SUserLogin_LoginResType `protobuf:"varint,1,opt,name=loginRes,enum=csproto.SUserLogin_LoginResType" json:"loginRes,omitempty"`
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
	Res  SCreateRole_ResultType `protobuf:"varint,1,opt,name=res,enum=csproto.SCreateRole_ResultType" json:"res,omitempty"`
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

type CEnterWorld struct {
	RoleId uint64 `protobuf:"varint,1,opt,name=roleId" json:"roleId,omitempty"`
}

func (m *CEnterWorld) Reset()                    { *m = CEnterWorld{} }
func (m *CEnterWorld) String() string            { return proto.CompactTextString(m) }
func (*CEnterWorld) ProtoMessage()               {}
func (*CEnterWorld) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CEnterWorld) GetRoleId() uint64 {
	if m != nil {
		return m.RoleId
	}
	return 0
}

type SEnterWorld struct {
	RoleId uint64 `protobuf:"varint,1,opt,name=roleId" json:"roleId,omitempty"`
}

func (m *SEnterWorld) Reset()                    { *m = SEnterWorld{} }
func (m *SEnterWorld) String() string            { return proto.CompactTextString(m) }
func (*SEnterWorld) ProtoMessage()               {}
func (*SEnterWorld) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *SEnterWorld) GetRoleId() uint64 {
	if m != nil {
		return m.RoleId
	}
	return 0
}

type CAddMoney struct {
	Num uint32 `protobuf:"varint,1,opt,name=num" json:"num,omitempty"`
}

func (m *CAddMoney) Reset()                    { *m = CAddMoney{} }
func (m *CAddMoney) String() string            { return proto.CompactTextString(m) }
func (*CAddMoney) ProtoMessage()               {}
func (*CAddMoney) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *CAddMoney) GetNum() uint32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type SMoneyInfo struct {
	Money uint32 `protobuf:"varint,1,opt,name=money" json:"money,omitempty"`
}

func (m *SMoneyInfo) Reset()                    { *m = SMoneyInfo{} }
func (m *SMoneyInfo) String() string            { return proto.CompactTextString(m) }
func (*SMoneyInfo) ProtoMessage()               {}
func (*SMoneyInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *SMoneyInfo) GetMoney() uint32 {
	if m != nil {
		return m.Money
	}
	return 0
}

type CAddLevel struct {
	Num uint32 `protobuf:"varint,1,opt,name=num" json:"num,omitempty"`
}

func (m *CAddLevel) Reset()                    { *m = CAddLevel{} }
func (m *CAddLevel) String() string            { return proto.CompactTextString(m) }
func (*CAddLevel) ProtoMessage()               {}
func (*CAddLevel) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *CAddLevel) GetNum() uint32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type SLevelInfo struct {
	Level uint32 `protobuf:"varint,1,opt,name=level" json:"level,omitempty"`
}

func (m *SLevelInfo) Reset()                    { *m = SLevelInfo{} }
func (m *SLevelInfo) String() string            { return proto.CompactTextString(m) }
func (*SLevelInfo) ProtoMessage()               {}
func (*SLevelInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *SLevelInfo) GetLevel() uint32 {
	if m != nil {
		return m.Level
	}
	return 0
}

type CReqServerRoleInfos struct {
}

func (m *CReqServerRoleInfos) Reset()                    { *m = CReqServerRoleInfos{} }
func (m *CReqServerRoleInfos) String() string            { return proto.CompactTextString(m) }
func (*CReqServerRoleInfos) ProtoMessage()               {}
func (*CReqServerRoleInfos) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type SServerRoleInfos struct {
	RoleId uint64                       `protobuf:"varint,1,opt,name=roleId" json:"roleId,omitempty"`
	Info   []*SServerRoleInfos_RoleInfo `protobuf:"bytes,2,rep,name=info" json:"info,omitempty"`
}

func (m *SServerRoleInfos) Reset()                    { *m = SServerRoleInfos{} }
func (m *SServerRoleInfos) String() string            { return proto.CompactTextString(m) }
func (*SServerRoleInfos) ProtoMessage()               {}
func (*SServerRoleInfos) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *SServerRoleInfos) GetRoleId() uint64 {
	if m != nil {
		return m.RoleId
	}
	return 0
}

func (m *SServerRoleInfos) GetInfo() []*SServerRoleInfos_RoleInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type SServerRoleInfos_RoleInfo struct {
	ZoneId   uint32 `protobuf:"varint,1,opt,name=zoneId" json:"zoneId,omitempty"`
	RoleId   uint64 `protobuf:"varint,2,opt,name=roleId" json:"roleId,omitempty"`
	RoleName string `protobuf:"bytes,3,opt,name=roleName" json:"roleName,omitempty"`
	Level    uint32 `protobuf:"varint,4,opt,name=level" json:"level,omitempty"`
	School   uint32 `protobuf:"varint,5,opt,name=school" json:"school,omitempty"`
	Sex      uint32 `protobuf:"varint,6,opt,name=sex" json:"sex,omitempty"`
	Lasttime uint64 `protobuf:"varint,7,opt,name=lasttime" json:"lasttime,omitempty"`
}

func (m *SServerRoleInfos_RoleInfo) Reset()                    { *m = SServerRoleInfos_RoleInfo{} }
func (m *SServerRoleInfos_RoleInfo) String() string            { return proto.CompactTextString(m) }
func (*SServerRoleInfos_RoleInfo) ProtoMessage()               {}
func (*SServerRoleInfos_RoleInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13, 0} }

func (m *SServerRoleInfos_RoleInfo) GetZoneId() uint32 {
	if m != nil {
		return m.ZoneId
	}
	return 0
}

func (m *SServerRoleInfos_RoleInfo) GetRoleId() uint64 {
	if m != nil {
		return m.RoleId
	}
	return 0
}

func (m *SServerRoleInfos_RoleInfo) GetRoleName() string {
	if m != nil {
		return m.RoleName
	}
	return ""
}

func (m *SServerRoleInfos_RoleInfo) GetLevel() uint32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *SServerRoleInfos_RoleInfo) GetSchool() uint32 {
	if m != nil {
		return m.School
	}
	return 0
}

func (m *SServerRoleInfos_RoleInfo) GetSex() uint32 {
	if m != nil {
		return m.Sex
	}
	return 0
}

func (m *SServerRoleInfos_RoleInfo) GetLasttime() uint64 {
	if m != nil {
		return m.Lasttime
	}
	return 0
}

func init() {
	proto.RegisterType((*CUserLogin)(nil), "csproto.CUserLogin")
	proto.RegisterType((*SUserLogin)(nil), "csproto.SUserLogin")
	proto.RegisterType((*CRoleList)(nil), "csproto.CRoleList")
	proto.RegisterType((*SRoleList)(nil), "csproto.SRoleList")
	proto.RegisterType((*SRoleList_RoleInfo)(nil), "csproto.SRoleList.RoleInfo")
	proto.RegisterType((*CCreateRole)(nil), "csproto.CCreateRole")
	proto.RegisterType((*SCreateRole)(nil), "csproto.SCreateRole")
	proto.RegisterType((*CEnterWorld)(nil), "csproto.CEnterWorld")
	proto.RegisterType((*SEnterWorld)(nil), "csproto.SEnterWorld")
	proto.RegisterType((*CAddMoney)(nil), "csproto.CAddMoney")
	proto.RegisterType((*SMoneyInfo)(nil), "csproto.SMoneyInfo")
	proto.RegisterType((*CAddLevel)(nil), "csproto.CAddLevel")
	proto.RegisterType((*SLevelInfo)(nil), "csproto.SLevelInfo")
	proto.RegisterType((*CReqServerRoleInfos)(nil), "csproto.CReqServerRoleInfos")
	proto.RegisterType((*SServerRoleInfos)(nil), "csproto.SServerRoleInfos")
	proto.RegisterType((*SServerRoleInfos_RoleInfo)(nil), "csproto.SServerRoleInfos.RoleInfo")
	proto.RegisterEnum("csproto.CUserLogin_PlatformType", CUserLogin_PlatformType_name, CUserLogin_PlatformType_value)
	proto.RegisterEnum("csproto.SUserLogin_LoginResType", SUserLogin_LoginResType_name, SUserLogin_LoginResType_value)
	proto.RegisterEnum("csproto.SCreateRole_ResultType", SCreateRole_ResultType_name, SCreateRole_ResultType_value)
}

func init() { proto.RegisterFile("login.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 780 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdb, 0x6e, 0xe2, 0x46,
	0x18, 0x8e, 0x4f, 0x1c, 0x7e, 0x03, 0x71, 0x27, 0x69, 0x85, 0xa8, 0xaa, 0xa2, 0x91, 0x5a, 0x21,
	0x55, 0x72, 0x94, 0x54, 0xaa, 0xaa, 0x2a, 0xbd, 0xa0, 0xc6, 0x55, 0x50, 0x8c, 0x89, 0xc6, 0x90,
	0x34, 0x57, 0x88, 0xc2, 0xa4, 0x41, 0x31, 0x1e, 0x6a, 0x9b, 0xb4, 0xe9, 0x6d, 0xb5, 0xcf, 0xb0,
	0xab, 0x7d, 0x8c, 0x7d, 0x97, 0x7d, 0x98, 0xbd, 0x5b, 0xcd, 0xd8, 0xd8, 0x06, 0x85, 0xec, 0xde,
	0xfd, 0x87, 0xef, 0x3f, 0x7c, 0x9e, 0xef, 0x37, 0xe8, 0x3e, 0xfb, 0x6b, 0x11, 0x98, 0xab, 0x90,
	0xc5, 0x0c, 0x95, 0x67, 0x91, 0x30, 0xf0, 0x7b, 0x09, 0xc0, 0x1a, 0x47, 0x34, 0x74, 0x78, 0x16,
	0x7d, 0x05, 0xa5, 0x75, 0x44, 0xc3, 0xfe, 0xbc, 0x29, 0xb5, 0xa5, 0x4e, 0x95, 0xa4, 0x1e, 0x3a,
	0x06, 0x2d, 0x66, 0x0f, 0x34, 0x68, 0xca, 0x22, 0x9c, 0x38, 0x1c, 0xfd, 0x1f, 0x0b, 0x68, 0x7f,
	0xde, 0x54, 0xda, 0x52, 0xa7, 0x4e, 0x52, 0x0f, 0x9d, 0x43, 0x65, 0xe5, 0x4f, 0xe3, 0x3b, 0x16,
	0x2e, 0x9b, 0x6a, 0x5b, 0xea, 0x34, 0xce, 0xda, 0x66, 0x3a, 0xd0, 0xcc, 0x87, 0x99, 0x57, 0x29,
	0x66, 0xf4, 0xb4, 0xa2, 0x24, 0xab, 0xc0, 0x16, 0xd4, 0x8a, 0x19, 0x04, 0x50, 0x1a, 0xbb, 0x97,
	0xee, 0xf0, 0xc6, 0x38, 0x40, 0x65, 0x50, 0xfa, 0x43, 0xcf, 0x90, 0x90, 0x0e, 0xe5, 0xae, 0xdb,
	0x23, 0xc3, 0x7e, 0xcf, 0x90, 0x79, 0xf4, 0xf6, 0xf6, 0x37, 0x43, 0x41, 0x15, 0x50, 0x6f, 0x2f,
	0x9c, 0x81, 0xa1, 0xe2, 0x37, 0x12, 0x80, 0x97, 0xf3, 0x3a, 0x87, 0x8a, 0xa0, 0x4f, 0x68, 0x24,
	0x98, 0x15, 0x37, 0xca, 0x61, 0xa6, 0x93, 0x62, 0x92, 0x8d, 0x36, 0x15, 0xd8, 0x81, 0x5a, 0x31,
	0x83, 0x1a, 0x00, 0xc9, 0x46, 0x13, 0x9b, 0x10, 0xe3, 0x80, 0x2f, 0xe3, 0x8d, 0x2d, 0xcb, 0xf6,
	0xf8, 0x66, 0x0d, 0x80, 0xab, 0xae, 0xe7, 0xdd, 0xf4, 0x44, 0x52, 0x46, 0x87, 0xa0, 0x8f, 0xfa,
	0x03, 0x7b, 0x62, 0xff, 0x61, 0xd9, 0x76, 0xcf, 0x50, 0xf0, 0x09, 0x54, 0x2d, 0xc2, 0x7c, 0xea,
	0x2c, 0xa2, 0x18, 0x61, 0xa8, 0x45, 0xd4, 0xa7, 0xb3, 0x98, 0x47, 0xd2, 0xcf, 0xae, 0x92, 0xad,
	0x18, 0xfe, 0x5f, 0x81, 0xaa, 0x97, 0x55, 0x9c, 0x82, 0x16, 0x32, 0x5f, 0xf0, 0x50, 0x3a, 0xfa,
	0xd9, 0xd7, 0x39, 0x8f, 0x0d, 0xc4, 0x14, 0x75, 0xc1, 0x1d, 0x23, 0x09, 0x12, 0x7d, 0x0f, 0x8d,
	0x55, 0x48, 0x13, 0x0a, 0xc9, 0x18, 0x59, 0x8c, 0xd9, 0x89, 0xb6, 0xde, 0xca, 0x50, 0xd9, 0xd4,
	0xf2, 0xc7, 0x0d, 0x8b, 0x3b, 0xa5, 0x1e, 0x6a, 0x41, 0x85, 0x5b, 0xee, 0x74, 0x49, 0x53, 0x35,
	0x64, 0x3e, 0x97, 0x89, 0x4f, 0x1f, 0xa9, 0x9f, 0xea, 0x21, 0x71, 0x78, 0xa7, 0x68, 0x76, 0xcf,
	0x98, 0x2f, 0xc4, 0x50, 0x27, 0xa9, 0x87, 0x2e, 0x01, 0x66, 0x6c, 0xb9, 0x62, 0x01, 0x0d, 0xe2,
	0xa8, 0xa9, 0x09, 0x3a, 0x3f, 0xbc, 0x40, 0xc7, 0xb4, 0x32, 0xb4, 0x1d, 0xc4, 0xe1, 0x13, 0x29,
	0x94, 0xa3, 0x36, 0xe8, 0xd1, 0x3d, 0xfb, 0xe7, 0xf7, 0x69, 0x74, 0xbf, 0x60, 0x41, 0xb3, 0xd4,
	0x96, 0x3a, 0x15, 0x52, 0x0c, 0xb5, 0x7e, 0x85, 0xc3, 0x9d, 0x06, 0xc8, 0x00, 0xe5, 0x81, 0x3e,
	0x09, 0x82, 0x1a, 0xe1, 0x26, 0x67, 0xf0, 0x38, 0xf5, 0xd7, 0x09, 0x35, 0x8d, 0x24, 0xce, 0x2f,
	0xf2, 0xcf, 0x12, 0xbe, 0x04, 0xdd, 0xb2, 0x42, 0x3a, 0x8d, 0x29, 0xdf, 0x07, 0x21, 0x50, 0x03,
	0xfe, 0x09, 0x92, 0x3b, 0x11, 0x76, 0x81, 0xa8, 0xbc, 0x45, 0xd4, 0x00, 0x25, 0xa2, 0xff, 0xa6,
	0xec, 0xb9, 0x89, 0x3f, 0x48, 0xa0, 0x7b, 0x85, 0x6e, 0xa7, 0xa0, 0x84, 0x99, 0x34, 0xbf, 0xcd,
	0xbf, 0x41, 0x0e, 0x31, 0x09, 0x8d, 0xd6, 0x7e, 0x2c, 0x94, 0xc9, 0xb1, 0xe8, 0x04, 0xd4, 0x45,
	0x70, 0xc7, 0xc4, 0xa8, 0x4f, 0xc8, 0x40, 0x00, 0xf1, 0x2b, 0x09, 0x20, 0x6f, 0xf2, 0xb2, 0x88,
	0x0d, 0xa8, 0xb9, 0xdd, 0x81, 0x3d, 0xe9, 0xbb, 0xd7, 0x5d, 0x47, 0xdc, 0xd8, 0x11, 0x1c, 0x8a,
	0x48, 0x6f, 0x7c, 0xe5, 0xf4, 0xad, 0xee, 0x88, 0x4b, 0x39, 0x83, 0x0d, 0xaf, 0x6d, 0xe2, 0xd8,
	0xae, 0xa1, 0xa2, 0x2f, 0xa0, 0x2e, 0x22, 0xde, 0xc5, 0x90, 0x8c, 0x78, 0x48, 0xe3, 0x83, 0x78,
	0x7e, 0x62, 0x0d, 0xc7, 0xee, 0xc8, 0x28, 0xe1, 0xef, 0x40, 0xb7, 0xec, 0x20, 0xa6, 0xe1, 0x0d,
	0x0b, 0xfd, 0xf9, 0x3e, 0x9d, 0x71, 0x98, 0xf7, 0x19, 0xb0, 0x6f, 0xa0, 0x6a, 0x75, 0xe7, 0xf3,
	0x01, 0x0b, 0xa8, 0x78, 0xcf, 0x60, 0xbd, 0x14, 0x88, 0x3a, 0xe1, 0x26, 0xc6, 0x00, 0x9e, 0xc8,
	0x09, 0x4d, 0x1f, 0x83, 0xb6, 0xe4, 0x4e, 0x8a, 0x48, 0x9c, 0x4d, 0x0b, 0x47, 0x88, 0xf5, 0xf9,
	0x16, 0x22, 0xb7, 0x69, 0x91, 0x48, 0x5c, 0x2a, 0x48, 0x1c, 0x7f, 0x09, 0x47, 0x16, 0xa1, 0x7f,
	0x7b, 0x34, 0x7c, 0xa4, 0xe1, 0xe6, 0xbb, 0x47, 0xf8, 0xb5, 0x0c, 0x86, 0xb7, 0x13, 0xdc, 0x7b,
	0x58, 0x3f, 0x65, 0x0f, 0xca, 0x0f, 0x01, 0xe7, 0x0f, 0xba, 0xd3, 0x60, 0xe7, 0x5d, 0x5b, 0xef,
	0xa4, 0xed, 0xab, 0x4d, 0x7f, 0xc9, 0xd2, 0xd6, 0x2f, 0x39, 0x1f, 0x2a, 0xef, 0xbd, 0x66, 0x65,
	0xdf, 0x35, 0xab, 0xcf, 0x5f, 0xb3, 0xf6, 0x9c, 0xc8, 0x4b, 0x99, 0xc8, 0x79, 0x6f, 0x7f, 0x1a,
	0xc5, 0xf1, 0x62, 0x49, 0x9b, 0x65, 0x31, 0x35, 0xf3, 0xff, 0x2c, 0x09, 0x6e, 0x3f, 0x7e, 0x0c,
	0x00, 0x00, 0xff, 0xff, 0x5c, 0x8a, 0x71, 0xeb, 0x96, 0x06, 0x00, 0x00,
}