package main

import (
	"context"
	"fmt"
	pb "gitlab.com/otis-team/backend/service/merchant/proto/merchant"
)

type handler struct {
	repository
}

func (s *handler) CreateMerchant(ctx context.Context, req *pb.Merchant, res *pb.CreateResponse) error {
	fmt.Println("Handler called!")

	uuid, err := s.repository.Create(ctx, MarshalMerchant(req))
	if err != nil {
		return err
	}

	res.Created = true
	res.MerchantID = uuid.String()

	return nil
}

func (s *handler) GetMerchant(ctx context.Context, req *pb.GetRequest, res *pb.GetResponse) error {
	var merchants []*Merchant
	var err error

	if len(req.Id) == 0 {
		merchants, err = s.repository.GetAll(ctx)
	} else {
		merchants, err = s.repository.Get(ctx, req.Id)
	}

	res.Merchants = UnmarshalMerchantCollection(merchants)
	return err
}