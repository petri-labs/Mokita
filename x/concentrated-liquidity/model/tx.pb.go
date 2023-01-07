// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mokita/concentrated-liquidity/pool-model/tx.proto

package model

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

// ===================== MsgCreateConcentratedPool
type MsgCreateConcentratedPool struct {
	Sender      string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty" yaml:"sender"`
	Denom0      string `protobuf:"bytes,2,opt,name=denom0,proto3" json:"denom0,omitempty" yaml:"denom0"`
	Denom1      string `protobuf:"bytes,3,opt,name=denom1,proto3" json:"denom1,omitempty" yaml:"denom1"`
	TickSpacing uint64 `protobuf:"varint,4,opt,name=tick_spacing,json=tickSpacing,proto3" json:"tick_spacing,omitempty" yaml:"tick_spacing"`
}

func (m *MsgCreateConcentratedPool) Reset()         { *m = MsgCreateConcentratedPool{} }
func (m *MsgCreateConcentratedPool) String() string { return proto.CompactTextString(m) }
func (*MsgCreateConcentratedPool) ProtoMessage()    {}
func (*MsgCreateConcentratedPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c324e8c9dd2851d, []int{0}
}
func (m *MsgCreateConcentratedPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateConcentratedPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateConcentratedPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateConcentratedPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateConcentratedPool.Merge(m, src)
}
func (m *MsgCreateConcentratedPool) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateConcentratedPool) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateConcentratedPool.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateConcentratedPool proto.InternalMessageInfo

func (m *MsgCreateConcentratedPool) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgCreateConcentratedPool) GetDenom0() string {
	if m != nil {
		return m.Denom0
	}
	return ""
}

func (m *MsgCreateConcentratedPool) GetDenom1() string {
	if m != nil {
		return m.Denom1
	}
	return ""
}

func (m *MsgCreateConcentratedPool) GetTickSpacing() uint64 {
	if m != nil {
		return m.TickSpacing
	}
	return 0
}

// Returns a unique poolID to identify the pool with.
type MsgCreateConcentratedPoolResponse struct {
	PoolID uint64 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
}

func (m *MsgCreateConcentratedPoolResponse) Reset()         { *m = MsgCreateConcentratedPoolResponse{} }
func (m *MsgCreateConcentratedPoolResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateConcentratedPoolResponse) ProtoMessage()    {}
func (*MsgCreateConcentratedPoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c324e8c9dd2851d, []int{1}
}
func (m *MsgCreateConcentratedPoolResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateConcentratedPoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateConcentratedPoolResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateConcentratedPoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateConcentratedPoolResponse.Merge(m, src)
}
func (m *MsgCreateConcentratedPoolResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateConcentratedPoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateConcentratedPoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateConcentratedPoolResponse proto.InternalMessageInfo

func (m *MsgCreateConcentratedPoolResponse) GetPoolID() uint64 {
	if m != nil {
		return m.PoolID
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgCreateConcentratedPool)(nil), "mokita.concentratedliquidity.v1beta1.MsgCreateConcentratedPool")
	proto.RegisterType((*MsgCreateConcentratedPoolResponse)(nil), "mokita.concentratedliquidity.v1beta1.MsgCreateConcentratedPoolResponse")
}

func init() {
	proto.RegisterFile("mokita/concentrated-liquidity/pool-model/tx.proto", fileDescriptor_6c324e8c9dd2851d)
}

