package client

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"gitlab.com/otis-team/backend/dtypes/user/proto"
)


// CreateUser : Creates a new user in the db
func (c*RDSClient) CreateUser(user *user.User) (*user.User, error) {
	av, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return nil, err
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("User"),
	}
	_, err = c.Client.PutItem(input)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetAllUsers : Retrieves all users from the db
func (c *RDSClient) GetAllUsers() ([]*user.User, error) {
	result, err := c.Client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("User"),
	})
	if err != nil {
		return nil, err
	}
	users := make([]*user.User,0)
	err = dynamodbattribute.UnmarshalMap(result.Item,&users)
	return users, err
}


// GetUserByID : Retrieves a single user from the db
func (c*RDSClient) GetUserByID(userID string) (*user.User, error) {
	result, err := c.Client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("User"),
		Key: map[string]*dynamodb.AttributeValue{
			"user_id": {
				S: aws.String(userID),
			},
		},
	})
	if err != nil {
		return nil, err
	}
	users := user.User{}
	err = dynamodbattribute.UnmarshalMap(result.Item,&users)
	return &users, err
}

// DeleteUser : Deletes the user with the given ID from the DB
func (c*RDSClient) DeleteUser(userID string) error {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"user_id": {
				S: aws.String(userID),
			},
		},
		TableName: aws.String("User"),
	}
	_, err := c.Client.DeleteItem(input)
	return err
}