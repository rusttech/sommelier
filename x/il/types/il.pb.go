// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: il/v1/il.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// Stoploss defines a set of parameters that together trigger a stoploss withdrawal.
type Stoploss struct {
	// uniswap pair hex address
	UniswapPairID string `protobuf:"bytes,1,opt,name=uniswap_pair_id,json=uniswapPairId,proto3" json:"uniswap_pair_id,omitempty" yaml:"uniswap_pair_id"`
	// amount of shares from the liquidity pool to redeem if current slippage > max slipage
	LiquidityPoolShares uint64 `protobuf:"varint,2,opt,name=liquidity_pool_shares,json=liquidityPoolShares,proto3" json:"liquidity_pool_shares,omitempty" yaml:"liquidity_pool_shares"`
	// max slippage allowed before the stoploss is triggered
	MaxSlippage github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=max_slippage,json=maxSlippage,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"max_slippage" yaml:"max_slippage"`
	// starting token pair ratio of the uniswap pool
	ReferencePairRatio github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=reference_pair_ratio,json=referencePairRatio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"reference_pair_ratio" yaml:"reference_pair_ratio"`
	// ethereum receiving address in hex format
	ReceiverAddress string `protobuf:"bytes,5,opt,name=receiver_address,json=receiverAddress,proto3" json:"receiver_address,omitempty" yaml:"receiver_address"`
}

func (m *Stoploss) Reset()         { *m = Stoploss{} }
func (m *Stoploss) String() string { return proto.CompactTextString(m) }
func (*Stoploss) ProtoMessage()    {}
func (*Stoploss) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccb151ade557e4dd, []int{0}
}
func (m *Stoploss) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Stoploss) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Stoploss.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Stoploss) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stoploss.Merge(m, src)
}
func (m *Stoploss) XXX_Size() int {
	return m.Size()
}
func (m *Stoploss) XXX_DiscardUnknown() {
	xxx_messageInfo_Stoploss.DiscardUnknown(m)
}

var xxx_messageInfo_Stoploss proto.InternalMessageInfo

func (m *Stoploss) GetUniswapPairID() string {
	if m != nil {
		return m.UniswapPairID
	}
	return ""
}

func (m *Stoploss) GetLiquidityPoolShares() uint64 {
	if m != nil {
		return m.LiquidityPoolShares
	}
	return 0
}

func (m *Stoploss) GetReceiverAddress() string {
	if m != nil {
		return m.ReceiverAddress
	}
	return ""
}

