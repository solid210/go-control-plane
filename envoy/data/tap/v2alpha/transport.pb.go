// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/data/tap/v2alpha/transport.proto

package envoy_data_tap_v2alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Connection struct {
	LocalAddress         *core.Address `protobuf:"bytes,2,opt,name=local_address,json=localAddress,proto3" json:"local_address,omitempty"`
	RemoteAddress        *core.Address `protobuf:"bytes,3,opt,name=remote_address,json=remoteAddress,proto3" json:"remote_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Connection) Reset()         { *m = Connection{} }
func (m *Connection) String() string { return proto.CompactTextString(m) }
func (*Connection) ProtoMessage()    {}
func (*Connection) Descriptor() ([]byte, []int) {
	return fileDescriptor_03a9cebdb27ee552, []int{0}
}

func (m *Connection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Connection.Unmarshal(m, b)
}
func (m *Connection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Connection.Marshal(b, m, deterministic)
}
func (m *Connection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Connection.Merge(m, src)
}
func (m *Connection) XXX_Size() int {
	return xxx_messageInfo_Connection.Size(m)
}
func (m *Connection) XXX_DiscardUnknown() {
	xxx_messageInfo_Connection.DiscardUnknown(m)
}

var xxx_messageInfo_Connection proto.InternalMessageInfo

func (m *Connection) GetLocalAddress() *core.Address {
	if m != nil {
		return m.LocalAddress
	}
	return nil
}

func (m *Connection) GetRemoteAddress() *core.Address {
	if m != nil {
		return m.RemoteAddress
	}
	return nil
}

type SocketEvent struct {
	Timestamp *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Types that are valid to be assigned to EventSelector:
	//	*SocketEvent_Read_
	//	*SocketEvent_Write_
	//	*SocketEvent_Closed_
	EventSelector        isSocketEvent_EventSelector `protobuf_oneof:"event_selector"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *SocketEvent) Reset()         { *m = SocketEvent{} }
func (m *SocketEvent) String() string { return proto.CompactTextString(m) }
func (*SocketEvent) ProtoMessage()    {}
func (*SocketEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_03a9cebdb27ee552, []int{1}
}

func (m *SocketEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SocketEvent.Unmarshal(m, b)
}
func (m *SocketEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SocketEvent.Marshal(b, m, deterministic)
}
func (m *SocketEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SocketEvent.Merge(m, src)
}
func (m *SocketEvent) XXX_Size() int {
	return xxx_messageInfo_SocketEvent.Size(m)
}
func (m *SocketEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_SocketEvent.DiscardUnknown(m)
}

var xxx_messageInfo_SocketEvent proto.InternalMessageInfo

func (m *SocketEvent) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type isSocketEvent_EventSelector interface {
	isSocketEvent_EventSelector()
}

type SocketEvent_Read_ struct {
	Read *SocketEvent_Read `protobuf:"bytes,2,opt,name=read,proto3,oneof"`
}

type SocketEvent_Write_ struct {
	Write *SocketEvent_Write `protobuf:"bytes,3,opt,name=write,proto3,oneof"`
}

type SocketEvent_Closed_ struct {
	Closed *SocketEvent_Closed `protobuf:"bytes,4,opt,name=closed,proto3,oneof"`
}

func (*SocketEvent_Read_) isSocketEvent_EventSelector() {}

func (*SocketEvent_Write_) isSocketEvent_EventSelector() {}

func (*SocketEvent_Closed_) isSocketEvent_EventSelector() {}

func (m *SocketEvent) GetEventSelector() isSocketEvent_EventSelector {
	if m != nil {
		return m.EventSelector
	}
	return nil
}

func (m *SocketEvent) GetRead() *SocketEvent_Read {
	if x, ok := m.GetEventSelector().(*SocketEvent_Read_); ok {
		return x.Read
	}
	return nil
}

func (m *SocketEvent) GetWrite() *SocketEvent_Write {
	if x, ok := m.GetEventSelector().(*SocketEvent_Write_); ok {
		return x.Write
	}
	return nil
}

func (m *SocketEvent) GetClosed() *SocketEvent_Closed {
	if x, ok := m.GetEventSelector().(*SocketEvent_Closed_); ok {
		return x.Closed
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SocketEvent) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SocketEvent_Read_)(nil),
		(*SocketEvent_Write_)(nil),
		(*SocketEvent_Closed_)(nil),
	}
}

