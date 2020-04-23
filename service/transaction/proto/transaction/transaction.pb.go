// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service/transaction/proto/transaction/transaction.proto

package transaction

import (
	proto1 "gitlab.com/otis_team/backend/dtypes/transaction/proto"
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

// Query transactions by User, Merchant, or ID
type TransactionQuery struct {
	Transaction          *proto1.Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	MerchantID           string              `protobuf:"bytes,2,opt,name=MerchantID,proto3" json:"MerchantID,omitempty"`
	UserID               string              `protobuf:"bytes,3,opt,name=UserID,proto3" json:"UserID,omitempty"`
	TransactionID        string              `protobuf:"bytes,4,opt,name=TransactionID,proto3" json:"TransactionID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *TransactionQuery) Reset()         { *m = TransactionQuery{} }
func (m *TransactionQuery) String() string { return proto.CompactTextString(m) }
func (*TransactionQuery) ProtoMessage()    {}
func (*TransactionQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0975306caeaf870, []int{0}
}

func (m *TransactionQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionQuery.Unmarshal(m, b)
}
func (m *TransactionQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionQuery.Marshal(b, m, deterministic)
}
func (m *TransactionQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionQuery.Merge(m, src)
}
func (m *TransactionQuery) XXX_Size() int {
	return xxx_messageInfo_TransactionQuery.Size(m)
}
func (m *TransactionQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionQuery.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionQuery proto.InternalMessageInfo

func (m *TransactionQuery) GetTransaction() *proto1.Transaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func (m *TransactionQuery) GetMerchantID() string {
	if m != nil {
		return m.MerchantID
	}
	return ""
}

func (m *TransactionQuery) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *TransactionQuery) GetTransactionID() string {
	if m != nil {
		return m.TransactionID
	}
	return ""
}

// Generic CRUD Response
type TransactionResponse struct {
	Executed             bool                  `protobuf:"varint,1,opt,name=executed,proto3" json:"executed,omitempty"`
	Transactions         []*proto1.Transaction `protobuf:"bytes,2,rep,name=transactions,proto3" json:"transactions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *TransactionResponse) Reset()         { *m = TransactionResponse{} }
func (m *TransactionResponse) String() string { return proto.CompactTextString(m) }
func (*TransactionResponse) ProtoMessage()    {}
func (*TransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0975306caeaf870, []int{1}
}

func (m *TransactionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionResponse.Unmarshal(m, b)
}
func (m *TransactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionResponse.Marshal(b, m, deterministic)
}
func (m *TransactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionResponse.Merge(m, src)
}
func (m *TransactionResponse) XXX_Size() int {
	return xxx_messageInfo_TransactionResponse.Size(m)
}
func (m *TransactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionResponse proto.InternalMessageInfo

func (m *TransactionResponse) GetExecuted() bool {
	if m != nil {
		return m.Executed
	}
	return false
}

func (m *TransactionResponse) GetTransactions() []*proto1.Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func init() {
	proto.RegisterType((*TransactionQuery)(nil), "transaction.TransactionQuery")
	proto.RegisterType((*TransactionResponse)(nil), "transaction.TransactionResponse")
}

func init() {
	proto.RegisterFile("service/transaction/proto/transaction/transaction.proto", fileDescriptor_c0975306caeaf870)
}

var fileDescriptor_c0975306caeaf870 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x4d, 0x2a, 0xa5, 0x4e, 0x15, 0x75, 0x04, 0x59, 0x02, 0x4a, 0x08, 0x1e, 0x8a, 0x87,
	0x16, 0xea, 0x41, 0x10, 0x6f, 0x06, 0x24, 0x07, 0x0f, 0xc6, 0xaa, 0xe7, 0x98, 0xfe, 0x60, 0x41,
	0xb2, 0x61, 0x77, 0x2b, 0xf6, 0x59, 0x7c, 0x05, 0x1f, 0x52, 0xd8, 0x16, 0x9d, 0x3d, 0xb4, 0xa7,
	0xde, 0x76, 0xfe, 0x7f, 0xe6, 0x9b, 0x99, 0xdd, 0xa5, 0x6b, 0x0b, 0xf3, 0x39, 0xab, 0x31, 0x72,
	0xa6, 0x6a, 0x6c, 0x55, 0xbb, 0x99, 0x6e, 0x46, 0xad, 0xd1, 0x4e, 0x07, 0x8a, 0x38, 0x0f, 0xbd,
	0xcb, 0x7d, 0x21, 0x25, 0x97, 0x53, 0xb7, 0x68, 0x61, 0x37, 0x43, 0x96, 0x85, 0xd9, 0x4f, 0x44,
	0x47, 0x93, 0x7f, 0xf5, 0x71, 0x0e, 0xb3, 0xe0, 0x1b, 0x92, 0x3c, 0x15, 0xa5, 0xd1, 0xa0, 0x3f,
	0x56, 0x43, 0x59, 0x2d, 0x6a, 0x4a, 0x99, 0xcc, 0xe7, 0x44, 0x0f, 0x30, 0xf5, 0x7b, 0xd5, 0xb8,
	0x22, 0x57, 0x71, 0x1a, 0x0d, 0xf6, 0x4a, 0xa1, 0xf0, 0x29, 0x75, 0x9f, 0x2d, 0x4c, 0x91, 0xab,
	0x8e, 0xf7, 0x56, 0x11, 0x5f, 0xd0, 0x81, 0x60, 0x16, 0xb9, 0xda, 0xf5, 0x76, 0x28, 0x66, 0x9a,
	0x4e, 0x64, 0x67, 0xd8, 0x56, 0x37, 0x16, 0x9c, 0x50, 0x0f, 0x5f, 0xa8, 0xe7, 0x0e, 0x53, 0x3f,
	0x6d, 0xaf, 0xfc, 0x8b, 0xf9, 0x96, 0xf6, 0xc5, 0x7c, 0x56, 0xc5, 0x69, 0x67, 0xe3, 0x36, 0x41,
	0xf6, 0xf8, 0x3b, 0x26, 0x16, 0xee, 0xd3, 0xf2, 0x7d, 0xf8, 0x85, 0x8e, 0xef, 0x0c, 0x2a, 0x07,
	0xe1, 0xf1, 0xd9, 0x3a, 0xa6, 0xbf, 0xd5, 0x24, 0x5d, 0xdb, 0x72, 0xb5, 0x46, 0xb6, 0xc3, 0x13,
	0x3a, 0xbc, 0x87, 0x13, 0x9e, 0xdd, 0x06, 0xf5, 0x95, 0x38, 0xc7, 0x07, 0x82, 0x69, 0xb7, 0x01,
	0x7e, 0xeb, 0xfa, 0x4f, 0x74, 0xf5, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x5f, 0xc4, 0x38, 0xb8,
	0x02, 0x00, 0x00,
}
