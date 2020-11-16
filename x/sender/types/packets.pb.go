// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sender/v1beta/packets.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/tendermint/tendermint/abci/types"
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

type PacketDataValidatorSet struct {
	ValidatorsCommit []byte                   `protobuf:"bytes,1,opt,name=validators_commit,json=validatorsCommit,proto3" json:"validators_commit,omitempty"`
	ValidatorUpdates []*types.ValidatorUpdate `protobuf:"bytes,2,rep,name=validator_updates,json=validatorUpdates,proto3" json:"validator_updates,omitempty"`
}

func (m *PacketDataValidatorSet) Reset()         { *m = PacketDataValidatorSet{} }
func (m *PacketDataValidatorSet) String() string { return proto.CompactTextString(m) }
func (*PacketDataValidatorSet) ProtoMessage()    {}
func (*PacketDataValidatorSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_a33f01fc8e155314, []int{0}
}
func (m *PacketDataValidatorSet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PacketDataValidatorSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PacketDataValidatorSet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PacketDataValidatorSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PacketDataValidatorSet.Merge(m, src)
}
func (m *PacketDataValidatorSet) XXX_Size() int {
	return m.Size()
}
func (m *PacketDataValidatorSet) XXX_DiscardUnknown() {
	xxx_messageInfo_PacketDataValidatorSet.DiscardUnknown(m)
}

var xxx_messageInfo_PacketDataValidatorSet proto.InternalMessageInfo

func (m *PacketDataValidatorSet) GetValidatorsCommit() []byte {
	if m != nil {
		return m.ValidatorsCommit
	}
	return nil
}

func (m *PacketDataValidatorSet) GetValidatorUpdates() []*types.ValidatorUpdate {
	if m != nil {
		return m.ValidatorUpdates
	}
	return nil
}

func init() {
	proto.RegisterType((*PacketDataValidatorSet)(nil), "hackatomibcapp.sender.v1beta1.PacketDataValidatorSet")
}

func init() { proto.RegisterFile("sender/v1beta/packets.proto", fileDescriptor_a33f01fc8e155314) }

var fileDescriptor_a33f01fc8e155314 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x18, 0x84, 0x63, 0x90, 0x18, 0x02, 0x03, 0x64, 0x40, 0x15, 0x15, 0x56, 0xc4, 0x54, 0x09, 0xd5,
	0x56, 0xcb, 0x1b, 0x00, 0x23, 0x48, 0x08, 0x04, 0x03, 0x4b, 0xf5, 0xfb, 0x8f, 0x45, 0xad, 0xe2,
	0xd8, 0x8a, 0xff, 0x56, 0xf0, 0x16, 0x0c, 0x3c, 0x14, 0x63, 0x47, 0x46, 0x94, 0xbc, 0x08, 0xc2,
	0x69, 0x49, 0xc6, 0x3b, 0x7f, 0x3e, 0xdf, 0x39, 0x1d, 0x06, 0x5d, 0x16, 0xba, 0x92, 0xab, 0x89,
	0xd2, 0x04, 0xd2, 0x03, 0x2e, 0x34, 0x05, 0xe1, 0x2b, 0x47, 0x2e, 0x3b, 0x9d, 0x03, 0x2e, 0x80,
	0x9c, 0x35, 0x0a, 0xc1, 0x7b, 0xd1, 0xb2, 0xa2, 0x65, 0x27, 0x27, 0x43, 0x8a, 0xda, 0x9a, 0x92,
	0x24, 0x28, 0x34, 0x92, 0xde, 0xbd, 0xde, 0xdc, 0x3d, 0xfb, 0x64, 0xe9, 0xf1, 0x5d, 0x4c, 0xbb,
	0x06, 0x82, 0x27, 0x78, 0x35, 0x05, 0x90, 0xab, 0x1e, 0x34, 0x65, 0xe7, 0xe9, 0xd1, 0x6a, 0xab,
	0xc3, 0x0c, 0x9d, 0xb5, 0x86, 0x06, 0x2c, 0x67, 0xa3, 0x83, 0xfb, 0xc3, 0xee, 0xe0, 0x2a, 0xfa,
	0xd9, 0x6d, 0x0f, 0x9e, 0x2d, 0x7d, 0x01, 0xa4, 0xc3, 0x60, 0x27, 0xdf, 0x1d, 0xed, 0x4f, 0x73,
	0xd1, 0x15, 0x10, 0x7f, 0x05, 0xc4, 0xff, 0x33, 0x8f, 0x11, 0xec, 0xc5, 0xb5, 0x46, 0xb8, 0xbc,
	0xf9, 0xaa, 0x39, 0x5b, 0xd7, 0x9c, 0xfd, 0xd4, 0x9c, 0x7d, 0x34, 0x3c, 0x59, 0x37, 0x3c, 0xf9,
	0x6e, 0x78, 0xf2, 0x3c, 0x7d, 0x31, 0x34, 0x5f, 0x2a, 0x81, 0xce, 0x4a, 0x8b, 0xae, 0x44, 0x20,
	0xb9, 0xdd, 0x3f, 0x36, 0x0a, 0xc7, 0xe0, 0xbd, 0x7c, 0x93, 0x9b, 0xff, 0x8a, 0x53, 0xd5, 0x5e,
	0xdc, 0x7a, 0xf1, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x20, 0xeb, 0xb1, 0x46, 0x01, 0x00, 0x00,
}

func (m *PacketDataValidatorSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PacketDataValidatorSet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PacketDataValidatorSet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorUpdates) > 0 {
		for iNdEx := len(m.ValidatorUpdates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ValidatorUpdates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPackets(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.ValidatorsCommit) > 0 {
		i -= len(m.ValidatorsCommit)
		copy(dAtA[i:], m.ValidatorsCommit)
		i = encodeVarintPackets(dAtA, i, uint64(len(m.ValidatorsCommit)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPackets(dAtA []byte, offset int, v uint64) int {
	offset -= sovPackets(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PacketDataValidatorSet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValidatorsCommit)
	if l > 0 {
		n += 1 + l + sovPackets(uint64(l))
	}
	if len(m.ValidatorUpdates) > 0 {
		for _, e := range m.ValidatorUpdates {
			l = e.Size()
			n += 1 + l + sovPackets(uint64(l))
		}
	}
	return n
}

func sovPackets(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPackets(x uint64) (n int) {
	return sovPackets(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PacketDataValidatorSet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPackets
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
			return fmt.Errorf("proto: PacketDataValidatorSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PacketDataValidatorSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorsCommit", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPackets
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPackets
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPackets
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorsCommit = append(m.ValidatorsCommit[:0], dAtA[iNdEx:postIndex]...)
			if m.ValidatorsCommit == nil {
				m.ValidatorsCommit = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorUpdates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPackets
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
				return ErrInvalidLengthPackets
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPackets
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorUpdates = append(m.ValidatorUpdates, &types.ValidatorUpdate{})
			if err := m.ValidatorUpdates[len(m.ValidatorUpdates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPackets(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPackets
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPackets
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
func skipPackets(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPackets
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
					return 0, ErrIntOverflowPackets
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
					return 0, ErrIntOverflowPackets
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
				return 0, ErrInvalidLengthPackets
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPackets
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPackets
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPackets        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPackets          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPackets = fmt.Errorf("proto: unexpected end of group")
)
