// Code generated by protoc-gen-go. DO NOT EDIT.
// source: options.proto

package brpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TalkType int32

const (
	TalkType_TALK_TYPE_NORMAL TalkType = 0
	TalkType_TALK_TYPE_ONEWAY TalkType = 1
)

var TalkType_name = map[int32]string{
	0: "TALK_TYPE_NORMAL",
	1: "TALK_TYPE_ONEWAY",
}
var TalkType_value = map[string]int32{
	"TALK_TYPE_NORMAL": 0,
	"TALK_TYPE_ONEWAY": 1,
}

func (x TalkType) Enum() *TalkType {
	p := new(TalkType)
	*p = x
	return p
}
func (x TalkType) String() string {
	return proto.EnumName(TalkType_name, int32(x))
}
func (x *TalkType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TalkType_value, data, "TalkType")
	if err != nil {
		return err
	}
	*x = TalkType(value)
	return nil
}
func (TalkType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_options_ce1982f1af64af84, []int{0}
}

type ConnectionType int32

const (
	// bit-exclusive values since we may OR them to represent supported types.
	ConnectionType_CONNECTION_TYPE_UNKNOWN ConnectionType = 0
	ConnectionType_CONNECTION_TYPE_SINGLE  ConnectionType = 1
	ConnectionType_CONNECTION_TYPE_POOLED  ConnectionType = 2
	ConnectionType_CONNECTION_TYPE_SHORT   ConnectionType = 4
)

var ConnectionType_name = map[int32]string{
	0: "CONNECTION_TYPE_UNKNOWN",
	1: "CONNECTION_TYPE_SINGLE",
	2: "CONNECTION_TYPE_POOLED",
	4: "CONNECTION_TYPE_SHORT",
}
var ConnectionType_value = map[string]int32{
	"CONNECTION_TYPE_UNKNOWN": 0,
	"CONNECTION_TYPE_SINGLE":  1,
	"CONNECTION_TYPE_POOLED":  2,
	"CONNECTION_TYPE_SHORT":   4,
}

func (x ConnectionType) Enum() *ConnectionType {
	p := new(ConnectionType)
	*p = x
	return p
}
func (x ConnectionType) String() string {
	return proto.EnumName(ConnectionType_name, int32(x))
}
func (x *ConnectionType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ConnectionType_value, data, "ConnectionType")
	if err != nil {
		return err
	}
	*x = ConnectionType(value)
	return nil
}
func (ConnectionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_options_ce1982f1af64af84, []int{1}
}

type ProtocolType int32

const (
	ProtocolType_PROTOCOL_UNKNOWN           ProtocolType = 0
	ProtocolType_PROTOCOL_BAIDU_STD         ProtocolType = 1
	ProtocolType_PROTOCOL_STREAMING_RPC     ProtocolType = 2
	ProtocolType_PROTOCOL_HULU_PBRPC        ProtocolType = 3
	ProtocolType_PROTOCOL_SOFA_PBRPC        ProtocolType = 4
	ProtocolType_PROTOCOL_RTMP              ProtocolType = 5
	ProtocolType_PROTOCOL_THRIFT            ProtocolType = 6
	ProtocolType_PROTOCOL_HTTP              ProtocolType = 7
	ProtocolType_PROTOCOL_PUBLIC_PBRPC      ProtocolType = 8
	ProtocolType_PROTOCOL_NOVA_PBRPC        ProtocolType = 9
	ProtocolType_PROTOCOL_NSHEAD_CLIENT     ProtocolType = 10
	ProtocolType_PROTOCOL_NSHEAD            ProtocolType = 11
	ProtocolType_PROTOCOL_HADOOP_RPC        ProtocolType = 12
	ProtocolType_PROTOCOL_HADOOP_SERVER_RPC ProtocolType = 13
	ProtocolType_PROTOCOL_MONGO             ProtocolType = 14
	ProtocolType_PROTOCOL_UBRPC_COMPACK     ProtocolType = 15
	ProtocolType_PROTOCOL_DIDX_CLIENT       ProtocolType = 16
	ProtocolType_PROTOCOL_REDIS             ProtocolType = 17
	ProtocolType_PROTOCOL_MEMCACHE          ProtocolType = 18
	ProtocolType_PROTOCOL_ITP               ProtocolType = 19
	ProtocolType_PROTOCOL_NSHEAD_MCPACK     ProtocolType = 20
	ProtocolType_PROTOCOL_DISP_IDL          ProtocolType = 21
	ProtocolType_PROTOCOL_ERSDA_CLIENT      ProtocolType = 22
	ProtocolType_PROTOCOL_UBRPC_MCPACK2     ProtocolType = 23
	// Reserve special protocol for cds-agent, which depends on FIFO right now
	ProtocolType_PROTOCOL_CDS_AGENT ProtocolType = 24
	ProtocolType_PROTOCOL_ESP       ProtocolType = 25
	ProtocolType_PROTOCOL_H2        ProtocolType = 26
)

