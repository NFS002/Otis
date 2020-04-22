package transaction

import (               
	"context"
	"github.com/satori/go.uuid"
	pb "gitlab.com/otis_team/backend/service/transaction/proto/transaction"
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
		TransactionID: transaction.TransactionID,
		UserID: transaction.UserID,
		MerchantID: transaction.MerchantID,
		StoreID: transaction.StoreID,
		Date: transaction.Date,
		Time: transaction.Time,
	}
}

// UnmarshalTransaction converts transaction struct to transaction protobuf
func UnmarshalTransaction(transaction *Transaction) *pb.Transaction {
	return &pb.Transaction {
		TransactionID: transaction.TransactionID,
		UserID: transaction.UserID,
		MerchantID: transaction.MerchantID,
		StoreID: transaction.StoreID,
		Date: transaction.Date,
		Time: transaction.Time,
	}
}


//Repository : Interface describes all available repository methods and is an abstraction over a Mongo collection */
type Repository interface {
	Create(ctx context.Context, transaction *Transaction) (*Transaction, error)
	GetAllTransactions(ctx context.Context) ([]*Transaction, error)
	GetTransactionByID(ctx context.Context, transactionID string) ([]*Transaction, error)
	GetTransactionsByUserID(ctx context.Context, userID string) ([]*Transaction, error)
	GetTransactionsByMerchantID(ctx context.Context, merchantID string) ([]*Transaction, error)
	GetTransactionsFromBSONQuery(ctx context.Context, query bson.M) ([]*Transaction, error)

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

// GetTransactionsFromBSONQuery : Method implements functionality to retrieve all transactions based on a BSON query
func (repository *MongoRepository) GetTransactionsFromBSONQuery(ctx context.Context, query bson.M) ([]*Transaction, error) {
	cur, err := repository.Collection.Find(ctx, query, nil)
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

// GetAllTransactions : Implements functionality to retrieve all transactions from the DB
func (repository *MongoRepository) GetAllTransactions(ctx context.Context) ([]*Transaction, error) {
	query := bson.M{}
	return repository.GetTransactionsFromBSONQuery(ctx, query)
}

// GetTransactionByID : Retrieves a single transaction from the DB
func (repository *MongoRepository) GetTransactionByID(ctx context.Context, transactionID string) ([]*Transaction, error) {
	query := bson.M{"transaction_id": transactionID }
	return repository.GetTransactionsFromBSONQuery(ctx, query)
}

// GetTransactionsByMerchantID : Retrieves all transactions from a single merchant
func (repository *MongoRepository) GetTransactionsByMerchantID(ctx context.Context, merchantID string) ([]*Transaction, error) {
	query := bson.M{"merchant_id": merchantID }
	return repository.GetTransactionsFromBSONQuery(ctx, query)
}

// GetTransactionsByUserID : Retrieves all transactions from a single user
func (repository *MongoRepository) GetTransactionsByUserID(ctx context.Context, userID string) ([]*Transaction, error) {
	query := bson.M{"user_id": userID }
	return repository.GetTransactionsFromBSONQuery(ctx, query)
}


func generateID() (uuid.UUID, error){
	var err error

	uuid := uuid.Must(uuid.NewV4(), err)

	return uuid, err
}
