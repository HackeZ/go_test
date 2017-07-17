// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/common.proto

/*
Package common is a generated protocol buffer package.

It is generated from these files:
	common/common.proto

It has these top-level messages:
	Rslt
*/
package common

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

type Rslt struct {
	// *
	// 返回结果值
	Rslt int32 `protobuf:"varint,1,opt,name=rslt" json:"rslt,omitempty"`
	// *
	// 描述
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
}

func (m *Rslt) Reset()                    { *m = Rslt{} }
func (m *Rslt) String() string            { return proto.CompactTextString(m) }
func (*Rslt) ProtoMessage()               {}
func (*Rslt) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Rslt) GetRslt() int32 {
	if m != nil {
		return m.Rslt
	}
	return 0
}

func (m *Rslt) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*Rslt)(nil), "common.Rslt")
}

func init() { proto.RegisterFile("common/common.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 98 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x87, 0x50, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x6c, 0x10, 0x9e, 0x92,
	0x0d, 0x17, 0x4b, 0x50, 0x71, 0x4e, 0x89, 0x90, 0x10, 0x17, 0x4b, 0x51, 0x71, 0x4e, 0x89, 0x04,
	0xa3, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x98, 0x2d, 0xa4, 0xc0, 0xc5, 0x9d, 0x92, 0x5a, 0x9c, 0x5c,
	0x94, 0x59, 0x50, 0x92, 0x99, 0x9f, 0x27, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x19, 0x84, 0x2c, 0x94,
	0xc4, 0x06, 0x36, 0xcc, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x4b, 0x86, 0xe5, 0x63, 0x00,
	0x00, 0x00,
}
