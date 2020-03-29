package transaction

import (               
	"context"
	"github.com/satori/go.uuid"
	pb "gitlab.com/otis-team/backend/service/transaction/proto/transaction"
	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Transaction struct maps protbuf definition. Contains json and bson key mappings
type Transaction struct {
	TransactionID string `json:"transaction_id,omitempty" bson:"transaction_id"`
	UserID string `json:"user_id,omitempty" bson:"user_id"`
	MerchantID string `json:"merchant_id" bson:"merchant_id"`
	StoreID string `json:"store_id" bson:"store_id"`
	Date string `json:"date" bson:"date"`
	Time string `json:"time" bson:"time"`
}

// MarshalTransactionCollection converts slice of transaction protobufs to slice of transaction structs.
func MarshalTransactionCollection(transactions []*pb.Transaction) []*Transaction {
	Collection := make([]*Transaction, 0)
	for _, t := range transactions {
		Collection = append(Collection, MarshalTransaction(t))
	}
	return Collection
}

// UnmarshalTransactionCollection converts slice of transaction structs to slice of transaction protobufs
func UnmarshalTransactionCollection(transactions []*Transaction) []*pb.Transaction {
	Collection := make([]*pb.Transaction, 0)
	for _, t := range transactions {
		Collection = append(Collection, UnmarshalTransaction(t))
	}
	return Collection
}

// MarshalTransaction converts a transaction protobuf to transaction struct
func MarshalTransaction(transaction *pb.Transaction) *Transaction {
	return &Transaction{
		TransactionID: transaction.TransactionId,
		UserID: transaction.UserId,
		MerchantID: transaction.MerchantId,
		StoreID: transaction.StoreId,
		Date: transaction.Date,
		Time: transaction.Time,
	}
}

// UnmarshalTransaction converts transaction struct to transaction protobuf
func UnmarshalTransaction(transaction *Transaction) *pb.Transaction {
	return &pb.Transaction {
		TransactionId: transaction.TransactionID,
		UserId: transaction.UserID,
		MerchantId: transaction.MerchantID,
		StoreId: transaction.StoreID,
		Date: transaction.Date,
		Time: transaction.Time,
	}
}

// Repository

// Repository interface describes all available repository methods. Currently Create and retrieve
type Repository interface {
	Create(ctx context.Context, transaction *Transaction) (*Transaction, error)
	GetAll(ctx context.Context) ([]*Transaction, error)
	Get(ctx context.Context, transactionID string) ([]*Transaction, error)
}

// MongoRepository struct describes specif collection relevant to the repository being used
type MongoRepository struct {
	Collection *mongo.Collection
}

// Create method implements functionlaity to create a transaction in the DB. UUID is generate to fill transaction_id.
func (repository *MongoRepository) Create(ctx context.Context, transaction *Transaction) (*Transaction, error){
	uuid, err := generateID()
	if err != nil {
		return nil, err
	}

	transaction.TransactionID = uuid.String()

	_, err = repository.Collection.InsertOne(ctx, transaction)

	return transaction, err
}

// GetAll method implements functionality to retrieve all transactions from the DB
func (repository *MongoRepository) GetAll(ctx context.Context) ([]*Transaction, error) {
	cur, err := repository.Collection.Find(ctx, bson.D{}, nil)
	if err != nil {
		return nil, err
	}

	var transactions []*Transaction
	for cur.Next(ctx) {
		var transaction *Transaction
		if err := cur.Decode(&transaction); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, err
}

// Get method implements functionality to retrieve a single transaction from the DB
func (repository *MongoRepository) Get(ctx context.Context, transactionID string) ([]*Transaction, error) {
	cur, err := repository.Collection.Find(ctx, bson.M{"transaction_id": transactionID}, nil)
	var transactions []*Transaction
	for cur.Next(ctx) {
		var transaction *Transaction
		if err := cur.Decode(&transaction); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, err
}

// UUID

func generateID() (uuid.UUID, error){
	var err error

	uuid := uuid.Must(uuid.NewV4(), err)

	return uuid, err
}