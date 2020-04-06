package client

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"gitlab.com/otis-team/backend/db/model"
)


// CreateUser : Creates a new user in the db
func (c* DynamoClient) CreateUser(user *model.User) (*model.User, error) {
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
func (c *DynamoClient) GetAllUsers() (*model.Users, error) {
	result, err := c.Client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("User"),
	})
	if err != nil {
		return nil, err
	}
	users := model.Users{}
	err = dynamodbattribute.UnmarshalMap(result.Item,&users)
	return &users, err
}


// GetUser : Retrieves a single user from the db
func (c* DynamoClient) GetUserById(userID string) (*model.Users, error) {
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
	users := model.Users{}
	err = dynamodbattribute.UnmarshalMap(result.Item,&users)
	return &users, err
}


// UpdateUser : Updates a user with the given id
func (c* DynamoClient) UpdateUserById(userID string) (*model.User, error) {
	/* Deprecated, call CreateUser instead */
	return nil, nil
}

func (c* DynamoClient) DeleteUser(userID string) error {
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