// Code generated by protoc-gen-go. DO NOT EDIT.
// source: data/ordervo/ordervo.proto

/*
Package ordervo is a generated protocol buffer package.

It is generated from these files:
	data/ordervo/ordervo.proto

It has these top-level messages:
	OrderVo
	SuborderVo
	SuborderProductVo
*/
package ordervo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type OrderVo struct {
	OrderId               string  `protobuf:"bytes,1,opt,name=orderId" json:"orderId,omitempty"`
	MemberCode            int64   `protobuf:"varint,2,opt,name=memberCode" json:"memberCode,omitempty"`
	Channel               string  `protobuf:"bytes,3,opt,name=channel" json:"channel,omitempty"`
	TotalAmount           float64 `protobuf:"fixed64,4,opt,name=totalAmount" json:"totalAmount,omitempty"`
	TotalFreight          float64 `protobuf:"fixed64,5,opt,name=totalFreight" json:"totalFreight,omitempty"`
	VipDiscount           float64 `protobuf:"fixed64,6,opt,name=vipDiscount" json:"vipDiscount,omitempty"`
	CashAmount            float64 `protobuf:"fixed64,7,opt,name=cashAmount" json:"cashAmount,omitempty"`
	PointAmount           float64 `protobuf:"fixed64,8,opt,name=pointAmount" json:"pointAmount,omitempty"`
	PointRate             int32   `protobuf:"varint,9,opt,name=pointRate" json:"pointRate,omitempty"`
	OrderStatus           int32   `protobuf:"varint,10,opt,name=orderStatus" json:"orderStatus,omitempty"`
	RecvAreaId            int32   `protobuf:"varint,11,opt,name=recvAreaId" json:"recvAreaId,omitempty"`
	RecvAddress           string  `protobuf:"bytes,12,opt,name=recvAddress" json:"recvAddress,omitempty"`
	PayMethod             int32   `protobuf:"varint,13,opt,name=payMethod" json:"payMethod,omitempty"`
	PayStartTime          int32   `protobuf:"varint,14,opt,name=payStartTime" json:"payStartTime,omitempty"`
	PayEndTime            int32   `protobuf:"varint,15,opt,name=payEndTime" json:"payEndTime,omitempty"`
	CreateTime            int32   `protobuf:"varint,16,opt,name=createTime" json:"createTime,omitempty"`
	UpdateTime            int32   `protobuf:"varint,17,opt,name=updateTime" json:"updateTime,omitempty"`
	WillBeClosedAt        int32   `protobuf:"varint,18,opt,name=willBeClosedAt" json:"willBeClosedAt,omitempty"`
	RedPacketID           string  `protobuf:"bytes,19,opt,name=redPacketID" json:"redPacketID,omitempty"`
	RedPacketDenomination float64 `protobuf:"fixed64,20,opt,name=redPacketDenomination" json:"redPacketDenomination,omitempty"`
	PointType             int32   `protobuf:"varint,21,opt,name=pointType" json:"pointType,omitempty"`
	RedPacketActivityID   string  `protobuf:"bytes,22,opt,name=redPacketActivityID" json:"redPacketActivityID,omitempty"`
	CoveredRedPacketID    string  `protobuf:"bytes,23,opt,name=coveredRedPacketID" json:"coveredRedPacketID,omitempty"`
	RedPacketUsedAmount   float64 `protobuf:"fixed64,24,opt,name=redPacketUsedAmount" json:"redPacketUsedAmount,omitempty"`
	Platform              int32   `protobuf:"varint,25,opt,name=platform" json:"platform,omitempty"`
	RedPacketType         int32   `protobuf:"varint,26,opt,name=redPacketType" json:"redPacketType,omitempty"`
	RedPacketDiscountRate float64 `protobuf:"fixed64,27,opt,name=redPacketDiscountRate" json:"redPacketDiscountRate,omitempty"`
}

