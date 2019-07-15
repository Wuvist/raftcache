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
	AdminAddr            string        `protobuf:"bytes,2,opt,name=AdminAddr,proto3" json:"AdminAddr,omitempty"`
	ListenAddr           string        `protobuf:"bytes,3,opt,name=ListenAddr,proto3" json:"ListenAddr,omitempty"`
	Status               Node_Statuses `protobuf:"varint,4,opt,name=Status,proto3,enum=raftcache.Node_Statuses" json:"Status,omitempty"`
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

func (m *Node) GetAdminAddr() string {
	if m != nil {
		return m.AdminAddr
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
	GroupNodes           []*Node          `protobuf:"bytes,2,rep,name=GroupNodes,proto3" json:"GroupNodes,omitempty"`
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

func (m *JoinResp) GetGroupNodes() []*Node {
	if m != nil {
		return m.GroupNodes
	}
	return nil
}

type LeaveResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaveResp) Reset()         { *m = LeaveResp{} }
func (m *LeaveResp) String() string { return proto.CompactTextString(m) }
func (*LeaveResp) ProtoMessage()    {}
func (*LeaveResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f886bd89b3f134d3, []int{2}
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
	proto.RegisterType((*LeaveResp)(nil), "raftcache.LeaveResp")
}

func init() { proto.RegisterFile("raftcache.proto", fileDescriptor_f886bd89b3f134d3) }

var fileDescriptor_f886bd89b3f134d3 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x6e, 0x82, 0x40,
	0x18, 0xc4, 0x59, 0xff, 0xf3, 0x69, 0x95, 0x7e, 0xf5, 0x40, 0xda, 0xa6, 0x31, 0x9c, 0x3c, 0x51,
	0x83, 0x4d, 0x7a, 0x26, 0xb0, 0x31, 0x1a, 0x02, 0xcd, 0x52, 0x0f, 0x3d, 0x52, 0x59, 0x53, 0x93,
	0x2a, 0x86, 0x5d, 0xfb, 0x4e, 0x7d, 0x8d, 0x3e, 0x59, 0x03, 0x88, 0x12, 0xdb, 0x23, 0x33, 0xc3,
	0x37, 0xbf, 0xc9, 0xc2, 0x20, 0x8d, 0xd6, 0x72, 0x15, 0xad, 0x3e, 0xb8, 0xb9, 0x4f, 0x13, 0x99,
	0xa0, 0x7a, 0x12, 0x8c, 0x1f, 0x02, 0x0d, 0x3f, 0x89, 0x39, 0x0e, 0xa1, 0x39, 0x4b, 0x93, 0xc3,
	0x5e, 0x27, 0x23, 0x32, 0x56, 0x59, 0xf1, 0x81, 0xf7, 0xa0, 0xda, 0xf1, 0x76, 0xb3, 0xb3, 0xe3,
	0x38, 0xd5, 0x6b, 0xb9, 0x73, 0x16, 0xf0, 0x01, 0xc0, 0xdb, 0x08, 0xc9, 0x0b, 0xbb, 0x9e, 0xdb,
	0x15, 0x05, 0x27, 0xd0, 0x0a, 0x65, 0x24, 0x0f, 0x42, 0x6f, 0x8c, 0xc8, 0xb8, 0x6f, 0xe9, 0xe6,
	0x99, 0x24, 0x2b, 0x35, 0x0b, 0x97, 0x0b, 0x76, 0xcc, 0x19, 0x4f, 0xd0, 0x29, 0x35, 0x54, 0xa1,
	0x69, 0x7b, 0x81, 0x4f, 0x35, 0x05, 0xbb, 0xd0, 0x9e, 0xfb, 0x33, 0x16, 0x2c, 0x5f, 0x34, 0x82,
	0x1a, 0xf4, 0xdc, 0x79, 0xe8, 0x04, 0xbe, 0x4f, 0x9d, 0x57, 0xea, 0x6a, 0x35, 0xe3, 0x9b, 0x40,
	0x67, 0x91, 0x6c, 0x76, 0x8c, 0x8b, 0x3d, 0x4e, 0xa1, 0xc5, 0xb8, 0x38, 0x7c, 0xca, 0x7c, 0x49,
	0xdf, 0xba, 0xab, 0x94, 0x96, 0x21, 0xb3, 0x48, 0x08, 0x76, 0x8c, 0xe2, 0x23, 0x40, 0x3e, 0x38,
	0xa3, 0x12, 0x7a, 0x6d, 0x54, 0x1f, 0x77, 0xad, 0xc1, 0x05, 0x2d, 0xab, 0x44, 0x8c, 0x67, 0x68,
	0x1f, 0x6f, 0x64, 0x70, 0xe1, 0xd2, 0x71, 0x68, 0x18, 0x6a, 0x0a, 0x5e, 0xc3, 0x95, 0xed, 0x31,
	0x6a, 0xbb, 0x6f, 0x8b, 0x60, 0xee, 0x53, 0x57, 0x23, 0xd8, 0x83, 0x0e, 0xa3, 0x8b, 0x92, 0xb5,
	0x0b, 0xaa, 0xc7, 0xa3, 0x2f, 0x9e, 0x61, 0x58, 0x5b, 0x50, 0x59, 0xb4, 0x96, 0x4e, 0xd6, 0x81,
	0x26, 0x34, 0x32, 0x3e, 0xbc, 0xec, 0xbd, 0xbd, 0xf9, 0x67, 0x81, 0xa1, 0xe0, 0x04, 0x9a, 0xf9,
	0xa5, 0xbf, 0x3f, 0x0c, 0x2b, 0xc2, 0xa9, 0xcc, 0x50, 0xde, 0x5b, 0xf9, 0xf3, 0x4f, 0x7f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x7b, 0xad, 0xc9, 0x62, 0x11, 0x02, 0x00, 0x00,
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
	Join(ctx context.Context, in *Node, opts ...grpc.CallOption) (*JoinResp, error)
	Leave(ctx context.Context, in *Node, opts ...grpc.CallOption) (*LeaveResp, error)
}

type raftCacheClient struct {
	cc *grpc.ClientConn
}

func NewRaftCacheClient(cc *grpc.ClientConn) RaftCacheClient {
	return &raftCacheClient{cc}
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
	Join(context.Context, *Node) (*JoinResp, error)
	Leave(context.Context, *Node) (*LeaveResp, error)
}

// UnimplementedRaftCacheServer can be embedded to have forward compatible implementations.
type UnimplementedRaftCacheServer struct {
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
