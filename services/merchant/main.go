package main

import (
	"context"
	"sync"
	"fmt"

	pb "gitlab.com/otis-team/backend/services/merchant/proto/merchant"
	"github.com/micro/go-micro"

)

const (
	port = ":50051"
)

type repository interface {
	Create(merchant *pb.Merchant) (*pb.Merchant, error)
	GetAll() []*pb.Merchant
	Get(id string) ([]*pb.Merchant, error)
}

type Repository struct {
	mu sync.RWMutex
	merchants []*pb.Merchant
}

func (repo *Repository) Create(merchant *pb.Merchant) (*pb.Merchant, error) {
	repo.mu.Lock()
	updated := append(repo.merchants, merchant)
	repo.merchants = updated
	repo.mu.Unlock()
	return merchant, nil
}

func (repo *Repository) GetAll() []*pb.Merchant {
	return repo.merchants
}

func (repo *Repository) Get(id string) ([]*pb.Merchant, error) {
	var res []*pb.Merchant

	for _, merchant := range repo.merchants {
		if merchant.Id == id {
			res = append(res, merchant)
		}
	}

	return res, nil
}

type service struct {
	repo repository
}

func (s *service) CreateMerchant(ctx context.Context, req *pb.Merchant, res *pb.CreateResponse) error {

	merchant, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	res.Created = true
	res.Merchant = merchant
	res.Merchant.Id = "1234"
	return nil
}

// REFACTOR
func (s *service) GetMerchant(ctx context.Context, req *pb.GetRequest, res *pb.GetResponse) error {
	if req.Id != "" {
		merchants, err := s.repo.Get(req.Id)
		if err != nil {
			return err
		}
		res.Merchants = merchants
	} else {
		merchants := s.repo.GetAll()

		res.Merchants = merchants
	}

	return nil
}

func main() {
	repo := &Repository{}

	srv := micro.NewService(
			micro.Name("otis-app.com.merchant.service"),
		)

	srv.Init()

	pb.RegisterMerchantServiceHandler(srv.Server(), &service{repo})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}