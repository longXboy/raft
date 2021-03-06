// Code generated by protoc-gen-gogo.
// source: raftRPC.proto
// DO NOT EDIT!

/*
	Package raft is a generated protocol buffer package.

	It is generated from these files:
		raftRPC.proto

	It has these top-level messages:
		VoteRequest
		VoteReply
		AppendRequest
		AppendReply
		Entry
*/
package raft

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type VoteRequest struct {
	Term         uint64 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	CandidateId  []byte `protobuf:"bytes,2,opt,name=candidateId,proto3" json:"candidateId,omitempty"`
	LastLogIndex uint64 `protobuf:"varint,3,opt,name=lastLogIndex,proto3" json:"lastLogIndex,omitempty"`
	LastLogTerm  uint64 `protobuf:"varint,4,opt,name=lastLogTerm,proto3" json:"lastLogTerm,omitempty"`
}

func (m *VoteRequest) Reset()                    { *m = VoteRequest{} }
func (m *VoteRequest) String() string            { return proto.CompactTextString(m) }
func (*VoteRequest) ProtoMessage()               {}
func (*VoteRequest) Descriptor() ([]byte, []int) { return fileDescriptorRaftRPC, []int{0} }

func (m *VoteRequest) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *VoteRequest) GetCandidateId() []byte {
	if m != nil {
		return m.CandidateId
	}
	return nil
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

type VoteReply struct {
	Term        uint64 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	VoteGranted bool   `protobuf:"varint,2,opt,name=voteGranted,proto3" json:"voteGranted,omitempty"`
}

func (m *VoteReply) Reset()                    { *m = VoteReply{} }
func (m *VoteReply) String() string            { return proto.CompactTextString(m) }
func (*VoteReply) ProtoMessage()               {}
func (*VoteReply) Descriptor() ([]byte, []int) { return fileDescriptorRaftRPC, []int{1} }

func (m *VoteReply) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *VoteReply) GetVoteGranted() bool {
	if m != nil {
		return m.VoteGranted
	}
	return false
}

type AppendRequest struct {
	Term         uint64   `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	LeaderId     []byte   `protobuf:"bytes,2,opt,name=leaderId,proto3" json:"leaderId,omitempty"`
	PreLogIndex  uint64   `protobuf:"varint,3,opt,name=preLogIndex,proto3" json:"preLogIndex,omitempty"`
	PreLogTerm   uint64   `protobuf:"varint,4,opt,name=preLogTerm,proto3" json:"preLogTerm,omitempty"`
	Entries      []*Entry `protobuf:"bytes,5,rep,name=entries" json:"entries,omitempty"`
	LeaderCommit uint64   `protobuf:"varint,6,opt,name=leaderCommit,proto3" json:"leaderCommit,omitempty"`
}

func (m *AppendRequest) Reset()                    { *m = AppendRequest{} }
func (m *AppendRequest) String() string            { return proto.CompactTextString(m) }
func (*AppendRequest) ProtoMessage()               {}
func (*AppendRequest) Descriptor() ([]byte, []int) { return fileDescriptorRaftRPC, []int{2} }

func (m *AppendRequest) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *AppendRequest) GetLeaderId() []byte {
	if m != nil {
		return m.LeaderId
	}
	return nil
}

func (m *AppendRequest) GetPreLogIndex() uint64 {
	if m != nil {
		return m.PreLogIndex
	}
	return 0
}

func (m *AppendRequest) GetPreLogTerm() uint64 {
	if m != nil {
		return m.PreLogTerm
	}
	return 0
}

func (m *AppendRequest) GetEntries() []*Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func (m *AppendRequest) GetLeaderCommit() uint64 {
	if m != nil {
		return m.LeaderCommit
	}
	return 0
}

type AppendReply struct {
	Term    uint64 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	Success bool   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
}

func (m *AppendReply) Reset()                    { *m = AppendReply{} }
func (m *AppendReply) String() string            { return proto.CompactTextString(m) }
func (*AppendReply) ProtoMessage()               {}
func (*AppendReply) Descriptor() ([]byte, []int) { return fileDescriptorRaftRPC, []int{3} }

func (m *AppendReply) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *AppendReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type Entry struct {
	Term     uint64 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	Index    uint64 `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	DataType uint32 `protobuf:"varint,3,opt,name=dataType,proto3" json:"dataType,omitempty"`
	Data     []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptorRaftRPC, []int{4} }

func (m *Entry) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *Entry) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Entry) GetDataType() uint32 {
	if m != nil {
		return m.DataType
	}
	return 0
}

func (m *Entry) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*VoteRequest)(nil), "raft.VoteRequest")
	proto.RegisterType((*VoteReply)(nil), "raft.VoteReply")
	proto.RegisterType((*AppendRequest)(nil), "raft.AppendRequest")
	proto.RegisterType((*AppendReply)(nil), "raft.AppendReply")
	proto.RegisterType((*Entry)(nil), "raft.Entry")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RaftRPC service