var ProtocolType_name = map[int32]string{
	0:  "PROTOCOL_UNKNOWN",
	1:  "PROTOCOL_BAIDU_STD",
	2:  "PROTOCOL_STREAMING_RPC",
	3:  "PROTOCOL_HULU_PBRPC",
	4:  "PROTOCOL_SOFA_PBRPC",
	5:  "PROTOCOL_RTMP",
	6:  "PROTOCOL_THRIFT",
	7:  "PROTOCOL_HTTP",
	8:  "PROTOCOL_PUBLIC_PBRPC",
	9:  "PROTOCOL_NOVA_PBRPC",
	10: "PROTOCOL_NSHEAD_CLIENT",
	11: "PROTOCOL_NSHEAD",
	12: "PROTOCOL_HADOOP_RPC",
	13: "PROTOCOL_HADOOP_SERVER_RPC",
	14: "PROTOCOL_MONGO",
	15: "PROTOCOL_UBRPC_COMPACK",
	16: "PROTOCOL_DIDX_CLIENT",
	17: "PROTOCOL_REDIS",
	18: "PROTOCOL_MEMCACHE",
	19: "PROTOCOL_ITP",
	20: "PROTOCOL_NSHEAD_MCPACK",
	21: "PROTOCOL_DISP_IDL",
	22: "PROTOCOL_ERSDA_CLIENT",
	23: "PROTOCOL_UBRPC_MCPACK2",
	24: "PROTOCOL_CDS_AGENT",
	25: "PROTOCOL_ESP",
	26: "PROTOCOL_H2",
}
var ProtocolType_value = map[string]int32{
	"PROTOCOL_UNKNOWN":           0,
	"PROTOCOL_BAIDU_STD":         1,
	"PROTOCOL_STREAMING_RPC":     2,
	"PROTOCOL_HULU_PBRPC":        3,
	"PROTOCOL_SOFA_PBRPC":        4,
	"PROTOCOL_RTMP":              5,
	"PROTOCOL_THRIFT":            6,
	"PROTOCOL_HTTP":              7,
	"PROTOCOL_PUBLIC_PBRPC":      8,
	"PROTOCOL_NOVA_PBRPC":        9,
	"PROTOCOL_NSHEAD_CLIENT":     10,
	"PROTOCOL_NSHEAD":            11,
	"PROTOCOL_HADOOP_RPC":        12,
	"PROTOCOL_HADOOP_SERVER_RPC": 13,
	"PROTOCOL_MONGO":             14,
	"PROTOCOL_UBRPC_COMPACK":     15,
	"PROTOCOL_DIDX_CLIENT":       16,
	"PROTOCOL_REDIS":             17,
	"PROTOCOL_MEMCACHE":          18,
	"PROTOCOL_ITP":               19,
	"PROTOCOL_NSHEAD_MCPACK":     20,
	"PROTOCOL_DISP_IDL":          21,
	"PROTOCOL_ERSDA_CLIENT":      22,
	"PROTOCOL_UBRPC_MCPACK2":     23,
	"PROTOCOL_CDS_AGENT":         24,
	"PROTOCOL_ESP":               25,
	"PROTOCOL_H2":                26,
}

