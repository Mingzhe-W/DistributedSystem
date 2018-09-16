// Code generated by protoc-gen-go. DO NOT EDIT.
// source: log_query.proto

package log_query

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// [START messages]
type ClientQuery struct {
	ClientIP             string   `protobuf:"bytes,1,opt,name=clientIP,proto3" json:"clientIP,omitempty"`
	ServerIP             string   `protobuf:"bytes,2,opt,name=serverIP,proto3" json:"serverIP,omitempty"`
	Pattern              string   `protobuf:"bytes,3,opt,name=pattern,proto3" json:"pattern,omitempty"`
	Regex                string   `protobuf:"bytes,4,opt,name=regex,proto3" json:"regex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientQuery) Reset()         { *m = ClientQuery{} }
func (m *ClientQuery) String() string { return proto.CompactTextString(m) }
func (*ClientQuery) ProtoMessage()    {}
func (*ClientQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_02e0a1acc6a3e478, []int{0}
}
func (m *ClientQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientQuery.Unmarshal(m, b)
}
func (m *ClientQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientQuery.Marshal(b, m, deterministic)
}
func (m *ClientQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientQuery.Merge(m, src)
}
func (m *ClientQuery) XXX_Size() int {
	return xxx_messageInfo_ClientQuery.Size(m)
}
func (m *ClientQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ClientQuery proto.InternalMessageInfo

func (m *ClientQuery) GetClientIP() string {
	if m != nil {
		return m.ClientIP
	}
	return ""
}

func (m *ClientQuery) GetServerIP() string {
	if m != nil {
		return m.ServerIP
	}
	return ""
}

func (m *ClientQuery) GetPattern() string {
	if m != nil {
		return m.Pattern
	}
	return ""
}

func (m *ClientQuery) GetRegex() string {
	if m != nil {
		return m.Regex
	}
	return ""
}

type ServerOneMatched struct {
	Line                 string   `protobuf:"bytes,1,opt,name=line,proto3" json:"line,omitempty"`
	MatchedPattern       string   `protobuf:"bytes,2,opt,name=matchedPattern,proto3" json:"matchedPattern,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerOneMatched) Reset()         { *m = ServerOneMatched{} }
func (m *ServerOneMatched) String() string { return proto.CompactTextString(m) }
func (*ServerOneMatched) ProtoMessage()    {}
func (*ServerOneMatched) Descriptor() ([]byte, []int) {
	return fileDescriptor_02e0a1acc6a3e478, []int{1}
}
func (m *ServerOneMatched) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerOneMatched.Unmarshal(m, b)
}
func (m *ServerOneMatched) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerOneMatched.Marshal(b, m, deterministic)
}
func (m *ServerOneMatched) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerOneMatched.Merge(m, src)
}
func (m *ServerOneMatched) XXX_Size() int {
	return xxx_messageInfo_ServerOneMatched.Size(m)
}
func (m *ServerOneMatched) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerOneMatched.DiscardUnknown(m)
}

var xxx_messageInfo_ServerOneMatched proto.InternalMessageInfo

func (m *ServerOneMatched) GetLine() string {
	if m != nil {
		return m.Line
	}
	return ""
}

func (m *ServerOneMatched) GetMatchedPattern() string {
	if m != nil {
		return m.MatchedPattern
	}
	return ""
}

