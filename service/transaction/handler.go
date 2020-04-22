package main

import (
	"context"
	"gitlab.com/otis_team/backend/db/client"
	transaction "gitlab.com/otis_team/backend/dtypes/transaction/proto"
	transactionService "gitlab.com/otis_team/backend/service/transaction/proto/transaction"
	"log"
)


// Handler struct contains the client connection to the DB, to be used by Handler functions.
type Handler struct {
	Client client.RDSClient
}


// CreateTransaction : Persists the given transaction in the DB
func (h *Handler) CreateTransaction(ctx context.Context, req *transactionService.TransactionQuery, res *transactionService.TransactionResponse) error {
	log.Print("CreateTransaction handler fired")
	t, err := h.Client.CreateTransaction(req.Transaction)
	res.Executed = err == nil
	res.Transactions = []*transaction.Transaction{ t }
	return err
}

// GetTransactions handles gRPC requests to retrieve one (if Transaction ID is supplied) or many transaction from the DB.
func (h *Handler) GetTransactions(ctx context.Context, req *transactionService.TransactionQuery, res *transactionService.TransactionResponse) error {
	log.Print("GetTransactions handler fired")

	var transactions []*transaction.Transaction
	var transaction *transaction.Transaction
	var err error

	tID := req.TransactionID
	mID := req.MerchantID
	uID := req.UserID

	if len(tID) > 0 {
		transaction, err = h.Client.GetTransactionByID(tID)
		transactions = append(transactions, transaction)
	} else if len(mID) > 0 {
		transactions, err = h.Client.GetTransactionsByMerchantID(mID)
	} else if len(uID) > 0 {
		transactions, err = h.Client.GetTransactionsByUserID(uID)
	} else {
		transactions, err = h.Client.GetAllTransactions()
	}

	res.Transactions = transactions
	res.Executed = err == nil
	return err
}

// DeleteTransactions : Handles grpc requests to delete transactions in the DB
func (h *Handler) DeleteTransactions(ctx context.Context, req *transactionService.TransactionQuery, res *transactionService.TransactionResponse) error {
	log.Print("DeleteTransactions handler fired!")
	err := h.Client.DeleteTransaction(req.TransactionID)
	res.Executed = err == nil
	return err
}