func (m *OrderVo) Reset()                    { *m = OrderVo{} }
func (m *OrderVo) String() string            { return proto.CompactTextString(m) }
func (*OrderVo) ProtoMessage()               {}
func (*OrderVo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *OrderVo) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *OrderVo) GetMemberCode() int64 {
	if m != nil {
		return m.MemberCode
	}
	return 0
}

func (m *OrderVo) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *OrderVo) GetTotalAmount() float64 {
	if m != nil {
		return m.TotalAmount
	}
	return 0
}

func (m *OrderVo) GetTotalFreight() float64 {
	if m != nil {
		return m.TotalFreight
	}
	return 0
}

func (m *OrderVo) GetVipDiscount() float64 {
	if m != nil {
		return m.VipDiscount
	}
	return 0
}

func (m *OrderVo) GetCashAmount() float64 {
	if m != nil {
		return m.CashAmount
	}
	return 0
}

func (m *OrderVo) GetPointAmount() float64 {
	if m != nil {
		return m.PointAmount
	}
	return 0
}

func (m *OrderVo) GetPointRate() int32 {
	if m != nil {
		return m.PointRate
	}
	return 0
}

func (m *OrderVo) GetOrderStatus() int32 {
	if m != nil {
		return m.OrderStatus
	}
	return 0
}

func (m *OrderVo) GetRecvAreaId() int32 {
	if m != nil {
		return m.RecvAreaId
	}
	return 0
}

func (m *OrderVo) GetRecvAddress() string {
	if m != nil {
		return m.RecvAddress
	}
	return ""
}

func (m *OrderVo) GetPayMethod() int32 {
	if m != nil {
		return m.PayMethod
	}
	return 0
}

func (m *OrderVo) GetPayStartTime() int32 {
	if m != nil {
		return m.PayStartTime
	}
	return 0
}

func (m *OrderVo) GetPayEndTime() int32 {
	if m != nil {
		return m.PayEndTime
	}
	return 0
}

func (m *OrderVo) GetCreateTime() int32 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *OrderVo) GetUpdateTime() int32 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

func (m *OrderVo) GetWillBeClosedAt() int32 {
	if m != nil {
		return m.WillBeClosedAt
	}
	return 0
}

func (m *OrderVo) GetRedPacketID() string {
	if m != nil {
		return m.RedPacketID
	}
	return ""
}

func (m *OrderVo) GetRedPacketDenomination() float64 {
	if m != nil {
		return m.RedPacketDenomination
	}
	return 0
}

func (m *OrderVo) GetPointType() int32 {
	if m != nil {
		return m.PointType
	}
	return 0
}

func (m *OrderVo) GetRedPacketActivityID() string {
	if m != nil {
		return m.RedPacketActivityID
	}
	return ""
}

func (m *OrderVo) GetCoveredRedPacketID() string {
	if m != nil {
		return m.CoveredRedPacketID
	}
	return ""
}

func (m *OrderVo) GetRedPacketUsedAmount() float64 {
	if m != nil {
		return m.RedPacketUsedAmount
	}
	return 0
}

func (m *OrderVo) GetPlatform() int32 {
	if m != nil {
		return m.Platform
	}
	return 0
}

func (m *OrderVo) GetRedPacketType() int32 {
	if m != nil {
		return m.RedPacketType
	}
	return 0
}

func (m *OrderVo) GetRedPacketDiscountRate() float64 {
	if m != nil {
		return m.RedPacketDiscountRate
	}
	return 0
}

