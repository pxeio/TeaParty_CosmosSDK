// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: partychain/party/orders_under_watch.proto

package types

import (
	fmt "fmt"
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

type OrdersUnderWatch struct {
	Index            string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	NknAddress       string `protobuf:"bytes,2,opt,name=nknAddress,proto3" json:"nknAddress,omitempty"`
	WalletPrivateKey string `protobuf:"bytes,3,opt,name=walletPrivateKey,proto3" json:"walletPrivateKey,omitempty"`
	WalletPublicKey  string `protobuf:"bytes,4,opt,name=walletPublicKey,proto3" json:"walletPublicKey,omitempty"`
	ShippingAddress  string `protobuf:"bytes,5,opt,name=shippingAddress,proto3" json:"shippingAddress,omitempty"`
	RefundAddress    string `protobuf:"bytes,6,opt,name=refundAddress,proto3" json:"refundAddress,omitempty"`
	Amount           string `protobuf:"bytes,7,opt,name=amount,proto3" json:"amount,omitempty"`
	Chain            string `protobuf:"bytes,8,opt,name=chain,proto3" json:"chain,omitempty"`
	PaymentComplete  bool   `protobuf:"varint,9,opt,name=paymentComplete,proto3" json:"paymentComplete,omitempty"`
}

func (m *OrdersUnderWatch) Reset()         { *m = OrdersUnderWatch{} }
func (m *OrdersUnderWatch) String() string { return proto.CompactTextString(m) }
func (*OrdersUnderWatch) ProtoMessage()    {}
func (*OrdersUnderWatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d0c17beedce3b4c, []int{0}
}
func (m *OrdersUnderWatch) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OrdersUnderWatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OrdersUnderWatch.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OrdersUnderWatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrdersUnderWatch.Merge(m, src)
}
func (m *OrdersUnderWatch) XXX_Size() int {
	return m.Size()
}
func (m *OrdersUnderWatch) XXX_DiscardUnknown() {
	xxx_messageInfo_OrdersUnderWatch.DiscardUnknown(m)
}

var xxx_messageInfo_OrdersUnderWatch proto.InternalMessageInfo

func (m *OrdersUnderWatch) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *OrdersUnderWatch) GetNknAddress() string {
	if m != nil {
		return m.NknAddress
	}
	return ""
}

func (m *OrdersUnderWatch) GetWalletPrivateKey() string {
	if m != nil {
		return m.WalletPrivateKey
	}
	return ""
}

func (m *OrdersUnderWatch) GetWalletPublicKey() string {
	if m != nil {
		return m.WalletPublicKey
	}
	return ""
}

func (m *OrdersUnderWatch) GetShippingAddress() string {
	if m != nil {
		return m.ShippingAddress
	}
	return ""
}

func (m *OrdersUnderWatch) GetRefundAddress() string {
	if m != nil {
		return m.RefundAddress
	}
	return ""
}

func (m *OrdersUnderWatch) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *OrdersUnderWatch) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *OrdersUnderWatch) GetPaymentComplete() bool {
	if m != nil {
		return m.PaymentComplete
	}
	return false
}

func init() {
	proto.RegisterType((*OrdersUnderWatch)(nil), "teapartycrypto.partychain.party.OrdersUnderWatch")
}

func init() {
	proto.RegisterFile("partychain/party/orders_under_watch.proto", fileDescriptor_1d0c17beedce3b4c)
}

