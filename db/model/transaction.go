package model

import pb "gitlab.com/otis-team/backend/service/user/proto/transaction"

// Transaction struct represents a dynamodb transaction resource
type Transaction struct {
	TransactionID string `json:"transaction_id,omitempty" dynamodbav:"transaction_id,omitempty"`
	UserID string `json:"user_id,omitempty" dynamodbav:"user_id"`
	MerchantID string `json:"merchant_id" dynamodbav:"merchant_id"`
	StoreID string `json:"store_id" dynamodbav:"store_id"`
	Date string `json:"date" dynamodbav:"date"`
	Time string `json:"time" dynamodbav:"time"`
}

// Transactions struct represents a slice of User structs
type Transactions []*Transaction

// ProtobufToTransaction : Converts a protobuf transaction message to Transaction struct
func ProtobufToTransaction(transaction *pb.Transaction) Transaction  {
	return &Transaction{
		TransactionID: 	transaction.TransactionID
		UserID:     	transaction.UserID,
		MerchantID: 	transaction.MerchantID,
		StoreID:  		transaction.StoreID,
		Date:       	transaction.Date,
		Time:     		transaction.Time,
	}
}

// ProtobufToTransactionCollection : Converts a protobuf message to a collection of Transaction structs
func ProtobufToTransactionCollection(transactions []*pb.Transaction) []*Transaction  {
	Collection := make([]*Transaction, 0)
	for _, transaction := range users {
		Collection = append(Collection, ProtobufToTransaction(transaction))
	}
	return Collection
}

// TransactionToProtobuf : Converts a Transaction struct into a protobuf message
func TransactionToProtobuf(transaction Transaction) *pb.Transaction {
	return &pb.Transaction{
		TransactionID:	transaction.TransactionID
		UserID:     	transaction.UserID,
		MerchantID: 	transaction.MerchantID,
		StoreID:  		transaction.StoreID,
		Date:       	transaction.Date,
		Time:     		transaction.Time,
	}
}

// TransactionCollectionToProtobuf : Converts a collection of 
func TransactionCollectionToProtobuf(transactions []*pb.Transaction) []*Transaction {
	Collection := make([]*pb.Transaction, 0)
	for _, transaction := range users {
		Collection = append(Collection, TransactionToProtobuf(transaction))
	}
	return Collection
}