package client

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"gitlab.com/otis-team/backend/dtypes/partner-merchant/proto"
)

// CreatePartnerMerchant : Creates a new partner merchant in the db
func (c*RDSClient) CreatePartnerMerchant(merchant *partnermerchant.PartnerMerchant) (*partnermerchant.PartnerMerchant, error) {
	av, err := dynamodbattribute.MarshalMap(merchant)
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


// GetAllPartnerMerchants : Retrieves all partner merchants from the DB
func (c*RDSClient) GetAllPartnerMerchants() ([]*partnermerchant.PartnerMerchant, error) {
	result, err := c.Client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("Merchant"),
	})
	if err != nil {
		return nil, err
	}
	merchants := make([]*partnermerchant.PartnerMerchant,0)
	err = dynamodbattribute.UnmarshalMap(result.Item,&merchants)
	return merchants, err
}

// GetPartnerMerchantByID : Retrieves the Merchant from the DB with the given ID
func (c*RDSClient) GetPartnerMerchantByID(merchantID string) (*partnermerchant.PartnerMerchant, error) {
	result, err := c.Client.GetItem(&dynamodb.GetItemInput{
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
	merchant := &partnermerchant.PartnerMerchant{}
	err = dynamodbattribute.UnmarshalMap(result.Item,merchant)
	return merchant, err
}

// DeletePartnerMerchant : Deletes a merchant with the given ID from the DB
func (c*RDSClient) DeletePartnerMerchant(merchantID string) error {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"merchant_id": {
				S: aws.String(merchantID),
			},
		},
		TableName: aws.String("Merchant"),
	}
	_, err := c.Client.DeleteItem(input)
	return err
}