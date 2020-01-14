// Code generated by protoc-gen-go. DO NOT EDIT.
// source: goshare/position.proto

package goshare

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PositionSummationType int32

const (
	PositionSummationType_TOTAL     PositionSummationType = 0
	PositionSummationType_TODAY     PositionSummationType = 1
	PositionSummationType_YESTERDAY PositionSummationType = 2
)

var PositionSummationType_name = map[int32]string{
	0: "TOTAL",
	1: "TODAY",
	2: "YESTERDAY",
}

var PositionSummationType_value = map[string]int32{
	"TOTAL":     0,
	"TODAY":     1,
	"YESTERDAY": 2,
}

func (x PositionSummationType) String() string {
	return proto.EnumName(PositionSummationType_name, int32(x))
}

func (PositionSummationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0f94fc40d6c0c65e, []int{0}
}

// 持仓统计
type PositionSummation struct {
	// 方向
	Direction int32 `protobuf:"varint,1,opt,name=direction,proto3" json:"direction"`
	// 类型
	Type int32 `protobuf:"varint,2,opt,name=type,proto3" json:"type"`
	// 持仓
	Position int32 `protobuf:"varint,3,opt,name=position,proto3" json:"position"`
	// 成本
	Cost float64 `protobuf:"fixed64,4,opt,name=cost,proto3" json:"cost"`
	// 金额
	Amount float64 `protobuf:"fixed64,5,opt,name=amount,proto3" json:"amount"`
	// 冻结
	Frozen int32 `protobuf:"varint,6,opt,name=frozen,proto3" json:"frozen"`
	// 可用
	Available int32 `protobuf:"varint,7,opt,name=available,proto3" json:"available"`
	// 盈亏
	FloatProfit float64 `protobuf:"fixed64,9,opt,name=float_profit,json=floatProfit,proto3" json:"floatProfit"`
	// 保证金
	Margin               float64  `protobuf:"fixed64,11,opt,name=margin,proto3" json:"margin"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PositionSummation) Reset()         { *m = PositionSummation{} }
func (m *PositionSummation) String() string { return proto.CompactTextString(m) }
func (*PositionSummation) ProtoMessage()    {}
func (*PositionSummation) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f94fc40d6c0c65e, []int{0}
}

func (m *PositionSummation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PositionSummation.Unmarshal(m, b)
}
func (m *PositionSummation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PositionSummation.Marshal(b, m, deterministic)
}
func (m *PositionSummation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PositionSummation.Merge(m, src)
}
func (m *PositionSummation) XXX_Size() int {
	return xxx_messageInfo_PositionSummation.Size(m)
}
func (m *PositionSummation) XXX_DiscardUnknown() {
	xxx_messageInfo_PositionSummation.DiscardUnknown(m)
}

var xxx_messageInfo_PositionSummation proto.InternalMessageInfo

func (m *PositionSummation) GetDirection() int32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

func (m *PositionSummation) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *PositionSummation) GetPosition() int32 {
	if m != nil {
		return m.Position
	}
	return 0
}

func (m *PositionSummation) GetCost() float64 {
	if m != nil {
		return m.Cost
	}
	return 0
}

func (m *PositionSummation) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *PositionSummation) GetFrozen() int32 {
	if m != nil {
		return m.Frozen
	}
	return 0
}

func (m *PositionSummation) GetAvailable() int32 {
	if m != nil {
		return m.Available
	}
	return 0
}

func (m *PositionSummation) GetFloatProfit() float64 {
	if m != nil {
		return m.FloatProfit
	}
	return 0
}

func (m *PositionSummation) GetMargin() float64 {
	if m != nil {
		return m.Margin
	}
	return 0
}

// 持仓
type Position struct {
	UserId               string             `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"userId"`
	TaId                 string             `protobuf:"bytes,2,opt,name=ta_id,json=taId,proto3" json:"taId"`
	BuId                 string             `protobuf:"bytes,3,opt,name=bu_id,json=buId,proto3" json:"buId"`
	Exchange             string             `protobuf:"bytes,4,opt,name=exchange,proto3" json:"exchange"`
	Symbol               string             `protobuf:"bytes,5,opt,name=symbol,proto3" json:"symbol"`
	Product              string             `protobuf:"bytes,6,opt,name=product,proto3" json:"product"`
	Direction            int32              `protobuf:"varint,7,opt,name=direction,proto3" json:"direction"`
	Total                *PositionSummation `protobuf:"bytes,8,opt,name=total,proto3" json:"total"`
	Today                *PositionSummation `protobuf:"bytes,9,opt,name=today,proto3" json:"today"`
	Yesterday            *PositionSummation `protobuf:"bytes,10,opt,name=yesterday,proto3" json:"yesterday"`
	PreSettlementPrice   float64            `protobuf:"fixed64,11,opt,name=pre_settlement_price,json=preSettlementPrice,proto3" json:"preSettlementPrice"`
	SettlementPrice      float64            `protobuf:"fixed64,12,opt,name=settlement_price,json=settlementPrice,proto3" json:"settlementPrice"`
	LastPrice            float64            `protobuf:"fixed64,13,opt,name=last_price,json=lastPrice,proto3" json:"lastPrice"`
	TradingDay           int32              `protobuf:"varint,14,opt,name=trading_day,json=tradingDay,proto3" json:"tradingDay"`
	ProductType          int32              `protobuf:"varint,15,opt,name=product_type,json=productType,proto3" json:"productType"`
	Multiple             int32              `protobuf:"varint,16,opt,name=multiple,proto3" json:"multiple"`
	PriceTick            float64            `protobuf:"fixed64,17,opt,name=price_tick,json=priceTick,proto3" json:"priceTick"`
	SymbolName           string             `protobuf:"bytes,18,opt,name=symbol_name,json=symbolName,proto3" json:"symbolName"`
	UserName             string             `protobuf:"bytes,19,opt,name=user_name,json=userName,proto3" json:"userName"`
	Branch               string             `protobuf:"bytes,20,opt,name=branch,proto3" json:"branch"`
	BranchName           string             `protobuf:"bytes,21,opt,name=branch_name,json=branchName,proto3" json:"branchName"`
	IsCloseTodayAllowed  bool               `protobuf:"varint,22,opt,name=is_close_today_allowed,json=isCloseTodayAllowed,proto3" json:"isCloseTodayAllowed"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Position) Reset()         { *m = Position{} }