type SocketEvent_Read struct {
	Data                 *Body    `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SocketEvent_Read) Reset()         { *m = SocketEvent_Read{} }
func (m *SocketEvent_Read) String() string { return proto.CompactTextString(m) }
func (*SocketEvent_Read) ProtoMessage()    {}
func (*SocketEvent_Read) Descriptor() ([]byte, []int) {
	return fileDescriptor_03a9cebdb27ee552, []int{1, 0}
}

func (m *SocketEvent_Read) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SocketEvent_Read.Unmarshal(m, b)
}
func (m *SocketEvent_Read) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SocketEvent_Read.Marshal(b, m, deterministic)
}
func (m *SocketEvent_Read) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SocketEvent_Read.Merge(m, src)
}
func (m *SocketEvent_Read) XXX_Size() int {
	return xxx_messageInfo_SocketEvent_Read.Size(m)
}
func (m *SocketEvent_Read) XXX_DiscardUnknown() {
	xxx_messageInfo_SocketEvent_Read.DiscardUnknown(m)
}

var xxx_messageInfo_SocketEvent_Read proto.InternalMessageInfo

func (m *SocketEvent_Read) GetData() *Body {
	if m != nil {
		return m.Data
	}
	return nil
}

type SocketEvent_Write struct {
	Data                 *Body    `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	EndStream            bool     `protobuf:"varint,2,opt,name=end_stream,json=endStream,proto3" json:"end_stream,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SocketEvent_Write) Reset()         { *m = SocketEvent_Write{} }
func (m *SocketEvent_Write) String() string { return proto.CompactTextString(m) }
func (*SocketEvent_Write) ProtoMessage()    {}
func (*SocketEvent_Write) Descriptor() ([]byte, []int) {
	return fileDescriptor_03a9cebdb27ee552, []int{1, 1}
}

func (m *SocketEvent_Write) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SocketEvent_Write.Unmarshal(m, b)
}
func (m *SocketEvent_Write) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SocketEvent_Write.Marshal(b, m, deterministic)
}
func (m *SocketEvent_Write) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SocketEvent_Write.Merge(m, src)
}
func (m *SocketEvent_Write) XXX_Size() int {
	return xxx_messageInfo_SocketEvent_Write.Size(m)
}
func (m *SocketEvent_Write) XXX_DiscardUnknown() {
	xxx_messageInfo_SocketEvent_Write.DiscardUnknown(m)
}

var xxx_messageInfo_SocketEvent_Write proto.InternalMessageInfo

func (m *SocketEvent_Write) GetData() *Body {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *SocketEvent_Write) GetEndStream() bool {
	if m != nil {
		return m.EndStream
	}
	return false
}

type SocketEvent_Closed struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SocketEvent_Closed) Reset()         { *m = SocketEvent_Closed{} }
func (m *SocketEvent_Closed) String() string { return proto.CompactTextString(m) }
func (*SocketEvent_Closed) ProtoMessage()    {}
func (*SocketEvent_Closed) Descriptor() ([]byte, []int) {
	return fileDescriptor_03a9cebdb27ee552, []int{1, 2}
}

func (m *SocketEvent_Closed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SocketEvent_Closed.Unmarshal(m, b)
}
func (m *SocketEvent_Closed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SocketEvent_Closed.Marshal(b, m, deterministic)
}
func (m *SocketEvent_Closed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SocketEvent_Closed.Merge(m, src)
}
func (m *SocketEvent_Closed) XXX_Size() int {
	return xxx_messageInfo_SocketEvent_Closed.Size(m)
}
func (m *SocketEvent_Closed) XXX_DiscardUnknown() {
	xxx_messageInfo_SocketEvent_Closed.DiscardUnknown(m)
}

var xxx_messageInfo_SocketEvent_Closed proto.InternalMessageInfo

type SocketBufferedTrace struct {
	TraceId              uint64         `protobuf:"varint,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	Connection           *Connection    `protobuf:"bytes,2,opt,name=connection,proto3" json:"connection,omitempty"`
	Events               []*SocketEvent `protobuf:"bytes,3,rep,name=events,proto3" json:"events,omitempty"`
	ReadTruncated        bool           `protobuf:"varint,4,opt,name=read_truncated,json=readTruncated,proto3" json:"read_truncated,omitempty"`
	WriteTruncated       bool           `protobuf:"varint,5,opt,name=write_truncated,json=writeTruncated,proto3" json:"write_truncated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SocketBufferedTrace) Reset()         { *m = SocketBufferedTrace{} }
func (m *SocketBufferedTrace) String() string { return proto.CompactTextString(m) }
func (*SocketBufferedTrace) ProtoMessage()    {}
func (*SocketBufferedTrace) Descriptor() ([]byte, []int) {
	return fileDescriptor_03a9cebdb27ee552, []int{2}
}

func (m *SocketBufferedTrace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SocketBufferedTrace.Unmarshal(m, b)
}
func (m *SocketBufferedTrace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SocketBufferedTrace.Marshal(b, m, deterministic)
}
func (m *SocketBufferedTrace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SocketBufferedTrace.Merge(m, src)
}
func (m *SocketBufferedTrace) XXX_Size() int {
	return xxx_messageInfo_SocketBufferedTrace.Size(m)
}
func (m *SocketBufferedTrace) XXX_DiscardUnknown() {
	xxx_messageInfo_SocketBufferedTrace.DiscardUnknown(m)
}

var xxx_messageInfo_SocketBufferedTrace proto.InternalMessageInfo

func (m *SocketBufferedTrace) GetTraceId() uint64 {
	if m != nil {
		return m.TraceId
	}
	return 0
}

func (m *SocketBufferedTrace) GetConnection() *Connection {
	if m != nil {
		return m.Connection
	}
	return nil
}

func (m *SocketBufferedTrace) GetEvents() []*SocketEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *SocketBufferedTrace) GetReadTruncated() bool {
	if m != nil {
		return m.ReadTruncated
	}
	return false
}

func (m *SocketBufferedTrace) GetWriteTruncated() bool {
	if m != nil {
		return m.WriteTruncated
	}
	return false
}

type SocketStreamedTraceSegment struct {
	TraceId uint64 `protobuf:"varint,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	// Types that are valid to be assigned to MessagePiece:
	//	*SocketStreamedTraceSegment_Connection
	//	*SocketStreamedTraceSegment_Event
	MessagePiece         isSocketStreamedTraceSegment_MessagePiece `protobuf_oneof:"message_piece"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *SocketStreamedTraceSegment) Reset()         { *m = SocketStreamedTraceSegment{} }
func (m *SocketStreamedTraceSegment) String() string { return proto.CompactTextString(m) }
func (*SocketStreamedTraceSegment) ProtoMessage()    {}
func (*SocketStreamedTraceSegment) Descriptor() ([]byte, []int) {
	return fileDescriptor_03a9cebdb27ee552, []int{3}
}

func (m *SocketStreamedTraceSegment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SocketStreamedTraceSegment.Unmarshal(m, b)
}
func (m *SocketStreamedTraceSegment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SocketStreamedTraceSegment.Marshal(b, m, deterministic)
}
func (m *SocketStreamedTraceSegment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SocketStreamedTraceSegment.Merge(m, src)
}
func (m *SocketStreamedTraceSegment) XXX_Size() int {
	return xxx_messageInfo_SocketStreamedTraceSegment.Size(m)
}
func (m *SocketStreamedTraceSegment) XXX_DiscardUnknown() {
	xxx_messageInfo_SocketStreamedTraceSegment.DiscardUnknown(m)
}

var xxx_messageInfo_SocketStreamedTraceSegment proto.InternalMessageInfo

func (m *SocketStreamedTraceSegment) GetTraceId() uint64 {
	if m != nil {
		return m.TraceId
	}
	return 0
}

type isSocketStreamedTraceSegment_MessagePiece interface {
	isSocketStreamedTraceSegment_MessagePiece()
}

type SocketStreamedTraceSegment_Connection struct {
	Connection *Connection `protobuf:"bytes,2,opt,name=connection,proto3,oneof"`
}

type SocketStreamedTraceSegment_Event struct {
	Event *SocketEvent `protobuf:"bytes,3,opt,name=event,proto3,oneof"`
}

func (*SocketStreamedTraceSegment_Connection) isSocketStreamedTraceSegment_MessagePiece() {}

func (*SocketStreamedTraceSegment_Event) isSocketStreamedTraceSegment_MessagePiece() {}

func (m *SocketStreamedTraceSegment) GetMessagePiece() isSocketStreamedTraceSegment_MessagePiece {
	if m != nil {
		return m.MessagePiece
	}
	return nil
}

func (m *SocketStreamedTraceSegment) GetConnection() *Connection {
	if x, ok := m.GetMessagePiece().(*SocketStreamedTraceSegment_Connection); ok {
		return x.Connection
	}
	return nil
}

func (m *SocketStreamedTraceSegment) GetEvent() *SocketEvent {
	if x, ok := m.GetMessagePiece().(*SocketStreamedTraceSegment_Event); ok {
		return x.Event
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SocketStreamedTraceSegment) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SocketStreamedTraceSegment_Connection)(nil),
		(*SocketStreamedTraceSegment_Event)(nil),
	}
}

func init() {
	proto.RegisterType((*Connection)(nil), "envoy.data.tap.v2alpha.Connection")
	proto.RegisterType((*SocketEvent)(nil), "envoy.data.tap.v2alpha.SocketEvent")
	proto.RegisterType((*SocketEvent_Read)(nil), "envoy.data.tap.v2alpha.SocketEvent.Read")
	proto.RegisterType((*SocketEvent_Write)(nil), "envoy.data.tap.v2alpha.SocketEvent.Write")
	proto.RegisterType((*SocketEvent_Closed)(nil), "envoy.data.tap.v2alpha.SocketEvent.Closed")
	proto.RegisterType((*SocketBufferedTrace)(nil), "envoy.data.tap.v2alpha.SocketBufferedTrace")
	proto.RegisterType((*SocketStreamedTraceSegment)(nil), "envoy.data.tap.v2alpha.SocketStreamedTraceSegment")
}

func init() {
	proto.RegisterFile("envoy/data/tap/v2alpha/transport.proto", fileDescriptor_03a9cebdb27ee552)
}

var fileDescriptor_03a9cebdb27ee552 = []byte{
	// 594 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x6f, 0x13, 0x31,
	0x14, 0x4c, 0xda, 0x24, 0xa4, 0xaf, 0x34, 0xad, 0x8c, 0x84, 0xca, 0x8a, 0xaa, 0x28, 0xe5, 0xa3,
	0x70, 0xf0, 0xa2, 0x70, 0xa9, 0x84, 0x04, 0x74, 0x5b, 0xa4, 0x70, 0xab, 0xb6, 0x91, 0xe0, 0x16,
	0xbd, 0xae, 0x5f, 0xc3, 0x8a, 0xac, 0xbd, 0xf2, 0x3a, 0x81, 0xde, 0xe0, 0x1f, 0xf0, 0x7b, 0x38,
	0x71, 0x44, 0xe2, 0xc4, 0x3f, 0x42, 0xfe, 0x48, 0x5a, 0x21, 0x02, 0x81, 0x5b, 0xec, 0xcc, 0x8c,
	0xe7, 0x8d, 0xc7, 0x0b, 0xf7, 0x49, 0x4e, 0xd5, 0x45, 0x2c, 0xd0, 0x60, 0x6c, 0xb0, 0x8c, 0xa7,
	0x3d, 0x1c, 0x97, 0x6f, 0x31, 0x36, 0x1a, 0x65, 0x55, 0x2a, 0x6d, 0x78, 0xa9, 0x95, 0x51, 0xec,
	0xa6, 0xc3, 0x71, 0x8b, 0xe3, 0x06, 0x4b, 0x1e, 0x70, 0xd1, 0xae, 0xe7, 0x63, 0x99, 0xc7, 0xd3,
	0x5e, 0x9c, 0x29, 0x4d, 0x31, 0x0a, 0xa1, 0xa9, 0xaa, 0x3c, 0x31, 0xda, 0x5b, 0x70, 0x40, 0xa6,
	0x8a, 0x42, 0xc9, 0x00, 0xda, 0x1d, 0x29, 0x35, 0x1a, 0x53, 0xec, 0x56, 0x67, 0x93, 0xf3, 0xd8,
	0xe4, 0x05, 0x55, 0x06, 0x8b, 0x32, 0x00, 0x76, 0x26, 0xa2, 0xc4, 0x18, 0xa5, 0x54, 0x06, 0x4d,
	0xae, 0x64, 0x15, 0x57, 0x06, 0xcd, 0x24, 0x1c, 0xd2, 0xfd, 0x5c, 0x07, 0x38, 0x52, 0x52, 0x52,
	0x66, 0xff, 0x64, 0xcf, 0x61, 0x63, 0xac, 0x32, 0x1c, 0x0f, 0x83, 0x95, 0xed, 0x95, 0x3b, 0xf5,
	0xfd, 0xf5, 0x5e, 0xc4, 0xfd, 0x10, 0x58, 0xe6, 0x7c, 0xda, 0xe3, 0xd6, 0x2c, 0x3f, 0xf4, 0x88,
	0xf4, 0xba, 0x23, 0x84, 0x15, 0x3b, 0x84, 0x8e, 0xa6, 0x42, 0x19, 0x9a, 0x2b, 0xac, 0xfe, 0x55,
	0x61, 0xc3, 0x33, 0xc2, 0xb2, 0xfb, 0x75, 0x15, 0xd6, 0x4f, 0x55, 0xf6, 0x8e, 0xcc, 0xcb, 0x29,
	0x49, 0xc3, 0x0e, 0x60, 0x6d, 0x3e, 0xd4, 0x76, 0x3d, 0xa8, 0xf9, 0xb1, 0xf9, 0x6c, 0x6c, 0x3e,
	0x98, 0x21, 0xd2, 0x4b, 0x30, 0x7b, 0x06, 0x0d, 0x4d, 0x28, 0xc2, 0x10, 0xfb, 0xfc, 0xf7, 0x37,
	0xc1, 0xaf, 0x1c, 0xc6, 0x53, 0x42, 0xd1, 0xaf, 0xa5, 0x8e, 0xc7, 0x0e, 0xa1, 0xf9, 0x5e, 0xe7,
	0x86, 0xc2, 0x0c, 0x0f, 0x97, 0x11, 0x78, 0x6d, 0x09, 0xfd, 0x5a, 0xea, 0x99, 0xec, 0x18, 0x5a,
	0xd9, 0x58, 0x55, 0x24, 0xb6, 0x1b, 0x4e, 0xe3, 0xd1, 0x32, 0x1a, 0x47, 0x8e, 0xd1, 0xaf, 0xa5,
	0x81, 0x1b, 0x1d, 0x40, 0xc3, 0x1a, 0x63, 0x8f, 0xa1, 0x61, 0x89, 0x21, 0x85, 0xdb, 0x8b, 0xb4,
	0x12, 0x25, 0x2e, 0x52, 0x87, 0x8c, 0xde, 0x40, 0xd3, 0x39, 0xfa, 0x77, 0x2a, 0xdb, 0x01, 0x20,
	0x29, 0x86, 0x95, 0xd1, 0x84, 0x85, 0xcb, 0xb0, 0x9d, 0xae, 0x91, 0x14, 0xa7, 0x6e, 0x23, 0x6a,
	0x43, 0xcb, 0xfb, 0x4c, 0xb6, 0xa0, 0x43, 0xd6, 0xf7, 0xb0, 0xa2, 0x31, 0x65, 0x46, 0xe9, 0xee,
	0xa7, 0x15, 0xb8, 0xe1, 0x07, 0x4a, 0x26, 0xe7, 0xe7, 0xa4, 0x49, 0x0c, 0x34, 0x66, 0xc4, 0x6e,
	0x41, 0xdb, 0xd8, 0x1f, 0xc3, 0x5c, 0x38, 0x23, 0x8d, 0xf4, 0x9a, 0x5b, 0xbf, 0x12, 0x2c, 0x01,
	0xc8, 0xe6, 0x3d, 0x0c, 0x37, 0xd6, 0x5d, 0xe4, 0xf2, 0xb2, 0xb1, 0xe9, 0x15, 0x16, 0x7b, 0x0a,
	0x2d, 0x67, 0xc4, 0x96, 0x6e, 0x75, 0x7f, 0xbd, 0xb7, 0xb7, 0x44, 0xd8, 0x69, 0xa0, 0xb0, 0x7b,
	0xb6, 0xb9, 0x28, 0x86, 0x46, 0x4f, 0x64, 0x86, 0x26, 0xdc, 0x58, 0xdb, 0xb6, 0x13, 0xc5, 0x60,
	0xb6, 0xc9, 0x1e, 0xc0, 0xa6, 0xbb, 0xd9, 0x2b, 0xb8, 0xa6, 0xc3, 0x75, 0xdc, 0xf6, 0x1c, 0xd8,
	0xfd, 0x5e, 0x87, 0xc8, 0x9f, 0xe3, 0x03, 0x0b, 0x19, 0x9c, 0xd2, 0xa8, 0xb0, 0xad, 0xfe, 0x43,
	0x14, 0xc7, 0xff, 0x17, 0x45, 0xbf, 0xf6, 0x4b, 0x18, 0x4d, 0x37, 0x59, 0x28, 0xef, 0x32, 0x59,
	0xd8, 0xda, 0x3a, 0x4e, 0xb2, 0x09, 0x1b, 0x05, 0x55, 0x15, 0x8e, 0x68, 0x58, 0xe6, 0x94, 0x51,
	0xf2, 0xe2, 0xcb, 0xc7, 0x6f, 0x3f, 0x5a, 0x2b, 0x5b, 0x75, 0xb8, 0x9b, 0x2b, 0x2f, 0x55, 0x6a,
	0xf5, 0xe1, 0x62, 0x81, 0x6a, 0xd2, 0x19, 0xcc, 0x3e, 0x83, 0x27, 0xf6, 0x89, 0x9e, 0xd4, 0xcf,
	0x5a, 0xee, 0xad, 0x3e, 0xf9, 0x19, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x22, 0x4a, 0x38, 0x38, 0x05,
	0x00, 0x00,
}