type SuborderVo struct {
	// *
	// 子订单ID
	SuborderId string `protobuf:"bytes,1,opt,name=suborderId" json:"suborderId,omitempty"`
	// *
	// 大订单id
	OrderId string `protobuf:"bytes,2,opt,name=orderId" json:"orderId,omitempty"`
	// *
	// 用户ID
	MemberCode int64 `protobuf:"varint,3,opt,name=memberCode" json:"memberCode,omitempty"`
	// *
	// 渠道 ID对应客户端ID（client_id）
	Channel string `protobuf:"bytes,4,opt,name=channel" json:"channel,omitempty"`
	// *
	// 供应商ID
	MerchantId string `protobuf:"bytes,5,opt,name=merchantId" json:"merchantId,omitempty"`
	// *
	// 供应商Name
	MerchantName string `protobuf:"bytes,6,opt,name=merchantName" json:"merchantName,omitempty"`
	// *
	// 子订单商品总价
	SuborderPrice float64 `protobuf:"fixed64,7,opt,name=suborderPrice" json:"suborderPrice,omitempty"`
	// *
	// 运费
	Freight float64 `protobuf:"fixed64,8,opt,name=freight" json:"freight,omitempty"`
	// *
	// 创建时间
	CreateTime int32 `protobuf:"varint,9,opt,name=createTime" json:"createTime,omitempty"`
	// *
	//  更新时间
	UpdateTime int32 `protobuf:"varint,10,opt,name=updateTime" json:"updateTime,omitempty"`
}

func (m *SuborderVo) Reset()                    { *m = SuborderVo{} }
func (m *SuborderVo) String() string            { return proto.CompactTextString(m) }
func (*SuborderVo) ProtoMessage()               {}
func (*SuborderVo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SuborderVo) GetSuborderId() string {
	if m != nil {
		return m.SuborderId
	}
	return ""
}

func (m *SuborderVo) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *SuborderVo) GetMemberCode() int64 {
	if m != nil {
		return m.MemberCode
	}
	return 0
}

func (m *SuborderVo) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *SuborderVo) GetMerchantId() string {
	if m != nil {
		return m.MerchantId
	}
	return ""
}

func (m *SuborderVo) GetMerchantName() string {
	if m != nil {
		return m.MerchantName
	}
	return ""
}

func (m *SuborderVo) GetSuborderPrice() float64 {
	if m != nil {
		return m.SuborderPrice
	}
	return 0
}

func (m *SuborderVo) GetFreight() float64 {
	if m != nil {
		return m.Freight
	}
	return 0
}

func (m *SuborderVo) GetCreateTime() int32 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *SuborderVo) GetUpdateTime() int32 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

