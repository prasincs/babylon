// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: babylon/incentive/incentive.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Gauge is an object that stores rewards to be distributed
// code adapted from https://github.com/osmosis-labs/osmosis/blob/v18.0.0/proto/osmosis/incentives/gauge.proto
type Gauge struct {
	// coins are coins that have been in the gauge
	// Can have multiple coin denoms
	Coins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
}

func (m *Gauge) Reset()         { *m = Gauge{} }
func (m *Gauge) String() string { return proto.CompactTextString(m) }
func (*Gauge) ProtoMessage()    {}
func (*Gauge) Descriptor() ([]byte, []int) {
	return fileDescriptor_3954bc4942045a7a, []int{0}
}
func (m *Gauge) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Gauge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Gauge.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Gauge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Gauge.Merge(m, src)
}
func (m *Gauge) XXX_Size() int {
	return m.Size()
}
func (m *Gauge) XXX_DiscardUnknown() {
	xxx_messageInfo_Gauge.DiscardUnknown(m)
}

var xxx_messageInfo_Gauge proto.InternalMessageInfo

func (m *Gauge) GetCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Coins
	}
	return nil
}

// RewardGauge is an object that stores rewards distributed to a BTC staking/timestamping stakeholder
// code adapted from https://github.com/osmosis-labs/osmosis/blob/v18.0.0/proto/osmosis/incentives/gauge.proto
type RewardGauge struct {
	// coins are coins that have been in the gauge
	// Can have multiple coin denoms
	Coins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
	// withdrawn_coins are coins that have been withdrawn by the stakeholder already
	WithdrawnCoins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=withdrawn_coins,json=withdrawnCoins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"withdrawn_coins"`
}

func (m *RewardGauge) Reset()         { *m = RewardGauge{} }
func (m *RewardGauge) String() string { return proto.CompactTextString(m) }
func (*RewardGauge) ProtoMessage()    {}
func (*RewardGauge) Descriptor() ([]byte, []int) {
	return fileDescriptor_3954bc4942045a7a, []int{1}
}
func (m *RewardGauge) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RewardGauge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RewardGauge.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RewardGauge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RewardGauge.Merge(m, src)
}
func (m *RewardGauge) XXX_Size() int {
	return m.Size()
}
func (m *RewardGauge) XXX_DiscardUnknown() {
	xxx_messageInfo_RewardGauge.DiscardUnknown(m)
}

var xxx_messageInfo_RewardGauge proto.InternalMessageInfo

func (m *RewardGauge) GetCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Coins
	}
	return nil
}

func (m *RewardGauge) GetWithdrawnCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.WithdrawnCoins
	}
	return nil
}

func init() {
	proto.RegisterType((*Gauge)(nil), "babylon.incentive.Gauge")
	proto.RegisterType((*RewardGauge)(nil), "babylon.incentive.RewardGauge")
}

func init() { proto.RegisterFile("babylon/incentive/incentive.proto", fileDescriptor_3954bc4942045a7a) }

var fileDescriptor_3954bc4942045a7a = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0x4a, 0x4c, 0xaa,
	0xcc, 0xc9, 0xcf, 0xd3, 0xcf, 0xcc, 0x4b, 0x4e, 0xcd, 0x2b, 0xc9, 0x2c, 0x4b, 0x45, 0xb0, 0xf4,
	0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x04, 0xa1, 0x4a, 0xf4, 0xe0, 0x12, 0x52, 0x22, 0xe9, 0xf9,
	0xe9, 0xf9, 0x60, 0x59, 0x7d, 0x10, 0x0b, 0xa2, 0x50, 0x4a, 0x2e, 0x39, 0xbf, 0x38, 0x37, 0xbf,
	0x58, 0x3f, 0x29, 0xb1, 0x38, 0x55, 0xbf, 0xcc, 0x30, 0x29, 0xb5, 0x24, 0xd1, 0x50, 0x3f, 0x39,
	0x3f, 0x33, 0x0f, 0x22, 0xaf, 0x94, 0xc5, 0xc5, 0xea, 0x9e, 0x58, 0x9a, 0x9e, 0x2a, 0x94, 0xc8,
	0xc5, 0x0a, 0x12, 0x2e, 0x96, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36, 0x92, 0xd4, 0x83, 0x68, 0xd4,
	0x03, 0x69, 0xd4, 0x83, 0x6a, 0xd4, 0x73, 0xce, 0xcf, 0xcc, 0x73, 0x32, 0x38, 0x71, 0x4f, 0x9e,
	0x61, 0xd5, 0x7d, 0x79, 0x8d, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d,
	0xa8, 0x2d, 0x10, 0x4a, 0xb7, 0x38, 0x25, 0x5b, 0xbf, 0xa4, 0xb2, 0x20, 0xb5, 0x18, 0xac, 0xa1,
	0x38, 0x08, 0x62, 0xb2, 0xd2, 0x33, 0x46, 0x2e, 0xee, 0xa0, 0xd4, 0xf2, 0xc4, 0xa2, 0x14, 0x7a,
	0x59, 0x29, 0x54, 0xc2, 0xc5, 0x5f, 0x9e, 0x59, 0x92, 0x91, 0x52, 0x94, 0x58, 0x9e, 0x17, 0x0f,
	0xb1, 0x8c, 0x89, 0xfa, 0x96, 0xf1, 0xc1, 0xed, 0x00, 0xf3, 0x9d, 0x7c, 0x4f, 0x3c, 0x92, 0x63,
	0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96,
	0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x18, 0xc9, 0x4c, 0x68, 0x14, 0xe6, 0x24, 0x26, 0x15,
	0xeb, 0x66, 0xe6, 0xc3, 0xb8, 0xfa, 0x15, 0x48, 0xd1, 0x0e, 0xb6, 0x24, 0x89, 0x0d, 0x1c, 0x55,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd1, 0xc1, 0x22, 0xfc, 0x18, 0x02, 0x00, 0x00,
}

func (m *Gauge) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Gauge) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Gauge) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintIncentive(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *RewardGauge) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RewardGauge) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RewardGauge) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.WithdrawnCoins) > 0 {
		for iNdEx := len(m.WithdrawnCoins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.WithdrawnCoins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintIncentive(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintIncentive(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintIncentive(dAtA []byte, offset int, v uint64) int {
	offset -= sovIncentive(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Gauge) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovIncentive(uint64(l))
		}
	}
	return n
}

func (m *RewardGauge) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovIncentive(uint64(l))
		}
	}
	if len(m.WithdrawnCoins) > 0 {
		for _, e := range m.WithdrawnCoins {
			l = e.Size()
			n += 1 + l + sovIncentive(uint64(l))
		}
	}
	return n
}

func sovIncentive(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIncentive(x uint64) (n int) {
	return sovIncentive(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Gauge) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIncentive
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
			return fmt.Errorf("proto: Gauge: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Gauge: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, types.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIncentive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIncentive
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
func (m *RewardGauge) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIncentive
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
			return fmt.Errorf("proto: RewardGauge: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RewardGauge: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, types.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawnCoins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WithdrawnCoins = append(m.WithdrawnCoins, types.Coin{})
			if err := m.WithdrawnCoins[len(m.WithdrawnCoins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIncentive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIncentive
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
func skipIncentive(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIncentive
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
					return 0, ErrIntOverflowIncentive
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
					return 0, ErrIntOverflowIncentive
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
				return 0, ErrInvalidLengthIncentive
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIncentive
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIncentive
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIncentive        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIncentive          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIncentive = fmt.Errorf("proto: unexpected end of group")
)