type RaftRPCClient interface {
	// Sends a vote request
	Vote(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*VoteReply, error)
	Append(ctx context.Context, in *AppendRequest, opts ...grpc.CallOption) (*AppendReply, error)
}

type raftRPCClient struct {
	cc *grpc.ClientConn
}

func NewRaftRPCClient(cc *grpc.ClientConn) RaftRPCClient {
	return &raftRPCClient{cc}
}

func (c *raftRPCClient) Vote(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*VoteReply, error) {
	out := new(VoteReply)
	err := grpc.Invoke(ctx, "/raft.RaftRPC/vote", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftRPCClient) Append(ctx context.Context, in *AppendRequest, opts ...grpc.CallOption) (*AppendReply, error) {
	out := new(AppendReply)
	err := grpc.Invoke(ctx, "/raft.RaftRPC/append", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RaftRPC service

type RaftRPCServer interface {
	// Sends a vote request
	Vote(context.Context, *VoteRequest) (*VoteReply, error)
	Append(context.Context, *AppendRequest) (*AppendReply, error)
}

func RegisterRaftRPCServer(s *grpc.Server, srv RaftRPCServer) {
	s.RegisterService(&_RaftRPC_serviceDesc, srv)
}

func _RaftRPC_Vote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftRPCServer).Vote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raft.RaftRPC/Vote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftRPCServer).Vote(ctx, req.(*VoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftRPC_Append_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftRPCServer).Append(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raft.RaftRPC/Append",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftRPCServer).Append(ctx, req.(*AppendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RaftRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "raft.RaftRPC",
	HandlerType: (*RaftRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "vote",
			Handler:    _RaftRPC_Vote_Handler,
		},
		{
			MethodName: "append",
			Handler:    _RaftRPC_Append_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "raftRPC.proto",
}

func (m *VoteRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VoteRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Term != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRaftRPC(dAtA, i, uint64(m.Term))
	}
	if len(m.CandidateId) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRaftRPC(dAtA, i, uint64(len(m.CandidateId)))
		i += copy(dAtA[i:], m.CandidateId)
	}
	if m.LastLogIndex != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintRaftRPC(dAtA, i, uint64(m.LastLogIndex))
	}
	if m.LastLogTerm != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintRaftRPC(dAtA, i, uint64(m.LastLogTerm))
	}
	return i, nil
}

func (m *VoteReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VoteReply) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Term != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRaftRPC(dAtA, i, uint64(m.Term))
	}
	if m.VoteGranted {
		dAtA[i] = 0x10
		i++
		if m.VoteGranted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *AppendRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AppendRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Term != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRaftRPC(dAtA, i, uint64(m.Term))
	}
	if len(m.LeaderId) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRaftRPC(dAtA, i, uint64(len(m.LeaderId)))
		i += copy(dAtA[i:], m.LeaderId)
	}
	if m.PreLogIndex != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintRaftRPC(dAtA, i, uint64(m.PreLogIndex))
	}
	if m.PreLogTerm != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintRaftRPC(dAtA, i, uint64(m.PreLogTerm))
	}
	if len(m.Entries) > 0 {
		for _, msg := range m.Entries {
			dAtA[i] = 0x2a
			i++
			i = encodeVarintRaftRPC(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.LeaderCommit != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintRaftRPC(dAtA, i, uint64(m.LeaderCommit))
	}
	return i, nil
}

func (m *AppendReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AppendReply) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Term != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRaftRPC(dAtA, i, uint64(m.Term))
	}
	if m.Success {
		dAtA[i] = 0x10
		i++
		if m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *Entry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Entry) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Term != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRaftRPC(dAtA, i, uint64(m.Term))
	}
	if m.Index != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintRaftRPC(dAtA, i, uint64(m.Index))
	}
	if m.DataType != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintRaftRPC(dAtA, i, uint64(m.DataType))
	}
	if len(m.Data) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintRaftRPC(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	return i, nil
}

