package client

import (
	"gitlab.com/otis_team/backend/dtypes/transaction/proto"
)

// CreateTransaction : Creates a new transaction in the db
func (c*RDSClient) CreateTransaction(transaction *transaction.Transaction) (*transaction.Transaction, error) {
	/*av, err := dynamodbattribute.MarshalMap(transaction)
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
	return transaction, nil8*/
	return nil, nil
}

// GetAllTransactions : Returns all Transactions from the DB
func (c*RDSClient) GetAllTransactions() ([]*transaction.Transaction, error) {
	/*result, err := c.Client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("Transaction"),
	})
	if err != nil {
		return nil, err
	}
	transactions := make([]*transaction.Transaction,0)
	err = dynamodbattribute.UnmarshalMap(result.Item,&transactions)
	return transactions, err*/
	return nil, nil
}

// GetTransactionByID : Returns the Transaction with the given ID from the DB
func (c*RDSClient) GetTransactionByID(transactionID string) (*transaction.Transaction, error) {
	/*result, err := c.Client.GetItem(&dynamodb.GetItemInput{
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
	return transaction, err*/
	return nil, nil
}

// GetTransactionsByMerchantID : Returns all transactions made at the given merchant from the DB
func (c *RDSClient) GetTransactionsByMerchantID(merchantID string) ([]*transaction.Transaction, error) {
	/*result, err := c.Client.GetItem(&dynamodb.GetItemInput{
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
	return transactions, err*/
	return nil, nil
}

// GetTransactionsByUserID : Retrieves all transactions made by the given user from the DB
func (c *RDSClient) GetTransactionsByUserID(userID string) ([]*transaction.Transaction, error) {
	/*result, err := c.Client.GetItem(&dynamodb.GetItemInput{
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
	return transactions, err*/
	return nil, nil
}

// DeleteTransaction : Deletes the transaction with the given ID
func (c *RDSClient) DeleteTransaction(transactionID string) error {
	/*input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"transaction_id": {
				S: aws.String(transactionID),
			},
		},
		TableName: aws.String("User"),
	}
	_, err := c.Client.DeleteItem(input)
	return err*/
	return nil
}
