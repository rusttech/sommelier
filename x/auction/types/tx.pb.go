// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: auction/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

type MsgSubmitBidRequest struct {
	AuctionId     uint32     `protobuf:"varint,1,opt,name=auction_id,json=auctionId,proto3" json:"auction_id,omitempty"`
	Bidder        string     `protobuf:"bytes,2,opt,name=bidder,proto3" json:"bidder,omitempty"`
	MaxBid        types.Coin `protobuf:"bytes,3,opt,name=max_bid,json=maxBid,proto3" json:"max_bid"`
	MinimumAmount types.Coin `protobuf:"bytes,4,opt,name=minimum_amount,json=minimumAmount,proto3" json:"minimum_amount"`
	Signer        string     `protobuf:"bytes,5,opt,name=signer,proto3" json:"signer,omitempty"`
}

func (m *MsgSubmitBidRequest) Reset()         { *m = MsgSubmitBidRequest{} }
func (m *MsgSubmitBidRequest) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitBidRequest) ProtoMessage()    {}
func (*MsgSubmitBidRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a60fe804de30894a, []int{0}
}
func (m *MsgSubmitBidRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitBidRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitBidRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitBidRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitBidRequest.Merge(m, src)
}
func (m *MsgSubmitBidRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitBidRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitBidRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitBidRequest proto.InternalMessageInfo

func (m *MsgSubmitBidRequest) GetAuctionId() uint32 {
	if m != nil {
		return m.AuctionId
	}
	return 0
}

func (m *MsgSubmitBidRequest) GetBidder() string {
	if m != nil {
		return m.Bidder
	}
	return ""
}

func (m *MsgSubmitBidRequest) GetMaxBid() types.Coin {
	if m != nil {
		return m.MaxBid
	}
	return types.Coin{}
}

func (m *MsgSubmitBidRequest) GetMinimumAmount() types.Coin {
	if m != nil {
		return m.MinimumAmount
	}
	return types.Coin{}
}

func (m *MsgSubmitBidRequest) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

type MsgSubmitBidResponse struct {
	Bid *Bid `protobuf:"bytes,1,opt,name=bid,proto3" json:"bid,omitempty"`
}

func (m *MsgSubmitBidResponse) Reset()         { *m = MsgSubmitBidResponse{} }
func (m *MsgSubmitBidResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitBidResponse) ProtoMessage()    {}
func (*MsgSubmitBidResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a60fe804de30894a, []int{1}
}
func (m *MsgSubmitBidResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitBidResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitBidResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitBidResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitBidResponse.Merge(m, src)
}
func (m *MsgSubmitBidResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitBidResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitBidResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitBidResponse proto.InternalMessageInfo

func (m *MsgSubmitBidResponse) GetBid() *Bid {
	if m != nil {
		return m.Bid
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgSubmitBidRequest)(nil), "auction.v1.MsgSubmitBidRequest")
	proto.RegisterType((*MsgSubmitBidResponse)(nil), "auction.v1.MsgSubmitBidResponse")
}

func init() { proto.RegisterFile("auction/v1/tx.proto", fileDescriptor_a60fe804de30894a) }

var fileDescriptor_a60fe804de30894a = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4f, 0xcb, 0xd3, 0x30,
	0x18, 0x6f, 0xdc, 0xeb, 0x64, 0x79, 0x99, 0x42, 0x36, 0xa4, 0x1b, 0xd8, 0xd5, 0x9d, 0x76, 0x4a,
	0xe8, 0xf4, 0xa0, 0x47, 0x2b, 0x08, 0x3b, 0xcc, 0x43, 0xc5, 0x8b, 0x97, 0xd2, 0xb4, 0x21, 0x46,
	0x4c, 0x53, 0x9b, 0xb4, 0x74, 0xdf, 0xc2, 0x8f, 0xb5, 0xe3, 0x8e, 0x9e, 0x44, 0x36, 0xf0, 0x73,
	0x48, 0xdb, 0x54, 0x27, 0x28, 0xbc, 0xb7, 0xfe, 0xfe, 0x35, 0xbf, 0x27, 0x4f, 0xe0, 0x2c, 0xa9,
	0x52, 0x23, 0x54, 0x4e, 0xea, 0x80, 0x98, 0x06, 0x17, 0xa5, 0x32, 0x0a, 0x41, 0x4b, 0xe2, 0x3a,
	0x58, 0xba, 0x57, 0x86, 0x81, 0xee, 0x5c, 0x4b, 0x2f, 0x55, 0x5a, 0x2a, 0x4d, 0x68, 0xa2, 0x19,
	0xa9, 0x03, 0xca, 0x4c, 0x12, 0x90, 0x54, 0x89, 0x41, 0x5f, 0xf4, 0x7a, 0xdc, 0x21, 0xd2, 0x03,
	0x2b, 0xcd, 0xb9, 0xe2, 0xaa, 0xe7, 0xdb, 0xaf, 0x9e, 0x5d, 0xff, 0x04, 0x70, 0xb6, 0xd7, 0xfc,
	0x5d, 0x45, 0xa5, 0x30, 0xa1, 0xc8, 0x22, 0xf6, 0xa5, 0x62, 0xda, 0xa0, 0x27, 0x70, 0x28, 0x14,
	0x8b, 0xcc, 0x05, 0x3e, 0xd8, 0x4c, 0xa3, 0x89, 0x65, 0x76, 0x19, 0x7a, 0x0c, 0xc7, 0x54, 0x64,
	0x19, 0x2b, 0xdd, 0x7b, 0x3e, 0xd8, 0x4c, 0x22, 0x8b, 0xd0, 0x0b, 0xf8, 0x40, 0x26, 0x4d, 0x4c,
	0x45, 0xe6, 0x8e, 0x7c, 0xb0, 0xb9, 0xdd, 0x2e, 0xb0, 0x2d, 0xd1, 0x36, 0xc6, 0xb6, 0x31, 0x7e,
	0xad, 0x44, 0x1e, 0xde, 0x1c, 0xbf, 0xaf, 0x9c, 0x68, 0x2c, 0x93, 0x26, 0x14, 0x19, 0x7a, 0x03,
	0x1f, 0x4a, 0x91, 0x0b, 0x59, 0xc9, 0x38, 0x91, 0xaa, 0xca, 0x8d, 0x7b, 0x73, 0xb7, 0x1f, 0x4c,
	0x6d, 0xec, 0x55, 0x97, 0x6a, 0x9b, 0x69, 0xc1, 0x73, 0x56, 0xba, 0xf7, 0xfb, 0x66, 0x3d, 0x5a,
	0xbf, 0x84, 0xf3, 0xbf, 0xe7, 0xd4, 0x85, 0xca, 0x35, 0x43, 0x4f, 0xe1, 0x88, 0xda, 0x09, 0x6f,
	0xb7, 0x8f, 0xf0, 0x9f, 0x2d, 0xe0, 0xd6, 0xd5, 0x6a, 0xdb, 0xf7, 0x70, 0xb4, 0xd7, 0x1c, 0xbd,
	0x85, 0x93, 0xdf, 0x71, 0xb4, 0xba, 0x76, 0xfe, 0xe3, 0x02, 0x97, 0xfe, 0xff, 0x0d, 0xfd, 0xc9,
	0xe1, 0xee, 0x78, 0xf6, 0xc0, 0xe9, 0xec, 0x81, 0x1f, 0x67, 0x0f, 0x7c, 0xbd, 0x78, 0xce, 0xe9,
	0xe2, 0x39, 0xdf, 0x2e, 0x9e, 0xf3, 0x81, 0x70, 0x61, 0x3e, 0x56, 0x14, 0xa7, 0x4a, 0x92, 0x82,
	0x71, 0x7e, 0xf8, 0x54, 0x13, 0xad, 0xa4, 0x64, 0x9f, 0x05, 0x2b, 0x49, 0xfd, 0x9c, 0x34, 0xc3,
	0xb3, 0x20, 0xe6, 0x50, 0x30, 0x4d, 0xc7, 0xdd, 0x32, 0x9f, 0xfd, 0x0a, 0x00, 0x00, 0xff, 0xff,
	0x2b, 0xee, 0xc6, 0x6a, 0x5a, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	SubmitBid(ctx context.Context, in *MsgSubmitBidRequest, opts ...grpc.CallOption) (*MsgSubmitBidResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SubmitBid(ctx context.Context, in *MsgSubmitBidRequest, opts ...grpc.CallOption) (*MsgSubmitBidResponse, error) {
	out := new(MsgSubmitBidResponse)
	err := c.cc.Invoke(ctx, "/auction.v1.Msg/SubmitBid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	SubmitBid(context.Context, *MsgSubmitBidRequest) (*MsgSubmitBidResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SubmitBid(ctx context.Context, req *MsgSubmitBidRequest) (*MsgSubmitBidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitBid not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SubmitBid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitBidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitBid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auction.v1.Msg/SubmitBid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitBid(ctx, req.(*MsgSubmitBidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auction.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitBid",
			Handler:    _Msg_SubmitBid_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auction/v1/tx.proto",
}

func (m *MsgSubmitBidRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitBidRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitBidRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0x2a
	}
	{
		size, err := m.MinimumAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.MaxBid.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Bidder) > 0 {
		i -= len(m.Bidder)
		copy(dAtA[i:], m.Bidder)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Bidder)))
		i--
		dAtA[i] = 0x12
	}
	if m.AuctionId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.AuctionId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgSubmitBidResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitBidResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitBidResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Bid != nil {
		{
			size, err := m.Bid.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSubmitBidRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AuctionId != 0 {
		n += 1 + sovTx(uint64(m.AuctionId))
	}
	l = len(m.Bidder)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.MaxBid.Size()
	n += 1 + l + sovTx(uint64(l))
	l = m.MinimumAmount.Size()
	n += 1 + l + sovTx(uint64(l))
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSubmitBidResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Bid != nil {
		l = m.Bid.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSubmitBidRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSubmitBidRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitBidRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionId", wireType)
			}
			m.AuctionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuctionId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bidder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bidder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxBid", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxBid.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinimumAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinimumAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSubmitBidResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSubmitBidResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitBidResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bid", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Bid == nil {
				m.Bid = &Bid{}
			}
			if err := m.Bid.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
