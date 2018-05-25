// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vote_request.proto

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type VoteRequest struct {
	Term          uint64 `protobuf:"varint,1,opt,name=Term" json:"Term,omitempty"`
	LastLogIndex  uint64 `protobuf:"varint,2,opt,name=LastLogIndex" json:"LastLogIndex,omitempty"`
	LastLogTerm   uint64 `protobuf:"varint,3,opt,name=LastLogTerm" json:"LastLogTerm,omitempty"`
	CandidateName string `protobuf:"bytes,4,opt,name=CandidateName" json:"CandidateName,omitempty"`
	Host          string `protobuf:"bytes,5,opt,name=Host" json:"Host,omitempty"`
}

func (m *VoteRequest) Reset()                    { *m = VoteRequest{} }
func (m *VoteRequest) String() string            { return proto1.CompactTextString(m) }
func (*VoteRequest) ProtoMessage()               {}
func (*VoteRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *VoteRequest) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *VoteRequest) GetLastLogIndex() uint64 {
	if m != nil {
		return m.LastLogIndex
	}
	return 0
}

func (m *VoteRequest) GetLastLogTerm() uint64 {
	if m != nil {
		return m.LastLogTerm
	}
	return 0
}

func (m *VoteRequest) GetCandidateName() string {
	if m != nil {
		return m.CandidateName
	}
	return ""
}

func (m *VoteRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

type VoteResponse struct {
	Term        uint64 `protobuf:"varint,1,opt,name=Term" json:"Term,omitempty"`
	VoteGranted bool   `protobuf:"varint,2,opt,name=VoteGranted" json:"VoteGranted,omitempty"`
}

func (m *VoteResponse) Reset()                    { *m = VoteResponse{} }
func (m *VoteResponse) String() string            { return proto1.CompactTextString(m) }
func (*VoteResponse) ProtoMessage()               {}
func (*VoteResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *VoteResponse) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *VoteResponse) GetVoteGranted() bool {
	if m != nil {
		return m.VoteGranted
	}
	return false
}

func init() {
	proto1.RegisterType((*VoteRequest)(nil), "proto.VoteRequest")
	proto1.RegisterType((*VoteResponse)(nil), "proto.VoteResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RequestVote service

type RequestVoteClient interface {
	RequestVoteMe(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*VoteResponse, error)
}

type requestVoteClient struct {
	cc *grpc.ClientConn
}

func NewRequestVoteClient(cc *grpc.ClientConn) RequestVoteClient {
	return &requestVoteClient{cc}
}

func (c *requestVoteClient) RequestVoteMe(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*VoteResponse, error) {
	out := new(VoteResponse)
	err := grpc.Invoke(ctx, "/proto.RequestVote/RequestVoteMe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RequestVote service

type RequestVoteServer interface {
	RequestVoteMe(context.Context, *VoteRequest) (*VoteResponse, error)
}

func RegisterRequestVoteServer(s *grpc.Server, srv RequestVoteServer) {
	s.RegisterService(&_RequestVote_serviceDesc, srv)
}

func _RequestVote_RequestVoteMe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestVoteServer).RequestVoteMe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RequestVote/RequestVoteMe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestVoteServer).RequestVoteMe(ctx, req.(*VoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RequestVote_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.RequestVote",
	HandlerType: (*RequestVoteServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestVoteMe",
			Handler:    _RequestVote_RequestVoteMe_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vote_request.proto",
}

func init() { proto1.RegisterFile("vote_request.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0xcb, 0x2f, 0x49,
	0x8d, 0x2f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x05, 0x53, 0x4a, 0x0b, 0x19, 0xb9, 0xb8, 0xc3, 0xf2, 0x4b, 0x52, 0x83, 0x20, 0x92, 0x42, 0x42,
	0x5c, 0x2c, 0x21, 0xa9, 0x45, 0xb9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0x60, 0xb6, 0x90,
	0x12, 0x17, 0x8f, 0x4f, 0x62, 0x71, 0x89, 0x4f, 0x7e, 0xba, 0x67, 0x5e, 0x4a, 0x6a, 0x85, 0x04,
	0x13, 0x58, 0x0e, 0x45, 0x4c, 0x48, 0x81, 0x8b, 0x1b, 0xca, 0x07, 0x6b, 0x67, 0x06, 0x2b, 0x41,
	0x16, 0x12, 0x52, 0xe1, 0xe2, 0x75, 0x4e, 0xcc, 0x4b, 0xc9, 0x4c, 0x49, 0x2c, 0x49, 0xf5, 0x4b,
	0xcc, 0x4d, 0x95, 0x60, 0x51, 0x60, 0xd4, 0xe0, 0x0c, 0x42, 0x15, 0x04, 0xd9, 0xef, 0x91, 0x5f,
	0x5c, 0x22, 0xc1, 0x0a, 0x96, 0x04, 0xb3, 0x95, 0x5c, 0xb8, 0x78, 0x20, 0x4e, 0x2c, 0x2e, 0xc8,
	0xcf, 0x2b, 0x4e, 0xc5, 0xea, 0x46, 0x05, 0x88, 0x37, 0xdc, 0x8b, 0x12, 0xf3, 0x4a, 0x52, 0x53,
	0xc0, 0x4e, 0xe4, 0x08, 0x42, 0x16, 0x32, 0xf2, 0xe4, 0xe2, 0x86, 0x7a, 0x12, 0x24, 0x2a, 0x64,
	0xc5, 0xc5, 0x8b, 0xc4, 0xf5, 0x4d, 0x15, 0x12, 0x82, 0x04, 0x8c, 0x1e, 0x52, 0x68, 0x48, 0x09,
	0xa3, 0x88, 0x41, 0xac, 0x57, 0x62, 0x48, 0x62, 0x03, 0x8b, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x09, 0xc2, 0x8b, 0xbb, 0x58, 0x01, 0x00, 0x00,
}