// Params define the impermanent loss module parameters
type Params struct {
	// contract address for impermanent loss handling on ethereum
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty" yaml:"contract_address"`
	// timeout block height value for the custom ethereum outgoing logic. This is value added to the last
	// seen ethereum height when executing the stoploss logic on EndBlock.
	EthTimeoutBlocks uint64 `protobuf:"varint,2,opt,name=eth_timeout_blocks,json=ethTimeoutBlocks,proto3" json:"eth_timeout_blocks,omitempty" yaml:"eth_timeout_blocks"`
	// timeout timestamp value for the redeemLiquidity deadline. This is value added to the block timestamp unix value
	// when executing the stoploss logic on EndBlock.
	EthTimeoutTimestamp uint64 `protobuf:"varint,3,opt,name=eth_timeout_timestamp,json=ethTimeoutTimestamp,proto3" json:"eth_timeout_timestamp,omitempty" yaml:"eth_timeout_timestamp"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccb151ade557e4dd, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *Params) GetEthTimeoutBlocks() uint64 {
	if m != nil {
		return m.EthTimeoutBlocks
	}
	return 0
}

func (m *Params) GetEthTimeoutTimestamp() uint64 {
	if m != nil {
		return m.EthTimeoutTimestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*Stoploss)(nil), "il.v1.Stoploss")
	proto.RegisterType((*Params)(nil), "il.v1.Params")
}

func init() { proto.RegisterFile("il/v1/il.proto", fileDescriptor_ccb151ade557e4dd) }

var fileDescriptor_ccb151ade557e4dd = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x3f, 0x6f, 0xd3, 0x40,
	0x18, 0xc6, 0x63, 0xfa, 0x47, 0x60, 0x28, 0xa9, 0x9c, 0x02, 0x86, 0x82, 0x1d, 0xdd, 0x00, 0x95,
	0x10, 0xb1, 0x2a, 0x36, 0xb6, 0x46, 0x05, 0x09, 0x21, 0xa4, 0xc8, 0x09, 0x0b, 0x8b, 0x75, 0xb1,
	0x0f, 0xfb, 0xe8, 0x5d, 0xde, 0xe3, 0xee, 0x1c, 0x92, 0x89, 0xaf, 0xc0, 0xc7, 0xea, 0xd8, 0x11,
	0x31, 0x58, 0x28, 0x59, 0x98, 0xcd, 0x17, 0x40, 0xb6, 0xe3, 0x90, 0x84, 0x2c, 0x9d, 0x2e, 0xf9,
	0xdd, 0xe3, 0xdf, 0x63, 0xfb, 0x3d, 0x9b, 0x77, 0x29, 0xf3, 0xc6, 0xa7, 0x1e, 0x65, 0x1d, 0x21,
	0x41, 0x83, 0xb5, 0x47, 0x59, 0x67, 0x7c, 0xfa, 0xe8, 0x28, 0x86, 0x18, 0x4a, 0xe2, 0x15, 0xbf,
	0xaa, 0x4d, 0xf4, 0x7b, 0xc7, 0xbc, 0xd9, 0xd7, 0x20, 0x18, 0x28, 0x65, 0xf5, 0xcd, 0x66, 0x3a,
	0xa2, 0xea, 0x2b, 0x16, 0x81, 0xc0, 0x54, 0x06, 0x34, 0xb2, 0x8d, 0xb6, 0x71, 0x72, 0xab, 0xfb,
	0x7c, 0x96, 0xb9, 0x07, 0x1f, 0xaa, 0xad, 0x1e, 0xa6, 0xf2, 0xed, 0x79, 0x9e, 0xb9, 0xf7, 0xa7,
	0x98, 0xb3, 0x57, 0x68, 0xe3, 0x0a, 0xe4, 0x1f, 0xa4, 0x2b, 0xc1, 0xc8, 0x1a, 0x98, 0xf7, 0x18,
	0xfd, 0x92, 0xd2, 0x88, 0xea, 0x69, 0x20, 0x00, 0x58, 0xa0, 0x12, 0x2c, 0x89, 0xb2, 0x6f, 0xb4,
	0x8d, 0x93, 0xdd, 0x6e, 0x3b, 0xcf, 0xdc, 0xc7, 0x95, 0x69, 0x6b, 0x0c, 0xf9, 0xad, 0x25, 0xef,
	0x01, 0xb0, 0x7e, 0x49, 0xad, 0xc4, 0xbc, 0xc3, 0xf1, 0x24, 0x50, 0x8c, 0x0a, 0x81, 0x63, 0x62,
	0xef, 0x94, 0xf7, 0xf9, 0xfa, 0x32, 0x73, 0x1b, 0x3f, 0x33, 0xf7, 0x69, 0x4c, 0x75, 0x92, 0x0e,
	0x3b, 0x21, 0x70, 0x2f, 0x04, 0xc5, 0x41, 0x2d, 0x96, 0x17, 0x2a, 0xba, 0xf0, 0xf4, 0x54, 0x10,
	0xd5, 0x39, 0x27, 0x61, 0x9e, 0xb9, 0xad, 0xaa, 0x7a, 0xd5, 0x85, 0xfc, 0xdb, 0x1c, 0x4f, 0xfa,
	0x8b, 0x7f, 0xd6, 0x37, 0xf3, 0x48, 0x92, 0x4f, 0x44, 0x92, 0x51, 0x48, 0xaa, 0x87, 0x94, 0x58,
	0x53, 0xb0, 0x77, 0xcb, 0xc6, 0xf7, 0xd7, 0x6e, 0x3c, 0xae, 0x1a, 0xb7, 0x39, 0x91, 0x6f, 0x2d,
	0x71, 0xf1, 0xf6, 0xfc, 0x02, 0x5a, 0x6f, 0xcc, 0x43, 0x49, 0x42, 0x42, 0xc7, 0x44, 0x06, 0x38,
	0x8a, 0x24, 0x51, 0xca, 0xde, 0x2b, 0xcb, 0x8f, 0xf3, 0xcc, 0x7d, 0x50, 0xeb, 0xd6, 0x13, 0xc8,
	0x6f, 0xd6, 0xe8, 0x6c, 0x41, 0xfe, 0x18, 0xe6, 0x7e, 0x0f, 0x4b, 0xcc, 0x55, 0xa1, 0x0c, 0x61,
	0xa4, 0x25, 0x0e, 0xf5, 0x52, 0x69, 0x6c, 0x2a, 0x37, 0x13, 0xc8, 0x6f, 0xd6, 0x68, 0xa1, 0xb4,
	0xde, 0x99, 0x16, 0xd1, 0x49, 0xa0, 0x29, 0x27, 0x90, 0xea, 0x60, 0xc8, 0x20, 0xbc, 0xa8, 0x07,
	0xfb, 0x24, 0xcf, 0xdc, 0x87, 0x95, 0xe9, 0xff, 0x0c, 0xf2, 0x0f, 0x89, 0x4e, 0x06, 0x15, 0xeb,
	0x96, 0xa8, 0x38, 0x28, 0xab, 0xc1, 0x62, 0x55, 0x1a, 0x73, 0x51, 0xce, 0x76, 0xed, 0xa0, 0x6c,
	0x8d, 0x21, 0xbf, 0xf5, 0x4f, 0x39, 0xa8, 0x69, 0xf7, 0xec, 0x72, 0xe6, 0x18, 0x57, 0x33, 0xc7,
	0xf8, 0x35, 0x73, 0x8c, 0xef, 0x73, 0xa7, 0x71, 0x35, 0x77, 0x1a, 0x3f, 0xe6, 0x4e, 0xe3, 0xe3,
	0xb3, 0x95, 0x91, 0x09, 0x12, 0xc7, 0xd3, 0xcf, 0x63, 0x4f, 0x01, 0xe7, 0x84, 0x51, 0x22, 0xbd,
	0x89, 0x47, 0x59, 0x35, 0xb7, 0xe1, 0x7e, 0xf9, 0xa9, 0xbc, 0xfc, 0x1b, 0x00, 0x00, 0xff, 0xff,
	0x19, 0x90, 0xcd, 0x77, 0x59, 0x03, 0x00, 0x00,
}

func (m *Stoploss) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Stoploss) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Stoploss) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ReceiverAddress) > 0 {
		i -= len(m.ReceiverAddress)
		copy(dAtA[i:], m.ReceiverAddress)
		i = encodeVarintIl(dAtA, i, uint64(len(m.ReceiverAddress)))
		i--
		dAtA[i] = 0x2a
	}
	{
		size := m.ReferencePairRatio.Size()
		i -= size
		if _, err := m.ReferencePairRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIl(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.MaxSlippage.Size()
		i -= size
		if _, err := m.MaxSlippage.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIl(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.LiquidityPoolShares != 0 {
		i = encodeVarintIl(dAtA, i, uint64(m.LiquidityPoolShares))
		i--
		dAtA[i] = 0x10
	}
	if len(m.UniswapPairID) > 0 {
		i -= len(m.UniswapPairID)
		copy(dAtA[i:], m.UniswapPairID)
		i = encodeVarintIl(dAtA, i, uint64(len(m.UniswapPairID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EthTimeoutTimestamp != 0 {
		i = encodeVarintIl(dAtA, i, uint64(m.EthTimeoutTimestamp))
		i--
		dAtA[i] = 0x18
	}
	if m.EthTimeoutBlocks != 0 {
		i = encodeVarintIl(dAtA, i, uint64(m.EthTimeoutBlocks))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintIl(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintIl(dAtA []byte, offset int, v uint64) int {
	offset -= sovIl(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Stoploss) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UniswapPairID)
	if l > 0 {
		n += 1 + l + sovIl(uint64(l))
	}
	if m.LiquidityPoolShares != 0 {
		n += 1 + sovIl(uint64(m.LiquidityPoolShares))
	}
	l = m.MaxSlippage.Size()
	n += 1 + l + sovIl(uint64(l))
	l = m.ReferencePairRatio.Size()
	n += 1 + l + sovIl(uint64(l))
	l = len(m.ReceiverAddress)
	if l > 0 {
		n += 1 + l + sovIl(uint64(l))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovIl(uint64(l))
	}
	if m.EthTimeoutBlocks != 0 {
		n += 1 + sovIl(uint64(m.EthTimeoutBlocks))
	}
	if m.EthTimeoutTimestamp != 0 {
		n += 1 + sovIl(uint64(m.EthTimeoutTimestamp))
	}
	return n
}

func sovIl(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIl(x uint64) (n int) {
	return sovIl(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Stoploss) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIl
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
			return fmt.Errorf("proto: Stoploss: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Stoploss: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UniswapPairID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIl
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
				return ErrInvalidLengthIl
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIl
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UniswapPairID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidityPoolShares", wireType)
			}
			m.LiquidityPoolShares = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LiquidityPoolShares |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSlippage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIl
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
				return ErrInvalidLengthIl
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIl
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxSlippage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReferencePairRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIl
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
				return ErrInvalidLengthIl
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIl
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReferencePairRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReceiverAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIl
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
				return ErrInvalidLengthIl
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIl
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReceiverAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIl
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthIl
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIl
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIl
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
				return ErrInvalidLengthIl
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIl
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthTimeoutBlocks", wireType)
			}
			m.EthTimeoutBlocks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EthTimeoutBlocks |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthTimeoutTimestamp", wireType)
			}
			m.EthTimeoutTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EthTimeoutTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipIl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIl
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthIl
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
func skipIl(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIl
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
					return 0, ErrIntOverflowIl
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
					return 0, ErrIntOverflowIl
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
				return 0, ErrInvalidLengthIl
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIl
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIl
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIl        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIl          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIl = fmt.Errorf("proto: unexpected end of group")
)