// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: interchaintxs/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_85130b102faab7ea, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_85130b102faab7ea, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// this line is used by starport scaffolding # 3
type QueryInterchainAccountAddressRequest struct {
	// owner_address is the owner of the interchain account on the controller chain
	OwnerAddress string `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	// interchain_account_id is an identifier of your interchain account from which you want to execute msgs
	InterchainAccountId string `protobuf:"bytes,2,opt,name=interchain_account_id,json=interchainAccountId,proto3" json:"interchain_account_id,omitempty"`
	// connection_id is an IBC connection identifier between Neutron and remote chain
	ConnectionId string `protobuf:"bytes,3,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
}

func (m *QueryInterchainAccountAddressRequest) Reset()         { *m = QueryInterchainAccountAddressRequest{} }
func (m *QueryInterchainAccountAddressRequest) String() string { return proto.CompactTextString(m) }
func (*QueryInterchainAccountAddressRequest) ProtoMessage()    {}
func (*QueryInterchainAccountAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_85130b102faab7ea, []int{2}
}
func (m *QueryInterchainAccountAddressRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryInterchainAccountAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryInterchainAccountAddressRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryInterchainAccountAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryInterchainAccountAddressRequest.Merge(m, src)
}
func (m *QueryInterchainAccountAddressRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryInterchainAccountAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryInterchainAccountAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryInterchainAccountAddressRequest proto.InternalMessageInfo

// Query response for an interchain account address
type QueryInterchainAccountAddressResponse struct {
	// The corresponding interchain account address on the host chain
	InterchainAccountAddress string `protobuf:"bytes,1,opt,name=interchain_account_address,json=interchainAccountAddress,proto3" json:"interchain_account_address,omitempty"`
}

func (m *QueryInterchainAccountAddressResponse) Reset()         { *m = QueryInterchainAccountAddressResponse{} }
func (m *QueryInterchainAccountAddressResponse) String() string { return proto.CompactTextString(m) }
func (*QueryInterchainAccountAddressResponse) ProtoMessage()    {}
func (*QueryInterchainAccountAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_85130b102faab7ea, []int{3}
}
func (m *QueryInterchainAccountAddressResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryInterchainAccountAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryInterchainAccountAddressResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryInterchainAccountAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryInterchainAccountAddressResponse.Merge(m, src)
}
func (m *QueryInterchainAccountAddressResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryInterchainAccountAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryInterchainAccountAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryInterchainAccountAddressResponse proto.InternalMessageInfo

func (m *QueryInterchainAccountAddressResponse) GetInterchainAccountAddress() string {
	if m != nil {
		return m.InterchainAccountAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "neutron.interchainadapter.interchaintxs.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "neutron.interchainadapter.interchaintxs.QueryParamsResponse")
	proto.RegisterType((*QueryInterchainAccountAddressRequest)(nil), "neutron.interchainadapter.interchaintxs.QueryInterchainAccountAddressRequest")
	proto.RegisterType((*QueryInterchainAccountAddressResponse)(nil), "neutron.interchainadapter.interchaintxs.QueryInterchainAccountAddressResponse")
}

func init() { proto.RegisterFile("interchaintxs/v1/query.proto", fileDescriptor_85130b102faab7ea) }

var fileDescriptor_85130b102faab7ea = []byte{
	// 464 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xbf, 0x6e, 0xd4, 0x30,
	0x18, 0x8f, 0x5b, 0x38, 0x81, 0x81, 0xc5, 0x2d, 0x52, 0x14, 0x95, 0x1c, 0x0a, 0x20, 0x10, 0xa8,
	0xb1, 0xee, 0xca, 0x04, 0x5d, 0xda, 0xed, 0x86, 0x56, 0x70, 0x23, 0xcb, 0xc9, 0x49, 0xac, 0xd4,
	0x12, 0xe7, 0x2f, 0xb5, 0x9d, 0xd2, 0x6e, 0x88, 0x89, 0x11, 0x89, 0x17, 0xe8, 0x13, 0x30, 0x23,
	0xf1, 0x02, 0x1d, 0x2b, 0xb1, 0x30, 0x21, 0x74, 0xc7, 0xc0, 0x63, 0xa0, 0xd8, 0x6e, 0x4f, 0x69,
	0xa9, 0x7a, 0x20, 0xb6, 0xe4, 0xcb, 0xef, 0xdf, 0xf7, 0xb3, 0x83, 0x57, 0x84, 0x34, 0x5c, 0xe5,
	0x3b, 0x4c, 0x48, 0xb3, 0xaf, 0xe9, 0x5e, 0x8f, 0xee, 0xd6, 0x5c, 0x1d, 0xa4, 0x95, 0x02, 0x03,
	0xe4, 0xa1, 0xe4, 0xb5, 0x51, 0x20, 0xd3, 0x19, 0x8a, 0x15, 0xac, 0x32, 0x5c, 0xa5, 0x2d, 0x5e,
	0xb4, 0x5c, 0x42, 0x09, 0x96, 0x43, 0x9b, 0x27, 0x47, 0x8f, 0x56, 0x4a, 0x80, 0xf2, 0x35, 0xa7,
	0xac, 0x12, 0x94, 0x49, 0x09, 0x86, 0x19, 0x01, 0x52, 0xfb, 0xaf, 0x8f, 0x73, 0xd0, 0x63, 0xd0,
	0x34, 0x63, 0x9a, 0x3b, 0x57, 0xba, 0xd7, 0xcb, 0xb8, 0x61, 0x3d, 0x5a, 0xb1, 0x52, 0x48, 0x0b,
	0xf6, 0xd8, 0x3b, 0xe7, 0x62, 0x56, 0x4c, 0xb1, 0xb1, 0x97, 0x4a, 0x96, 0x31, 0x79, 0xd9, 0x08,
	0xbc, 0xb0, 0xc3, 0x21, 0xdf, 0xad, 0xb9, 0x36, 0x49, 0x81, 0x97, 0x5a, 0x53, 0x5d, 0x81, 0xd4,
	0x9c, 0x6c, 0xe1, 0x8e, 0x23, 0x87, 0xe8, 0x2e, 0x7a, 0x74, 0xa3, 0x4f, 0xd3, 0x39, 0xb7, 0x4c,
	0x9d, 0xd0, 0xe6, 0x95, 0xa3, 0xef, 0xdd, 0x60, 0xe8, 0x45, 0x92, 0x4f, 0x08, 0xdf, 0xb7, 0x36,
	0x83, 0x53, 0xec, 0x46, 0x9e, 0x43, 0x2d, 0xcd, 0x46, 0x51, 0x28, 0xae, 0x4f, 0xe2, 0x90, 0x7b,
	0xf8, 0x16, 0xbc, 0x91, 0x5c, 0x8d, 0x98, 0x9b, 0x5b, 0xfb, 0xeb, 0xc3, 0x9b, 0x76, 0xe8, 0xb1,
	0xa4, 0x8f, 0x6f, 0xcf, 0x3c, 0x47, 0xcc, 0x09, 0x8d, 0x44, 0x11, 0x2e, 0x58, 0xf0, 0x92, 0x38,
	0x6b, 0x32, 0x28, 0x1a, 0xe1, 0x1c, 0xa4, 0xe4, 0x79, 0x53, 0x58, 0x83, 0x5d, 0x74, 0xc2, 0xb3,
	0xe1, 0xa0, 0x78, 0x76, 0xed, 0xfd, 0x61, 0x37, 0xf8, 0x75, 0xd8, 0x0d, 0x12, 0x8e, 0x1f, 0x5c,
	0x92, 0xd7, 0x17, 0xb5, 0x8e, 0xa3, 0x3f, 0x64, 0x69, 0xa7, 0x0f, 0xc5, 0x05, 0x2a, 0xfd, 0xb7,
	0x8b, 0xf8, 0xaa, 0xf5, 0x21, 0x9f, 0x11, 0xee, 0xb8, 0xea, 0xc8, 0xf3, 0xb9, 0xbb, 0x3e, 0x7f,
	0x9e, 0xd1, 0xfa, 0xbf, 0x91, 0xdd, 0x36, 0xc9, 0xda, 0xbb, 0xaf, 0x3f, 0x3f, 0x2e, 0xac, 0x92,
	0x27, 0xd4, 0xab, 0xd0, 0xf6, 0x9d, 0x6a, 0xbf, 0xb9, 0xc3, 0x25, 0x5f, 0x10, 0x0e, 0x2f, 0xea,
	0x89, 0x6c, 0xfd, 0x5d, 0x9e, 0x4b, 0xee, 0x47, 0xb4, 0xfd, 0xbf, 0xe4, 0xfc, 0xc2, 0xc1, 0xe6,
	0xf6, 0xd1, 0x24, 0x46, 0xc7, 0x93, 0x18, 0xfd, 0x98, 0xc4, 0xe8, 0xc3, 0x34, 0x0e, 0x8e, 0xa7,
	0x71, 0xf0, 0x6d, 0x1a, 0x07, 0xaf, 0x9e, 0x96, 0xc2, 0xec, 0xd4, 0x59, 0x9a, 0xc3, 0xf8, 0xa4,
	0x8e, 0x55, 0x50, 0xe5, 0x69, 0x35, 0xfb, 0x67, 0xea, 0x30, 0x07, 0x15, 0xd7, 0x59, 0xc7, 0xfe,
	0x6d, 0x6b, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xaf, 0xae, 0x1c, 0xe2, 0x35, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	InterchainAccountAddress(ctx context.Context, in *QueryInterchainAccountAddressRequest, opts ...grpc.CallOption) (*QueryInterchainAccountAddressResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/neutron.interchainadapter.interchaintxs.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) InterchainAccountAddress(ctx context.Context, in *QueryInterchainAccountAddressRequest, opts ...grpc.CallOption) (*QueryInterchainAccountAddressResponse, error) {
	out := new(QueryInterchainAccountAddressResponse)
	err := c.cc.Invoke(ctx, "/neutron.interchainadapter.interchaintxs.Query/InterchainAccountAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	InterchainAccountAddress(context.Context, *QueryInterchainAccountAddressRequest) (*QueryInterchainAccountAddressResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) InterchainAccountAddress(ctx context.Context, req *QueryInterchainAccountAddressRequest) (*QueryInterchainAccountAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InterchainAccountAddress not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neutron.interchainadapter.interchaintxs.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_InterchainAccountAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryInterchainAccountAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).InterchainAccountAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neutron.interchainadapter.interchaintxs.Query/InterchainAccountAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).InterchainAccountAddress(ctx, req.(*QueryInterchainAccountAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "neutron.interchainadapter.interchaintxs.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "InterchainAccountAddress",
			Handler:    _Query_InterchainAccountAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interchaintxs/v1/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryInterchainAccountAddressRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryInterchainAccountAddressRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryInterchainAccountAddressRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ConnectionId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.InterchainAccountId) > 0 {
		i -= len(m.InterchainAccountId)
		copy(dAtA[i:], m.InterchainAccountId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.InterchainAccountId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.OwnerAddress) > 0 {
		i -= len(m.OwnerAddress)
		copy(dAtA[i:], m.OwnerAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.OwnerAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryInterchainAccountAddressResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryInterchainAccountAddressResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryInterchainAccountAddressResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.InterchainAccountAddress) > 0 {
		i -= len(m.InterchainAccountAddress)
		copy(dAtA[i:], m.InterchainAccountAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.InterchainAccountAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryInterchainAccountAddressRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OwnerAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.InterchainAccountId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryInterchainAccountAddressResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.InterchainAccountAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryInterchainAccountAddressRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryInterchainAccountAddressRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryInterchainAccountAddressRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OwnerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InterchainAccountId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InterchainAccountId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryInterchainAccountAddressResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryInterchainAccountAddressResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryInterchainAccountAddressResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InterchainAccountAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InterchainAccountAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
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
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
