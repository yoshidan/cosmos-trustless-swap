// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: swap/swap/swap.proto

package types

import (
	fmt "fmt"
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

type SwapStatus int32

const (
	SwapStatus_Active    SwapStatus = 0
	SwapStatus_Closed    SwapStatus = 1
	SwapStatus_Cancelled SwapStatus = 2
)

var SwapStatus_name = map[int32]string{
	0: "Active",
	1: "Closed",
	2: "Cancelled",
}

var SwapStatus_value = map[string]int32{
	"Active":    0,
	"Closed":    1,
	"Cancelled": 2,
}

func (x SwapStatus) String() string {
	return proto.EnumName(SwapStatus_name, int32(x))
}

func (SwapStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f25efd28bf67e8da, []int{0}
}

// QueryParamsResponse is response type for the Query/Params RPC method.
type Swap struct {
	Id              uint64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Sender          string     `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Receiver        string     `protobuf:"bytes,3,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Amount          string     `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"`
	AmountToReceive string     `protobuf:"bytes,5,opt,name=amountToReceive,proto3" json:"amountToReceive,omitempty"`
	Status          SwapStatus `protobuf:"varint,6,opt,name=status,proto3,enum=swap.swap.SwapStatus" json:"status,omitempty"`
}

func (m *Swap) Reset()         { *m = Swap{} }
func (m *Swap) String() string { return proto.CompactTextString(m) }
func (*Swap) ProtoMessage()    {}
func (*Swap) Descriptor() ([]byte, []int) {
	return fileDescriptor_f25efd28bf67e8da, []int{0}
}
func (m *Swap) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Swap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Swap.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Swap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Swap.Merge(m, src)
}
func (m *Swap) XXX_Size() int {
	return m.Size()
}
func (m *Swap) XXX_DiscardUnknown() {
	xxx_messageInfo_Swap.DiscardUnknown(m)
}

var xxx_messageInfo_Swap proto.InternalMessageInfo

func (m *Swap) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Swap) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *Swap) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *Swap) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *Swap) GetAmountToReceive() string {
	if m != nil {
		return m.AmountToReceive
	}
	return ""
}

func (m *Swap) GetStatus() SwapStatus {
	if m != nil {
		return m.Status
	}
	return SwapStatus_Active
}

type SwapNFT struct {
	Id              uint64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Sender          string     `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Receiver        string     `protobuf:"bytes,3,opt,name=receiver,proto3" json:"receiver,omitempty"`
	ClassId         string     `protobuf:"bytes,4,opt,name=classId,proto3" json:"classId,omitempty"`
	NftId           string     `protobuf:"bytes,5,opt,name=nftId,proto3" json:"nftId,omitempty"`
	AmountToReceive string     `protobuf:"bytes,6,opt,name=amountToReceive,proto3" json:"amountToReceive,omitempty"`
	Status          SwapStatus `protobuf:"varint,7,opt,name=status,proto3,enum=swap.swap.SwapStatus" json:"status,omitempty"`
}

func (m *SwapNFT) Reset()         { *m = SwapNFT{} }
func (m *SwapNFT) String() string { return proto.CompactTextString(m) }
func (*SwapNFT) ProtoMessage()    {}
func (*SwapNFT) Descriptor() ([]byte, []int) {
	return fileDescriptor_f25efd28bf67e8da, []int{1}
}
func (m *SwapNFT) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SwapNFT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SwapNFT.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SwapNFT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SwapNFT.Merge(m, src)
}
func (m *SwapNFT) XXX_Size() int {
	return m.Size()
}
func (m *SwapNFT) XXX_DiscardUnknown() {
	xxx_messageInfo_SwapNFT.DiscardUnknown(m)
}

var xxx_messageInfo_SwapNFT proto.InternalMessageInfo

func (m *SwapNFT) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SwapNFT) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *SwapNFT) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *SwapNFT) GetClassId() string {
	if m != nil {
		return m.ClassId
	}
	return ""
}

func (m *SwapNFT) GetNftId() string {
	if m != nil {
		return m.NftId
	}
	return ""
}

func (m *SwapNFT) GetAmountToReceive() string {
	if m != nil {
		return m.AmountToReceive
	}
	return ""
}

func (m *SwapNFT) GetStatus() SwapStatus {
	if m != nil {
		return m.Status
	}
	return SwapStatus_Active
}

func init() {
	proto.RegisterEnum("swap.swap.SwapStatus", SwapStatus_name, SwapStatus_value)
	proto.RegisterType((*Swap)(nil), "swap.swap.Swap")
	proto.RegisterType((*SwapNFT)(nil), "swap.swap.SwapNFT")
}

func init() { proto.RegisterFile("swap/swap/swap.proto", fileDescriptor_f25efd28bf67e8da) }

