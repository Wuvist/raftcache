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
	Node_INITIATING   Node_Statuses = 3
	Node_HANDSHAKING  Node_Statuses = 4
)

var Node_Statuses_name = map[int32]string{
	0: "ALONE",
	1: "INGROUP",
	2: "DISCONNECTED",
	3: "INITIATING",
	4: "HANDSHAKING",
}

var Node_Statuses_value = map[string]int32{
	"ALONE":        0,
	"INGROUP":      1,
	"DISCONNECTED": 2,
	"INITIATING":   3,
	"HANDSHAKING":  4,
}

func (x Node_Statuses) String() string {
	return proto.EnumName(Node_Statuses_name, int32(x))
}

func (Node_Statuses) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f886bd89b3f134d3, []int{1, 0}
}

type JoinResp_Results int32

const (
	JoinResp_SUCCESS       JoinResp_Results = 0
	JoinResp_ALREADYJOINED JoinResp_Results = 1
	JoinResp_REJECTED      JoinResp_Results = 2
	JoinResp_PINGFAIL      JoinResp_Results = 3
	JoinResp_TRYLATER      JoinResp_Results = 4
)

var JoinResp_Results_name = map[int32]string{
	0: "SUCCESS",
	1: "ALREADYJOINED",
	2: "REJECTED",
	3: "PINGFAIL",
	4: "TRYLATER",
}

var JoinResp_Results_value = map[string]int32{
	"SUCCESS":       0,
	"ALREADYJOINED": 1,
	"REJECTED":      2,
	"PINGFAIL":      3,
	"TRYLATER":      4,
}

func (x JoinResp_Results) String() string {
	return proto.EnumName(JoinResp_Results_name, int32(x))
}

func (JoinResp_Results) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f886bd89b3f134d3, []int{2, 0}
}

type LeaveResp_Results int32

const (
	LeaveResp_SUCCESS    LeaveResp_Results = 0
	LeaveResp_REJECTED   LeaveResp_Results = 1
	LeaveResp_NOTINGROUP LeaveResp_Results = 2
)

var LeaveResp_Results_name = map[int32]string{
	0: "SUCCESS",
	1: "REJECTED",
	2: "NOTINGROUP",
}

var LeaveResp_Results_value = map[string]int32{
	"SUCCESS":    0,
	"REJECTED":   1,
	"NOTINGROUP": 2,
}

func (x LeaveResp_Results) String() string {
	return proto.EnumName(LeaveResp_Results_name, int32(x))
}

func (LeaveResp_Results) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f886bd89b3f134d3, []int{3, 0}
}

type HandshakeResp_Results int32

const (
	HandshakeResp_SUCCESS  HandshakeResp_Results = 0
	HandshakeResp_PINGFAIL HandshakeResp_Results = 1
	HandshakeResp_TRYLATER HandshakeResp_Results = 2
)

var HandshakeResp_Results_name = map[int32]string{
	0: "SUCCESS",
	1: "PINGFAIL",
	2: "TRYLATER",
}

var HandshakeResp_Results_value = map[string]int32{
	"SUCCESS":  0,
	"PINGFAIL": 1,
	"TRYLATER": 2,
}

func (x HandshakeResp_Results) String() string {
	return proto.EnumName(HandshakeResp_Results_name, int32(x))
}

func (HandshakeResp_Results) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f886bd89b3f134d3, []int{4, 0}
}

type JoinConfirmResp_Results int32

const (
	JoinConfirmResp_SUCCESS JoinConfirmResp_Results = 0
	JoinConfirmResp_FAILED  JoinConfirmResp_Results = 1
)

var JoinConfirmResp_Results_name = map[int32]string{
	0: "SUCCESS",
	1: "FAILED",
}

var JoinConfirmResp_Results_value = map[string]int32{
	"SUCCESS": 0,
	"FAILED":  1,
}

func (x JoinConfirmResp_Results) String() string {
	return proto.EnumName(JoinConfirmResp_Results_name, int32(x))
}

func (JoinConfirmResp_Results) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f886bd89b3f134d3, []int{5, 0}
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
	return fileDescriptor_f886bd89b3f134d3, []int{0}
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
	return fileDescriptor_f886bd89b3f134d3, []int{1}
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
	return fileDescriptor_f886bd89b3f134d3, []int{2}
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