func (x ProtocolType) Enum() *ProtocolType {
	p := new(ProtocolType)
	*p = x
	return p
}
func (x ProtocolType) String() string {
	return proto.EnumName(ProtocolType_name, int32(x))
}
func (x *ProtocolType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ProtocolType_value, data, "ProtocolType")
	if err != nil {
		return err
	}
	*x = ProtocolType(value)
	return nil
}
func (ProtocolType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_options_ce1982f1af64af84, []int{2}
}

type CompressType int32

const (
	CompressType_COMPRESS_TYPE_NONE   CompressType = 0
	CompressType_COMPRESS_TYPE_SNAPPY CompressType = 1
	CompressType_COMPRESS_TYPE_GZIP   CompressType = 2
	CompressType_COMPRESS_TYPE_ZLIB   CompressType = 3
	CompressType_COMPRESS_TYPE_LZ4    CompressType = 4
)

var CompressType_name = map[int32]string{
	0: "COMPRESS_TYPE_NONE",
	1: "COMPRESS_TYPE_SNAPPY",
	2: "COMPRESS_TYPE_GZIP",
	3: "COMPRESS_TYPE_ZLIB",
	4: "COMPRESS_TYPE_LZ4",
}
var CompressType_value = map[string]int32{
	"COMPRESS_TYPE_NONE":   0,
	"COMPRESS_TYPE_SNAPPY": 1,
	"COMPRESS_TYPE_GZIP":   2,
	"COMPRESS_TYPE_ZLIB":   3,
	"COMPRESS_TYPE_LZ4":    4,
}

func (x CompressType) Enum() *CompressType {
	p := new(CompressType)
	*p = x
	return p
}
func (x CompressType) String() string {
	return proto.EnumName(CompressType_name, int32(x))
}
func (x *CompressType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CompressType_value, data, "CompressType")
	if err != nil {
		return err
	}
	*x = CompressType(value)
	return nil
}
func (CompressType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_options_ce1982f1af64af84, []int{3}
}