var fileDescriptor_f25efd28bf67e8da = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x2e, 0x4f, 0x2c,
	0xd0, 0x87, 0x13, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x9c, 0x60, 0x36, 0x88, 0x90, 0x12,
	0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x8b, 0xea, 0x83, 0x58, 0x10, 0x05, 0x4a, 0xbb, 0x19, 0xb9, 0x58,
	0x82, 0xcb, 0x13, 0x0b, 0x84, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58,
	0x82, 0x98, 0x32, 0x53, 0x84, 0xc4, 0xb8, 0xd8, 0x8a, 0x53, 0xf3, 0x52, 0x52, 0x8b, 0x24, 0x98,
	0x14, 0x18, 0x35, 0x38, 0x83, 0xa0, 0x3c, 0x21, 0x29, 0x2e, 0x8e, 0xa2, 0xd4, 0xe4, 0xd4, 0xcc,
	0xb2, 0xd4, 0x22, 0x09, 0x66, 0xb0, 0x0c, 0x9c, 0x0f, 0xd2, 0x93, 0x98, 0x9b, 0x5f, 0x9a, 0x57,
	0x22, 0xc1, 0x02, 0xd1, 0x03, 0xe1, 0x09, 0x69, 0x70, 0xf1, 0x43, 0x58, 0x21, 0xf9, 0x41, 0x10,
	0xb5, 0x12, 0xac, 0x60, 0x05, 0xe8, 0xc2, 0x42, 0xba, 0x5c, 0x6c, 0xc5, 0x25, 0x89, 0x25, 0xa5,
	0xc5, 0x12, 0x6c, 0x0a, 0x8c, 0x1a, 0x7c, 0x46, 0xa2, 0x7a, 0x70, 0x0f, 0xe8, 0x81, 0x9c, 0x19,
	0x0c, 0x96, 0x0c, 0x82, 0x2a, 0x52, 0xba, 0xc6, 0xc8, 0xc5, 0x0e, 0x12, 0xf6, 0x73, 0x0b, 0xa1,
	0x8a, 0x07, 0x24, 0xb8, 0xd8, 0x93, 0x73, 0x12, 0x8b, 0x8b, 0x3d, 0x53, 0xa0, 0x3e, 0x80, 0x71,
	0x85, 0x44, 0xb8, 0x58, 0xf3, 0xd2, 0x4a, 0x3c, 0x53, 0xa0, 0x0e, 0x87, 0x70, 0xb0, 0x79, 0x8c,
	0x8d, 0x90, 0xc7, 0xd8, 0x89, 0xf0, 0x98, 0x96, 0x31, 0x17, 0x17, 0x42, 0x54, 0x88, 0x8b, 0x8b,
	0xcd, 0x31, 0xb9, 0x24, 0xb3, 0x2c, 0x55, 0x80, 0x01, 0xc4, 0x76, 0xce, 0xc9, 0x2f, 0x4e, 0x4d,
	0x11, 0x60, 0x14, 0xe2, 0xe5, 0xe2, 0x74, 0x4e, 0xcc, 0x4b, 0x4e, 0xcd, 0xc9, 0x49, 0x4d, 0x11,
	0x60, 0x72, 0xd2, 0x3e, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18,
	0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0x41, 0x70,
	0xba, 0xa8, 0x80, 0x24, 0x8f, 0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0x70, 0xfc, 0x1b, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xf4, 0xed, 0x57, 0x69, 0x38, 0x02, 0x00, 0x00,
}

func (m *Swap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Swap) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Swap) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintSwap(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x30
	}
	if len(m.AmountToReceive) > 0 {
		i -= len(m.AmountToReceive)
		copy(dAtA[i:], m.AmountToReceive)
		i = encodeVarintSwap(dAtA, i, uint64(len(m.AmountToReceive)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintSwap(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintSwap(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintSwap(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintSwap(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SwapNFT) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SwapNFT) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SwapNFT) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintSwap(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x38
	}
	if len(m.AmountToReceive) > 0 {
		i -= len(m.AmountToReceive)
		copy(dAtA[i:], m.AmountToReceive)
		i = encodeVarintSwap(dAtA, i, uint64(len(m.AmountToReceive)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.NftId) > 0 {
		i -= len(m.NftId)
		copy(dAtA[i:], m.NftId)
		i = encodeVarintSwap(dAtA, i, uint64(len(m.NftId)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ClassId) > 0 {
		i -= len(m.ClassId)
		copy(dAtA[i:], m.ClassId)
		i = encodeVarintSwap(dAtA, i, uint64(len(m.ClassId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintSwap(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintSwap(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintSwap(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSwap(dAtA []byte, offset int, v uint64) int {
	offset -= sovSwap(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Swap) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovSwap(uint64(m.Id))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovSwap(uint64(l))
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovSwap(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovSwap(uint64(l))
	}
	l = len(m.AmountToReceive)
	if l > 0 {
		n += 1 + l + sovSwap(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovSwap(uint64(m.Status))
	}
	return n
}

func (m *SwapNFT) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovSwap(uint64(m.Id))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovSwap(uint64(l))
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovSwap(uint64(l))
	}
	l = len(m.ClassId)
	if l > 0 {
		n += 1 + l + sovSwap(uint64(l))
	}
	l = len(m.NftId)
	if l > 0 {
		n += 1 + l + sovSwap(uint64(l))
	}
	l = len(m.AmountToReceive)
	if l > 0 {
		n += 1 + l + sovSwap(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovSwap(uint64(m.Status))
	}
	return n
}

func sovSwap(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSwap(x uint64) (n int) {
	return sovSwap(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Swap) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSwap
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
			return fmt.Errorf("proto: Swap: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Swap: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwap
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
				return ErrInvalidLengthSwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwap
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
				return ErrInvalidLengthSwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwap
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
				return ErrInvalidLengthSwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountToReceive", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwap
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
				return ErrInvalidLengthSwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AmountToReceive = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= SwapStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSwap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSwap
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
func (m *SwapNFT) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSwap
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
			return fmt.Errorf("proto: SwapNFT: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SwapNFT: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwap
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
				return ErrInvalidLengthSwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwap
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
				return ErrInvalidLengthSwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwap
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
				return ErrInvalidLengthSwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClassId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwap
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
				return ErrInvalidLengthSwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NftId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountToReceive", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwap
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
				return ErrInvalidLengthSwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AmountToReceive = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= SwapStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSwap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSwap
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
func skipSwap(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSwap
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
					return 0, ErrIntOverflowSwap
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
					return 0, ErrIntOverflowSwap
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
				return 0, ErrInvalidLengthSwap
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSwap
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSwap
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSwap        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSwap          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSwap = fmt.Errorf("proto: unexpected end of group")
)
