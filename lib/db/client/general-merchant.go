package client

import (
	//"github.com/aws/aws-sdk-go/service/dynamodb"
	//"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"gitlab.com/otis_team/backend/dtypes/general-merchant/proto"
)


// CreateGeneralMerchant : Creates a new general merchant in the db
func (c*RDSClient) CreateGeneralMerchant(merchant *generalmerchant.GeneralMerchant) (*generalmerchant.GeneralMerchant, error) {
	/*av, err := dynamodbattribute.MarshalMap(merchant)
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
	Not implemented */
	return nil, nil
}

// GetAllGeneralMerchants : Retrieves all general merchants from the DB
func (c*RDSClient) GetAllGeneralMerchants() ([]*generalmerchant.GeneralMerchant, error) {
	/*result, err := c.Client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("Merchant"),
	})
	if err != nil {
		return nil, err
	}
	merchants := make([]*generalmerchant.GeneralMerchant, 0)
	err = dynamodbattribute.UnmarshalMap(result.Item,&merchants)
	return merchants, err
	Not implemented */
	return nil, nil
}

// GetGeneralMerchantByID : Retrieves the general merchant from the DB with the given ID
func (c*RDSClient) GetGeneralMerchantByID(merchantID string) (*generalmerchant.GeneralMerchant, error) {
	/*result, err := c.Client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("Merchant"),
		Key: map[string]*dynamodb.AttributeValue{
			"merchant_id": {
				S: aws.String(merchantID),
			},
		},
	})
	if err != nil {
		return nil, err
	}
	merchant := &generalmerchant.GeneralMerchant{}
	err = dynamodbattribute.UnmarshalMap(result.Item,merchant)
	return merchant, err
	Not implemented */
	return nil, nil
}

// DeleteGeneralMerchant : Deletes a general merchant with the given ID from the DB
func (c*RDSClient) DeleteGeneralMerchant(merchantID string) error {
	/*input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"merchant_id": {
				S: aws.String(merchantID),
			},
		},
		TableName: aws.String("Merchant"),
	}
	_, err := c.Client.DeleteItem(input)
	return err
	Not implemented */
	return nil

}