type SuborderProductVo struct {
	// *
	// 订单ID
	SuborderId string `protobuf:"bytes,1,opt,name=suborderId" json:"suborderId,omitempty"`
	// *
	// 商品ID
	ProductId string `protobuf:"bytes,2,opt,name=productId" json:"productId,omitempty"`
	// *
	// 订单ID
	OrderId string `protobuf:"bytes,3,opt,name=orderId" json:"orderId,omitempty"`
	// *
	// member_code
	MemberCode int64 `protobuf:"varint,4,opt,name=memberCode" json:"memberCode,omitempty"`
	// *
	// 商品名称
	ProductName string `protobuf:"bytes,5,opt,name=productName" json:"productName,omitempty"`
	// *
	// 商品类型
	ProductType int32 `protobuf:"varint,6,opt,name=productType" json:"productType,omitempty"`
	// *
	// 普通单价
	UnitPrice float64 `protobuf:"fixed64,7,opt,name=unitPrice" json:"unitPrice,omitempty"`
	// *
	// 普通购买数量
	Quantity int32 `protobuf:"varint,8,opt,name=quantity" json:"quantity,omitempty"`
	// *
	// Vip 单价
	VipUnitPrice float64 `protobuf:"fixed64,9,opt,name=vipUnitPrice" json:"vipUnitPrice,omitempty"`
	// *
	// VIP购买数量
	VipQuantity int32 `protobuf:"varint,10,opt,name=vipQuantity" json:"vipQuantity,omitempty"`
	// *
	// 商品图片URL
	ImgUrl string `protobuf:"bytes,11,opt,name=imgUrl" json:"imgUrl,omitempty"`
	// *
	// 渠道
	Channel string `protobuf:"bytes,12,opt,name=channel" json:"channel,omitempty"`
	// *
	// 预计发货时间
	DeliveryTime int32 `protobuf:"varint,13,opt,name=deliveryTime" json:"deliveryTime,omitempty"`
	// *
	// shipping_id
	ShippingId int32 `protobuf:"varint,14,opt,name=shippingId" json:"shippingId,omitempty"`
	// *
	// refund_id
	RefundId string `protobuf:"bytes,15,opt,name=refundId" json:"refundId,omitempty"`
	// *
	// 退款类型
	RefundType int32 `protobuf:"varint,16,opt,name=refundType" json:"refundType,omitempty"`
	// *
	// 创建时间
	CreateTime int32 `protobuf:"varint,17,opt,name=createTime" json:"createTime,omitempty"`
	// *
	// 更新时间
	UpdateTime int32 `protobuf:"varint,18,opt,name=updateTime" json:"updateTime,omitempty"`
	// *
	// 确认物流状态
	ConfirmState int32 `protobuf:"varint,19,opt,name=confirmState" json:"confirmState,omitempty"`
	// *
	// 确认物流状态时间
	ConfirmTime int32 `protobuf:"varint,20,opt,name=confirmTime" json:"confirmTime,omitempty"`
	// *
	// 拆单现金值
	CashAmount float64 `protobuf:"fixed64,21,opt,name=cashAmount" json:"cashAmount,omitempty"`
	// *
	// 拆单积分值
	PointAmount float64 `protobuf:"fixed64,22,opt,name=pointAmount" json:"pointAmount,omitempty"`
	// *
	// 积分类型
	PointType int32 `protobuf:"varint,23,opt,name=pointType" json:"pointType,omitempty"`
	// *
	// 积分现金兑换比率
	PointRate int32 `protobuf:"varint,24,opt,name=pointRate" json:"pointRate,omitempty"`
	// *
	// 积分价值（抵扣现金额）
	PointValue float64 `protobuf:"fixed64,25,opt,name=pointValue" json:"pointValue,omitempty"`
	// *
	// 红包抵扣现金额
	RedPacketValue float64 `protobuf:"fixed64,26,opt,name=redPacketValue" json:"redPacketValue,omitempty"`
	// *
	// 商品全称
	ProductFullName string `protobuf:"bytes,27,opt,name=productFullName" json:"productFullName,omitempty"`
	// *
	// 兑换积分价格
	ExchangePointPrice float64 `protobuf:"fixed64,28,opt,name=exchangePointPrice" json:"exchangePointPrice,omitempty"`
	// *
	// 兑换现金价格
	ExchangeCashPrice float64 `protobuf:"fixed64,29,opt,name=exchangeCashPrice" json:"exchangeCashPrice,omitempty"`
	//
	// 售后详情
	AfterSales string `protobuf:"bytes,30,opt,name=afterSales" json:"afterSales,omitempty"`
}

func (m *SuborderProductVo) Reset()                    { *m = SuborderProductVo{} }
func (m *SuborderProductVo) String() string            { return proto.CompactTextString(m) }
func (*SuborderProductVo) ProtoMessage()               {}
func (*SuborderProductVo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SuborderProductVo) GetSuborderId() string {
	if m != nil {
		return m.SuborderId
	}
	return ""
}

func (m *SuborderProductVo) GetProductId() string {
	if m != nil {
		return m.ProductId
	}
	return ""
}

func (m *SuborderProductVo) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *SuborderProductVo) GetMemberCode() int64 {
	if m != nil {
		return m.MemberCode
	}
	return 0
}

func (m *SuborderProductVo) GetProductName() string {
	if m != nil {
		return m.ProductName
	}
	return ""
}

