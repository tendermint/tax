// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tendermint/budget/v1beta1/budget.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Params defines the parameters for the budget module.
type Params struct {
	// The universal epoch length in number of blocks
	// A collection of budgets is executed with this epoch_blocks parameter
	EpochBlocks uint32 `protobuf:"varint,1,opt,name=epoch_blocks,json=epochBlocks,proto3" json:"epoch_blocks,omitempty" yaml:"epoch_blocks"`
	// Budgets parameter can be added, modified, and deleted through
	// parameter change governance proposal
	Budgets []Budget `protobuf:"bytes,2,rep,name=budgets,proto3" json:"budgets" yaml:"budgets"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_9df5cca239dd8691, []int{0}
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

func (m *Params) GetEpochBlocks() uint32 {
	if m != nil {
		return m.EpochBlocks
	}
	return 0
}

func (m *Params) GetBudgets() []Budget {
	if m != nil {
		return m.Budgets
	}
	return nil
}

// Budget defines a budget object.
type Budget struct {
	// name defines the name of the budget
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" yaml:"name"`
	// rate specifies the distributing amount by ratio of total budget source
	Rate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=rate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"rate" yaml:"rate"`
	// budget_source_address defines the bech32-encoded address that source of the budget
	// There does not require any interaction from the BudgetSourceAddress to accept that money will be withdrawn.
	// If this address is for e.g. is dormant, participants of the blockchain can vote to transfer his money and be
	// used by the community. Is this intentional?
	BudgetSourceAddress string `protobuf:"bytes,3,opt,name=budget_source_address,json=budgetSourceAddress,proto3" json:"budget_source_address,omitempty" yaml:"budget_source_address"`
	// collection_address defines the bech32-encoded address of the budget pool to distribute
	CollectionAddress string `protobuf:"bytes,4,opt,name=collection_address,json=collectionAddress,proto3" json:"collection_address,omitempty" yaml:"collection_address"`
	// start_time specifies the start time of the budget
	StartTime time.Time `protobuf:"bytes,5,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time" yaml:"start_time"`
	// end_time specifies the end time of the budget
	EndTime time.Time `protobuf:"bytes,6,opt,name=end_time,json=endTime,proto3,stdtime" json:"end_time" yaml:"end_time"`
}

func (m *Budget) Reset()      { *m = Budget{} }
func (*Budget) ProtoMessage() {}
func (*Budget) Descriptor() ([]byte, []int) {
	return fileDescriptor_9df5cca239dd8691, []int{1}
}
func (m *Budget) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Budget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Budget.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Budget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Budget.Merge(m, src)
}
func (m *Budget) XXX_Size() int {
	return m.Size()
}
func (m *Budget) XXX_DiscardUnknown() {
	xxx_messageInfo_Budget.DiscardUnknown(m)
}

var xxx_messageInfo_Budget proto.InternalMessageInfo

// TotalCollectedCoins defines total collected coins with relevant metadata.
type TotalCollectedCoins struct {
	// total_collected_coins specifis the total collected coins in a budget ever since the budget is created
	TotalCollectedCoins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=total_collected_coins,json=totalCollectedCoins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"total_collected_coins" yaml:"total_collected_coins"`
}

func (m *TotalCollectedCoins) Reset()         { *m = TotalCollectedCoins{} }
func (m *TotalCollectedCoins) String() string { return proto.CompactTextString(m) }
func (*TotalCollectedCoins) ProtoMessage()    {}
func (*TotalCollectedCoins) Descriptor() ([]byte, []int) {
	return fileDescriptor_9df5cca239dd8691, []int{2}
}
func (m *TotalCollectedCoins) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TotalCollectedCoins) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TotalCollectedCoins.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TotalCollectedCoins) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TotalCollectedCoins.Merge(m, src)
}
func (m *TotalCollectedCoins) XXX_Size() int {
	return m.Size()
}
func (m *TotalCollectedCoins) XXX_DiscardUnknown() {
	xxx_messageInfo_TotalCollectedCoins.DiscardUnknown(m)
}

