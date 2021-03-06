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

// QuerySyncValidatorsRequest is request type for Query/SyncValidators RPC method.
type QuerySyncValidatorsRequest struct {
	// pagination defines an optional pagination for the request.
	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QuerySyncValidatorsRequest) Reset()         { *m = QuerySyncValidatorsRequest{} }
func (m *QuerySyncValidatorsRequest) String() string { return proto.CompactTextString(m) }
func (*QuerySyncValidatorsRequest) ProtoMessage()    {}
func (*QuerySyncValidatorsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_83430a08489a14c4, []int{0}
}
func (m *QuerySyncValidatorsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySyncValidatorsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySyncValidatorsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySyncValidatorsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySyncValidatorsRequest.Merge(m, src)
}
func (m *QuerySyncValidatorsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QuerySyncValidatorsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySyncValidatorsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySyncValidatorsRequest proto.InternalMessageInfo

func (m *QuerySyncValidatorsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QuerySyncValidatorsResponse is response type for the Query/SyncValidators RPC method
type QuerySyncValidatorsResponse struct {
	// validators contains all the queried validators.
	Validators []types.Validator `protobuf:"bytes,1,rep,name=validators,proto3" json:"validators"`
	// pagination defines the pagination in the response.
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QuerySyncValidatorsResponse) Reset()         { *m = QuerySyncValidatorsResponse{} }
func (m *QuerySyncValidatorsResponse) String() string { return proto.CompactTextString(m) }
func (*QuerySyncValidatorsResponse) ProtoMessage()    {}
func (*QuerySyncValidatorsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_83430a08489a14c4, []int{1}
}
func (m *QuerySyncValidatorsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySyncValidatorsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySyncValidatorsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySyncValidatorsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySyncValidatorsResponse.Merge(m, src)
}
func (m *QuerySyncValidatorsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QuerySyncValidatorsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySyncValidatorsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySyncValidatorsResponse proto.InternalMessageInfo

func (m *QuerySyncValidatorsResponse) GetValidators() []types.Validator {
	if m != nil {
		return m.Validators
	}
	return nil
}

func (m *QuerySyncValidatorsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QuerySyncValidatorRequest is response type for the Query/SyncValidator RPC method
type QuerySyncValidatorRequest struct {
	// validator_addr defines the validator address to query for.
	ValidatorAddr string `protobuf:"bytes,1,opt,name=validator_addr,json=validatorAddr,proto3" json:"validator_addr,omitempty"`
}

func (m *QuerySyncValidatorRequest) Reset()         { *m = QuerySyncValidatorRequest{} }
func (m *QuerySyncValidatorRequest) String() string { return proto.CompactTextString(m) }
func (*QuerySyncValidatorRequest) ProtoMessage()    {}
func (*QuerySyncValidatorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_83430a08489a14c4, []int{2}
}
func (m *QuerySyncValidatorRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySyncValidatorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySyncValidatorRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySyncValidatorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySyncValidatorRequest.Merge(m, src)
}
func (m *QuerySyncValidatorRequest) XXX_Size() int {
	return m.Size()
}
func (m *QuerySyncValidatorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySyncValidatorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySyncValidatorRequest proto.InternalMessageInfo

func (m *QuerySyncValidatorRequest) GetValidatorAddr() string {
	if m != nil {
		return m.ValidatorAddr
	}
	return ""
}

// QuerySyncValidatorResponse is response type for the Query/SyncValidator RPC method
type QuerySyncValidatorResponse struct {
	// validator defines the the validator info.
	Validator types.Validator `protobuf:"bytes,1,opt,name=validator,proto3" json:"validator"`
}

func (m *QuerySyncValidatorResponse) Reset()         { *m = QuerySyncValidatorResponse{} }
func (m *QuerySyncValidatorResponse) String() string { return proto.CompactTextString(m) }
func (*QuerySyncValidatorResponse) ProtoMessage()    {}
func (*QuerySyncValidatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_83430a08489a14c4, []int{3}
}
func (m *QuerySyncValidatorResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySyncValidatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySyncValidatorResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySyncValidatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySyncValidatorResponse.Merge(m, src)
}
func (m *QuerySyncValidatorResponse) XXX_Size() int {
	return m.Size()
}
func (m *QuerySyncValidatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySyncValidatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySyncValidatorResponse proto.InternalMessageInfo

func (m *QuerySyncValidatorResponse) GetValidator() types.Validator {
	if m != nil {
		return m.Validator
	}
	return types.Validator{}
}

func init() {
	proto.RegisterType((*QuerySyncValidatorsRequest)(nil), "receiver.v1beta1.QuerySyncValidatorsRequest")
	proto.RegisterType((*QuerySyncValidatorsResponse)(nil), "receiver.v1beta1.QuerySyncValidatorsResponse")
	proto.RegisterType((*QuerySyncValidatorRequest)(nil), "receiver.v1beta1.QuerySyncValidatorRequest")
	proto.RegisterType((*QuerySyncValidatorResponse)(nil), "receiver.v1beta1.QuerySyncValidatorResponse")
}

func init() { proto.RegisterFile("receiver/v1beta/querier.proto", fileDescriptor_83430a08489a14c4) }

var fileDescriptor_83430a08489a14c4 = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x4f, 0x8b, 0x13, 0x31,
	0x18, 0xc6, 0x9b, 0xfa, 0x07, 0x36, 0xcb, 0x2e, 0x12, 0x3c, 0xd4, 0x59, 0x1d, 0x4b, 0x41, 0x2d,
	0xba, 0x4d, 0x68, 0xd5, 0xeb, 0xb2, 0xf6, 0xb0, 0xde, 0x44, 0x2b, 0x78, 0x10, 0x41, 0x32, 0x99,
	0x30, 0x1b, 0xdc, 0x49, 0x66, 0x93, 0xb4, 0x58, 0xc4, 0x8b, 0x9f, 0x40, 0xf0, 0x2b, 0x78, 0xf0,
	0xec, 0xa7, 0xd8, 0xe3, 0x82, 0x17, 0x4f, 0x22, 0xad, 0x1f, 0x44, 0x26, 0xf3, 0xaf, 0xb3, 0x5b,
	0x71, 0xf6, 0x16, 0x92, 0xf7, 0x79, 0x9e, 0x5f, 0xde, 0x37, 0x81, 0xb7, 0x34, 0x67, 0x5c, 0xcc,
	0xb8, 0x26, 0xb3, 0x61, 0xc0, 0x2d, 0x25, 0xc7, 0x53, 0xae, 0x05, 0xd7, 0x38, 0xd1, 0xca, 0x2a,
	0x74, 0xad, 0x38, 0xc6, 0xd9, 0xf1, 0xd0, 0xdb, 0xb1, 0x5c, 0x86, 0x5c, 0xc7, 0x42, 0x5a, 0x42,
	0x03, 0x26, 0x88, 0x9d, 0x27, 0xdc, 0x64, 0xe5, 0xde, 0x7d, 0xa6, 0x4c, 0xac, 0x0c, 0x09, 0xa8,
	0xe1, 0xce, 0x69, 0x9e, 0xdb, 0x0e, 0x49, 0x42, 0x23, 0x21, 0xa9, 0x15, 0x4a, 0xe6, 0xb5, 0xd7,
	0x23, 0x15, 0x29, 0xb7, 0x24, 0xe9, 0x2a, 0xdf, 0xbd, 0x19, 0x29, 0x15, 0x1d, 0x71, 0x42, 0x13,
	0x41, 0xa8, 0x94, 0xca, 0x3a, 0x49, 0xee, 0xdf, 0x0b, 0xa1, 0xf7, 0x22, 0x75, 0x7d, 0x39, 0x97,
	0xec, 0x15, 0x3d, 0x12, 0x21, 0xb5, 0x4a, 0x9b, 0x09, 0x3f, 0x9e, 0x72, 0x63, 0xd1, 0x01, 0x84,
	0x55, 0x4a, 0xa7, 0xdd, 0x05, 0xfd, 0xcd, 0xd1, 0x5d, 0x9c, 0x21, 0xe1, 0x14, 0x09, 0x3b, 0xa4,
	0xe2, 0x2a, 0xf8, 0x39, 0x8d, 0x78, 0xae, 0x9d, 0xac, 0x28, 0x7b, 0xdf, 0x00, 0xdc, 0x59, 0x1b,
	0x63, 0x12, 0x25, 0x0d, 0x47, 0xfb, 0x10, 0xce, 0xca, 0xdd, 0x0e, 0xe8, 0x5e, 0xea, 0x6f, 0x8e,
	0x3c, 0x5c, 0xf5, 0x05, 0xa7, 0x7d, 0xc1, 0xa5, 0x70, 0x7c, 0xf9, 0xe4, 0xd7, 0xed, 0xd6, 0x64,
	0x45, 0x83, 0x9e, 0xae, 0x21, 0xbd, 0xf7, 0x5f, 0xd2, 0x2c, 0xbe, 0x86, 0x3a, 0x86, 0x37, 0xce,
	0x93, 0x16, 0xfd, 0xb8, 0x03, 0xb7, 0xcb, 0xcc, 0xb7, 0x34, 0x0c, 0x75, 0x07, 0x74, 0x41, 0x7f,
	0x63, 0xb2, 0x55, 0xee, 0x3e, 0x09, 0x43, 0xdd, 0x7b, 0xb3, 0xae, 0xa9, 0xe5, 0x65, 0xf7, 0xe0,
	0x46, 0x59, 0xee, 0xf4, 0x4d, 0xee, 0x5a, 0x49, 0x46, 0xcb, 0x36, 0xbc, 0xe2, 0xec, 0xd1, 0x57,
	0x00, 0xb7, 0xeb, 0x1d, 0x45, 0xbb, 0xf8, 0xec, 0xfb, 0xc2, 0xff, 0x9e, 0xaf, 0x37, 0x68, 0x58,
	0x9d, 0x91, 0xf7, 0x1e, 0x7f, 0xfa, 0xf1, 0xe7, 0x4b, 0x9b, 0xa0, 0x01, 0x39, 0xa4, 0xec, 0x1d,
	0xb5, 0x2a, 0x1e, 0x88, 0x80, 0x0d, 0x68, 0x92, 0x90, 0x33, 0x8f, 0x7e, 0x48, 0x56, 0x66, 0xf3,
	0x1d, 0xc0, 0xad, 0x9a, 0x23, 0x7a, 0xd0, 0x24, 0xb7, 0x80, 0xdc, 0x6d, 0x56, 0x9c, 0x33, 0x1e,
	0x38, 0xc6, 0x7d, 0xb4, 0x77, 0x21, 0x46, 0xf2, 0xa1, 0x3e, 0xd7, 0x8f, 0xe3, 0x67, 0x27, 0x0b,
	0x1f, 0x9c, 0x2e, 0x7c, 0xf0, 0x7b, 0xe1, 0x83, 0xcf, 0x4b, 0xbf, 0x75, 0xba, 0xf4, 0x5b, 0x3f,
	0x97, 0x7e, 0xeb, 0xf5, 0xa3, 0x48, 0xd8, 0xc3, 0x69, 0x80, 0x99, 0x8a, 0x49, 0xcc, 0x94, 0x64,
	0xd4, 0x9e, 0xcf, 0x7a, 0x5f, 0xa5, 0xb9, 0xef, 0x1c, 0x5c, 0x75, 0xff, 0xed, 0xe1, 0xdf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x57, 0xfe, 0x7b, 0x3e, 0x1f, 0x04, 0x00, 0x00,
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
	// SyncValidators queries all validators that match the given status.
	SyncValidators(ctx context.Context, in *QuerySyncValidatorsRequest, opts ...grpc.CallOption) (*QuerySyncValidatorsResponse, error)
	// SyncValidator queries validator info for given validator address.
	SyncValidator(ctx context.Context, in *QuerySyncValidatorRequest, opts ...grpc.CallOption) (*QuerySyncValidatorResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) SyncValidators(ctx context.Context, in *QuerySyncValidatorsRequest, opts ...grpc.CallOption) (*QuerySyncValidatorsResponse, error) {
	out := new(QuerySyncValidatorsResponse)
	err := c.cc.Invoke(ctx, "/receiver.v1beta1.Query/SyncValidators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SyncValidator(ctx context.Context, in *QuerySyncValidatorRequest, opts ...grpc.CallOption) (*QuerySyncValidatorResponse, error) {
	out := new(QuerySyncValidatorResponse)
	err := c.cc.Invoke(ctx, "/receiver.v1beta1.Query/SyncValidator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// SyncValidators queries all validators that match the given status.
	SyncValidators(context.Context, *QuerySyncValidatorsRequest) (*QuerySyncValidatorsResponse, error)
	// SyncValidator queries validator info for given validator address.
	SyncValidator(context.Context, *QuerySyncValidatorRequest) (*QuerySyncValidatorResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) SyncValidators(ctx context.Context, req *QuerySyncValidatorsRequest) (*QuerySyncValidatorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncValidators not implemented")
}
func (*UnimplementedQueryServer) SyncValidator(ctx context.Context, req *QuerySyncValidatorRequest) (*QuerySyncValidatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncValidator not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_SyncValidators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySyncValidatorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SyncValidators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/receiver.v1beta1.Query/SyncValidators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SyncValidators(ctx, req.(*QuerySyncValidatorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SyncValidator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySyncValidatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SyncValidator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/receiver.v1beta1.Query/SyncValidator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SyncValidator(ctx, req.(*QuerySyncValidatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "receiver.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SyncValidators",
			Handler:    _Query_SyncValidators_Handler,
		},
		{
			MethodName: "SyncValidator",
			Handler:    _Query_SyncValidator_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "receiver/v1beta/querier.proto",
}

func (m *QuerySyncValidatorsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySyncValidatorsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySyncValidatorsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *QuerySyncValidatorsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySyncValidatorsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySyncValidatorsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *QuerySyncValidatorRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySyncValidatorRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySyncValidatorRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *QuerySyncValidatorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySyncValidatorResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySyncValidatorResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *QuerySyncValidatorsRequest) Size() (n int) {
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

func (m *QuerySyncValidatorsResponse) Size() (n int) {
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

func (m *QuerySyncValidatorRequest) Size() (n int) {
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

func (m *QuerySyncValidatorResponse) Size() (n int) {
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
func (m *QuerySyncValidatorsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QuerySyncValidatorsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySyncValidatorsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QuerySyncValidatorsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QuerySyncValidatorsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySyncValidatorsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QuerySyncValidatorRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QuerySyncValidatorRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySyncValidatorRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QuerySyncValidatorResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QuerySyncValidatorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySyncValidatorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
