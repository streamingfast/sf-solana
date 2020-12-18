// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dfuse/solana/serumhist/v1/serumhist.proto

package pbaccounthist

import (
	context "context"
	fmt "fmt"
	math "math"

	v1 "github.com/dfuse-io/dfuse-solana/pb/dfuse/solana/serum/v1"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetFillsRequest struct {
	Trader               string   `protobuf:"bytes,1,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Market               string   `protobuf:"bytes,2,opt,name=market,proto3" json:"market,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFillsRequest) Reset()         { *m = GetFillsRequest{} }
func (m *GetFillsRequest) String() string { return proto.CompactTextString(m) }
func (*GetFillsRequest) ProtoMessage()    {}
func (*GetFillsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c18832c9aa9279e6, []int{0}
}

func (m *GetFillsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFillsRequest.Unmarshal(m, b)
}
func (m *GetFillsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFillsRequest.Marshal(b, m, deterministic)
}
func (m *GetFillsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFillsRequest.Merge(m, src)
}
func (m *GetFillsRequest) XXX_Size() int {
	return xxx_messageInfo_GetFillsRequest.Size(m)
}
func (m *GetFillsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFillsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFillsRequest proto.InternalMessageInfo

func (m *GetFillsRequest) GetPubkey() string {
	if m != nil {
		return m.Trader
	}
	return ""
}

func (m *GetFillsRequest) GetMarket() string {
	if m != nil {
		return m.Market
	}
	return ""
}

type FillsResponse struct {
	Fill                 []*v1.Fill `protobuf:"bytes,1,rep,name=fill,proto3" json:"fill,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *FillsResponse) Reset()         { *m = FillsResponse{} }
func (m *FillsResponse) String() string { return proto.CompactTextString(m) }
func (*FillsResponse) ProtoMessage()    {}
func (*FillsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c18832c9aa9279e6, []int{1}
}

func (m *FillsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FillsResponse.Unmarshal(m, b)
}
func (m *FillsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FillsResponse.Marshal(b, m, deterministic)
}
func (m *FillsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FillsResponse.Merge(m, src)
}
func (m *FillsResponse) XXX_Size() int {
	return xxx_messageInfo_FillsResponse.Size(m)
}
func (m *FillsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FillsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FillsResponse proto.InternalMessageInfo

func (m *FillsResponse) GetFill() []*v1.Fill {
	if m != nil {
		return m.Fill
	}
	return nil
}

type Checkpoint struct {
	LastWrittenSlotNum   uint64   `protobuf:"varint,1,opt,name=last_written_slot_num,json=lastWrittenSlotNum,proto3" json:"last_written_slot_num,omitempty"`
	LastWrittenLostId    string   `protobuf:"bytes,2,opt,name=last_written_lost_id,json=lastWrittenLostId,proto3" json:"last_written_lost_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Checkpoint) Reset()         { *m = Checkpoint{} }
func (m *Checkpoint) String() string { return proto.CompactTextString(m) }
func (*Checkpoint) ProtoMessage()    {}
func (*Checkpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_c18832c9aa9279e6, []int{2}
}

func (m *Checkpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Checkpoint.Unmarshal(m, b)
}
func (m *Checkpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Checkpoint.Marshal(b, m, deterministic)
}
func (m *Checkpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Checkpoint.Merge(m, src)
}
func (m *Checkpoint) XXX_Size() int {
	return xxx_messageInfo_Checkpoint.Size(m)
}
func (m *Checkpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Checkpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Checkpoint proto.InternalMessageInfo

func (m *Checkpoint) GetLastWrittenSlotNum() uint64 {
	if m != nil {
		return m.LastWrittenSlotNum
	}
	return 0
}

func (m *Checkpoint) GetLastWrittenLostId() string {
	if m != nil {
		return m.LastWrittenLostId
	}
	return ""
}

func init() {
	proto.RegisterType((*GetFillsRequest)(nil), "dfuse.eosio.accounthist.v1.GetFillsRequest")
	proto.RegisterType((*FillsResponse)(nil), "dfuse.eosio.accounthist.v1.FillsResponse")
	proto.RegisterType((*Checkpoint)(nil), "dfuse.eosio.accounthist.v1.Checkpoint")
}

func init() {
	proto.RegisterFile("dfuse/solana/serumhist/v1/serumhist.proto", fileDescriptor_c18832c9aa9279e6)
}

var fileDescriptor_c18832c9aa9279e6 = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcd, 0x4b, 0xeb, 0x40,
	0x14, 0xc5, 0xe9, 0x7b, 0xa5, 0xbc, 0x37, 0x2a, 0xe2, 0xa0, 0x52, 0xea, 0xa6, 0x74, 0xd5, 0x22,
	0xce, 0x90, 0xba, 0x74, 0xe3, 0x07, 0xf8, 0x81, 0x1f, 0x8b, 0x74, 0x21, 0xb8, 0x09, 0x49, 0x3a,
	0xb5, 0x43, 0x27, 0xb9, 0x31, 0xf7, 0x4e, 0x4a, 0xff, 0x7b, 0xc9, 0x4c, 0xc4, 0x56, 0xd1, 0xdd,
	0x9c, 0x7b, 0xef, 0xef, 0x70, 0x38, 0xc3, 0x46, 0xd3, 0x99, 0x45, 0x25, 0x11, 0x4c, 0x9c, 0xc7,
	0x12, 0x55, 0x69, 0xb3, 0xb9, 0x46, 0x92, 0x55, 0xf0, 0x29, 0x44, 0x51, 0x02, 0x01, 0xef, 0xb9,
	0x53, 0xa1, 0x00, 0x35, 0x88, 0x38, 0x4d, 0xc1, 0xe6, 0xe4, 0xd6, 0x55, 0xd0, 0xeb, 0x7f, 0xb7,
	0xa9, 0x2d, 0x4a, 0x58, 0xa2, 0xa7, 0x07, 0x17, 0x6c, 0xf7, 0x46, 0xd1, 0xb5, 0x36, 0x06, 0x43,
	0xf5, 0x66, 0x15, 0x12, 0x3f, 0x64, 0x9d, 0xc2, 0x26, 0x0b, 0xb5, 0xea, 0xb6, 0xfa, 0xad, 0xe1,
	0xff, 0xb0, 0x51, 0xf5, 0x3c, 0x8b, 0xcb, 0x85, 0xa2, 0xee, 0x1f, 0x3f, 0xf7, 0x6a, 0x70, 0xce,
	0x76, 0x1a, 0x1e, 0x0b, 0xc8, 0x51, 0x71, 0xc9, 0xda, 0x33, 0x6d, 0x4c, 0xb7, 0xd5, 0xff, 0x3b,
	0xdc, 0x1a, 0x1f, 0x09, 0x1f, 0xd0, 0x87, 0x10, 0x2e, 0x84, 0xa8, 0x02, 0x51, 0x33, 0xa1, 0x3b,
	0x1c, 0x14, 0x8c, 0x5d, 0xcd, 0x55, 0xba, 0x28, 0x40, 0xe7, 0xc4, 0x03, 0x76, 0x60, 0x62, 0xa4,
	0x68, 0x59, 0x6a, 0x22, 0x95, 0x47, 0x68, 0x80, 0xa2, 0xdc, 0x66, 0x2e, 0x4e, 0x3b, 0xe4, 0xf5,
	0xf2, 0xd9, 0xef, 0x26, 0x06, 0xe8, 0xc9, 0x66, 0x5c, 0xb2, 0xfd, 0x0d, 0xc4, 0x00, 0x52, 0xa4,
	0xa7, 0x4d, 0xd0, 0xbd, 0x35, 0xe2, 0x01, 0x90, 0xee, 0xa6, 0xe3, 0x92, 0x6d, 0x4f, 0xea, 0x20,
	0xb7, 0x1a, 0x09, 0xca, 0x15, 0x4f, 0xd8, 0xbf, 0x8f, 0x1a, 0xf8, 0xb1, 0xf8, 0xb9, 0x51, 0xf1,
	0xa5, 0xac, 0xde, 0xe8, 0xb7, 0xe3, 0x8d, 0x5a, 0x2e, 0x1f, 0x5f, 0xee, 0x5f, 0x35, 0xcd, 0x6d,
	0x22, 0x52, 0xc8, 0xa4, 0xc3, 0x4e, 0x34, 0x34, 0x0f, 0xc7, 0xcb, 0x22, 0xf1, 0x52, 0x7a, 0xb9,
	0x66, 0x27, 0xab, 0xe0, 0xac, 0x48, 0xd6, 0x06, 0x49, 0xc7, 0x7d, 0xe0, 0xe9, 0x7b, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x4e, 0xf5, 0x85, 0xba, 0x2b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SerumHistoryClient is the client API for SerumHistory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SerumHistoryClient interface {
	GetFills(ctx context.Context, in *GetFillsRequest, opts ...grpc.CallOption) (*FillsResponse, error)
}

type serumHistoryClient struct {
	cc *grpc.ClientConn
}

func NewSerumHistoryClient(cc *grpc.ClientConn) SerumHistoryClient {
	return &serumHistoryClient{cc}
}

func (c *serumHistoryClient) GetFills(ctx context.Context, in *GetFillsRequest, opts ...grpc.CallOption) (*FillsResponse, error) {
	out := new(FillsResponse)
	err := c.cc.Invoke(ctx, "/dfuse.eosio.accounthist.v1.SerumHistory/GetFills", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SerumHistoryServer is the server API for SerumHistory service.
type SerumHistoryServer interface {
	GetFills(context.Context, *GetFillsRequest) (*FillsResponse, error)
}

// UnimplementedSerumHistoryServer can be embedded to have forward compatible implementations.
type UnimplementedSerumHistoryServer struct {
}

func (*UnimplementedSerumHistoryServer) GetFills(ctx context.Context, req *GetFillsRequest) (*FillsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFills not implemented")
}

func RegisterSerumHistoryServer(s *grpc.Server, srv SerumHistoryServer) {
	s.RegisterService(&_SerumHistory_serviceDesc, srv)
}

func _SerumHistory_GetFills_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFillsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SerumHistoryServer).GetFills(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dfuse.eosio.accounthist.v1.SerumHistory/GetFills",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SerumHistoryServer).GetFills(ctx, req.(*GetFillsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SerumHistory_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dfuse.eosio.accounthist.v1.SerumHistory",
	HandlerType: (*SerumHistoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFills",
			Handler:    _SerumHistory_GetFills_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dfuse/solana/serumhist/v1/serumhist.proto",
}