type ChunkInfo struct {
	StreamId             *int64   `protobuf:"varint,1,req,name=stream_id,json=streamId" json:"stream_id,omitempty"`
	ChunkId              *int64   `protobuf:"varint,2,req,name=chunk_id,json=chunkId" json:"chunk_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChunkInfo) Reset()         { *m = ChunkInfo{} }
func (m *ChunkInfo) String() string { return proto.CompactTextString(m) }
func (*ChunkInfo) ProtoMessage()    {}
func (*ChunkInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_options_ce1982f1af64af84, []int{0}
}
func (m *ChunkInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChunkInfo.Unmarshal(m, b)
}
func (m *ChunkInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChunkInfo.Marshal(b, m, deterministic)
}
func (dst *ChunkInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChunkInfo.Merge(dst, src)
}
func (m *ChunkInfo) XXX_Size() int {
	return xxx_messageInfo_ChunkInfo.Size(m)
}
func (m *ChunkInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ChunkInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ChunkInfo proto.InternalMessageInfo

func (m *ChunkInfo) GetStreamId() int64 {
	if m != nil && m.StreamId != nil {
		return *m.StreamId
	}
	return 0
}

func (m *ChunkInfo) GetChunkId() int64 {
	if m != nil && m.ChunkId != nil {
		return *m.ChunkId
	}
	return 0
}

var E_ServiceTimeout = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.ServiceOptions)(nil),
	ExtensionType: (*int64)(nil),
	Field:         90000,
	Name:          "brpc.service_timeout",
	Tag:           "varint,90000,opt,name=service_timeout,json=serviceTimeout,def=10000",
	Filename:      "options.proto",
}

var E_RequestTalkType = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*TalkType)(nil),
	Field:         90001,
	Name:          "brpc.request_talk_type",
	Tag:           "varint,90001,opt,name=request_talk_type,json=requestTalkType,enum=brpc.TalkType,def=0",
	Filename:      "options.proto",
}

var E_ResponseTalkType = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*TalkType)(nil),
	Field:         90002,
	Name:          "brpc.response_talk_type",
	Tag:           "varint,90002,opt,name=response_talk_type,json=responseTalkType,enum=brpc.TalkType,def=0",
	Filename:      "options.proto",
}

var E_MethodTimeout = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*int64)(nil),
	Field:         90003,
	Name:          "brpc.method_timeout",
	Tag:           "varint,90003,opt,name=method_timeout,json=methodTimeout",
	Filename:      "options.proto",
}

var E_RequestCompression = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*CompressType)(nil),
	Field:         90004,
	Name:          "brpc.request_compression",
	Tag:           "varint,90004,opt,name=request_compression,json=requestCompression,enum=brpc.CompressType,def=0",
	Filename:      "options.proto",
}

var E_ResponseCompression = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*CompressType)(nil),
	Field:         90005,
	Name:          "brpc.response_compression",
	Tag:           "varint,90005,opt,name=response_compression,json=responseCompression,enum=brpc.CompressType,def=0",
	Filename:      "options.proto",
}

func init() {
	proto.RegisterType((*ChunkInfo)(nil), "brpc.ChunkInfo")
	proto.RegisterEnum("brpc.TalkType", TalkType_name, TalkType_value)
	proto.RegisterEnum("brpc.ConnectionType", ConnectionType_name, ConnectionType_value)
	proto.RegisterEnum("brpc.ProtocolType", ProtocolType_name, ProtocolType_value)
	proto.RegisterEnum("brpc.CompressType", CompressType_name, CompressType_value)
	proto.RegisterExtension(E_ServiceTimeout)
	proto.RegisterExtension(E_RequestTalkType)
	proto.RegisterExtension(E_ResponseTalkType)
	proto.RegisterExtension(E_MethodTimeout)
	proto.RegisterExtension(E_RequestCompression)
	proto.RegisterExtension(E_ResponseCompression)
}

func init() { proto.RegisterFile("options.proto", fileDescriptor_options_ce1982f1af64af84) }

var fileDescriptor_options_ce1982f1af64af84 = []byte{
	// 783 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4d, 0x6f, 0xdb, 0x46,
	0x10, 0xb5, 0x64, 0xbb, 0x96, 0x37, 0xb6, 0xb4, 0x5e, 0xf9, 0x3b, 0x40, 0x6a, 0xf4, 0x14, 0xf8,
	0xa0, 0xb8, 0x46, 0xd1, 0x83, 0x6e, 0xd4, 0x72, 0x23, 0x2d, 0x4c, 0xee, 0x6e, 0x97, 0xab, 0xa4,
	0xf6, 0x85, 0x70, 0x28, 0x26, 0x11, 0x2c, 0x71, 0x59, 0x92, 0x2a, 0x10, 0xf4, 0xd2, 0x5b, 0x8f,
	0xfd, 0xfe, 0x3d, 0xfd, 0x6b, 0x05, 0x3f, 0x6d, 0x46, 0x02, 0x1c, 0x1f, 0xf9, 0xde, 0xcc, 0x7b,
	0x33, 0x0f, 0xcb, 0x01, 0xbb, 0x3a, 0x4c, 0xa6, 0x3a, 0x88, 0x7b, 0x61, 0xa4, 0x13, 0x8d, 0x36,
	0xde, 0x45, 0xa1, 0x77, 0x7a, 0xf6, 0x41, 0xeb, 0x0f, 0x33, 0xff, 0x55, 0x86, 0xbd, 0x5b, 0xbc,
	0x7f, 0x35, 0xf1, 0x63, 0x2f, 0x9a, 0x86, 0x89, 0x8e, 0xf2, 0xba, 0x6f, 0x30, 0xd8, 0xc6, 0x1f,
	0x17, 0xc1, 0x1d, 0x0d, 0xde, 0x6b, 0xf4, 0x1c, 0x6c, 0xc7, 0x49, 0xe4, 0xdf, 0xce, 0xdd, 0xe9,
	0xe4, 0xb8, 0x71, 0xd6, 0x7c, 0xb9, 0x2e, 0x5b, 0x39, 0x40, 0x27, 0xe8, 0x04, 0xb4, 0xbc, 0xb4,
	0x32, 0xe5, 0x9a, 0x19, 0xb7, 0x95, 0x7d, 0xd3, 0xc9, 0xf9, 0xf7, 0xa0, 0xa5, 0x6e, 0x67, 0x77,
	0xea, 0x53, 0xe8, 0xa3, 0x7d, 0x00, 0x95, 0x61, 0x5d, 0xb9, 0xea, 0x5a, 0x10, 0x97, 0x71, 0x69,
	0x1b, 0x16, 0x5c, 0xab, 0xa3, 0x9c, 0x91, 0xb7, 0xc6, 0x35, 0x6c, 0x9c, 0xff, 0xda, 0x00, 0x6d,
	0xac, 0x83, 0xc0, 0xf7, 0xd2, 0xd1, 0xb3, 0xf6, 0xe7, 0xe0, 0x08, 0x73, 0xc6, 0x08, 0x56, 0x94,
	0xb3, 0xbc, 0x7c, 0xcc, 0xae, 0x18, 0x7f, 0xcb, 0xe0, 0x1a, 0x3a, 0x05, 0x87, 0x9f, 0x93, 0x0e,
	0x65, 0x43, 0x8b, 0xc0, 0xc6, 0x2a, 0x4e, 0x70, 0x6e, 0x11, 0x13, 0x36, 0xd1, 0x09, 0x38, 0x58,
	0xea, 0x1b, 0x71, 0xa9, 0xe0, 0xc6, 0xf9, 0xef, 0x9b, 0x60, 0x47, 0xa4, 0x49, 0x78, 0x7a, 0x56,
	0xce, 0x2f, 0x24, 0x57, 0x1c, 0x73, 0xeb, 0x81, 0xf3, 0x21, 0x40, 0x15, 0x3a, 0x30, 0xa8, 0x39,
	0x76, 0x1d, 0x65, 0xe6, 0xae, 0x15, 0xee, 0x28, 0x49, 0x0c, 0x9b, 0xb2, 0xa1, 0x2b, 0x05, 0x86,
	0x4d, 0x74, 0x04, 0xba, 0x15, 0x37, 0x1a, 0x5b, 0x63, 0x57, 0x0c, 0x52, 0x62, 0xbd, 0x46, 0x38,
	0xfc, 0xb5, 0x51, 0x10, 0x1b, 0x68, 0x0f, 0xec, 0x56, 0x84, 0x54, 0xb6, 0x80, 0x9b, 0xa8, 0x0b,
	0x3a, 0x15, 0xa4, 0x46, 0x92, 0xbe, 0x56, 0xf0, 0xab, 0x5a, 0xdd, 0x48, 0x29, 0x01, 0xb7, 0xd2,
	0x15, 0x2b, 0x48, 0x8c, 0x07, 0x16, 0xc5, 0x85, 0x6a, 0xab, 0x66, 0xc7, 0xf8, 0x9b, 0xd2, 0x6e,
	0xbb, 0x36, 0x3c, 0x73, 0x46, 0xc4, 0x30, 0x5d, 0x6c, 0x51, 0xc2, 0x14, 0x04, 0x35, 0xdf, 0x9c,
	0x83, 0xcf, 0xea, 0x1b, 0x19, 0x26, 0xe7, 0x22, 0x5b, 0x75, 0x07, 0xbd, 0x00, 0xa7, 0x9f, 0x13,
	0x0e, 0x91, 0x6f, 0x88, 0xcc, 0xf8, 0x5d, 0x84, 0x40, 0xbb, 0xe2, 0x6d, 0xce, 0x86, 0x1c, 0xb6,
	0x6b, 0xee, 0xe3, 0x74, 0x22, 0x17, 0x73, 0x5b, 0x18, 0xf8, 0x0a, 0x76, 0xd0, 0x31, 0xd8, 0xaf,
	0x38, 0x93, 0x9a, 0x3f, 0x96, 0x73, 0xc1, 0x9a, 0x92, 0x24, 0x26, 0x75, 0xe0, 0x1e, 0x3a, 0x00,
	0x7b, 0xf7, 0xea, 0xc4, 0xc6, 0x06, 0x1e, 0x11, 0x88, 0x10, 0x04, 0x3b, 0x15, 0x4c, 0x95, 0x80,
	0xdd, 0x55, 0x0b, 0xdb, 0x38, 0xb3, 0xdc, 0xaf, 0x89, 0x98, 0xd4, 0x11, 0x2e, 0x35, 0x2d, 0x78,
	0x50, 0xcb, 0x95, 0x48, 0xc7, 0x34, 0xca, 0x51, 0x0e, 0x57, 0x2c, 0x90, 0x8b, 0x5d, 0xc2, 0xa3,
	0xda, 0x7b, 0xc1, 0xa6, 0xe3, 0x1a, 0xc3, 0xb4, 0xe7, 0xb8, 0x36, 0x13, 0x71, 0x04, 0x3c, 0x41,
	0x1d, 0xf0, 0xec, 0x3e, 0xba, 0x4b, 0x78, 0x7a, 0xfe, 0x5b, 0x03, 0xec, 0x60, 0x3d, 0x0f, 0x23,
	0x3f, 0x8e, 0xb3, 0x17, 0x79, 0x08, 0x50, 0x9a, 0x8c, 0x24, 0x8e, 0x53, 0xfe, 0x55, 0x8c, 0xc0,
	0xb5, 0x34, 0xa4, 0x3a, 0xee, 0x30, 0x43, 0x88, 0x6b, 0xd8, 0x58, 0xee, 0x18, 0xde, 0x50, 0x01,
	0x9b, 0xcb, 0xf8, 0x8d, 0x45, 0x07, 0x70, 0x3d, 0xdd, 0xbd, 0x8e, 0x5b, 0x37, 0xdf, 0xc1, 0x8d,
	0xfe, 0x0f, 0xa0, 0x13, 0xfb, 0xd1, 0xcf, 0x53, 0xcf, 0x77, 0x93, 0xe9, 0xdc, 0xd7, 0x8b, 0x04,
	0x7d, 0xdd, 0xcb, 0x2f, 0x4a, 0xaf, 0xbc, 0x28, 0x3d, 0x27, 0xaf, 0xe0, 0xf9, 0xf5, 0x39, 0xfe,
	0xe3, 0xbf, 0xcd, 0xb3, 0xc6, 0xcb, 0xf5, 0xfe, 0xe6, 0xb7, 0x17, 0x17, 0x17, 0x17, 0xb2, 0x5d,
	0x08, 0xa8, 0xbc, 0xbf, 0x1f, 0x80, 0xbd, 0xc8, 0xff, 0x69, 0xe1, 0xc7, 0x89, 0x9b, 0xdc, 0xce,
	0xee, 0xdc, 0x24, 0x5d, 0xf0, 0xc5, 0x92, 0xa8, 0xed, 0x27, 0x1f, 0xf5, 0xa4, 0xd4, 0xfc, 0x33,
	0xd3, 0x6c, 0x5f, 0xb6, 0x7b, 0xe9, 0x51, 0xeb, 0x95, 0xa7, 0xa6, 0xbf, 0x74, 0x68, 0x64, 0xa7,
	0x10, 0xaf, 0x4a, 0x42, 0x80, 0x22, 0x3f, 0x0e, 0x75, 0x10, 0xfb, 0x4f, 0x30, 0xfc, 0xeb, 0x8b,
	0x0d, 0x61, 0xa9, 0x5e, 0xd5, 0x0c, 0x41, 0x7b, 0x9e, 0x89, 0x55, 0x99, 0x3d, 0xe6, 0xf6, 0x77,
	0x1e, 0x99, 0xdc, 0xcd, 0xfb, 0xca, 0xa8, 0x3e, 0x81, 0x6e, 0x19, 0x95, 0x57, 0x3c, 0x87, 0xa9,
	0x0e, 0x1e, 0x55, 0xfb, 0xa7, 0x98, 0x1d, 0xe5, 0xb3, 0x3f, 0x7c, 0x49, 0xfd, 0x15, 0xef, 0x48,
	0xa2, 0xc2, 0x04, 0xdf, 0x7b, 0xf4, 0x7f, 0x01, 0xfb, 0x55, 0x6a, 0x4f, 0xf1, 0xfe, 0xf7, 0x89,
	0xde, 0xdd, 0xd2, 0xe5, 0x81, 0xf9, 0xa0, 0x0b, 0x5a, 0x9e, 0x9e, 0x67, 0x0a, 0x83, 0xad, 0x42,
	0xf9, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x09, 0x80, 0x90, 0xae, 0xde, 0x06, 0x00, 0x00,
}