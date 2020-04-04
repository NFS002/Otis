package client

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"gitlab.com/otis-team/backend/db/model"
)



// DynamoClient : Struct to represent a connection to a DynamoDB instance 
type DynamoClient struct {
	Session *session.Session
	Client *dynamodb.DynamoDB
}

// Init : Function called on startup to initialize the dynamodb connection 
// AWS credentials and access keys must also be made available seperately.
// Please see the section on 'Specifying credentials' at https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html
func (c *DynamoClient) Init() error {
	var err error
	if c.Session == nil || c.Client == nil {
		err = c.newSession()
		if err == nil {
			c.newClient()
		}
	}
	return err
}

func (c *DynamoClient) newSession() error {
	sess, err := session.NewSession()
	c.Session = sess
	return err
}

func (c *DynamoClient) newClient() {
	c.Client = dynamodb.New(c.Session)
}