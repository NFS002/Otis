package client

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"gitlab.com/otis-team/backend/dtypes/transaction/proto"
)

// CreateTransaction : Creates a new transaction in the db
func (c* DynamoClient) CreateTransaction(transaction *transaction.Transaction) (*transaction.Transaction, error) {
	av, err := dynamodbattribute.MarshalMap(transaction)
	if err != nil {
		return nil, err
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("Transaction"),
	}
	_, err = c.Client.PutItem(input)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

// GetAllTransactions : Returns all Transactions from the DB
func (c* DynamoClient) GetAllTransactions() ([]*transaction.Transaction, error) {
	result, err := c.Client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("Transaction"),
	})
	if err != nil {
		return nil, err
	}
	transactions := make([]*transaction.Transaction,0)
	err = dynamodbattribute.UnmarshalMap(result.Item,&transactions)
	return transactions, err
}

// GetTransactionByID : Returns the Transaction with the given ID from the DB
func (c* DynamoClient) GetTransactionByID(transactionID string) (*transaction.Transaction, error) {
	result, err := c.Client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("Transaction"),
		Key: map[string]*dynamodb.AttributeValue{
			"transaction_id": {
				S: aws.String(transactionID),
			},
		},
	})
	if err != nil {
		return nil, err
	}
	transaction := &transaction.Transaction{}
	err = dynamodbattribute.UnmarshalMap(result.Item,transaction)
	return transaction, err
}

// GetTransactionsByMerchantID : Returns all transactions made at the given merchant from the DB
func (c *DynamoClient) GetTransactionsByMerchantID(merchantID string) ([]*transaction.Transaction, error) {
	result, err := c.Client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("Transaction"),
		Key: map[string]*dynamodb.AttributeValue{
			"merchant_id": {
				S: aws.String(merchantID),
			},
		},
	})
	if err != nil {
		return nil, err
	}
	transactions := make([]*transaction.Transaction,0)
	err = dynamodbattribute.UnmarshalMap(result.Item,&transactions)
	return transactions, err
}

// GetTransactionsByUserID : Retrieves all transactions made by the given user from the DB
func (c *DynamoClient) GetTransactionsByUserID(userID string) ([]*transaction.Transaction, error) {
	result, err := c.Client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("Transaction"),
		Key: map[string]*dynamodb.AttributeValue{
			"user_id": {
				S: aws.String(userID),
			},
		},
	})
	if err != nil {
		return nil, err
	}
	transactions := make([]*transaction.Transaction,0)
	err = dynamodbattribute.UnmarshalMap(result.Item, &transactions)
	return transactions, err
}

// DeleteTransaction : Deletes the transaction with the given ID
func (c *DynamoClient) DeleteTransaction(transactionID string) error {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"transaction_id": {
				S: aws.String(transactionID),
			},
		},
		TableName: aws.String("User"),
	}
	_, err := c.Client.DeleteItem(input)
	return err
}