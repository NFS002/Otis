package client

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"gitlab.com/otis-team/backend/db/merchant"
)

// CreateMerchant : Creates a new merchant in the db
func (c* DynamoClient) CreateMerchant(merchant *model.Merchant) (*model.Merchant, error) {
	av, err := dynamodbav.Marshal(merchant)
	if err != nil {
		return nil, err
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("Merchant"),
	}
	_, err = c.Client.PutItem(input)
	if err != nil {
		return nil, err
	}
	return merchant, nil
}

// GetAllMerchants : Retrieves all merchants from the DB
func (c* DynamoClient) GetAllMerchants() (merchants []*model.Merchant, error) {
	result, err := c.Client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("Merchant"),
	})
	if err != nil {
		return nil, err
	}
	merchant = model.Merchant{}
	err := dynamodbav.Unmarshal(result,&merchant)
	return merchant, err
}

// GetMerchantById : Retrieves the Merchant from the DB with the given ID
func (c* DynamoClient) GetMerchantById(merchantId string) (merchant *model.Merchant, error) {
	result, err := c.Client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("Merchant"),
		Key: map[string]*dynamodb.AttributeValue{
			"merchant_id": {
				S: aws.String(merchantId),
			}
		},
	})
	if err != nil {
		return nil, err
	}
	merchant = model.Merchant{}
	err := dynamodbav.Unmarshal(result,&merchant)
	return merchant, err
}

// UpdateMerchant : Updates a merchant in the DB
func (c* DynamoClient) UpdateMerchant(merchant *model.Merchant) (*model.Merchant, error) {
	/* Deprecated. 
	* Call CreateMerchant(merchant *model.Merchant) (*model.Merchant, error) instead */
}

// Delete Merchant : Deletes a merchant with the given ID from the DB
func (c* DynamoClient) DeleteMerchant(merchantId string) (error) {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"merchant_id": {
				S: aws.String(merchantId),
			}
		},
		TableName: aws.String("Merchant"),
	}
	_, err := c.Client.DeleteItem(input)
	return err
}
