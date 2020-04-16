package client

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
)



// RDSClient : Struct to represent a connection to a DynamoDB instance
type RDSClient struct {
	Session *session.Session
	Client  *rds.RDS
}

// Init : Function called on startup to initialize the amazon RDS connection
// AWS credentials and access keys must also be made available seperately.
// Please see the section on 'Specifying credentials' at https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html
func (c *RDSClient) Init() error {
	var err error
	if c.Session == nil || c.Client == nil {
		err = c.newSession()
		if err == nil {
			c.newClient()
		}
	}
	return err
}

func (c *RDSClient) newSession() error {
	sess, err := session.NewSession()
	c.Session = sess
	return err
}

func (c *RDSClient) newClient() {
	c.Client = rds.New(c.Session)
}