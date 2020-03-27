package merchant

import (
	"context"
	pb "gitlab.com/otis-team/backend/service/merchant/proto/merchant"
	"log"
)

type Handler struct {
	Repository
}

func (s *Handler) CreateMerchant(ctx context.Context, req *pb.Merchant, res *pb.CreateResponse) error {
	log.Print("CreateMerchant handler fired!")

	uuid, err := s.Repository.Create(ctx, MarshalMerchant(req))
	if err != nil {
		return err
	}

	res.Created = true
	res.MerchantID = uuid.String()

	return nil
}

func (s *Handler) GetMerchant(ctx context.Context, req *pb.GetRequest, res *pb.GetResponse) error {
	log.Print("GetMerchant handler fired!")

	var merchants []*Merchant
	var err error

	if len(req.Id) == 0 {
		merchants, err = s.Repository.GetAll(ctx)
	} else {
		merchants, err = s.Repository.Get(ctx, req.Id)
	}

	res.Merchants = UnmarshalMerchantCollection(merchants)
	return err
}

func (s *Handler) UpdateMerchant(ctx context.Context, req *pb.Merchant, res *pb.UpdateResponse) error {
	log.Print("UpdateMerchant handler fired!")

	err := s.Repository.Update(ctx, MarshalMerchant(req))
	if err != nil {
		return err
	}

	res.Created = true
	res.Merchant = req

	return nil
}

func (s *Handler) DeleteMerchant(ctx context.Context, req *pb.DeleteRequest, res *pb.DeleteResponse) error {
	log.Print("DeleteMerchant handler fired!")

	err := s.Repository.Delete(ctx, req.MerchantID)
	if err != nil {
		return err
	}

	res.Deleted = true

	return nil
}