var fileDescriptor_1d0c17beedce3b4c = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xc1, 0x4a, 0x3b, 0x31,
	0x10, 0xc6, 0xbb, 0xfd, 0xff, 0x5b, 0xdb, 0x80, 0x58, 0x82, 0xc8, 0x9e, 0x62, 0x11, 0x0f, 0xd5,
	0x43, 0x17, 0xf4, 0x09, 0xb4, 0x47, 0x05, 0x4b, 0x51, 0x04, 0x2f, 0x25, 0xdd, 0x8c, 0xdd, 0xe0,
	0x6e, 0x12, 0xb2, 0xb3, 0xb6, 0xfb, 0x16, 0x3e, 0x95, 0x78, 0xec, 0xd1, 0xa3, 0xb4, 0x2f, 0x22,
	0xc9, 0xb6, 0x58, 0xeb, 0x6d, 0xe6, 0x9b, 0xdf, 0xc7, 0x30, 0xf3, 0x91, 0x33, 0xc3, 0x2d, 0x96,
	0x71, 0xc2, 0xa5, 0x8a, 0x7c, 0x19, 0x69, 0x2b, 0xc0, 0xe6, 0xe3, 0x42, 0x09, 0xb0, 0xe3, 0x19,
	0xc7, 0x38, 0xe9, 0x1b, 0xab, 0x51, 0xd3, 0x63, 0x04, 0x5e, 0xd1, 0xb6, 0x34, 0xa8, 0xfb, 0x3f,
	0xce, 0xaa, 0x3c, 0x79, 0xaf, 0x93, 0xce, 0x9d, 0x77, 0x3f, 0x38, 0xf3, 0xa3, 0xf3, 0xd2, 0x43,
	0xd2, 0x90, 0x4a, 0xc0, 0x3c, 0x0c, 0xba, 0x41, 0xaf, 0x3d, 0xaa, 0x1a, 0xca, 0x08, 0x51, 0x2f,
	0xea, 0x4a, 0x08, 0x0b, 0x79, 0x1e, 0xd6, 0xfd, 0x68, 0x4b, 0xa1, 0xe7, 0xa4, 0x33, 0xe3, 0x69,
	0x0a, 0x38, 0xb4, 0xf2, 0x95, 0x23, 0xdc, 0x40, 0x19, 0xfe, 0xf3, 0xd4, 0x1f, 0x9d, 0xf6, 0xc8,
	0xc1, 0x5a, 0x2b, 0x26, 0xa9, 0x8c, 0x1d, 0xfa, 0xdf, 0xa3, 0xbb, 0xb2, 0x23, 0xf3, 0x44, 0x1a,
	0x23, 0xd5, 0x74, 0xb3, 0xba, 0x51, 0x91, 0x3b, 0x32, 0x3d, 0x25, 0xfb, 0x16, 0x9e, 0x0b, 0x25,
	0x36, 0x5c, 0xd3, 0x73, 0xbf, 0x45, 0x7a, 0x44, 0x9a, 0x3c, 0xd3, 0x85, 0xc2, 0x70, 0xcf, 0x8f,
	0xd7, 0x9d, 0xbb, 0xd9, 0xff, 0x25, 0x6c, 0x55, 0x37, 0xfb, 0xc6, 0x6d, 0x37, 0xbc, 0xcc, 0x40,
	0xe1, 0x40, 0x67, 0x26, 0x05, 0x84, 0xb0, 0xdd, 0x0d, 0x7a, 0xad, 0xd1, 0xae, 0x7c, 0x7d, 0xfb,
	0xb1, 0x64, 0xc1, 0x62, 0xc9, 0x82, 0xaf, 0x25, 0x0b, 0xde, 0x56, 0xac, 0xb6, 0x58, 0xb1, 0xda,
	0xe7, 0x8a, 0xd5, 0x9e, 0x2e, 0xa6, 0x12, 0x93, 0x62, 0xd2, 0x8f, 0x75, 0x16, 0xdd, 0x03, 0x1f,
	0xba, 0xbf, 0x0f, 0x7c, 0x1c, 0xd1, 0x56, 0x90, 0xf3, 0x75, 0x94, 0x58, 0x1a, 0xc8, 0x27, 0x4d,
	0x1f, 0xdf, 0xe5, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6b, 0xba, 0x22, 0xf8, 0xeb, 0x01, 0x00,
	0x00,
}