var xxx_messageInfo_TotalCollectedCoins proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "cosmos.budget.v1beta1.Params")
	proto.RegisterType((*Budget)(nil), "cosmos.budget.v1beta1.Budget")
	proto.RegisterType((*TotalCollectedCoins)(nil), "cosmos.budget.v1beta1.TotalCollectedCoins")
}

func init() {
	proto.RegisterFile("tendermint/budget/v1beta1/budget.proto", fileDescriptor_9df5cca239dd8691)
}

var fileDescriptor_9df5cca239dd8691 = []byte{
	// 590 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xc7, 0xed, 0x36, 0xbf, 0xb4, 0xdd, 0xfc, 0xa0, 0xaa, 0x43, 0x21, 0x09, 0xc4, 0x1b, 0x19,
	0xa9, 0x8a, 0x84, 0xb0, 0xd5, 0x72, 0x8b, 0xc4, 0x01, 0x17, 0x6e, 0x48, 0x14, 0x93, 0x03, 0xe2,
	0x62, 0xad, 0xed, 0xc5, 0xb5, 0x6a, 0x7b, 0x23, 0xef, 0x06, 0xd1, 0x37, 0xe0, 0xd8, 0x23, 0x12,
	0x07, 0x7a, 0xe6, 0x1d, 0xb8, 0xf7, 0xd8, 0x23, 0x42, 0xc8, 0x45, 0xc9, 0x1b, 0xe4, 0x09, 0xd0,
	0xfe, 0x71, 0x12, 0x91, 0x4a, 0xf4, 0x14, 0xcf, 0xcc, 0x77, 0x3e, 0x33, 0xbb, 0xfb, 0x55, 0xc0,
	0x1e, 0xc3, 0x79, 0x84, 0x8b, 0x2c, 0xc9, 0x99, 0x13, 0x8c, 0xa3, 0x18, 0x33, 0xe7, 0xc3, 0x7e,
	0x80, 0x19, 0xda, 0x57, 0xa1, 0x3d, 0x2a, 0x08, 0x23, 0xc6, 0x6e, 0x48, 0x68, 0x46, 0xa8, 0xad,
	0x92, 0x4a, 0xd3, 0xb9, 0x13, 0x93, 0x98, 0x08, 0x85, 0xc3, 0xbf, 0xa4, 0xb8, 0xd3, 0x96, 0x62,
	0x5f, 0x16, 0x54, 0xa7, 0x2c, 0x99, 0x32, 0x72, 0x02, 0x44, 0xf1, 0x7c, 0x52, 0x48, 0x92, 0x5c,
	0xd5, 0x61, 0x4c, 0x48, 0x9c, 0x62, 0x47, 0x44, 0xc1, 0xf8, 0xbd, 0xc3, 0x92, 0x0c, 0x53, 0x86,
	0xb2, 0xd1, 0xcd, 0x00, 0xd6, 0x17, 0x1d, 0xd4, 0x8f, 0x50, 0x81, 0x32, 0x6a, 0x0c, 0xc0, 0xff,
	0x78, 0x44, 0xc2, 0x63, 0x3f, 0x48, 0x49, 0x78, 0x42, 0x5b, 0x7a, 0x4f, 0xef, 0xdf, 0x72, 0xef,
	0xcd, 0x4a, 0xd8, 0x3c, 0x45, 0x59, 0x3a, 0xb0, 0x96, 0xab, 0x96, 0xd7, 0x10, 0xa1, 0x2b, 0x22,
	0xe3, 0x15, 0xd8, 0x90, 0x47, 0xa5, 0xad, 0xb5, 0xde, 0x7a, 0xbf, 0x71, 0xd0, 0xb5, 0xaf, 0xbd,
	0x01, 0xdb, 0x15, 0xa1, 0x7b, 0xf7, 0xa2, 0x84, 0xda, 0xac, 0x84, 0xb7, 0x25, 0x59, 0xf5, 0x5a,
	0x5e, 0x45, 0x19, 0xd4, 0x3e, 0x9f, 0x43, 0xcd, 0xfa, 0xb5, 0x0e, 0xea, 0xb2, 0xc3, 0x78, 0x08,
	0x6a, 0x39, 0xca, 0xb0, 0xd8, 0x6a, 0xcb, 0xdd, 0x9e, 0x95, 0xb0, 0x21, 0x7b, 0x79, 0xd6, 0xf2,
	0x44, 0xd1, 0x78, 0x0d, 0x6a, 0x05, 0x62, 0xb8, 0xb5, 0x26, 0x44, 0x4f, 0xf9, 0x90, 0x9f, 0x25,
	0xdc, 0x8b, 0x13, 0x76, 0x3c, 0x0e, 0xec, 0x90, 0x64, 0xea, 0x76, 0xd5, 0xcf, 0x63, 0x1a, 0x9d,
	0x38, 0xec, 0x74, 0x84, 0xa9, 0xfd, 0x1c, 0x87, 0x0b, 0x24, 0x67, 0x58, 0x9e, 0x40, 0x19, 0x43,
	0xb0, 0x2b, 0x77, 0xf2, 0x29, 0x19, 0x17, 0x21, 0xf6, 0x51, 0x14, 0x15, 0x98, 0xd2, 0xd6, 0xba,
	0x98, 0xd1, 0x9b, 0x95, 0xf0, 0xc1, 0xf2, 0x21, 0xfe, 0x92, 0x59, 0x5e, 0x53, 0xe6, 0xdf, 0x88,
	0xf4, 0x33, 0x99, 0x35, 0x5e, 0x02, 0x23, 0x24, 0x69, 0x8a, 0x43, 0x96, 0x90, 0x7c, 0x8e, 0xac,
	0x09, 0x64, 0x77, 0x56, 0xc2, 0xb6, 0x44, 0xae, 0x6a, 0x2c, 0x6f, 0x67, 0x91, 0xac, 0x68, 0x6f,
	0x01, 0xa0, 0x0c, 0x15, 0xcc, 0xe7, 0xaf, 0xdf, 0xfa, 0xaf, 0xa7, 0xf7, 0x1b, 0x07, 0x1d, 0x5b,
	0x5a, 0xc3, 0xae, 0xac, 0x61, 0x0f, 0x2b, 0x6b, 0xb8, 0x5d, 0x75, 0xfb, 0x3b, 0x72, 0xca, 0xa2,
	0xd7, 0x3a, 0xbb, 0x82, 0xba, 0xb7, 0x25, 0x12, 0x5c, 0x6e, 0x78, 0x60, 0x13, 0xe7, 0x91, 0xe4,
	0xd6, 0xff, 0xc9, 0xbd, 0xaf, 0xb8, 0xdb, 0xca, 0x2f, 0xaa, 0x53, 0x52, 0x37, 0x70, 0x1e, 0x71,
	0xe9, 0x60, 0xf3, 0xd3, 0x39, 0xd4, 0xc4, 0xf3, 0x7e, 0xd7, 0x41, 0x73, 0x48, 0x18, 0x4a, 0x0f,
	0xe5, 0x91, 0x70, 0x74, 0x48, 0x92, 0x9c, 0x1a, 0x5f, 0x75, 0xb0, 0xcb, 0x78, 0xde, 0x0f, 0xab,
	0x82, 0xcf, 0x3d, 0xcb, 0x3d, 0xc9, 0xcd, 0xd5, 0x9e, 0x9b, 0x0b, 0x51, 0x3c, 0xb7, 0x16, 0xef,
	0x75, 0x8f, 0xd4, 0x0a, 0xea, 0x4d, 0xae, 0xa5, 0x58, 0xdf, 0xae, 0x60, 0xff, 0x06, 0x9e, 0x10,
	0xcb, 0x78, 0x4d, 0xb6, 0xba, 0xe1, 0xa0, 0xc6, 0xcf, 0xe0, 0xbe, 0xb8, 0x98, 0x98, 0xfa, 0xe5,
	0xc4, 0xd4, 0x7f, 0x4f, 0x4c, 0xfd, 0x6c, 0x6a, 0x6a, 0x97, 0x53, 0x53, 0xfb, 0x31, 0x35, 0xb5,
	0x77, 0x8f, 0x96, 0xf0, 0xab, 0x7f, 0x19, 0x1f, 0xab, 0x0f, 0x31, 0x27, 0xa8, 0x8b, 0xab, 0x7c,
	0xf2, 0x27, 0x00, 0x00, 0xff, 0xff, 0x39, 0xe3, 0x2a, 0x1f, 0x5d, 0x04, 0x00, 0x00,
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
	if len(m.Budgets) > 0 {
		for iNdEx := len(m.Budgets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Budgets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBudget(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.EpochBlocks != 0 {
		i = encodeVarintBudget(dAtA, i, uint64(m.EpochBlocks))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Budget) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Budget) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Budget) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.EndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintBudget(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x32
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintBudget(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x2a
	if len(m.CollectionAddress) > 0 {
		i -= len(m.CollectionAddress)
		copy(dAtA[i:], m.CollectionAddress)
		i = encodeVarintBudget(dAtA, i, uint64(len(m.CollectionAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.BudgetSourceAddress) > 0 {
		i -= len(m.BudgetSourceAddress)
		copy(dAtA[i:], m.BudgetSourceAddress)
		i = encodeVarintBudget(dAtA, i, uint64(len(m.BudgetSourceAddress)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.Rate.Size()
		i -= size
		if _, err := m.Rate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBudget(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintBudget(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TotalCollectedCoins) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TotalCollectedCoins) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TotalCollectedCoins) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TotalCollectedCoins) > 0 {
		for iNdEx := len(m.TotalCollectedCoins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TotalCollectedCoins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBudget(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintBudget(dAtA []byte, offset int, v uint64) int {
	offset -= sovBudget(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EpochBlocks != 0 {
		n += 1 + sovBudget(uint64(m.EpochBlocks))
	}
	if len(m.Budgets) > 0 {
		for _, e := range m.Budgets {
			l = e.Size()
			n += 1 + l + sovBudget(uint64(l))
		}
	}
	return n
}

func (m *Budget) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovBudget(uint64(l))
	}
	l = m.Rate.Size()
	n += 1 + l + sovBudget(uint64(l))
	l = len(m.BudgetSourceAddress)
	if l > 0 {
		n += 1 + l + sovBudget(uint64(l))
	}
	l = len(m.CollectionAddress)
	if l > 0 {
		n += 1 + l + sovBudget(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovBudget(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime)
	n += 1 + l + sovBudget(uint64(l))
	return n
}

func (m *TotalCollectedCoins) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.TotalCollectedCoins) > 0 {
		for _, e := range m.TotalCollectedCoins {
			l = e.Size()
			n += 1 + l + sovBudget(uint64(l))
		}
	}
	return n
}

func sovBudget(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBudget(x uint64) (n int) {
	return sovBudget(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBudget
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochBlocks", wireType)
			}
			m.EpochBlocks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBudget
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochBlocks |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Budgets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBudget
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
				return ErrInvalidLengthBudget
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBudget
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Budgets = append(m.Budgets, Budget{})
			if err := m.Budgets[len(m.Budgets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBudget(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBudget
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
func (m *Budget) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBudget
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
			return fmt.Errorf("proto: Budget: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Budget: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBudget
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
				return ErrInvalidLengthBudget
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBudget
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBudget
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
				return ErrInvalidLengthBudget
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBudget
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Rate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BudgetSourceAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBudget
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
				return ErrInvalidLengthBudget
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBudget
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BudgetSourceAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollectionAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBudget
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
				return ErrInvalidLengthBudget
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBudget
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollectionAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBudget
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
				return ErrInvalidLengthBudget
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBudget
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBudget
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
				return ErrInvalidLengthBudget
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBudget
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.EndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBudget(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBudget
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
func (m *TotalCollectedCoins) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBudget
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
			return fmt.Errorf("proto: TotalCollectedCoins: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TotalCollectedCoins: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalCollectedCoins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBudget
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
				return ErrInvalidLengthBudget
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBudget
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TotalCollectedCoins = append(m.TotalCollectedCoins, types.Coin{})
			if err := m.TotalCollectedCoins[len(m.TotalCollectedCoins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBudget(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBudget
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
func skipBudget(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBudget
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
					return 0, ErrIntOverflowBudget
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
					return 0, ErrIntOverflowBudget
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
				return 0, ErrInvalidLengthBudget
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBudget
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBudget
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBudget        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBudget          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBudget = fmt.Errorf("proto: unexpected end of group")
)
