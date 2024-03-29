// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/getImg/getImg.proto

package go_micro_srv_getImg

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Request struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_44bdb7cd3d3a0a3e, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type Response struct {
	Errno                string   `protobuf:"bytes,1,opt,name=Errno,proto3" json:"Errno,omitempty"`
	Errmsg               string   `protobuf:"bytes,2,opt,name=Errmsg,proto3" json:"Errmsg,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_44bdb7cd3d3a0a3e, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetErrno() string {
	if m != nil {
		return m.Errno
	}
	return ""
}

func (m *Response) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func (m *Response) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.getImg.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.getImg.Response")
}

func init() { proto.RegisterFile("proto/getImg/getImg.proto", fileDescriptor_44bdb7cd3d3a0a3e) }

var fileDescriptor_44bdb7cd3d3a0a3e = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x4f, 0x2d, 0xf1, 0xcc, 0x4d, 0x87, 0x52, 0x7a, 0x60, 0x31, 0x21, 0xe1, 0xf4,
	0x7c, 0xbd, 0xdc, 0xcc, 0xe4, 0xa2, 0x7c, 0xbd, 0xe2, 0xa2, 0x32, 0x3d, 0x88, 0x94, 0x92, 0x2c,
	0x17, 0x7b, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x10, 0x17, 0x4b, 0x69, 0x69, 0x66,
	0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0xad, 0xe4, 0xc3, 0xc5, 0x11, 0x94, 0x5a,
	0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x24, 0xc2, 0xc5, 0xea, 0x5a, 0x54, 0x94, 0x97, 0x0f, 0x55,
	0x00, 0xe1, 0x08, 0x89, 0x71, 0xb1, 0xb9, 0x16, 0x15, 0xe5, 0x16, 0xa7, 0x4b, 0x30, 0x81, 0x85,
	0xa1, 0x3c, 0x90, 0x69, 0x29, 0x89, 0x25, 0x89, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0x3c, 0x41, 0x60,
	0xb6, 0x51, 0x18, 0x17, 0x9b, 0x3b, 0xd8, 0x5a, 0x21, 0x1f, 0x2e, 0x6e, 0x5f, 0x90, 0x53, 0xa0,
	0x5c, 0x19, 0x3d, 0x2c, 0x6e, 0xd3, 0x83, 0x3a, 0x4c, 0x4a, 0x16, 0x87, 0x2c, 0xc4, 0x5d, 0x4a,
	0x0c, 0x49, 0x6c, 0x60, 0x0f, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc0, 0x74, 0x7b, 0x61,
	0xfd, 0x00, 0x00, 0x00,
}