func (m *OrdersUnderWatch) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OrdersUnderWatch) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OrdersUnderWatch) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PaymentComplete {
		i--
		if m.PaymentComplete {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintOrdersUnderWatch(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintOrdersUnderWatch(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.RefundAddress) > 0 {
		i -= len(m.RefundAddress)
		copy(dAtA[i:], m.RefundAddress)
		i = encodeVarintOrdersUnderWatch(dAtA, i, uint64(len(m.RefundAddress)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.ShippingAddress) > 0 {
		i -= len(m.ShippingAddress)
		copy(dAtA[i:], m.ShippingAddress)
		i = encodeVarintOrdersUnderWatch(dAtA, i, uint64(len(m.ShippingAddress)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.WalletPublicKey) > 0 {
		i -= len(m.WalletPublicKey)
		copy(dAtA[i:], m.WalletPublicKey)
		i = encodeVarintOrdersUnderWatch(dAtA, i, uint64(len(m.WalletPublicKey)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.WalletPrivateKey) > 0 {
		i -= len(m.WalletPrivateKey)
		copy(dAtA[i:], m.WalletPrivateKey)
		i = encodeVarintOrdersUnderWatch(dAtA, i, uint64(len(m.WalletPrivateKey)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.NknAddress) > 0 {
		i -= len(m.NknAddress)
		copy(dAtA[i:], m.NknAddress)
		i = encodeVarintOrdersUnderWatch(dAtA, i, uint64(len(m.NknAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintOrdersUnderWatch(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOrdersUnderWatch(dAtA []byte, offset int, v uint64) int {
	offset -= sovOrdersUnderWatch(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OrdersUnderWatch) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovOrdersUnderWatch(uint64(l))
	}
	l = len(m.NknAddress)
	if l > 0 {
		n += 1 + l + sovOrdersUnderWatch(uint64(l))
	}
	l = len(m.WalletPrivateKey)
	if l > 0 {
		n += 1 + l + sovOrdersUnderWatch(uint64(l))
	}
	l = len(m.WalletPublicKey)
	if l > 0 {
		n += 1 + l + sovOrdersUnderWatch(uint64(l))
	}
	l = len(m.ShippingAddress)
	if l > 0 {
		n += 1 + l + sovOrdersUnderWatch(uint64(l))
	}
	l = len(m.RefundAddress)
	if l > 0 {
		n += 1 + l + sovOrdersUnderWatch(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovOrdersUnderWatch(uint64(l))
	}
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovOrdersUnderWatch(uint64(l))
	}
	if m.PaymentComplete {
		n += 2
	}
	return n
}

func sovOrdersUnderWatch(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOrdersUnderWatch(x uint64) (n int) {
	return sovOrdersUnderWatch(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OrdersUnderWatch) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrdersUnderWatch
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
			return fmt.Errorf("proto: OrdersUnderWatch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OrdersUnderWatch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrdersUnderWatch
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
				return ErrInvalidLengthOrdersUnderWatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrdersUnderWatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NknAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrdersUnderWatch
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
				return ErrInvalidLengthOrdersUnderWatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrdersUnderWatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NknAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WalletPrivateKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrdersUnderWatch
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
				return ErrInvalidLengthOrdersUnderWatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrdersUnderWatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WalletPrivateKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WalletPublicKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrdersUnderWatch
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
				return ErrInvalidLengthOrdersUnderWatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrdersUnderWatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WalletPublicKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShippingAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrdersUnderWatch
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
				return ErrInvalidLengthOrdersUnderWatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrdersUnderWatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ShippingAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RefundAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrdersUnderWatch
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
				return ErrInvalidLengthOrdersUnderWatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrdersUnderWatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RefundAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrdersUnderWatch
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
				return ErrInvalidLengthOrdersUnderWatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrdersUnderWatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrdersUnderWatch
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
				return ErrInvalidLengthOrdersUnderWatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrdersUnderWatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PaymentComplete", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrdersUnderWatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.PaymentComplete = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipOrdersUnderWatch(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOrdersUnderWatch
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
func skipOrdersUnderWatch(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOrdersUnderWatch
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
					return 0, ErrIntOverflowOrdersUnderWatch
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
					return 0, ErrIntOverflowOrdersUnderWatch
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
				return 0, ErrInvalidLengthOrdersUnderWatch
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOrdersUnderWatch
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOrdersUnderWatch
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOrdersUnderWatch        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOrdersUnderWatch          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOrdersUnderWatch = fmt.Errorf("proto: unexpected end of group")
)
