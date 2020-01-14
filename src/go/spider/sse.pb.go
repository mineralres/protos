// Code generated by protoc-gen-go. DO NOT EDIT.
// source: spider/sse.proto

package goshare_spider

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

// SSEStockOption 上证交易所ETF期权
type SSEStockOption struct {
	// 行权价
	ExercisePrice        string   `protobuf:"bytes,1,opt,name=exercise_price,json=exercisePrice,proto3" json:"exercisePrice"`
	UpdateVersion        string   `protobuf:"bytes,2,opt,name=update_version,json=updateVersion,proto3" json:"updateVersion"`
	OptionType           string   `protobuf:"bytes,3,opt,name=option_type,json=optionType,proto3" json:"optionType"`
	DailyPriceUpLimit    string   `protobuf:"bytes,4,opt,name=daily_price_up_limit,json=dailyPriceUpLimit,proto3" json:"dailyPriceUpLimit"`
	TimeSave             string   `protobuf:"bytes,5,opt,name=time_save,json=timeSave,proto3" json:"timeSave"`
	DELIST_Flag          string   `protobuf:"bytes,6,opt,name=DELIST_Flag,json=DELISTFlag,proto3" json:"DELISTFlag"`
	StartDate            string   `protobuf:"bytes,7,opt,name=start_date,json=startDate,proto3" json:"startDate"`
	ExpireDate           string   `protobuf:"bytes,8,opt,name=expire_date,json=expireDate,proto3" json:"expireDate"`
	ContractUnit         string   `protobuf:"bytes,9,opt,name=contract_unit,json=contractUnit,proto3" json:"contractUnit"`
	CallOrPut            string   `protobuf:"bytes,10,opt,name=call_or_put,json=callOrPut,proto3" json:"callOrPut"`
	LmtOrdMaxFloor       string   `protobuf:"bytes,11,opt,name=lmt_ord_max_floor,json=lmtOrdMaxFloor,proto3" json:"lmtOrdMaxFloor"`
	DeliveryDate         string   `protobuf:"bytes,12,opt,name=delivery_date,json=deliveryDate,proto3" json:"deliveryDate"`
	ChangeFlag           string   `protobuf:"bytes,13,opt,name=change_flag,json=changeFlag,proto3" json:"changeFlag"`
	MktOrdMaxFloor       string   `protobuf:"bytes,14,opt,name=mkt_ord_max_floor,json=mktOrdMaxFloor,proto3" json:"mktOrdMaxFloor"`
	UnderlyingType       string   `protobuf:"bytes,15,opt,name=underlying_type,json=underlyingType,proto3" json:"underlyingType"`
	DailyPriceDownLimit  string   `protobuf:"bytes,16,opt,name=daily_price_down_limit,json=dailyPriceDownLimit,proto3" json:"dailyPriceDownLimit"`
	RoundLot             string   `protobuf:"bytes,17,opt,name=round_lot,json=roundLot,proto3" json:"roundLot"`
	SecurityClosePX      string   `protobuf:"bytes,18,opt,name=security_closePX,json=securityClosePX,proto3" json:"securityClosePX"`
	SettlPrice           string   `protobuf:"bytes,19,opt,name=settl_price,json=settlPrice,proto3" json:"settlPrice"`
	ContractSymbol       string   `protobuf:"bytes,20,opt,name=contract_symbol,json=contractSymbol,proto3" json:"contractSymbol"`
	Num                  string   `protobuf:"bytes,21,opt,name=num,proto3" json:"num"`
	ContractID           string   `protobuf:"bytes,22,opt,name=contractID,proto3" json:"contractID"`
	MarginRatioParam1    string   `protobuf:"bytes,23,opt,name=margin_ratio_param1,json=marginRatioParam1,proto3" json:"marginRatioParam1"`
	MarginRatioParam2    string   `protobuf:"bytes,24,opt,name=margin_ratio_param2,json=marginRatioParam2,proto3" json:"marginRatioParam2"`
	LmtOrdMinFloor       string   `protobuf:"bytes,25,opt,name=lmt_ord_min_floor,json=lmtOrdMinFloor,proto3" json:"lmtOrdMinFloor"`
	MktOrdMinFloor       string   `protobuf:"bytes,26,opt,name=mkt_ord_min_floor,json=mktOrdMinFloor,proto3" json:"mktOrdMinFloor"`
	EndDate              string   `protobuf:"bytes,27,opt,name=end_date,json=endDate,proto3" json:"endDate"`
	PriceLimitType       string   `protobuf:"bytes,28,opt,name=price_limit_type,json=priceLimitType,proto3" json:"priceLimitType"`
	ExerciseDate         string   `protobuf:"bytes,29,opt,name=exercise_date,json=exerciseDate,proto3" json:"exerciseDate"`
	MarginUnit           string   `protobuf:"bytes,30,opt,name=margin_unit,json=marginUnit,proto3" json:"marginUnit"`
	SecurityID           string   `protobuf:"bytes,31,opt,name=securityID,proto3" json:"securityID"`
	SecurityNameByID     string   `protobuf:"bytes,32,opt,name=security_name_byID,json=securityNameByID,proto3" json:"securityNameByID"`
	ContractFlag         string   `protobuf:"bytes,33,opt,name=contract_flag,json=contractFlag,proto3" json:"contractFlag"`
	UnderlyingClosePX    string   `protobuf:"bytes,34,opt,name=underlying_closePX,json=underlyingClosePX,proto3" json:"underlyingClosePX"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SSEStockOption) Reset()         { *m = SSEStockOption{} }
func (m *SSEStockOption) String() string { return proto.CompactTextString(m) }
func (*SSEStockOption) ProtoMessage()    {}
func (*SSEStockOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_eab2a290fad82415, []int{0}
}

func (m *SSEStockOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSEStockOption.Unmarshal(m, b)
}
func (m *SSEStockOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSEStockOption.Marshal(b, m, deterministic)
}
func (m *SSEStockOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSEStockOption.Merge(m, src)
}
func (m *SSEStockOption) XXX_Size() int {
	return xxx_messageInfo_SSEStockOption.Size(m)
}
func (m *SSEStockOption) XXX_DiscardUnknown() {
	xxx_messageInfo_SSEStockOption.DiscardUnknown(m)
}

var xxx_messageInfo_SSEStockOption proto.InternalMessageInfo

func (m *SSEStockOption) GetExercisePrice() string {
	if m != nil {
		return m.ExercisePrice
	}
	return ""
}

func (m *SSEStockOption) GetUpdateVersion() string {
	if m != nil {
		return m.UpdateVersion
	}
	return ""
}

func (m *SSEStockOption) GetOptionType() string {
	if m != nil {
		return m.OptionType
	}
	return ""
}

func (m *SSEStockOption) GetDailyPriceUpLimit() string {
	if m != nil {
		return m.DailyPriceUpLimit
	}
	return ""
}

func (m *SSEStockOption) GetTimeSave() string {
	if m != nil {
		return m.TimeSave
	}
	return ""
}

func (m *SSEStockOption) GetDELIST_Flag() string {
	if m != nil {
		return m.DELIST_Flag
	}
	return ""
}

func (m *SSEStockOption) GetStartDate() string {
	if m != nil {
		return m.StartDate
	}
	return ""
}

func (m *SSEStockOption) GetExpireDate() string {
	if m != nil {
		return m.ExpireDate
	}
	return ""
}

func (m *SSEStockOption) GetContractUnit() string {
	if m != nil {
		return m.ContractUnit
	}
	return ""
}

func (m *SSEStockOption) GetCallOrPut() string {
	if m != nil {
		return m.CallOrPut
	}
	return ""
}

func (m *SSEStockOption) GetLmtOrdMaxFloor() string {
	if m != nil {
		return m.LmtOrdMaxFloor
	}
	return ""
}

func (m *SSEStockOption) GetDeliveryDate() string {
	if m != nil {
		return m.DeliveryDate
	}
	return ""
}

func (m *SSEStockOption) GetChangeFlag() string {
	if m != nil {
		return m.ChangeFlag
	}
	return ""
}

func (m *SSEStockOption) GetMktOrdMaxFloor() string {
	if m != nil {
		return m.MktOrdMaxFloor
	}
	return ""
}

func (m *SSEStockOption) GetUnderlyingType() string {
	if m != nil {
		return m.UnderlyingType
	}
	return ""
}

func (m *SSEStockOption) GetDailyPriceDownLimit() string {
	if m != nil {
		return m.DailyPriceDownLimit
	}
	return ""
}

func (m *SSEStockOption) GetRoundLot() string {
	if m != nil {
		return m.RoundLot
	}
	return ""
}

func (m *SSEStockOption) GetSecurityClosePX() string {
	if m != nil {
		return m.SecurityClosePX
	}
	return ""
}

func (m *SSEStockOption) GetSettlPrice() string {
	if m != nil {
		return m.SettlPrice
	}
	return ""
}

func (m *SSEStockOption) GetContractSymbol() string {
	if m != nil {
		return m.ContractSymbol
	}
	return ""
}

func (m *SSEStockOption) GetNum() string {
	if m != nil {
		return m.Num
	}
	return ""
}

func (m *SSEStockOption) GetContractID() string {
	if m != nil {
		return m.ContractID
	}
	return ""
}

func (m *SSEStockOption) GetMarginRatioParam1() string {
	if m != nil {
		return m.MarginRatioParam1
	}
	return ""
}

func (m *SSEStockOption) GetMarginRatioParam2() string {
	if m != nil {
		return m.MarginRatioParam2
	}
	return ""
}

func (m *SSEStockOption) GetLmtOrdMinFloor() string {
	if m != nil {
		return m.LmtOrdMinFloor
	}
	return ""
}

func (m *SSEStockOption) GetMktOrdMinFloor() string {
	if m != nil {
		return m.MktOrdMinFloor
	}
	return ""
}

func (m *SSEStockOption) GetEndDate() string {
	if m != nil {
		return m.EndDate
	}
	return ""
}

func (m *SSEStockOption) GetPriceLimitType() string {
	if m != nil {
		return m.PriceLimitType
	}
	return ""
}

func (m *SSEStockOption) GetExerciseDate() string {
	if m != nil {
		return m.ExerciseDate
	}
	return ""
}

func (m *SSEStockOption) GetMarginUnit() string {
	if m != nil {
		return m.MarginUnit
	}
	return ""
}

func (m *SSEStockOption) GetSecurityID() string {
	if m != nil {
		return m.SecurityID
	}
	return ""
}

func (m *SSEStockOption) GetSecurityNameByID() string {
	if m != nil {
		return m.SecurityNameByID
	}
	return ""
}

func (m *SSEStockOption) GetContractFlag() string {
	if m != nil {
		return m.ContractFlag
	}
	return ""
}

func (m *SSEStockOption) GetUnderlyingClosePX() string {
	if m != nil {
		return m.UnderlyingClosePX
	}
	return ""
}

func init() {
	proto.RegisterType((*SSEStockOption)(nil), "goshare.spider.SSEStockOption")
}

func init() { proto.RegisterFile("spider/sse.proto", fileDescriptor_eab2a290fad82415) }

var fileDescriptor_eab2a290fad82415 = []byte{
	// 666 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x94, 0x6f, 0x4f, 0x1b, 0x39,
	0x10, 0xc6, 0xc5, 0x71, 0x07, 0x64, 0x80, 0xfc, 0x31, 0x1c, 0x67, 0x8e, 0xf2, 0xa7, 0xa0, 0x0a,
	0x90, 0x5a, 0x50, 0xe1, 0x1b, 0xb4, 0x01, 0x09, 0x89, 0x96, 0x88, 0x40, 0xd5, 0x77, 0x96, 0xd9,
	0x35, 0xc1, 0x62, 0xd7, 0x5e, 0x79, 0xbd, 0x21, 0xf9, 0x1e, 0xfd, 0xc0, 0x95, 0x67, 0xe2, 0x25,
	0xa9, 0x78, 0x17, 0xfd, 0xe6, 0x89, 0xc7, 0xe3, 0xe7, 0x99, 0x85, 0x76, 0x59, 0xe8, 0x54, 0xb9,
	0xd3, 0xb2, 0x54, 0x27, 0x85, 0xb3, 0xde, 0xb2, 0xe6, 0xc0, 0x96, 0x4f, 0xd2, 0xa9, 0x13, 0xaa,
	0xec, 0xff, 0x02, 0x68, 0xf6, 0xfb, 0x17, 0x7d, 0x6f, 0x93, 0xe7, 0x9b, 0xc2, 0x6b, 0x6b, 0xd8,
	0x07, 0x68, 0xaa, 0x91, 0x72, 0x89, 0x2e, 0x95, 0x28, 0x9c, 0x4e, 0x14, 0x9f, 0xdb, 0x9b, 0x3b,
	0x6a, 0xdc, 0xae, 0x46, 0xda, 0x0b, 0x30, 0xc8, 0xaa, 0x22, 0x95, 0x5e, 0x89, 0xa1, 0x72, 0xa5,
	0xb6, 0x86, 0xff, 0x45, 0x32, 0xa2, 0x3f, 0x08, 0xb2, 0x5d, 0x58, 0xb6, 0x78, 0xae, 0xf0, 0xe3,
	0x42, 0xf1, 0x79, 0xd4, 0x00, 0xa1, 0xbb, 0x71, 0xa1, 0xd8, 0x29, 0xac, 0xa7, 0x52, 0x67, 0x63,
	0xea, 0x25, 0xaa, 0x42, 0x64, 0x3a, 0xd7, 0x9e, 0xff, 0x8d, 0xca, 0x0e, 0xd6, 0xb0, 0xe3, 0x7d,
	0x71, 0x1d, 0x0a, 0x6c, 0x0b, 0x1a, 0x5e, 0xe7, 0x4a, 0x94, 0x72, 0xa8, 0xf8, 0x3f, 0xa8, 0x5a,
	0x0a, 0xa0, 0x2f, 0x87, 0x2a, 0xb4, 0xeb, 0x5e, 0x5c, 0x5f, 0xf5, 0xef, 0xc4, 0x65, 0x26, 0x07,
	0x7c, 0x81, 0xda, 0x11, 0x0a, 0x84, 0x6d, 0x03, 0x94, 0x5e, 0x3a, 0x2f, 0xc2, 0x25, 0xf9, 0x22,
	0xd6, 0x1b, 0x48, 0xba, 0xd2, 0xe3, 0xff, 0xd5, 0xa8, 0xd0, 0x4e, 0x51, 0x7d, 0x89, 0xfe, 0x4f,
	0x08, 0x05, 0x07, 0xb0, 0x9a, 0x58, 0xe3, 0x9d, 0x4c, 0xbc, 0xa8, 0x8c, 0xf6, 0xbc, 0x81, 0x92,
	0x95, 0x08, 0xef, 0x8d, 0xf6, 0x6c, 0x07, 0x96, 0x13, 0x99, 0x65, 0xc2, 0x3a, 0x51, 0x54, 0x9e,
	0x03, 0x75, 0x09, 0xe8, 0xc6, 0xf5, 0x2a, 0xcf, 0x8e, 0xa1, 0x93, 0xe5, 0x5e, 0x58, 0x97, 0x8a,
	0x5c, 0x8e, 0xc4, 0x63, 0x66, 0xad, 0xe3, 0xcb, 0xa8, 0x6a, 0x66, 0xb9, 0xbf, 0x71, 0xe9, 0x37,
	0x39, 0xba, 0x0c, 0x34, 0xf4, 0x4b, 0x55, 0xa6, 0x87, 0xca, 0x8d, 0xe9, 0x4a, 0x2b, 0xd4, 0x2f,
	0xc2, 0x78, 0xeb, 0xe4, 0x49, 0x9a, 0x81, 0x12, 0x8f, 0x61, 0xea, 0x55, 0xba, 0x35, 0x21, 0x9c,
	0xfa, 0x18, 0x3a, 0xf9, 0xf3, 0x9f, 0x0d, 0x9b, 0xd4, 0x30, 0x7f, 0x9e, 0x69, 0x78, 0x08, 0xad,
	0xca, 0xa4, 0xca, 0x65, 0x63, 0x6d, 0x06, 0x64, 0x5a, 0x8b, 0x84, 0xaf, 0x18, 0x8d, 0x3b, 0x87,
	0x8d, 0x69, 0xe3, 0x52, 0xfb, 0x62, 0x26, 0xd6, 0xb5, 0x51, 0xbf, 0xf6, 0x6a, 0x5d, 0xd7, 0xbe,
	0x98, 0xda, 0x3c, 0x67, 0x2b, 0x93, 0x8a, 0xcc, 0x7a, 0xde, 0x21, 0xf3, 0x10, 0x5c, 0xdb, 0xf0,
	0x2c, 0xed, 0x52, 0x25, 0x95, 0xd3, 0x7e, 0x2c, 0x92, 0xcc, 0x96, 0xaa, 0xf7, 0x93, 0x33, 0xd4,
	0xb4, 0x22, 0xff, 0x4a, 0x38, 0x4c, 0x5c, 0x2a, 0xef, 0xb3, 0x49, 0x42, 0xd7, 0x68, 0x62, 0x44,
	0x14, 0xcf, 0x43, 0x68, 0xd5, 0x3e, 0x95, 0xe3, 0xfc, 0xc1, 0x66, 0x7c, 0x9d, 0xc6, 0x88, 0xb8,
	0x8f, 0x94, 0xb5, 0x61, 0xde, 0x54, 0x39, 0xff, 0x17, 0x8b, 0xe1, 0x27, 0xdb, 0x01, 0x88, 0x9a,
	0xab, 0x2e, 0xdf, 0x98, 0x3c, 0x66, 0x4d, 0xd8, 0x09, 0xac, 0xe5, 0xd2, 0x0d, 0xb4, 0x11, 0x4e,
	0x7a, 0x6d, 0x45, 0x21, 0x9d, 0xcc, 0x3f, 0xf3, 0xff, 0x28, 0xb0, 0x54, 0xba, 0x0d, 0x95, 0x1e,
	0x16, 0xde, 0xd6, 0x9f, 0x71, 0xfe, 0xb6, 0xfe, 0x6c, 0x26, 0x1d, 0xda, 0x4c, 0xcc, 0xda, 0x9c,
	0x49, 0x87, 0x36, 0x64, 0xd6, 0xb4, 0xaf, 0xb5, 0xf4, 0xff, 0x19, 0x5f, 0xa3, 0x74, 0x13, 0x96,
	0x94, 0x49, 0x29, 0x43, 0x5b, 0xa8, 0x58, 0x54, 0x26, 0xc5, 0xf8, 0x1c, 0x41, 0x9b, 0x3c, 0x44,
	0xfb, 0xc8, 0xf3, 0x77, 0x74, 0x08, 0x72, 0xb4, 0x0e, 0x3d, 0x3f, 0x80, 0xfa, 0x2b, 0x40, 0x27,
	0x6d, 0x53, 0x1a, 0x23, 0x8c, 0x69, 0x9c, 0xcc, 0x8b, 0x0b, 0xb2, 0x43, 0x0f, 0x48, 0x68, 0xb2,
	0x1e, 0x10, 0xfd, 0xbc, 0xea, 0xf2, 0xdd, 0xe8, 0x5d, 0x24, 0xec, 0x23, 0xb0, 0x3a, 0x07, 0x46,
	0xe6, 0x4a, 0x3c, 0x04, 0xdd, 0x1e, 0xea, 0xea, 0x84, 0x7c, 0x97, 0xb9, 0xfa, 0x12, 0xd4, 0xd3,
	0x1b, 0x89, 0xf1, 0x7f, 0x3f, 0xbb, 0x91, 0xb8, 0x00, 0x9f, 0x80, 0x4d, 0xa5, 0x3a, 0x86, 0x6b,
	0x9f, 0x2c, 0x78, 0xad, 0x4c, 0xe2, 0xf5, 0xb0, 0x80, 0x5f, 0xcb, 0xf3, 0xdf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xe2, 0xab, 0xd5, 0x90, 0x41, 0x05, 0x00, 0x00,
}