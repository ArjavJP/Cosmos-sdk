// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/core/types/v1/genesis.proto

package types

import (
	fmt "fmt"
	types "https://github.com/ArjavJP/Cosmos-sdk/x/ibc/core/02-client/types"
	types1 "https://github.com/ArjavJP/Cosmos-sdk/x/ibc/core/03-connection/types"
	types2 "https://github.com/ArjavJP/Cosmos-sdk/x/ibc/core/04-channel/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// GenesisState defines the ibc module's genesis state.
type GenesisState struct {
	// ICS002 - Clients genesis state
	ClientGenesis types.GenesisState `protobuf:"bytes,1,opt,name=client_genesis,json=clientGenesis,proto3" json:"client_genesis" yaml:"client_genesis"`
	// ICS003 - Connections genesis state
	ConnectionGenesis types1.GenesisState `protobuf:"bytes,2,opt,name=connection_genesis,json=connectionGenesis,proto3" json:"connection_genesis" yaml:"connection_genesis"`
	// ICS004 - Channel genesis state
	ChannelGenesis types2.GenesisState `protobuf:"bytes,3,opt,name=channel_genesis,json=channelGenesis,proto3" json:"channel_genesis" yaml:"channel_genesis"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a49c5663e6fc59, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetClientGenesis() types.GenesisState {
	if m != nil {
		return m.ClientGenesis
	}
	return types.GenesisState{}
}

func (m *GenesisState) GetConnectionGenesis() types1.GenesisState {
	if m != nil {
		return m.ConnectionGenesis
	}
	return types1.GenesisState{}
}

func (m *GenesisState) GetChannelGenesis() types2.GenesisState {
	if m != nil {
		return m.ChannelGenesis
	}
	return types2.GenesisState{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "ibc.core.types.v1.GenesisState")
}

func init() { proto.RegisterFile("ibc/core/types/v1/genesis.proto", fileDescriptor_b9a49c5663e6fc59) }

var fileDescriptor_b9a49c5663e6fc59 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x3d, 0x4e, 0xc3, 0x30,
	0x18, 0x86, 0x93, 0x22, 0x31, 0x04, 0x28, 0x6a, 0x04, 0x08, 0x2a, 0xe1, 0x36, 0x51, 0x07, 0x96,
	0xda, 0x0a, 0x6c, 0x8c, 0x5d, 0xba, 0x87, 0x8d, 0x05, 0x25, 0xc6, 0xa4, 0x86, 0xc4, 0xae, 0x6a,
	0x13, 0xd1, 0x5b, 0x70, 0xac, 0x8e, 0x1d, 0x11, 0x43, 0x85, 0x92, 0x1b, 0x70, 0x02, 0xd4, 0xd8,
	0xe4, 0x47, 0x9e, 0x12, 0xbd, 0x7e, 0xbe, 0xf7, 0xf9, 0x94, 0xd8, 0x19, 0xd1, 0x18, 0x23, 0xcc,
	0x57, 0x04, 0xc9, 0xf5, 0x92, 0x08, 0x94, 0x07, 0x28, 0x21, 0x8c, 0x08, 0x2a, 0xe0, 0x72, 0xc5,
	0x25, 0x77, 0x07, 0x34, 0xc6, 0x70, 0x0f, 0xc0, 0x0a, 0x80, 0x79, 0x30, 0x3c, 0x4b, 0x78, 0xc2,
	0xab, 0x53, 0xb4, 0x7f, 0x53, 0xe0, 0x70, 0x5c, 0x37, 0xe1, 0x94, 0x12, 0x26, 0x8d, 0xaa, 0xe1,
	0xa4, 0x21, 0x38, 0x63, 0x04, 0x4b, 0xca, 0x99, 0x49, 0x79, 0x0d, 0xb5, 0x88, 0x18, 0x23, 0xa9,
	0x81, 0xf8, 0xdf, 0x3d, 0xe7, 0x78, 0xae, 0x92, 0x07, 0x19, 0x49, 0xe2, 0xbe, 0x38, 0x7d, 0x25,
	0x7d, 0xd2, 0xe0, 0xa5, 0x3d, 0xb6, 0x6f, 0x8e, 0x6e, 0xc7, 0xb0, 0xde, 0x5e, 0x9d, 0xc3, 0x3c,
	0x80, 0xed, 0xc9, 0xd9, 0xf5, 0x66, 0x37, 0xb2, 0x7e, 0x77, 0xa3, 0xf3, 0x75, 0x94, 0xa5, 0xf7,
	0x7e, 0xb7, 0xc5, 0x0f, 0x4f, 0x54, 0xa0, 0x47, 0xdc, 0xdc, 0x71, 0x9b, 0xd5, 0x6b, 0x57, 0xaf,
	0x72, 0x4d, 0x5a, 0xae, 0x9a, 0x31, 0x7c, 0x9e, 0xf6, 0x5d, 0x69, 0x9f, 0xd1, 0xe6, 0x87, 0x83,
	0x26, 0xfc, 0xf7, 0xbe, 0x3a, 0xa7, 0xfa, 0x63, 0xd4, 0xd2, 0x83, 0x4a, 0xea, 0xb5, 0xa4, 0x0a,
	0x30, 0x8c, 0x40, 0x1b, 0x2f, 0xb4, 0xb1, 0xdb, 0xe3, 0x87, 0x7d, 0x9d, 0xe8, 0xa1, 0xd9, 0x7c,
	0x53, 0x00, 0x7b, 0x5b, 0x00, 0xfb, 0xa7, 0x00, 0xf6, 0x67, 0x09, 0xac, 0x6d, 0x09, 0xac, 0xaf,
	0x12, 0x58, 0x8f, 0xd3, 0x84, 0xca, 0xc5, 0x7b, 0x0c, 0x31, 0xcf, 0x10, 0xe6, 0x22, 0xe3, 0x42,
	0x3f, 0xa6, 0xe2, 0xf9, 0x0d, 0x7d, 0xa0, 0xee, 0x55, 0x8a, 0x0f, 0xab, 0x9f, 0x75, 0xf7, 0x17,
	0x00, 0x00, 0xff, 0xff, 0x90, 0x81, 0x26, 0xd1, 0x63, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.ChannelGenesis.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.ConnectionGenesis.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.ClientGenesis.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ClientGenesis.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.ConnectionGenesis.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.ChannelGenesis.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientGenesis", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ClientGenesis.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionGenesis", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ConnectionGenesis.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelGenesis", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ChannelGenesis.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
