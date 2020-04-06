package main

import (
	"context"
	"gitlab.com/otis-team/backend/db/client"
	"gitlab.com/otis-team/backend/db/model"
	pb "gitlab.com/otis-team/backend/service/transaction/proto/transaction"
	"log"
)


// Handler struct contains the client connection to the DB, to be used by Handler functions.
type Handler struct {
	Client client.DynamoClient
}


// CreateTransaction : Persists the given transaction in the DB
func (h *Handler) CreateTransaction(ctx context.Context, req *pb.Transaction, res *pb.CRUDResponse) error {
	log.Print("CreateTransaction handler fired")
	transaction := model.ProtobufToTransaction(req)
	_, err := h.Client.CreateTransaction(transaction)
	res.Executed = (err == nil)
	res.Transactions = []*pb.Transaction{ req }
	return err
}

// GetTransactions handles gRPC requests to retrieve one (if Transaction ID is supplied) or many transaction from the DB.
func (h *Handler) GetTransactions(ctx context.Context, req *pb.IDRequest, res *pb.CRUDResponse) error {
	log.Print("GetTransactions handler fired")

	var transactions []*model.Transaction
	var err error

	tID := req.TransactionID
	mID := req.MerchantID
	uID := req.UserID

	if len(tID) > 0 {
		transactions, err = h.Client.GetTransactionByID(tID)
	} else if len(mID) > 0 {
		transactions, err = h.Client.GetTransactionsByMerchantID(mID)
	} else if len(uID) > 0 {
		transactions, err = h.Client.GetTransactionsByUserId(uID)
	} else {
		transactions, err = h.Client.GetAllTransactions()
	}

	res.Transactions = model.TransactionCollectionToProtobuf(transactions)
	res.Executed = true
	return err
}

// DeleteTransaction : Handles grpc requests to delete transactions in the DB
func (h *Handler) DeleteTransactions(ctx context.Context, req *pb.IDRequest, res *pb.CRUDResponse) error {
	log.Print("DeleteTransactions handler fired!")
	err := h.Client.DeleteTransaction(req.TransactionID)
	res.Executed = (err == nil)
	return err
}