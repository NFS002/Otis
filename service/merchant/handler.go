package main

import (
	"log"
	"github.com/micro/go-micro"
	client "gitlab.com/otis-team/backend/db/client"
	model "gitlab.com/otis-team/backend/db/merchant"
	pb "gitlab.com/otis-team/backend/service/merchant/proto/merchant"
)


// Handler struct contains the client connection to the DB, to be used by Handler functions.
type Handler struct {
	Client client.DynamoClient
}


// CreateMerchant handles gRPC requests to create a new merchant in the DB.
func (h *Handler) CreateMerchant(ctx context.Context, req *pb.Merchant, res *pb.CreateResponse) error {
	log.Print("CreateMerchant handler fired!")
	merchant := model.ProtobufToMerchant(req)
	_, err := s.Client.Create(merchant)
	res.Created = (err == nil)
	res.Merchant = req
	return err
}

// GetMerchant handles gRPC requests to retrieve one (if Merchant ID is supplied) or many merchants from the DB.
func (h *Handler) GetMerchant(ctx context.Context, req *pb.GetRequest, res *pb.GetResponse) error {
	log.Print("GetMerchant handler fired!")

	var merchants []*Merchant
	var err error

	if len(req.MerchantID) == 0 {
		merchants, err = h.Client.GetAllMerchants()
	} else {
		merchants, err = h.Client.GetMerchantById(req.MerchantID)
	}

	res.Merchants = model.MerchantCollectionToProtobuf(merchants)
	return err
}

// UpdateMerchant handles gRPC requests to update a new merchant in the DB
func (h *Handler) UpdateMerchant(ctx context.Context, req *pb.Merchant, res *pb.UpdateResponse) error {
	log.Print("UpdateMerchant handler fired!")
	merchant := model.ProtobufToMerchant(req)
	/* UpdateMerchant behaves the same as CreateMerchant if the primary key is given */
	err := s.Client.CreateMerchant(merchant)
	res.Updated = (err == nil)
	res.Merchant = req
	return err
}

// DeleteMerchant handles gRPC requests to delete a new merchant from the DB
func (h *Handler) DeleteMerchant(ctx context.Context, req *pb.DeleteRequest, res *pb.DeleteResponse) error {
	log.Print("DeleteMerchant handler fired!")
	err := h.Client.Delete(req.MerchantID)
	res.Deleted = (err == nil)
	return err
}