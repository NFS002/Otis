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

func (s *Handler) CreateTransaction(ctx context.Context, req *pb.Transaction, res *pb.CRUDResponse) error {
	log.Print("CreateTransaction handler fired!")

	transaction, err := s.Repository.Create(ctx, MarshalTransaction(req))

	res.Created = true
	Collection := make([]*pb.Transaction, 0)
	Collection = append(Collection, UnmarshalTransaction(transaction))
	res.Transactions = Collection

	return err
}

// GetTransactions handles gRPC requests to retrieve one (if Transaction ID is supplied) or many transaction from the DB.
func (s *Handler) GetTransactions(ctx context.Context, req *pb.IDRequest, res *pb.CRUDResponse) error {
	log.Print("GetTransactions handler fired")

	var transactions []*Transaction
	var err error
	mID := req.MerchantID
	tID := req.TransactionID
	uID := req.UserID

	if tID != nil && len(req.Id) > 0 {
		transactions, err = s.Repository.GetTransactionByID(ctx, tID)
	} else if mID != nil && len(mID) > 0 {
		transactions, err = s.Repository.GetTransactionsByMerchantID(ctx, mID)
	} else if uID != nil && len(uID) > 0 {
		transactions, err = s.Repository.GetTransactionsByUserID(ctx, req.UID)
	}

	res.Transactions = UnmarshalTransactionCollection(transactions)
	return err
}

// Handles grpc requests to delete transactions in the DB
func (s *Handler) DeleteTransactions(ctx context.Context, req *pb.IdRequest, res *pb.CRUDResponse) error {
	log.Print("DeleteTransactions handler fired!")

	/* Not yet implemented */

	return nil
}