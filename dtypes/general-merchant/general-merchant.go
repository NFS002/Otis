package generalmerchant

// GeneralMerchant : Struct is a local representation of a merchant who is not part of the Otis platform
// GeneralMerchant resources retrieved from the DB are serialized to an instance of this struct
type GeneralMerchant struct {
	ID          string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Locations   []string `protobuf:"bytes,3,rep,name=locations,proto3" json:"locations,omitempty"`
	APITags     []string `protobuf:"bytes,4,rep,name=api_tags,json=apiTags,proto3" json:"api_tags,omitempty"`
	ScrapeTags  []string `protobuf:"bytes,5,rep,name=scrape_tags,json=scrapeTags,proto3" json:"scrape_tags,omitempty"`
	Sector      string   `protobuf:"bytes,6,opt,name=sector,proto3" json:"sector,omitempty"`
	Size        string   `protobuf:"bytes,7,opt,name=size,proto3" json:"size,omitempty"`
	ExpenseBand string   `protobuf:"bytes,12,opt,name=expenseBand,proto3" json:"expenseBand,omitempty"`
}

// GeneralMerchants represents a slice of GeneralMerchant structs
type GeneralMerchants []*GeneralMerchant