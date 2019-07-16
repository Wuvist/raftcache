// Code generated by protoc-gen-go. DO NOT EDIT.
// source: raftcache.proto

package raftcache

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Node_Statuses int32

const (
	Node_ALONE        Node_Statuses = 0
	Node_INGROUP      Node_Statuses = 1
	Node_DISCONNECTED Node_Statuses = 2
)

var Node_Statuses_name = map[int32]string{
	0: "ALONE",
	1: "INGROUP",
	2: "DISCONNECTED",
}

var Node_Statuses_value = map[string]int32{
	"ALONE":        0,
	"INGROUP":      1,
	"DISCONNECTED": 2,
}

func (x Node_Statuses) String() string {
	return proto.EnumName(Node_Statuses_name, int32(x))
}

func (Node_Statuses) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f886bd89b3f134d3, []int{0, 0}
}

type JoinResp_Results int32

const (
	JoinResp_SUCCESS       JoinResp_Results = 0
	JoinResp_ALREADYJOINED JoinResp_Results = 1
	JoinResp_REJECTED      JoinResp_Results = 2
)

var JoinResp_Results_name = map[int32]string{
	0: "SUCCESS",
	1: "ALREADYJOINED",
	2: "REJECTED",
}

var JoinResp_Results_value = map[string]int32{
	"SUCCESS":       0,
	"ALREADYJOINED": 1,
	"REJECTED":      2,
}

func (x JoinResp_Results) String() string {
	return proto.EnumName(JoinResp_Results_name, int32(x))
}

func (JoinResp_Results) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f886bd89b3f134d3, []int{1, 0}
}

