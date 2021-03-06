// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/listener/original_src/v2alpha1/original_src.proto

package envoy_config_filter_listener_original_src_v2alpha1

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type OriginalSrc struct {
	BindPort             bool     `protobuf:"varint,1,opt,name=bind_port,json=bindPort,proto3" json:"bind_port,omitempty"`
	Mark                 uint32   `protobuf:"varint,2,opt,name=mark,proto3" json:"mark,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OriginalSrc) Reset()         { *m = OriginalSrc{} }
func (m *OriginalSrc) String() string { return proto.CompactTextString(m) }
func (*OriginalSrc) ProtoMessage()    {}
func (*OriginalSrc) Descriptor() ([]byte, []int) {
	return fileDescriptor_19e4bd40556a6972, []int{0}
}

func (m *OriginalSrc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OriginalSrc.Unmarshal(m, b)
}
func (m *OriginalSrc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OriginalSrc.Marshal(b, m, deterministic)
}
func (m *OriginalSrc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OriginalSrc.Merge(m, src)
}
func (m *OriginalSrc) XXX_Size() int {
	return xxx_messageInfo_OriginalSrc.Size(m)
}
func (m *OriginalSrc) XXX_DiscardUnknown() {
	xxx_messageInfo_OriginalSrc.DiscardUnknown(m)
}

var xxx_messageInfo_OriginalSrc proto.InternalMessageInfo

func (m *OriginalSrc) GetBindPort() bool {
	if m != nil {
		return m.BindPort
	}
	return false
}

func (m *OriginalSrc) GetMark() uint32 {
	if m != nil {
		return m.Mark
	}
	return 0
}

func init() {
	proto.RegisterType((*OriginalSrc)(nil), "envoy.config.filter.listener.original_src.v2alpha1.OriginalSrc")
}

func init() {
	proto.RegisterFile("envoy/config/filter/listener/original_src/v2alpha1/original_src.proto", fileDescriptor_19e4bd40556a6972)
}

var fileDescriptor_19e4bd40556a6972 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xd9, 0xa2, 0xa5, 0x46, 0x84, 0x92, 0x8b, 0xa5, 0xa2, 0x14, 0x4f, 0x3d, 0xed, 0xd2,
	0xe6, 0x2e, 0x52, 0xf0, 0x6c, 0xa8, 0x0f, 0x50, 0xa6, 0xc9, 0x36, 0x2e, 0xa6, 0x3b, 0x61, 0x76,
	0x1a, 0xda, 0x9b, 0x6f, 0xe0, 0x51, 0x9f, 0xc5, 0x27, 0xf0, 0xea, 0xab, 0xf8, 0x00, 0x22, 0xc9,
	0x26, 0x60, 0x11, 0x0f, 0xbd, 0x0d, 0xf3, 0xf3, 0x7f, 0xbb, 0xdf, 0x04, 0x77, 0xda, 0x96, 0xb8,
	0x53, 0x09, 0xda, 0x95, 0xc9, 0xd4, 0xca, 0xe4, 0xac, 0x49, 0xe5, 0xc6, 0xb1, 0xb6, 0x9a, 0x14,
	0x92, 0xc9, 0x8c, 0x85, 0x7c, 0xe1, 0x28, 0x51, 0xe5, 0x14, 0xf2, 0xe2, 0x11, 0x26, 0x7b, 0x5b,
	0x59, 0x10, 0x32, 0x86, 0xd3, 0x1a, 0x23, 0x3d, 0x46, 0x7a, 0x8c, 0x6c, 0x31, 0x72, 0xaf, 0xd0,
	0x62, 0x86, 0x57, 0x9b, 0xb4, 0x00, 0x05, 0xd6, 0x22, 0x03, 0x1b, 0xb4, 0x4e, 0xad, 0x4d, 0x46,
	0xc0, 0xda, 0x33, 0x87, 0x97, 0x7f, 0x72, 0xc7, 0xc0, 0x1b, 0xd7, 0xc4, 0xe7, 0x25, 0xe4, 0x26,
	0x05, 0xd6, 0xaa, 0x1d, 0x7c, 0x70, 0x7d, 0x13, 0x9c, 0xde, 0x37, 0x0f, 0x3e, 0x50, 0x12, 0x5e,
	0x04, 0x27, 0x4b, 0x63, 0xd3, 0x45, 0x81, 0xc4, 0x03, 0x31, 0x12, 0xe3, 0xde, 0xbc, 0x57, 0x2d,
	0x62, 0x24, 0x0e, 0xc3, 0xe0, 0x68, 0x0d, 0xf4, 0x34, 0xe8, 0x8c, 0xc4, 0xf8, 0x6c, 0x5e, 0xcf,
	0xb3, 0x57, 0xf1, 0xf5, 0xf6, 0xfd, 0x72, 0x1c, 0x85, 0x13, 0x2f, 0xa5, 0xb7, 0xac, 0xad, 0xab,
	0x3e, 0xd0, 0x88, 0xb9, 0xff, 0xcc, 0xa2, 0xf7, 0xe7, 0x8f, 0xcf, 0x6e, 0xa7, 0x2f, 0x82, 0x5b,
	0x83, 0xb2, 0x6e, 0x17, 0x84, 0xdb, 0x9d, 0x3c, 0xfc, 0x3a, 0xb3, 0xfe, 0x2f, 0x87, 0xb8, 0xf2,
	0x8a, 0xc5, 0xb2, 0x5b, 0x0b, 0x46, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x57, 0x2c, 0x7c, 0x08,
	0xb5, 0x01, 0x00, 0x00,
}