type LeaveResp struct {
	Result               LeaveResp_Results `protobuf:"varint,1,opt,name=Result,proto3,enum=raftcache.LeaveResp_Results" json:"Result,omitempty"`
	Message              string            `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
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

func (m *LeaveResp) GetResult() LeaveResp_Results {
	if m != nil {
		return m.Result
	}
	return LeaveResp_SUCCESS
}

func (m *LeaveResp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type HandshakeResp struct {
	Result               HandshakeResp_Results `protobuf:"varint,1,opt,name=Result,proto3,enum=raftcache.HandshakeResp_Results" json:"Result,omitempty"`
	Message              string                `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *HandshakeResp) Reset()         { *m = HandshakeResp{} }
func (m *HandshakeResp) String() string { return proto.CompactTextString(m) }
func (*HandshakeResp) ProtoMessage()    {}
func (*HandshakeResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f886bd89b3f134d3, []int{4}
}

func (m *HandshakeResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HandshakeResp.Unmarshal(m, b)
}
func (m *HandshakeResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HandshakeResp.Marshal(b, m, deterministic)
}
func (m *HandshakeResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandshakeResp.Merge(m, src)
}
func (m *HandshakeResp) XXX_Size() int {
	return xxx_messageInfo_HandshakeResp.Size(m)
}
func (m *HandshakeResp) XXX_DiscardUnknown() {
	xxx_messageInfo_HandshakeResp.DiscardUnknown(m)
}

var xxx_messageInfo_HandshakeResp proto.InternalMessageInfo

func (m *HandshakeResp) GetResult() HandshakeResp_Results {
	if m != nil {
		return m.Result
	}
	return HandshakeResp_SUCCESS
}

