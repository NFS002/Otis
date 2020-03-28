package main

import (
	"context"
	"sync"
	"fmt"

	pb "gitlab.com/otis-team/backend/service/transaction/proto/transaction"
	"github.com/micro/go-micro"

)

const (
	port = ":60051"
)

type repository interface {
	Create(transaction *pb.Transaction) (*pb.Transaction, error)
	GetAll() []*pb.Transaction
	Get(id string) ([]*pb.Transaction, error)
}

type Repository struct {
	mu sync.RWMutex
	transactions []*pb.Transaction
}

func (repo *Repository) Create(transaction *pb.Transaction) (*pb.Transaction, error) {
	repo.mu.Lock()
	updated := append(repo.transactions, transactions)
	repo.transactions = updated
	repo.mu.Unlock()
	return transaction, nil
}

func (repo *Repository) GetAll() []*pb.Transaction {
	return repo.transactions
}

func (repo *Repository) Get(id string) ([]*pb.Transaction, error) {
	
	var res []*pb.Transaction
	for _, t := range repo.transactions {
		if t.TransactionID == id {
			res = append(res, t)
		}
	}

	return res, nil
}

type service struct {
	repo repository
}

func (s *service) CreateTransaction(ctx context.Context, req *pb.Transaction, res *pb.CreateResponse) error {
	transaction, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	res.Created = true
	res.Transaction = transaction
	res.Transaction.TransactionID = "5678"

	return nil
}

func (s *service) GetTransaction(ctx context.Context, req *pb.IdRequest, res *pb.GetResponse) error {
	if req.Id == "" {
		transactions := s.repo.GetAll()
	} else {
		transactions, err := s.repo.Get(req.Id)
		if err != nil {
			return err
		}
	}
	res.Transactions = transactions
	return nil
}

func main() {
	repo := &Repository{}

	srv := micro.NewService(
		micro.Name("otis-app.com.transaction.service"),
	)

	srv.Init()

	pb.RegisterTransactionServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}