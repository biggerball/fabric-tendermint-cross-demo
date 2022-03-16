// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: erc20mgr/msgs.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_datachainlab_cross_x_core_auth_types "github.com/datachainlab/cross/x/core/auth/types"
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

type MsgContractCallTx struct {
	Request *ContractCallRequest                                        `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Signers []github_com_datachainlab_cross_x_core_auth_types.AccountID `protobuf:"bytes,2,rep,name=signers,proto3,casttype=github.com/datachainlab/cross/x/core/auth/types.AccountID" json:"signers,omitempty"`
}

func (m *MsgContractCallTx) Reset()         { *m = MsgContractCallTx{} }
func (m *MsgContractCallTx) String() string { return proto.CompactTextString(m) }
func (*MsgContractCallTx) ProtoMessage()    {}
func (*MsgContractCallTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d777cf67545a374, []int{0}
}
func (m *MsgContractCallTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgContractCallTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgContractCallTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgContractCallTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgContractCallTx.Merge(m, src)
}
func (m *MsgContractCallTx) XXX_Size() int {
	return m.Size()
}
func (m *MsgContractCallTx) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgContractCallTx.DiscardUnknown(m)
}

var xxx_messageInfo_MsgContractCallTx proto.InternalMessageInfo

// MsgSignTxResponse defines the Msg/SignTx response type.
type MsgContractCallTxResponse struct {
}

func (m *MsgContractCallTxResponse) Reset()         { *m = MsgContractCallTxResponse{} }
func (m *MsgContractCallTxResponse) String() string { return proto.CompactTextString(m) }
func (*MsgContractCallTxResponse) ProtoMessage()    {}
func (*MsgContractCallTxResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d777cf67545a374, []int{1}
}
func (m *MsgContractCallTxResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgContractCallTxResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgContractCallTxResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgContractCallTxResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgContractCallTxResponse.Merge(m, src)
}
func (m *MsgContractCallTxResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgContractCallTxResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgContractCallTxResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgContractCallTxResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgContractCallTx)(nil), "erc20mgr.MsgContractCallTx")
	proto.RegisterType((*MsgContractCallTxResponse)(nil), "erc20mgr.MsgContractCallTxResponse")
}

func init() { proto.RegisterFile("erc20mgr/msgs.proto", fileDescriptor_9d777cf67545a374) }

var fileDescriptor_9d777cf67545a374 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xb1, 0x4e, 0xc2, 0x40,
	0x18, 0xc7, 0x5b, 0x49, 0x84, 0x54, 0x63, 0x62, 0x75, 0x00, 0x8c, 0x07, 0xc1, 0x98, 0xb0, 0xd0,
	0x33, 0x38, 0x18, 0x4d, 0x1c, 0x04, 0x17, 0x07, 0x96, 0xc6, 0xc4, 0xa8, 0xd3, 0xf5, 0x7a, 0x1c,
	0x4d, 0xda, 0xfb, 0xf0, 0xee, 0x9a, 0xc0, 0x1b, 0x38, 0xfa, 0x08, 0xbc, 0x81, 0xaf, 0xc1, 0xc8,
	0xe8, 0x64, 0x14, 0x16, 0x9f, 0xc1, 0xc9, 0x58, 0x5a, 0x30, 0x12, 0xdd, 0xbe, 0xfe, 0xfb, 0xbb,
	0xff, 0xf7, 0xff, 0xe7, 0xb3, 0x76, 0x98, 0xa4, 0xcd, 0xa3, 0x88, 0x4b, 0x1c, 0x29, 0xae, 0x9c,
	0xbe, 0x04, 0x0d, 0x76, 0x21, 0x13, 0xcb, 0x25, 0x0e, 0xc0, 0x43, 0x86, 0x13, 0xdd, 0x8b, 0xbb,
	0x98, 0x88, 0xe1, 0x1c, 0x2a, 0xef, 0x72, 0xe0, 0x90, 0x8c, 0xf8, 0x7b, 0xca, 0xd4, 0x85, 0x9f,
	0x1e, 0xf6, 0x59, 0x6a, 0x58, 0x7b, 0x36, 0xad, 0xed, 0x8e, 0xe2, 0x6d, 0x10, 0x5a, 0x12, 0xaa,
	0xdb, 0x24, 0x0c, 0xaf, 0x07, 0xf6, 0x89, 0x95, 0x97, 0xec, 0x21, 0x66, 0x4a, 0x17, 0xcd, 0xaa,
	0x59, 0xdf, 0x68, 0xee, 0x3b, 0xd9, 0x6b, 0xe7, 0x27, 0xea, 0xce, 0x21, 0x37, 0xa3, 0xed, 0x1b,
	0x2b, 0xaf, 0x02, 0x2e, 0x98, 0x54, 0xc5, 0xb5, 0x6a, 0xae, 0xbe, 0xd9, 0x3a, 0xff, 0x7c, 0xad,
	0x9c, 0xf2, 0x40, 0xf7, 0x62, 0xcf, 0xa1, 0x10, 0x61, 0x9f, 0x68, 0x42, 0x7b, 0x24, 0x10, 0x21,
	0xf1, 0x30, 0x95, 0xa0, 0x14, 0x1e, 0x60, 0x0a, 0x92, 0x61, 0x12, 0xeb, 0x5e, 0x1a, 0xed, 0x82,
	0x52, 0x88, 0x85, 0xbe, 0xba, 0x74, 0x33, 0xb7, 0xb3, 0xc2, 0xe3, 0xa8, 0x62, 0x7c, 0x8c, 0x2a,
	0x46, 0xed, 0xd0, 0x2a, 0xad, 0x04, 0x76, 0x99, 0xea, 0x83, 0x50, 0x6c, 0x89, 0x35, 0x6f, 0xad,
	0x5c, 0x47, 0x71, 0xdb, 0xb5, 0xb6, 0x7e, 0x75, 0xdb, 0x5b, 0x56, 0x59, 0xf1, 0x29, 0x1f, 0xfc,
	0xf3, 0x33, 0x5b, 0xd2, 0x1a, 0x8e, 0xdf, 0x91, 0x31, 0x9e, 0x22, 0x73, 0x32, 0x45, 0xe6, 0xdb,
	0x14, 0x99, 0x4f, 0x33, 0x64, 0x4c, 0x66, 0xc8, 0x78, 0x99, 0x21, 0xe3, 0xee, 0xfe, 0xaf, 0xb6,
	0x5d, 0xe2, 0xc9, 0x80, 0x36, 0x34, 0x13, 0x3e, 0x93, 0x51, 0x20, 0x74, 0x23, 0xe9, 0xdf, 0xf0,
	0x59, 0x04, 0x98, 0xa6, 0xab, 0x14, 0x4e, 0x02, 0xe0, 0x08, 0xfc, 0x38, 0x64, 0xe9, 0xd7, 0xe2,
	0x68, 0xde, 0x7a, 0x72, 0xb5, 0xe3, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xea, 0x3f, 0x32, 0x10,
	0x1d, 0x02, 0x00, 0x00,
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
	ContractCallTx(ctx context.Context, in *MsgContractCallTx, opts ...grpc.CallOption) (*MsgContractCallTxResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) ContractCallTx(ctx context.Context, in *MsgContractCallTx, opts ...grpc.CallOption) (*MsgContractCallTxResponse, error) {
	out := new(MsgContractCallTxResponse)
	err := c.cc.Invoke(ctx, "/erc20mgr.Msg/ContractCallTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	ContractCallTx(context.Context, *MsgContractCallTx) (*MsgContractCallTxResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) ContractCallTx(ctx context.Context, req *MsgContractCallTx) (*MsgContractCallTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContractCallTx not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_ContractCallTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgContractCallTx)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ContractCallTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/erc20mgr.Msg/ContractCallTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ContractCallTx(ctx, req.(*MsgContractCallTx))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "erc20mgr.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ContractCallTx",
			Handler:    _Msg_ContractCallTx_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "erc20mgr/msgs.proto",
}

func (m *MsgContractCallTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgContractCallTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgContractCallTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signers) > 0 {
		for iNdEx := len(m.Signers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Signers[iNdEx])
			copy(dAtA[i:], m.Signers[iNdEx])
			i = encodeVarintMsgs(dAtA, i, uint64(len(m.Signers[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Request != nil {
		{
			size, err := m.Request.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMsgs(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgContractCallTxResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgContractCallTxResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgContractCallTxResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintMsgs(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsgs(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgContractCallTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Request != nil {
		l = m.Request.Size()
		n += 1 + l + sovMsgs(uint64(l))
	}
	if len(m.Signers) > 0 {
		for _, b := range m.Signers {
			l = len(b)
			n += 1 + l + sovMsgs(uint64(l))
		}
	}
	return n
}

func (m *MsgContractCallTxResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovMsgs(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsgs(x uint64) (n int) {
	return sovMsgs(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgContractCallTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgs
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
			return fmt.Errorf("proto: MsgContractCallTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgContractCallTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Request", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
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
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Request == nil {
				m.Request = &ContractCallRequest{}
			}
			if err := m.Request.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signers", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
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
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signers = append(m.Signers, make([]byte, postIndex-iNdEx))
			copy(m.Signers[len(m.Signers)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsgs
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMsgs
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
func (m *MsgContractCallTxResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgs
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
			return fmt.Errorf("proto: MsgContractCallTxResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgContractCallTxResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsgs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsgs
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMsgs
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
func skipMsgs(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgs
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
					return 0, ErrIntOverflowMsgs
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
					return 0, ErrIntOverflowMsgs
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
				return 0, ErrInvalidLengthMsgs
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsgs
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsgs
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsgs        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgs          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsgs = fmt.Errorf("proto: unexpected end of group")
)