func (m *HandshakeResp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type JoinConfirmResp struct {
	Result               JoinConfirmResp_Results `protobuf:"varint,1,opt,name=Result,proto3,enum=raftcache.JoinConfirmResp_Results" json:"Result,omitempty"`
	Message              string                  `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *JoinConfirmResp) Reset()         { *m = JoinConfirmResp{} }
func (m *JoinConfirmResp) String() string { return proto.CompactTextString(m) }
func (*JoinConfirmResp) ProtoMessage()    {}
func (*JoinConfirmResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f886bd89b3f134d3, []int{5}
}

func (m *JoinConfirmResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinConfirmResp.Unmarshal(m, b)
}
func (m *JoinConfirmResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinConfirmResp.Marshal(b, m, deterministic)
}
func (m *JoinConfirmResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinConfirmResp.Merge(m, src)
}
func (m *JoinConfirmResp) XXX_Size() int {
	return xxx_messageInfo_JoinConfirmResp.Size(m)
}
func (m *JoinConfirmResp) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinConfirmResp.DiscardUnknown(m)
}

var xxx_messageInfo_JoinConfirmResp proto.InternalMessageInfo

func (m *JoinConfirmResp) GetResult() JoinConfirmResp_Results {
	if m != nil {
		return m.Result
	}
	return JoinConfirmResp_SUCCESS
}

func (m *JoinConfirmResp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterEnum("raftcache.Node_Statuses", Node_Statuses_name, Node_Statuses_value)
	proto.RegisterEnum("raftcache.JoinResp_Results", JoinResp_Results_name, JoinResp_Results_value)
	proto.RegisterEnum("raftcache.LeaveResp_Results", LeaveResp_Results_name, LeaveResp_Results_value)
	proto.RegisterEnum("raftcache.HandshakeResp_Results", HandshakeResp_Results_name, HandshakeResp_Results_value)
	proto.RegisterEnum("raftcache.JoinConfirmResp_Results", JoinConfirmResp_Results_name, JoinConfirmResp_Results_value)
	proto.RegisterType((*Empty)(nil), "raftcache.Empty")
	proto.RegisterType((*Node)(nil), "raftcache.Node")
	proto.RegisterType((*JoinResp)(nil), "raftcache.JoinResp")
	proto.RegisterType((*LeaveResp)(nil), "raftcache.LeaveResp")
	proto.RegisterType((*HandshakeResp)(nil), "raftcache.HandshakeResp")
	proto.RegisterType((*JoinConfirmResp)(nil), "raftcache.JoinConfirmResp")
}

func init() { proto.RegisterFile("raftcache.proto", fileDescriptor_f886bd89b3f134d3) }

var fileDescriptor_f886bd89b3f134d3 = []byte{
	// 533 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xcd, 0x6e, 0xd3, 0x40,
	0x14, 0x85, 0x3d, 0xf9, 0xf7, 0x4d, 0x9b, 0x0c, 0x97, 0x2e, 0xac, 0x80, 0x50, 0x34, 0xab, 0x2c,
	0x50, 0xa8, 0xd2, 0x2e, 0x2a, 0xc4, 0xc6, 0x72, 0x4c, 0xe2, 0x60, 0x26, 0xd1, 0xd8, 0x59, 0x74,
	0x69, 0x9a, 0x49, 0x1b, 0x41, 0xe3, 0x28, 0x76, 0x90, 0x78, 0x06, 0x56, 0x88, 0x05, 0x4f, 0xc4,
	0x73, 0xf0, 0x2a, 0x68, 0x9c, 0x1f, 0x5c, 0xa7, 0x96, 0xe8, 0x72, 0xee, 0x9c, 0x99, 0xf9, 0xce,
	0xb9, 0xd7, 0x86, 0xe6, 0x3a, 0x98, 0xc7, 0x37, 0xc1, 0xcd, 0x9d, 0xec, 0xae, 0xd6, 0x61, 0x1c,
	0xa2, 0x7e, 0x28, 0xb0, 0x2a, 0x94, 0xed, 0xfb, 0x55, 0xfc, 0x8d, 0xfd, 0x26, 0x50, 0xe2, 0xe1,
	0x4c, 0xe2, 0x19, 0x94, 0x07, 0xeb, 0x70, 0xb3, 0x32, 0x48, 0x9b, 0x74, 0x74, 0xb1, 0x5d, 0xe0,
	0x2b, 0x00, 0x77, 0x11, 0xc5, 0x72, 0x69, 0xce, 0x66, 0x6b, 0xa3, 0x90, 0x6c, 0xa5, 0x2a, 0x78,
	0x0e, 0x15, 0x2f, 0x0e, 0xe2, 0x4d, 0x64, 0x14, 0xdb, 0xa4, 0xd3, 0xe8, 0x19, 0xdd, 0x7f, 0x8f,
	0xaa, 0x6b, 0xbb, 0xdb, 0x5d, 0x19, 0x89, 0x9d, 0x8e, 0x4d, 0xa1, 0xb6, 0xaf, 0xa1, 0x0e, 0x65,
	0xd3, 0x1d, 0x73, 0x9b, 0x6a, 0x58, 0x87, 0xaa, 0xc3, 0x07, 0x62, 0x3c, 0x9d, 0x50, 0x82, 0x14,
	0x4e, 0xfa, 0x8e, 0x67, 0x8d, 0x39, 0xb7, 0x2d, 0xdf, 0xee, 0xd3, 0x02, 0x36, 0x00, 0x1c, 0xee,
	0xf8, 0x8e, 0xe9, 0x3b, 0x7c, 0x40, 0x8b, 0xd8, 0x84, 0xfa, 0xd0, 0xe4, 0x7d, 0x6f, 0x68, 0x7e,
	0x50, 0x85, 0x12, 0xfb, 0x43, 0xa0, 0x36, 0x0a, 0x17, 0x4b, 0x21, 0xa3, 0x15, 0x5e, 0x40, 0x45,
	0xc8, 0x68, 0xf3, 0x25, 0x4e, 0xcc, 0x34, 0x7a, 0x2f, 0x52, 0x54, 0x7b, 0x51, 0x77, 0xab, 0x88,
	0xc4, 0x4e, 0x8a, 0x06, 0x54, 0x3f, 0xca, 0x28, 0x0a, 0x6e, 0xe5, 0xce, 0xe7, 0x7e, 0x89, 0x6f,
	0x00, 0x92, 0x34, 0x94, 0x21, 0x65, 0xb4, 0xd8, 0xa9, 0xf7, 0x9a, 0x19, 0xa3, 0x22, 0x25, 0x61,
	0x1e, 0x54, 0x77, 0xb7, 0x2b, 0x5f, 0xde, 0xd4, 0xb2, 0x6c, 0xcf, 0xa3, 0x1a, 0x3e, 0x83, 0x53,
	0xd3, 0x15, 0xb6, 0xd9, 0xbf, 0x1e, 0x8d, 0x1d, 0x6e, 0xf7, 0x29, 0xc1, 0x13, 0xa8, 0x09, 0x7b,
	0xb4, 0xb7, 0x79, 0x02, 0xb5, 0x89, 0xc3, 0x07, 0xef, 0x4d, 0xc7, 0xa5, 0x45, 0xb5, 0xf2, 0xc5,
	0xb5, 0x6b, 0xfa, 0xb6, 0xa0, 0x25, 0xf6, 0x83, 0x80, 0xee, 0xca, 0xe0, 0xab, 0x4c, 0x2c, 0x5e,
	0x66, 0x2c, 0xbe, 0x4c, 0xf1, 0x1c, 0x54, 0xff, 0xef, 0x91, 0x5d, 0xe6, 0x20, 0xa7, 0xf9, 0x88,
	0x6a, 0x03, 0x1f, 0xfb, 0xfb, 0x46, 0x15, 0xd8, 0x2f, 0x02, 0xa7, 0xc3, 0x60, 0x39, 0x8b, 0xee,
	0x82, 0xcf, 0x5b, 0xae, 0xab, 0x0c, 0x57, 0x3b, 0xc5, 0xf5, 0x40, 0xf9, 0x04, 0xb6, 0x5e, 0x3e,
	0xdb, 0x21, 0x2d, 0xf2, 0x20, 0xad, 0x02, 0xfb, 0x4e, 0xa0, 0xa9, 0x5a, 0x6d, 0x85, 0xcb, 0xf9,
	0x62, 0x7d, 0x9f, 0xb0, 0xbd, 0xcd, 0xb0, 0xb1, 0xcc, 0x58, 0xa4, 0xb4, 0x4f, 0xa0, 0x63, 0x39,
	0x74, 0x00, 0x15, 0x45, 0xa6, 0x72, 0xeb, 0xfd, 0x2c, 0x80, 0x2e, 0x82, 0x79, 0x6c, 0xa9, 0xb7,
	0xf0, 0x35, 0x94, 0x26, 0x8b, 0xe5, 0x2d, 0xd2, 0xd4, 0xfb, 0xc9, 0xd7, 0xd8, 0x3a, 0xaa, 0x30,
	0x0d, 0xbb, 0x50, 0x52, 0x70, 0x98, 0x9d, 0xb8, 0xd6, 0xf3, 0x47, 0xa6, 0x9a, 0x69, 0x78, 0x0e,
	0xe5, 0x64, 0x00, 0x8e, 0x0f, 0x9c, 0x3d, 0x36, 0x23, 0x4c, 0xc3, 0x2b, 0xd0, 0x0f, 0xad, 0x39,
	0x3e, 0x65, 0xe4, 0x75, 0x90, 0x69, 0xf8, 0x0e, 0xea, 0xa9, 0xe0, 0x8e, 0xcf, 0xb6, 0xf2, 0x13,
	0x66, 0xda, 0xa7, 0x4a, 0xf2, 0x5b, 0xba, 0xf8, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x10, 0x1f, 0x84,
	0x89, 0xa9, 0x04, 0x00, 0x00,
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
	Handshake(ctx context.Context, in *Node, opts ...grpc.CallOption) (*HandshakeResp, error)
	JoinConfirm(ctx context.Context, in *Node, opts ...grpc.CallOption) (*JoinConfirmResp, error)
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

func (c *raftCacheClient) Handshake(ctx context.Context, in *Node, opts ...grpc.CallOption) (*HandshakeResp, error) {
	out := new(HandshakeResp)
	err := c.cc.Invoke(ctx, "/raftcache.RaftCache/Handshake", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftCacheClient) JoinConfirm(ctx context.Context, in *Node, opts ...grpc.CallOption) (*JoinConfirmResp, error) {
	out := new(JoinConfirmResp)
	err := c.cc.Invoke(ctx, "/raftcache.RaftCache/JoinConfirm", in, out, opts...)
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
	Handshake(context.Context, *Node) (*HandshakeResp, error)
	JoinConfirm(context.Context, *Node) (*JoinConfirmResp, error)
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
func (*UnimplementedRaftCacheServer) Handshake(ctx context.Context, req *Node) (*HandshakeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handshake not implemented")
}
func (*UnimplementedRaftCacheServer) JoinConfirm(ctx context.Context, req *Node) (*JoinConfirmResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinConfirm not implemented")
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

func _RaftCache_Handshake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftCacheServer).Handshake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raftcache.RaftCache/Handshake",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftCacheServer).Handshake(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftCache_JoinConfirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftCacheServer).JoinConfirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raftcache.RaftCache/JoinConfirm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftCacheServer).JoinConfirm(ctx, req.(*Node))
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
		{
			MethodName: "Handshake",
			Handler:    _RaftCache_Handshake_Handler,
		},
		{
			MethodName: "JoinConfirm",
			Handler:    _RaftCache_JoinConfirm_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "raftcache.proto",
}
