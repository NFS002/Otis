package merchant

import (
	"context"
	pb "gitlab.com/otis-team/backend/service/merchant/proto/merchant"
	"log"
)

// Handler struct contains other structs (mainly Repository, aka the client connection to the DB) to be used by Handler functions.
type Handler struct {
	Repository
}

// CreateMerchant handles gRPC requests to create a new merchant in the DB.
func (s *Handler) CreateMerchant(ctx context.Context, req *pb.Merchant, res *pb.CreateResponse) error {
	log.Print("CreateMerchant handler fired!")

	merchant, err := s.Repository.Create(ctx, MarshalMerchant(req))
	if err != nil {
		return err
	}

	res.Created = true
	res.Merchant = UnmarshalMerchant(merchant)

	return nil
}

// GetMerchant handles gRPC requests to retrieve one (if Merchant ID is upplied) or many merchants from the DB.
func (s *Handler) GetMerchant(ctx context.Context, req *pb.GetRequest, res *pb.GetResponse) error {
	log.Print("GetMerchant handler fired!")

	var merchants []*Merchant
	var err error

	if len(req.MerchantID) == 0 {
		merchants, err = s.Repository.GetAll(ctx)
	} else {
		merchants, err = s.Repository.Get(ctx, req.MerchantID)
	}

	res.Merchants = UnmarshalMerchantCollection(merchants)
	return err
}

// UpdateMerchant handles gRPC requests to update a new merchant in the DB
func (s *Handler) UpdateMerchant(ctx context.Context, req *pb.Merchant, res *pb.UpdateResponse) error {
	log.Print("UpdateMerchant handler fired!")

	err := s.Repository.Update(ctx, MarshalMerchant(req))
	if err != nil {
		return err
	}

	res.Updated = true
	res.Merchant = req

	return nil
}

// DeleteMerchant handles gRPC requests to delete a new merchant from the DB
func (s *Handler) DeleteMerchant(ctx context.Context, req *pb.DeleteRequest, res *pb.DeleteResponse) error {
	log.Print("DeleteMerchant handler fired!")

	err := s.Repository.Delete(ctx, req.MerchantID)
	if err != nil {
		return err
	}

	res.Deleted = true

	return nil
}