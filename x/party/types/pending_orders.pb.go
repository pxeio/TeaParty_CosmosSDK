// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: partychain/party/pending_orders.proto

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

type PendingOrders struct {
	Index                        string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	BuyerEscrowWalletPublicKey   string `protobuf:"bytes,2,opt,name=buyerEscrowWalletPublicKey,proto3" json:"buyerEscrowWalletPublicKey,omitempty"`
	BuyerEscrowWalletPrivateKey  string `protobuf:"bytes,3,opt,name=buyerEscrowWalletPrivateKey,proto3" json:"buyerEscrowWalletPrivateKey,omitempty"`
	SellerEscrowWalletPublicKey  string `protobuf:"bytes,4,opt,name=sellerEscrowWalletPublicKey,proto3" json:"sellerEscrowWalletPublicKey,omitempty"`
	SellerEscrowWalletPrivateKey string `protobuf:"bytes,5,opt,name=sellerEscrowWalletPrivateKey,proto3" json:"sellerEscrowWalletPrivateKey,omitempty"`
	SellerPaymentComplete        bool   `protobuf:"varint,6,opt,name=sellerPaymentComplete,proto3" json:"sellerPaymentComplete,omitempty"`
	BuyerPaymentComplete         bool   `protobuf:"varint,7,opt,name=buyerPaymentComplete,proto3" json:"buyerPaymentComplete,omitempty"`
	Amount                       string `protobuf:"bytes,8,opt,name=amount,proto3" json:"amount,omitempty"`
	BuyerShippingAddress         string `protobuf:"bytes,9,opt,name=buyerShippingAddress,proto3" json:"buyerShippingAddress,omitempty"`
	BuyerRefundAddress           string `protobuf:"bytes,10,opt,name=buyerRefundAddress,proto3" json:"buyerRefundAddress,omitempty"`
	BuyerNKNAddress              string `protobuf:"bytes,11,opt,name=buyerNKNAddress,proto3" json:"buyerNKNAddress,omitempty"`
	SellerRefundAddress          string `protobuf:"bytes,12,opt,name=sellerRefundAddress,proto3" json:"sellerRefundAddress,omitempty"`
	SellerShippingAddress        string `protobuf:"bytes,13,opt,name=sellerShippingAddress,proto3" json:"sellerShippingAddress,omitempty"`
	SellerNKNAddress             string `protobuf:"bytes,14,opt,name=sellerNKNAddress,proto3" json:"sellerNKNAddress,omitempty"`
	TradeAsset                   string `protobuf:"bytes,15,opt,name=tradeAsset,proto3" json:"tradeAsset,omitempty"`
	Currency                     string `protobuf:"bytes,16,opt,name=currency,proto3" json:"currency,omitempty"`
	Price                        string `protobuf:"bytes,17,opt,name=price,proto3" json:"price,omitempty"`
	BlockHeight                  string `protobuf:"bytes,18,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
}

func (m *PendingOrders) Reset()         { *m = PendingOrders{} }
func (m *PendingOrders) String() string { return proto.CompactTextString(m) }
func (*PendingOrders) ProtoMessage()    {}
func (*PendingOrders) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6bac0cc0b8c6155, []int{0}
}
func (m *PendingOrders) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PendingOrders) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PendingOrders.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PendingOrders) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PendingOrders.Merge(m, src)
}
func (m *PendingOrders) XXX_Size() int {
	return m.Size()
}
func (m *PendingOrders) XXX_DiscardUnknown() {
	xxx_messageInfo_PendingOrders.DiscardUnknown(m)
}

var xxx_messageInfo_PendingOrders proto.InternalMessageInfo

func (m *PendingOrders) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *PendingOrders) GetBuyerEscrowWalletPublicKey() string {
	if m != nil {
		return m.BuyerEscrowWalletPublicKey
	}
	return ""
}

func (m *PendingOrders) GetBuyerEscrowWalletPrivateKey() string {
	if m != nil {
		return m.BuyerEscrowWalletPrivateKey
	}
	return ""
}

func (m *PendingOrders) GetSellerEscrowWalletPublicKey() string {
	if m != nil {
		return m.SellerEscrowWalletPublicKey
	}
	return ""
}

func (m *PendingOrders) GetSellerEscrowWalletPrivateKey() string {
	if m != nil {
		return m.SellerEscrowWalletPrivateKey
	}
	return ""
}

func (m *PendingOrders) GetSellerPaymentComplete() bool {
	if m != nil {
		return m.SellerPaymentComplete
	}
	return false
}

func (m *PendingOrders) GetBuyerPaymentComplete() bool {
	if m != nil {
		return m.BuyerPaymentComplete
	}
	return false
}

func (m *PendingOrders) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *PendingOrders) GetBuyerShippingAddress() string {
	if m != nil {
		return m.BuyerShippingAddress
	}
	return ""
}

func (m *PendingOrders) GetBuyerRefundAddress() string {
	if m != nil {
		return m.BuyerRefundAddress
	}
	return ""
}

func (m *PendingOrders) GetBuyerNKNAddress() string {
	if m != nil {
		return m.BuyerNKNAddress
	}
	return ""
}

func (m *PendingOrders) GetSellerRefundAddress() string {
	if m != nil {
		return m.SellerRefundAddress
	}
	return ""
}

func (m *PendingOrders) GetSellerShippingAddress() string {
	if m != nil {
		return m.SellerShippingAddress
	}
	return ""
}

func (m *PendingOrders) GetSellerNKNAddress() string {
	if m != nil {
		return m.SellerNKNAddress
	}
	return ""
}

func (m *PendingOrders) GetTradeAsset() string {
	if m != nil {
		return m.TradeAsset
	}
	return ""
}

func (m *PendingOrders) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *PendingOrders) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *PendingOrders) GetBlockHeight() string {
	if m != nil {
		return m.BlockHeight
	}
	return ""
}

func init() {
	proto.RegisterType((*PendingOrders)(nil), "teapartycrypto.partychain.party.PendingOrders")
}

func init() {
	proto.RegisterFile("partychain/party/pending_orders.proto", fileDescriptor_a6bac0cc0b8c6155)
}

var fileDescriptor_a6bac0cc0b8c6155 = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x51, 0x6b, 0x13, 0x41,
	0x10, 0xc7, 0x73, 0x6a, 0x63, 0x3a, 0xb5, 0xb6, 0x8e, 0x55, 0x96, 0x2a, 0x67, 0x10, 0x84, 0xe0,
	0x43, 0x22, 0xd5, 0x67, 0xb1, 0x2d, 0x82, 0x50, 0xa9, 0x21, 0x0a, 0x82, 0x2f, 0xb2, 0x77, 0x37,
	0x26, 0x8b, 0x97, 0xdd, 0x65, 0x6f, 0x4f, 0x7b, 0xdf, 0xc2, 0x2f, 0x25, 0xf8, 0xd8, 0x47, 0x1f,
	0x25, 0xf9, 0x22, 0x92, 0xd9, 0xb6, 0x89, 0xc9, 0x99, 0xb7, 0x9b, 0xf9, 0xfd, 0xff, 0x7f, 0x66,
	0x8e, 0x1d, 0x78, 0x62, 0xa5, 0xf3, 0x55, 0x3a, 0x92, 0x4a, 0xf7, 0xf8, 0xb3, 0x67, 0x49, 0x67,
	0x4a, 0x0f, 0x3f, 0x1b, 0x97, 0x91, 0x2b, 0xba, 0xd6, 0x19, 0x6f, 0xf0, 0x91, 0x27, 0x19, 0x94,
	0xae, 0xb2, 0xde, 0x74, 0xe7, 0xae, 0xf0, 0xf9, 0xf8, 0x67, 0x13, 0xb6, 0xfb, 0xc1, 0xf9, 0x8e,
	0x8d, 0xb8, 0x07, 0x1b, 0x4a, 0x67, 0x74, 0x26, 0xa2, 0x76, 0xd4, 0xd9, 0x1c, 0x84, 0x02, 0x5f,
	0xc2, 0x7e, 0x52, 0x56, 0xe4, 0x5e, 0x17, 0xa9, 0x33, 0xdf, 0x3f, 0xca, 0x3c, 0x27, 0xdf, 0x2f,
	0x93, 0x5c, 0xa5, 0x27, 0x54, 0x89, 0x6b, 0x2c, 0x5d, 0xa3, 0xc0, 0x57, 0xf0, 0x60, 0x95, 0x3a,
	0xf5, 0x4d, 0x7a, 0x9a, 0x05, 0x5c, 0xe7, 0x80, 0x75, 0x92, 0x59, 0x42, 0x41, 0x79, 0xfe, 0xbf,
	0x11, 0x6e, 0x84, 0x84, 0x35, 0x12, 0x3c, 0x82, 0x87, 0x35, 0x78, 0x3e, 0xc4, 0x06, 0x47, 0xac,
	0xd5, 0xe0, 0x0b, 0xb8, 0x17, 0x78, 0x5f, 0x56, 0x63, 0xd2, 0xfe, 0xd8, 0x8c, 0x6d, 0x4e, 0x9e,
	0x44, 0xb3, 0x1d, 0x75, 0x5a, 0x83, 0x7a, 0x88, 0x07, 0xb0, 0xc7, 0xab, 0x2d, 0x9b, 0x6e, 0xb2,
	0xa9, 0x96, 0xe1, 0x7d, 0x68, 0xca, 0xb1, 0x29, 0xb5, 0x17, 0x2d, 0x9e, 0xeb, 0xa2, 0xba, 0xca,
	0x7a, 0x3f, 0x52, 0xd6, 0x2a, 0x3d, 0x3c, 0xcc, 0x32, 0x47, 0x45, 0x21, 0x36, 0x59, 0x55, 0xcb,
	0xb0, 0x0b, 0xc8, 0xfd, 0x01, 0x7d, 0x29, 0x75, 0x76, 0xe9, 0x00, 0x76, 0xd4, 0x10, 0xec, 0xc0,
	0x0e, 0x77, 0x4f, 0x4f, 0x4e, 0x2f, 0xc5, 0x5b, 0x2c, 0x5e, 0x6e, 0xe3, 0x33, 0xb8, 0x1b, 0x56,
	0xfe, 0x37, 0xfa, 0x16, 0xab, 0xeb, 0xd0, 0xfc, 0x0f, 0x2e, 0x2f, 0xb0, 0xcd, 0x9e, 0x7a, 0x88,
	0x4f, 0x61, 0x37, 0x80, 0x85, 0x91, 0x6e, 0xb3, 0x61, 0xa5, 0x8f, 0x31, 0x80, 0x77, 0x32, 0xa3,
	0xc3, 0xa2, 0x20, 0x2f, 0x76, 0x58, 0xb5, 0xd0, 0xc1, 0x7d, 0x68, 0xa5, 0xa5, 0x73, 0xa4, 0xd3,
	0x4a, 0xec, 0x32, 0xbd, 0xaa, 0x67, 0xaf, 0xdf, 0x3a, 0x95, 0x92, 0xb8, 0x13, 0x5e, 0x3f, 0x17,
	0xd8, 0x86, 0xad, 0x24, 0x37, 0xe9, 0xd7, 0x37, 0xa4, 0x86, 0x23, 0x2f, 0x90, 0xd9, 0x62, 0xeb,
	0xe8, 0xed, 0xaf, 0x49, 0x1c, 0x9d, 0x4f, 0xe2, 0xe8, 0xcf, 0x24, 0x8e, 0x7e, 0x4c, 0xe3, 0xc6,
	0xf9, 0x34, 0x6e, 0xfc, 0x9e, 0xc6, 0x8d, 0x4f, 0x07, 0x43, 0xe5, 0x47, 0x65, 0xd2, 0x4d, 0xcd,
	0xb8, 0xf7, 0x81, 0x64, 0x7f, 0x76, 0x76, 0xc7, 0x7c, 0x8d, 0xbd, 0x85, 0x1b, 0x3e, 0xbb, 0xb8,
	0x62, 0x5f, 0x59, 0x2a, 0x92, 0x26, 0x5f, 0xef, 0xf3, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xda,
	0x9f, 0xff, 0x52, 0xe6, 0x03, 0x00, 0x00,
}

func (m *PendingOrders) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PendingOrders) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PendingOrders) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BlockHeight) > 0 {
		i -= len(m.BlockHeight)
		copy(dAtA[i:], m.BlockHeight)
		i = encodeVarintPendingOrders(dAtA, i, uint64(len(m.BlockHeight)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x92
	}
	if len(m.Price) > 0 {
		i -= len(m.Price)
		copy(dAtA[i:], m.Price)
		i = encodeVarintPendingOrders(dAtA, i, uint64(len(m.Price)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x8a
	}
	if len(m.Currency) > 0 {
		i -= len(m.Currency)
		copy(dAtA[i:], m.Currency)
		i = encodeVarintPendingOrders(dAtA, i, uint64(len(m.Currency)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x82
	}
	if len(m.TradeAsset) > 0 {
		i -= len(m.TradeAsset)
		copy(dAtA[i:], m.TradeAsset)
		i = encodeVarintPendingOrders(dAtA, i, uint64(len(m.TradeAsset)))
		i--
		dAtA[i] = 0x7a
	}
	if len(m.SellerNKNAddress) > 0 {
		i -= len(m.SellerNKNAddress)
		copy(dAtA[i:], m.SellerNKNAddress)
		i = encodeVarintPendingOrders(dAtA, i, uint64(len(m.SellerNKNAddress)))
		i--
		dAtA[i] = 0x72
	}
	if len(m.SellerShippingAddress) > 0 {
		i -= len(m.SellerShippingAddress)
		copy(dAtA[i:], m.SellerShippingAddress)
		i = encodeVarintPendingOrders(dAtA, i, uint64(len(m.SellerShippingAddress)))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.SellerRefundAddress) > 0 {
		i -= len(m.SellerRefundAddress)
		copy(dAtA[i:], m.SellerRefundAddress)
		i = encodeVarintPendingOrders(dAtA, i, uint64(len(m.SellerRefundAddress)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.BuyerNKNAddress) > 0 {
		i -= len(m.BuyerNKNAddress)
		copy(dAtA[i:], m.BuyerNKNAddress)
		i = encodeVarintPendingOrders(dAtA, i, uint64(len(m.BuyerNKNAddress)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.BuyerRefundAddress) > 0 {
		i -= len(m.BuyerRefundAddress)
		copy(dAtA[i:], m.BuyerRefundAddress)
		i = encodeVarintPendingOrders(dAtA, i, uint64(len(m.BuyerRefundAddress)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.BuyerShippingAddress) > 0 {
		i -= len(m.BuyerShippingAddress)
		copy(dAtA[i:], m.BuyerShippingAddress)
		i = encodeVarintPendingOrders(dAtA, i, uint64(len(m.BuyerShippingAddress)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintPendingOrders(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x42
	}
	if m.BuyerPaymentComplete {
		i--
		if m.BuyerPaymentComplete {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.SellerPaymentComplete {
		i--
		if m.SellerPaymentComplete {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if len(m.SellerEscrowWalletPrivateKey) > 0 {
		i -= len(m.SellerEscrowWalletPrivateKey)
		copy(dAtA[i:], m.SellerEscrowWalletPrivateKey)
		i = encodeVarintPendingOrders(dAtA, i, uint64(len(m.SellerEscrowWalletPrivateKey)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.SellerEscrowWalletPublicKey) > 0 {
		i -= len(m.SellerEscrowWalletPublicKey)
		copy(dAtA[i:], m.SellerEscrowWalletPublicKey)
		i = encodeVarintPendingOrders(dAtA, i, uint64(len(m.SellerEscrowWalletPublicKey)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.BuyerEscrowWalletPrivateKey) > 0 {
		i -= len(m.BuyerEscrowWalletPrivateKey)
		copy(dAtA[i:], m.BuyerEscrowWalletPrivateKey)
		i = encodeVarintPendingOrders(dAtA, i, uint64(len(m.BuyerEscrowWalletPrivateKey)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.BuyerEscrowWalletPublicKey) > 0 {
		i -= len(m.BuyerEscrowWalletPublicKey)
		copy(dAtA[i:], m.BuyerEscrowWalletPublicKey)
		i = encodeVarintPendingOrders(dAtA, i, uint64(len(m.BuyerEscrowWalletPublicKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintPendingOrders(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPendingOrders(dAtA []byte, offset int, v uint64) int {
	offset -= sovPendingOrders(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PendingOrders) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovPendingOrders(uint64(l))
	}
	l = len(m.BuyerEscrowWalletPublicKey)
	if l > 0 {
		n += 1 + l + sovPendingOrders(uint64(l))
	}
	l = len(m.BuyerEscrowWalletPrivateKey)
	if l > 0 {
		n += 1 + l + sovPendingOrders(uint64(l))
	}
	l = len(m.SellerEscrowWalletPublicKey)
	if l > 0 {
		n += 1 + l + sovPendingOrders(uint64(l))
	}
	l = len(m.SellerEscrowWalletPrivateKey)
	if l > 0 {
		n += 1 + l + sovPendingOrders(uint64(l))
	}
	if m.SellerPaymentComplete {
		n += 2
	}
	if m.BuyerPaymentComplete {
		n += 2
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovPendingOrders(uint64(l))
	}
	l = len(m.BuyerShippingAddress)
	if l > 0 {
		n += 1 + l + sovPendingOrders(uint64(l))
	}
	l = len(m.BuyerRefundAddress)
	if l > 0 {
		n += 1 + l + sovPendingOrders(uint64(l))
	}
	l = len(m.BuyerNKNAddress)
	if l > 0 {
		n += 1 + l + sovPendingOrders(uint64(l))
	}
	l = len(m.SellerRefundAddress)
	if l > 0 {
		n += 1 + l + sovPendingOrders(uint64(l))
	}
	l = len(m.SellerShippingAddress)
	if l > 0 {
		n += 1 + l + sovPendingOrders(uint64(l))
	}
	l = len(m.SellerNKNAddress)
	if l > 0 {
		n += 1 + l + sovPendingOrders(uint64(l))
	}
	l = len(m.TradeAsset)
	if l > 0 {
		n += 1 + l + sovPendingOrders(uint64(l))
	}
	l = len(m.Currency)
	if l > 0 {
		n += 2 + l + sovPendingOrders(uint64(l))
	}
	l = len(m.Price)
	if l > 0 {
		n += 2 + l + sovPendingOrders(uint64(l))
	}
	l = len(m.BlockHeight)
	if l > 0 {
		n += 2 + l + sovPendingOrders(uint64(l))
	}
	return n
}

func sovPendingOrders(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPendingOrders(x uint64) (n int) {
	return sovPendingOrders(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PendingOrders) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPendingOrders
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
			return fmt.Errorf("proto: PendingOrders: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PendingOrders: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPendingOrders
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
				return ErrInvalidLengthPendingOrders
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPendingOrders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BuyerEscrowWalletPublicKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPendingOrders
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
				return ErrInvalidLengthPendingOrders
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPendingOrders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BuyerEscrowWalletPublicKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BuyerEscrowWalletPrivateKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPendingOrders
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
				return ErrInvalidLengthPendingOrders
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPendingOrders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BuyerEscrowWalletPrivateKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SellerEscrowWalletPublicKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPendingOrders
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
				return ErrInvalidLengthPendingOrders
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPendingOrders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SellerEscrowWalletPublicKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SellerEscrowWalletPrivateKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPendingOrders
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
				return ErrInvalidLengthPendingOrders
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPendingOrders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SellerEscrowWalletPrivateKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SellerPaymentComplete", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPendingOrders
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
			m.SellerPaymentComplete = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BuyerPaymentComplete", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPendingOrders
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
			m.BuyerPaymentComplete = bool(v != 0)
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPendingOrders
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
				return ErrInvalidLengthPendingOrders
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPendingOrders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BuyerShippingAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPendingOrders
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
				return ErrInvalidLengthPendingOrders
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPendingOrders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BuyerShippingAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BuyerRefundAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPendingOrders
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
				return ErrInvalidLengthPendingOrders
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPendingOrders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BuyerRefundAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BuyerNKNAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPendingOrders
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
				return ErrInvalidLengthPendingOrders
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPendingOrders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BuyerNKNAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SellerRefundAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPendingOrders
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
				return ErrInvalidLengthPendingOrders
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPendingOrders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SellerRefundAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SellerShippingAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPendingOrders
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
				return ErrInvalidLengthPendingOrders
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPendingOrders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SellerShippingAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SellerNKNAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPendingOrders
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
				return ErrInvalidLengthPendingOrders
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPendingOrders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SellerNKNAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradeAsset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPendingOrders
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
				return ErrInvalidLengthPendingOrders
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPendingOrders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TradeAsset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Currency", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPendingOrders
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
				return ErrInvalidLengthPendingOrders
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPendingOrders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Currency = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 17:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPendingOrders
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
				return ErrInvalidLengthPendingOrders
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPendingOrders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Price = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 18:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPendingOrders
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
				return ErrInvalidLengthPendingOrders
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPendingOrders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockHeight = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPendingOrders(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPendingOrders
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
func skipPendingOrders(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPendingOrders
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
					return 0, ErrIntOverflowPendingOrders
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
					return 0, ErrIntOverflowPendingOrders
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
				return 0, ErrInvalidLengthPendingOrders
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPendingOrders
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPendingOrders
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPendingOrders        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPendingOrders          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPendingOrders = fmt.Errorf("proto: unexpected end of group")
)
