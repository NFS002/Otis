package partnermerchant

// PartnerMerchant : Struct is a local representation of a merchant who is registered with the Otis platform
// PartnerMerchant resources retrieved from the DB are serialized to an instance of this struct
type PartnerMerchant struct {
	ID                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Locations            []string `protobuf:"bytes,3,rep,name=locations,proto3" json:"locations,omitempty"`
	APITags              []string `protobuf:"bytes,4,rep,name=api_tags,json=apiTags,proto3" json:"api_tags,omitempty"`
	OtisTags             []string `protobuf:"bytes,5,rep,name=otis_tags,json=otisTags,proto3" json:"otis_tags,omitempty"`
	Sector               string   `protobuf:"bytes,6,opt,name=sector,proto3" json:"sector,omitempty"`
	Size                 string   `protobuf:"bytes,7,opt,name=size,proto3" json:"size,omitempty"`
	Rate                 int32    `protobuf:"varint,8,opt,name=rate,proto3" json:"rate,omitempty"`
	JoinDate             string   `protobuf:"bytes,9,opt,name=joinDate,proto3" json:"joinDate,omitempty"`
	NextBillingDate      string   `protobuf:"bytes,10,opt,name=nextBillingDate,proto3" json:"nextBillingDate,omitempty"`
	GocardlessMandate    string   `protobuf:"bytes,11,opt,name=gocardlessMandate,proto3" json:"gocardlessMandate,omitempty"`
	ExpenseBand          string   `protobuf:"bytes,12,opt,name=expenseBand,proto3" json:"expenseBand,omitempty"`
}

// PartnerMerchants represents a slice of PartnerMerchant structs
type PartnerMerchants []*PartnerMerchant