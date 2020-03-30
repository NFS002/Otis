package transaction

import (
	"context"
	pb "gitlab.com/otis-team/backend/service/transaction/proto/transaction"
	"log"
)

// Handler struct contains other structs (mainly Repository, aka the client connection to the DB) to be used by Handler functions.
type Handler struct {
	Repository
}

// CreateTransaction : persists the given transaction in the db
func (s *Handler) CreateTransaction(ctx context.Context, req *pb.Transaction, res *pb.CRUDResponse) error {
	log.Print("CreateTransaction handler fired!")

	transaction, err := s.Repository.Create(ctx, MarshalTransaction(req))

	Collection := make([]*pb.Transaction, 0)
	Collection = append(Collection, UnmarshalTransaction(transaction))
	res.Transactions = Collection
	res.Executed = true

	return err
}

// GetTransactions handles gRPC requests to retrieve one (if Transaction ID is supplied) or many transaction from the DB.
func (s *Handler) GetTransactions(ctx context.Context, req *pb.IDRequest, res *pb.CRUDResponse) error {
	log.Print("GetTransactions handler fired")

	var transactions []*Transaction
	var err error

	tID := req.TransactionID
	mID := req.MerchantID
	uID := req.UserID

	if len(tID) > 0 {
		transactions, err = s.Repository.GetTransactionByID(ctx, tID)
	} else if len(mID) > 0 {
		transactions, err = s.Repository.GetTransactionsByMerchantID(ctx, mID)
	} else if len(uID) > 0 {
		transactions, err = s.Repository.GetTransactionsByUserID(ctx, uID)
	}

	res.Transactions = UnmarshalTransactionCollection(transactions)
	res.Executed = true
	return err
}

// DeleteTransactions : Handles grpc requests to delete transactions in the DB
func (s *Handler) DeleteTransactions(ctx context.Context, req *pb.IDRequest, res *pb.CRUDResponse) error {
	log.Print("DeleteTransactions handler fired!")

	/* Not yet implemented */

	return nil
}