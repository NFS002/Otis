package model

import pb "gitlab.com/otis-team/backend/service/merchant/proto/merchant"

//Merchant : struct reprents a merchant db resource
type Merchant struct {
	MerchantID string `json:"merchant_id,omitempty" dybamodbav:"merchant_id"`
	Name string `json:"name" dybamodbav:"name"`
	Locations Locations `json:"locations" dybamodbav:"locations"`
	Tags Tags `json:"tags" dybamodbav:"tags"`
	ContactPhone string `json:"contact_phone" dybamodbav:"contact_phone"`
	ContactEmail string `json:"contact_email" dybamodbav:"contact_email"`
	Rate float32 `json:"rate" dybamodbav:"rate"`
}

// Location struct maps a tag protobuf message to a db resource
type Location struct {
	Location string `json:"location" dybamodbav:"location"`
}

// Locations struct represents slice of Location structs
type Locations []*Location

// Tag struct maps a tag protobuf message to a db resource
type Tag struct {
	Tag string `json:"tag" dybamodbav:"tag"`
}

// Tags struct represents slice of Tag structs
type Tags []*Tag

// ProtobufToMerchant : Converts a protobuf merchant message to db merchant resource
func ProtobufToMerchant(merchant *pb.Merchant) Merchant {
	locations := ProtobufToLocationCollection(merchant.Locations)
	tags := ProtobufToTagCollection(merchant.Tags)

	return &Merchant{
		MerchantID: merchant.MerchantID,
		Name: merchant.Name,
		Locations: locations,
		Tags: tags,
		ContactPhone: merchant.ContactPhone,
		ContactEmail: merchant.ContactEmail,
		Rate: merchant.Rate,
	}
}

// ProtobufToMerchantCollection : Converts a collection of protobuf merchant messages to db merchant resource
func ProtobufToMerchantCollection(merchant []*pb.Merchant) []*Merchant {
	Collection := make([]*Merchant, 0)
	for _, merchant := range merchants {
		Collection = append(Collection, ProtobufToMerchant(merchant))
	}
	return Collection
}

// ProtobufToLocation : Converts a protobuf location message to a db resource
func ProtobufToLocation(location *pb.Location) Location {
	return &Location{
		Location: location.Location,
	}
}

// ProtobufToLocationCollection : Converts a protobuf location message to a collection of db resources
func ProtobufToLocationCollection(location []*pb.Location) []*Location {
	Collection := make([]*Location, 0)
	for _, location := range locations {
		Collection = append(Collection, ProtobufToLocation(location))
	}
	return Collection
}

// ProtobufToTag : Converts a protobuf tag message to a db resource
func ProtobufToTag(tag *pb.Tag) Tag {
	return &Tag {
		Tag: tag.Tag,
	}
}

// ProtobufToTagCollection : Converts a protobuf tag message to a collection of db resources
func ProtobufToTagCollection(tag []*pb.Tag) []*Tag {
	Collection := make([]*Tag, 0)
	for _, tag := range tags {
		Collection = append(Collection, ProtobufToTag(tag))
	}
	return Collection
}

// MerchantToProtobuf : Converts a merchant db resource to a protobuf message
func MerchantToProtobuf(merchant Merchant) *pb.Merchant {
	locations := UnmarshalLocationCollection(merchant.Locations)
	tags := UnmarshalTagCollection(merchant.Tags)

	return &pb.Merchant{
		MerchantID: merchant.MerchantID,
		Name: merchant.Name,
		Locations: locations,
		Tags: tags,
		ContactPhone: merchant.ContactPhone,
		ContactEmail: merchant.ContactEmail,
		Rate: merchant.Rate,
	}
}

// MerchantCollectionToProtobuf : Converts a collection of merchant db resources to a protobuf message
func MerchantCollectionToProtobuf(merchants []*Merchant) []*pb.Merchant {
	Collection := make([]*pb.Merchant, 0)
	for _, merchant := range merchants {
		Collection = append(Collection, MerchantToProtobuf(merchant))
	}
	return Collection
}

// LocationToProtobuf : Converts a location db resource to a protobuf messsge
func LocationToProtobuf(location Location) *pb.Location  {
	return &pb.Location{
		Location: location.Location,
	}
}

// LocationCollectionToProtobuf : Converts a collection of location db resources to a protobuf message
func LocationCollectionToProtobuf(locations []*Location) []*pb.Location  {
	Collection := make([]*pb.Location, 0)
	for _, location := range locations {
		Collection = append(Collection, LocationToProtobuf(location))
	}
	return Collection
}

// TagToProtobuf : Converts a tag db resource to a protobuf message
func TagToProtobuf(tag Tag) *pb.Tag {
	return &pb.Tag{
		Tag: tag.Tag,
	}
}

// TagCollectionToProtobuf : Converts a collection 
func TagCollectionToProtobuf(tags []*Tag) []*pb.Tag {
	Collection := make([]*pb.Tag, 0)
	for _, tag := range tags {
		Collection = append(Collection, TagToProtobuf(tag))
	}
	return Collection
}