func (m *SuborderProductVo) GetProductType() int32 {
	if m != nil {
		return m.ProductType
	}
	return 0
}

func (m *SuborderProductVo) GetUnitPrice() float64 {
	if m != nil {
		return m.UnitPrice
	}
	return 0
}

func (m *SuborderProductVo) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *SuborderProductVo) GetVipUnitPrice() float64 {
	if m != nil {
		return m.VipUnitPrice
	}
	return 0
}

func (m *SuborderProductVo) GetVipQuantity() int32 {
	if m != nil {
		return m.VipQuantity
	}
	return 0
}

func (m *SuborderProductVo) GetImgUrl() string {
	if m != nil {
		return m.ImgUrl
	}
	return ""
}

func (m *SuborderProductVo) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *SuborderProductVo) GetDeliveryTime() int32 {
	if m != nil {
		return m.DeliveryTime
	}
	return 0
}

func (m *SuborderProductVo) GetShippingId() int32 {
	if m != nil {
		return m.ShippingId
	}
	return 0
}

func (m *SuborderProductVo) GetRefundId() string {
	if m != nil {
		return m.RefundId
	}
	return ""
}

func (m *SuborderProductVo) GetRefundType() int32 {
	if m != nil {
		return m.RefundType
	}
	return 0
}

func (m *SuborderProductVo) GetCreateTime() int32 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *SuborderProductVo) GetUpdateTime() int32 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

func (m *SuborderProductVo) GetConfirmState() int32 {
	if m != nil {
		return m.ConfirmState
	}
	return 0
}

func (m *SuborderProductVo) GetConfirmTime() int32 {
	if m != nil {
		return m.ConfirmTime
	}
	return 0
}

func (m *SuborderProductVo) GetCashAmount() float64 {
	if m != nil {
		return m.CashAmount
	}
	return 0
}

func (m *SuborderProductVo) GetPointAmount() float64 {
	if m != nil {
		return m.PointAmount
	}
	return 0
}

func (m *SuborderProductVo) GetPointType() int32 {
	if m != nil {
		return m.PointType
	}
	return 0
}

func (m *SuborderProductVo) GetPointRate() int32 {
	if m != nil {
		return m.PointRate
	}
	return 0
}

func (m *SuborderProductVo) GetPointValue() float64 {
	if m != nil {
		return m.PointValue
	}
	return 0
}

func (m *SuborderProductVo) GetRedPacketValue() float64 {
	if m != nil {
		return m.RedPacketValue
	}
	return 0
}

func (m *SuborderProductVo) GetProductFullName() string {
	if m != nil {
		return m.ProductFullName
	}
	return ""
}

func (m *SuborderProductVo) GetExchangePointPrice() float64 {
	if m != nil {
		return m.ExchangePointPrice
	}
	return 0
}

func (m *SuborderProductVo) GetExchangeCashPrice() float64 {
	if m != nil {
		return m.ExchangeCashPrice
	}
	return 0
}

func (m *SuborderProductVo) GetAfterSales() string {
	if m != nil {
		return m.AfterSales
	}
	return ""
}

func init() {
	proto.RegisterType((*OrderVo)(nil), "ordervo.OrderVo")
	proto.RegisterType((*SuborderVo)(nil), "ordervo.SuborderVo")
	proto.RegisterType((*SuborderProductVo)(nil), "ordervo.SuborderProductVo")
}

