package campaign

// Campaign : Struct is a local representation of a campaign resource
// Campaign resources retrieved fron the DB are serialized to an instance of this struct
type Campaign struct {
	ID                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	MerchantID           string   `protobuf:"bytes,2,opt,name=merchantID,proto3" json:"merchantID,omitempty"`
	Rate                 string   `protobuf:"bytes,3,opt,name=rate,proto3" json:"rate,omitempty"`
	StartTimestamp       string   `protobuf:"bytes,4,opt,name=startTimestamp,proto3" json:"startTimestamp,omitempty"`
	EndTimestamp         string   `protobuf:"bytes,5,opt,name=endTimestamp,proto3" json:"endTimestamp,omitempty"`
	UserSet              []string `protobuf:"bytes,6,rep,name=userSet,proto3" json:"userSet,omitempty"`
	TransactionSet       []string `protobuf:"bytes,7,rep,name=transactionSet,proto3" json:"transactionSet,omitempty"`
	Title                string   `protobuf:"bytes,8,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,9,opt,name=description,proto3" json:"description,omitempty"`
	ValidFor             int32    `protobuf:"varint,10,opt,name=validFor,proto3" json:"validFor,omitempty"`
	Active               bool     `protobuf:"varint,11,opt,name=active,proto3" json:"active,omitempty"`
}