var fileDescriptor_6c324e8c9dd2851d = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xcd, 0x6a, 0xdb, 0x40,
	0x10, 0x80, 0xbd, 0xad, 0x51, 0xe9, 0xb6, 0x3d, 0x54, 0x2d, 0xad, 0xea, 0x83, 0xec, 0xaa, 0x14,
	0x9c, 0x83, 0xb5, 0x91, 0x7d, 0xf3, 0x29, 0xc8, 0x39, 0xd8, 0x87, 0x40, 0x50, 0x6e, 0x21, 0x60,
	0xf4, 0xb3, 0x28, 0x4b, 0x24, 0x8d, 0xa2, 0x5d, 0x1b, 0xfb, 0x2d, 0xf2, 0x10, 0x79, 0x83, 0xbc,
	0x44, 0x8e, 0x3e, 0xfa, 0x64, 0x82, 0xfc, 0x06, 0x7e, 0x82, 0xa0, 0x3f, 0x63, 0x42, 0x04, 0x81,
	0xdc, 0x34, 0x3b, 0xdf, 0xb7, 0xa3, 0x99, 0x1d, 0xdc, 0x07, 0x1e, 0x02, 0x67, 0x9c, 0xb8, 0x10,
	0xb9, 0x34, 0x12, 0x89, 0x2d, 0xa8, 0xd7, 0x0b, 0xd8, 0xed, 0x8c, 0x79, 0x4c, 0x2c, 0x49, 0x0c,
	0x10, 0xf4, 0x42, 0xf0, 0x68, 0x40, 0xc4, 0x42, 0x8f, 0x13, 0x10, 0x20, 0xff, 0x2f, 0x1d, 0xfd,
	0xd0, 0xd9, 0x2b, 0xfa, 0xdc, 0x70, 0xa8, 0xb0, 0x8d, 0xd6, 0x4f, 0x1f, 0x7c, 0xc8, 0x0d, 0x92,
	0x7d, 0x15, 0x72, 0x4b, 0x75, 0x73, 0x9b, 0x38, 0x36, 0xa7, 0xa4, 0x44, 0x89, 0x0b, 0x2c, 0x2a,
	0xf2, 0xda, 0x1a, 0xe1, 0x3f, 0x67, 0xdc, 0x1f, 0x25, 0xd4, 0x16, 0x74, 0x74, 0x50, 0xe0, 0x1c,
	0x20, 0x90, 0x8f, 0xb0, 0xc4, 0x69, 0xe4, 0xd1, 0x44, 0x41, 0x1d, 0xd4, 0xfd, 0x6c, 0x7e, 0xdf,
	0x6d, 0xda, 0xdf, 0x96, 0x76, 0x18, 0x0c, 0xb5, 0xe2, 0x5c, 0xb3, 0x4a, 0x20, 0x43, 0x3d, 0x1a,
	0x41, 0x78, 0xac, 0x7c, 0x78, 0x89, 0x16, 0xe7, 0x9a, 0x55, 0x02, 0x7b, 0xd4, 0x50, 0x3e, 0xbe,
	0x8a, 0x1a, 0x15, 0x6a, 0xc8, 0x43, 0xfc, 0x55, 0x30, 0xf7, 0x66, 0xca, 0x63, 0xdb, 0x65, 0x91,
	0xaf, 0x34, 0x3b, 0xa8, 0xdb, 0x34, 0x7f, 0xef, 0x36, 0xed, 0x1f, 0x85, 0x70, 0x98, 0xd5, 0xac,
	0x2f, 0x59, 0x78, 0x51, 0x46, 0x63, 0xfc, 0xb7, 0xb6, 0x33, 0x8b, 0xf2, 0x18, 0x22, 0x4e, 0xe5,
	0x7f, 0xf8, 0x53, 0x36, 0xf3, 0x29, 0xf3, 0xf2, 0x16, 0x9b, 0x26, 0x4e, 0x37, 0x6d, 0x29, 0x43,
	0x26, 0xa7, 0x96, 0x94, 0xa5, 0x26, 0x5e, 0xff, 0x01, 0x61, 0x5c, 0x5d, 0x05, 0x89, 0x7c, 0x8f,
	0xf0, 0xaf, 0x9a, 0x81, 0x9d, 0xe8, 0x6f, 0x7a, 0x2c, 0xbd, 0xf6, 0xc7, 0x5a, 0xe3, 0xf7, 0xde,
	0x50, 0xb5, 0x66, 0x5e, 0x3d, 0xa6, 0x2a, 0x5a, 0xa5, 0x2a, 0x7a, 0x4a, 0x55, 0x74, 0xb7, 0x55,
	0x1b, 0xab, 0xad, 0xda, 0x58, 0x6f, 0xd5, 0xc6, 0xa5, 0xe9, 0x33, 0x71, 0x3d, 0x73, 0x74, 0x17,
	0x42, 0x52, 0x56, 0xeb, 0x05, 0xb6, 0xc3, 0xab, 0x80, 0xcc, 0x8d, 0x01, 0x59, 0xd4, 0xed, 0x68,
	0xbe, 0x9e, 0x8e, 0x94, 0xef, 0xcf, 0xe0, 0x39, 0x00, 0x00, 0xff, 0xff, 0x87, 0x0e, 0x15, 0x36,
	0xd2, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgCreatorClient is the client API for MsgCreator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgCreatorClient interface {
	CreateConcentratedPool(ctx context.Context, in *MsgCreateConcentratedPool, opts ...grpc.CallOption) (*MsgCreateConcentratedPoolResponse, error)
}