func (m *Position) String() string { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()    {}
func (*Position) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f94fc40d6c0c65e, []int{1}
}

func (m *Position) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Position.Unmarshal(m, b)
}
func (m *Position) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Position.Marshal(b, m, deterministic)
}
func (m *Position) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Position.Merge(m, src)
}
func (m *Position) XXX_Size() int {
	return xxx_messageInfo_Position.Size(m)
}
func (m *Position) XXX_DiscardUnknown() {
	xxx_messageInfo_Position.DiscardUnknown(m)
}

var xxx_messageInfo_Position proto.InternalMessageInfo

func (m *Position) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Position) GetTaId() string {
	if m != nil {
		return m.TaId
	}
	return ""
}

func (m *Position) GetBuId() string {
	if m != nil {
		return m.BuId
	}
	return ""
}

func (m *Position) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *Position) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *Position) GetProduct() string {
	if m != nil {
		return m.Product
	}
	return ""
}

func (m *Position) GetDirection() int32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

func (m *Position) GetTotal() *PositionSummation {
	if m != nil {
		return m.Total
	}
	return nil
}

func (m *Position) GetToday() *PositionSummation {
	if m != nil {
		return m.Today
	}
	return nil
}

func (m *Position) GetYesterday() *PositionSummation {
	if m != nil {
		return m.Yesterday
	}
	return nil
}

func (m *Position) GetPreSettlementPrice() float64 {
	if m != nil {
		return m.PreSettlementPrice
	}
	return 0
}

func (m *Position) GetSettlementPrice() float64 {
	if m != nil {
		return m.SettlementPrice
	}
	return 0
}

func (m *Position) GetLastPrice() float64 {
	if m != nil {
		return m.LastPrice
	}
	return 0
}

func (m *Position) GetTradingDay() int32 {
	if m != nil {
		return m.TradingDay
	}
	return 0
}

func (m *Position) GetProductType() int32 {
	if m != nil {
		return m.ProductType
	}
	return 0
}

func (m *Position) GetMultiple() int32 {
	if m != nil {
		return m.Multiple
	}
	return 0
}

func (m *Position) GetPriceTick() float64 {
	if m != nil {
		return m.PriceTick
	}
	return 0
}

func (m *Position) GetSymbolName() string {
	if m != nil {
		return m.SymbolName
	}
	return ""
}

func (m *Position) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *Position) GetBranch() string {
	if m != nil {
		return m.Branch
	}
	return ""
}

func (m *Position) GetBranchName() string {
	if m != nil {
		return m.BranchName
	}
	return ""
}

