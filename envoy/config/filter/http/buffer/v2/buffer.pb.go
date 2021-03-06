// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/buffer/v2/buffer.proto

package envoy_config_filter_http_buffer_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type Buffer struct {
	MaxRequestBytes      *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=max_request_bytes,json=maxRequestBytes,proto3" json:"max_request_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Buffer) Reset()         { *m = Buffer{} }
func (m *Buffer) String() string { return proto.CompactTextString(m) }
func (*Buffer) ProtoMessage()    {}
func (*Buffer) Descriptor() ([]byte, []int) {
	return fileDescriptor_e782fc75ce4c789f, []int{0}
}

func (m *Buffer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Buffer.Unmarshal(m, b)
}
func (m *Buffer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Buffer.Marshal(b, m, deterministic)
}
func (m *Buffer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Buffer.Merge(m, src)
}
func (m *Buffer) XXX_Size() int {
	return xxx_messageInfo_Buffer.Size(m)
}
func (m *Buffer) XXX_DiscardUnknown() {
	xxx_messageInfo_Buffer.DiscardUnknown(m)
}

var xxx_messageInfo_Buffer proto.InternalMessageInfo

func (m *Buffer) GetMaxRequestBytes() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxRequestBytes
	}
	return nil
}

type BufferPerRoute struct {
	// Types that are valid to be assigned to Override:
	//	*BufferPerRoute_Disabled
	//	*BufferPerRoute_Buffer
	Override             isBufferPerRoute_Override `protobuf_oneof:"override"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *BufferPerRoute) Reset()         { *m = BufferPerRoute{} }
func (m *BufferPerRoute) String() string { return proto.CompactTextString(m) }
func (*BufferPerRoute) ProtoMessage()    {}
func (*BufferPerRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_e782fc75ce4c789f, []int{1}
}

func (m *BufferPerRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BufferPerRoute.Unmarshal(m, b)
}
func (m *BufferPerRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BufferPerRoute.Marshal(b, m, deterministic)
}
func (m *BufferPerRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BufferPerRoute.Merge(m, src)
}
func (m *BufferPerRoute) XXX_Size() int {
	return xxx_messageInfo_BufferPerRoute.Size(m)
}
func (m *BufferPerRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_BufferPerRoute.DiscardUnknown(m)
}

var xxx_messageInfo_BufferPerRoute proto.InternalMessageInfo

type isBufferPerRoute_Override interface {
	isBufferPerRoute_Override()
}

type BufferPerRoute_Disabled struct {
	Disabled bool `protobuf:"varint,1,opt,name=disabled,proto3,oneof"`
}

type BufferPerRoute_Buffer struct {
	Buffer *Buffer `protobuf:"bytes,2,opt,name=buffer,proto3,oneof"`
}

func (*BufferPerRoute_Disabled) isBufferPerRoute_Override() {}

func (*BufferPerRoute_Buffer) isBufferPerRoute_Override() {}

func (m *BufferPerRoute) GetOverride() isBufferPerRoute_Override {
	if m != nil {
		return m.Override
	}
	return nil
}

func (m *BufferPerRoute) GetDisabled() bool {
	if x, ok := m.GetOverride().(*BufferPerRoute_Disabled); ok {
		return x.Disabled
	}
	return false
}

func (m *BufferPerRoute) GetBuffer() *Buffer {
	if x, ok := m.GetOverride().(*BufferPerRoute_Buffer); ok {
		return x.Buffer
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*BufferPerRoute) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*BufferPerRoute_Disabled)(nil),
		(*BufferPerRoute_Buffer)(nil),
	}
}

func init() {
	proto.RegisterType((*Buffer)(nil), "envoy.config.filter.http.buffer.v2.Buffer")
	proto.RegisterType((*BufferPerRoute)(nil), "envoy.config.filter.http.buffer.v2.BufferPerRoute")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/buffer/v2/buffer.proto", fileDescriptor_e782fc75ce4c789f)
}

var fileDescriptor_e782fc75ce4c789f = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xc1, 0x8a, 0xd3, 0x40,
	0x1c, 0xc6, 0x3b, 0x69, 0x1b, 0xc3, 0x14, 0x6c, 0xcd, 0xc5, 0x52, 0xb4, 0x94, 0x82, 0xa8, 0x3d,
	0xcc, 0x48, 0xfa, 0x06, 0x73, 0xaa, 0x1e, 0xa4, 0x04, 0x14, 0x3c, 0x95, 0x89, 0xf9, 0x27, 0x8e,
	0xa4, 0x99, 0x38, 0x33, 0x89, 0xe9, 0xcd, 0xb3, 0x17, 0xaf, 0x7d, 0x16, 0x9f, 0xc0, 0xab, 0xaf,
	0xe2, 0x49, 0x7a, 0x58, 0x96, 0x64, 0x92, 0x85, 0x65, 0x0f, 0xbb, 0xb7, 0x81, 0x6f, 0xbe, 0xdf,
	0xff, 0xc7, 0x87, 0x29, 0xe4, 0x95, 0x3c, 0xd1, 0xcf, 0x32, 0x4f, 0x44, 0x4a, 0x13, 0x91, 0x19,
	0x50, 0xf4, 0x8b, 0x31, 0x05, 0x8d, 0xca, 0x24, 0x01, 0x45, 0xab, 0xa0, 0x7b, 0x91, 0x42, 0x49,
	0x23, 0xfd, 0x75, 0x5b, 0x20, 0xb6, 0x40, 0x6c, 0x81, 0x34, 0x05, 0xd2, 0x7d, 0xab, 0x82, 0xc5,
	0x32, 0x95, 0x32, 0xcd, 0x80, 0xb6, 0x8d, 0xa8, 0x4c, 0xe8, 0x77, 0xc5, 0x8b, 0x02, 0x94, 0xb6,
	0x8c, 0xc5, 0xb2, 0x8c, 0x0b, 0x4e, 0x79, 0x9e, 0x4b, 0xc3, 0x8d, 0x90, 0xb9, 0xa6, 0x47, 0x91,
	0x2a, 0x6e, 0xa0, 0xcb, 0x9f, 0xdf, 0xc9, 0xb5, 0xe1, 0xa6, 0xec, 0xeb, 0x4f, 0x2b, 0x9e, 0x89,
	0x98, 0x1b, 0xa0, 0xfd, 0xc3, 0x06, 0x6b, 0x81, 0x5d, 0xd6, 0x4a, 0xf8, 0x9f, 0xf0, 0x93, 0x23,
	0xaf, 0x0f, 0x0a, 0xbe, 0x95, 0xa0, 0xcd, 0x21, 0x3a, 0x19, 0xd0, 0x73, 0xb4, 0x42, 0xaf, 0x26,
	0xc1, 0x33, 0x62, 0xed, 0x48, 0x6f, 0x47, 0x3e, 0xbc, 0xcd, 0xcd, 0x36, 0xf8, 0xc8, 0xb3, 0x12,
	0xd8, 0xf4, 0xc2, 0x46, 0x1b, 0x67, 0x35, 0xb8, 0xb0, 0xf1, 0x4f, 0xe4, 0xcc, 0x50, 0x38, 0x3d,
	0xf2, 0x3a, 0xb4, 0x18, 0xd6, 0x50, 0xde, 0x8d, 0x3c, 0x67, 0x36, 0x5c, 0x9f, 0x11, 0x7e, 0x6c,
	0x6f, 0xed, 0x41, 0x85, 0xb2, 0x34, 0xe0, 0xbf, 0xc0, 0x5e, 0x2c, 0x34, 0x8f, 0x32, 0x88, 0xdb,
	0x53, 0x1e, 0x7b, 0x74, 0x61, 0xa3, 0xaf, 0x8e, 0x87, 0x76, 0x83, 0xf0, 0x26, 0xf2, 0xdf, 0x63,
	0xd7, 0x2e, 0x35, 0x77, 0x5a, 0x9f, 0x0d, 0xb9, 0x7f, 0x51, 0x62, 0x4f, 0x31, 0xaf, 0xd7, 0xda,
	0x0d, 0xc2, 0x8e, 0xc2, 0xa6, 0xd8, 0x93, 0x15, 0x28, 0x25, 0x62, 0xf0, 0x87, 0xff, 0x19, 0x62,
	0xf5, 0xbf, 0xf3, 0xd5, 0xaf, 0xf1, 0x6b, 0xff, 0xa5, 0xe5, 0x42, 0x6d, 0x20, 0xd7, 0xcd, 0x8a,
	0x1d, 0x5b, 0xdf, 0x86, 0x6f, 0x7f, 0xff, 0xf8, 0xf3, 0xd7, 0x75, 0x66, 0x08, 0xbf, 0x11, 0xd2,
	0xba, 0x14, 0x4a, 0xd6, 0xa7, 0x07, 0x68, 0xb1, 0x49, 0x37, 0x41, 0xb3, 0xe4, 0x1e, 0x45, 0x6e,
	0x3b, 0xe9, 0xf6, 0x3a, 0x00, 0x00, 0xff, 0xff, 0x53, 0x16, 0x8e, 0x67, 0x55, 0x02, 0x00, 0x00,
}