type msgCreatorClient struct {
	cc grpc1.ClientConn
}

func NewMsgCreatorClient(cc grpc1.ClientConn) MsgCreatorClient {
	return &msgCreatorClient{cc}
}

func (c *msgCreatorClient) CreateConcentratedPool(ctx context.Context, in *MsgCreateConcentratedPool, opts ...grpc.CallOption) (*MsgCreateConcentratedPoolResponse, error) {
	out := new(MsgCreateConcentratedPoolResponse)
	err := c.cc.Invoke(ctx, "/mokita.concentratedliquidity.v1beta1.MsgCreator/CreateConcentratedPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgCreatorServer is the server API for MsgCreator service.
type MsgCreatorServer interface {
	CreateConcentratedPool(context.Context, *MsgCreateConcentratedPool) (*MsgCreateConcentratedPoolResponse, error)
}

// UnimplementedMsgCreatorServer can be embedded to have forward compatible implementations.
type UnimplementedMsgCreatorServer struct {
}

func (*UnimplementedMsgCreatorServer) CreateConcentratedPool(ctx context.Context, req *MsgCreateConcentratedPool) (*MsgCreateConcentratedPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConcentratedPool not implemented")
}

func RegisterMsgCreatorServer(s grpc1.Server, srv MsgCreatorServer) {
	s.RegisterService(&_MsgCreator_serviceDesc, srv)
}

func _MsgCreator_CreateConcentratedPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateConcentratedPool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgCreatorServer).CreateConcentratedPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mokita.concentratedliquidity.v1beta1.MsgCreator/CreateConcentratedPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgCreatorServer).CreateConcentratedPool(ctx, req.(*MsgCreateConcentratedPool))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgCreator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mokita.concentratedliquidity.v1beta1.MsgCreator",
	HandlerType: (*MsgCreatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateConcentratedPool",
			Handler:    _MsgCreator_CreateConcentratedPool_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mokita/concentrated-liquidity/pool-model/tx.proto",
}

func (m *MsgCreateConcentratedPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateConcentratedPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateConcentratedPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TickSpacing != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.TickSpacing))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Denom1) > 0 {
		i -= len(m.Denom1)
		copy(dAtA[i:], m.Denom1)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Denom1)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Denom0) > 0 {
		i -= len(m.Denom0)
		copy(dAtA[i:], m.Denom0)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Denom0)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateConcentratedPoolResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateConcentratedPoolResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateConcentratedPoolResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PoolID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.PoolID))
		i--
		dAtA[i] = 0x8
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
func (m *MsgCreateConcentratedPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Denom0)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Denom1)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.TickSpacing != 0 {
		n += 1 + sovTx(uint64(m.TickSpacing))
	}
	return n
}

func (m *MsgCreateConcentratedPoolResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolID != 0 {
		n += 1 + sovTx(uint64(m.PoolID))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateConcentratedPool) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateConcentratedPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateConcentratedPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
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
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom0", wireType)
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
			m.Denom0 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom1", wireType)
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
			m.Denom1 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TickSpacing", wireType)
			}
			m.TickSpacing = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TickSpacing |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *MsgCreateConcentratedPoolResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateConcentratedPoolResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateConcentratedPoolResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolID", wireType)
			}
			m.PoolID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
