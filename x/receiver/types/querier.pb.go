// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: receiver/v1beta/querier.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/tendermint/tendermint/abci/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

// QueryValidatorsRequest is request type for Query/Validators RPC method.
type QueryValidatorsRequest struct {
	// pagination defines an optional pagination for the request.
	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryValidatorsRequest) Reset()         { *m = QueryValidatorsRequest{} }
func (m *QueryValidatorsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryValidatorsRequest) ProtoMessage()    {}
func (*QueryValidatorsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_83430a08489a14c4, []int{0}
}
func (m *QueryValidatorsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryValidatorsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryValidatorsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryValidatorsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryValidatorsRequest.Merge(m, src)
}
func (m *QueryValidatorsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryValidatorsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryValidatorsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryValidatorsRequest proto.InternalMessageInfo

func (m *QueryValidatorsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryValidatorsResponse is response type for the Query/Validators RPC method
type QueryValidatorsResponse struct {
	// validators contains all the queried validators.
	Validators []types.Validator `protobuf:"bytes,1,rep,name=validators,proto3" json:"validators"`
	// pagination defines the pagination in the response.
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryValidatorsResponse) Reset()         { *m = QueryValidatorsResponse{} }
func (m *QueryValidatorsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryValidatorsResponse) ProtoMessage()    {}
func (*QueryValidatorsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_83430a08489a14c4, []int{1}
}
func (m *QueryValidatorsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryValidatorsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryValidatorsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryValidatorsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryValidatorsResponse.Merge(m, src)
}
func (m *QueryValidatorsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryValidatorsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryValidatorsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryValidatorsResponse proto.InternalMessageInfo

func (m *QueryValidatorsResponse) GetValidators() []types.Validator {
	if m != nil {
		return m.Validators
	}
	return nil
}

func (m *QueryValidatorsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryValidatorRequest is response type for the Query/Validator RPC method
type QueryValidatorRequest struct {
	// validator_addr defines the validator address to query for.
	ValidatorAddr string `protobuf:"bytes,1,opt,name=validator_addr,json=validatorAddr,proto3" json:"validator_addr,omitempty"`
}

func (m *QueryValidatorRequest) Reset()         { *m = QueryValidatorRequest{} }
func (m *QueryValidatorRequest) String() string { return proto.CompactTextString(m) }
func (*QueryValidatorRequest) ProtoMessage()    {}
func (*QueryValidatorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_83430a08489a14c4, []int{2}
}
func (m *QueryValidatorRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryValidatorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryValidatorRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryValidatorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryValidatorRequest.Merge(m, src)
}
func (m *QueryValidatorRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryValidatorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryValidatorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryValidatorRequest proto.InternalMessageInfo

func (m *QueryValidatorRequest) GetValidatorAddr() string {
	if m != nil {
		return m.ValidatorAddr
	}
	return ""
}

// QueryValidatorResponse is response type for the Query/Validator RPC method
type QueryValidatorResponse struct {
	// validator defines the the validator info.
	Validator types.Validator `protobuf:"bytes,1,opt,name=validator,proto3" json:"validator"`
}

func (m *QueryValidatorResponse) Reset()         { *m = QueryValidatorResponse{} }
func (m *QueryValidatorResponse) String() string { return proto.CompactTextString(m) }
func (*QueryValidatorResponse) ProtoMessage()    {}
func (*QueryValidatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_83430a08489a14c4, []int{3}
}
func (m *QueryValidatorResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryValidatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryValidatorResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryValidatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryValidatorResponse.Merge(m, src)
}
func (m *QueryValidatorResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryValidatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryValidatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryValidatorResponse proto.InternalMessageInfo

func (m *QueryValidatorResponse) GetValidator() types.Validator {
	if m != nil {
		return m.Validator
	}
	return types.Validator{}
}

func init() {
	proto.RegisterType((*QueryValidatorsRequest)(nil), "cosmos.staking.v1beta1.QueryValidatorsRequest")
	proto.RegisterType((*QueryValidatorsResponse)(nil), "cosmos.staking.v1beta1.QueryValidatorsResponse")
	proto.RegisterType((*QueryValidatorRequest)(nil), "cosmos.staking.v1beta1.QueryValidatorRequest")
	proto.RegisterType((*QueryValidatorResponse)(nil), "cosmos.staking.v1beta1.QueryValidatorResponse")
}

func init() { proto.RegisterFile("receiver/v1beta/querier.proto", fileDescriptor_83430a08489a14c4) }

var fileDescriptor_83430a08489a14c4 = []byte{
	// 475 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcf, 0x8a, 0x13, 0x41,
	0x10, 0xc6, 0xd3, 0xf1, 0x0f, 0xa4, 0x17, 0x3d, 0x34, 0xba, 0x86, 0xa8, 0x63, 0x08, 0xa8, 0x41,
	0x48, 0x17, 0x89, 0x7a, 0x0d, 0xeb, 0x1e, 0xd6, 0x9b, 0x68, 0x0e, 0x22, 0x5e, 0xb4, 0x66, 0xa6,
	0x99, 0x6d, 0x76, 0xa7, 0x7b, 0xb6, 0xbb, 0x13, 0x5c, 0xc4, 0x8b, 0x4f, 0x20, 0xf8, 0x0a, 0xfb,
	0x04, 0x3e, 0xc5, 0x1e, 0x17, 0xbc, 0x78, 0x12, 0x49, 0x7c, 0x10, 0x49, 0xcf, 0x64, 0x26, 0xd9,
	0x0d, 0x98, 0xbd, 0x35, 0x4d, 0x7d, 0xdf, 0xf7, 0xab, 0xaa, 0x6e, 0x7a, 0xdf, 0x88, 0x48, 0xc8,
	0x89, 0x30, 0x30, 0xe9, 0x87, 0xc2, 0x21, 0x1c, 0x8d, 0x85, 0x91, 0xc2, 0xf0, 0xcc, 0x68, 0xa7,
	0xd9, 0x76, 0xa4, 0x6d, 0xaa, 0x2d, 0xb7, 0x0e, 0x0f, 0xa4, 0x4a, 0x78, 0x5e, 0xd4, 0x6f, 0xdd,
	0x75, 0x42, 0xc5, 0xc2, 0xa4, 0x52, 0x39, 0xc0, 0x30, 0x92, 0xe0, 0x8e, 0x33, 0x61, 0x73, 0x51,
	0xeb, 0x49, 0x2e, 0x82, 0x10, 0xad, 0xf0, 0x7e, 0xc7, 0x85, 0x79, 0x1f, 0x32, 0x4c, 0xa4, 0x42,
	0x27, 0xb5, 0x2a, 0x6a, 0x6f, 0x25, 0x3a, 0xd1, 0xfe, 0x08, 0xf3, 0x53, 0x71, 0x7b, 0x2f, 0xd1,
	0x3a, 0x39, 0x14, 0x80, 0x99, 0x04, 0x54, 0x4a, 0x3b, 0x2f, 0x29, 0xfc, 0x3b, 0x1f, 0xe9, 0xf6,
	0x9b, 0xb9, 0xeb, 0x5b, 0x3c, 0x94, 0x31, 0x3a, 0x6d, 0xec, 0x48, 0x1c, 0x8d, 0x85, 0x75, 0x6c,
	0x8f, 0xd2, 0x2a, 0xa1, 0x59, 0x6f, 0x93, 0xee, 0xd6, 0xe0, 0x11, 0x2f, 0x7a, 0x98, 0xe3, 0x70,
	0x8f, 0xb3, 0x68, 0x83, 0xbf, 0xc6, 0x44, 0x14, 0xda, 0xd1, 0x92, 0xb2, 0x73, 0x42, 0xe8, 0x9d,
	0x0b, 0x11, 0x36, 0xd3, 0xca, 0x0a, 0xb6, 0x43, 0xe9, 0xa4, 0xbc, 0x6d, 0x92, 0xf6, 0x95, 0xee,
	0xd6, 0xa0, 0xc5, 0xab, 0x79, 0xf0, 0xf9, 0x3c, 0x78, 0x29, 0xdc, 0xbd, 0x7a, 0xfa, 0xfb, 0x41,
	0x6d, 0xb4, 0xa4, 0x61, 0x2f, 0xd7, 0x50, 0x3e, 0xfe, 0x2f, 0x65, 0x1e, 0xbf, 0x82, 0x39, 0xa4,
	0xb7, 0x57, 0x29, 0x17, 0x73, 0x78, 0x48, 0x6f, 0x96, 0x79, 0x1f, 0x30, 0x8e, 0x4d, 0x93, 0xb4,
	0x49, 0xb7, 0x31, 0xba, 0x51, 0xde, 0xbe, 0x88, 0x63, 0xd3, 0x79, 0x77, 0x7e, 0x90, 0x65, 0x93,
	0x43, 0xda, 0x28, 0x4b, 0xbd, 0x76, 0x93, 0x1e, 0x2b, 0xc9, 0x60, 0x56, 0xa7, 0xd7, 0xbc, 0x35,
	0x3b, 0x21, 0x94, 0x56, 0x53, 0x64, 0x9c, 0xaf, 0x7f, 0x51, 0x7c, 0xfd, 0x46, 0x5b, 0xb0, 0x71,
	0x7d, 0x4e, 0xde, 0x79, 0xfe, 0xf5, 0xe7, 0xdf, 0xef, 0x75, 0x60, 0x3d, 0xd8, 0xc7, 0xe8, 0x00,
	0x9d, 0x4e, 0x7b, 0x32, 0x8c, 0x7a, 0x98, 0x65, 0x70, 0xee, 0xa9, 0xf7, 0x61, 0x69, 0x27, 0x3f,
	0x08, 0x6d, 0x94, 0x6e, 0xac, 0xb7, 0x59, 0xea, 0x02, 0x92, 0x6f, 0x5a, 0x5e, 0x30, 0xee, 0x79,
	0xc6, 0x1d, 0x36, 0xbc, 0x14, 0x23, 0x7c, 0x5e, 0xdd, 0xe9, 0x97, 0xdd, 0x57, 0xa7, 0xd3, 0x80,
	0x9c, 0x4d, 0x03, 0xf2, 0x67, 0x1a, 0x90, 0x6f, 0xb3, 0xa0, 0x76, 0x36, 0x0b, 0x6a, 0xbf, 0x66,
	0x41, 0xed, 0xfd, 0xb3, 0x44, 0xba, 0xfd, 0x71, 0xc8, 0x23, 0x9d, 0x42, 0x1a, 0x69, 0x15, 0xa1,
	0xbb, 0x98, 0xf5, 0xa9, 0x4a, 0xf3, 0xdf, 0x37, 0xbc, 0xee, 0xff, 0xd7, 0xd3, 0x7f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x0f, 0x24, 0xf1, 0xe3, 0x15, 0x04, 0x00, 0x00,
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
	// Validators queries all validators that match the given status.
	Validators(ctx context.Context, in *QueryValidatorsRequest, opts ...grpc.CallOption) (*QueryValidatorsResponse, error)
	// Validator queries validator info for given validator address.
	Validator(ctx context.Context, in *QueryValidatorRequest, opts ...grpc.CallOption) (*QueryValidatorResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Validators(ctx context.Context, in *QueryValidatorsRequest, opts ...grpc.CallOption) (*QueryValidatorsResponse, error) {
	out := new(QueryValidatorsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.staking.v1beta1.Query/Validators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Validator(ctx context.Context, in *QueryValidatorRequest, opts ...grpc.CallOption) (*QueryValidatorResponse, error) {
	out := new(QueryValidatorResponse)
	err := c.cc.Invoke(ctx, "/cosmos.staking.v1beta1.Query/Validator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Validators queries all validators that match the given status.
	Validators(context.Context, *QueryValidatorsRequest) (*QueryValidatorsResponse, error)
	// Validator queries validator info for given validator address.
	Validator(context.Context, *QueryValidatorRequest) (*QueryValidatorResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Validators(ctx context.Context, req *QueryValidatorsRequest) (*QueryValidatorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validators not implemented")
}
func (*UnimplementedQueryServer) Validator(ctx context.Context, req *QueryValidatorRequest) (*QueryValidatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validator not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Validators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryValidatorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Validators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.Query/Validators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Validators(ctx, req.(*QueryValidatorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Validator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryValidatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Validator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.Query/Validator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Validator(ctx, req.(*QueryValidatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.staking.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Validators",
			Handler:    _Query_Validators_Handler,
		},
		{
			MethodName: "Validator",
			Handler:    _Query_Validator_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "receiver/v1beta/querier.proto",
}

func (m *QueryValidatorsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryValidatorsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryValidatorsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuerier(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func (m *QueryValidatorsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryValidatorsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryValidatorsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuerier(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Validators) > 0 {
		for iNdEx := len(m.Validators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Validators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuerier(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryValidatorRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryValidatorRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryValidatorRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorAddr) > 0 {
		i -= len(m.ValidatorAddr)
		copy(dAtA[i:], m.ValidatorAddr)
		i = encodeVarintQuerier(dAtA, i, uint64(len(m.ValidatorAddr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryValidatorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryValidatorResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryValidatorResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Validator.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuerier(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuerier(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuerier(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryValidatorsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuerier(uint64(l))
	}
	return n
}

func (m *QueryValidatorsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Validators) > 0 {
		for _, e := range m.Validators {
			l = e.Size()
			n += 1 + l + sovQuerier(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuerier(uint64(l))
	}
	return n
}

func (m *QueryValidatorRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValidatorAddr)
	if l > 0 {
		n += 1 + l + sovQuerier(uint64(l))
	}
	return n
}

func (m *QueryValidatorResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Validator.Size()
	n += 1 + l + sovQuerier(uint64(l))
	return n
}

func sovQuerier(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuerier(x uint64) (n int) {
	return sovQuerier(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryValidatorsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuerier
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
			return fmt.Errorf("proto: QueryValidatorsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryValidatorsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerier
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
				return ErrInvalidLengthQuerier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuerier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuerier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuerier
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuerier
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
func (m *QueryValidatorsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuerier
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
			return fmt.Errorf("proto: QueryValidatorsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryValidatorsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerier
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
				return ErrInvalidLengthQuerier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuerier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validators = append(m.Validators, types.Validator{})
			if err := m.Validators[len(m.Validators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerier
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
				return ErrInvalidLengthQuerier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuerier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuerier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuerier
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuerier
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
func (m *QueryValidatorRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuerier
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
			return fmt.Errorf("proto: QueryValidatorRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryValidatorRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerier
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
				return ErrInvalidLengthQuerier
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuerier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuerier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuerier
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuerier
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
func (m *QueryValidatorResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuerier
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
			return fmt.Errorf("proto: QueryValidatorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryValidatorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerier
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
				return ErrInvalidLengthQuerier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuerier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Validator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuerier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuerier
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuerier
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
func skipQuerier(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuerier
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
					return 0, ErrIntOverflowQuerier
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
					return 0, ErrIntOverflowQuerier
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
				return 0, ErrInvalidLengthQuerier
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuerier
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuerier
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuerier        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuerier          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuerier = fmt.Errorf("proto: unexpected end of group")
)
