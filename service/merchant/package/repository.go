package merchant

import (
	"context"
	"github.com/satori/go.uuid"
	pb "gitlab.com/otis_team/backend/service/merchant/proto/merchant"
	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Merchant struct maps protobuf definition. Contains json and bson key mappings
type Merchant struct {
	MerchantID string `json:"merchant_id,omitempty" bson:"merchant_id"`
	Name string `json:"name" bson:"name"`
	Locations Locations `json:"locations" bson:"locations"`
	Tags Tags `json:"tags" bson:"tags"`
	ContactPhone string `json:"contact_phone" bson:"contact_phone"`
	ContactEmail string `json:"contact_email" bson:"contact_email"`
	Rate float32 `json:"rate" bson:"rate"`
}

// Location struct maps protobuf definition. Contains json and bson key mappings
type Location struct {
	Location string `json:"location" bson:"location"`
}

// Locations struct represents slice of Location structs
type Locations []*Location

// Tag struct maps protobuf definition. Contains json and bson key mappings
type Tag struct {
	Tag string `json:"tag" bson:"tag"`
}

// Tags struct represents slice of Tag structs
type Tags []*Tag


// MarshalLocationCollection converts slice of location protobufs into slice of location structs
func MarshalLocationCollection(locations []*pb.Location) []*Location {
	Collection := make([]*Location, 0)
	for _, location := range locations {
		Collection = append(Collection, MarshalLocation(location))
	}
	return Collection
}

// MarshalTagCollection converts slice of tag protobufs into slice of tag structs
func MarshalTagCollection(tags []*pb.Tag) []*Tag {
	Collection := make([]*Tag, 0)
	for _, tag := range tags {
		Collection = append(Collection, MarshalTag(tag))
	}
	return Collection
}

// UnmarshalLocationCollection converts slice of location structs into slice of location protobufs
func UnmarshalLocationCollection(locations []*Location) []*pb.Location {
	Collection := make([]*pb.Location, 0)
	for _, location := range locations {
		Collection = append(Collection, UnmarshalLocation(location))
	}
	return Collection
}

// UnmarshalTagCollection converts slice of tag structs into slice of tag protobufs
func UnmarshalTagCollection(tags []*Tag) []*pb.Tag {
	Collection := make([]*pb.Tag, 0)
	for _, tag := range tags {
		Collection = append(Collection, UnmarshalTag(tag))
	}
	return Collection
}

// MarshalLocation converts location protobuf into location struct
func MarshalLocation(location *pb.Location) *Location {
	return &Location{
		Location: location.Location,
	}
}

// MarshalTag converts tag protobuf into tag struct
func MarshalTag(tag *pb.Tag) *Tag {
	return &Tag {
		Tag: tag.Tag,
	}
}

// UnmarshalLocation converts location struct into location protobuf
func UnmarshalLocation(location *Location) *pb.Location {
	return &pb.Location{
		Location: location.Location,
	}
}

// UnmarshalTag converts tag struct into tag protobuf
func UnmarshalTag(tag *Tag) *pb.Tag {
	return &pb.Tag{
		Tag: tag.Tag,
	}
}


// Merchant


// MarshalMerchantCollection converts slice of merhcnat protobufs into slice of merchant structs
func MarshalMerchantCollection(merchants []*pb.Merchant) []*Merchant {
	Collection := make([]*Merchant, 0)
	for _, merchant := range merchants {
		Collection = append(Collection, MarshalMerchant(merchant))
	}
	return Collection
}

// UnmarshalMerchantCollection converts slice of merchant structs into slice of merchant protobufs
func UnmarshalMerchantCollection(merchants []*Merchant) []*pb.Merchant {
	Collection := make([]*pb.Merchant, 0)
	for _, merchant := range merchants {
		Collection = append(Collection, UnmarshalMerchant(merchant))
	}
	return Collection
}

// MarshalMerchant converts merchant protobuf into merchant struct
func MarshalMerchant(merchant *pb.Merchant) *Merchant {
	locations := MarshalLocationCollection(merchant.Locations)
	tags := MarshalTagCollection(merchant.Tags)

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

// UnmarshalMerchant converts merchant struct into merchant protobuf
func UnmarshalMerchant(merchant *Merchant) *pb.Merchant {
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

// Repository


// Repository interface describes all available repository methods. Currently basic CRUD.
type Repository interface {
	Create(ctx context.Context, merchant *Merchant) (*Merchant, error)
	GetAll(ctx context.Context) ([]*Merchant, error)
	Get(ctx context.Context, merchantID string) ([]*Merchant, error)
	Update(ctx context.Context, merchant *Merchant) error
	Delete(ctx context.Context, merchantID string) error
}

// MongoRepository struct describes specific collection relevant to the repository being used.
type MongoRepository struct {
	Collection *mongo.Collection
}

// Create method implements functionality to create a merchant in the DB. UUID is generated to fill merchant_id.
func (repository *MongoRepository) Create(ctx context.Context, merchant *Merchant) (*Merchant, error){
	uuid, err := generateUUID()
	if err != nil {
		return nil, err
	}

	merchant.MerchantID = uuid.String()

	_, err = repository.Collection.InsertOne(ctx, merchant)

	return merchant, err
}

// GetAll method implements functionality to retrieve all merchants from the DB.
func (repository *MongoRepository) GetAll(ctx context.Context) ([]*Merchant, error) {
	cur, err := repository.Collection.Find(ctx, bson.D{}, nil)
	var merchants []*Merchant
	for cur.Next(ctx) {
		var merchant *Merchant
		if err := cur.Decode(&merchant); err != nil {
			return nil, err
		}
		merchants = append(merchants, merchant)
	}

	return merchants, err
}

// Get method implements functionality to retrieve merchants from the DB matching merchantID. Possible for multiple
// merchants to be returned.
func (repository *MongoRepository) Get(ctx context.Context, merchantID string) ([]*Merchant, error) {
	cur, err := repository.Collection.Find(ctx, bson.M{"merchant_id": merchantID}, nil)
	var merchants []*Merchant
	for cur.Next(ctx) {
		var merchant *Merchant
		if err := cur.Decode(&merchant); err != nil {
			return nil, err
		}
		merchants = append(merchants, merchant)
	}

	return merchants, err
}

// Update method implements functionality to update a merchant in the DB. Only requirement for supplied new merchant is
// that the merchant_id field matches an existing merchant in the DB.
func (repository *MongoRepository) Update(ctx context.Context, merchant *Merchant) error {
	err := repository.Collection.FindOneAndReplace(ctx, bson.M{"merchant_id": merchant.MerchantID}, merchant)
	if err != nil {
		return err.Err()
	}

	return nil
}

// Delete method implements functionality to delete a merchant from the DB matching a user ID.
func (repository *MongoRepository) Delete(ctx context.Context, merchantID string) error {
	err := repository.Collection.FindOneAndDelete(ctx, bson.M{"merchant_id": merchantID})
	if err != nil {
		return err.Err()
	}

	return nil
}

// UUID

// generateUUID generates a random UUID
func generateUUID() (uuid.UUID, error){
	var err error

	uuid := uuid.Must(uuid.NewV4(), err)

	return uuid, err
}
