// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/CentrexMsg.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	pb/CentrexMsg.proto

It has these top-level messages:
	CentrexMsg
	Login
	LoginRsp
*/
package pb

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

type CentrexMsg_MessageType int32

const (
	CentrexMsg_Login    CentrexMsg_MessageType = 0
	CentrexMsg_LoginRsp CentrexMsg_MessageType = 1
)

var CentrexMsg_MessageType_name = map[int32]string{
	0: "Login",
	1: "LoginRsp",
}
var CentrexMsg_MessageType_value = map[string]int32{
	"Login":    0,
	"LoginRsp": 1,
}

func (x CentrexMsg_MessageType) String() string {
	return proto.EnumName(CentrexMsg_MessageType_name, int32(x))
}
func (CentrexMsg_MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type CentrexMsg struct {
	Type     CentrexMsg_MessageType `protobuf:"varint,1,opt,name=type,enum=CentrexMsg_MessageType" json:"type,omitempty"`
	Login    *Login                 `protobuf:"bytes,2,opt,name=login" json:"login,omitempty"`
	LoginRsp *LoginRsp              `protobuf:"bytes,3,opt,name=loginRsp" json:"loginRsp,omitempty"`
}

func (m *CentrexMsg) Reset()                    { *m = CentrexMsg{} }
func (m *CentrexMsg) String() string            { return proto.CompactTextString(m) }
func (*CentrexMsg) ProtoMessage()               {}
func (*CentrexMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CentrexMsg) GetType() CentrexMsg_MessageType {
	if m != nil {
		return m.Type
	}
	return CentrexMsg_Login
}

func (m *CentrexMsg) GetLogin() *Login {
	if m != nil {
		return m.Login
	}
	return nil
}

func (m *CentrexMsg) GetLoginRsp() *LoginRsp {
	if m != nil {
		return m.LoginRsp
	}
	return nil
}

type Login struct {
	RequestId uint32 `protobuf:"varint,1,opt,name=requestId" json:"requestId,omitempty"`
	Username  string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Password  string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
}

func (m *Login) Reset()                    { *m = Login{} }
func (m *Login) String() string            { return proto.CompactTextString(m) }
func (*Login) ProtoMessage()               {}
func (*Login) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Login) GetRequestId() uint32 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

func (m *Login) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Login) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginRsp struct {
	RequestId uint32 `protobuf:"varint,1,opt,name=requestId" json:"requestId,omitempty"`
	ErrorCode int32  `protobuf:"varint,2,opt,name=errorCode" json:"errorCode,omitempty"`
	ErrorMsg  string `protobuf:"bytes,3,opt,name=errorMsg" json:"errorMsg,omitempty"`
}

func (m *LoginRsp) Reset()                    { *m = LoginRsp{} }
func (m *LoginRsp) String() string            { return proto.CompactTextString(m) }
func (*LoginRsp) ProtoMessage()               {}
func (*LoginRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LoginRsp) GetRequestId() uint32 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

func (m *LoginRsp) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *LoginRsp) GetErrorMsg() string {
	if m != nil {
		return m.ErrorMsg
	}
	return ""
}

func init() {
	proto.RegisterType((*CentrexMsg)(nil), "CentrexMsg")
	proto.RegisterType((*Login)(nil), "Login")
	proto.RegisterType((*LoginRsp)(nil), "LoginRsp")
	proto.RegisterEnum("CentrexMsg_MessageType", CentrexMsg_MessageType_name, CentrexMsg_MessageType_value)
}

func init() { proto.RegisterFile("pb/CentrexMsg.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x4d, 0x35, 0x25, 0x73, 0xeb, 0x4f, 0x19, 0x05, 0x83, 0x64, 0x51, 0x02, 0x4a, 0xa1,
	0x30, 0x42, 0x7d, 0x83, 0x76, 0x25, 0x18, 0x90, 0xc1, 0x95, 0xe0, 0x22, 0x3f, 0x97, 0x50, 0xa8,
	0x99, 0x71, 0xee, 0xd4, 0xda, 0x17, 0xf2, 0x39, 0x25, 0x93, 0xbf, 0xee, 0xdc, 0xcd, 0x39, 0xf7,
	0x63, 0xce, 0xe1, 0x5e, 0xb8, 0xd6, 0xd9, 0xe3, 0x1a, 0x2b, 0x6b, 0xf0, 0x27, 0xa1, 0x52, 0x68,
	0xa3, 0xac, 0x8a, 0x7f, 0x3d, 0x80, 0xc1, 0xe4, 0x0b, 0x38, 0xb3, 0x07, 0x8d, 0xa1, 0x37, 0xf3,
	0xe6, 0x97, 0xcb, 0x5b, 0x71, 0xc4, 0x27, 0x48, 0x94, 0x96, 0xf8, 0x76, 0xd0, 0x28, 0x1d, 0xc4,
	0x23, 0xf0, 0xb7, 0xaa, 0xdc, 0x54, 0xe1, 0x68, 0xe6, 0xcd, 0x27, 0xcb, 0xb1, 0x78, 0xa9, 0x95,
	0x6c, 0x4c, 0x7e, 0x0f, 0x81, 0x7b, 0x48, 0xd2, 0xe1, 0xa9, 0x03, 0x58, 0x0b, 0x90, 0x96, 0xfd,
	0x28, 0x7e, 0x80, 0xc9, 0xd1, 0xcf, 0x9c, 0x81, 0xef, 0xa0, 0xe9, 0x09, 0x3f, 0x87, 0xa0, 0xe3,
	0xa7, 0x5e, 0xfc, 0xd1, 0x0e, 0x78, 0x04, 0xcc, 0xe0, 0xd7, 0x0e, 0xc9, 0x3e, 0x17, 0xae, 0xe7,
	0x85, 0x1c, 0x0c, 0x7e, 0x07, 0xc1, 0x8e, 0xd0, 0x54, 0xe9, 0x27, 0xba, 0x5a, 0x4c, 0xf6, 0xba,
	0x9e, 0xe9, 0x94, 0x68, 0xaf, 0x4c, 0xe1, 0x1a, 0x31, 0xd9, 0xeb, 0x38, 0x1b, 0xc2, 0xfe, 0x49,
	0x88, 0x80, 0xa1, 0x31, 0xca, 0xac, 0x55, 0xd1, 0x44, 0xf8, 0x72, 0x30, 0xea, 0x0c, 0x27, 0x12,
	0x2a, 0xbb, 0x8c, 0x4e, 0xaf, 0x16, 0x70, 0x63, 0x95, 0x55, 0xdf, 0x1b, 0xdc, 0x8b, 0xbc, 0x59,
	0xac, 0x30, 0x3a, 0x5f, 0x5d, 0xb5, 0x5b, 0x7e, 0xad, 0x2f, 0x92, 0xab, 0xed, 0xfb, 0x48, 0x67,
	0xd9, 0xd8, 0xdd, 0xe7, 0xe9, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xb5, 0x7e, 0x96, 0xc2, 0xb6, 0x01,
	0x00, 0x00,
}
