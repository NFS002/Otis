// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service/merchant/proto/merchant/merchant.proto

package merchant

import (
	proto2 "dtypes/general-merchant/proto"
	proto1 "dtypes/partner-merchant/proto"
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

type MerchantQuery struct {
	MerchantID           string   `protobuf:"bytes,1,opt,name=merchantID,proto3" json:"merchantID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MerchantQuery) Reset()         { *m = MerchantQuery{} }
func (m *MerchantQuery) String() string { return proto.CompactTextString(m) }
func (*MerchantQuery) ProtoMessage()    {}
func (*MerchantQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa912c7099ca584b, []int{0}
}

func (m *MerchantQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MerchantQuery.Unmarshal(m, b)
}
func (m *MerchantQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MerchantQuery.Marshal(b, m, deterministic)
}
func (m *MerchantQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerchantQuery.Merge(m, src)
}
func (m *MerchantQuery) XXX_Size() int {
	return xxx_messageInfo_MerchantQuery.Size(m)
}
func (m *MerchantQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_MerchantQuery.DiscardUnknown(m)
}

var xxx_messageInfo_MerchantQuery proto.InternalMessageInfo

func (m *MerchantQuery) GetMerchantID() string {
	if m != nil {
		return m.MerchantID
	}
	return ""
}

type MerchantRequest struct {
	PartnerMerchant      *proto1.PartnerMerchant `protobuf:"bytes,1,opt,name=partnerMerchant,proto3" json:"partnerMerchant,omitempty"`
	GeneralMerchant      *proto2.GeneralMerchant `protobuf:"bytes,2,opt,name=generalMerchant,proto3" json:"generalMerchant,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *MerchantRequest) Reset()         { *m = MerchantRequest{} }
func (m *MerchantRequest) String() string { return proto.CompactTextString(m) }
func (*MerchantRequest) ProtoMessage()    {}
func (*MerchantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa912c7099ca584b, []int{1}
}

func (m *MerchantRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MerchantRequest.Unmarshal(m, b)
}
func (m *MerchantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MerchantRequest.Marshal(b, m, deterministic)
}
func (m *MerchantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerchantRequest.Merge(m, src)
}
func (m *MerchantRequest) XXX_Size() int {
	return xxx_messageInfo_MerchantRequest.Size(m)
}
func (m *MerchantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MerchantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MerchantRequest proto.InternalMessageInfo

func (m *MerchantRequest) GetPartnerMerchant() *proto1.PartnerMerchant {
	if m != nil {
		return m.PartnerMerchant
	}
	return nil
}

func (m *MerchantRequest) GetGeneralMerchant() *proto2.GeneralMerchant {
	if m != nil {
		return m.GeneralMerchant
	}
	return nil
}

type MerchantsResponse struct {
	Executed             bool                      `protobuf:"varint,1,opt,name=executed,proto3" json:"executed,omitempty"`
	PartnerMerchants     []*proto1.PartnerMerchant `protobuf:"bytes,2,rep,name=partnerMerchants,proto3" json:"partnerMerchants,omitempty"`
	GeneralMerchants     []*proto2.GeneralMerchant `protobuf:"bytes,3,rep,name=generalMerchants,proto3" json:"generalMerchants,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *MerchantsResponse) Reset()         { *m = MerchantsResponse{} }
func (m *MerchantsResponse) String() string { return proto.CompactTextString(m) }
func (*MerchantsResponse) ProtoMessage()    {}
func (*MerchantsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa912c7099ca584b, []int{2}
}

func (m *MerchantsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MerchantsResponse.Unmarshal(m, b)
}
func (m *MerchantsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MerchantsResponse.Marshal(b, m, deterministic)
}
func (m *MerchantsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerchantsResponse.Merge(m, src)
}
func (m *MerchantsResponse) XXX_Size() int {
	return xxx_messageInfo_MerchantsResponse.Size(m)
}
func (m *MerchantsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MerchantsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MerchantsResponse proto.InternalMessageInfo

func (m *MerchantsResponse) GetExecuted() bool {
	if m != nil {
		return m.Executed
	}
	return false
}

func (m *MerchantsResponse) GetPartnerMerchants() []*proto1.PartnerMerchant {
	if m != nil {
		return m.PartnerMerchants
	}
	return nil
}

func (m *MerchantsResponse) GetGeneralMerchants() []*proto2.GeneralMerchant {
	if m != nil {
		return m.GeneralMerchants
	}
	return nil
}

func init() {
	proto.RegisterType((*MerchantQuery)(nil), "merchant.MerchantQuery")
	proto.RegisterType((*MerchantRequest)(nil), "merchant.MerchantRequest")
	proto.RegisterType((*MerchantsResponse)(nil), "merchant.MerchantsResponse")
}

func init() {
	proto.RegisterFile("service/merchant/proto/merchant/merchant.proto", fileDescriptor_fa912c7099ca584b)
}

var fileDescriptor_fa912c7099ca584b = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x51, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x4d, 0x03, 0x12, 0x47, 0xa4, 0xba, 0x50, 0x8c, 0x11, 0x24, 0xe4, 0xc9, 0x17, 0x13,
	0xa8, 0xde, 0xc0, 0x42, 0x51, 0x2a, 0xda, 0xf5, 0x04, 0x31, 0x1d, 0x54, 0xa8, 0x49, 0xdc, 0xdd,
	0x88, 0x3d, 0x92, 0x67, 0xf1, 0x3c, 0xbe, 0x4b, 0x93, 0xdd, 0x8d, 0x9d, 0x3c, 0x58, 0xac, 0x6f,
	0x99, 0x7f, 0xe7, 0xff, 0xf2, 0xcf, 0x64, 0x03, 0xb1, 0x44, 0xf1, 0xf6, 0x9c, 0x61, 0xf2, 0x82,
	0x22, 0x7b, 0x4a, 0x73, 0x95, 0x94, 0xa2, 0x50, 0x45, 0x5b, 0x9a, 0x87, 0xb8, 0xd6, 0x99, 0x67,
	0xea, 0xe0, 0x62, 0xa6, 0x16, 0x25, 0xca, 0xa4, 0x4c, 0x85, 0xca, 0x51, 0x9c, 0x11, 0x00, 0x95,
	0x1b, 0xbf, 0x75, 0x3d, 0x62, 0x8e, 0x22, 0x9d, 0x53, 0x17, 0x95, 0x1b, 0x57, 0x94, 0xc0, 0xde,
	0x8d, 0x56, 0xa6, 0x15, 0x8a, 0x05, 0x3b, 0x01, 0x30, 0x2d, 0x57, 0x23, 0xdf, 0x09, 0x9d, 0xd3,
	0x1d, 0xfe, 0x43, 0x89, 0x3e, 0x1c, 0xe8, 0x1b, 0x07, 0xc7, 0xd7, 0x0a, 0xa5, 0x62, 0xd7, 0xd0,
	0xd7, 0xa1, 0xcc, 0x49, 0x6d, 0xdc, 0x1d, 0x86, 0xb1, 0xd6, 0xed, 0x5b, 0xef, 0x56, 0xfb, 0x38,
	0x35, 0x2e, 0x59, 0x3a, 0xaa, 0x65, 0xf5, 0x34, 0x4b, 0xeb, 0x96, 0x35, 0x5e, 0xed, 0xe3, 0xd4,
	0x18, 0x7d, 0x3a, 0x70, 0x60, 0x0a, 0xc9, 0x51, 0x96, 0x45, 0x2e, 0x91, 0x05, 0xe0, 0xe1, 0x3b,
	0x66, 0x95, 0xc2, 0x59, 0x1d, 0xd3, 0xe3, 0xb6, 0x66, 0x13, 0xd8, 0x27, 0x81, 0xa4, 0xdf, 0x0b,
	0xdd, 0xb5, 0x46, 0xe9, 0x38, 0x97, 0x34, 0x12, 0x49, 0xfa, 0xae, 0xa6, 0xfd, 0x36, 0x4c, 0xc7,
	0x39, 0xfc, 0x72, 0xdb, 0xcd, 0xdf, 0x37, 0x77, 0x8b, 0x4d, 0x61, 0x70, 0x29, 0x30, 0x55, 0x48,
	0xec, 0xec, 0x28, 0xb6, 0x64, 0xf2, 0xb5, 0x82, 0xe3, 0xee, 0x91, 0x5d, 0x4e, 0xb4, 0xd5, 0x22,
	0xc9, 0x7c, 0x1b, 0x20, 0x27, 0xc0, 0xc6, 0xa8, 0x68, 0xc4, 0xc3, 0xae, 0xa9, 0xbe, 0x82, 0xeb,
	0xd1, 0x68, 0xba, 0xbf, 0xd2, 0x6e, 0x61, 0x30, 0xc2, 0x39, 0x76, 0x37, 0xb8, 0x31, 0xf0, 0x9f,
	0x12, 0x3e, 0x6c, 0xd7, 0x7f, 0xea, 0xf9, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe9, 0x30, 0xbd,
	0xaa, 0x51, 0x04, 0x00, 0x00,
}
