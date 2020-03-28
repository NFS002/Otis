package merchant

import (
	"context"
	"github.com/satori/go.uuid"
	pb "gitlab.com/otis-team/backend/service/merchant/proto/merchant"
	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
	Collection := make([]*Location, 0)
	for _, location := range locations {
		Collection = append(Collection, MarshalLocation(location))
	}
	return Collection
}

func MarshalTagCollection(tags []*pb.Tag) []*Tag {
	Collection := make([]*Tag, 0)
	for _, tag := range tags {
		Collection = append(Collection, MarshalTag(tag))
	}
	return Collection
}

func UnmarshalLocationCollection(locations []*Location) []*pb.Location {
	Collection := make([]*pb.Location, 0)
	for _, location := range locations {
		Collection = append(Collection, UnmarshalLocation(location))
	}
	return Collection
}

func UnmarshalTagCollection(tags []*Tag) []*pb.Tag {
	Collection := make([]*pb.Tag, 0)
	for _, tag := range tags {
		Collection = append(Collection, UnmarshalTag(tag))
	}
	return Collection
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
	Collection := make([]*Merchant, 0)
	for _, merchant := range merchants {
		Collection = append(Collection, MarshalMerchant(merchant))
	}
	return Collection
}

func UnmarshalMerchantCollection(merchants []*Merchant) []*pb.Merchant {
	Collection := make([]*pb.Merchant, 0)
	for _, merchant := range merchants {
		Collection = append(Collection, UnmarshalMerchant(merchant))
	}
	return Collection
}

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

type Repository interface {
	Create(ctx context.Context, merchant *Merchant) (*Merchant, error)
	GetAll(ctx context.Context) ([]*Merchant, error)
	Get(ctx context.Context, merchantID string) ([]*Merchant, error)
	Update(ctx context.Context, merchant *Merchant) error
	Delete(ctx context.Context, merchantID string) error
}

type MongoRepository struct {
	Collection *mongo.Collection
}

func (repository *MongoRepository) Create(ctx context.Context, merchant *Merchant) (*Merchant, error){
	uuid, err := generateUUID()
	if err != nil {
		return nil, err
	}

	merchant.MerchantID = uuid.String()

	_, err = repository.Collection.InsertOne(ctx, merchant)

	return merchant, err
}

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

func (repository *MongoRepository) Update(ctx context.Context, merchant *Merchant) error {
	err := repository.Collection.FindOneAndReplace(ctx, bson.M{"merchant_id": merchant.MerchantID}, merchant)
	if err != nil {
		return err.Err()
	}

	return nil
}

func (repository *MongoRepository) Delete(ctx context.Context, merchantID string) error {
	err := repository.Collection.FindOneAndDelete(ctx, bson.M{"merchant_id": merchantID})
	if err != nil {
		return err.Err()
	}

	return nil
}

// UUID

func generateUUID() (uuid.UUID, error){
	var err error

	uuid := uuid.Must(uuid.NewV4(), err)

	return uuid, err
}