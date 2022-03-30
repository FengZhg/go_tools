// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: go_protocol/protocol.proto

package go_protocol

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// 基础返回体 具体信息会在message里面
type StandardRsp struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StandardRsp) Reset()         { *m = StandardRsp{} }
func (m *StandardRsp) String() string { return proto.CompactTextString(m) }
func (*StandardRsp) ProtoMessage()    {}
func (*StandardRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_25134bbfdae23f9b, []int{0}
}
func (m *StandardRsp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StandardRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StandardRsp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StandardRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StandardRsp.Merge(m, src)
}
func (m *StandardRsp) XXX_Size() int {
	return m.Size()
}
func (m *StandardRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_StandardRsp.DiscardUnknown(m)
}

var xxx_messageInfo_StandardRsp proto.InternalMessageInfo

func (m *StandardRsp) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *StandardRsp) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *StandardRsp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// 登录态结构
type LoginStatus struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginStatus) Reset()         { *m = LoginStatus{} }
func (m *LoginStatus) String() string { return proto.CompactTextString(m) }
func (*LoginStatus) ProtoMessage()    {}
func (*LoginStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_25134bbfdae23f9b, []int{1}
}
func (m *LoginStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LoginStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LoginStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LoginStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginStatus.Merge(m, src)
}
func (m *LoginStatus) XXX_Size() int {
	return m.Size()
}
func (m *LoginStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginStatus.DiscardUnknown(m)
}

var xxx_messageInfo_LoginStatus proto.InternalMessageInfo