// Our address book file is just one of these.
type ServerResponse struct {
	ClientIP             string              `protobuf:"bytes,1,opt,name=clientIP,proto3" json:"clientIP,omitempty"`
	ServerIP             string              `protobuf:"bytes,2,opt,name=serverIP,proto3" json:"serverIP,omitempty"`
	Answers              []*ServerOneMatched `protobuf:"bytes,3,rep,name=answers,proto3" json:"answers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ServerResponse) Reset()         { *m = ServerResponse{} }
func (m *ServerResponse) String() string { return proto.CompactTextString(m) }
func (*ServerResponse) ProtoMessage()    {}
func (*ServerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_02e0a1acc6a3e478, []int{2}
}
func (m *ServerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerResponse.Unmarshal(m, b)
}
func (m *ServerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerResponse.Marshal(b, m, deterministic)
}
func (m *ServerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerResponse.Merge(m, src)
}
func (m *ServerResponse) XXX_Size() int {
	return xxx_messageInfo_ServerResponse.Size(m)
}
func (m *ServerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServerResponse proto.InternalMessageInfo

func (m *ServerResponse) GetClientIP() string {
	if m != nil {
		return m.ClientIP
	}
	return ""
}

func (m *ServerResponse) GetServerIP() string {
	if m != nil {
		return m.ServerIP
	}
	return ""
}

func (m *ServerResponse) GetAnswers() []*ServerOneMatched {
	if m != nil {
		return m.Answers
	}
	return nil
}

func init() {
	proto.RegisterType((*ClientQuery)(nil), "log_query.ClientQuery")
	proto.RegisterType((*ServerOneMatched)(nil), "log_query.ServerOneMatched")
	proto.RegisterType((*ServerResponse)(nil), "log_query.ServerResponse")
}

func init() { proto.RegisterFile("log_query.proto", fileDescriptor_02e0a1acc6a3e478) }

var fileDescriptor_02e0a1acc6a3e478 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x90, 0xcf, 0x4e, 0x84, 0x30,
	0x10, 0x87, 0xb3, 0xb2, 0xba, 0xee, 0x6c, 0xb2, 0x9a, 0xc6, 0x43, 0xb3, 0x1e, 0xdc, 0x70, 0x30,
	0x9c, 0x20, 0xd1, 0xf8, 0x04, 0x9e, 0x38, 0xa8, 0x88, 0x0f, 0x60, 0x0a, 0x8e, 0x95, 0xa4, 0xb4,
	0xb5, 0x2d, 0xfe, 0x39, 0xf9, 0xea, 0x86, 0x16, 0x30, 0xe1, 0xe8, 0xad, 0xdf, 0x7c, 0xbf, 0xce,
	0xb4, 0x03, 0x27, 0x42, 0xf1, 0xe7, 0xf7, 0x0e, 0xcd, 0x77, 0xaa, 0x8d, 0x72, 0x8a, 0xac, 0xa7,
	0xc2, 0xee, 0x82, 0x2b, 0xc5, 0x05, 0x66, 0x5e, 0x54, 0xdd, 0x6b, 0xe6, 0x9a, 0x16, 0xad, 0x63,
	0xad, 0x0e, 0xd9, 0xb8, 0x83, 0xcd, 0xad, 0x68, 0x50, 0xba, 0xc7, 0x3e, 0x4f, 0x76, 0x70, 0x5c,
	0x7b, 0xcc, 0x0b, 0xba, 0xd8, 0x2f, 0x92, 0x75, 0x39, 0x71, 0xef, 0x2c, 0x9a, 0x0f, 0x34, 0x79,
	0x41, 0x0f, 0x82, 0x1b, 0x99, 0x50, 0x58, 0x69, 0xe6, 0x1c, 0x1a, 0x49, 0x23, 0xaf, 0x46, 0x24,
	0x67, 0x70, 0x68, 0x90, 0xe3, 0x17, 0x5d, 0xfa, 0x7a, 0x80, 0xf8, 0x1e, 0x4e, 0x9f, 0xfc, 0xdd,
	0x07, 0x89, 0x77, 0xcc, 0xd5, 0x6f, 0xf8, 0x42, 0x08, 0x2c, 0x45, 0x23, 0x71, 0x98, 0xeb, 0xcf,
	0xe4, 0x12, 0xb6, 0x6d, 0xd0, 0xc5, 0xd0, 0x3e, 0x4c, 0x9e, 0x55, 0xe3, 0x1f, 0xd8, 0x86, 0x7e,
	0x25, 0x5a, 0xad, 0xa4, 0xc5, 0x7f, 0xff, 0xe4, 0x06, 0x56, 0x4c, 0xda, 0x4f, 0x34, 0x96, 0x46,
	0xfb, 0x28, 0xd9, 0x5c, 0x9d, 0xa7, 0x7f, 0xfb, 0x9d, 0xbf, 0xb9, 0x1c, 0xb3, 0xd5, 0x91, 0x5f,
	0xe7, 0xf5, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xef, 0xfe, 0xe9, 0xca, 0x8d, 0x01, 0x00, 0x00,
}