package transaction

import (
	"context"
	pb "gitlab.com/otis-team/backend/service/transaction/proto/transaction"
	"log"
)

type Handler struct {
	Repository
}

func (s *Handler) CreateTransaction(ctx context.Context, req *pb.Transaction, res *pb.CreateResponse) error {
	log.Print("CreateTransaction handler fired!")

	transaction, err := s.Repository.Create(ctx, MarshalTransaction(req))

	res.Created = true
	res.Transaction = UnmarshalTransaction(transaction)

	return err
}

func (s *Handler) GetTransactions(ctx context.Context, req *pb.IdRequest, res *pb.Transactions) error {
	log.Print("GetTransactions handler fired!")

	var transactions []*Transaction
	var err error

	if len(req.Id) == 0 {
		transactions, err = s.Repository.GetAll(ctx)
	} else {
		transactions, err = s.Repository.Get(ctx, req.Id)
	}

	res.Transactions = UnmarshalTransactionCollection(transactions)
	return err
}