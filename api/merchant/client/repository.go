package repository

import (
	pb "gitlab.com/otis-team/backend/api/merchant/client/proto"
)

type Merchant struct {
	MerchantID string `json:"merchant_id,omitempty" bson:"merchant_id"`
	Name string `json:"name" bson:"name"`
	Locations Locations `json:"locations" bson:"locations"`
	Tags Tags `json:"tags" bson:"tags"`
	ContactPhone string `json:"contact_phone" bson:"contact_phone"`
	ContactEmail string `json:"contact_email" bson:"contact_email"`
	Rate float32 `json:"rate" bson:"rate"`
}

type Location struct {
	Location string `json:"location" bson:"location"`
}

type Locations []*Location

type Tag struct {
	Tag string `json:"tag" bson:"tag"`
}

type Tags []*Tag

// Location and Tags

func MarshalLocationCollection(locations []*pb.Location) []*Location {
	collection := make([]*Location, 0)
	for _, location := range locations {
		collection = append(collection, MarshalLocation(location))
	}
	return collection
}

func MarshalTagCollection(tags []*pb.Tag) []*Tag {
	collection := make([]*Tag, 0)
	for _, tag := range tags {
		collection = append(collection, MarshalTag(tag))
	}
	return collection
}

func UnmarshalLocationCollection(locations []*Location) []*pb.Location {
	collection := make([]*pb.Location, 0)
	for _, location := range locations {
		collection = append(collection, UnmarshalLocation(location))
	}
	return collection
}

func UnmarshalTagCollection(tags []*Tag) []*pb.Tag {
	collection := make([]*pb.Tag, 0)
	for _, tag := range tags {
		collection = append(collection, UnmarshalTag(tag))
	}
	return collection
}

func MarshalLocation(location *pb.Location) *Location {
	return &Location{
		Location: location.Location,
	}
}

func MarshalTag(tag *pb.Tag) *Tag {
	return &Tag {
		Tag: tag.Tag,
	}
}

func UnmarshalLocation(location *Location) *pb.Location {
	return &pb.Location{
		Location: location.Location,
	}
}

func UnmarshalTag(tag *Tag) *pb.Tag {
	return &pb.Tag{
		Tag: tag.Tag,
	}
}

// Merchant

func MarshalMerchantCollection(merchants []*pb.Merchant) []*Merchant {
	collection := make([]*Merchant, 0)
	for _, merchant := range merchants {
		collection = append(collection, MarshalMerchant(merchant))
	}
	return collection
}

func UnmarshalMerchantCollection(merchants []*Merchant) []*pb.Merchant {
	collection := make([]*pb.Merchant, 0)
	for _, merchant := range merchants {
		collection = append(collection, UnmarshalMerchant(merchant))
	}
	return collection
}

func MarshalMerchant(merchant *pb.Merchant) *Merchant {
	locations := MarshalLocationCollection(merchant.Locations)
	tags := MarshalTagCollection(merchant.Tags)

	return &Merchant{
		MerchantID: merchant.Id,
		Name: merchant.Name,
		Locations: locations,
		Tags: tags,
		ContactPhone: merchant.ContactPhone,
		ContactEmail: merchant.ContactEmail,
		Rate: merchant.Rate,
	}
}

func UnmarshalMerchant(merchant *Merchant) *pb.Merchant {
	locations := UnmarshalLocationCollection(merchant.Locations)
	tags := UnmarshalTagCollection(merchant.Tags)

	return &pb.Merchant{
		Id: merchant.MerchantID,
		Name: merchant.Name,
		Locations: locations,
		Tags: tags,
		ContactPhone: merchant.ContactPhone,
		ContactEmail: merchant.ContactEmail,
		Rate: merchant.Rate,
	}
}