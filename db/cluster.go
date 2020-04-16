package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rds"
	"gitlab.com/otis-team/backend/db/client"
	"os"
)

type rdsClient struct {
	*client.RDSClient
}

func getEnvOrPanic( key string ) string {
	v, exists := os.LookupEnv( key )
	if !exists {
		panic("Environment variabl not found: " + key)
	}
	return v
}


func (c* rdsClient) createCluster() (*rds.CreateDBClusterOutput, error) {
	input := &rds.CreateDBClusterInput{
		AvailabilityZones: []*string{
			aws.String(os.Getenv("DB_CLUSTER_ZONE")),
		},
		BackupRetentionPeriod:       	 aws.Int64(1),
		DBClusterIdentifier:         	 aws.String(getEnvOrPanic("DB_CLUSTER_NAME")),
		DBClusterParameterGroupName: 	 aws.String(getEnvOrPanic("DB_NAME")),
		DatabaseName:                	 aws.String(getEnvOrPanic("DB_CLUSTER_ZONE")),
		Engine:                      	 aws.String("aurora-postgresql"),
		EngineVersion:                	 aws.String("10.7"),
		EngineMode:						 aws.String("serverless"),
		MasterUserPassword:              aws.String(getEnvOrPanic("DB_MASTER_PASSWORD")),
		MasterUsername:                  aws.String(getEnvOrPanic("DB_MASTER_USERNAME")),
		StorageEncrypted:                aws.Bool(true),
		EnableHttpEndpoint:			     aws.Bool( true ),
		EnableIAMDatabaseAuthentication: aws.Bool( true ),

	}

	return c.Client.CreateDBCluster(input)
}

func main() {
	cli := client.RDSClient{}
	err := cli.Init()
	if err != nil {
		panic(err)
	}
	rdCli := rdsClient{ &cli }
	result, err := rdCli.createCluster()
	fmt.Println(result, err)
	if err != nil {
		panic(err)
	}
}