type Node struct {
	Group                string        `protobuf:"bytes,1,opt,name=Group,proto3" json:"Group,omitempty"`
	ListenAddr           string        `protobuf:"bytes,2,opt,name=ListenAddr,proto3" json:"ListenAddr,omitempty"`
	Status               Node_Statuses `protobuf:"varint,3,opt,name=Status,proto3,enum=raftcache.Node_Statuses" json:"Status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_f886bd89b3f134d3, []int{0}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *Node) GetListenAddr() string {
	if m != nil {
		return m.ListenAddr
	}
	return ""
}

func (m *Node) GetStatus() Node_Statuses {
	if m != nil {
		return m.Status
	}
	return Node_ALONE
}

type JoinResp struct {
	Result               JoinResp_Results `protobuf:"varint,1,opt,name=Result,proto3,enum=raftcache.JoinResp_Results" json:"Result,omitempty"`
	Message              string           `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	GroupNodes           []*Node          `protobuf:"bytes,3,rep,name=GroupNodes,proto3" json:"GroupNodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *JoinResp) Reset()         { *m = JoinResp{} }
func (m *JoinResp) String() string { return proto.CompactTextString(m) }
func (*JoinResp) ProtoMessage()    {}
func (*JoinResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f886bd89b3f134d3, []int{1}
}

func (m *JoinResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinResp.Unmarshal(m, b)
}
func (m *JoinResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinResp.Marshal(b, m, deterministic)
}
func (m *JoinResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinResp.Merge(m, src)
}
func (m *JoinResp) XXX_Size() int {
	return xxx_messageInfo_JoinResp.Size(m)
}
func (m *JoinResp) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinResp.DiscardUnknown(m)
}

var xxx_messageInfo_JoinResp proto.InternalMessageInfo

func (m *JoinResp) GetResult() JoinResp_Results {
	if m != nil {
		return m.Result
	}
	return JoinResp_SUCCESS
}

func (m *JoinResp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *JoinResp) GetGroupNodes() []*Node {
	if m != nil {
		return m.GroupNodes
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_f886bd89b3f134d3, []int{2}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type LeaveResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaveResp) Reset()         { *m = LeaveResp{} }
func (m *LeaveResp) String() string { return proto.CompactTextString(m) }
func (*LeaveResp) ProtoMessage()    {}
func (*LeaveResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f886bd89b3f134d3, []int{3}
}

func (m *LeaveResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeaveResp.Unmarshal(m, b)
}
func (m *LeaveResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeaveResp.Marshal(b, m, deterministic)
}
func (m *LeaveResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaveResp.Merge(m, src)
}
func (m *LeaveResp) XXX_Size() int {
	return xxx_messageInfo_LeaveResp.Size(m)
}
func (m *LeaveResp) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaveResp.DiscardUnknown(m)
}

var xxx_messageInfo_LeaveResp proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("raftcache.Node_Statuses", Node_Statuses_name, Node_Statuses_value)
	proto.RegisterEnum("raftcache.JoinResp_Results", JoinResp_Results_name, JoinResp_Results_value)
	proto.RegisterType((*Node)(nil), "raftcache.Node")
	proto.RegisterType((*JoinResp)(nil), "raftcache.JoinResp")
	proto.RegisterType((*Empty)(nil), "raftcache.Empty")
	proto.RegisterType((*LeaveResp)(nil), "raftcache.LeaveResp")
}

func init() { proto.RegisterFile("raftcache.proto", fileDescriptor_f886bd89b3f134d3) }

var fileDescriptor_f886bd89b3f134d3 = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xc1, 0xce, 0x9a, 0x40,
	0x14, 0x85, 0x99, 0xff, 0x17, 0x91, 0xab, 0x55, 0x7a, 0xeb, 0x82, 0xd8, 0xa4, 0x31, 0xb3, 0x72,
	0xd1, 0x50, 0x83, 0x4d, 0xba, 0x26, 0x30, 0x31, 0x1a, 0x0a, 0x66, 0xa8, 0x8b, 0x2e, 0xa9, 0x8e,
	0xd6, 0xa4, 0x15, 0xc2, 0x8c, 0x4d, 0xfa, 0x2c, 0xdd, 0xf6, 0x4d, 0xfa, 0x62, 0x0d, 0x08, 0x96,
	0x68, 0x97, 0x73, 0xee, 0xb9, 0xf7, 0x7c, 0x87, 0x00, 0xa3, 0x22, 0x3d, 0xa8, 0x5d, 0xba, 0xfb,
	0x2a, 0x9c, 0xbc, 0xc8, 0x54, 0x86, 0xe6, 0x4d, 0xa0, 0xbf, 0x09, 0x74, 0xa2, 0x6c, 0x2f, 0x70,
	0x0c, 0xfa, 0xb2, 0xc8, 0x2e, 0xb9, 0x4d, 0xa6, 0x64, 0x66, 0xf2, 0xeb, 0x03, 0xdf, 0x00, 0x84,
	0x27, 0xa9, 0xc4, 0xd9, 0xdb, 0xef, 0x0b, 0xfb, 0xa9, 0x1a, 0xb5, 0x14, 0x9c, 0x43, 0x37, 0x51,
	0xa9, 0xba, 0x48, 0xfb, 0x79, 0x4a, 0x66, 0x43, 0xd7, 0x76, 0xfe, 0x65, 0x95, 0x67, 0x9d, 0xeb,
	0x54, 0x48, 0x5e, 0xfb, 0xe8, 0x7b, 0xe8, 0x35, 0x1a, 0x9a, 0xa0, 0x7b, 0x61, 0x1c, 0x31, 0x4b,
	0xc3, 0x3e, 0x18, 0xab, 0x68, 0xc9, 0xe3, 0xed, 0xc6, 0x22, 0x68, 0xc1, 0x20, 0x58, 0x25, 0x7e,
	0x1c, 0x45, 0xcc, 0xff, 0xc4, 0x02, 0xeb, 0x89, 0xfe, 0x21, 0xd0, 0x5b, 0x67, 0xa7, 0x33, 0x17,
	0x32, 0xc7, 0x05, 0x74, 0xb9, 0x90, 0x97, 0x6f, 0xaa, 0x62, 0x1d, 0xba, 0xaf, 0x5b, 0xa1, 0x8d,
	0xc9, 0xb9, 0x3a, 0x24, 0xaf, 0xad, 0x68, 0x83, 0xf1, 0x51, 0x48, 0x99, 0x1e, 0x45, 0x5d, 0xa3,
	0x79, 0xe2, 0x3b, 0x80, 0xaa, 0x6c, 0xc9, 0x5b, 0xf6, 0x78, 0x9e, 0xf5, 0xdd, 0xd1, 0x5d, 0x0f,
	0xde, 0xb2, 0xd0, 0x0f, 0x60, 0xd4, 0xd7, 0x4b, 0xec, 0x64, 0xeb, 0xfb, 0x2c, 0x49, 0x2c, 0x0d,
	0x5f, 0xc2, 0x0b, 0x2f, 0xe4, 0xcc, 0x0b, 0x3e, 0xaf, 0xe3, 0x55, 0xc4, 0x02, 0x8b, 0xe0, 0x00,
	0x7a, 0x9c, 0xad, 0x9b, 0x16, 0x06, 0xe8, 0xec, 0x7b, 0xae, 0x7e, 0xd2, 0x3e, 0x98, 0xa1, 0x48,
	0x7f, 0x88, 0x92, 0xd4, 0xfd, 0x45, 0xc0, 0xe4, 0xe9, 0x41, 0xf9, 0x65, 0x1a, 0xbe, 0x85, 0xce,
	0xe6, 0x74, 0x3e, 0xa2, 0xd5, 0x22, 0xa8, 0x96, 0x26, 0x0f, 0x0a, 0xd5, 0xd0, 0x81, 0x4e, 0xd9,
	0x18, 0xef, 0x79, 0x27, 0xaf, 0xfe, 0xf3, 0x4d, 0xa8, 0x86, 0x73, 0xd0, 0xab, 0xe0, 0xc7, 0x85,
	0x71, 0x4b, 0xb8, 0xb1, 0x51, 0xed, 0x4b, 0xb7, 0xfa, 0x65, 0x16, 0x7f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x66, 0xb0, 0x93, 0x83, 0x45, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RaftCacheClient is the client API for RaftCache service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RaftCacheClient interface {
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	Join(ctx context.Context, in *Node, opts ...grpc.CallOption) (*JoinResp, error)
	Leave(ctx context.Context, in *Node, opts ...grpc.CallOption) (*LeaveResp, error)
}

type raftCacheClient struct {
	cc *grpc.ClientConn
}

func NewRaftCacheClient(cc *grpc.ClientConn) RaftCacheClient {
	return &raftCacheClient{cc}
}

func (c *raftCacheClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/raftcache.RaftCache/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftCacheClient) Join(ctx context.Context, in *Node, opts ...grpc.CallOption) (*JoinResp, error) {
	out := new(JoinResp)
	err := c.cc.Invoke(ctx, "/raftcache.RaftCache/Join", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftCacheClient) Leave(ctx context.Context, in *Node, opts ...grpc.CallOption) (*LeaveResp, error) {
	out := new(LeaveResp)
	err := c.cc.Invoke(ctx, "/raftcache.RaftCache/Leave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RaftCacheServer is the server API for RaftCache service.
type RaftCacheServer interface {
	Ping(context.Context, *Empty) (*Empty, error)
	Join(context.Context, *Node) (*JoinResp, error)
	Leave(context.Context, *Node) (*LeaveResp, error)
}

// UnimplementedRaftCacheServer can be embedded to have forward compatible implementations.
type UnimplementedRaftCacheServer struct {
}

func (*UnimplementedRaftCacheServer) Ping(ctx context.Context, req *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedRaftCacheServer) Join(ctx context.Context, req *Node) (*JoinResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (*UnimplementedRaftCacheServer) Leave(ctx context.Context, req *Node) (*LeaveResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Leave not implemented")
}

func RegisterRaftCacheServer(s *grpc.Server, srv RaftCacheServer) {
	s.RegisterService(&_RaftCache_serviceDesc, srv)
}

func _RaftCache_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftCacheServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raftcache.RaftCache/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftCacheServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftCache_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftCacheServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raftcache.RaftCache/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftCacheServer).Join(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftCache_Leave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftCacheServer).Leave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raftcache.RaftCache/Leave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftCacheServer).Leave(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

var _RaftCache_serviceDesc = grpc.ServiceDesc{
	ServiceName: "raftcache.RaftCache",
	HandlerType: (*RaftCacheServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _RaftCache_Ping_Handler,
		},
		{
			MethodName: "Join",
			Handler:    _RaftCache_Join_Handler,
		},
		{
			MethodName: "Leave",
			Handler:    _RaftCache_Leave_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "raftcache.proto",
}
