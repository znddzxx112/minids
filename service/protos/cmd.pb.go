// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cmd.proto

package protos

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

type Cmd struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ArgInfo              []byte   `protobuf:"bytes,2,opt,name=argInfo,proto3" json:"argInfo,omitempty"`
	ResInfo              []byte   `protobuf:"bytes,3,opt,name=resInfo,proto3" json:"resInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cmd) Reset()         { *m = Cmd{} }
func (m *Cmd) String() string { return proto.CompactTextString(m) }
func (*Cmd) ProtoMessage()    {}
func (*Cmd) Descriptor() ([]byte, []int) {
	return fileDescriptor_7520252fb01eaf30, []int{0}
}

func (m *Cmd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cmd.Unmarshal(m, b)
}
func (m *Cmd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cmd.Marshal(b, m, deterministic)
}
func (m *Cmd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cmd.Merge(m, src)
}
func (m *Cmd) XXX_Size() int {
	return xxx_messageInfo_Cmd.Size(m)
}
func (m *Cmd) XXX_DiscardUnknown() {
	xxx_messageInfo_Cmd.DiscardUnknown(m)
}

var xxx_messageInfo_Cmd proto.InternalMessageInfo

func (m *Cmd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Cmd) GetArgInfo() []byte {
	if m != nil {
		return m.ArgInfo
	}
	return nil
}

func (m *Cmd) GetResInfo() []byte {
	if m != nil {
		return m.ResInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*Cmd)(nil), "protos.cmd")
}

func init() { proto.RegisterFile("cmd.proto", fileDescriptor_7520252fb01eaf30) }

var fileDescriptor_7520252fb01eaf30 = []byte{
	// 99 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0xce, 0x4d, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53, 0xc5, 0x4a, 0xbe, 0x5c, 0xcc, 0xc9, 0xb9,
	0x29, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41,
	0x60, 0xb6, 0x90, 0x04, 0x17, 0x7b, 0x62, 0x51, 0xba, 0x67, 0x5e, 0x5a, 0xbe, 0x04, 0x93, 0x02,
	0xa3, 0x06, 0x4f, 0x10, 0x8c, 0x0b, 0x92, 0x29, 0x4a, 0x2d, 0x06, 0xcb, 0x30, 0x43, 0x64, 0xa0,
	0xdc, 0x24, 0x88, 0xb1, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9a, 0xea, 0xc3, 0x92, 0x6a,
	0x00, 0x00, 0x00,
}