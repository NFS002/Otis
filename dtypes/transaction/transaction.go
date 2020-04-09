package transaction

// Transaction : Struct is a local representation of a transaction
// Transaction resources retrieved from the DB are serialized to an instance of this struct
type Transaction struct {
	ID                   string           	`protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Timestamp            string           	`protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ProcessedTimestamp   string           	`protobuf:"bytes,3,opt,name=processed_timestamp,json=processedTimestamp,proto3" json:"processed_timestamp,omitempty"`
	UserID               string           	`protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AccountID            string           	`protobuf:"bytes,5,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	MerchantID           string           	`protobuf:"bytes,6,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
	Amount               float64          	`protobuf:"fixed64,7,opt,name=amount,proto3" json:"amount,omitempty"`
	Currency             string           	`protobuf:"bytes,8,opt,name=currency,proto3" json:"currency,omitempty"`
	Location             string           	`protobuf:"bytes,9,opt,name=location,proto3" json:"location,omitempty"`
	Rating               int32            	`protobuf:"varint,10,opt,name=rating,proto3" json:"rating,omitempty"`
	APITags              []string         	`protobuf:"bytes,11,rep,name=api_tags,json=apiTags,proto3" json:"api_tags,omitempty"`
	OtisTags             []string         	`protobuf:"bytes,12,rep,name=otis_tags,json=otisTags,proto3" json:"otis_tags,omitempty"`
	Holidays             []string         	`protobuf:"bytes,14,rep,name=holidays,proto3" json:"holidays,omitempty"`
	Type                 int32 		  		`protobuf:"varint,15,opt,name=type,proto3,enum=transaction.Transaction_Type" json:"type,omitempty"`
	Online               bool             	`protobuf:"varint,16,opt,name=online,proto3" json:"online,omitempty"`
	CashbackValid        bool             	`protobuf:"varint,17,opt,name=cashback_valid,json=cashbackValid,proto3" json:"cashback_valid,omitempty"`
	MerchantName         string           	`protobuf:"bytes,18,opt,name=merchant_name,json=merchantName,proto3" json:"merchant_name,omitempty"`
}

// Transactions represents a slice of Transaction structs
type Transactions []*Transaction