func encodeFixed64RaftRPC(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32RaftRPC(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintRaftRPC(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *VoteRequest) Size() (n int) {
	var l int
	_ = l
	if m.Term != 0 {
		n += 1 + sovRaftRPC(uint64(m.Term))
	}
	l = len(m.CandidateId)
	if l > 0 {
		n += 1 + l + sovRaftRPC(uint64(l))
	}
	if m.LastLogIndex != 0 {
		n += 1 + sovRaftRPC(uint64(m.LastLogIndex))
	}
	if m.LastLogTerm != 0 {
		n += 1 + sovRaftRPC(uint64(m.LastLogTerm))
	}
	return n
}

func (m *VoteReply) Size() (n int) {
	var l int
	_ = l
	if m.Term != 0 {
		n += 1 + sovRaftRPC(uint64(m.Term))
	}
	if m.VoteGranted {
		n += 2
	}
	return n
}

func (m *AppendRequest) Size() (n int) {
	var l int
	_ = l
	if m.Term != 0 {
		n += 1 + sovRaftRPC(uint64(m.Term))
	}
	l = len(m.LeaderId)
	if l > 0 {
		n += 1 + l + sovRaftRPC(uint64(l))
	}
	if m.PreLogIndex != 0 {
		n += 1 + sovRaftRPC(uint64(m.PreLogIndex))
	}
	if m.PreLogTerm != 0 {
		n += 1 + sovRaftRPC(uint64(m.PreLogTerm))
	}
	if len(m.Entries) > 0 {
		for _, e := range m.Entries {
			l = e.Size()
			n += 1 + l + sovRaftRPC(uint64(l))
		}
	}
	if m.LeaderCommit != 0 {
		n += 1 + sovRaftRPC(uint64(m.LeaderCommit))
	}
	return n
}

func (m *AppendReply) Size() (n int) {
	var l int
	_ = l
	if m.Term != 0 {
		n += 1 + sovRaftRPC(uint64(m.Term))
	}
	if m.Success {
		n += 2
	}
	return n
}

func (m *Entry) Size() (n int) {
	var l int
	_ = l
	if m.Term != 0 {
		n += 1 + sovRaftRPC(uint64(m.Term))
	}
	if m.Index != 0 {
		n += 1 + sovRaftRPC(uint64(m.Index))
	}
	if m.DataType != 0 {
		n += 1 + sovRaftRPC(uint64(m.DataType))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovRaftRPC(uint64(l))
	}
	return n
}

func sovRaftRPC(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRaftRPC(x uint64) (n int) {
	return sovRaftRPC(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VoteRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaftRPC
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: VoteRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VoteRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Term", wireType)
			}
			m.Term = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftRPC
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Term |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CandidateId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftRPC
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRaftRPC
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CandidateId = append(m.CandidateId[:0], dAtA[iNdEx:postIndex]...)
			if m.CandidateId == nil {
				m.CandidateId = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastLogIndex", wireType)
			}
			m.LastLogIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftRPC
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastLogIndex |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastLogTerm", wireType)
			}
			m.LastLogTerm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftRPC
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastLogTerm |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRaftRPC(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaftRPC
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *VoteReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaftRPC
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: VoteReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VoteReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Term", wireType)
			}
			m.Term = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftRPC
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Term |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteGranted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftRPC
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.VoteGranted = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRaftRPC(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaftRPC
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AppendRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaftRPC
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AppendRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AppendRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Term", wireType)
			}
			m.Term = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftRPC
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Term |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeaderId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftRPC
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRaftRPC
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LeaderId = append(m.LeaderId[:0], dAtA[iNdEx:postIndex]...)
			if m.LeaderId == nil {
				m.LeaderId = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreLogIndex", wireType)
			}
			m.PreLogIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftRPC
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PreLogIndex |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreLogTerm", wireType)
			}
			m.PreLogTerm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftRPC
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PreLogTerm |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftRPC
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaftRPC
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Entries = append(m.Entries, &Entry{})
			if err := m.Entries[len(m.Entries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeaderCommit", wireType)
			}
			m.LeaderCommit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftRPC
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LeaderCommit |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRaftRPC(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaftRPC
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AppendReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaftRPC
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AppendReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AppendReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Term", wireType)
			}
			m.Term = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftRPC
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Term |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftRPC
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Success = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRaftRPC(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaftRPC
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Entry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaftRPC
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Entry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Entry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Term", wireType)
			}
			m.Term = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftRPC
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Term |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftRPC
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataType", wireType)
			}
			m.DataType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftRPC
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DataType |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftRPC
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRaftRPC
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRaftRPC(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaftRPC
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRaftRPC(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRaftRPC
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRaftRPC
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRaftRPC
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthRaftRPC
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRaftRPC
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipRaftRPC(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthRaftRPC = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRaftRPC   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("raftRPC.proto", fileDescriptorRaftRPC) }

var fileDescriptorRaftRPC = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x4e, 0xea, 0x40,
	0x14, 0xc6, 0x19, 0x28, 0x7f, 0xee, 0x29, 0xe4, 0x5e, 0xe6, 0xde, 0x45, 0xc3, 0xa2, 0x69, 0x9a,
	0xdc, 0x84, 0x85, 0x61, 0x81, 0x4b, 0x57, 0x48, 0x8c, 0x21, 0x71, 0x61, 0x26, 0xc4, 0xfd, 0x48,
	0x8f, 0xa6, 0xb1, 0xb4, 0x75, 0x3a, 0x18, 0xfb, 0x08, 0xbe, 0x81, 0x8f, 0xe4, 0xc2, 0x85, 0x8f,
	0x60, 0xf0, 0x45, 0xcc, 0x9c, 0x02, 0x19, 0x8c, 0xb8, 0x9b, 0xef, 0x6b, 0xcf, 0x9c, 0xf3, 0xfb,
	0xe6, 0x40, 0x4f, 0xc9, 0x1b, 0x2d, 0x2e, 0xa7, 0xa3, 0x5c, 0x65, 0x3a, 0xe3, 0x8e, 0x91, 0xe1,
	0x13, 0x03, 0xf7, 0x2a, 0xd3, 0x28, 0xf0, 0x7e, 0x85, 0x85, 0xe6, 0x1c, 0x1c, 0x8d, 0x6a, 0xe9,
	0xb1, 0x80, 0x0d, 0x1d, 0x41, 0x67, 0x1e, 0x80, 0xbb, 0x90, 0x69, 0x14, 0x47, 0x52, 0xe3, 0x2c,
	0xf2, 0xea, 0x01, 0x1b, 0x76, 0x85, 0x6d, 0xf1, 0x10, 0xba, 0x89, 0x2c, 0xf4, 0x45, 0x76, 0x3b,
	0x4b, 0x23, 0x7c, 0xf4, 0x1a, 0x54, 0xbd, 0xe7, 0x99, 0x5b, 0x36, 0x7a, 0x6e, 0x1a, 0x38, 0xf4,
	0x8b, 0x6d, 0x85, 0x13, 0xf8, 0x55, 0x8d, 0x92, 0x27, 0xe5, 0xa1, 0x41, 0x1e, 0x32, 0x8d, 0xe7,
	0x4a, 0xa6, 0x1a, 0xab, 0x41, 0x3a, 0xc2, 0xb6, 0xc2, 0x57, 0x06, 0xbd, 0x49, 0x9e, 0x63, 0x1a,
	0xfd, 0x04, 0x34, 0x80, 0x4e, 0x82, 0x32, 0x42, 0xb5, 0xa3, 0xd9, 0x69, 0xd3, 0x23, 0x57, 0xf8,
	0x85, 0xc4, 0xb6, 0xb8, 0x0f, 0x50, 0x49, 0x8b, 0xc3, 0x72, 0xf8, 0x7f, 0x68, 0x63, 0xaa, 0x55,
	0x8c, 0x85, 0xd7, 0x0c, 0x1a, 0x43, 0x77, 0xec, 0x8e, 0x4c, 0xd4, 0xa3, 0xb3, 0x54, 0xab, 0x52,
	0x6c, 0xbf, 0x51, 0x66, 0xd4, 0x74, 0x9a, 0x2d, 0x97, 0xb1, 0xf6, 0x5a, 0x9b, 0xcc, 0x2c, 0x2f,
	0x3c, 0x01, 0x77, 0x4b, 0x73, 0x28, 0x13, 0x0f, 0xda, 0xc5, 0x6a, 0xb1, 0xc0, 0xa2, 0xd8, 0xe4,
	0xb1, 0x95, 0xa1, 0x84, 0x26, 0xb5, 0xfc, 0xb6, 0xec, 0x1f, 0x34, 0x63, 0x02, 0xac, 0x93, 0x59,
	0x09, 0x13, 0x4c, 0x24, 0xb5, 0x9c, 0x97, 0x39, 0x12, 0x79, 0x4f, 0xec, 0xb4, 0xb9, 0xc5, 0x9c,
	0x09, 0xb8, 0x2b, 0xe8, 0x3c, 0xbe, 0x83, 0xb6, 0xa8, 0x96, 0x8a, 0x1f, 0x81, 0x63, 0x1e, 0x82,
	0xf7, 0x2b, 0x58, 0x6b, 0xa7, 0x06, 0xbf, 0x6d, 0x2b, 0x4f, 0xca, 0xb0, 0xc6, 0xc7, 0xd0, 0x92,
	0x04, 0xc6, 0xff, 0x56, 0x1f, 0xf7, 0x1e, 0x6d, 0xd0, 0xdf, 0x37, 0xa9, 0xe6, 0xf4, 0xcf, 0xcb,
	0xda, 0x67, 0x6f, 0x6b, 0x9f, 0xbd, 0xaf, 0x7d, 0xf6, 0xfc, 0xe1, 0xd7, 0xae, 0x5b, 0xb4, 0xc9,
	0xc7, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x37, 0xd3, 0x34, 0xda, 0x02, 0x00, 0x00,
}
