package client

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"gitlab.com/otis-team/backend/db/model"
)

// CreateTransaction : Creates a new transaction in the db
func (c* DynamoClient) CreateTransaction(transaction *model.Transaction) (*model.Transaction, error) {
	av, err := dynamodbav.Marshal(transaction)
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

// GetAllTransactiins : Returns all Transactions from the DB
func (c* DynamoClient) GetAllTransactions() ([]*model.Transaction, error) {
	result, err := c.Client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("Transaction")
	})
	if err != nil {
		return nil, err
	}
	transactions = model.Transactions{}
	err := dynamodbav.Unmarshal(result,&transactions)
	return transactions, err
}

// GetTransactionById : Returns the Transaction with the given ID from the DB
func (c* DynamoClient) GetTransactionById(transactionId string) (*model.Transactions, error) {
	result, err := c.Client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("Transaction"),
		Key: map[string]*dynamodb.AttributeValue{
			"transaction_id": {
				S: aws.String(transactionId),
			}
		},
	})
	if err != nil {
		return nil, err
	}
	transactions = model.Transactions{}
	err := dynamodbav.Unmarshal(result,&transactions)
	return transactions, err
}

// GetTransactionsByMerchantId : Returns all transactions made at the given merchant from the DB
func (c *DynamoClient) GetTransactionsByMerchantId(merchantId string) ([]*model.Transaction, error) {
	result, err := c.Client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("Transaction"),
		Key: map[string]*dynamodb.AttributeValue{
			"merchant_id": {
				S: aws.String(merchantId),
			}
		},
	})
	if err != nil {
		return nil, err
	}
	transactions = model.Transactions{}
	err := dynamodbav.Unmarshal(result,&transactions)
	return transactions, err
}

// GetTransactionsByUserId : Retrieves all transactions made by the given user from the DB
func (c *DynamoClient) GetTransactionsByUserId(userId string) ([]*model.Transaction, error) {
	result, err := c.Client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("Transaction"),
		Key: map[string]*dynamodb.AttributeValue{
			"user_id": {
				S: aws.String(userId),
			}
		},
	})
	if err != nil {
		return nil, err
	}
	transactions = model.Transactions{}
	err := dynamodbav.Unmarshal(result,&transactions)
	return transactions, err
}

// DeleteTransaction : Deletes the transaction with the given ID
func (c *DynamoClient) DeleteTransaction(transactionId string) error {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"transaction_id": {
				S: aws.String(transactionId),
			}
		},
		TableName: aws.String("User"),
	}
	_, err := c.Client.DeleteItem(input)
	return err
}