func (m *LoginStatus) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *LoginStatus) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LoginStatus) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type SingleLogInfo struct {
	LogType              int32    `protobuf:"varint,1,opt,name=logType,proto3" json:"logType,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string   `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	FullPath             string   `protobuf:"bytes,4,opt,name=fullPath,proto3" json:"fullPath,omitempty"`
	Status               string   `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Req                  string   `protobuf:"bytes,6,opt,name=req,proto3" json:"req,omitempty"`
	Message              string   `protobuf:"bytes,7,opt,name=message,proto3" json:"message,omitempty"`
	Time                 string   `protobuf:"bytes,8,opt,name=time,proto3" json:"time,omitempty"`
	TimeStamp            int64    `protobuf:"varint,9,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
	ImgUrl               string   `protobuf:"bytes,10,opt,name=imgUrl,proto3" json:"imgUrl,omitempty"`
	Reserve1             string   `protobuf:"bytes,11,opt,name=reserve1,proto3" json:"reserve1,omitempty"`
	Reserve2             string   `protobuf:"bytes,12,opt,name=reserve2,proto3" json:"reserve2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SingleLogInfo) Reset()         { *m = SingleLogInfo{} }
func (m *SingleLogInfo) String() string { return proto.CompactTextString(m) }
func (*SingleLogInfo) ProtoMessage()    {}
func (*SingleLogInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_25134bbfdae23f9b, []int{2}
}
func (m *SingleLogInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SingleLogInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SingleLogInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SingleLogInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SingleLogInfo.Merge(m, src)
}
func (m *SingleLogInfo) XXX_Size() int {
	return m.Size()
}
func (m *SingleLogInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SingleLogInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SingleLogInfo proto.InternalMessageInfo

func (m *SingleLogInfo) GetLogType() int32 {
	if m != nil {
		return m.LogType
	}
	return 0
}

func (m *SingleLogInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SingleLogInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SingleLogInfo) GetFullPath() string {
	if m != nil {
		return m.FullPath
	}
	return ""
}

func (m *SingleLogInfo) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *SingleLogInfo) GetReq() string {
	if m != nil {
		return m.Req
	}
	return ""
}

func (m *SingleLogInfo) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *SingleLogInfo) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func (m *SingleLogInfo) GetTimeStamp() int64 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

func (m *SingleLogInfo) GetImgUrl() string {
	if m != nil {
		return m.ImgUrl
	}
	return ""
}

func (m *SingleLogInfo) GetReserve1() string {
	if m != nil {
		return m.Reserve1
	}
	return ""
}

func (m *SingleLogInfo) GetReserve2() string {
	if m != nil {
		return m.Reserve2
	}
	return ""
}

func init() {
	proto.RegisterType((*StandardRsp)(nil), "go_protocol.StandardRsp")
	proto.RegisterType((*LoginStatus)(nil), "go_protocol.LoginStatus")
	proto.RegisterType((*SingleLogInfo)(nil), "go_protocol.SingleLogInfo")
}

func init() { proto.RegisterFile("go_protocol/protocol.proto", fileDescriptor_25134bbfdae23f9b) }

var fileDescriptor_25134bbfdae23f9b = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x4d, 0x4a, 0xfb, 0x40,
	0x1c, 0xfd, 0x4f, 0xd2, 0xaf, 0xfc, 0xf2, 0xaf, 0x94, 0x41, 0x64, 0x28, 0x12, 0x4b, 0x57, 0x5d,
	0xb5, 0x58, 0x6f, 0xe0, 0x46, 0x84, 0x2e, 0x24, 0xb1, 0x1b, 0x37, 0x32, 0x9a, 0xe9, 0x38, 0x90,
	0x64, 0xd2, 0xc9, 0x54, 0xf0, 0x26, 0xde, 0xc2, 0x6b, 0xb8, 0xf4, 0x08, 0x52, 0x2f, 0x22, 0x33,
	0x93, 0x68, 0x0a, 0xae, 0xf2, 0x3e, 0x92, 0xf7, 0x7b, 0x8f, 0xc0, 0x98, 0xcb, 0xfb, 0x52, 0x49,
	0x2d, 0x1f, 0x65, 0xb6, 0x68, 0xc0, 0xdc, 0x02, 0x1c, 0xb6, 0xbc, 0xe9, 0x1a, 0xc2, 0x44, 0xd3,
	0x22, 0xa5, 0x2a, 0x8d, 0xab, 0x12, 0x9f, 0x40, 0xaf, 0xd2, 0x54, 0xef, 0x2a, 0x82, 0x26, 0x68,
	0xd6, 0x8d, 0x6b, 0x86, 0x8f, 0xa1, 0xcb, 0x94, 0x92, 0x8a, 0x78, 0x13, 0x34, 0x0b, 0x62, 0x47,
	0x30, 0x81, 0x7e, 0xce, 0xaa, 0x8a, 0x72, 0x46, 0x7c, 0xab, 0x37, 0x74, 0x7a, 0x05, 0xe1, 0x4a,
	0x72, 0x51, 0x24, 0xee, 0xf3, 0x11, 0xf8, 0x3b, 0x91, 0xda, 0xcc, 0x20, 0x36, 0x10, 0x63, 0xe8,
	0x14, 0x34, 0x67, 0x75, 0x9e, 0xc5, 0x46, 0xd3, 0x2f, 0x65, 0x93, 0x65, 0xf1, 0xf4, 0xcd, 0x83,
	0x61, 0x22, 0x0a, 0x9e, 0xb1, 0x95, 0xe4, 0xd7, 0xc5, 0x46, 0x9a, 0xa3, 0x99, 0xe4, 0xb7, 0xe6,
	0x45, 0xd7, 0xb1, 0xa1, 0x7f, 0x66, 0x1e, 0x81, 0x27, 0xd2, 0x3a, 0xd1, 0x13, 0x29, 0x1e, 0xc3,
	0x60, 0xb3, 0xcb, 0xb2, 0x1b, 0xaa, 0x9f, 0x48, 0xc7, 0xaa, 0x3f, 0xbc, 0x35, 0xbe, 0x6b, 0x9d,
	0x66, 0xfc, 0x08, 0x7c, 0xc5, 0xb6, 0xa4, 0xe7, 0xda, 0x2b, 0xb6, 0x6d, 0x0f, 0xef, 0x1f, 0x0c,
	0xb7, 0x1b, 0x44, 0xce, 0xc8, 0xa0, 0xde, 0x20, 0x72, 0x86, 0x4f, 0x21, 0x30, 0xcf, 0x44, 0xd3,
	0xbc, 0x24, 0xc1, 0x04, 0xcd, 0xfc, 0xf8, 0x57, 0x30, 0x57, 0x45, 0xce, 0xd7, 0x2a, 0x23, 0xe0,
	0xae, 0x3a, 0x66, 0x9a, 0x2a, 0x56, 0x31, 0xf5, 0xcc, 0xce, 0x49, 0xe8, 0x9a, 0x36, 0xbc, 0xe5,
	0x2d, 0xc9, 0xff, 0x03, 0x6f, 0x79, 0x79, 0xf6, 0xbe, 0x8f, 0xd0, 0xc7, 0x3e, 0x42, 0x9f, 0xfb,
	0x08, 0xbd, 0x7e, 0x45, 0xff, 0xee, 0x86, 0xf3, 0x45, 0xeb, 0x97, 0x3f, 0xf4, 0x2c, 0xba, 0xf8,
	0x0e, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x54, 0xa0, 0x2d, 0x24, 0x02, 0x00, 0x00,
}

func (m *StandardRsp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StandardRsp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StandardRsp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintProtocol(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Error) > 0 {
		i -= len(m.Error)
		copy(dAtA[i:], m.Error)
		i = encodeVarintProtocol(dAtA, i, uint64(len(m.Error)))
		i--
		dAtA[i] = 0x12
	}
	if m.Status != 0 {
		i = encodeVarintProtocol(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *LoginStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoginStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LoginStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintProtocol(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintProtocol(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Uid) > 0 {
		i -= len(m.Uid)
		copy(dAtA[i:], m.Uid)
		i = encodeVarintProtocol(dAtA, i, uint64(len(m.Uid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SingleLogInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SingleLogInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SingleLogInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Reserve2) > 0 {
		i -= len(m.Reserve2)
		copy(dAtA[i:], m.Reserve2)
		i = encodeVarintProtocol(dAtA, i, uint64(len(m.Reserve2)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.Reserve1) > 0 {
		i -= len(m.Reserve1)
		copy(dAtA[i:], m.Reserve1)
		i = encodeVarintProtocol(dAtA, i, uint64(len(m.Reserve1)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.ImgUrl) > 0 {
		i -= len(m.ImgUrl)
		copy(dAtA[i:], m.ImgUrl)
		i = encodeVarintProtocol(dAtA, i, uint64(len(m.ImgUrl)))
		i--
		dAtA[i] = 0x52
	}
	if m.TimeStamp != 0 {
		i = encodeVarintProtocol(dAtA, i, uint64(m.TimeStamp))
		i--
		dAtA[i] = 0x48
	}
	if len(m.Time) > 0 {
		i -= len(m.Time)
		copy(dAtA[i:], m.Time)
		i = encodeVarintProtocol(dAtA, i, uint64(len(m.Time)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintProtocol(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Req) > 0 {
		i -= len(m.Req)
		copy(dAtA[i:], m.Req)
		i = encodeVarintProtocol(dAtA, i, uint64(len(m.Req)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintProtocol(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.FullPath) > 0 {
		i -= len(m.FullPath)
		copy(dAtA[i:], m.FullPath)
		i = encodeVarintProtocol(dAtA, i, uint64(len(m.FullPath)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintProtocol(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintProtocol(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.LogType != 0 {
		i = encodeVarintProtocol(dAtA, i, uint64(m.LogType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintProtocol(dAtA []byte, offset int, v uint64) int {
	offset -= sovProtocol(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StandardRsp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovProtocol(uint64(m.Status))
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovProtocol(uint64(l))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovProtocol(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LoginStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Uid)
	if l > 0 {
		n += 1 + l + sovProtocol(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovProtocol(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovProtocol(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SingleLogInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LogType != 0 {
		n += 1 + sovProtocol(uint64(m.LogType))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovProtocol(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovProtocol(uint64(l))
	}
	l = len(m.FullPath)
	if l > 0 {
		n += 1 + l + sovProtocol(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovProtocol(uint64(l))
	}
	l = len(m.Req)
	if l > 0 {
		n += 1 + l + sovProtocol(uint64(l))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovProtocol(uint64(l))
	}
	l = len(m.Time)
	if l > 0 {
		n += 1 + l + sovProtocol(uint64(l))
	}
	if m.TimeStamp != 0 {
		n += 1 + sovProtocol(uint64(m.TimeStamp))
	}
	l = len(m.ImgUrl)
	if l > 0 {
		n += 1 + l + sovProtocol(uint64(l))
	}
	l = len(m.Reserve1)
	if l > 0 {
		n += 1 + l + sovProtocol(uint64(l))
	}
	l = len(m.Reserve2)
	if l > 0 {
		n += 1 + l + sovProtocol(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovProtocol(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProtocol(x uint64) (n int) {
	return sovProtocol(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StandardRsp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtocol
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StandardRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StandardRsp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtocol
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtocol
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProtocol(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProtocol
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LoginStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtocol
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LoginStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LoginStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtocol
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtocol
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtocol
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProtocol(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProtocol
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SingleLogInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtocol
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SingleLogInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SingleLogInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogType", wireType)
			}
			m.LogType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LogType |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtocol
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtocol
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FullPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtocol
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FullPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtocol
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Req", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtocol
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Req = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtocol
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtocol
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Time = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeStamp", wireType)
			}
			m.TimeStamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeStamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImgUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtocol
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ImgUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reserve1", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtocol
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reserve1 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reserve2", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtocol
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reserve2 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProtocol(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProtocol
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipProtocol(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProtocol
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthProtocol
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProtocol
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProtocol
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProtocol        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProtocol          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProtocol = fmt.Errorf("proto: unexpected end of group")
)