func (m *Position) GetIsCloseTodayAllowed() bool {
	if m != nil {
		return m.IsCloseTodayAllowed
	}
	return false
}

func init() {
	proto.RegisterEnum("goshare.PositionSummationType", PositionSummationType_name, PositionSummationType_value)
	proto.RegisterType((*PositionSummation)(nil), "goshare.PositionSummation")
	proto.RegisterType((*Position)(nil), "goshare.Position")
}

func init() { proto.RegisterFile("goshare/position.proto", fileDescriptor_0f94fc40d6c0c65e) }

var fileDescriptor_0f94fc40d6c0c65e = []byte{
	// 615 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0x4f, 0x4f, 0x1b, 0x3d,
	0x10, 0xc6, 0xdf, 0x05, 0xf2, 0x67, 0x27, 0xf0, 0x12, 0x0c, 0xa4, 0x16, 0x6d, 0xd5, 0x94, 0x53,
	0xda, 0xaa, 0x09, 0x2a, 0x97, 0x1e, 0x7a, 0xa1, 0x85, 0x03, 0x52, 0xd5, 0xa2, 0x25, 0x17, 0x7a,
	0x59, 0x79, 0x77, 0x4d, 0x62, 0xe1, 0x5d, 0xaf, 0x6c, 0x6f, 0xdb, 0xf4, 0xc3, 0xa3, 0xca, 0x63,
	0x27, 0xa8, 0xe4, 0xc0, 0xcd, 0xcf, 0xef, 0x99, 0xf1, 0xac, 0x67, 0x46, 0x0b, 0x83, 0x99, 0x32,
	0x73, 0xa6, 0xf9, 0xa4, 0x56, 0x46, 0x58, 0xa1, 0xaa, 0x71, 0xad, 0x95, 0x55, 0xa4, 0x13, 0xf8,
	0xf1, 0x7d, 0x04, 0x7b, 0x57, 0xc1, 0xbb, 0x6e, 0xca, 0x92, 0xb9, 0x03, 0x79, 0x01, 0x71, 0x21,
	0x34, 0xcf, 0x9d, 0xa0, 0xd1, 0x30, 0x1a, 0xb5, 0x92, 0x07, 0x40, 0x08, 0x6c, 0xd9, 0x45, 0xcd,
	0xe9, 0x06, 0x1a, 0x78, 0x26, 0x47, 0xd0, 0x5d, 0x96, 0xa0, 0x9b, 0xc8, 0x57, 0xda, 0xc5, 0xe7,
	0xca, 0x58, 0xba, 0x35, 0x8c, 0x46, 0x51, 0x82, 0x67, 0x32, 0x80, 0x36, 0x2b, 0x55, 0x53, 0x59,
	0xda, 0x42, 0x1a, 0x94, 0xe3, 0xb7, 0x5a, 0xfd, 0xe1, 0x15, 0x6d, 0xe3, 0x2d, 0x41, 0xb9, 0x2f,
	0x62, 0x3f, 0x99, 0x90, 0x2c, 0x93, 0x9c, 0x76, 0xfc, 0x17, 0xad, 0x00, 0x79, 0x0d, 0xdb, 0xb7,
	0x52, 0x31, 0x9b, 0xd6, 0x5a, 0xdd, 0x0a, 0x4b, 0x63, 0xbc, 0xb3, 0x87, 0xec, 0x0a, 0x91, 0xbb,
	0xb8, 0x64, 0x7a, 0x26, 0x2a, 0xda, 0xf3, 0x05, 0xbd, 0x3a, 0xbe, 0x6f, 0x41, 0x77, 0xd9, 0x00,
	0xf2, 0x0c, 0x3a, 0x8d, 0xe1, 0x3a, 0x15, 0x05, 0xbe, 0x3a, 0x4e, 0xda, 0x4e, 0x5e, 0x16, 0x64,
	0x1f, 0x5a, 0x96, 0x39, 0xbc, 0x81, 0x78, 0xcb, 0x32, 0x0f, 0xb3, 0xc6, 0xc1, 0x4d, 0x0f, 0xb3,
	0xe6, 0xb2, 0x70, 0x8d, 0xe0, 0xbf, 0xf3, 0x39, 0xab, 0x66, 0x1c, 0x1f, 0x1c, 0x27, 0x2b, 0xed,
	0xbe, 0xc1, 0x2c, 0xca, 0x4c, 0x49, 0x7c, 0x74, 0x9c, 0x04, 0x45, 0x28, 0x74, 0x6a, 0xad, 0x8a,
	0x26, 0xb7, 0xf8, 0xea, 0x38, 0x59, 0xca, 0x7f, 0x07, 0xd1, 0x79, 0x3c, 0x88, 0x13, 0x68, 0x59,
	0x65, 0x99, 0xa4, 0xdd, 0x61, 0x34, 0xea, 0x7d, 0x38, 0x1a, 0x87, 0xa9, 0x8e, 0xd7, 0x26, 0x9a,
	0xf8, 0x40, 0x9f, 0x51, 0xb0, 0x05, 0x76, 0xe8, 0xc9, 0x8c, 0x82, 0x2d, 0xc8, 0x47, 0x88, 0x17,
	0xdc, 0x58, 0xae, 0x5d, 0x16, 0x3c, 0x99, 0xf5, 0x10, 0x4c, 0x4e, 0xe0, 0xa0, 0xd6, 0x3c, 0x35,
	0xdc, 0x5a, 0xc9, 0x4b, 0x5e, 0xb9, 0xe9, 0x88, 0x9c, 0x87, 0xfe, 0x93, 0x5a, 0xf3, 0xeb, 0x95,
	0x75, 0xe5, 0x1c, 0xf2, 0x06, 0xfa, 0x6b, 0xd1, 0xdb, 0x18, 0xbd, 0x6b, 0x1e, 0x85, 0xbe, 0x04,
	0x90, 0xcc, 0x2c, 0x83, 0x76, 0x30, 0x28, 0x76, 0xc4, 0xdb, 0xaf, 0xa0, 0x67, 0x35, 0x2b, 0x44,
	0x35, 0x4b, 0xdd, 0x77, 0xff, 0x8f, 0x9d, 0x83, 0x80, 0xce, 0xd9, 0xc2, 0x6d, 0x4c, 0xe8, 0x71,
	0x8a, 0xbb, 0xbc, 0x8b, 0x11, 0xbd, 0xc0, 0xa6, 0x61, 0xa5, 0xcb, 0x46, 0x5a, 0x51, 0x4b, 0x4e,
	0xfb, 0x7e, 0xa5, 0x97, 0xda, 0x95, 0xc7, 0xca, 0xa9, 0x15, 0xf9, 0x1d, 0xdd, 0xf3, 0xe5, 0x91,
	0x4c, 0x45, 0x7e, 0xe7, 0xca, 0xfb, 0xd1, 0xa6, 0x15, 0x2b, 0x39, 0x25, 0x38, 0x54, 0xf0, 0xe8,
	0x1b, 0x2b, 0x39, 0x79, 0x0e, 0x31, 0x2e, 0x1a, 0xda, 0xfb, 0x7e, 0x4d, 0x1c, 0x40, 0x73, 0x00,
	0xed, 0x4c, 0xb3, 0x2a, 0x9f, 0xd3, 0x03, 0xbf, 0x26, 0x5e, 0xb9, 0x5b, 0xfd, 0xc9, 0xa7, 0x1d,
	0xfa, 0x5b, 0x3d, 0xc2, 0xc4, 0x53, 0x18, 0x08, 0x93, 0xe6, 0x52, 0x19, 0x9e, 0xe2, 0xf4, 0x52,
	0x26, 0xa5, 0xfa, 0xc5, 0x0b, 0x3a, 0x18, 0x46, 0xa3, 0x6e, 0xb2, 0x2f, 0xcc, 0x17, 0x67, 0x4e,
	0x9d, 0x77, 0xe6, 0xad, 0xb7, 0x9f, 0xe0, 0x70, 0x6d, 0x8c, 0xf8, 0xfe, 0x18, 0x5a, 0xd3, 0xef,
	0xd3, 0xb3, 0xaf, 0xfd, 0xff, 0xfc, 0xf1, 0xfc, 0xec, 0xa6, 0x1f, 0x91, 0x1d, 0x88, 0x6f, 0x2e,
	0xae, 0xa7, 0x17, 0x89, 0x93, 0x1b, 0x9f, 0xdf, 0xff, 0x78, 0x37, 0x13, 0x76, 0xde, 0x64, 0xe3,
	0x5c, 0x95, 0x93, 0x52, 0x54, 0x5c, 0x33, 0xa9, 0xb9, 0x99, 0xe0, 0x7f, 0xc6, 0x4c, 0x8c, 0xce,
	0x27, 0x33, 0x35, 0x09, 0x0b, 0x93, 0xb5, 0x11, 0x9f, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x72,
	0xa6, 0x02, 0x9f, 0x98, 0x04, 0x00, 0x00,
}