func init() { proto.RegisterFile("data/ordervo/ordervo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 872 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x96, 0xcd, 0x72, 0xdb, 0x36,
	0x10, 0xc7, 0x47, 0xfe, 0x52, 0x88, 0x28, 0x1f, 0x46, 0x62, 0x07, 0x75, 0x52, 0x8f, 0x46, 0xd3,
	0xe9, 0xe8, 0xd0, 0x49, 0x3b, 0xd3, 0xbe, 0x80, 0x6b, 0x35, 0x33, 0x3a, 0xb4, 0x75, 0xa9, 0x38,
	0x77, 0x98, 0x58, 0x49, 0x98, 0x92, 0x04, 0x0b, 0x82, 0x6a, 0xf5, 0x7e, 0x7d, 0x85, 0x3e, 0x4a,
	0xef, 0x9d, 0x5d, 0x90, 0x14, 0x48, 0xab, 0x76, 0x4e, 0xd6, 0xfe, 0x76, 0x01, 0x2e, 0xfe, 0xd8,
	0x5d, 0x98, 0x5d, 0x28, 0xe9, 0xe4, 0xb7, 0xc6, 0x2a, 0xb0, 0x1b, 0xd3, 0xfc, 0x7d, 0x5f, 0x58,
	0xe3, 0x0c, 0x1f, 0xd6, 0xe6, 0xe4, 0x9f, 0x21, 0x1b, 0xfe, 0x8a, 0xbf, 0x3f, 0x19, 0x2e, 0x98,
	0xc7, 0x73, 0x25, 0x06, 0xe3, 0xc1, 0x34, 0x8a, 0x1b, 0x93, 0x5f, 0x32, 0x96, 0x41, 0x76, 0x07,
	0xf6, 0xda, 0x28, 0x10, 0x07, 0xe3, 0xc1, 0xf4, 0x30, 0x0e, 0x08, 0xae, 0x4c, 0xd6, 0x32, 0xcf,
	0x21, 0x15, 0x87, 0x7e, 0x65, 0x6d, 0xf2, 0x31, 0x7b, 0xea, 0x8c, 0x93, 0xe9, 0x55, 0x66, 0xaa,
	0xdc, 0x89, 0xa3, 0xf1, 0x60, 0x3a, 0x88, 0x43, 0xc4, 0x27, 0x6c, 0x44, 0xe6, 0x07, 0x0b, 0x7a,
	0xb5, 0x76, 0xe2, 0x98, 0x42, 0x3a, 0x0c, 0x77, 0xd9, 0xe8, 0x62, 0xa6, 0xcb, 0x84, 0x76, 0x39,
	0xf1, 0xbb, 0x04, 0x08, 0x33, 0x4c, 0x64, 0xb9, 0xae, 0x3f, 0x33, 0xa4, 0x80, 0x80, 0xe0, 0x0e,
	0x85, 0xd1, 0xb9, 0xab, 0x03, 0x9e, 0xf8, 0x1d, 0x02, 0xc4, 0xdf, 0xb1, 0x88, 0xcc, 0x58, 0x3a,
	0x10, 0xd1, 0x78, 0x30, 0x3d, 0x8e, 0x77, 0x00, 0xd7, 0x93, 0x18, 0x0b, 0x27, 0x5d, 0x55, 0x0a,
	0x46, 0xfe, 0x10, 0x61, 0x06, 0x16, 0x92, 0xcd, 0x95, 0x05, 0x39, 0x57, 0xe2, 0x29, 0x05, 0x04,
	0x04, 0x77, 0x20, 0x4b, 0x29, 0x0b, 0x65, 0x29, 0x46, 0xa4, 0x53, 0x88, 0x28, 0x03, 0xb9, 0xfd,
	0x19, 0xdc, 0xda, 0x28, 0xf1, 0xac, 0xce, 0xa0, 0x01, 0xa8, 0x53, 0x21, 0xb7, 0x0b, 0x27, 0xad,
	0xfb, 0xa8, 0x33, 0x10, 0xcf, 0x29, 0xa0, 0xc3, 0x30, 0x87, 0x42, 0x6e, 0x7f, 0xca, 0x15, 0x45,
	0xbc, 0xf0, 0x39, 0xec, 0x08, 0xa9, 0x64, 0x41, 0x3a, 0x20, 0xff, 0x4b, 0xef, 0xdf, 0x11, 0xf4,
	0x57, 0x85, 0x6a, 0xfc, 0xa7, 0xde, 0xbf, 0x23, 0xfc, 0x6b, 0xf6, 0xfc, 0x4f, 0x9d, 0xa6, 0x3f,
	0xc2, 0x75, 0x6a, 0x4a, 0x50, 0x57, 0x4e, 0x70, 0x8a, 0xe9, 0x51, 0x7f, 0x56, 0x75, 0x23, 0x93,
	0xdf, 0xc1, 0xcd, 0x67, 0xe2, 0x55, 0x73, 0xd6, 0x16, 0xf1, 0x1f, 0xd8, 0x59, 0x6b, 0xce, 0x20,
	0x37, 0x99, 0xce, 0xa5, 0xd3, 0x26, 0x17, 0xaf, 0xe9, 0x66, 0xf6, 0x3b, 0xdb, 0x3b, 0xfa, 0xb8,
	0x2d, 0x40, 0x9c, 0x05, 0x77, 0x84, 0x80, 0x7f, 0xc7, 0x5e, 0xb5, 0xcb, 0xae, 0x12, 0xa7, 0x37,
	0xda, 0x6d, 0xe7, 0x33, 0x71, 0x4e, 0x5f, 0xdf, 0xe7, 0xe2, 0xef, 0x19, 0x4f, 0xcc, 0x06, 0x2c,
	0xa8, 0x38, 0x48, 0xf7, 0x0d, 0x2d, 0xd8, 0xe3, 0xe9, 0x7c, 0xe1, 0x16, 0x8f, 0xea, 0xab, 0x49,
	0x50, 0xce, 0xfb, 0x5c, 0xfc, 0x82, 0x3d, 0x29, 0x52, 0xe9, 0x96, 0xc6, 0x66, 0xe2, 0x0b, 0x4a,
	0xb8, 0xb5, 0xf9, 0x57, 0xec, 0x59, 0xbb, 0x84, 0x4e, 0x74, 0x41, 0x01, 0x5d, 0xd8, 0x55, 0xaa,
	0x2e, 0x77, 0xaa, 0xd1, 0xb7, 0x7d, 0xa5, 0x02, 0xe7, 0xe4, 0xef, 0x03, 0xc6, 0x16, 0xd5, 0x9d,
	0xa9, 0x5b, 0xfb, 0x92, 0xb1, 0xb2, 0xb6, 0xda, 0xee, 0x0e, 0x48, 0xd8, 0xfa, 0x07, 0x0f, 0xb5,
	0xfe, 0xe1, 0x43, 0xad, 0x7f, 0xd4, 0x6d, 0x7d, 0x5a, 0x69, 0xd1, 0x72, 0x73, 0x45, 0x6d, 0x1d,
	0xc5, 0x01, 0xc1, 0x82, 0x6e, 0xac, 0x5f, 0x64, 0x06, 0xd4, 0xd5, 0x51, 0xdc, 0x61, 0x28, 0x51,
	0x93, 0xe5, 0x8d, 0xd5, 0x09, 0xd4, 0x9d, 0xdd, 0x85, 0x98, 0xc3, 0xb2, 0x9e, 0x1e, 0xbe, 0xb1,
	0x1b, 0xb3, 0x57, 0xf0, 0xd1, 0x23, 0x05, 0xcf, 0xfa, 0x05, 0x3f, 0xf9, 0x77, 0xc8, 0x4e, 0x17,
	0xed, 0xb7, 0x8c, 0xaa, 0x12, 0xf7, 0x19, 0x6a, 0x62, 0x99, 0xfa, 0xe0, 0x56, 0xcf, 0x1d, 0x08,
	0xb5, 0x3e, 0x7c, 0x48, 0xeb, 0xa3, 0x7b, 0x5a, 0xe3, 0x10, 0xf3, 0xdb, 0x90, 0x60, 0x5e, 0xd2,
	0x10, 0x05, 0x11, 0x54, 0x50, 0x27, 0x7e, 0x4c, 0x05, 0x08, 0x73, 0xab, 0x72, 0xed, 0x42, 0x35,
	0x77, 0x00, 0xcb, 0xf5, 0x8f, 0x4a, 0xe6, 0x4e, 0xbb, 0x2d, 0x49, 0x79, 0x1c, 0xb7, 0x36, 0xde,
	0xd7, 0x46, 0x17, 0xb7, 0xed, 0xe2, 0xc8, 0x0f, 0xea, 0x90, 0xd5, 0x83, 0xfa, 0xb7, 0x66, 0x8b,
	0x7a, 0x4c, 0x06, 0x88, 0x9f, 0xb3, 0x13, 0x9d, 0xad, 0x6e, 0x6d, 0x4a, 0x23, 0x32, 0x8a, 0x6b,
	0x2b, 0xac, 0xa3, 0x51, 0xb7, 0x8e, 0x26, 0x6c, 0xa4, 0x20, 0xd5, 0x1b, 0xb0, 0x5b, 0xba, 0x25,
	0x3f, 0x19, 0x3b, 0x8c, 0x6e, 0x64, 0xad, 0x8b, 0x42, 0xe7, 0xab, 0xb9, 0xaa, 0x47, 0x63, 0x40,
	0xf0, 0x5c, 0x16, 0x96, 0x55, 0xae, 0xe6, 0x8a, 0xc6, 0x62, 0x14, 0xb7, 0xb6, 0x1f, 0xdc, 0xf8,
	0x9b, 0x24, 0x7b, 0xd9, 0x0c, 0xee, 0x86, 0xf4, 0x6a, 0xe8, 0xf4, 0x91, 0x1a, 0xe2, 0xf7, 0x86,
	0xe6, 0x84, 0x8d, 0x12, 0x93, 0x2f, 0xb5, 0xcd, 0xf0, 0xa5, 0x00, 0x9a, 0x86, 0xc7, 0x71, 0x87,
	0xa1, 0x6e, 0xb5, 0x4d, 0x9b, 0xbc, 0xf6, 0xba, 0x05, 0xa8, 0xf7, 0xc0, 0x9d, 0x3d, 0xf6, 0xc0,
	0x9d, 0xff, 0xff, 0x03, 0x47, 0xc7, 0x7c, 0xd3, 0x1f, 0x9e, 0x9d, 0xe7, 0x4f, 0xf4, 0x9f, 0x3f,
	0x7c, 0x58, 0xd0, 0xf8, 0x24, 0xd3, 0x0a, 0x68, 0x90, 0x0d, 0xe2, 0x80, 0xe0, 0xc3, 0xd0, 0xce,
	0x21, 0x1f, 0x73, 0x41, 0x31, 0x3d, 0xca, 0xa7, 0xec, 0x45, 0x5d, 0x8c, 0x1f, 0xaa, 0x34, 0xa5,
	0x2a, 0x7e, 0x4b, 0xd7, 0xd1, 0xc7, 0x38, 0x9a, 0xe1, 0x2f, 0x2c, 0x81, 0x15, 0xdc, 0xe0, 0x77,
	0x7c, 0xcd, 0xbd, 0xa3, 0x5d, 0xf7, 0x78, 0xf8, 0x37, 0xec, 0xb4, 0xa1, 0xd7, 0xb2, 0x5c, 0xfb,
	0xf0, 0x2f, 0x29, 0xfc, 0xbe, 0x03, 0xcf, 0x23, 0x97, 0x0e, 0xec, 0x42, 0xa6, 0x50, 0x8a, 0x4b,
	0xdf, 0xc1, 0x3b, 0x72, 0x77, 0x42, 0xff, 0x26, 0x7d, 0xff, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xc4, 0x39, 0xd0, 0x2d, 0x44, 0x09, 0x00, 0x00,
}